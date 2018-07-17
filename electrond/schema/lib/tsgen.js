
module.exports = class Generator {
    constructor() {
        this.stack = []
        this.stmts = []
    }

    decl(...args) {
        if (typeof args[args.length-1] == "function") {
            var blockFn = args.pop()
            this.stmts.push(`${args.join(" ")} {\n${this.block(blockFn)}\n}`)
        } else {
            var line = args.join(" ")
            var end = (line.startsWith("/") || line.endsWith("}")) ? "" : ";"
            this.stmts.push(line+end)
        }
    }

    inline() {
        var stmt = this.stmts.pop()
        return (stmt.endsWith(";")) ? stmt.slice(0, -1) : stmt
    }

    export() {
        this.decl("export", this.inline())
    }

    async() {
        this.decl("async", this.inline())
    }

    func(name, args, type, blockFn) {
        this.decl("function", `${name}${this.signature(args, type)}`, blockFn)
    }

    call(name, ...args) {
        this.decl(`${name}(${(args||[]).join(", ")})`)
    }

    comment(str) {
        this.decl(`// ${str}`)
    }

    commentBlock(str) {
        this.decl("/**\n * "+(str||"").split("\n").join("\n * ")+"\n */")
    }

    block(blockFn) {
        this.stack.push(this.stmts)
        this.stmts = []
        blockFn(this)
        var blockStmts = this.stmts
        this.stmts = this.stack.pop()
        return blockStmts.join("\n")
    }

    signature(args, type) {
        return `(${(args||[]).join(", ")})${(type) ? `: ${type} ` : ""}`
    }

    str(v) {
        return `"${v}"`
    }

    arr(...args) {
        return `[${args.join(", ")}]`
    }

    idx(name, key) {
        return `${name}[${key}]`
    }

    chain(...args) {
        return args.join(".")
    }

    lambda(...args) {
        var async = ""
        if (args[0] === "async") {
            async = args.shift()+" "
        }
        var blockFn = args.pop()
        var type = args.pop()
        return `${async}${this.signature(args, type)} => {\n${this.block(blockFn)}\n}`
    }

    var(name, type, value) {
        return (value) ? `${name}: ${type} = ${value}` : `${name}: ${type}`
    }

    toString() {
        return this.stmts.join("\n")
    }
}
