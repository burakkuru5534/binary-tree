package model

import "gopkg.in/guregu/null.v4"

type Node struct {
	ID    string      `json:"id"`
	Value int64       `json:"value"`
	Left  null.String `json:"left"`
	Right null.String `json:"right"`
}

type Tree struct {
	Nodes []Node `json:"nodes"`
	Root    string `json:"root"`

}

type RequestBody struct {
	Tree Tree `json:"tree"`
}
