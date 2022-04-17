# binary-tree

## Introduction

In this project, we write a web server that accepts a Binary Tree and returns its max path sum.
A path is a collection of connected nodes in a tree, where no node is connected to more than
two other nodes; a path sum is the sum of the values of the nodes in a particular path.
Each BinaryTree node has an integer value, a left child node, and a right child node. Children
nodes can either be BinaryTree nodes themselves on nil.

### Languages and frameworks

Technologies used in this project:

Golang

Gorilla Mux

Cors

net/http

log

encoding json



Test Environments:

postman
golang testing

## Problem solution

We take the json from the request's body and we unmarshall it to our struct.

Sample Input:
```json
{
"tree": {
"nodes": [
{"id": "1", "left": "2", "right": "3", "value": 1},
{"id": "3", "left": "6", "right": "7", "value": 3},
{"id": "7", "left": null, "right": null, "value": 7},
{"id": "6", "left": null, "right": null, "value": 6},
{"id": "2", "left": "4", "right": "5", "value": 2},
{"id": "5", "left": null, "right": null, "value": 5},
{"id": "4", "left": null, "right": null, "value": 4}
],
"root": "1"
}
}
```

After that we create a map string from our struct to generate our tree and calculate max path sum.

Then, we used a function which is calculate the max path sum of the tree.

### Usage Example

Method: POST

 http://localhost:8081/binary/tree
 
 request Body Example:
 ```json
{
"tree": {
"nodes": [
{"id": "1", "left": "2", "right": "3", "value": 1},
{"id": "3", "left": "6", "right": "7", "value": 3},
{"id": "7", "left": null, "right": null, "value": 7},
{"id": "6", "left": null, "right": null, "value": 6},
{"id": "2", "left": "4", "right": "5", "value": 2},
{"id": "5", "left": null, "right": null, "value": 5},
{"id": "4", "left": null, "right": null, "value": 4}
],
"root": "1"
}
}
 ```
 response example:
 
 18
 


