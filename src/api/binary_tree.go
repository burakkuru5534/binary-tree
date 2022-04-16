package api

import (
	"encoding/json"
	"example.com/m/src/helper"
	"example.com/m/src/model"
	"net/http"
)

func BinaryTreeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var reqData model.RequestBody
	//	var bTree *model.TreeNode
		err := helper.BodyToJsonReq(r,&reqData)
		if err != nil {
			panic(err)
		}

		var tree = &model.BinarySearchTree{}

		tree.InitTree(reqData)

		tree.MaxPathSum()

		json.NewEncoder(w).Encode(tree.MaxPathSum())

	})
}