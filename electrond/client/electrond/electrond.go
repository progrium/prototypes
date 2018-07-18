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
	App             *AppModule
	AppDock         *DockModule
	Clipboard       *ClipboardModule
	DesktopCapturer *DesktopCapturerModule
	Dialog          *DialogModule
	GlobalShortcut  *GlobalShortcutModule
	PowerMonitor    *PowerMonitorModule
	Process         *ProcessModule
	Protocol        *ProtocolModule
	Screen          *ScreenModule
	Shell           *ShellModule
	NativeImage     *NativeImageModule
	Menu            *MenuModule
	MenuItem        *MenuItemModule
	Tray            *TrayModule
	Notification    *NotificationModule
}

func (c *Client) setup() *Client {
	c.App = &AppModule{Client: c.Client}
	c.AppDock = &DockModule{Client: c.Client}
	c.Clipboard = &ClipboardModule{Client: c.Client}
	c.DesktopCapturer = &DesktopCapturerModule{Client: c.Client}
	c.Dialog = &DialogModule{Client: c.Client}
	c.GlobalShortcut = &GlobalShortcutModule{Client: c.Client}
	c.PowerMonitor = &PowerMonitorModule{Client: c.Client}
	c.Process = &ProcessModule{Client: c.Client}
	c.Protocol = &ProtocolModule{Client: c.Client}
	c.Screen = &ScreenModule{Client: c.Client}
	c.Shell = &ShellModule{Client: c.Client}
	c.NativeImage = &NativeImageModule{Client: c.Client}
	c.Menu = &MenuModule{Client: c.Client}
	c.MenuItem = &MenuItemModule{Client: c.Client}
	c.Tray = &TrayModule{Client: c.Client}
	c.Notification = &NotificationModule{Client: c.Client}
	return c
}

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
	Id         string            `msgpack:"id"`
	Name       string            `msgpack:"name"`
	Thumbnail  qrpc.ObjectHandle `msgpack:"thumbnail"`
	Display_id string            `msgpack:"display_id"`
}
type UploadData struct {
	Bytes    []byte `msgpack:"bytes"`
	File     string `msgpack:"file"`
	BlobUUID string `msgpack:"blobUUID"`
}
type NotificationAction struct {
	Type string `msgpack:"type"`
	Text string `msgpack:"text,omitempty"`
}
type MenuItemConstructorOptions struct {
	Click       *qrpc.ObjectHandle          `msgpack:"click,omitempty"`
	Role        string                      `msgpack:"role,omitempty"`
	Type        string                      `msgpack:"type,omitempty"`
	Label       string                      `msgpack:"label,omitempty"`
	Sublabel    string                      `msgpack:"sublabel,omitempty"`
	Accelerator string                      `msgpack:"accelerator,omitempty"`
	Icon        *qrpc.ObjectHandle          `msgpack:"icon,omitempty"`
	Enabled     bool                        `msgpack:"enabled,omitempty"`
	Visible     bool                        `msgpack:"visible,omitempty"`
	Checked     bool                        `msgpack:"checked,omitempty"`
	Submenu     *MenuItemConstructorOptions `msgpack:"submenu,omitempty"`
	Id          string                      `msgpack:"id,omitempty"`
	Position    string                      `msgpack:"position,omitempty"`
}
type AppModule struct {
	*qrpc.Client
}

func (c *AppModule) Quit(ret interface{}) error {
	return c.Call("app.quit", nil, ret)
}
func (c *AppModule) Focus(ret interface{}) error {
	return c.Call("app.focus", nil, ret)
}
func (c *AppModule) Hide(ret interface{}) error {
	return c.Call("app.hide", nil, ret)
}
func (c *AppModule) Show(ret interface{}) error {
	return c.Call("app.show", nil, ret)
}
func (c *AppModule) GetAppPath(ret interface{}) error {
	return c.Call("app.getAppPath", nil, ret)
}
func (c *AppModule) GetPath(params AppGetPathParams, ret interface{}) error {
	return c.Call("app.getPath", params, ret)
}

type AppGetPathParams struct {
	Name string `msgpack:"name"`
}

func (c *AppModule) GetFileIcon(params AppGetFileIconParams, ret interface{}) error {
	return c.Call("app.getFileIcon", params, ret)
}

type AppGetFileIconParamsOptions struct {
	Size string `msgpack:"size,omitempty"`
}
type AppGetFileIconParams struct {
	Path     string                      `msgpack:"path"`
	Options  AppGetFileIconParamsOptions `msgpack:"options,omitempty"`
	Callback qrpc.ObjectHandle           `msgpack:"callback"`
}

func (c *AppModule) GetVersion(ret interface{}) error {
	return c.Call("app.getVersion", nil, ret)
}
func (c *AppModule) GetLocale(ret interface{}) error {
	return c.Call("app.getLocale", nil, ret)
}
func (c *AppModule) GetAppMetrics(ret interface{}) error {
	return c.Call("app.getAppMetrics", nil, ret)
}
func (c *AppModule) SetBadgeCount(params AppSetBadgeCountParams, ret interface{}) error {
	return c.Call("app.setBadgeCount", params, ret)
}

type AppSetBadgeCountParams struct {
	Count int `msgpack:"count"`
}

func (c *AppModule) GetBadgeCount(ret interface{}) error {
	return c.Call("app.getBadgeCount", nil, ret)
}

type DockModule struct {
	*qrpc.Client
}

func (c *DockModule) Bounce(params DockBounceParams, ret interface{}) error {
	return c.Call("app.dock.bounce", params, ret)
}

type DockBounceParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *DockModule) CancelBounce(params DockCancelBounceParams, ret interface{}) error {
	return c.Call("app.dock.cancelBounce", params, ret)
}

type DockCancelBounceParams struct {
	Id int `msgpack:"id"`
}

func (c *DockModule) DownloadFinished(params DockDownloadFinishedParams, ret interface{}) error {
	return c.Call("app.dock.downloadFinished", params, ret)
}

type DockDownloadFinishedParams struct {
	FilePath string `msgpack:"filePath"`
}

func (c *DockModule) SetBadge(params DockSetBadgeParams, ret interface{}) error {
	return c.Call("app.dock.setBadge", params, ret)
}

type DockSetBadgeParams struct {
	Text string `msgpack:"text"`
}

func (c *DockModule) GetBadge(ret interface{}) error {
	return c.Call("app.dock.getBadge", nil, ret)
}
func (c *DockModule) Hide(ret interface{}) error {
	return c.Call("app.dock.hide", nil, ret)
}
func (c *DockModule) Show(ret interface{}) error {
	return c.Call("app.dock.show", nil, ret)
}
func (c *DockModule) IsVisible(ret interface{}) error {
	return c.Call("app.dock.isVisible", nil, ret)
}

type ClipboardModule struct {
	*qrpc.Client
}

func (c *ClipboardModule) ReadText(params ClipboardReadTextParams, ret interface{}) error {
	return c.Call("clipboard.readText", params, ret)
}

type ClipboardReadTextParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) WriteText(params ClipboardWriteTextParams, ret interface{}) error {
	return c.Call("clipboard.writeText", params, ret)
}

type ClipboardWriteTextParams struct {
	Text string `msgpack:"text"`
	Type string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) ReadHTML(params ClipboardReadHTMLParams, ret interface{}) error {
	return c.Call("clipboard.readHTML", params, ret)
}

type ClipboardReadHTMLParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) WriteHTML(params ClipboardWriteHTMLParams, ret interface{}) error {
	return c.Call("clipboard.writeHTML", params, ret)
}

type ClipboardWriteHTMLParams struct {
	Markup string `msgpack:"markup"`
	Type   string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) ReadImage(params ClipboardReadImageParams, ret interface{}) error {
	return c.Call("clipboard.readImage", params, ret)
}

type ClipboardReadImageParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) WriteImage(params ClipboardWriteImageParams, ret interface{}) error {
	return c.Call("clipboard.writeImage", params, ret)
}

type ClipboardWriteImageParams struct {
	Image qrpc.ObjectHandle `msgpack:"image"`
	Type  string            `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) ReadRTF(params ClipboardReadRTFParams, ret interface{}) error {
	return c.Call("clipboard.readRTF", params, ret)
}

type ClipboardReadRTFParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) WriteRTF(params ClipboardWriteRTFParams, ret interface{}) error {
	return c.Call("clipboard.writeRTF", params, ret)
}

type ClipboardWriteRTFParams struct {
	Text string `msgpack:"text"`
	Type string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) ReadBookmark(ret interface{}) error {
	return c.Call("clipboard.readBookmark", nil, ret)
}
func (c *ClipboardModule) WriteBookmark(params ClipboardWriteBookmarkParams, ret interface{}) error {
	return c.Call("clipboard.writeBookmark", params, ret)
}

type ClipboardWriteBookmarkParams struct {
	Title string `msgpack:"title"`
	Url   string `msgpack:"url"`
	Type  string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) Clear(params ClipboardClearParams, ret interface{}) error {
	return c.Call("clipboard.clear", params, ret)
}

type ClipboardClearParams struct {
	Type string `msgpack:"type,omitempty"`
}

func (c *ClipboardModule) AvailableFormats(params ClipboardAvailableFormatsParams, ret interface{}) error {
	return c.Call("clipboard.availableFormats", params, ret)
}

type ClipboardAvailableFormatsParams struct {
	Type string `msgpack:"type,omitempty"`
}
type DesktopCapturerModule struct {
	*qrpc.Client
}

func (c *DesktopCapturerModule) GetSources(params DesktopCapturerGetSourcesParams, ret interface{}) error {
	return c.Call("desktopCapturer.getSources", params, ret)
}

type DesktopCapturerGetSourcesParamsOptions struct {
	Types         []string `msgpack:"types,omitempty"`
	ThumbnailSize Size     `msgpack:"thumbnailSize,omitempty"`
}
type DesktopCapturerGetSourcesParams struct {
	Options  DesktopCapturerGetSourcesParamsOptions `msgpack:"options"`
	Callback qrpc.ObjectHandle                      `msgpack:"callback"`
}
type DialogModule struct {
	*qrpc.Client
}

func (c *DialogModule) ShowOpenDialog(params DialogShowOpenDialogParams, ret interface{}) error {
	return c.Call("dialog.showOpenDialog", params, ret)
}

type DialogShowOpenDialogParamsOptions struct {
	Title                   string       `msgpack:"title,omitempty"`
	DefaultPath             string       `msgpack:"defaultPath,omitempty"`
	ButtonLabel             string       `msgpack:"buttonLabel,omitempty"`
	Filters                 []FileFilter `msgpack:"filters,omitempty"`
	Properties              []string     `msgpack:"properties,omitempty"`
	Message                 string       `msgpack:"message,omitempty"`
	SecurityScopedBookmarks bool         `msgpack:"securityScopedBookmarks,omitempty"`
}
type DialogShowOpenDialogParams struct {
	Options DialogShowOpenDialogParamsOptions `msgpack:"options"`
}

func (c *DialogModule) ShowSaveDialog(params DialogShowSaveDialogParams, ret interface{}) error {
	return c.Call("dialog.showSaveDialog", params, ret)
}

type DialogShowSaveDialogParamsOptions struct {
	Title                   string       `msgpack:"title,omitempty"`
	DefaultPath             string       `msgpack:"defaultPath,omitempty"`
	ButtonLabel             string       `msgpack:"buttonLabel,omitempty"`
	Filters                 []FileFilter `msgpack:"filters,omitempty"`
	Message                 string       `msgpack:"message,omitempty"`
	NameFieldLabel          string       `msgpack:"nameFieldLabel,omitempty"`
	ShowsTagField           bool         `msgpack:"showsTagField,omitempty"`
	SecurityScopedBookmarks bool         `msgpack:"securityScopedBookmarks,omitempty"`
}
type DialogShowSaveDialogParams struct {
	Options DialogShowSaveDialogParamsOptions `msgpack:"options"`
}

func (c *DialogModule) ShowMessageBox(params DialogShowMessageBoxParams, ret interface{}) error {
	return c.Call("dialog.showMessageBox", params, ret)
}

type DialogShowMessageBoxParamsOptions struct {
	Type                string             `msgpack:"type,omitempty"`
	Buttons             []string           `msgpack:"buttons,omitempty"`
	DefaultId           int                `msgpack:"defaultId,omitempty"`
	Title               string             `msgpack:"title,omitempty"`
	Message             string             `msgpack:"message,omitempty"`
	Detail              string             `msgpack:"detail,omitempty"`
	CheckboxLabel       string             `msgpack:"checkboxLabel,omitempty"`
	CheckboxChecked     bool               `msgpack:"checkboxChecked,omitempty"`
	Icon                *qrpc.ObjectHandle `msgpack:"icon,omitempty"`
	CancelId            int                `msgpack:"cancelId,omitempty"`
	NoLink              bool               `msgpack:"noLink,omitempty"`
	NormalizeAccessKeys bool               `msgpack:"normalizeAccessKeys,omitempty"`
}
type DialogShowMessageBoxParams struct {
	Options DialogShowMessageBoxParamsOptions `msgpack:"options"`
}

func (c *DialogModule) ShowErrorBox(params DialogShowErrorBoxParams, ret interface{}) error {
	return c.Call("dialog.showErrorBox", params, ret)
}

type DialogShowErrorBoxParams struct {
	Title   string `msgpack:"title"`
	Content string `msgpack:"content"`
}
type GlobalShortcutModule struct {
	*qrpc.Client
}

func (c *GlobalShortcutModule) Register(params GlobalShortcutRegisterParams, ret interface{}) error {
	return c.Call("globalShortcut.register", params, ret)
}

type GlobalShortcutRegisterParams struct {
	Accelerator string            `msgpack:"accelerator"`
	Callback    qrpc.ObjectHandle `msgpack:"callback"`
}

func (c *GlobalShortcutModule) IsRegistered(params GlobalShortcutIsRegisteredParams, ret interface{}) error {
	return c.Call("globalShortcut.isRegistered", params, ret)
}

type GlobalShortcutIsRegisteredParams struct {
	Accelerator string `msgpack:"accelerator"`
}

func (c *GlobalShortcutModule) Unregister(params GlobalShortcutUnregisterParams, ret interface{}) error {
	return c.Call("globalShortcut.unregister", params, ret)
}

type GlobalShortcutUnregisterParams struct {
	Accelerator string `msgpack:"accelerator"`
}
type PowerMonitorModule struct {
	*qrpc.Client
}
type ProcessModule struct {
	*qrpc.Client
}

func (c *ProcessModule) GetCPUUsage(ret interface{}) error {
	return c.Call("process.getCPUUsage", nil, ret)
}
func (c *ProcessModule) GetHeapStatistics(ret interface{}) error {
	return c.Call("process.getHeapStatistics", nil, ret)
}
func (c *ProcessModule) GetProcessMemoryInfo(ret interface{}) error {
	return c.Call("process.getProcessMemoryInfo", nil, ret)
}
func (c *ProcessModule) GetSystemMemoryInfo(ret interface{}) error {
	return c.Call("process.getSystemMemoryInfo", nil, ret)
}

type ProtocolModule struct {
	*qrpc.Client
}

func (c *ProtocolModule) RegisterStandardSchemes(params ProtocolRegisterStandardSchemesParams, ret interface{}) error {
	return c.Call("protocol.registerStandardSchemes", params, ret)
}

type ProtocolRegisterStandardSchemesParamsOptions struct {
	Secure bool `msgpack:"secure,omitempty"`
}
type ProtocolRegisterStandardSchemesParams struct {
	Schemes []string                                     `msgpack:"schemes"`
	Options ProtocolRegisterStandardSchemesParamsOptions `msgpack:"options,omitempty"`
}

func (c *ProtocolModule) RegisterFileProtocol(params ProtocolRegisterFileProtocolParams, ret interface{}) error {
	return c.Call("protocol.registerFileProtocol", params, ret)
}

type ProtocolRegisterFileProtocolParams struct {
	Scheme  string            `msgpack:"scheme"`
	Handler qrpc.ObjectHandle `msgpack:"handler"`
}

func (c *ProtocolModule) RegisterStringProtocol(params ProtocolRegisterStringProtocolParams, ret interface{}) error {
	return c.Call("protocol.registerStringProtocol", params, ret)
}

type ProtocolRegisterStringProtocolParams struct {
	Scheme  string            `msgpack:"scheme"`
	Handler qrpc.ObjectHandle `msgpack:"handler"`
}

func (c *ProtocolModule) RegisterHttpProtocol(params ProtocolRegisterHttpProtocolParams, ret interface{}) error {
	return c.Call("protocol.registerHttpProtocol", params, ret)
}

type ProtocolRegisterHttpProtocolParams struct {
	Scheme  string            `msgpack:"scheme"`
	Handler qrpc.ObjectHandle `msgpack:"handler"`
}

func (c *ProtocolModule) UnregisterProtocol(params ProtocolUnregisterProtocolParams, ret interface{}) error {
	return c.Call("protocol.unregisterProtocol", params, ret)
}

type ProtocolUnregisterProtocolParams struct {
	Scheme string `msgpack:"scheme"`
}

func (c *ProtocolModule) IsProtocolHandled(params ProtocolIsProtocolHandledParams, ret interface{}) error {
	return c.Call("protocol.isProtocolHandled", params, ret)
}

type ProtocolIsProtocolHandledParams struct {
	Scheme string `msgpack:"scheme"`
}
type ScreenModule struct {
	*qrpc.Client
}

func (c *ScreenModule) GetCursorScreenPoint(ret interface{}) error {
	return c.Call("screen.getCursorScreenPoint", nil, ret)
}
func (c *ScreenModule) GetPrimaryDisplay(ret interface{}) error {
	return c.Call("screen.getPrimaryDisplay", nil, ret)
}
func (c *ScreenModule) GetAllDisplays(ret interface{}) error {
	return c.Call("screen.getAllDisplays", nil, ret)
}
func (c *ScreenModule) GetDisplayNearestPoint(params ScreenGetDisplayNearestPointParams, ret interface{}) error {
	return c.Call("screen.getDisplayNearestPoint", params, ret)
}

type ScreenGetDisplayNearestPointParams struct {
	Point Point `msgpack:"point"`
}

func (c *ScreenModule) GetDisplayMatching(params ScreenGetDisplayMatchingParams, ret interface{}) error {
	return c.Call("screen.getDisplayMatching", params, ret)
}

type ScreenGetDisplayMatchingParams struct {
	Rect Rectangle `msgpack:"rect"`
}

func (c *ScreenModule) ScreenToDipPoint(params ScreenScreenToDipPointParams, ret interface{}) error {
	return c.Call("screen.screenToDipPoint", params, ret)
}

type ScreenScreenToDipPointParams struct {
	Point Point `msgpack:"point"`
}

func (c *ScreenModule) DipToScreenPoint(params ScreenDipToScreenPointParams, ret interface{}) error {
	return c.Call("screen.dipToScreenPoint", params, ret)
}

type ScreenDipToScreenPointParams struct {
	Point Point `msgpack:"point"`
}
type ShellModule struct {
	*qrpc.Client
}

func (c *ShellModule) ShowItemInFolder(params ShellShowItemInFolderParams, ret interface{}) error {
	return c.Call("shell.showItemInFolder", params, ret)
}

type ShellShowItemInFolderParams struct {
	FullPath string `msgpack:"fullPath"`
}

func (c *ShellModule) OpenItem(params ShellOpenItemParams, ret interface{}) error {
	return c.Call("shell.openItem", params, ret)
}

type ShellOpenItemParams struct {
	FullPath string `msgpack:"fullPath"`
}

func (c *ShellModule) OpenExternal(params ShellOpenExternalParams, ret interface{}) error {
	return c.Call("shell.openExternal", params, ret)
}

type ShellOpenExternalParamsOptions struct {
	Activate bool `msgpack:"activate,omitempty"`
}
type ShellOpenExternalParams struct {
	Url      string                         `msgpack:"url"`
	Options  ShellOpenExternalParamsOptions `msgpack:"options,omitempty"`
	Callback qrpc.ObjectHandle              `msgpack:"callback,omitempty"`
}

func (c *ShellModule) MoveItemToTrash(params ShellMoveItemToTrashParams, ret interface{}) error {
	return c.Call("shell.moveItemToTrash", params, ret)
}

type ShellMoveItemToTrashParams struct {
	FullPath string `msgpack:"fullPath"`
}

func (c *ShellModule) Beep(ret interface{}) error {
	return c.Call("shell.beep", nil, ret)
}
func (c *ShellModule) WriteShortcutLink(params ShellWriteShortcutLinkParams, ret interface{}) error {
	return c.Call("shell.writeShortcutLink", params, ret)
}

type ShellWriteShortcutLinkParams struct {
	ShortcutPath string          `msgpack:"shortcutPath"`
	Operation    string          `msgpack:"operation,omitempty"`
	Options      ShortcutDetails `msgpack:"options"`
}

func (c *ShellModule) ReadShortcutLink(params ShellReadShortcutLinkParams, ret interface{}) error {
	return c.Call("shell.readShortcutLink", params, ret)
}

type ShellReadShortcutLinkParams struct {
	ShortcutPath string `msgpack:"shortcutPath"`
}
type NativeImageModule struct {
	*qrpc.Client
}

func (c *NativeImageModule) CreateEmpty(ret interface{}) error {
	return c.Call("nativeImage.createEmpty", nil, ret)
}
func (c *NativeImageModule) CreateFromPath(params NativeImageCreateFromPathParams, ret interface{}) error {
	return c.Call("nativeImage.createFromPath", params, ret)
}

type NativeImageCreateFromPathParams struct {
	Path string `msgpack:"path"`
}

func (c *NativeImageModule) CreateFromBuffer(params NativeImageCreateFromBufferParams, ret interface{}) error {
	return c.Call("nativeImage.createFromBuffer", params, ret)
}

type NativeImageCreateFromBufferParamsOptions struct {
	Width       int     `msgpack:"width,omitempty"`
	Height      int     `msgpack:"height,omitempty"`
	ScaleFactor float64 `msgpack:"scaleFactor,omitempty"`
}
type NativeImageCreateFromBufferParams struct {
	Buffer  []byte                                   `msgpack:"buffer"`
	Options NativeImageCreateFromBufferParamsOptions `msgpack:"options,omitempty"`
}

func (c *NativeImageModule) CreateFromDataURL(params NativeImageCreateFromDataURLParams, ret interface{}) error {
	return c.Call("nativeImage.createFromDataURL", params, ret)
}

type NativeImageCreateFromDataURLParams struct {
	DataURL string `msgpack:"dataURL"`
}

func (c *NativeImageModule) CreateFromNamedImage(params NativeImageCreateFromNamedImageParams, ret interface{}) error {
	return c.Call("nativeImage.createFromNamedImage", params, ret)
}

type NativeImageCreateFromNamedImageParams struct {
	ImageName string `msgpack:"imageName"`
	HslShift  []int  `msgpack:"hslShift"`
}
type Menu struct{}

func (c *MenuModule) SetApplicationMenu(params MenuSetApplicationMenuParams, ret interface{}) error {
	return c.Call("Menu.setApplicationMenu", params, ret)
}

type MenuSetApplicationMenuParams struct {
	Menu qrpc.ObjectHandle `msgpack:"menu"`
}

func (c *MenuModule) GetApplicationMenu(ret interface{}) error {
	return c.Call("Menu.getApplicationMenu", nil, ret)
}
func (c *MenuModule) SendActionToFirstResponder(params MenuSendActionToFirstResponderParams, ret interface{}) error {
	return c.Call("Menu.sendActionToFirstResponder", params, ret)
}

type MenuSendActionToFirstResponderParams struct {
	Action string `msgpack:"action"`
}

func (c *MenuModule) BuildFromTemplate(params MenuBuildFromTemplateParams, ret interface{}) error {
	return c.Call("Menu.buildFromTemplate", params, ret)
}

type MenuBuildFromTemplateParams struct {
	Template []*MenuItemConstructorOptions `msgpack:"template"`
}
type MenuModule struct {
	*qrpc.Client
}
type MenuItem struct{}
type MenuItemParamsOptions struct {
	Click       qrpc.ObjectHandle           `msgpack:"click,omitempty"`
	Role        string                      `msgpack:"role,omitempty"`
	Type        string                      `msgpack:"type,omitempty"`
	Label       string                      `msgpack:"label,omitempty"`
	Sublabel    string                      `msgpack:"sublabel,omitempty"`
	Accelerator string                      `msgpack:"accelerator,omitempty"`
	Icon        qrpc.ObjectHandle           `msgpack:"icon,omitempty"`
	Enabled     bool                        `msgpack:"enabled,omitempty"`
	Visible     bool                        `msgpack:"visible,omitempty"`
	Checked     bool                        `msgpack:"checked,omitempty"`
	Submenu     *MenuItemConstructorOptions `msgpack:"submenu,omitempty"`
	Id          string                      `msgpack:"id,omitempty"`
	Position    string                      `msgpack:"position,omitempty"`
}
type MenuItemParams struct {
	Options MenuItemParamsOptions `msgpack:"options"`
}
type MenuItemModule struct {
	*qrpc.Client
}
type NativeImage struct{}
type Tray struct{}
type TrayParams struct {
	Image qrpc.ObjectHandle `msgpack:"image"`
}
type TrayModule struct {
	*qrpc.Client
}
type Notification struct{}

func (c *NotificationModule) IsSupported(ret interface{}) error {
	return c.Call("Notification.isSupported", nil, ret)
}

type NotificationParamsOptions struct {
	Title            string               `msgpack:"title,omitempty"`
	Subtitle         string               `msgpack:"subtitle,omitempty"`
	Body             string               `msgpack:"body,omitempty"`
	Silent           bool                 `msgpack:"silent,omitempty"`
	Icon             string               `msgpack:"icon,omitempty"`
	HasReply         bool                 `msgpack:"hasReply,omitempty"`
	ReplyPlaceholder string               `msgpack:"replyPlaceholder,omitempty"`
	Sound            string               `msgpack:"sound,omitempty"`
	Actions          []NotificationAction `msgpack:"actions,omitempty"`
	CloseButtonText  string               `msgpack:"closeButtonText,omitempty"`
}
type NotificationParams struct {
	Options NotificationParamsOptions `msgpack:"options"`
}
type NotificationModule struct {
	*qrpc.Client
}
