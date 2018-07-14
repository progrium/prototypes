const { spawn, spawnSync } = require('child_process');
const fs = require("fs");
const ejs = require("ejs");

var electronDev = null;

process.on("SIGINT", () => {
    if (electronDev) {
        spawnSync("kill", ["-9", "-"+electronDev.pid]);
    }
});



(function() {
    var schema = JSON.parse(fs.readFileSync("./schema/electron-api.json"));
    var transform = require("./schema/transform");
    var api = transform(schema);
    var helpers = require("./schema/helpers");
    var template = fs.readFileSync("./schema/template.ejs");
    fs.writeFileSync("./src/main/api.ts", ejs.render(template.toString('ascii'), {
        api: api,
        helpers: helpers,
    }));
    console.log("Compiling TypeScript...");
    var tsc = spawnSync('tsc');
    console.log(tsc.stdout.toString('ascii'));
    console.log(tsc.stderr.toString('ascii'));
    if (tsc.status != 0) {
        return;
    }
    electronDev = spawn('./node_modules/.bin/electron', ['./dist/main.js']);
    electronDev.stdout.on('data', (data) => {
        console.log(data.toString());
    });
    electronDev.stderr.on('data', (data) => {
        console.log(data.toString());
    });
    electronDev.on('exit', (code) => {
        console.log(`Child exited with code ${code}`);
    });
})();