class Integer {
}
class Float {
}
class Double {
}
class Accelerator {
}
class MenuItemConstructorOptions {
}
class Protocol {
}
class TouchBarItem {
}
class WebFrame {
}
class BrowserWindow {
}
var electron;
(function (electron) {
    /**
     *
     */
    class CPUUsage {
    }
    /**
     *
     */
    class Display {
    }
    /**
     *
     */
    class FileFilter {
    }
    /**
     *
     */
    class Point {
    }
    /**
     *
     */
    class ProcessMetric {
    }
    /**
     *
     */
    class Rectangle {
    }
    /**
     *
     */
    class ShortcutDetails {
    }
    /**
     *
     */
    class Size {
    }
    /**
     *
     */
    class MemoryInfo {
    }
    /**
     *
     */
    class DesktopCapturerSource {
    }
    /**
     *
     */
    class UploadData {
    }
    /**
     * Create native application menus and context menus.
     */
    class Menu {
    }
    /**
     * Add items to native application menus and context menus.
     */
    class MenuItem {
    }
    /**
     * Natively wrap images such as tray, dock, and application icons.
     */
    class NativeImage {
    }
    /**
     * Add icons and context menus to the system's notification area.
     */
    class Tray {
    }
    /**
     * Control your application's event lifecycle.
     */
    let app;
    (function (app) {
        let dock;
        (function (dock) {
            /**
             * When critical is passed, the dock icon will bounce until either the application becomes active or the request is canceled. When informational is passed, the dock icon will bounce for one second. However, the request remains active until either the application becomes active or the request is canceled.
             */
            function bounce(type) {
                return 0;
            }
            /**
             * Cancel the bounce of id.
             */
            function cancelBounce(id) {
            }
            /**
             * Bounces the Downloads stack if the filePath is inside the Downloads folder.
             */
            function downloadFinished(filePath) {
            }
            /**
             * Sets the string to be displayed in the dock’s badging area.
             */
            function setBadge(text) {
            }
            /**
             *
             */
            function getBadge() {
                return "";
            }
            /**
             * Hides the dock icon.
             */
            function hide() {
            }
            /**
             * Shows the dock icon.
             */
            function show() {
            }
            /**
             *
             */
            function isVisible() {
                return false;
            }
            /**
             * Sets the application's dock menu.
             */
            function setMenu(menu) {
            }
        })(dock || (dock = {}));
        /**
         * Try to close all windows. The before-quit event will be emitted first. If all windows are successfully closed, the will-quit event will be emitted and by default the application will terminate. This method guarantees that all beforeunload and unload event handlers are correctly executed. It is possible that a window cancels the quitting by returning false in the beforeunload event handler.
         */
        function quit() {
        }
        /**
         * On Linux, focuses on the first visible window. On macOS, makes the application the active app. On Windows, focuses on the application's first window.
         */
        function focus() {
        }
        /**
         * Hides all application windows without minimizing them.
         */
        function hide() {
        }
        /**
         * Shows application windows after they were hidden. Does not automatically focus them.
         */
        function show() {
        }
        /**
         *
         */
        function getAppPath() {
            return "";
        }
        /**
         * You can request the following paths by the name:
         */
        function getPath(name) {
            return "";
        }
        /**
         * Fetches a path's associated icon. On Windows, there a 2 kinds of icons: On Linux and macOS, icons depend on the application associated with file mime type.
         */
        function getFileIcon(path, options, callback) {
        }
        /**
         *
         */
        function getVersion() {
            return "";
        }
        /**
         * To set the locale, you'll want to use a command line switch at app startup, which may be found here. Note: When distributing your packaged app, you have to also ship the locales folder. Note: On Windows you have to call it after the ready events gets emitted.
         */
        function getLocale() {
            return "";
        }
        /**
         *
         */
        function getAppMetrics() {
            return new Array();
        }
        /**
         * Sets the counter badge for current app. Setting the count to 0 will hide the badge. On macOS it shows on the dock icon. On Linux it only works for Unity launcher, Note: Unity launcher requires the existence of a .desktop file to work, for more information please read Desktop Environment Integration.
         */
        function setBadgeCount(count) {
            return false;
        }
        /**
         *
         */
        function getBadgeCount() {
            return 0;
        }
        /**
         * Emitted when all windows have been closed and the application will quit. Calling event.preventDefault() will prevent the default behaviour, which is terminating the application. See the description of the window-all-closed event for the differences between the will-quit and window-all-closed events. Note: On Windows, this event will not be emitted if the app is closed due to a shutdown/restart of the system or a user logout.
         * @event
         */
        var WillQuit = "will-quit";
        /**
         * Emitted when the application is quitting. Note: On Windows, this event will not be emitted if the app is closed due to a shutdown/restart of the system or a user logout.
         * @event
         */
        var Quit = "quit";
    })(app || (app = {}));
    /**
     * Perform copy and paste operations on the system clipboard.
     */
    let clipboard;
    (function (clipboard) {
        /**
         *
         */
        function readText(type) {
            return "";
        }
        /**
         * Writes the text into the clipboard as plain text.
         */
        function writeText(text, type) {
        }
        /**
         *
         */
        function readHTML(type) {
            return "";
        }
        /**
         * Writes markup to the clipboard.
         */
        function writeHTML(markup, type) {
        }
        /**
         *
         */
        function readImage(type) {
            return new NativeImage();
        }
        /**
         * Writes image to the clipboard.
         */
        function writeImage(image, type) {
        }
        /**
         *
         */
        function readRTF(type) {
            return "";
        }
        /**
         * Writes the text into the clipboard in RTF.
         */
        function writeRTF(text, type) {
        }
        /**
         * Returns an Object containing title and url keys representing the bookmark in the clipboard. The title and url values will be empty strings when the bookmark is unavailable.
         */
        function readBookmark() {
            return null;
        }
        /**
         * Writes the title and url into the clipboard as a bookmark. Note: Most apps on Windows don't support pasting bookmarks into them so you can use clipboard.write to write both a bookmark and fallback text to the clipboard.
         */
        function writeBookmark(title, url, type) {
        }
        /**
         * Clears the clipboard content.
         */
        function clear(type) {
        }
        /**
         *
         */
        function availableFormats(type) {
            return [""];
        }
    })(clipboard || (clipboard = {}));
    /**
     * Access information about media sources that can be used to capture audio and
     * video from the desktop using the navigator.mediaDevices.getUserMedia API.
     */
    let desktopCapturer;
    (function (desktopCapturer) {
        /**
         * Starts gathering information about all available desktop media sources, and calls callback(error, sources) when finished. sources is an array of DesktopCapturerSource objects, each DesktopCapturerSource represents a screen or an individual window that can be captured.
         */
        function getSources(options, callback) {
        }
    })(desktopCapturer || (desktopCapturer = {}));
    /**
     * Display native system dialogs for opening and saving files, alerting, etc.
     */
    let dialog;
    (function (dialog) {
        /**
         * The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed or selected when you want to limit the user to a specific type. For example: The extensions array should contain extensions without wildcards or dots (e.g. 'png' is good but '.png' and '*.png' are bad). To show all files, use the '*' wildcard (no other wildcard is supported). If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filenames). Note: On Windows and Linux an open dialog can not be both a file selector and a directory selector, so if you set properties to ['openFile', 'openDirectory'] on these platforms, a directory selector will be shown.
         */
        function showOpenDialog(options) {
            return null;
        }
        /**
         * The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed, see dialog.showOpenDialog for an example. If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filename).
         */
        function showSaveDialog(options) {
            return null;
        }
        /**
         * Shows a message box, it will block the process until the message box is closed. It returns the index of the clicked button. The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. If a callback is passed, the dialog will not block the process. The API call will be asynchronous and the result will be passed via callback(response).
         */
        function showMessageBox(options) {
            return null;
        }
        /**
         * Displays a modal dialog that shows an error message. This API can be called safely before the ready event the app module emits, it is usually used to report errors in early stage of startup. If called before the app readyevent on Linux, the message will be emitted to stderr, and no GUI dialog will appear.
         */
        function showErrorBox(title, content) {
        }
    })(dialog || (dialog = {}));
    /**
     * Detect keyboard events when the application does not have keyboard focus.
     */
    let globalShortcut;
    (function (globalShortcut) {
        /**
         * Registers a global shortcut of accelerator. The callback is called when the registered shortcut is pressed by the user. When the accelerator is already taken by other applications, this call will silently fail. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
         */
        function register(accelerator, callback) {
        }
        /**
         * When the accelerator is already taken by other applications, this call will still return false. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
         */
        function isRegistered(accelerator) {
            return false;
        }
        /**
         * Unregisters the global shortcut of accelerator.
         */
        function unregister(accelerator) {
        }
    })(globalShortcut || (globalShortcut = {}));
    /**
     * Monitor power state changes.
     */
    let powerMonitor;
    (function (powerMonitor) {
        /**
         * Emitted when the system is suspending.
         * @event
         */
        var Suspend = "suspend";
        /**
         * Emitted when system is resuming.
         * @event
         */
        var Resume = "resume";
        /**
         * Emitted when the system changes to AC power.
         * @event
         */
        var OnAc = "on-ac";
        /**
         * Emitted when system changes to battery power.
         * @event
         */
        var OnBattery = "on-battery";
        /**
         * Emitted when the system is about to reboot or shut down. If the event handler invokes e.preventDefault(), Electron will attempt to delay system shutdown in order for the app to exit cleanly. If e.preventDefault() is called, the app should exit as soon as possible by calling something like app.quit().
         * @event
         */
        var Shutdown = "shutdown";
        /**
         * Emitted when the system is about to lock the screen.
         * @event
         */
        var LockScreen = "lock-screen";
        /**
         * Emitted as soon as the systems screen is unlocked.
         * @event
         */
        var UnlockScreen = "unlock-screen";
    })(powerMonitor || (powerMonitor = {}));
    /**
     * Extensions to process object.
     */
    let process;
    (function (process) {
        /**
         *
         */
        function getCPUUsage() {
            return new CPUUsage();
        }
        /**
         * Returns an object with V8 heap statistics. Note that all statistics are reported in Kilobytes.
         */
        function getHeapStatistics() {
            return null;
        }
        /**
         * Returns an object giving memory usage statistics about the current process. Note that all statistics are reported in Kilobytes.
         */
        function getProcessMemoryInfo() {
            return null;
        }
        /**
         * Returns an object giving memory usage statistics about the entire system. Note that all statistics are reported in Kilobytes.
         */
        function getSystemMemoryInfo() {
            return null;
        }
    })(process || (process = {}));
    /**
     * Register a custom protocol and intercept existing protocol requests.
     */
    let protocol;
    (function (protocol) {
        /**
         * A standard scheme adheres to what RFC 3986 calls generic URI syntax. For example http and https are standard schemes, while file is not. Registering a scheme as standard, will allow relative and absolute resources to be resolved correctly when served. Otherwise the scheme will behave like the file protocol, but without the ability to resolve relative URLs. For example when you load following page with custom protocol without registering it as standard scheme, the image will not be loaded because non-standard schemes can not recognize relative URLs: Registering a scheme as standard will allow access to files through the FileSystem API. Otherwise the renderer will throw a security error for the scheme. By default web storage apis (localStorage, sessionStorage, webSQL, indexedDB, cookies) are disabled for non standard schemes. So in general if you want to register a custom protocol to replace the http protocol, you have to register it as a standard scheme: Note: This method can only be used before the ready event of the app module gets emitted.
         */
        function registerStandardSchemes(schemes, options) {
        }
        /**
         * Registers a protocol of scheme that will send the file as a response. The handler will be called with handler(request, callback) when a request is going to be created with scheme. completion will be called with completion(null) when scheme is successfully registered or completion(error) when failed. To handle the request, the callback should be called with either the file's path or an object that has a path property, e.g. callback(filePath) or callback({path: filePath}). When callback is called with nothing, a number, or an object that has an error property, the request will fail with the error number you specified. For the available error numbers you can use, please see the net error list. By default the scheme is treated like http:, which is parsed differently than protocols that follow the "generic URI syntax" like file:, so you probably want to call protocol.registerStandardSchemes to have your scheme treated as a standard scheme.
         */
        function registerFileProtocol(scheme, handler) {
            return null;
        }
        /**
         * Registers a protocol of scheme that will send a String as a response. The usage is the same with registerFileProtocol, except that the callback should be called with either a String or an object that has the data, mimeType, and charset properties.
         */
        function registerStringProtocol(scheme, handler) {
            return null;
        }
        /**
         * Registers a protocol of scheme that will send an HTTP request as a response. The usage is the same with registerFileProtocol, except that the callback should be called with a redirectRequest object that has the url, method, referrer, uploadData and session properties. By default the HTTP request will reuse the current session. If you want the request to have a different session you should set session to null. For POST requests the uploadData object must be provided.
         */
        function registerHttpProtocol(scheme, handler) {
            return null;
        }
        /**
         * Unregisters the custom protocol of scheme.
         */
        function unregisterProtocol(scheme) {
            return null;
        }
        /**
         * The callback will be called with a boolean that indicates whether there is already a handler for scheme.
         */
        function isProtocolHandled(scheme) {
            return null;
        }
    })(protocol || (protocol = {}));
    /**
     * Retrieve information about screen size, displays, cursor position, etc.
     */
    let screen;
    (function (screen) {
        /**
         * The current absolute position of the mouse pointer.
         */
        function getCursorScreenPoint() {
            return new Point();
        }
        /**
         *
         */
        function getPrimaryDisplay() {
            return new Display();
        }
        /**
         *
         */
        function getAllDisplays() {
            return new Array();
        }
        /**
         *
         */
        function getDisplayNearestPoint(point) {
            return new Display();
        }
        /**
         *
         */
        function getDisplayMatching(rect) {
            return new Display();
        }
        /**
         * Converts a screen physical point to a screen DIP point. The DPI scale is performed relative to the display containing the physical point.
         */
        function screenToDipPoint(point) {
            return new Point();
        }
        /**
         * Converts a screen DIP point to a screen physical point. The DPI scale is performed relative to the display containing the DIP point.
         */
        function dipToScreenPoint(point) {
            return new Point();
        }
        /**
         * Converts a screen physical rect to a screen DIP rect. The DPI scale is performed relative to the display nearest to window. If window is null, scaling will be performed to the display nearest to rect.
         */
        function screenToDipRect(window, rect) {
            return new Rectangle();
        }
        /**
         * Converts a screen DIP rect to a screen physical rect. The DPI scale is performed relative to the display nearest to window. If window is null, scaling will be performed to the display nearest to rect.
         */
        function dipToScreenRect(window, rect) {
            return new Rectangle();
        }
        /**
         * Emitted when newDisplay has been added.
         * @event
         */
        var DisplayAdded = "display-added";
        /**
         * Emitted when oldDisplay has been removed.
         * @event
         */
        var DisplayRemoved = "display-removed";
        /**
         * Emitted when one or more metrics change in a display. The changedMetrics is an array of strings that describe the changes. Possible changes are bounds, workArea, scaleFactor and rotation.
         * @event
         */
        var DisplayMetricsChanged = "display-metrics-changed";
    })(screen || (screen = {}));
    /**
     * Manage files and URLs using their default applications.
     */
    let shell;
    (function (shell) {
        /**
         * Show the given file in a file manager. If possible, select the file.
         */
        function showItemInFolder(fullPath) {
            return false;
        }
        /**
         * Open the given file in the desktop's default manner.
         */
        function openItem(fullPath) {
            return false;
        }
        /**
         * Open the given external protocol URL in the desktop's default manner. (For example, mailto: URLs in the user's default mail agent).
         */
        function openExternal(url, options, callback) {
            return false;
        }
        /**
         * Move the given file to trash and returns a boolean status for the operation.
         */
        function moveItemToTrash(fullPath) {
            return false;
        }
        /**
         * Play the beep sound.
         */
        function beep() {
        }
        /**
         * Creates or updates a shortcut link at shortcutPath.
         */
        function writeShortcutLink(shortcutPath, operation, options) {
            return false;
        }
        /**
         * Resolves the shortcut link at shortcutPath. An exception will be thrown when any error happens.
         */
        function readShortcutLink(shortcutPath) {
            return new ShortcutDetails();
        }
    })(shell || (shell = {}));
})(electron || (electron = {}));
//# sourceMappingURL=schema.js.map