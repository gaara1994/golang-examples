package main

import (
	"fmt"
	"path"
)

func main() {
	//获取文件名
	fmt.Println(path.Base("./log/2023/06/07/aaa.txt")) //aaa.txt

	//获取文件所在的目录名
	fmt.Println(path.Dir("./log/2023/06/07/aaa.txt")) //log/2023/06/07

	//获取文件的后缀名
	fmt.Println(path.Ext("./log/2023/06/07/aaa.txt")) //.txt

	//将目录与文件名分开
	fmt.Println(path.Split("./log/2023/06/07/aaa.txt")) //./log/2023/06/07/ aaa.txt

	//添加路径分隔符“/”
	fmt.Println(path.Join("log", "2023", "06", "aaa.txt")) //log/2023/06/aaa.txt

	//判断是否为绝对路径，其实就是判断是不是以 / 开始，只能是unix类的系统
	fmt.Println(path.IsAbs("./log/2023/06/07/aaa.txt"))                                                 //false
	fmt.Println(path.IsAbs("/log/2023/06/07/aaa.txt"))                                                  //true
	fmt.Println(path.IsAbs("D:\\go\\src\\go_code\\standard_library\\path\\log\\2023\\06\\07\\aaa.txt")) //false

}
