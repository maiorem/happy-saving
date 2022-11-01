package repository

import (
	"happy/dto"
	"happy/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repository interface {

	// User
	Join(user dto.CreateUserRequest)       // 회원가입
	UpdateUser(user dto.UpdateUserRequest) // 회원정보수정
	FindAllUser() []entity.User            // 회원 리스트
	DeleteUser(user entity.User)           // 회원탈퇴
	FindByEmail(email string) entity.User  // 회원 정보보기
	Login(email string) entity.User        // 로그인

	// Savebox
	Save(savebox dto.CreateBoxRequest)           // 저금통 등록
	UpdateBox(savebox dto.UpdateBoxRequest)      // 저금통 수정
	FindAllBox() []entity.SaveBox                // 저금통 전체 리스트 (어드민)
	DeleteBox(savebox entity.SaveBox)            // 저금통 삭제
	ActivateBox() entity.SaveBox                 // 활성화 상태인 저금통 (회원별로 수정)
	HistoryAllBox() []entity.SaveBox             // 비활성화 상태인 모든 저금통 (회원별로 수정)
	FindHistoryBoxById(id uint64) entity.SaveBox // 비활성화 상태인 저금통 중 하나

	// Diary
	Write(diary dto.CreateDiaryRequest)       // 일기 쓰기
	UpdateDiary(diary dto.UpdateDiaryRequest) // 일기 수정
	FindAllDiary() []entity.Diary             // 일기 전체 리스트 (어드민)
	// 일기 카운트 (저금통 별 조회)
	// 일기 전체 (저금통 별 조회)
	// 일기 하나 열기
	DeleteDiary(diary entity.Diary) // 일기 삭제

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

// ========================== User
// ///// 회원가입
func (db *database) Join(user dto.CreateUserRequest) {
	db.connection.Create(&user)
}

// 이메일 중복 체크 (bool)
func (db *database) EmailDuplicationCheck(email string) bool {
	var user entity.User
	result := true
	resultquery := db.connection.Model(entity.User{Email: email}).First(&user)
	if resultquery.RowsAffected > 0 {
		result = false
	}
	return result
}

// ///// 회원정보 수정 (상세 수정 기능 포함해야함)
func (db *database) UpdateUser(user dto.UpdateUserRequest) {
	db.connection.Save(&user)
}

// ///// (어드민) 회원전체 리스트
func (db *database) FindAllUser() []entity.User {
	var members []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&members)
	return members
}

// ////// 내 정보 보기
func (db *database) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Select("email = ?", email).First(&user)
	return user
}

////////// 로그인

func (db *database) Login(email string) entity.User {
	var user entity.User
	db.connection.Select("email = ?", email).First(&user)
	return user
}

// ///// 회원탈퇴
func (db *database) DeleteUser(user entity.User) {
	db.connection.Save(&user)
}

// ======================= Savebox
// 박스 등록
func (db *database) Save(savebox dto.CreateBoxRequest) {
	db.connection.Create(&savebox)
}

// 박스 수정
func (db *database) UpdateBox(savebox dto.UpdateBoxRequest) {
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
	var viewbox entity.SaveBox
	db.connection.Where("activate = ?", true).First(&viewbox)
	return viewbox
}

// 비활성화 박스 전체 리스트
func (db *database) HistoryAllBox() []entity.SaveBox {
	var boxes []entity.SaveBox
	db.connection.Where("activate = ?", false).Find(&boxes)
	return boxes
}

// activate == false 박스 중 특정 박스 하나 (findbyid)
func (db *database) FindHistoryBoxById(id uint64) entity.SaveBox {
	var historybox entity.SaveBox
	db.connection.Where("activate = ?", false).First(&historybox, id)
	return historybox
}

// 박스 삭제
func (db *database) DeleteBox(savebox entity.SaveBox) {
	db.connection.Save(&savebox)
}

// ================================ Diary
// 일기 쓰기
func (db *database) Write(diary dto.CreateDiaryRequest) {
	db.connection.Create(&diary)
}

// 일기 수정
func (db *database) UpdateDiary(diary dto.UpdateDiaryRequest) {
	db.connection.Save(&diary)
}

// 일기장 보기 (1건)

// 일기 전체 (박스별 리스트로 수정해야 함)
func (db *database) FindAllDiary() []entity.Diary {
	var diaries []entity.Diary
	db.connection.Set("gorm:auto_preload", true).Find(&diaries)
	return diaries
}

// 일기 삭제
func (db *database) DeleteDiary(diary entity.Diary) {
	db.connection.Save(&diary)
}
