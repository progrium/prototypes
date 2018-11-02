package foundation

type NSAffineTransformStruct struct {
	M11 float32
	M12 float32
	M21 float32
	M22 float32
	TX  float32
	TY  float32
}
type NSDecimal struct {
	_exponent   int32
	_length     uint32
	_isNegative uint32
	_isCompact  uint32
	_reserved   uint32
}
type NSEdgeInsets struct {
	Top    float32
	Left   float32
	Bottom float32
	Right  float32
}
type NSFastEnumerationState struct {
	State        uint32
	ItemsPtr     uintptr
	MutationsPtr uintptr
	Extra        slice
}
type NSHashEnumerator struct {
	_pi uint32
	_si uint32
	_bs uintptr
}
type NSHashTableCallBacks struct {
	Hash     uintptr
	IsEqual  uintptr
	Retain   uintptr
	Release  uintptr
	Describe uintptr
}
type NSMapEnumerator struct {
	_pi uint32
	_si uint32
	_bs uintptr
}
type NSMapTableKeyCallBacks struct {
	Hash     uintptr
	IsEqual  uintptr
	Retain   uintptr
	Release  uintptr
	Describe uintptr
}
type NSMapTableValueCallBacks struct {
	Retain   uintptr
	Release  uintptr
	Describe uintptr
}
type NSOperatingSystemVersion struct {
	MajorVersion int32
	MinorVersion int32
	PatchVersion int32
}
type NSPoint struct {
	X float32
	Y float32
}
type NSPointArray struct {
	X float32
	Y float32
}
type NSPointPointer struct {
	X float32
	Y float32
}
type NSRange struct {
	Location uint32
	Length   uint32
}
type NSRangePointer struct {
	Location uint32
	Length   uint32
}
type NSRect struct {
	Origin NSPoint
	Size   NSSize
}
type NSRectArray struct {
	Origin NSPoint
	Size   NSSize
}
type NSRectPointer struct {
	Origin NSPoint
	Size   NSSize
}
type NSSize struct {
	Width  float32
	Height float32
}
type NSSizeArray struct {
	Width  float32
	Height float32
}
type NSSizePointer struct {
	Width  float32
	Height float32
}
type NSSwappedDouble struct {
	V uint64
}
type NSSwappedFloat struct {
	V uint32
}
