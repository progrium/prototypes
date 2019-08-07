import React from 'react';
import { Row, Button, Col, Dropdown, Tree, Menu } from 'antd';
import {pathsToTree} from './util';

const DirectoryTree = Tree.DirectoryTree;
const { TreeNode } = Tree;

export function Hierarchy(props) {
    return (
        <div>
            <Row style={{padding: "5px"}}> 
                <Col span={6}><Button size="small" onClick={props.onAddNode}>Add</Button></Col>
                <Col span={6}><Button size="small" onClick={props.onReload}>Reload</Button></Col>
                
            </Row>
            <DirectoryTree 
                defaultExpandAll
                onSelect={props.onSelect}>
            
                {pathsToTree(props.hierarchy, props.nodePaths).map(convertNodeToTree)}
            </DirectoryTree>
        </div>
    )
}

class HierarchyNode extends TreeNode {
    render() {
      var onClick = ({item, key, keyPath}) => {
        //console.log(key, this.props.eventKey);
        if (key === "delete") {
          window.rpc.call("deleteNode", this.props.eventKey);
        }
        if (key === "add") {
          let name = prompt("New Node");
          window.rpc.call("appendNode", {"ID": this.props.eventKey, "Name": name});
        }
      }
      const contextMenu = (
        <Menu onClick={onClick}>
          <Menu.Item key="add">Add Node</Menu.Item>
          <Menu.Item key="delete">Delete</Menu.Item>
        </Menu>
      )
      return <Dropdown overlay={contextMenu} trigger={['contextMenu']}>{super.render()}</Dropdown>;
    }
  }
  
  function convertNodeToTree(node) {
    var isLeaf = node.children.length === 0;
    return <HierarchyNode 
        title={node.name} 
        key={node.id} 
        isLeaf={isLeaf}>
        
        {node.children.map((n) => { return convertNodeToTree(n) })}
    </HierarchyNode>
  }

  
