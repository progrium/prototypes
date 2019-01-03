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

function LabeledField(props) {
  return (<Row>
      <Col span={10} style={{fontSize: "smaller"}}>{props.label}</Col>
      <Col span={14}>
        {props.children}
      </Col>
    </Row>);
}

function EmbeddedFields(props) {
  return (
    <Collapse className="embedded" bordered={false} defaultActiveKey="1">
        <Panel header={props.label} key="1">
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

const exampleFields = [
  {type: "string", name: "StringField", value: "Foobar"},
  {type: "number", name: "NumberField", value: 6},
  {type: "boolean", name: "BoolField", value: true},
  {type: "struct", name: "SomeObject", value: [
    {type: "string", name: "StringField", value: "Foobar"},
  ]},
  {type: "array", name: "List", value: [
    {type: "number", name: "NumberField", value: 6},
    {type: "number", name: "NumberField", value: 6},
  ]},
]

function PropField(props) {
  var onChange = () => {

  }
  switch (props.type) {
    case "boolean":
      return <LabeledField key={props.key} label={props.name}><Checkbox checked={props.value} /></LabeledField>
    case "string":
      return <LabeledField key={props.key} label={props.name}><Input size="small" onChange={onChange} value={props.value} /></LabeledField>
    case "number":
      return <LabeledField key={props.key} label={props.name}><InputNumber size="small" value={props.value} /></LabeledField>
    case "struct":
      return (
        <EmbeddedFields key={props.key} label={props.name}>
          {props.children}
        </EmbeddedFields>
      );
    case "pointer":
      return (
        <LabeledField label={props.name} key={props.key} >
          <TreeSelect
            size="small"
            showSearch
            style={{ width: "100%"}}
            dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
            placeholder="Please select"
            value="/server/server.go"
            allowClear
            treeDefaultExpandAll
          >
            {props.children}
          </TreeSelect>
        </LabeledField>
      );
    default:
      return <LabeledField key={props.key} label={props.name}><Input size="small" placeholder="small size" /></LabeledField>              
  }

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
    const treeSelectNodes = fileTree.map(convertNodeToSelectTree);
    const self = this;
    return (
      <div style={{height: "100%"}}>
        <header id="header" className="clearfix">
          <Row>
            <Col span={24}>
           
            <Menu
          
          mode="horizontal"
          defaultSelectedKeys={['1']}
          style={{ lineHeight: '64px' }}
        >
          <Menu.Item style={{float: 'right'}} key="0"></Menu.Item>
          <Menu.Item style={{float: 'right'}} key="1">progrium</Menu.Item>
        </Menu>
            </Col>
          </Row>
        
          
        </header>
        <Layout style={{height: "100%"}}>
          <SplitterLayout customClassName="tree" primaryIndex={1} secondaryMinSize={150} secondaryInitialSize={200}>
              
                <DirectoryTree
                  defaultExpandAll
                  onSelect={this.onSelect.bind(this)}
                >
                  {treeNodes}
                </DirectoryTree>
               

                <SplitterLayout customClassName="tree" primaryIndex={1} secondaryMinSize={250} secondaryInitialSize={250}>
                  <div>
                    <Row style={{padding: "5px"}}>
                      <Col span={2}><Icon type="code" /></Col>
                      <Col span={2}><Checkbox /></Col>
                      <Col span={20}><Input size="small" placeholder="small size" /></Col>
                    </Row>
                  
                  
                  <Collapse bordered={false} defaultActiveKey={['1']}>
                    <Panel header="FooComponent" key="1">
                      <PropField type="boolean" name="BoolField" value={true} />
                      <PropField type="number" name="NumberField" value={5} />
                      <PropField type="string" name="StringField" value="Hello" />
                      <PropField type="struct" name="Somethinggg">
                        <PropField type="string" name="StringField" />
                      </PropField>
                      <PropField type="pointer" name="Foobar">
                        {treeSelectNodes}
                      </PropField>
                      
                    </Panel>
                    <Panel header="Foobar" key="2">
                      <PropFields fields={exampleFields} />
                    </Panel>
                    <Panel header="TwilioProvider" key="3">
                      Three
                    </Panel>
                  </Collapse>
                  <Dropdown overlay={menu}>
                    <Button style={{margin: "20px"}}>
                      Add Component <Icon type="down" />
                    </Button>
                  </Dropdown>

                  </div>

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
