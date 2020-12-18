package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

var mydata = `{
"course":100
}`

func Test_parseJson(t *testing.T) {

	var res Result2

	if err := json.Unmarshal([]byte(mydata), &res); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(res.Course)
}

type Result2 struct {
	Course int64 `json:"course"`
}
