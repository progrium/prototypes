
module.exports = class Generator {
    constructor() {
        this.stack = []
        this.stmts = []
    }

    decl(...args) {
        var blockFn
        if (typeof args[args.length-1] == "function") {
            blockFn = args.pop()
            this.stack.push(this.stmts)
            this.stmts = []
            blockFn(this)
            var blockStmts = this.stmts
            this.stmts = this.stack.pop()
            this.stmts.push(args.join(" ")+" {\n"+blockStmts.join("\n")+"\n}\n")
        } else {
            var semi = ";"
            if (args[0].startsWith("/")) semi = ""
            this.stmts.push(args.join(" ")+semi)
        }
    }

    func(name, args, type, blockFn) {
        if (type) {
            type = `: ${type} `
        } else {
            type = ""
        }
        this.decl("function", `${name}(${(args||[]).join(", ")})`+type, blockFn)
    }

    comment(str) {
        this.decl(`// ${str}`)
    }

    commentBlock(str) {
        this.decl("/**\n * "+(str||"").split("\n").join("\n * ")+"\n */")
    }

    str(v) {
        return `"${v}"`
    }

    var(name, type, value) {
        if (value) {
            return `${name}: ${type} = ${value}`
        }
        return `${name}: ${type}`
    }

    toString() {
        return this.stmts.join("\n")
    }
}
