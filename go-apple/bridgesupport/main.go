package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const (
	encId          = "@"
	encClass       = "#"
	encSelector    = ":"
	encChar        = "c"
	encUChar       = "C"
	encShort       = "s"
	encUShort      = "S"
	encInt         = "i"
	encUInt        = "I"
	encLong        = "l"
	encULong       = "L"
	encLongLong    = "q"
	encULongLong   = "Q"
	encFloat       = "f"
	encDouble      = "d"
	encDFLD        = "b"
	encBool        = "B"
	encVoid        = "v"
	encUndef       = "?"
	encPtr         = "^"
	encCharPtr     = "*"
	encAtom        = "%"
	encArrayBegin  = "["
	encArrayEnd    = "]"
	encUnionBegin  = "("
	encUnionEnd    = ")"
	encStructBegin = "{"
	encStructEnd   = "}"
	encVector      = "!"
	encConst       = "r"
	encNameQuote   = `"`
)

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func extractCompound(data, open, close []byte, atEOF bool) (advance int, token []byte, err error) {
	start := bytes.Index(data, open)
	if start == -1 {
		return 0, nil, io.EOF
	}
	inner := 0
	for i := start; i < len(data); i++ {
		d := data[i : i+1]
		switch {
		case bytes.HasPrefix(d, open):
			inner++
		case bytes.HasPrefix(d, close):
			if inner > 1 {
				inner--
			} else if inner == 1 {
				return i + 1, data[start : i+2], nil
			}
		}
	}
	if atEOF && len(data) > start {
		return len(data[start:]), data[start:], nil
	}
	return start, nil, nil
}

func split(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 { // atEOF &&
		return 0, nil, nil
	}
	switch string(data[0]) {
	case encStructBegin:
		return extractCompound(data, []byte(encStructBegin), []byte(encStructEnd), atEOF)
	case encPtr:
		width, _, err := split(data[1:], atEOF)
		return width + 1, data[0 : width+1], err
	case encDFLD:
		width := 1
		for _, b := range data[1:] {
			if !isNumber(string(b)) {
				break
			}
			width++
		}
		return width, data[0:width], nil
	case encArrayBegin:
		return extractCompound(data, []byte(encArrayBegin), []byte(encArrayEnd), atEOF)
	case encUnionBegin:
		return extractCompound(data, []byte(encUnionBegin), []byte(encUnionEnd), atEOF)
	case encNameQuote:
		closer := bytes.Index(data[1:], []byte(encNameQuote))
		if closer == -1 {
			return 0, nil, io.EOF
		}
		start := closer + 2
		width, _, err := split(data[start:], atEOF)
		return start + width, data[0 : start+width], err
	default:
		return 1, data[0:1], nil
	}

}

func NewParser(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(split)
	buf := make([]byte, 64)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	return scanner
}

func parseType(t string) TypeInfo {
	if len(t) == 0 {
		return TypeInfo{
			Error: "empty type",
		}
	}
	switch t[0] {
	case '@', '#', ':', 'c', 'C', 's', 'S', 'i', 'I', 'l', 'L', 'q', 'Q', 'f', 'd', 'B', 'v', '?', '*', '%', '!', 'r':
		return TypeInfo{
			Code: t,
		}
	// case 'T':
	// 	return TypeInfo{
	// 		Code: t,
	// 	}
	case 'b':
		// TODO
		return TypeInfo{
			Code: t,
		}
	case '^':
		if len(t) == 1 {
			return TypeInfo{
				Code: t,
			}
		}
		typ := parseType(t[1 : len(t)-1])
		return TypeInfo{
			Code: t,
			Type: &typ,
		}
	case '[':
		// TODO
		return TypeInfo{
			Code: t,
		}
	case '(':
		// TODO
		return TypeInfo{
			Code: t,
		}
	case '{':
		closer := strings.LastIndex(t, "}")
		if closer == -1 {
			return TypeInfo{
				Code:  t,
				Error: "invalid struct",
			}
		}
		inner := t[1:closer]
		if strings.Index(inner, "=") == -1 {
			return TypeInfo{
				Code: t,
				Name: inner,
			}
		}
		parts := strings.SplitN(inner, "=", 2)
		var types []*TypeInfo
		var names []string
		if len(parts) > 1 {
			scanner := NewParser(strings.NewReader(parts[1]))
			for scanner.Scan() {
				token := scanner.Text()
				if strings.Index(token, `"`) == -1 {
					ti := parseType(token)
					types = append(types, &ti)
				} else {
					p := strings.Split(token, `"`)
					if len(p) != 3 {
						continue
					}
					ti := parseType(p[2])
					types = append(types, &ti)
					names = append(names, p[1])
				}
			}
		}
		return TypeInfo{
			Code:  t,
			Name:  parts[0],
			Types: types,
			Names: names,
		}
	default:
		return TypeInfo{
			Code:  t,
			Error: "cannot parse",
		}
	}

}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

type TypeInfo struct {
	Code  string
	Type  *TypeInfo
	Names []string
	Types []*TypeInfo
	Bits  int
	Name  string
	Error string
}

func (ti *TypeInfo) UnmarshalXMLAttr(attr xml.Attr) error {
	*ti = parseType(attr.Value)
	return nil
}

func main() {

	// f, err := os.Open("../appkit/AppKit.bridgesupport")
	f, err := os.Open("../foundation/Foundation.bridgesupport")
	fatal(err)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	fatal(err)
	var sigs Signatures
	fatal(xml.Unmarshal(b, &sigs))
	spew.Dump(sigs.Classes)
}

type Signatures struct {
	XMLName           xml.Name           `xml:"signatures"`
	Version           string             `xml:"version,attr"`
	Opaques           []Opaque           `xml:"opaque"`
	Constants         []Constant         `xml:"constant"`
	Enums             []Enum             `xml:"enum"`
	Functions         []Function         `xml:"function"`
	Classes           []Class            `xml:"class"`
	Structs           []Struct           `xml:"struct"`
	InformalProtocols []InformalProtocol `xml:"informal_protocol"`
}

type Struct struct {
	XMLName xml.Name  `xml:"struct"`
	Name    string    `xml:"name,attr"`
	Type    TypeInfo  `xml:"type,attr"`
	Type64  *TypeInfo `xml:"type64,attr"`
}

type Opaque struct {
	XMLName xml.Name  `xml:"opaque"`
	Name    string    `xml:"name,attr"`
	Type    TypeInfo  `xml:"type,attr"`
	Type64  *TypeInfo `xml:"type64,attr"`
}

type Constant struct {
	XMLName xml.Name  `xml:"constant"`
	Name    string    `xml:"name,attr"`
	Type    TypeInfo  `xml:"type,attr"`
	Type64  *TypeInfo `xml:"type64,attr"`
}

type InformalProtocol struct {
	XMLName xml.Name `xml:"informal_protocol"`
	Name    string   `xml:"name,attr"`
	Methods []Method `xml:"method"`
}

type Enum struct {
	XMLName xml.Name `xml:"enum"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
	Value64 string   `xml:"value64,attr"`
	Ignore  bool     `xml:"ignore,attr"`
}

type Class struct {
	XMLName xml.Name `xml:"class"`
	Name    string   `xml:"name,attr"`
	Methods []Method `xml:"method"`
}

type Method struct {
	XMLName     xml.Name `xml:"method"`
	Selector    string   `xml:"selector,attr"`
	ClassMethod bool     `xml:"class_method,attr"`
	Variadic    bool     `xml:"variadic,attr"`
	Ignore      bool     `xml:"ignore,attr"`
	Suggestion  string   `xml:"suggestion,attr"`
	Sentinal    int      `xml:"sentinal,attr"`
	Type        string   `xml:"type,attr"`
	Type64      string   `xml:"type64,attr"`
	Args        []Arg    `xml:"arg"`
	RetVal      RetVal   `xml:"retval"`
}

type Function struct {
	XMLName  xml.Name `xml:"function"`
	Name     string   `xml:"name,attr"`
	Variadic bool     `xml:"variadic,attr"`
	Inline   bool     `xml:"inline,attr"`
	Sentinal int      `xml:"sentinal,attr"`
	Args     []Arg    `xml:"arg"`
	RetVal   RetVal   `xml:"retval"`
}

type Arg struct {
	XMLName      xml.Name  `xml:"arg"`
	Index        *int      `xml:"index,attr"`
	TypeModifier string    `xml:"type_modifier,attr"`
	Type         TypeInfo  `xml:"type,attr"`
	Type64       *TypeInfo `xml:"type64,attr"`
}

type RetVal struct {
	XMLName xml.Name  `xml:"retval"`
	Type    TypeInfo  `xml:"type,attr"`
	Type64  *TypeInfo `xml:"type64,attr"`
}
