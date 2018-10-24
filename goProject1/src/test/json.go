package main

import (
	"encoding/json"
	"fmt"
)


func main() {
	m1 := map[string]interface{}{"name": "John", "age": 10}
	fmt.Println(Map2Json(m1))


	str := `{"IsOk":true,"Subjects":["A","B","C"],"company":"KingSoft"}`
	fmt.Println(Json2Map(str))
}

func Map2Json(m map[string]interface{}) string {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	return string(b)
}

func Json2Map(str string) map[string]interface{} {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(str),&m)
	fmt.Println(m)
	return m
}