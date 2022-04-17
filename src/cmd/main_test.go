package main

import (
	"encoding/json"
	"example.com/m/src/model"
	"testing"
)


type maxPathSumTest struct {
	treeJson string
	expected int64
}

const (
	jsonVal1 = `{
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
}`
	jsonVal2 = `{
"tree": {
"nodes": [
{"id": "1", "left": "2", "right": "3", "value": 1},
{"id": "3", "left": null, "right": null, "value": 3},
{"id": "2", "left": null, "right": null, "value": 2}
],
"root": "1"
}
}`
	jsonVal3 = `{
"tree": {
"nodes": [
{"id": "1", "left": "-10", "right": "-5", "value": 1},
{"id": "-5", "left": "-20", "right": "-21", "value": -5},
{"id": "-21", "left": "100-2", "right": "1-3", "value": -21},
{"id": "1-3", "left": null, "right": null, "value": 1},
{"id": "100-2", "left": null, "right": null, "value": 100},
{"id": "-20", "left": "100", "right": "2", "value": -20},
{"id": "2", "left": null, "right": null, "value": 2},
{"id": "100", "left": null, "right": null, "value": 100},
{"id": "-10", "left": "30", "right": "45", "value": -10},
{"id": "45", "left": "3", "right": "-3", "value": 45},
{"id": "-3", "left": null, "right": null, "value": -3},
{"id": "3", "left": null, "right": null, "value": 3},
{"id": "30", "left": "5", "right": "1-2", "value": 30},
{"id": "1-2", "left": null, "right": null, "value": 1},
{"id": "5", "left": null, "right": null, "value": 5}
],
"root": "1"
}
}`
)


var maxPathSumTests = []maxPathSumTest{
	maxPathSumTest{jsonVal1, 18},
	maxPathSumTest{jsonVal2, 6},
	maxPathSumTest{jsonVal3, 154},

}

func TestMaxPathSum(t *testing.T){

	var reqData model.RequestBody

	for _, test := range maxPathSumTests{

		err := json.Unmarshal([]byte(test.treeJson),&reqData)
		if err != nil {
			t.Errorf("body unmarshall error")
		}

		myReqData := make(map[string]*model.ReqMapData, 0)

		for _, row := range reqData.Tree.Nodes {

			myReqData[row.ID] = &model.ReqMapData{
				Value:       row.Value,
				LeftNodeID:  row.Left,
				RightNodeID: row.Right,
			}
		}

		if output := model.MaxPathSum(myReqData,reqData.Tree.Root); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
