"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class ListenerHandle {
}
class Accelerator {
}
class Buffer {
}
class BrowserWindow {
}
class Session {
}
class Product {
}
class Transaction {
}
var api;
(function (api) {
    /**
     * Natively wrap images such as tray, dock, and application icons.
     */
    class NativeImage {
    }
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
    class MemoryInfo {
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
    class UploadData {
    }
    /**
     *
     */
    class NewMenuItem {
    }
    /**
     * Control your application's event lifecycle.
     */
    let app;
    (function (app) {
        /**
         * On Linux, focuses on the first visible window. On macOS, makes the application the active app. On Windows, focuses on the application's first window.
         */
        function focus() {
            console.log("focus");
        }
        app.focus = focus;
        /**
         * Hides all application windows without minimizing them.
         */
        function hide() {
            console.log("hide");
        }
        app.hide = hide;
        /**
         * Shows application windows after they were hidden. Does not automatically focus them.
         */
        function show() {
            console.log("show");
        }
        app.show = show;
        /**
         *
         */
        function getAppPath() {
            console.log("getAppPath");
            return "";
        }
        app.getAppPath = getAppPath;
        /**
         * You can request the following paths by the name:
         */
        function getPath(name) {
            console.log("getPath");
            return "";
        }
        app.getPath = getPath;
        /**
         * Fetches a path's associated icon. On Windows, there a 2 kinds of icons: On Linux and macOS, icons depend on the application associated with file mime type.
         */
        function getFileIcon(path, options) {
            console.log("getFileIcon");
            return null;
        }
        app.getFileIcon = getFileIcon;
        /**
         *
         */
        function getVersion() {
            console.log("getVersion");
            return "";
        }
        app.getVersion = getVersion;
        /**
         * To set the locale, you'll want to use a command line switch at app startup, which may be found here. Note: When distributing your packaged app, you have to also ship the locales folder. Note: On Windows you have to call it after the ready events gets emitted.
         */
        function getLocale() {
            console.log("getLocale");
            return "";
        }
        app.getLocale = getLocale;
        /**
         * This method makes your application a Single Instance Application - instead of allowing multiple instances of your app to run, this will ensure that only a single instance of your app is running, and other instances signal this instance and exit. The return value of this method indicates whether or not this instance of your application successfully obtained the lock.  If it failed to obtain the lock you can assume that another instance of your application is already running with the lock and exit immediately. I.e. This method returns true if your process is the primary instance of your application and your app should continue loading.  It returns false if your process should immediately quit as it has sent its parameters to another instance that has already acquired the lock. On macOS the system enforces single instance automatically when users try to open a second instance of your app in Finder, and the open-file and open-url events will be emitted for that. However when users start your app in command line the system's single instance mechanism will be bypassed and you have to use this method to ensure single instance. An example of activating the window of primary instance when a second instance starts:
         */
        function requestSingleInstanceLock() {
            console.log("requestSingleInstanceLock");
            return false;
        }
        app.requestSingleInstanceLock = requestSingleInstanceLock;
        /**
         * This method returns whether or not this instance of your app is currently holding the single instance lock.  You can request the lock with app.requestSingleInstanceLock() and release with app.releaseSingleInstanceLock()
         */
        function hasSingleInstanceLock() {
            console.log("hasSingleInstanceLock");
            return false;
        }
        app.hasSingleInstanceLock = hasSingleInstanceLock;
        /**
         * Releases all locks that were created by requestSingleInstanceLock. This will allow multiple instances of the application to once again run side by side.
         */
        function releaseSingleInstanceLock() {
            console.log("releaseSingleInstanceLock");
        }
        app.releaseSingleInstanceLock = releaseSingleInstanceLock;
        /**
         *
         */
        function getAppMetrics() {
            console.log("getAppMetrics");
            return new Array();
        }
        app.getAppMetrics = getAppMetrics;
        /**
         * Sets the counter badge for current app. Setting the count to 0 will hide the badge. On macOS it shows on the dock icon. On Linux it only works for Unity launcher, Note: Unity launcher requires the existence of a .desktop file to work, for more information please read Desktop Environment Integration.
         */
        function setBadgeCount(count) {
            console.log("setBadgeCount");
            return false;
        }
        app.setBadgeCount = setBadgeCount;
        /**
         *
         */
        function getBadgeCount() {
            console.log("getBadgeCount");
            return 0;
        }
        app.getBadgeCount = getBadgeCount;
        /**
         * Start accessing a security scoped resource. With this method electron applications that are packaged for the Mac App Store may reach outside their sandbox to access files chosen by the user. See Apple's documentation for a description of how this system works.
         */
        function startAccessingSecurityScopedResource(bookmarkData) {
            console.log("startAccessingSecurityScopedResource");
            return function () { };
        }
        app.startAccessingSecurityScopedResource = startAccessingSecurityScopedResource;
        /**
         * Emitted when all windows have been closed and the application will quit. Calling event.preventDefault() will prevent the default behaviour, which is terminating the application. See the description of the window-all-closed event for the differences between the will-quit and window-all-closed events. Note: On Windows, this event will not be emitted if the app is closed due to a shutdown/restart of the system or a user logout.
         * @event
         */
        app.WillQuit = 'will-quit';
        /**
         * Emitted when the application is quitting. Note: On Windows, this event will not be emitted if the app is closed due to a shutdown/restart of the system or a user logout.
         * @event
         */
        app.Quit = 'quit';
        /**
         * Emitted when Electron has created a new session.
         * @event
         */
        app.SessionCreated = 'session-created';
        /**
         * This event will be emitted inside the primary instance of your application when a second instance has been executed. argv is an Array of the second instance's command line arguments, and workingDirectory is its current working directory. Usually applications respond to this by making their primary window focused and non-minimized. This event is guaranteed to be emitted after the ready event of app gets emitted.
         * @event
         */
        app.SecondInstance = 'second-instance';
        function on(event, listener) { return new ListenerHandle(); }
        app.on = on;
        function off(event, handle) { }
        app.off = off;
    })(app = api.app || (api.app = {}));
    /**
     * Perform copy and paste operations on the system clipboard.
     */
    let clipboard;
    (function (clipboard) {
        /**
         *
         */
        function readText(type) {
            console.log("readText");
            return "";
        }
        clipboard.readText = readText;
        /**
         * Writes the text into the clipboard as plain text.
         */
        function writeText(text, type) {
            console.log("writeText");
        }
        clipboard.writeText = writeText;
        /**
         *
         */
        function readHTML(type) {
            console.log("readHTML");
            return "";
        }
        clipboard.readHTML = readHTML;
        /**
         * Writes markup to the clipboard.
         */
        function writeHTML(markup, type) {
            console.log("writeHTML");
        }
        clipboard.writeHTML = writeHTML;
        /**
         *
         */
        function readImage(type) {
            console.log("readImage");
            return new NativeImage();
        }
        clipboard.readImage = readImage;
        /**
         * Writes image to the clipboard.
         */
        function writeImage(image, type) {
            console.log("writeImage");
        }
        clipboard.writeImage = writeImage;
        /**
         *
         */
        function readRTF(type) {
            console.log("readRTF");
            return "";
        }
        clipboard.readRTF = readRTF;
        /**
         * Writes the text into the clipboard in RTF.
         */
        function writeRTF(text, type) {
            console.log("writeRTF");
        }
        clipboard.writeRTF = writeRTF;
        /**
         * Returns an Object containing title and url keys representing the bookmark in the clipboard. The title and url values will be empty strings when the bookmark is unavailable.
         */
        function readBookmark() {
            console.log("readBookmark");
            return null;
        }
        clipboard.readBookmark = readBookmark;
        /**
         * Writes the title and url into the clipboard as a bookmark. Note: Most apps on Windows don't support pasting bookmarks into them so you can use clipboard.write to write both a bookmark and fallback text to the clipboard.
         */
        function writeBookmark(title, url, type) {
            console.log("writeBookmark");
        }
        clipboard.writeBookmark = writeBookmark;
        /**
         * Clears the clipboard content.
         */
        function clear(type) {
            console.log("clear");
        }
        clipboard.clear = clear;
        /**
         *
         */
        function availableFormats(type) {
            console.log("availableFormats");
            return [""];
        }
        clipboard.availableFormats = availableFormats;
    })(clipboard = api.clipboard || (api.clipboard = {}));
    /**
     * Display native system dialogs for opening and saving files, alerting, etc.
     */
    let dialog;
    (function (dialog) {
        /**
         * The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed or selected when you want to limit the user to a specific type. For example: The extensions array should contain extensions without wildcards or dots (e.g. 'png' is good but '.png' and '*.png' are bad). To show all files, use the '*' wildcard (no other wildcard is supported). If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filenames). Note: On Windows and Linux an open dialog can not be both a file selector and a directory selector, so if you set properties to ['openFile', 'openDirectory'] on these platforms, a directory selector will be shown.
         */
        function showOpenDialog(options) {
            console.log("showOpenDialog");
            return null;
        }
        dialog.showOpenDialog = showOpenDialog;
        /**
         * The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed, see dialog.showOpenDialog for an example. If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filename).
         */
        function showSaveDialog(options) {
            console.log("showSaveDialog");
            return null;
        }
        dialog.showSaveDialog = showSaveDialog;
        /**
         * Shows a message box, it will block the process until the message box is closed. It returns the index of the clicked button. The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. If a callback is passed, the dialog will not block the process. The API call will be asynchronous and the result will be passed via callback(response).
         */
        function showMessageBox(options) {
            console.log("showMessageBox");
            return null;
        }
        dialog.showMessageBox = showMessageBox;
        /**
         * Displays a modal dialog that shows an error message. This API can be called safely before the ready event the app module emits, it is usually used to report errors in early stage of startup. If called before the app readyevent on Linux, the message will be emitted to stderr, and no GUI dialog will appear.
         */
        function showErrorBox(title, content) {
            console.log("showErrorBox");
        }
        dialog.showErrorBox = showErrorBox;
    })(dialog = api.dialog || (api.dialog = {}));
    /**
     * Detect keyboard events when the application does not have keyboard focus.
     */
    let globalShortcut;
    (function (globalShortcut) {
        /**
         * Registers a global shortcut of accelerator. The callback is called when the registered shortcut is pressed by the user. When the accelerator is already taken by other applications, this call will silently fail. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
         */
        function register(accelerator, callback) {
            console.log("register");
        }
        globalShortcut.register = register;
        /**
         * When the accelerator is already taken by other applications, this call will still return false. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
         */
        function isRegistered(accelerator) {
            console.log("isRegistered");
            return false;
        }
        globalShortcut.isRegistered = isRegistered;
        /**
         * Unregisters the global shortcut of accelerator.
         */
        function unregister(accelerator) {
            console.log("unregister");
        }
        globalShortcut.unregister = unregister;
    })(globalShortcut = api.globalShortcut || (api.globalShortcut = {}));
    /**
     * In-app purchases on Mac App Store.
     */
    let inAppPurchase;
    (function (inAppPurchase) {
        /**
         * You should listen for the transactions-updated event as soon as possible and certainly before you call purchaseProduct.
         */
        function purchaseProduct(productID, quantity, callback) {
            console.log("purchaseProduct");
        }
        inAppPurchase.purchaseProduct = purchaseProduct;
        /**
         * Retrieves the product descriptions.
         */
        function getProducts(productIDs, callback) {
            console.log("getProducts");
        }
        inAppPurchase.getProducts = getProducts;
        /**
         *
         */
        function canMakePayments() {
            console.log("canMakePayments");
            return false;
        }
        inAppPurchase.canMakePayments = canMakePayments;
        /**
         *
         */
        function getReceiptURL() {
            console.log("getReceiptURL");
            return "";
        }
        inAppPurchase.getReceiptURL = getReceiptURL;
        /**
         * Completes all pending transactions.
         */
        function finishAllTransactions() {
            console.log("finishAllTransactions");
        }
        inAppPurchase.finishAllTransactions = finishAllTransactions;
        /**
         * Completes the pending transactions corresponding to the date.
         */
        function finishTransactionByDate(date) {
            console.log("finishTransactionByDate");
        }
        inAppPurchase.finishTransactionByDate = finishTransactionByDate;
        /**
         * Emitted when one or more transactions have been updated.
         * @event
         */
        inAppPurchase.TransactionsUpdated = 'transactions-updated';
        function on(event, listener) { return new ListenerHandle(); }
        inAppPurchase.on = on;
        function off(event, handle) { }
        inAppPurchase.off = off;
    })(inAppPurchase = api.inAppPurchase || (api.inAppPurchase = {}));
    /**
     * Logging network events.
     */
    let netLog;
    (function (netLog) {
        /**
         * Starts recording network events to path.
         */
        function startLogging(path) {
            console.log("startLogging");
        }
        netLog.startLogging = startLogging;
        /**
         * Stops recording network events. If not called, net logging will automatically end when app quits.
         */
        function stopLogging(callback) {
            console.log("stopLogging");
        }
        netLog.stopLogging = stopLogging;
    })(netLog = api.netLog || (api.netLog = {}));
    /**
     * Monitor power state changes.
     */
    let powerMonitor;
    (function (powerMonitor) {
        /**
         * Emitted when the system is suspending.
         * @event
         */
        powerMonitor.Suspend = 'suspend';
        /**
         * Emitted when system is resuming.
         * @event
         */
        powerMonitor.Resume = 'resume';
        /**
         * Emitted when the system changes to AC power.
         * @event
         */
        powerMonitor.OnAc = 'on-ac';
        /**
         * Emitted when system changes to battery power.
         * @event
         */
        powerMonitor.OnBattery = 'on-battery';
        /**
         * Emitted when the system is about to reboot or shut down. If the event handler invokes e.preventDefault(), Electron will attempt to delay system shutdown in order for the app to exit cleanly. If e.preventDefault() is called, the app should exit as soon as possible by calling something like app.quit().
         * @event
         */
        powerMonitor.Shutdown = 'shutdown';
        /**
         * Emitted when the system is about to lock the screen.
         * @event
         */
        powerMonitor.LockScreen = 'lock-screen';
        /**
         * Emitted as soon as the systems screen is unlocked.
         * @event
         */
        powerMonitor.UnlockScreen = 'unlock-screen';
        function on(event, listener) { return new ListenerHandle(); }
        powerMonitor.on = on;
        function off(event, handle) { }
        powerMonitor.off = off;
    })(powerMonitor = api.powerMonitor || (api.powerMonitor = {}));
    /**
     * Extensions to process object.
     */
    let process;
    (function (process) {
        /**
         *
         */
        function getCPUUsage() {
            console.log("getCPUUsage");
            return new CPUUsage();
        }
        process.getCPUUsage = getCPUUsage;
        /**
         * Returns an object with V8 heap statistics. Note that all statistics are reported in Kilobytes.
         */
        function getHeapStatistics() {
            console.log("getHeapStatistics");
            return null;
        }
        process.getHeapStatistics = getHeapStatistics;
        /**
         * Returns an object giving memory usage statistics about the current process. Note that all statistics are reported in Kilobytes.
         */
        function getProcessMemoryInfo() {
            console.log("getProcessMemoryInfo");
            return null;
        }
        process.getProcessMemoryInfo = getProcessMemoryInfo;
        /**
         * Returns an object giving memory usage statistics about the entire system. Note that all statistics are reported in Kilobytes.
         */
        function getSystemMemoryInfo() {
            console.log("getSystemMemoryInfo");
            return null;
        }
        process.getSystemMemoryInfo = getSystemMemoryInfo;
    })(process = api.process || (api.process = {}));
    /**
     * Register a custom protocol and intercept existing protocol requests.
     */
    let protocol;
    (function (protocol) {
        /**
         * A standard scheme adheres to what RFC 3986 calls generic URI syntax. For example http and https are standard schemes, while file is not. Registering a scheme as standard, will allow relative and absolute resources to be resolved correctly when served. Otherwise the scheme will behave like the file protocol, but without the ability to resolve relative URLs. For example when you load following page with custom protocol without registering it as standard scheme, the image will not be loaded because non-standard schemes can not recognize relative URLs: Registering a scheme as standard will allow access to files through the FileSystem API. Otherwise the renderer will throw a security error for the scheme. By default web storage apis (localStorage, sessionStorage, webSQL, indexedDB, cookies) are disabled for non standard schemes. So in general if you want to register a custom protocol to replace the http protocol, you have to register it as a standard scheme: Note: This method can only be used before the ready event of the app module gets emitted.
         */
        function registerStandardSchemes(schemes, options) {
            console.log("registerStandardSchemes");
        }
        protocol.registerStandardSchemes = registerStandardSchemes;
        /**
         * Registers a protocol of scheme that will send the file as a response. The handler will be called with handler(request, callback) when a request is going to be created with scheme. completion will be called with completion(null) when scheme is successfully registered or completion(error) when failed. To handle the request, the callback should be called with either the file's path or an object that has a path property, e.g. callback(filePath) or callback({path: filePath}). When callback is called with nothing, a number, or an object that has an error property, the request will fail with the error number you specified. For the available error numbers you can use, please see the net error list. By default the scheme is treated like http:, which is parsed differently than protocols that follow the "generic URI syntax" like file:, so you probably want to call protocol.registerStandardSchemes to have your scheme treated as a standard scheme.
         */
        function registerFileProtocol(scheme, handler) {
            console.log("registerFileProtocol");
            return null;
        }
        protocol.registerFileProtocol = registerFileProtocol;
        /**
         * Registers a protocol of scheme that will send a String as a response. The usage is the same with registerFileProtocol, except that the callback should be called with either a String or an object that has the data, mimeType, and charset properties.
         */
        function registerStringProtocol(scheme, handler) {
            console.log("registerStringProtocol");
            return null;
        }
        protocol.registerStringProtocol = registerStringProtocol;
        /**
         * Registers a protocol of scheme that will send an HTTP request as a response. The usage is the same with registerFileProtocol, except that the callback should be called with a redirectRequest object that has the url, method, referrer, uploadData and session properties. By default the HTTP request will reuse the current session. If you want the request to have a different session you should set session to null. For POST requests the uploadData object must be provided.
         */
        function registerHttpProtocol(scheme, handler) {
            console.log("registerHttpProtocol");
            return null;
        }
        protocol.registerHttpProtocol = registerHttpProtocol;
        /**
         * Unregisters the custom protocol of scheme.
         */
        function unregisterProtocol(scheme) {
            console.log("unregisterProtocol");
            return null;
        }
        protocol.unregisterProtocol = unregisterProtocol;
        /**
         * The callback will be called with a boolean that indicates whether there is already a handler for scheme.
         */
        function isProtocolHandled(scheme) {
            console.log("isProtocolHandled");
            return null;
        }
        protocol.isProtocolHandled = isProtocolHandled;
    })(protocol = api.protocol || (api.protocol = {}));
    /**
     * Retrieve information about screen size, displays, cursor position, etc.
     */
    let screen;
    (function (screen) {
        /**
         * The current absolute position of the mouse pointer.
         */
        function getCursorScreenPoint() {
            console.log("getCursorScreenPoint");
            return new Point();
        }
        screen.getCursorScreenPoint = getCursorScreenPoint;
        /**
         *
         */
        function getPrimaryDisplay() {
            console.log("getPrimaryDisplay");
            return new Display();
        }
        screen.getPrimaryDisplay = getPrimaryDisplay;
        /**
         *
         */
        function getAllDisplays() {
            console.log("getAllDisplays");
            return new Array();
        }
        screen.getAllDisplays = getAllDisplays;
        /**
         *
         */
        function getDisplayNearestPoint(point) {
            console.log("getDisplayNearestPoint");
            return new Display();
        }
        screen.getDisplayNearestPoint = getDisplayNearestPoint;
        /**
         *
         */
        function getDisplayMatching(rect) {
            console.log("getDisplayMatching");
            return new Display();
        }
        screen.getDisplayMatching = getDisplayMatching;
        /**
         * Converts a screen physical point to a screen DIP point. The DPI scale is performed relative to the display containing the physical point.
         */
        function screenToDipPoint(point) {
            console.log("screenToDipPoint");
            return new Point();
        }
        screen.screenToDipPoint = screenToDipPoint;
        /**
         * Converts a screen DIP point to a screen physical point. The DPI scale is performed relative to the display containing the DIP point.
         */
        function dipToScreenPoint(point) {
            console.log("dipToScreenPoint");
            return new Point();
        }
        screen.dipToScreenPoint = dipToScreenPoint;
        /**
         * Converts a screen physical rect to a screen DIP rect. The DPI scale is performed relative to the display nearest to window. If window is null, scaling will be performed to the display nearest to rect.
         */
        function screenToDipRect(window, rect) {
            console.log("screenToDipRect");
            return new Rectangle();
        }
        screen.screenToDipRect = screenToDipRect;
        /**
         * Converts a screen DIP rect to a screen physical rect. The DPI scale is performed relative to the display nearest to window. If window is null, scaling will be performed to the display nearest to rect.
         */
        function dipToScreenRect(window, rect) {
            console.log("dipToScreenRect");
            return new Rectangle();
        }
        screen.dipToScreenRect = dipToScreenRect;
        /**
         * Emitted when newDisplay has been added.
         * @event
         */
        screen.DisplayAdded = 'display-added';
        /**
         * Emitted when oldDisplay has been removed.
         * @event
         */
        screen.DisplayRemoved = 'display-removed';
        /**
         * Emitted when one or more metrics change in a display. The changedMetrics is an array of strings that describe the changes. Possible changes are bounds, workArea, scaleFactor and rotation.
         * @event
         */
        screen.DisplayMetricsChanged = 'display-metrics-changed';
        function on(event, listener) { return new ListenerHandle(); }
        screen.on = on;
        function off(event, handle) { }
        screen.off = off;
    })(screen = api.screen || (api.screen = {}));
    /**
     * Manage files and URLs using their default applications.
     */
    let shell;
    (function (shell) {
        /**
         * Show the given file in a file manager. If possible, select the file.
         */
        function showItemInFolder(fullPath) {
            console.log("showItemInFolder");
            return false;
        }
        shell.showItemInFolder = showItemInFolder;
        /**
         * Open the given file in the desktop's default manner.
         */
        function openItem(fullPath) {
            console.log("openItem");
            return false;
        }
        shell.openItem = openItem;
        /**
         * Open the given external protocol URL in the desktop's default manner. (For example, mailto: URLs in the user's default mail agent).
         */
        function openExternal(url, options) {
            console.log("openExternal");
            return null;
        }
        shell.openExternal = openExternal;
        /**
         * Move the given file to trash and returns a boolean status for the operation.
         */
        function moveItemToTrash(fullPath) {
            console.log("moveItemToTrash");
            return false;
        }
        shell.moveItemToTrash = moveItemToTrash;
        /**
         * Play the beep sound.
         */
        function beep() {
            console.log("beep");
        }
        shell.beep = beep;
        /**
         * Creates or updates a shortcut link at shortcutPath.
         */
        function writeShortcutLink(shortcutPath, operation, options) {
            console.log("writeShortcutLink");
            return false;
        }
        shell.writeShortcutLink = writeShortcutLink;
        /**
         * Resolves the shortcut link at shortcutPath. An exception will be thrown when any error happens.
         */
        function readShortcutLink(shortcutPath) {
            console.log("readShortcutLink");
            return new ShortcutDetails();
        }
        shell.readShortcutLink = readShortcutLink;
    })(shell = api.shell || (api.shell = {}));
    /**
     *
     */
    let menu;
    (function (menu_1) {
        /**
         * Sets menu as the application menu on macOS. On Windows and Linux, the menu will be set as each window's top menu. Passing null will remove the menu bar on Windows and Linux but has no effect on macOS. Note: This API has to be called after the ready event of app module.
         */
        function setApplicationMenu(menu) {
            console.log("setApplicationMenu");
        }
        menu_1.setApplicationMenu = setApplicationMenu;
        /**
         * Note: The returned Menu instance doesn't support dynamic addition or removal of menu items. Instance properties can still be dynamically modified.
         */
        function getApplicationMenu() {
            console.log("getApplicationMenu");
            return new Menu();
        }
        menu_1.getApplicationMenu = getApplicationMenu;
        /**
         * Sends the action to the first responder of application. This is used for emulating default macOS menu behaviors. Usually you would use the role property of a MenuItem. See the macOS Cocoa Event Handling Guide for more information on macOS' native actions.
         */
        function sendActionToFirstResponder(action) {
            console.log("sendActionToFirstResponder");
        }
        menu_1.sendActionToFirstResponder = sendActionToFirstResponder;
        /**
         * Generally, the template is an array of options for constructing a MenuItem. The usage can be referenced above. You can also attach other fields to the element of the template and they will become properties of the constructed menu items.
         */
        function buildFromTemplate(template) {
            console.log("buildFromTemplate");
            return new Menu();
        }
        menu_1.buildFromTemplate = buildFromTemplate;
        /**
         *
         */
        function make() {
            console.log("make");
            return new Menu();
        }
        menu_1.make = make;
        /**
         *
         */
        function ref(handle) {
            console.log("ref");
            return new Menu();
        }
        menu_1.ref = ref;
    })(menu = api.menu || (api.menu = {}));
    /**
     *
     */
    (function (menu) {
        let item;
        (function (item) {
            /**
             *
             */
            function make(options) {
                console.log("make");
                return new MenuItem();
            }
            item.make = make;
            /**
             *
             */
            function ref(handle) {
                console.log("ref");
                return new MenuItem();
            }
            item.ref = ref;
        })(item = menu.item || (menu.item = {}));
    })(menu = api.menu || (api.menu = {}));
    /**
     *
     */
    (function (app) {
        let dock;
        (function (dock) {
            /**
             * When critical is passed, the dock icon will bounce until either the application becomes active or the request is canceled. When informational is passed, the dock icon will bounce for one second. However, the request remains active until either the application becomes active or the request is canceled.
             */
            function bounce(type) {
                console.log("bounce");
                return 0;
            }
            dock.bounce = bounce;
            /**
             * Cancel the bounce of id.
             */
            function cancelBounce(id) {
                console.log("cancelBounce");
            }
            dock.cancelBounce = cancelBounce;
            /**
             * Bounces the Downloads stack if the filePath is inside the Downloads folder.
             */
            function downloadFinished(filePath) {
                console.log("downloadFinished");
            }
            dock.downloadFinished = downloadFinished;
            /**
             * Sets the string to be displayed in the dock’s badging area.
             */
            function setBadge(text) {
                console.log("setBadge");
            }
            dock.setBadge = setBadge;
            /**
             *
             */
            function getBadge() {
                console.log("getBadge");
                return "";
            }
            dock.getBadge = getBadge;
            /**
             * Hides the dock icon.
             */
            function hide() {
                console.log("hide");
            }
            dock.hide = hide;
            /**
             * Shows the dock icon.
             */
            function show() {
                console.log("show");
            }
            dock.show = show;
            /**
             *
             */
            function isVisible() {
                console.log("isVisible");
                return false;
            }
            dock.isVisible = isVisible;
            /**
             * Sets the application's dock menu.
             */
            function setMenu(menu) {
                console.log("setMenu");
            }
            dock.setMenu = setMenu;
        })(dock = app.dock || (app.dock = {}));
    })(app = api.app || (api.app = {}));
})(api || (api = {}));
exports.default = api;
//# sourceMappingURL=api.js.map