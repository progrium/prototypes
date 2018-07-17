"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function argX(arg) {
    if (arg == null)
        return null;
    switch (arg.constructor.name) {
        case "NativeImage":
            return arg.toDataURL();
        default:
            return arg;
    }
}
exports.argX = argX;
//# sourceMappingURL=util.js.map