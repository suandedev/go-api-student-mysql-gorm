package functions

import (
	"go-api-student-mysql-gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go-api-student?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&User{})

	return db
}

func Create(name string, email string, age uint8) *gorm.DB{
	db := connectDb()
	// create
	rs := db.Create(&model.User{Name: name, Email: email, Age: age})

	if rs.Error != nil {
		panic(rs.Error)
	}
	return rs
}

func ReadAll() *gorm.DB{
	db := connectDb()
	var users []model.User
	rs := db.Find(&users)
	return rs
}

func Update(id uint8, name string, email string, age uint8) *gorm.DB{
	db := connectDb()
	rs := db.Model(&model.User{}).Where("id = ?", id).Updates(model.User{Name: name, Email: email, Age: age})
	return rs
}

func Delete(id uint8) *gorm.DB{
	db := connectDb()
	db.Delete(&model.User{}, id)
	return db
}

func ReadById(id uint8) *gorm.DB{
	db := connectDb()
	var user model.User
	rs := db.Find(&user, id)
	return rs
}