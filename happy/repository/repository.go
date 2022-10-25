package repository

import (
	"happy/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repository interface {

	// User
	Join(user entity.User)
	UpdateUser(user entity.User)
	FindAllUser() []entity.User
	DeleteUser(user entity.User)

	// Savebox
	Save(savebox entity.SaveBox)
	UpdateBox(savebox entity.SaveBox)
	FindAllBox() []entity.SaveBox
	DeleteBox(savebox entity.SaveBox)

	// Diary
	Write(diary entity.Diary)
	UpdateDiary(diary entity.Diary)
	FindAllDiary() []entity.Diary
	DeleteDiary(diary entity.Diary)

	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewRepository() Repository {

	db, err := gorm.Open("mysql", "maiorem:123456@tcp(localhost:3306)/happysave?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect databse")
	}
	db.AutoMigrate(&entity.SaveBox{}, &entity.User{}, &entity.Diary{}, &entity.Emoji{})
	return &database{
		connection: db,
	}

}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

// User
func (db *database) Join(user entity.User) {
	db.connection.Create(&user)
}
func (db *database) UpdateUser(user entity.User) {
	db.connection.Save(&user)
}
func (db *database) FindAllUser() []entity.User {
	var members []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&members)
	return members
}

func (db *database) DeleteUser(user entity.User) {
	db.connection.Save(&user)
}

// Savebox
func (db *database) Save(savebox entity.SaveBox) {
	db.connection.Create(&savebox)
}
func (db *database) UpdateBox(savebox entity.SaveBox) {
	db.connection.Save(&savebox)
}

// 박스 전체 (어드민)
func (db *database) FindAllBox() []entity.SaveBox {
	var boxes []entity.SaveBox
	db.connection.Set("gorm:auto_preload", true).Find(&boxes)
	return boxes
}

// activate == true 상태인 박스 하나만 노출
func (db *database) ActivateBox() entity.SaveBox {
	var boxes []entity.SaveBox
	var viewbox entity.SaveBox
	db.connection.Set("gorm:auto_preload", true).Find(&boxes)
	for _, box := range boxes {
		if box.Activate {
			viewbox = box
		}
	}
	return viewbox
}

// activate == false 박스 중 특정 박스 하나 (findbyid)
func (db *database) FindHistoryBoxById(id uint64) entity.SaveBox {
	var historybox entity.SaveBox
	db.connection.First(&historybox, id)
	return historybox
}

func (db *database) DeleteBox(savebox entity.SaveBox) {
	db.connection.Save(&savebox)
}

// Diary
func (db *database) Write(diary entity.Diary) {
	db.connection.Create(&diary)
}
func (db *database) UpdateDiary(diary entity.Diary) {
	db.connection.Save(&diary)
}
func (db *database) FindAllDiary() []entity.Diary {
	var diaries []entity.Diary
	db.connection.Set("gorm:auto_preload", true).Find(&diaries)
	return diaries
}
func (db *database) DeleteDiary(diary entity.Diary) {
	db.connection.Save(&diary)
}
