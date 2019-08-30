package main

import (
	"encoding/json"
	"log"
	"strconv"
)

const (
	hour = -11
)

type FlexInt int

func (fi *FlexInt) UnmarshalJson(b []byte) error {
	log.Println(string(b))
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	*fi = FlexInt(i)
	return nil
}

func main() {
	var err error

	s1 := `{"num":123}`
	var MyNum1 struct {
		Num FlexInt `json:"num"`
	}
	err = json.Unmarshal([]byte(s1), &MyNum1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("s1", MyNum1)

	s2 := `{"num":"123"}`
	//var MyNum2 struct {
	//	Num FlexInt `json:"num"`
	//}
	//err = json.Unmarshal([]byte(s2), &MyNum2)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println("s2", MyNum2)

	var fi *FlexInt
	err = fi.UnmarshalJson([]byte(s2))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("fi", fi)

}