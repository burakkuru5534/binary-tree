package model

import (
	"example.com/m/src/helper"
	"gopkg.in/guregu/null.v4"
)


type ReqMapData struct {
	Value       int64
	LeftNodeID  null.String
	RightNodeID null.String
}

func  MaxPathSum(rd map[string]*ReqMapData,id string) int64 {

	if rd == nil {
		return 0
	}
	maxValue := rd[id].Value

	// dfs return the max sum path that starts from root
	var dfs func(root *ReqMapData) int64
	dfs = func(root *ReqMapData) int64 {
		if root == nil {
			return 0
		}
		var leftNode *ReqMapData
		var rightNode *ReqMapData


		if root.LeftNodeID.Valid {
			leftNode = rd[root.LeftNodeID.String]
		}
		if root.RightNodeID.Valid {
			rightNode = rd[root.RightNodeID.String]
		}
		left := helper.MaxVal(0, dfs(leftNode))
		right := helper.MaxVal(0, dfs(rightNode))
		sum := root.Value + left + right
		if sum > maxValue {
			maxValue = sum
		}

		return helper.MaxVal(left, right) + root.Value
	}

	dfs(rd[id])
	return maxValue
}
