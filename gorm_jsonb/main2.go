//package main
//
//import (
//	"fmt"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//	"log"
//)
//
//var DB *gorm.DB
//
//type Student struct {
//	gorm.Model
//	Name   string            `gorm:"not null"`
//	Member map[string]string `gorm:"type:jsonb"`
//}
//
//type Record struct {
//	gorm.Model
//	Name   string
//	Member string
//}
//
//func (Student) TableName() string {
//	return "students"
//}
//
//func InitDB() {
//	dsn := "host=localhost user=postgres password=MyNewPass4! dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = db.AutoMigrate(&Student{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	DB = db
//}
//func main() {
//	//初始化数据库
//	InitDB()
//
//	//插入一条数据
//	//student := Student{Name: "张小二", Member: map[string]string{"father": "张大二", "grandfather": "张二"}}
//	//DB.Create(&student)
//
//	// 查找 member 中包含键 "grandfather" 的所有记录
//	//var students []Record
//	//err := DB.Table("students").Where("member ? 'grandfather'").Find(&students).Error
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//// 现在 students 变量包含了所有 member 中包含 "grandfather" 键的学生记录
//	//for _, student := range students {
//	//	log.Printf("找到学生: ID=%d, Name=%s, Member=%v", student.ID, student.Name, student.Member)
//	//}
//
//	// 查找 member 中包含键 "grandfather" 的所有记录
//	//var students2 []Record
//	//err := DB.Table("students").Where("member ? 'mother'").Find(&students2).Error
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//// 现在 students 变量包含了所有 member 中包含 "grandfather" 键的学生记录
//	//for _, student := range students2 {
//	//	log.Printf("找到学生: ID=%d, Name=%s, Member=%v", student.ID, student.Name, student.Member)
//	//}
//
//	var students3 []Record
//	err := DB.Table("students").Find(&students3).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, student := range students3 {
//		if _, ok := student.Member["mother"]; ok {
//			log.Printf("找到学生: ID=%d, Name=%s, Member=%v", student.ID, student.Name, student.Member)
//		}
//	}
//}
