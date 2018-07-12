const libmux = require("./libmux");

(async () => {
    await libmux.TestError("Hello world");
})();