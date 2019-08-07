import React from 'react';
import { Row, Col, InputNumber, Dropdown, Checkbox, TreeSelect, Icon, Button, Input, Menu, Collapse } from 'antd';

import {pathsToTree} from './util';

const TreeSelectNode = TreeSelect.TreeNode;
const Panel = Collapse.Panel;

const TreeContext = React.createContext(null);

export function Inspector(props) {
    if (props.id === null) {
        return <div></div>
    }

    var onNameChange = (event) => {
      window.rpc.call("updateNode", {"ID": props.id, "Name": event.target.value});
    }
    var onActiveChange = (event) => {
      window.rpc.call("updateNode", {"ID": props.id, "Active": event.target.checked});
    }

    let panelHeader = (name) => {
      var onClick = async ({item, key, keyPath, domEvent}) => {
        domEvent.stopPropagation();
        if (key === "delete") {
          window.rpc.call("removeComponent", {"ID": props.id, "Component": name});    
        }
        if (key === "edit" && name === "Delegate") {
          var res = await window.rpc.call("readDelegate", {ID: props.id});
          props.openDelegate(res.reply);
        }
      }
      const menu = (
        <Menu>
          <Menu.Item onClick={onClick} key="edit">Edit</Menu.Item>
          <Menu.Item onClick={onClick} key="delete">Delete</Menu.Item>
        </Menu>
      )
      return <div><span>{name}</span><Dropdown overlay={menu} trigger={['hover']}><Icon type="setting" style={{float: "right", marginRight: "4px", marginTop: "2px"}} theme="outlined" /></Dropdown></div>;
    }

    let delegatePlaceholder = null;
    if (props.node.components.length === 0 || props.node.components[0].name !== "Delegate") {
      let onClick = () => {
        window.rpc.call("addDelegate", {"ID": props.id});
      };
      delegatePlaceholder = (
        <Panel header="Delegate" key="-1">
          <Button size="small" onClick={onClick} style={{marginTop: "10px", width: "100%"}}>
            Add Delegate
          </Button>
        </Panel>
      );
    }

    return <TreeContext.Provider value={pathsToTree(props.hierarchy, props.nodePaths)}>
      <Row style={{padding: "5px"}}> 
        <Col span={2}><Checkbox checked={props.node.active} onChange={onActiveChange} /></Col>
        <Col span={22}><Input size="small" value={props.node.name} onChange={onNameChange} /></Col>
      </Row>
      <Collapse bordered={false} defaultActiveKey={['1']}>
        {delegatePlaceholder}
        {props.node.components.map((com, idx) => {
          return <Panel header={panelHeader(com.name)} key={idx}>
            <PropFields fields={com.fields} />
            <Buttons buttons={com.buttons} />
          </Panel>
        })}
      </Collapse>
      <div style={{margin: "20px"}}>
        <Dropdown overlay={createComponentMenu(props.id, props.components)}>
          <Button style={{width: "100%"}}>
            Add Component <Icon type="down" />
          </Button>
        </Dropdown>
      </div>
    </TreeContext.Provider>
  }

  class FieldInput extends React.Component {
    static contextType = TreeContext;

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
      var onChange = (event) => {
        //console.log(event.target.value);
        window.rpc.call("setValue", {"Path": this.props.path, "Value": event.target.value});
      }
      if (this.state.expressionMode) {
        var onExprChange = (event) => {
          window.rpc.call("setExpression", {"Path": this.props.path, "Value": event.target.value});
        }
        return <div onDoubleClick={this.toggleExpressionMode}>
          <Input size="small" style={{color:"white", backgroundColor: "#555", fontFamily:"monospace"}} onChange={onExprChange} onDoubleClick={this.toggleExpressionMode} value={this.props.expression} />
        </div>
      }
      let readOnly = (this.props.expression||"").length > 0;
      switch (this.props.type) {
        case "boolean":
          var onBoolChange = (event) => {
            if (readOnly) {
              return;
            }
            window.rpc.call("setValue", {"Path": this.props.path, "Value": event.target.checked});
          }
          return <div onDoubleClick={this.toggleExpressionMode}><Checkbox onChange={onBoolChange} checked={this.props.value} /></div>;
        case "string":
          return <Input readOnly={readOnly} size="small" onChange={onChange} onDoubleClick={this.toggleExpressionMode} value={this.props.value} />;
        case "number":
          var onNumChange = (val) => {
            window.rpc.call("setValue", {"Path": this.props.path, "IntValue": val});
          }
          return <div onDoubleClick={this.toggleExpressionMode}><InputNumber readOnly={readOnly} style={{width: "100%"}} size="small" onChange={onNumChange} value={this.props.value} /></div>;
        default:
          if (this.props.type.startsWith("reference:")) {
            var refType = this.props.type.split(":")[1];
            var onRefChange = (val, label, extra) => {
              window.rpc.call("setValue", {"Path": this.props.path, "RefValue": val+"/"+refType});
            }
            return <div onDoubleClick={this.toggleExpressionMode}><TreeSelect
              size="small"
              showSearch
              onChange={onRefChange}
              style={{ width: "100%"}}
              dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
              placeholder="Please select"
              value={this.props.value}
              allowClear
              treeDefaultExpandAll
            >
              {this.context.map(convertNodeToSelectTree)}
            </TreeSelect></div>;
          } else {
            return "???";
          }
      }
    }
  }

  function convertNodeToSelectTree(node) {
    return <TreeSelectNode title={node.name} value={node.path} key={node.path}>{node.children.map((n) => { return convertNodeToSelectTree(n) })}</TreeSelectNode>
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
              return <PropField key={p.name} {...p} />
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
              return <KeyedField key={field.name} name={field.name}><FieldInput {...field} /></KeyedField>
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
              return <PropField key={p.name} {...p} />
            })}
            {props.children}
          </Panel>
      </Collapse>
    );
  }
  
  function PropFields(props) {
    return (props.fields||[]).map((el, idx) => {
      el.key = idx;
      return <PropField {...el} />
    });
  }
  
  function Buttons(props) {
    var methodClick = (event) => {
      window.rpc.call("callMethod", event.target.value);
    }
    return (props.buttons||[]).map((button, idx) => {
      var onClick = methodClick;
      if (button.onclick !== "") {
        onClick = (event) => {
          eval(button.onclick);
        }
      }
      return <Button size="small" onClick={onClick} key={idx} value={button.path} style={{marginTop: "10px", width: "100%"}}>
            {button.name}
        </Button>
    });
  }


  function createComponentMenu(id, components) {
    const onClick = ({item, key, keyPath}) => {
      window.rpc.call("appendComponent", {"ID": id, "Name": key});
    }
    return <Menu onClick={onClick}>
      {(components||[]).map((name) => {
        return <Menu.Item key={name}>{name}</Menu.Item>
      })}
    </Menu> 
  }