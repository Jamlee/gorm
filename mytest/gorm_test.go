package mytest

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Student struct {
	ID   int
	Name string
}

func TestOpen(t *testing.T) {
	dsn := "root:my-secret-pw@tcp(9.135.228.214:3306)/hellodb?charset=utf8mb4&parseTime=True" // invalid loc
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("should returns error but got nil")
	}
	stu := &Student{}
	result := mysqlDb.Table("student").Where("id", 1).First(stu)
	if result.Error != nil {
		t.Error(err)
	} else {
		t.Log(stu)
	}
	t.Log(mysqlDb)
}
