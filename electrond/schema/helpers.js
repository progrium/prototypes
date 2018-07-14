function convertType(node) {
    var exp = ""
    switch (node.type) {
      case "String":
        exp = "string"
        break
      case "Integer":
        exp = "number"
        break
      case "Object":
        if (node.properties === undefined) {
          exp = "any"
        } else {
          exp = objectType(node)
        }
        break
      case "Function":
        exp = fnType(node.parameters)
        break
      case "Number":
        exp = "number"
        break
      case "Boolean":
        exp = "boolean"
        break
      default:
        exp = node.type
    }
    if (node.collection) {
      return exp+"[]"
    } else {
      return exp
    }
  }
  
  function typeValue(node) {
    var exp = ""
    switch (node.type) {
      case "String":
        exp = '""'
        break
      case "Integer":
        exp = '0'
        break
      case "Object":
        exp = 'null'
        break
      case "Function":
        return 'function() {}'
        break
      case "Number":
        exp = '0'
        break
      case "Boolean":
        exp = 'false'
        break
      default:
        if (node.collection) {
          return 'new Array<'+node.type+'>()'
        } else {
          return 'new '+node.type+'()'
        }
    }
    if (node.collection) {
      return "["+exp+"]"
    } else {
      return exp
    }
  }
  
  function fnType(params) {
    if (params !== undefined && params.length > 0) {
      return "("+params.map(function(param) {
        return param.name + ": " + convertType(param)
      }).join(", ")+") => void"
    } else {
      return "() => void"
    }
  }
  
  function buildArgs(method) {
    return (method.parameters||[]).map(function(param) {
      return param.name + ": " + convertType(param)
    }).join(", ")
  }
  
  function objectType(node) {
    return "{"+(node.properties||[]).map(function(prop) {
      return prop.name + ": " + convertType(prop)
    }).join(", ")+"}"
  }
  
  function classProps(node) {
    var props = (node.properties||[]).map(function(prop) {
      return "    "+prop.name+": "+convertType(prop)+";"
    })
    return props.join("\n")
  }
  
  function callbackToReturn(arg) {
    return function(node) {
      var cbParamIdx = -1
      node.parameters.forEach(function(param,idx) {
        if (param.name == arg) {
          cbParamIdx = idx
        }
      })
      if (cbParamIdx > -1) {
        var param = node.parameters[cbParamIdx]
        node.parameters.splice(cbParamIdx,1)
        node.returns = {
          "type": "Object",
          "name": param.name,
          "callbackParam": cbParamIdx,
          "collection": false,
          "properties": param.parameters,
        }
      }
      return node
    }
  }
  
  function paramToStruct(param, struct) {
    return function(node, newNodes) {
      // TODO: make it not hard coded to a single param method
      var newStruct = node.constructorMethod.parameters[0]
      newStruct.type = "Structure"
      newStruct.name = struct
      newNodes.push(newStruct)
      node.constructorMethod.parameters = [{
        name: param,
        type: struct,
        collection: newStruct.collection,
        required: newStruct.required,
      }]
      return node
    }
  }
  
  function classToModule(modName, instance) {
    return function(node, newNodes) {
      node.type = "Structure"
      node.properties = node.instanceProperties||[]
      node.properties.splice(0,0, {
        name: "handle",
        type: "string",
        required: false,
        collection: false
      })
      node.constructorMethod.name = "make"
      node.constructorMethod.returns = {
        "type": node.name,
        "collection": false,
      }
      var methods = node.staticMethods||[]
      methods.push(node.constructorMethod)
      methods.push({
        name: "ref",
        parameters: [
          {
            name: "handle",
            type: "string"
          }
        ],
        returns: {
          type: node.name,
          collection: false
        }
      })
      if (instance === true) {
        node.instanceMethods.forEach(function(method) {
          method.parameters = method.parameters||[]
          method.parameters.splice(0, 0, {
            name: "tray",
            type: "Tray",
            required: true,
            collection: false,
          })
          methods.push(method)
        })
      }
      newNodes.push({
        "name": modName,
        "type": "Module",
        "process": {"main": true},
        "methods": methods
      })
      return node
    }
  }
  
  module.exports = {
    eventTitle: function(name) {
      var words = name.split('-');
  
      for(var i = 0; i < words.length; i++) {
        var word = words[i];
        words[i] = word.charAt(0).toUpperCase() + word.slice(1);
      }
  
      return words.join('');
    },
    fnType: fnType,
    buildArgs: buildArgs,
    classProps: classProps,
    callbackToReturn: callbackToReturn,
    classToModule: classToModule,
    paramToStruct: paramToStruct,
    buildReturn: function(method) {
      if (method.returns !== undefined) {
        return ": "+convertType(method.returns)
      } else {
        return ""
      }
    },
    buildBody: function(method) {
      if (method.returns !== undefined) {
        return "return "+typeValue(method.returns)+";"
      } else {
        return ""
      }
    },
  }
  