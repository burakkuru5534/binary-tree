package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BodyToJsonReq(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return errors.New(fmt.Sprintf("Body unmarshall error %s", string(body)))
	}

	defer r.Body.Close()

	return nil
}

func MaxVal(a,b int64) int64 {
	if a > b { return a }
	return b
}

