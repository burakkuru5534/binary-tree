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
		err := helper.BodyToJsonReq(r, &reqData)
		if err != nil {
			json.NewEncoder(w).Encode(http.StatusBadRequest)
			panic(err)
		}

		myReqData := make(map[string]*model.ReqMapData, 0)

		for _, row := range reqData.Tree.Nodes {

			myReqData[row.ID] = &model.ReqMapData{
				Value:       row.Value,
				LeftNodeID:  row.Left,
				RightNodeID: row.Right,
			}
		}

		result := model.MaxPathSum(myReqData,reqData.Tree.Root)

		json.NewEncoder(w).Encode(result)

	})
}
