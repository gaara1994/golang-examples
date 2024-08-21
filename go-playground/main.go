package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// 定义结构体User，每个字段都有一个validate标签，这些标签定义了字段应该遵循的验证规则
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	Gender         string     `validate:"oneof=male female prefer_not_to"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

// 用户地址信息
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func main() {
	//一：对结构体验证
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Gender:         "male",
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	//创建了一个新的验证器实例，并启用 WithRequiredStructEnabled 选项，这使得结构体本身也可以被验证是否为空。
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		//如果验证过程中遇到一个无法验证的值，如一个 nil 接口，验证器将返回 InvalidValidationError。
		//排除这种情况
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			fmt.Println(err)
			return
		}

		//当 validator 验证失败时，它会返回一个 validator.ValidationErrors 类型的错误对象。
		//这段代码遍历所有的验证错误，并打印出每个错误的相关信息。
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

		}

	}

	//二：对单独的变量验证
	// myEmail := "bill@gmail.com"
	myEmail := "bill&gmail.com"
	errs := validate.Var(myEmail, "required,email")
	if errs != nil {
		fmt.Println(errs)
		return
	}
}
