import _ from 'underscore';

export function pathsToTree(paths, nodePaths) {
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
                    id: nodePaths[currentPath],
                    children: [],
                }
  
                currentLevel.push(newPart);
                currentLevel = newPart.children;
            }
        });
    });
  
    return tree;
  }