import React, { Component } from 'react';
import MonacoEditor from 'react-monaco-editor';
import PanelGroup from 'react-panelgroup';
import { Layout, Tree, Menu } from 'antd';
import _ from 'underscore';

import './App.css';
import 'antd/dist/antd.css';

const DirectoryTree = Tree.DirectoryTree;
const { TreeNode } = Tree;
const {
  Header,
} = Layout;

const qmux = window.qmux;
const qrpc = window.qrpc;

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

function convertNodeToTree(node) {
  var isLeaf = node.children.length === 0;
  return <TreeNode title={node.name} key={node.path} isLeaf={isLeaf}>{node.children.map((n) => { return convertNodeToTree(n) })}</TreeNode>
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
    const self = this;
    return (
      <Layout style={{height: "100%"}}>
        <Header>
          <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={['2']}
          style={{ lineHeight: '64px' }}
        >
          <Menu.Item key="1">{this.state.Message}</Menu.Item>
          <Menu.Item key="2">{this.state.Counter}</Menu.Item>
          <Menu.Item key="3">nav 3</Menu.Item>
        </Menu>
        </Header>
        <Layout style={{height: "100%"}}>
          <PanelGroup borderColor="grey" panelWidths={[
              {size: 250, minSize:200, resize: "dynamic"},
              {minSize:100, resize: "stretch"}
            ]}>
              <div style={{height: "100%"}}>
                <DirectoryTree
                  defaultExpandAll
                  onSelect={this.onSelect.bind(this)}
                >
                  {treeNodes}
                </DirectoryTree>
              </div>
              <div style={{"width": "100%"}}>
                <MonacoEditor
                  language="javascript"
                  theme="vs-dark"
                  value=""
                  options={options}
                  onChange={self.onChange}
                  editorDidMount={self.editorDidMount.bind(self)}
                />
              </div>
          </PanelGroup>
        </Layout>
      </Layout>
    );
  }
}

export default App;
