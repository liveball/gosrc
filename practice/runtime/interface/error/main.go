package main

import (
	"encoding/json"
	"fmt"
)

var _ error = (*apiErr)(nil)

func main() {
	fmt.Println(test(-1))
	fmt.Println(test(1))
}

type apiErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *apiErr) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func test(a int) error {
	var err *apiErr

	if a > 0 {
		err = &apiErr{
			Code: a,
			Msg:  "msg",
		}
	}

	return err
}
