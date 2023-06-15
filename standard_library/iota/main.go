package main

import "fmt"

/*
在C语言中有枚举型
enum DAY

	{
		SATURDAY = 0,
		MONDAY,
		TUESDAY ,
		WEDNESDAY,
		THURSDAY,
		FRIDAY,
		SATURDAY
	};
*/
func main() {
	//golang没有枚举，但是可以通过cost实现同样的功能
	/*
		const (
			SATURDAY	= 0
			MONDAY		= 1
			TUESDAY		= 2
			WEDNESDAY	= 3
			THURSDAY	= 4
			FRIDAY		= 5
			SUNDAY		= 6
		)
	*/

	//iota：默认值为0,每一行自动加 1，不包括空行和注释行
	/*
		const (
			SATURDAY	=	iota	//0
			MONDAY					//1
			TUESDAY					//2
			WEDNESDAY				//3
			THURSDAY				//4
			FRIDAY					//5
			SUNDAY					//6
		)
	*/

	//还可以修改起始索引
	const baseIndex = 400
	const (
		SATURDAY  = baseIndex + iota //400
		MONDAY                       //401
		TUESDAY                      //402
		WEDNESDAY                    //403
		THURSDAY                     //404
		FRIDAY                       //405
		SUNDAY                       //406
	)
	fmt.Println(THURSDAY)

}
