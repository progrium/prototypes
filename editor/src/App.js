import React, { Component } from 'react';
import MonacoEditor from 'react-monaco-editor';
import { Badge, Avatar, Row, Col, Layout, Dropdown, Icon, Menu } from 'antd';
import { Console } from 'console-feed'

// import { XTerm } from 'react-xterm';
import './App.css';
import 'antd/dist/antd.css';
import 'xterm/src/xterm.css';
import SplitterLayout from 'react-splitter-layout';
import { Tabs } from 'antd';

import { Inspector } from './Inspector';
import { Hierarchy } from './Hierarchy';

const TabPane = Tabs.TabPane;

const qmux = window.qmux;
const qrpc = window.qrpc;


class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      remote: {
        components: [],
        projects: [],
        hierarchy: [],
        nodes: {},
        nodePaths: {},
        currentProject: null
      },
      files: [],
      activeTab: "-1",
      inspectNode: null,
      logs: []
    }

    this.api = new qrpc.API();
    this.api.handle("state", this);
  }

  async serveRPC(r, c) {
    var obj = await c.decode();
    this.setState({"remote": obj});
    r.return();
  }

  async connectConsole() {
    var conn = await qmux.DialWebsocket("ws://localhost:4242");
    var session = new qmux.Session(conn);
    var client = new qrpc.Client(session);
    var resp = await client.call("console");
    if (resp.hijacked === true) {
      this.readLogs(resp.channel);
    } else {
      console.log("log stream not hijacked");
    }
  }

  async connectRepl() {
    var conn = await qmux.DialWebsocket("ws://localhost:4243");
    var session = new qmux.Session(conn);
    var client = new qrpc.Client(session);
    var resp = await client.call("repl");
    if (resp.hijacked === true) {
      this.repl = resp.channel;
      this.readLogs(resp.channel);
    } else {
      console.log("repl stream not hijacked");
    }
  }

  async connectServer() {
    try {
      var conn = await qmux.DialWebsocket("ws://localhost:4243");
    } catch (e) {
      setTimeout(() => {
        this.connectServer();
      }, 200);
      return;
    }
    conn.socket.onclose = () => {
      conn.close();
      setTimeout(() => {
        this.connectServer();
      }, 200);
    }
    //console.clear();
    var session = new qmux.Session(conn);
    var client = new qrpc.Client(session, this.api);
    client.serveAPI();
    window.rpc = client;
    await client.call("subscribe");
  }

  async componentDidMount() {
    this.connectConsole();
    this.connectRepl();    
    this.connectServer();

    this.log({"Hello": "world"});
    this.prompt("Hello world");

  }

  prompt() {
    this.setState(({ logs }) => ({ logs: [...logs, {method: "command", data: Array.prototype.slice.call(arguments)}] }))
  }

  log() {
    this.setState(({ logs }) => ({ logs: [...logs, {method: "log", data: Array.prototype.slice.call(arguments)}] }))
  }

  async readLogs(ch) {
    var linebuf = new Buffer.from([]);
    while (true) {
      var data = await ch.read(1);
      if (data === undefined) {
          this.log("|Server got EOF");
          break;
      }
      if (data.toString('ascii') === "\n") {
        this.log(linebuf.toString('ascii'))
        //this.refs.xterm.writeln();
        linebuf = new Buffer.from([]);
      } else {
        linebuf = Buffer.concat([linebuf, data]);
      }
    }
  }

  

  consoleDidMount(editor, monaco) {
    editor.addCommand(monaco.KeyCode.Enter, () => {
      var line = editor.getValue();
      this.repl.write(Buffer.from(line+"\n"));
      this.prompt(line);
      editor.getModel().setValue("");
    });
  }

  editorDidMount(editor, monaco) {
    editor.focus();
    editor.addAction({
      id: 'save',
      label: 'Save',
      keybindings: [
        monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_S,
      ],  
      run: (ed) => {
        if (this.state.editor.id === "") {
          return;
        }
        window.rpc.call("writeDelegate", {"ID": this.state.editor.id, "Contents": ed.getValue()});
        return null;
      }
    });
  }

  onSelect(selectedKeys, info) {
    this.setState({"inspectNode": selectedKeys[0]});
  }

  onEditorChange() {
    
  }

  openDelegate(delegate) {
    this.editFile({
      name: "Delegate",
      contents: delegate.Contents
    });
  }

  editFile(file) {
    this.setState((state, props) => {
      var files = state.files.concat(file);
      return {
        files: files,
        activeTab: (files.length-1).toString()
      }
    });
  }

  closeFile(idx) {
    var files = this.state.files;
    files.splice(idx, 1);
    this.setState({
      "files": files,
      "activeTab": "-1"
    });
  }

  onReload() {
    window.rpc.call("reload");
  }

  async onAddNode() {
    let name = prompt("Node Name");
    await window.rpc.call("appendNode", {"ID": "", "Name": name});
  }

  onTabClick(key) {
    this.setState({
      "activeTab": key
    });
  }

  render() {
    return (
      <div style={{height: "100%"}}>
        {/* <Header 
          projects={this.state.remote.projects} 
          currentProject={this.state.remote.currentProject} /> */}
        <Layout style={{height: "100%"}}>
          <Layout.Sider width="62" style={{ background: '#fff'}}>
            <Menu className="projects" style={{width: "62px", height: "100%"}} mode="inline" defaultSelectedKeys={['1']}>
              <Menu.Item key="1"><Badge status="success" offset={[-35,35]}><Avatar size="large">U</Avatar></Badge></Menu.Item>
              <Menu.Item key="2"><Badge status="success" offset={[-35,35]}><Avatar size="large">U</Avatar></Badge></Menu.Item>
            </Menu>
          </Layout.Sider>
          <Layout.Content>
          <SplitterLayout 
            customClassName="tree" 
            primaryIndex={1} 
            secondaryMinSize={150} 
            secondaryInitialSize={200}>
            <Hierarchy
              onAddNode={this.onAddNode}
              onReload={this.onReload}
              hierarchy={this.state.remote.hierarchy}
              nodePaths={this.state.remote.nodePaths}
              onSelect={this.onSelect.bind(this)} />
            <SplitterLayout 
              customClassName="tree" 
              primaryIndex={1} 
              secondaryMinSize={250} 
              secondaryInitialSize={250}>
              <Inspector 
                openDelegate={this.openDelegate.bind(this)}
                id={this.state.inspectNode}
                nodePaths={this.state.remote.nodePaths}
                hierarchy={this.state.remote.hierarchy}  
                node={this.state.remote.nodes[this.state.inspectNode]} 
                components={this.state.remote.components} />
              {/* <SplitterLayout vertical style={{"width": "100%"}}> */}
                <Tabs activeKey={this.state.activeTab} size="small" onTabClick={this.onTabClick.bind(this)}>
                  <TabPane tab="Console" key="-1" style={{backgroundColor: "#242424"}}>
                    <Console logs={this.state.logs} variant="dark" />
                    <div data-method="command" className="css-y965hj"><div className="css-1wcpu1t"></div><div className="css-3tmnuj">
                      <div data-type="string" className="css-c0sqg9" style={{"height": "20px"}}>
                        <MonacoEditor
                            language="javascript"
                            theme="vs-dark"
                            width="800px"
                            height="20px"
                            // value={this.state.editor.contents}
                            options={{
                              selectOnLineNumbers: true,
                              automaticLayout: true,
                              lineNumbers: 'off',
                              glyphMargin: false,
                              folding: false,
                              // Undocumented see https://github.com/Microsoft/vscode/issues/30795#issuecomment-410998882
                              lineDecorationsWidth: 0,
                              lineNumbersMinChars: 0,
                              wordWrap: "off",
                              highlightActiveIndentGuide: false,
                              renderLineHighlight: "none",
                              scrollbar: {vertical: "hidden", horizontal: "hidden"},
                              minimap: {enabled: false}
                            }}
                            // onChange={this.onEditorChange.bind(this)}
                            editorDidMount={this.consoleDidMount.bind(this)}
                          />
                      </div>
                    </div></div>
                  </TabPane>
                  {(this.state.files||[]).map((file, idx) => {
                    var onClick = (evt) => {
                      this.closeFile(idx);
                      evt.stopPropagation();
                    }
                    var tabTitle = <div><span style={{marginRight: "8px"}}>{file.name}</span> <Icon style={{marginRight: "-12px"}} onClick={onClick} type="close" /></div>
                    return (
                    <TabPane tab={tabTitle} key={idx}>
                      <div style={{"height": "100%"}}>
                        <MonacoEditor
                          language="go"
                          theme="vs-light"
                          value={file.contents}
                          options={{
                            selectOnLineNumbers: true,
                            automaticLayout: true
                          }}
                          onChange={this.onEditorChange.bind(this)}
                          editorDidMount={this.editorDidMount.bind(this)}
                        />
                      </div>
                    </TabPane>)
                  })}
                </Tabs>
              {/* </SplitterLayout> */}
            </SplitterLayout>
          </SplitterLayout>
          </Layout.Content>
        </Layout>
      </div>
    );
  }
}

export default App;

function createProjectMenu(projects) {
  const onClick = (item, key, keyPath) => {
    window.rpc.call("selectProject", item.key);
  }
  return <Menu onClick={onClick}>
    {(projects||[]).map((project) => {
      return <Menu.Item key={project.name}>{project.name}</Menu.Item>
    })}
  </Menu> 
}

function ProjectSelector(props) {
  let project = props.projects.find((el) => {
    return el.name === props.currentProject;
  })
  project = project || props.projects[0] || {name: "", path: ""};
  return <Dropdown overlay={createProjectMenu(props.projects)} trigger={['click']}>
        <h3 style={{margin: "12px 0px 2px 20px", fontSize: "22px", lineHeight: "22px", display: "block", width: "400px"}}>
          {project.name} <Icon style={{fontSize: "18px"}} type="caret-down" />
          <span style={{display: "block", color: "gray", fontSize: "12px"}}>{project.path}</span>
        </h3>
      </Dropdown>
}

function Header(props) {
  return <header id="header" className="clearfix">
    <Row>
      <Col span={12} className="ant-menu-horizontal" style={{paddingBottom: "7px"}}>
        <ProjectSelector projects={props.projects} currentProject={props.currentProject} />
      </Col>
      <Col span={12}>
      <Menu
        mode="horizontal"
        defaultSelectedKeys={['1']}
        style={{lineHeight: '64px'}}
      >
        <Menu.Item style={{float: 'right'}} key="0"></Menu.Item>
        <Menu.Item style={{float: 'right'}} key="1">progrium</Menu.Item>
      </Menu>
      </Col>
    </Row>
  </header>
}
