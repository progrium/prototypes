Sidecar supervisor made to run [plugin] daemons alongside your main process.

- scripts/executables live in a directory
- watches the directory
- starts all executables in dir as named daemons
- a log mode allows their logs to aggregate into the parent stdout/err
- on unexpected exit, it'll try to restart several times and then leave it
    - write the recent log output to file
- on directory changes, it'll reload or start new daemons


- testing with exec.Cmd
- watching files: https://github.com/fsnotify/fsnotify
- line writer
- circbuf for recent log output on fail
- restart policies: imediate w/ 3 retries
- passing signals along (forego)

type Supervisor struct {
    sidecars map[string]*Process
    output *LineWriter
}
ReloadAll
Reload(path)
    checks that its executable
Logs(path)
Stop(path)
StopAll()

type Process struct {
    cmd *Cmd
    logs *circbuf.Buffer
    restarts int
    name string
}

type LineWriter struct {
    output io.Writer
}
ReadLines
ErrorOutput

read dir, find executables
reload executable
