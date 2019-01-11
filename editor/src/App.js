import React, { Component } from 'react';
import MonacoEditor from 'react-monaco-editor';
import { Row, Col, Layout, InputNumber, Dropdown, Checkbox, TreeSelect, Icon, Button, Tree, Input, Menu, Collapse } from 'antd';
import _ from 'underscore';
import { XTerm } from 'react-xterm';
import './App.css';
import 'antd/dist/antd.css';
import 'xterm/src/xterm.css';
import SplitterLayout from 'react-splitter-layout';
import { Tabs } from 'antd';

const TabPane = Tabs.TabPane;

const DirectoryTree = Tree.DirectoryTree;
const { TreeNode } = Tree;
const Panel = Collapse.Panel;
const TreeSelectNode = TreeSelect.TreeNode;

const qmux = window.qmux;
const qrpc = window.qrpc;

class FieldInput extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expressionMode: false
    };
    this.toggleExpressionMode = this.toggleExpressionMode.bind(this);
  }

  toggleExpressionMode() {
    this.setState({expressionMode: !this.state.expressionMode});
  }

  render() {
    var onChange = () => {}
    if (this.state.expressionMode) {
      return <div onDoubleClick={this.toggleExpressionMode}>
        <Input size="small" style={{color:"white", backgroundColor: "#555", fontFamily:"monospace"}} onChange={onChange} onDoubleClick={this.toggleExpressionMode} value={this.props.expression || this.props.value} />
      </div>
    }
    switch (this.props.type) {
      case "boolean":
        return <div onDoubleClick={this.toggleExpressionMode}><Checkbox checked={this.props.value} /></div>;
      case "string":
        return <Input size="small" onChange={onChange} onDoubleClick={this.toggleExpressionMode} value={this.props.value} />;
      case "number":
        return <div onDoubleClick={this.toggleExpressionMode}><InputNumber style={{width: "100%"}} size="small" value={this.props.value} /></div>;
      case "reference":
        return <div onDoubleClick={this.toggleExpressionMode}><TreeSelect
          size="small"
          showSearch
          style={{ width: "100%"}}
          dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
          placeholder="Please select"
          value="/server/server.go"
          allowClear
          treeDefaultExpandAll
        >
          {refTree}
        </TreeSelect></div>;
      default:
        return "???";
    }
  }
}

function LabeledField(props) {
  const onClick = () => {
    console.log("Click");
  }
  return (<Row>
      <Col span={9} style={{fontSize: "smaller", height: "25px", marginTop: "2px"}}>
        <span onClick={onClick}>{props.label}</span>
      </Col>
      <Col span={15}>
        {props.children}
      </Col>
    </Row>);
}

function KeyedField(props) {
  var onChange = () => {}
  return (<Row>
    <Col span={9} style={{fontSize: "smaller", height: "25px"}}>
    <Input size="small" onChange={onChange} value={props.name} />
    </Col>
    <Col span={15}>
      {props.children}
    </Col>
  </Row>);
}

function PropField(props) {
  switch (props.type) {
    case "boolean":
    case "string":
    case "number":
    case "reference":
      return <LabeledField key={props.eventKey} label={props.name}><FieldInput {...props} /></LabeledField>
    case "struct":
      return (
        <EmbeddedFields key={props.eventKey} label={props.name} fields={props.fields}>
          {props.children}
        </EmbeddedFields>
      );
    case "map":
      return (
        <MapFields key={props.eventKey} label={props.name} fields={props.fields}>
          {props.children}
        </MapFields>
      );
    case "array":
      return (
        <ArrayFields key={props.eventKey} label={props.name} fields={props.fields}>
          {props.children}
        </ArrayFields>
      );
    default:
      return <LabeledField key={props.eventKey} label={props.name}><FieldInput {...props} /></LabeledField>              
  }  
}

function EmbeddedFields(props) {
  let fields = props.fields || []
  return (
    <Collapse className="embedded" bordered={false} defaultActiveKey="1">
        <Panel header={props.label} key="1">
          {fields.map((p) => {
            return <PropField {...p} />
          })}
          {props.children}
        </Panel>
    </Collapse>
  );
}

function MapFields(props) {
  let fields = props.fields || []
  return (
    <Collapse className="embedded" bordered={false} defaultActiveKey="1">
        <Panel header={props.label} key="1">
          {fields.map((field) => {
            return <KeyedField name={field.name}><FieldInput {...field} /></KeyedField>
          })}
          {props.children}
        </Panel>
    </Collapse>
  );
}

function ArrayFields(props) {
  let fields = props.fields || []
  return (
    <Collapse className="embedded" bordered={false} defaultActiveKey="1">
        <Panel header={props.label} key="1">
          <LabeledField key="count" label="Count"><InputNumber style={{width: "100%"}} size="small" value={fields.length} /></LabeledField>
          {fields.map((field, idx) => {
            field.name = "Element "+idx;
            return field
          }).map((p) => {
            return <PropField {...p} />
          })}
          {props.children}
        </Panel>
    </Collapse>
  );
}

function PropFields(props) {
  return props.fields.map((el, idx) => {
    el.key = idx;
    return <PropField {...el} />
  });
}

function Buttons(props) {
  return (props.buttons||[]).map((button) => {
    return <Button size="small" style={{marginTop: "10px", width: "100%"}}>
          {button.name}
      </Button>
  });
}

// const demoMenu = (
//   <Menu>
//     <Menu.Item key="1"><Icon type="user" />1st menu item</Menu.Item>
//     <Menu.Item key="2"><Icon type="user" />2nd menu item</Menu.Item>
//     <Menu.Item key="3"><Icon type="user" />3rd item</Menu.Item>
//   </Menu>
// );

function createComponentMenu(components) {
  return <Menu>
    {(components||[]).map((name) => {
      return <Menu.Item key={name}>{name}</Menu.Item>
    })}
  </Menu> 
}

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

function arrangeIntoTree(paths) {
  var tree = [];

  // This example uses the underscore.js library.
  _.each(paths, function(path) {

      var pathParts = path.split('/');
      pathParts.shift(); // Remove first blank element from the parts array.

      var currentLevel = tree; // initialize currentLevel to root
      var currentPath = "";

      _.each(pathParts, function(part) {
          currentPath = currentPath+"/"+part;

          // check to see if the path already exists.
          var existingPath = _.findWhere(currentLevel, {
              name: part
          });

          if (existingPath) {
              // The path to this item was already in the tree, so don't add it again.
              // Set the current level to this path's children
              currentLevel = existingPath.children;
          } else {
              var newPart = {
                  name: part,
                  path: currentPath,
                  children: [],
              }

              currentLevel.push(newPart);
              currentLevel = newPart.children;
          }
      });
  });

  return tree;
}

class MyTreeNode extends TreeNode {
  render() {
    var onClick = (item, key, keyPath) => {
      console.log(this.props.eventKey);
    }
    const contextMenu = (
      <Menu onClick={onClick}>
        <Menu.Item key="1">Delete</Menu.Item>
        <Menu.Item key="2">Duplicate</Menu.Item>
      </Menu>
    )
    return <Dropdown overlay={contextMenu} trigger={['contextMenu']}>{super.render()}</Dropdown>;
  }
}

function convertNodeToTree(node) {
  var isLeaf = node.children.length === 0;
  return <MyTreeNode title={node.name} key={node.path} isLeaf={isLeaf}>{node.children.map((n) => { return convertNodeToTree(n) })}</MyTreeNode>
}

function convertNodeToSelectTree(node) {
  return <TreeSelectNode title={node.name} value={node.path} key={node.path}>{node.children.map((n) => { return convertNodeToSelectTree(n) })}</TreeSelectNode>
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

function Inspector(props) {
  var onChange = () => {}
  return <div>
    <Row style={{padding: "5px"}}> 
      <Col span={2}><Checkbox checked={props.node.active} /></Col>
      <Col span={22}><Input size="small" value={props.name} onChange={onChange} /></Col>
    </Row>
    <Collapse bordered={false} defaultActiveKey={['1']}>
      {props.node.components.map((com, idx) => {
        return <Panel header={com.name} key={idx}>
          <PropFields fields={com.fields} />
          <Buttons buttons={com.buttons} />
        </Panel>
      })}
    </Collapse>
    <div style={{margin: "20px"}}>
      <Dropdown overlay={createComponentMenu(props.components)}>
        <Button style={{width: "100%"}}>
          Add Component <Icon type="down" />
        </Button>
      </Dropdown>
    </div>
  </div>
}

let refTree = null;

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      remote: {
        components: [],
        projects: [],
        hierarchy: {},
        currentProject: null
      },
      inspectNode: null
    }
  }

  componentDidMount() {
    var api = new qrpc.API();
    const setState = this.setState.bind(this);
    api.handle("state", {
      "serveRPC": async (r, c) => {
        var obj = await c.decode();
        console.log(obj);
        setState({"remote": obj});
      }
    });
    (async () => {
        var conn = await qmux.DialWebsocket("ws://localhost:4242");
        var session = new qmux.Session(conn);
        var client = new qrpc.Client(session, api);
        client.serveAPI();
        window.rpc = client;
        await client.call("subscribe");
    
    })().catch(async (err) => { 
        console.log(err.stack);
    });
  }

  editorDidMount(editor, monaco) {
    editor.focus();
    this.editor = editor;
  }

  onChange(newValue, e) {
    //window.rpc.call("increment");
  }

  onSelect(selectedKeys, info) {
    this.setState({"inspectNode": selectedKeys[0]});
    //var data = await window.rpc.call("readFile", selectedKeys[0]);
    //this.editor.getModel().setValue(data);
  }


  render() {
    const options = {
      selectOnLineNumbers: true,
      automaticLayout: true
    };
    const nodeTree = arrangeIntoTree(this.state.remote.hierarchy);
    const treeNodes = nodeTree.map(convertNodeToTree);
    refTree = nodeTree.map(convertNodeToSelectTree);
    const self = this;
    let inspector = <div></div>
    if (this.state.inspectNode !== null) {
      let parts = this.state.inspectNode.split("/");
      let name = parts[parts.length-1];
      inspector = <Inspector name={name} node={this.state.remote.nodes[this.state.inspectNode]} components={this.state.remote.components} />
    }
    return (
      <div style={{height: "100%"}}>
        <Header projects={this.state.remote.projects} currentProject={this.state.remote.currentProject} />
        <Layout style={{height: "100%"}}>
          <SplitterLayout customClassName="tree" primaryIndex={1} secondaryMinSize={150} secondaryInitialSize={200}>
            <DirectoryTree
              defaultExpandAll
              onSelect={this.onSelect.bind(this)}
            >
              {treeNodes}
            </DirectoryTree>
            
            <SplitterLayout customClassName="tree" primaryIndex={1} secondaryMinSize={250} secondaryInitialSize={250}>
              {inspector}

              <SplitterLayout vertical style={{"width": "100%"}}>
                <Tabs defaultActiveKey="1" size="small">
                  <TabPane tab="Edit" key="1">
                    <div style={{"height": "100%"}}>
                      <MonacoEditor
                        language="javascript"
                        theme="vs-light"
                        value=""
                        options={options}
                        onChange={self.onChange}
                        editorDidMount={self.editorDidMount.bind(self)}
                      />
                    </div>
                  </TabPane>
                </Tabs>
                
                <div>
                  <Tabs defaultActiveKey="1" size="small">
                    <TabPane tab="Console" key="1">
                      <XTerm ref='xterm' style={{
                        addons:['fit', 'fullscreen'],
                        overflow: 'hidden',
                        position: 'relative',
                        width: '100%',
                        height: '100%'
                      }} />
                    </TabPane>
                  </Tabs>
                </div>
              </SplitterLayout>

            </SplitterLayout>
          </SplitterLayout>
        </Layout>
      </div>
    );
  }
}

export default App;
