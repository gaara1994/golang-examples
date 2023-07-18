package main

import (
	"encoding/json"
	"fmt"
)

//结构体的字段必须是首字母大写，否则无法编码

type User struct {
	Name    string
	Age     string
	Country string
	addr    string
	Sex     string
}

func main() {
	var tom = User{
		Name:    "Tom",
		Age:     "18",
		Country: "美国",
		addr:    "纽约",
		Sex:     "男",
	}

	str, _ := json.Marshal(tom)
	fmt.Println(string(str)) //addr 没有显示

	var tom2 User
	json.Unmarshal(str, &tom2)
	fmt.Println(tom2)
}
