import electron from "electron";

class ListenerHandle {}
class Accelerator {}
class Buffer {}
class BrowserWindow {}
class Session {}
class Product {}
class Transaction {}

namespace api {
  /**
   * Natively wrap images such as tray, dock, and application icons.
   */
  class NativeImage {}

  /**
   * 
   */
  class CPUUsage {
    percentCPUUsage: number;
    idleWakeupsPerSecond: number;
  }
  /**
   * 
   */
  class Display {
    id: number;
    rotation: number;
    scaleFactor: number;
    touchSupport: string;
    bounds: Rectangle;
    size: Size;
    workArea: Rectangle;
    workAreaSize: Size;
  }
  /**
   * 
   */
  class FileFilter {
    name: string;
    extensions: string[];
  }
  /**
   * 
   */
  class MemoryInfo {
    pid: number;
    workingSetSize: number;
    peakWorkingSetSize: number;
    privateBytes: number;
    sharedBytes: number;
  }
  /**
   * Create native application menus and context menus.
   */
  class Menu {
    handle: string;
    items: MenuItem[];
  }
  /**
   * Add items to native application menus and context menus.
   */
  class MenuItem {
    handle: string;
    enabled: boolean;
    visible: boolean;
    checked: boolean;
    label: string;
    click: () => void;
  }
  /**
   * 
   */
  class Point {
    x: number;
    y: number;
  }
  /**
   * 
   */
  class ProcessMetric {
    pid: number;
    type: string;
    memory: MemoryInfo;
    cpu: CPUUsage;
  }
  /**
   * 
   */
  class Rectangle {
    x: number;
    y: number;
    width: number;
    height: number;
  }
  /**
   * 
   */
  class ShortcutDetails {
    target: string;
    cwd: string;
    args: string;
    description: string;
    icon: string;
    iconIndex: number;
    appUserModelId: string;
  }
  /**
   * 
   */
  class Size {
    width: number;
    height: number;
  }
  /**
   * 
   */
  class UploadData {
    bytes: Buffer;
    file: string;
    blobUUID: string;
  }
  /**
   * 
   */
  class NewMenuItem {
    click: (menuItem: MenuItem, event: Event) => void;
    role: string;
    type: string;
    label: string;
    sublabel: string;
    accelerator: Accelerator;
    icon: NativeImage;
    enabled: boolean;
    visible: boolean;
    checked: boolean;
    submenu: NewMenuItem[];
    id: string;
    position: string;
  }

    /**
     * Control your application's event lifecycle.
     */
    export namespace app {

          /**
           * On Linux, focuses on the first visible window. On macOS, makes the application the active app. On Windows, focuses on the application's first window.
           */
          export function focus() {
              console.log("focus");
              
          }

          /**
           * Hides all application windows without minimizing them.
           */
          export function hide() {
              console.log("hide");
              
          }

          /**
           * Shows application windows after they were hidden. Does not automatically focus them.
           */
          export function show() {
              console.log("show");
              
          }

          /**
           * 
           */
          export function getAppPath(): string {
              console.log("getAppPath");
              return "";
          }

          /**
           * You can request the following paths by the name:
           */
          export function getPath(name: string): string {
              console.log("getPath");
              return "";
          }

          /**
           * Fetches a path's associated icon. On Windows, there a 2 kinds of icons: On Linux and macOS, icons depend on the application associated with file mime type.
           */
          export function getFileIcon(path: string, options: {size: string}): {error: Error, icon: NativeImage} {
              console.log("getFileIcon");
              return null;
          }

          /**
           * 
           */
          export function getVersion(): string {
              console.log("getVersion");
              return "";
          }

          /**
           * To set the locale, you'll want to use a command line switch at app startup, which may be found here. Note: When distributing your packaged app, you have to also ship the locales folder. Note: On Windows you have to call it after the ready events gets emitted.
           */
          export function getLocale(): string {
              console.log("getLocale");
              return "";
          }

          /**
           * This method makes your application a Single Instance Application - instead of allowing multiple instances of your app to run, this will ensure that only a single instance of your app is running, and other instances signal this instance and exit. The return value of this method indicates whether or not this instance of your application successfully obtained the lock.  If it failed to obtain the lock you can assume that another instance of your application is already running with the lock and exit immediately. I.e. This method returns true if your process is the primary instance of your application and your app should continue loading.  It returns false if your process should immediately quit as it has sent its parameters to another instance that has already acquired the lock. On macOS the system enforces single instance automatically when users try to open a second instance of your app in Finder, and the open-file and open-url events will be emitted for that. However when users start your app in command line the system's single instance mechanism will be bypassed and you have to use this method to ensure single instance. An example of activating the window of primary instance when a second instance starts:
           */
          export function requestSingleInstanceLock(): boolean {
              console.log("requestSingleInstanceLock");
              return false;
          }

          /**
           * This method returns whether or not this instance of your app is currently holding the single instance lock.  You can request the lock with app.requestSingleInstanceLock() and release with app.releaseSingleInstanceLock()
           */
          export function hasSingleInstanceLock(): boolean {
              console.log("hasSingleInstanceLock");
              return false;
          }

          /**
           * Releases all locks that were created by requestSingleInstanceLock. This will allow multiple instances of the application to once again run side by side.
           */
          export function releaseSingleInstanceLock() {
              console.log("releaseSingleInstanceLock");
              
          }

          /**
           * 
           */
          export function getAppMetrics(): ProcessMetric[] {
              console.log("getAppMetrics");
              return new Array<ProcessMetric>();
          }

          /**
           * Sets the counter badge for current app. Setting the count to 0 will hide the badge. On macOS it shows on the dock icon. On Linux it only works for Unity launcher, Note: Unity launcher requires the existence of a .desktop file to work, for more information please read Desktop Environment Integration.
           */
          export function setBadgeCount(count: number): boolean {
              console.log("setBadgeCount");
              return false;
          }

          /**
           * 
           */
          export function getBadgeCount(): number {
              console.log("getBadgeCount");
              return 0;
          }

          /**
           * Start accessing a security scoped resource. With this method electron applications that are packaged for the Mac App Store may reach outside their sandbox to access files chosen by the user. See Apple's documentation for a description of how this system works.
           */
          export function startAccessingSecurityScopedResource(bookmarkData: string): () => void {
              console.log("startAccessingSecurityScopedResource");
              return function() {};
          }


        /**
         * Emitted when all windows have been closed and the application will quit. Calling event.preventDefault() will prevent the default behaviour, which is terminating the application. See the description of the window-all-closed event for the differences between the will-quit and window-all-closed events. Note: On Windows, this event will not be emitted if the app is closed due to a shutdown/restart of the system or a user logout.
         * @event
         */
        export var WillQuit:string = 'will-quit';
        /**
         * Emitted when the application is quitting. Note: On Windows, this event will not be emitted if the app is closed due to a shutdown/restart of the system or a user logout.
         * @event
         */
        export var Quit:string = 'quit';
        /**
         * Emitted when Electron has created a new session.
         * @event
         */
        export var SessionCreated:string = 'session-created';
        /**
         * This event will be emitted inside the primary instance of your application when a second instance has been executed. argv is an Array of the second instance's command line arguments, and workingDirectory is its current working directory. Usually applications respond to this by making their primary window focused and non-minimized. This event is guaranteed to be emitted after the ready event of app gets emitted.
         * @event
         */
        export var SecondInstance:string = 'second-instance';
          export function on(event: 'will-quit', listener: (event: Event) => void): ListenerHandle;
          export function on(event: 'quit', listener: (event: Event, exitCode: number) => void): ListenerHandle;
          export function on(event: 'session-created', listener: (event: Event, session: Session) => void): ListenerHandle;
          export function on(event: 'second-instance', listener: (argv: string[], workingDirectory: string) => void): ListenerHandle;
          export function on(event: string, listener: Function): ListenerHandle { return new ListenerHandle() }
          export function off(event: 'will-quit', handle: ListenerHandle): void;
          export function off(event: 'quit', handle: ListenerHandle): void;
          export function off(event: 'session-created', handle: ListenerHandle): void;
          export function off(event: 'second-instance', handle: ListenerHandle): void;
          export function off(event: string, handle: ListenerHandle): void { }

    }

    /**
     * Perform copy and paste operations on the system clipboard.
     */
    export namespace clipboard {

          /**
           * 
           */
          export function readText(type: string): string {
              console.log("readText");
              return "";
          }

          /**
           * Writes the text into the clipboard as plain text.
           */
          export function writeText(text: string, type: string) {
              console.log("writeText");
              
          }

          /**
           * 
           */
          export function readHTML(type: string): string {
              console.log("readHTML");
              return "";
          }

          /**
           * Writes markup to the clipboard.
           */
          export function writeHTML(markup: string, type: string) {
              console.log("writeHTML");
              
          }

          /**
           * 
           */
          export function readImage(type: string): NativeImage {
              console.log("readImage");
              return new NativeImage();
          }

          /**
           * Writes image to the clipboard.
           */
          export function writeImage(image: NativeImage, type: string) {
              console.log("writeImage");
              
          }

          /**
           * 
           */
          export function readRTF(type: string): string {
              console.log("readRTF");
              return "";
          }

          /**
           * Writes the text into the clipboard in RTF.
           */
          export function writeRTF(text: string, type: string) {
              console.log("writeRTF");
              
          }

          /**
           * Returns an Object containing title and url keys representing the bookmark in the clipboard. The title and url values will be empty strings when the bookmark is unavailable.
           */
          export function readBookmark(): {title: string, url: string} {
              console.log("readBookmark");
              return null;
          }

          /**
           * Writes the title and url into the clipboard as a bookmark. Note: Most apps on Windows don't support pasting bookmarks into them so you can use clipboard.write to write both a bookmark and fallback text to the clipboard.
           */
          export function writeBookmark(title: string, url: string, type: string) {
              console.log("writeBookmark");
              
          }

          /**
           * Clears the clipboard content.
           */
          export function clear(type: string) {
              console.log("clear");
              
          }

          /**
           * 
           */
          export function availableFormats(type: string): string[] {
              console.log("availableFormats");
              return [""];
          }



    }

    /**
     * Display native system dialogs for opening and saving files, alerting, etc.
     */
    export namespace dialog {

          /**
           * The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed or selected when you want to limit the user to a specific type. For example: The extensions array should contain extensions without wildcards or dots (e.g. 'png' is good but '.png' and '*.png' are bad). To show all files, use the '*' wildcard (no other wildcard is supported). If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filenames). Note: On Windows and Linux an open dialog can not be both a file selector and a directory selector, so if you set properties to ['openFile', 'openDirectory'] on these platforms, a directory selector will be shown.
           */
          export function showOpenDialog(options: {title: string, defaultPath: string, buttonLabel: string, filters: FileFilter[], properties: string[], message: string, securityScopedBookmarks: boolean}): {filePaths: string[], bookmarks: string[]} {
              console.log("showOpenDialog");
              return null;
          }

          /**
           * The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed, see dialog.showOpenDialog for an example. If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filename).
           */
          export function showSaveDialog(options: {title: string, defaultPath: string, buttonLabel: string, filters: FileFilter[], message: string, nameFieldLabel: string, showsTagField: boolean, securityScopedBookmarks: boolean}): {filename: string, bookmark: string} {
              console.log("showSaveDialog");
              return null;
          }

          /**
           * Shows a message box, it will block the process until the message box is closed. It returns the index of the clicked button. The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. If a callback is passed, the dialog will not block the process. The API call will be asynchronous and the result will be passed via callback(response).
           */
          export function showMessageBox(options: {type: string, buttons: string[], defaultId: number, title: string, message: string, detail: string, checkboxLabel: string, checkboxChecked: boolean, icon: NativeImage, cancelId: number, noLink: boolean, normalizeAccessKeys: boolean}): {response: number, checkboxChecked: boolean} {
              console.log("showMessageBox");
              return null;
          }

          /**
           * Displays a modal dialog that shows an error message. This API can be called safely before the ready event the app module emits, it is usually used to report errors in early stage of startup. If called before the app readyevent on Linux, the message will be emitted to stderr, and no GUI dialog will appear.
           */
          export function showErrorBox(title: string, content: string) {
              console.log("showErrorBox");
              
          }



    }

    /**
     * Detect keyboard events when the application does not have keyboard focus.
     */
    export namespace globalShortcut {

          /**
           * Registers a global shortcut of accelerator. The callback is called when the registered shortcut is pressed by the user. When the accelerator is already taken by other applications, this call will silently fail. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
           */
          export function register(accelerator: Accelerator, callback: () => void) {
              console.log("register");
              
          }

          /**
           * When the accelerator is already taken by other applications, this call will still return false. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
           */
          export function isRegistered(accelerator: Accelerator): boolean {
              console.log("isRegistered");
              return false;
          }

          /**
           * Unregisters the global shortcut of accelerator.
           */
          export function unregister(accelerator: Accelerator) {
              console.log("unregister");
              
          }



    }

    /**
     * In-app purchases on Mac App Store.
     */
    export namespace inAppPurchase {

          /**
           * You should listen for the transactions-updated event as soon as possible and certainly before you call purchaseProduct.
           */
          export function purchaseProduct(productID: string, quantity: number, callback: (isProductValid: boolean) => void) {
              console.log("purchaseProduct");
              
          }

          /**
           * Retrieves the product descriptions.
           */
          export function getProducts(productIDs: string[], callback: (products: Product[]) => void) {
              console.log("getProducts");
              
          }

          /**
           * 
           */
          export function canMakePayments(): boolean {
              console.log("canMakePayments");
              return false;
          }

          /**
           * 
           */
          export function getReceiptURL(): string {
              console.log("getReceiptURL");
              return "";
          }

          /**
           * Completes all pending transactions.
           */
          export function finishAllTransactions() {
              console.log("finishAllTransactions");
              
          }

          /**
           * Completes the pending transactions corresponding to the date.
           */
          export function finishTransactionByDate(date: string) {
              console.log("finishTransactionByDate");
              
          }


        /**
         * Emitted when one or more transactions have been updated.
         * @event
         */
        export var TransactionsUpdated:string = 'transactions-updated';
          export function on(event: 'transactions-updated', listener: (event: Event, transactions: Transaction[]) => void): ListenerHandle;
          export function on(event: string, listener: Function): ListenerHandle { return new ListenerHandle() }
          export function off(event: 'transactions-updated', handle: ListenerHandle): void;
          export function off(event: string, handle: ListenerHandle): void { }

    }

    /**
     * Logging network events.
     */
    export namespace netLog {

          /**
           * Starts recording network events to path.
           */
          export function startLogging(path: string) {
              console.log("startLogging");
              
          }

          /**
           * Stops recording network events. If not called, net logging will automatically end when app quits.
           */
          export function stopLogging(callback: (path: string) => void) {
              console.log("stopLogging");
              
          }



    }

    /**
     * Monitor power state changes.
     */
    export namespace powerMonitor {


        /**
         * Emitted when the system is suspending.
         * @event
         */
        export var Suspend:string = 'suspend';
        /**
         * Emitted when system is resuming.
         * @event
         */
        export var Resume:string = 'resume';
        /**
         * Emitted when the system changes to AC power.
         * @event
         */
        export var OnAc:string = 'on-ac';
        /**
         * Emitted when system changes to battery power.
         * @event
         */
        export var OnBattery:string = 'on-battery';
        /**
         * Emitted when the system is about to reboot or shut down. If the event handler invokes e.preventDefault(), Electron will attempt to delay system shutdown in order for the app to exit cleanly. If e.preventDefault() is called, the app should exit as soon as possible by calling something like app.quit().
         * @event
         */
        export var Shutdown:string = 'shutdown';
        /**
         * Emitted when the system is about to lock the screen.
         * @event
         */
        export var LockScreen:string = 'lock-screen';
        /**
         * Emitted as soon as the systems screen is unlocked.
         * @event
         */
        export var UnlockScreen:string = 'unlock-screen';
          export function on(event: 'suspend', listener: () => void): ListenerHandle;
          export function on(event: 'resume', listener: () => void): ListenerHandle;
          export function on(event: 'on-ac', listener: () => void): ListenerHandle;
          export function on(event: 'on-battery', listener: () => void): ListenerHandle;
          export function on(event: 'shutdown', listener: () => void): ListenerHandle;
          export function on(event: 'lock-screen', listener: () => void): ListenerHandle;
          export function on(event: 'unlock-screen', listener: () => void): ListenerHandle;
          export function on(event: string, listener: Function): ListenerHandle { return new ListenerHandle() }
          export function off(event: 'suspend', handle: ListenerHandle): void;
          export function off(event: 'resume', handle: ListenerHandle): void;
          export function off(event: 'on-ac', handle: ListenerHandle): void;
          export function off(event: 'on-battery', handle: ListenerHandle): void;
          export function off(event: 'shutdown', handle: ListenerHandle): void;
          export function off(event: 'lock-screen', handle: ListenerHandle): void;
          export function off(event: 'unlock-screen', handle: ListenerHandle): void;
          export function off(event: string, handle: ListenerHandle): void { }

    }

    /**
     * Extensions to process object.
     */
    export namespace process {

          /**
           * 
           */
          export function getCPUUsage(): CPUUsage {
              console.log("getCPUUsage");
              return new CPUUsage();
          }

          /**
           * Returns an object with V8 heap statistics. Note that all statistics are reported in Kilobytes.
           */
          export function getHeapStatistics(): {totalHeapSize: number, totalHeapSizeExecutable: number, totalPhysicalSize: number, totalAvailableSize: number, usedHeapSize: number, heapSizeLimit: number, mallocedMemory: number, peakMallocedMemory: number, doesZapGarbage: boolean} {
              console.log("getHeapStatistics");
              return null;
          }

          /**
           * Returns an object giving memory usage statistics about the current process. Note that all statistics are reported in Kilobytes.
           */
          export function getProcessMemoryInfo(): {workingSetSize: number, peakWorkingSetSize: number, privateBytes: number, sharedBytes: number} {
              console.log("getProcessMemoryInfo");
              return null;
          }

          /**
           * Returns an object giving memory usage statistics about the entire system. Note that all statistics are reported in Kilobytes.
           */
          export function getSystemMemoryInfo(): {total: number, free: number, swapTotal: number, swapFree: number} {
              console.log("getSystemMemoryInfo");
              return null;
          }



    }

    /**
     * Register a custom protocol and intercept existing protocol requests.
     */
    export namespace protocol {

          /**
           * A standard scheme adheres to what RFC 3986 calls generic URI syntax. For example http and https are standard schemes, while file is not. Registering a scheme as standard, will allow relative and absolute resources to be resolved correctly when served. Otherwise the scheme will behave like the file protocol, but without the ability to resolve relative URLs. For example when you load following page with custom protocol without registering it as standard scheme, the image will not be loaded because non-standard schemes can not recognize relative URLs: Registering a scheme as standard will allow access to files through the FileSystem API. Otherwise the renderer will throw a security error for the scheme. By default web storage apis (localStorage, sessionStorage, webSQL, indexedDB, cookies) are disabled for non standard schemes. So in general if you want to register a custom protocol to replace the http protocol, you have to register it as a standard scheme: Note: This method can only be used before the ready event of the app module gets emitted.
           */
          export function registerStandardSchemes(schemes: string[], options: {secure: boolean}) {
              console.log("registerStandardSchemes");
              
          }

          /**
           * Registers a protocol of scheme that will send the file as a response. The handler will be called with handler(request, callback) when a request is going to be created with scheme. completion will be called with completion(null) when scheme is successfully registered or completion(error) when failed. To handle the request, the callback should be called with either the file's path or an object that has a path property, e.g. callback(filePath) or callback({path: filePath}). When callback is called with nothing, a number, or an object that has an error property, the request will fail with the error number you specified. For the available error numbers you can use, please see the net error list. By default the scheme is treated like http:, which is parsed differently than protocols that follow the "generic URI syntax" like file:, so you probably want to call protocol.registerStandardSchemes to have your scheme treated as a standard scheme.
           */
          export function registerFileProtocol(scheme: string, handler: (request: {url: string, referrer: string, method: string, uploadData: UploadData[]}, callback: (filePath: string) => void) => void): {error: Error} {
              console.log("registerFileProtocol");
              return null;
          }

          /**
           * Registers a protocol of scheme that will send a String as a response. The usage is the same with registerFileProtocol, except that the callback should be called with either a String or an object that has the data, mimeType, and charset properties.
           */
          export function registerStringProtocol(scheme: string, handler: (request: {url: string, referrer: string, method: string, uploadData: UploadData[]}, callback: (data: string) => void) => void): {error: Error} {
              console.log("registerStringProtocol");
              return null;
          }

          /**
           * Registers a protocol of scheme that will send an HTTP request as a response. The usage is the same with registerFileProtocol, except that the callback should be called with a redirectRequest object that has the url, method, referrer, uploadData and session properties. By default the HTTP request will reuse the current session. If you want the request to have a different session you should set session to null. For POST requests the uploadData object must be provided.
           */
          export function registerHttpProtocol(scheme: string, handler: (request: {url: string, referrer: string, method: string, uploadData: UploadData[]}, callback: (redirectRequest: {url: string, method: string, session: {}, uploadData: {contentType: string, data: string}}) => void) => void): {error: Error} {
              console.log("registerHttpProtocol");
              return null;
          }

          /**
           * Unregisters the custom protocol of scheme.
           */
          export function unregisterProtocol(scheme: string): {error: Error} {
              console.log("unregisterProtocol");
              return null;
          }

          /**
           * The callback will be called with a boolean that indicates whether there is already a handler for scheme.
           */
          export function isProtocolHandled(scheme: string): {error: Error} {
              console.log("isProtocolHandled");
              return null;
          }



    }

    /**
     * Retrieve information about screen size, displays, cursor position, etc.
     */
    export namespace screen {

          /**
           * The current absolute position of the mouse pointer.
           */
          export function getCursorScreenPoint(): Point {
              console.log("getCursorScreenPoint");
              return new Point();
          }

          /**
           * 
           */
          export function getPrimaryDisplay(): Display {
              console.log("getPrimaryDisplay");
              return new Display();
          }

          /**
           * 
           */
          export function getAllDisplays(): Display[] {
              console.log("getAllDisplays");
              return new Array<Display>();
          }

          /**
           * 
           */
          export function getDisplayNearestPoint(point: Point): Display {
              console.log("getDisplayNearestPoint");
              return new Display();
          }

          /**
           * 
           */
          export function getDisplayMatching(rect: Rectangle): Display {
              console.log("getDisplayMatching");
              return new Display();
          }

          /**
           * Converts a screen physical point to a screen DIP point. The DPI scale is performed relative to the display containing the physical point.
           */
          export function screenToDipPoint(point: Point): Point {
              console.log("screenToDipPoint");
              return new Point();
          }

          /**
           * Converts a screen DIP point to a screen physical point. The DPI scale is performed relative to the display containing the DIP point.
           */
          export function dipToScreenPoint(point: Point): Point {
              console.log("dipToScreenPoint");
              return new Point();
          }

          /**
           * Converts a screen physical rect to a screen DIP rect. The DPI scale is performed relative to the display nearest to window. If window is null, scaling will be performed to the display nearest to rect.
           */
          export function screenToDipRect(window: BrowserWindow, rect: Rectangle): Rectangle {
              console.log("screenToDipRect");
              return new Rectangle();
          }

          /**
           * Converts a screen DIP rect to a screen physical rect. The DPI scale is performed relative to the display nearest to window. If window is null, scaling will be performed to the display nearest to rect.
           */
          export function dipToScreenRect(window: BrowserWindow, rect: Rectangle): Rectangle {
              console.log("dipToScreenRect");
              return new Rectangle();
          }


        /**
         * Emitted when newDisplay has been added.
         * @event
         */
        export var DisplayAdded:string = 'display-added';
        /**
         * Emitted when oldDisplay has been removed.
         * @event
         */
        export var DisplayRemoved:string = 'display-removed';
        /**
         * Emitted when one or more metrics change in a display. The changedMetrics is an array of strings that describe the changes. Possible changes are bounds, workArea, scaleFactor and rotation.
         * @event
         */
        export var DisplayMetricsChanged:string = 'display-metrics-changed';
          export function on(event: 'display-added', listener: (event: Event, newDisplay: Display) => void): ListenerHandle;
          export function on(event: 'display-removed', listener: (event: Event, oldDisplay: Display) => void): ListenerHandle;
          export function on(event: 'display-metrics-changed', listener: (event: Event, display: Display, changedMetrics: string[]) => void): ListenerHandle;
          export function on(event: string, listener: Function): ListenerHandle { return new ListenerHandle() }
          export function off(event: 'display-added', handle: ListenerHandle): void;
          export function off(event: 'display-removed', handle: ListenerHandle): void;
          export function off(event: 'display-metrics-changed', handle: ListenerHandle): void;
          export function off(event: string, handle: ListenerHandle): void { }

    }

    /**
     * Manage files and URLs using their default applications.
     */
    export namespace shell {

          /**
           * Show the given file in a file manager. If possible, select the file.
           */
          export function showItemInFolder(fullPath: string): boolean {
              console.log("showItemInFolder");
              return false;
          }

          /**
           * Open the given file in the desktop's default manner.
           */
          export function openItem(fullPath: string): boolean {
              console.log("openItem");
              return false;
          }

          /**
           * Open the given external protocol URL in the desktop's default manner. (For example, mailto: URLs in the user's default mail agent).
           */
          export function openExternal(url: string, options: {activate: boolean}): {error: Error} {
              console.log("openExternal");
              return null;
          }

          /**
           * Move the given file to trash and returns a boolean status for the operation.
           */
          export function moveItemToTrash(fullPath: string): boolean {
              console.log("moveItemToTrash");
              return false;
          }

          /**
           * Play the beep sound.
           */
          export function beep() {
              console.log("beep");
              
          }

          /**
           * Creates or updates a shortcut link at shortcutPath.
           */
          export function writeShortcutLink(shortcutPath: string, operation: string, options: ShortcutDetails): boolean {
              console.log("writeShortcutLink");
              return false;
          }

          /**
           * Resolves the shortcut link at shortcutPath. An exception will be thrown when any error happens.
           */
          export function readShortcutLink(shortcutPath: string): ShortcutDetails {
              console.log("readShortcutLink");
              return new ShortcutDetails();
          }



    }

    /**
     * 
     */
    export namespace menu {

          /**
           * Sets menu as the application menu on macOS. On Windows and Linux, the menu will be set as each window's top menu. Passing null will remove the menu bar on Windows and Linux but has no effect on macOS. Note: This API has to be called after the ready event of app module.
           */
          export function setApplicationMenu(menu: Menu) {
              console.log("setApplicationMenu");
              
          }

          /**
           * Note: The returned Menu instance doesn't support dynamic addition or removal of menu items. Instance properties can still be dynamically modified.
           */
          export function getApplicationMenu(): Menu {
              console.log("getApplicationMenu");
              return new Menu();
          }

          /**
           * Sends the action to the first responder of application. This is used for emulating default macOS menu behaviors. Usually you would use the role property of a MenuItem. See the macOS Cocoa Event Handling Guide for more information on macOS' native actions.
           */
          export function sendActionToFirstResponder(action: string) {
              console.log("sendActionToFirstResponder");
              
          }

          /**
           * Generally, the template is an array of options for constructing a MenuItem. The usage can be referenced above. You can also attach other fields to the element of the template and they will become properties of the constructed menu items.
           */
          export function buildFromTemplate(template: NewMenuItem[]): Menu {
              console.log("buildFromTemplate");
              return new Menu();
          }

          /**
           * 
           */
          export function make(): Menu {
              console.log("make");
              return new Menu();
          }

          /**
           * 
           */
          export function ref(handle: string): Menu {
              console.log("ref");
              return new Menu();
          }



    }

    /**
     * 
     */
    export namespace menu.item {

          /**
           * 
           */
          export function make(options: NewMenuItem): MenuItem {
              console.log("make");
              return new MenuItem();
          }

          /**
           * 
           */
          export function ref(handle: string): MenuItem {
              console.log("ref");
              return new MenuItem();
          }



    }

    /**
     * 
     */
    export namespace app.dock {

          /**
           * When critical is passed, the dock icon will bounce until either the application becomes active or the request is canceled. When informational is passed, the dock icon will bounce for one second. However, the request remains active until either the application becomes active or the request is canceled.
           */
          export function bounce(type: string): number {
              console.log("bounce");
              return 0;
          }

          /**
           * Cancel the bounce of id.
           */
          export function cancelBounce(id: number) {
              console.log("cancelBounce");
              
          }

          /**
           * Bounces the Downloads stack if the filePath is inside the Downloads folder.
           */
          export function downloadFinished(filePath: string) {
              console.log("downloadFinished");
              
          }

          /**
           * Sets the string to be displayed in the dock’s badging area.
           */
          export function setBadge(text: string) {
              console.log("setBadge");
              
          }

          /**
           * 
           */
          export function getBadge(): string {
              console.log("getBadge");
              return "";
          }

          /**
           * Hides the dock icon.
           */
          export function hide() {
              console.log("hide");
              
          }

          /**
           * Shows the dock icon.
           */
          export function show() {
              console.log("show");
              
          }

          /**
           * 
           */
          export function isVisible(): boolean {
              console.log("isVisible");
              return false;
          }

          /**
           * Sets the application's dock menu.
           */
          export function setMenu(menu: Menu) {
              console.log("setMenu");
              
          }



    }


}

export default api;