package electrond

import (
	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
)

func Dial(addr string, api qrpc.API) (*Client, error) {
	sess, err := mux.DialWebsocket(addr)
	if err != nil {
		return nil, err
	}
	return (&Client{
		Client: &qrpc.Client{
			Session: sess,
			API:     api,
		},
	}).setup(), nil
}


type Client struct {
	*qrpc.Client
	App             *App
	AppDock         *Dock
	Clipboard       *Clipboard
	DesktopCapturer *DesktopCapturer
	Dialog          *Dialog
	GlobalShortcut  *GlobalShortcut
	PowerMonitor    *PowerMonitor
	Process         *Process
	Protocol        *Protocol
	Screen          *Screen
	Shell           *Shell
}

func (c *Client) setup() *Client {
	c.App = &App{Client: c.Client}
	c.AppDock = &Dock{Client: c.Client}
	c.Clipboard = &Clipboard{Client: c.Client}
	c.DesktopCapturer = &DesktopCapturer{Client: c.Client}
	c.Dialog = &Dialog{Client: c.Client}
	c.GlobalShortcut = &GlobalShortcut{Client: c.Client}
	c.PowerMonitor = &PowerMonitor{Client: c.Client}
	c.Process = &Process{Client: c.Client}
	c.Protocol = &Protocol{Client: c.Client}
	c.Screen = &Screen{Client: c.Client}
	c.Shell = &Shell{Client: c.Client}
	return c
}

type NativeImage struct{}
type CPUUsage struct {
	PercentCPUUsage      int `msgpack:"percentCPUUsage"`
	IdleWakeupsPerSecond int `msgpack:"idleWakeupsPerSecond"`
}
type Display struct {
	Id           int       `msgpack:"id"`
	Rotation     int       `msgpack:"rotation"`
	ScaleFactor  int       `msgpack:"scaleFactor"`
	TouchSupport string    `msgpack:"touchSupport"`
	Bounds       Rectangle `msgpack:"bounds"`
	Size         Size      `msgpack:"size"`
	WorkArea     Rectangle `msgpack:"workArea"`
	WorkAreaSize Size      `msgpack:"workAreaSize"`
}
type FileFilter struct {
	Name       string   `msgpack:"name"`
	Extensions []string `msgpack:"extensions"`
}
type Point struct {
	X int `msgpack:"x"`
	Y int `msgpack:"y"`
}
type ProcessMetric struct {
	Pid    int        `msgpack:"pid"`
	Type   string     `msgpack:"type"`
	Memory MemoryInfo `msgpack:"memory"`
	Cpu    CPUUsage   `msgpack:"cpu"`
}
type Rectangle struct {
	X      int `msgpack:"x"`
	Y      int `msgpack:"y"`
	Width  int `msgpack:"width"`
	Height int `msgpack:"height"`
}
type ShortcutDetails struct {
	Target         string `msgpack:"target"`
	Cwd            string `msgpack:"cwd,omitempty"`
	Args           string `msgpack:"args,omitempty"`
	Description    string `msgpack:"description,omitempty"`
	Icon           string `msgpack:"icon,omitempty"`
	IconIndex      int    `msgpack:"iconIndex,omitempty"`
	AppUserModelId string `msgpack:"appUserModelId,omitempty"`
}
type Size struct {
	Width  int `msgpack:"width"`
	Height int `msgpack:"height"`
}
type MemoryInfo struct {
	Pid                int `msgpack:"pid"`
	WorkingSetSize     int `msgpack:"workingSetSize"`
	PeakWorkingSetSize int `msgpack:"peakWorkingSetSize"`
	PrivateBytes       int `msgpack:"privateBytes"`
	SharedBytes        int `msgpack:"sharedBytes"`
}
type DesktopCapturerSource struct {
	Id         string      `msgpack:"id"`
	Name       string      `msgpack:"name"`
	Thumbnail  NativeImage `msgpack:"thumbnail"`
	Display_id string      `msgpack:"display_id"`
}
type UploadData struct {
	Bytes    []byte `msgpack:"bytes"`
	File     string `msgpack:"file"`
	BlobUUID string `msgpack:"blobUUID"`
}
type App struct {
	*qrpc.Client
}

func (c *App) Quit(ret interface{}) error {
	return c.Call("app.quit", nil, ret)
}
func (c *App) Focus(ret interface{}) error {
	return c.Call("app.focus", nil, ret)
}
func (c *App) Hide(ret interface{}) error {
	return c.Call("app.hide", nil, ret)
}
func (c *App) Show(ret interface{}) error {
	return c.Call("app.show", nil, ret)
}
func (c *App) GetAppPath(ret interface{}) error {
	return c.Call("app.getAppPath", nil, ret)
}
func (c *App) GetPath(params AppGetPathParams, ret interface{}) error {
	return c.Call("app.getPath", params, ret)
}

type AppGetPathParams struct {
	Name string `msgpack:"name"`
}

func (c *App) GetFileIcon(params AppGetFileIconParams, ret interface{}) error {
	return c.Call("app.getFileIcon", params, ret)
}

type AppGetFileIconParams struct {
	Path    string `msgpack:"path"`
	Options struct {
		Size string `msgpack:"size,omitempty"`
	} `msgpack:"options,omitempty"`
	Callback qrpc.ObjectHandle `msgpack:"callback"`
}

func (c *App) GetVersion(ret interface{}) error {
	return c.Call("app.getVersion", nil, ret)
}
func (c *App) GetLocale(ret interface{}) error {
	return c.Call("app.getLocale", nil, ret)
}
func (c *App) GetAppMetrics(ret interface{}) error {
	return c.Call("app.getAppMetrics", nil, ret)
}
func (c *App) SetBadgeCount(params AppSetBadgeCountParams, ret interface{}) error {
	return c.Call("app.setBadgeCount", params, ret)
}

type AppSetBadgeCountParams struct {
	Count int `msgpack:"count"`
}

func (c *App) GetBadgeCount(ret interface{}) error {
	return c.Call("app.getBadgeCount", nil, ret)
}

type Dock struct {
	*qrpc.Client
}

func (c *Dock) Bounce(params DockBounceParams, ret interface{}) error {
	return c.Call("app.dock.bounce", params, ret)
}

type DockBounceParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *Dock) CancelBounce(params DockCancelBounceParams, ret interface{}) error {
	return c.Call("app.dock.cancelBounce", params, ret)
}

type DockCancelBounceParams struct {
	Id int `msgpack:"id"`
}

func (c *Dock) DownloadFinished(params DockDownloadFinishedParams, ret interface{}) error {
	return c.Call("app.dock.downloadFinished", params, ret)
}

type DockDownloadFinishedParams struct {
	FilePath string `msgpack:"filePath"`
}

func (c *Dock) SetBadge(params DockSetBadgeParams, ret interface{}) error {
	return c.Call("app.dock.setBadge", params, ret)
}

type DockSetBadgeParams struct {
	Text string `msgpack:"text"`
}

func (c *Dock) GetBadge(ret interface{}) error {
	return c.Call("app.dock.getBadge", nil, ret)
}
func (c *Dock) Hide(ret interface{}) error {
	return c.Call("app.dock.hide", nil, ret)
}
func (c *Dock) Show(ret interface{}) error {
	return c.Call("app.dock.show", nil, ret)
}
func (c *Dock) IsVisible(ret interface{}) error {
	return c.Call("app.dock.isVisible", nil, ret)
}

type Clipboard struct {
	*qrpc.Client
}

func (c *Clipboard) ReadText(params ClipboardReadTextParams, ret interface{}) error {
	return c.Call("clipboard.readText", params, ret)
}

type ClipboardReadTextParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *Clipboard) WriteText(params ClipboardWriteTextParams, ret interface{}) error {
	return c.Call("clipboard.writeText", params, ret)
}

type ClipboardWriteTextParams struct {
	Text string `msgpack:"text"`
	Type string `msgpack:"type,omitempty"`
}

func (c *Clipboard) ReadHTML(params ClipboardReadHTMLParams, ret interface{}) error {
	return c.Call("clipboard.readHTML", params, ret)
}

type ClipboardReadHTMLParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *Clipboard) WriteHTML(params ClipboardWriteHTMLParams, ret interface{}) error {
	return c.Call("clipboard.writeHTML", params, ret)
}

type ClipboardWriteHTMLParams struct {
	Markup string `msgpack:"markup"`
	Type   string `msgpack:"type,omitempty"`
}

func (c *Clipboard) ReadImage(params ClipboardReadImageParams, ret interface{}) error {
	return c.Call("clipboard.readImage", params, ret)
}

type ClipboardReadImageParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *Clipboard) WriteImage(params ClipboardWriteImageParams, ret interface{}) error {
	return c.Call("clipboard.writeImage", params, ret)
}

type ClipboardWriteImageParams struct {
	Image NativeImage `msgpack:"image"`
	Type  string      `msgpack:"type,omitempty"`
}

func (c *Clipboard) ReadRTF(params ClipboardReadRTFParams, ret interface{}) error {
	return c.Call("clipboard.readRTF", params, ret)
}

type ClipboardReadRTFParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *Clipboard) WriteRTF(params ClipboardWriteRTFParams, ret interface{}) error {
	return c.Call("clipboard.writeRTF", params, ret)
}

type ClipboardWriteRTFParams struct {
	Text string `msgpack:"text"`
	Type string `msgpack:"type,omitempty"`
}

func (c *Clipboard) ReadBookmark(ret interface{}) error {
	return c.Call("clipboard.readBookmark", nil, ret)
}
func (c *Clipboard) WriteBookmark(params ClipboardWriteBookmarkParams, ret interface{}) error {
	return c.Call("clipboard.writeBookmark", params, ret)
}

type ClipboardWriteBookmarkParams struct {
	Title string `msgpack:"title"`
	Url   string `msgpack:"url"`
	Type  string `msgpack:"type,omitempty"`
}

func (c *Clipboard) Clear(params ClipboardClearParams, ret interface{}) error {
	return c.Call("clipboard.clear", params, ret)
}

type ClipboardClearParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *Clipboard) AvailableFormats(params ClipboardAvailableFormatsParams, ret interface{}) error {
	return c.Call("clipboard.availableFormats", params, ret)
}

type ClipboardAvailableFormatsParams struct {
	Type string `msgpack:"type,omitempty"`
}
type DesktopCapturer struct {
	*qrpc.Client
}

func (c *DesktopCapturer) GetSources(params DesktopCapturerGetSourcesParams, ret interface{}) error {
	return c.Call("desktopCapturer.getSources", params, ret)
}

type DesktopCapturerGetSourcesParams struct {
	Options struct {
		Types         []string `msgpack:"types,omitempty"`
		ThumbnailSize Size     `msgpack:"thumbnailSize,omitempty"`
	} `msgpack:"options"`
	Callback qrpc.ObjectHandle `msgpack:"callback"`
}
type Dialog struct {
	*qrpc.Client
}

func (c *Dialog) ShowOpenDialog(params DialogShowOpenDialogParams, ret interface{}) error {
	return c.Call("dialog.showOpenDialog", params, ret)
}

type DialogShowOpenDialogParams struct {
	Options struct {
		Title                   string       `msgpack:"title,omitempty"`
		DefaultPath             string       `msgpack:"defaultPath,omitempty"`
		ButtonLabel             string       `msgpack:"buttonLabel,omitempty"`
		Filters                 []FileFilter `msgpack:"filters,omitempty"`
		Properties              []string     `msgpack:"properties,omitempty"`
		Message                 string       `msgpack:"message,omitempty"`
		SecurityScopedBookmarks bool         `msgpack:"securityScopedBookmarks,omitempty"`
	} `msgpack:"options"`
}

func (c *Dialog) ShowSaveDialog(params DialogShowSaveDialogParams, ret interface{}) error {
	return c.Call("dialog.showSaveDialog", params, ret)
}

type DialogShowSaveDialogParams struct {
	Options struct {
		Title                   string       `msgpack:"title,omitempty"`
		DefaultPath             string       `msgpack:"defaultPath,omitempty"`
		ButtonLabel             string       `msgpack:"buttonLabel,omitempty"`
		Filters                 []FileFilter `msgpack:"filters,omitempty"`
		Message                 string       `msgpack:"message,omitempty"`
		NameFieldLabel          string       `msgpack:"nameFieldLabel,omitempty"`
		ShowsTagField           bool         `msgpack:"showsTagField,omitempty"`
		SecurityScopedBookmarks bool         `msgpack:"securityScopedBookmarks,omitempty"`
	} `msgpack:"options"`
}

func (c *Dialog) ShowMessageBox(params DialogShowMessageBoxParams, ret interface{}) error {
	return c.Call("dialog.showMessageBox", params, ret)
}

type DialogShowMessageBoxParams struct {
	Options struct {
		Type                string      `msgpack:"type,omitempty"`
		Buttons             []string    `msgpack:"buttons,omitempty"`
		DefaultId           int         `msgpack:"defaultId,omitempty"`
		Title               string      `msgpack:"title,omitempty"`
		Message             string      `msgpack:"message,omitempty"`
		Detail              string      `msgpack:"detail,omitempty"`
		CheckboxLabel       string      `msgpack:"checkboxLabel,omitempty"`
		CheckboxChecked     bool        `msgpack:"checkboxChecked,omitempty"`
		Icon                NativeImage `msgpack:"icon,omitempty"`
		CancelId            int         `msgpack:"cancelId,omitempty"`
		NoLink              bool        `msgpack:"noLink,omitempty"`
		NormalizeAccessKeys bool        `msgpack:"normalizeAccessKeys,omitempty"`
	} `msgpack:"options"`
}

func (c *Dialog) ShowErrorBox(params DialogShowErrorBoxParams, ret interface{}) error {
	return c.Call("dialog.showErrorBox", params, ret)
}

type DialogShowErrorBoxParams struct {
	Title   string `msgpack:"title"`
	Content string `msgpack:"content"`
}
type GlobalShortcut struct {
	*qrpc.Client
}

func (c *GlobalShortcut) Register(params GlobalShortcutRegisterParams, ret interface{}) error {
	return c.Call("globalShortcut.register", params, ret)
}

type GlobalShortcutRegisterParams struct {
	Accelerator string            `msgpack:"accelerator"`
	Callback    qrpc.ObjectHandle `msgpack:"callback"`
}

func (c *GlobalShortcut) IsRegistered(params GlobalShortcutIsRegisteredParams, ret interface{}) error {
	return c.Call("globalShortcut.isRegistered", params, ret)
}

type GlobalShortcutIsRegisteredParams struct {
	Accelerator string `msgpack:"accelerator"`
}

func (c *GlobalShortcut) Unregister(params GlobalShortcutUnregisterParams, ret interface{}) error {
	return c.Call("globalShortcut.unregister", params, ret)
}

type GlobalShortcutUnregisterParams struct {
	Accelerator string `msgpack:"accelerator"`
}
type PowerMonitor struct {
	*qrpc.Client
}
type Process struct {
	*qrpc.Client
}

func (c *Process) GetCPUUsage(ret interface{}) error {
	return c.Call("process.getCPUUsage", nil, ret)
}
func (c *Process) GetHeapStatistics(ret interface{}) error {
	return c.Call("process.getHeapStatistics", nil, ret)
}
func (c *Process) GetProcessMemoryInfo(ret interface{}) error {
	return c.Call("process.getProcessMemoryInfo", nil, ret)
}
func (c *Process) GetSystemMemoryInfo(ret interface{}) error {
	return c.Call("process.getSystemMemoryInfo", nil, ret)
}

type Protocol struct {
	*qrpc.Client
}

func (c *Protocol) RegisterStandardSchemes(params ProtocolRegisterStandardSchemesParams, ret interface{}) error {
	return c.Call("protocol.registerStandardSchemes", params, ret)
}

type ProtocolRegisterStandardSchemesParams struct {
	Schemes []string `msgpack:"schemes"`
	Options struct {
		Secure bool `msgpack:"secure,omitempty"`
	} `msgpack:"options,omitempty"`
}

func (c *Protocol) RegisterFileProtocol(params ProtocolRegisterFileProtocolParams, ret interface{}) error {
	return c.Call("protocol.registerFileProtocol", params, ret)
}

type ProtocolRegisterFileProtocolParams struct {
	Scheme  string            `msgpack:"scheme"`
	Handler qrpc.ObjectHandle `msgpack:"handler"`
}

func (c *Protocol) RegisterStringProtocol(params ProtocolRegisterStringProtocolParams, ret interface{}) error {
	return c.Call("protocol.registerStringProtocol", params, ret)
}

type ProtocolRegisterStringProtocolParams struct {
	Scheme  string            `msgpack:"scheme"`
	Handler qrpc.ObjectHandle `msgpack:"handler"`
}

func (c *Protocol) RegisterHttpProtocol(params ProtocolRegisterHttpProtocolParams, ret interface{}) error {
	return c.Call("protocol.registerHttpProtocol", params, ret)
}

type ProtocolRegisterHttpProtocolParams struct {
	Scheme  string            `msgpack:"scheme"`
	Handler qrpc.ObjectHandle `msgpack:"handler"`
}

func (c *Protocol) UnregisterProtocol(params ProtocolUnregisterProtocolParams, ret interface{}) error {
	return c.Call("protocol.unregisterProtocol", params, ret)
}

type ProtocolUnregisterProtocolParams struct {
	Scheme string `msgpack:"scheme"`
}

func (c *Protocol) IsProtocolHandled(params ProtocolIsProtocolHandledParams, ret interface{}) error {
	return c.Call("protocol.isProtocolHandled", params, ret)
}

type ProtocolIsProtocolHandledParams struct {
	Scheme string `msgpack:"scheme"`
}
type Screen struct {
	*qrpc.Client
}

func (c *Screen) GetCursorScreenPoint(ret interface{}) error {
	return c.Call("screen.getCursorScreenPoint", nil, ret)
}
func (c *Screen) GetPrimaryDisplay(ret interface{}) error {
	return c.Call("screen.getPrimaryDisplay", nil, ret)
}
func (c *Screen) GetAllDisplays(ret interface{}) error {
	return c.Call("screen.getAllDisplays", nil, ret)
}
func (c *Screen) GetDisplayNearestPoint(params ScreenGetDisplayNearestPointParams, ret interface{}) error {
	return c.Call("screen.getDisplayNearestPoint", params, ret)
}

type ScreenGetDisplayNearestPointParams struct {
	Point Point `msgpack:"point"`
}

func (c *Screen) GetDisplayMatching(params ScreenGetDisplayMatchingParams, ret interface{}) error {
	return c.Call("screen.getDisplayMatching", params, ret)
}

type ScreenGetDisplayMatchingParams struct {
	Rect Rectangle `msgpack:"rect"`
}

func (c *Screen) ScreenToDipPoint(params ScreenScreenToDipPointParams, ret interface{}) error {
	return c.Call("screen.screenToDipPoint", params, ret)
}

type ScreenScreenToDipPointParams struct {
	Point Point `msgpack:"point"`
}

func (c *Screen) DipToScreenPoint(params ScreenDipToScreenPointParams, ret interface{}) error {
	return c.Call("screen.dipToScreenPoint", params, ret)
}

type ScreenDipToScreenPointParams struct {
	Point Point `msgpack:"point"`
}
type Shell struct {
	*qrpc.Client
}

func (c *Shell) ShowItemInFolder(params ShellShowItemInFolderParams, ret interface{}) error {
	return c.Call("shell.showItemInFolder", params, ret)
}

type ShellShowItemInFolderParams struct {
	FullPath string `msgpack:"fullPath"`
}

func (c *Shell) OpenItem(params ShellOpenItemParams, ret interface{}) error {
	return c.Call("shell.openItem", params, ret)
}

type ShellOpenItemParams struct {
	FullPath string `msgpack:"fullPath"`
}

func (c *Shell) OpenExternal(params ShellOpenExternalParams, ret interface{}) error {
	return c.Call("shell.openExternal", params, ret)
}

type ShellOpenExternalParams struct {
	Url     string `msgpack:"url"`
	Options struct {
		Activate bool `msgpack:"activate,omitempty"`
	} `msgpack:"options,omitempty"`
	Callback qrpc.ObjectHandle `msgpack:"callback,omitempty"`
}

func (c *Shell) MoveItemToTrash(params ShellMoveItemToTrashParams, ret interface{}) error {
	return c.Call("shell.moveItemToTrash", params, ret)
}

type ShellMoveItemToTrashParams struct {
	FullPath string `msgpack:"fullPath"`
}

func (c *Shell) Beep(ret interface{}) error {
	return c.Call("shell.beep", nil, ret)
}
func (c *Shell) WriteShortcutLink(params ShellWriteShortcutLinkParams, ret interface{}) error {
	return c.Call("shell.writeShortcutLink", params, ret)
}

type ShellWriteShortcutLinkParams struct {
	ShortcutPath string          `msgpack:"shortcutPath"`
	Operation    string          `msgpack:"operation,omitempty"`
	Options      ShortcutDetails `msgpack:"options"`
}

func (c *Shell) ReadShortcutLink(params ShellReadShortcutLinkParams, ret interface{}) error {
	return c.Call("shell.readShortcutLink", params, ret)
}

type ShellReadShortcutLinkParams struct {
	ShortcutPath string `msgpack:"shortcutPath"`
}
