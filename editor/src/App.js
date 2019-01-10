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

function FieldInput(props) {
  var onChange = () => {}
  switch (props.type) {
    case "boolean":
      return <Checkbox checked={props.value} />;
    case "string":
      return <Input size="small" style={{color:"white", backgroundColor: "#555", fontFamily:"monospace"}} onChange={onChange} value={props.value} />;
    case "number":
      return <InputNumber style={{width: "100%"}} size="small" value={props.value} />;
    case "reference":
      return <TreeSelect
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
      </TreeSelect>;
    default:
      return "???";
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
      return <LabeledField key={props.key} label={props.name}>{FieldInput(props)}</LabeledField>
    case "struct":
      return (
        <EmbeddedFields key={props.key} label={props.name} fields={props.fields}>
          {props.children}
        </EmbeddedFields>
      );
    case "map":
      return (
        <MapFields key={props.key} label={props.name} fields={props.fields}>
          {props.children}
        </MapFields>
      );
    case "array":
      return (
        <ArrayFields key={props.key} label={props.name} fields={props.fields}>
          {props.children}
        </ArrayFields>
      );
    default:
      return <LabeledField key={props.key} label={props.name}>{FieldInput(props)}</LabeledField>              
  }
}

function EmbeddedFields(props) {
  let fields = props.fields || []
  return (
    <Collapse className="embedded" bordered={false} defaultActiveKey="1">
        <Panel header={props.label} key="1">
          {fields.map(PropField)}
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
            return <KeyedField name={field.name}>{FieldInput(field)}</KeyedField>
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
          <PropField type="number" name="Count" value={fields.length} />
          {fields.map((field, idx) => {
            field.name = "Element "+idx;
            return field
          }).map(PropField)}
          {props.children}
        </Panel>
    </Collapse>
  );
}

function PropFields(props) {
  return props.fields.map((el, idx) => {
    el.key = idx;
    return PropField(el);
  });
}



const menu = (
  <Menu>
    <Menu.Item key="1"><Icon type="user" />1st menu item</Menu.Item>
    <Menu.Item key="2"><Icon type="user" />2nd menu item</Menu.Item>
    <Menu.Item key="3"><Icon type="user" />3rd item</Menu.Item>
  </Menu>
);

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

function Header(props) {
  return <header id="header" className="clearfix">
    <Row>
      <Menu
        mode="horizontal"
        defaultSelectedKeys={['1']}
        style={{lineHeight: '64px'}}
      >
        <Menu.Item style={{float: 'right'}} key="0"></Menu.Item>
        <Menu.Item style={{float: 'right'}} key="1">progrium</Menu.Item>
      </Menu>
    </Row>
  </header>
}

function Inspector(props) {
  return <div>
    <Row style={{padding: "5px"}}> 
      <Col span={2}><Checkbox checked={props.active} /></Col>
      <Col span={22}><Input size="small" value={props.name} /></Col>
    </Row>
    <Collapse bordered={false} defaultActiveKey={['1']}>
      {props.components.map((com) => {
        return <Panel header={com.name} key="1">
          <PropFields fields={com.fields} />
          <Button size="small" style={{marginTop: "10px", width: "100%"}}>
            Test Button
          </Button>
        </Panel>
      })}
    </Collapse>
    <div style={{margin: "20px"}}>
      <Dropdown overlay={menu}>
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
      Message: "Loading...",
      Counter: 0,
      Files: []
    }
  }

  componentDidMount() {
    var api = new qrpc.API();
    const setState = this.setState.bind(this);
    api.handle("state", {
      "serveRPC": async (r, c) => {
        var obj = await c.decode();
        console.log(obj);
        setState(obj);
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

  async onSelect(selectedKeys, info) {
    var data = await window.rpc.call("readFile", selectedKeys[0]);
    this.editor.getModel().setValue(data);
  }


  render() {
    const options = {
      selectOnLineNumbers: true,
      automaticLayout: true
    };
    const fileTree = arrangeIntoTree(this.state.Files);
    const treeNodes = fileTree.map(convertNodeToTree);
    refTree = fileTree.map(convertNodeToSelectTree);
    const exampleComponents = [
      {name: "foo.Component", fields: [
        {type: "string", name: "StringField", value: "Foobar"},
        {type: "number", name: "NumberField", value: 6},
        {type: "boolean", name: "BoolField", value: true},
        {type: "struct", name: "SomeObject", fields: [
          {type: "string", name: "StringField", value: "Foobar"},
        ]},
        {type: "map", name: "MapValue", fields: [
          {type: "string", name: "str2", value: "hello1"},
          {type: "string", name: "str1", value: "hello2"},
        ]},
        {type: "array", name: "NumberList", fields: [
          {type: "number", value: 6},
          {type: "number", value: 6},
        ]},
      ]},
      {name: "twilio.Client", fields: [
        {type: "reference", name: "Ref", value: "/Foobar"},
        {type: "number", name: "NumberField", value: 6},
      ]}
    ]
    const self = this;
    return (
      <div style={{height: "100%"}}>
        <Header />
        <Layout style={{height: "100%"}}>
          <SplitterLayout customClassName="tree" primaryIndex={1} secondaryMinSize={150} secondaryInitialSize={200}>
            <DirectoryTree
              defaultExpandAll
              onSelect={this.onSelect.bind(this)}
            >
              {treeNodes}
            </DirectoryTree>
            
            <SplitterLayout customClassName="tree" primaryIndex={1} secondaryMinSize={250} secondaryInitialSize={250}>
              <Inspector name="TestObject" active={true} components={exampleComponents} />

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
