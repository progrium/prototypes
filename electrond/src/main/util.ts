export function argX(arg: any): any {
    if (arg == null) return null;
    switch (arg.constructor.name) {
        case "NativeImage":
            return arg.toDataURL();
        default:
            return arg;
    }
}