package model

import (
	"gopkg.in/guregu/null.v4"
)

type TreeNode struct {
	ID        string
	Value     int64
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type RequestBody struct {
	ID    string      `json:"id"`
	Value int64       `json:"value"`
	Left  null.String `json:"left"`
	Right null.String `json:"right"`
}

func CreateNode(id string,value int64) *TreeNode {
	return &TreeNode{id,value, nil, nil}
}
