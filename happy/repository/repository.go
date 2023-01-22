package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"happy-save-api/dto"
	"happy-save-api/entity"
	"time"
)

type Repository interface {

	// User
	Join(user dto.CreateUserRequest)       // 회원가입
	UpdateUser(user dto.UpdateUserRequest) // 회원정보수정
	DeleteUser(id uint64)                  // 회원탈퇴
	FindById(id uint64) entity.User        // 회원 정보보기
	Login(email string) entity.User        // 로그인

	FindAllUser() []entity.User // 회원 리스트 (관리자)

	// Savebox
	Save(savebox dto.CreateBoxRequest)            // 저금통 등록
	UpdateBox(savebox dto.UpdateBoxRequest)       // 저금통 수정
	DeleteBox(savebox entity.SaveBox)             // 저금통 삭제
	ActivateBox(userId uint64) entity.SaveBox     // 활성화 상태인 저금통 상세 (회원 ID) => 작성 중인 메인 저금통 (1개) / 작성한 일기 이모지 목록
	HistoryAllBox(userId uint64) []entity.SaveBox // 비활성화 상태인 모든 저금통 목록 (회원 ID)
	FindHistoryBoxById(id uint64) entity.SaveBox  // 비활성화 상태인 저금통 중 하나 상세 (저금통 ID) => 일기 목록

	FindAllBox() []entity.SaveBox // 저금통 전체 리스트 (관리자)

	// Diary
	Write(diary dto.CreateDiaryRequest)       // 일기 쓰기
	UpdateDiary(diary dto.UpdateDiaryRequest) // 일기 수정
	CountDiaryInBox(boxId uint64) int64       // 일기 카운트 (저금통 ID)
	FindDiaryById(id uint64) entity.Diary     // 일기 하나 열기 (다이어리 ID)

	FindAllDiary() []entity.Diary                    // 일기 전체 리스트 (관리자)
	FindAllDiaryByBoxId(boxid uint64) []entity.Diary // 박스에 해당되는 다이어리 전체

	// Emoji
	FindAllEmoji() []entity.Emoji // 이모지 전체
	EmojiOne(id uint64) string

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
		//panic("Failed to close database")
		panic(err.Error())
	}
}

// ========================== User
// ///// 회원가입
func (db *database) Join(user dto.CreateUserRequest) {
	var userentity entity.User
	userentity.Email = user.Email
	userentity.Name = user.Name
	userentity.Password = user.Password

	if db.EmailDuplicationCheck(user.Email) && user.Password == user.ConfirmPassword {
		db.connection.Create(&userentity)
	}
}

// 이메일 중복 체크 (bool)
func (db *database) EmailDuplicationCheck(email string) bool {
	var user entity.User
	result := true
	count := db.connection.Where("email = ?", email).Find(&user).RowsAffected
	if count > 0 {
		result = false
	}
	return result
}

// ///// 회원정보 수정 (상세 수정 기능 포함해야함)
func (db *database) UpdateUser(user dto.UpdateUserRequest) {
	var updateuser entity.User
	// 비밀번호 수정
	db.connection.Where("id = ?", user.ID).Find(&updateuser)

	if user.OldPassowrd == updateuser.Password && user.NewPassword == user.NewConfirmPassword {
		updateuser.Password = user.NewPassword
		db.connection.Save(&updateuser)
		return
	}

	// 이름 수정
	updateuser.Name = user.Name
	db.connection.Save(&updateuser)
}

// ///// (어드민) 회원전체 리스트
func (db *database) FindAllUser() []entity.User {
	var members []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&members)
	return members
}

// ////// 내 정보 보기
func (db *database) FindById(userid uint64) entity.User {
	var user entity.User
	db.connection.Where("id = ?", userid).First(&user)
	return user
}

// //////// 로그인
func (db *database) Login(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).First(&user)
	return user
}

// ///// 회원탈퇴
func (db *database) DeleteUser(id uint64) {
	var user entity.User
	db.connection.Where("id = ?", id).First(&user)
	user.IsActivate = false
	db.connection.Save(&user)
}

// ======================= Savebox
// 박스 등록
func (db *database) Save(savebox dto.CreateBoxRequest) {
	var boxEntity entity.SaveBox
	boxEntity.BoxName = savebox.BoxName
	boxEntity.UserID = savebox.UserID
	boxEntity.OpenBoxDate, _ = time.Parse("2006-01-02", savebox.OpenDate)
	boxEntity.Status = "activate"
	boxEntity.IsOpened = entity.IsOpenedChange(boxEntity.OpenBoxDate)
	db.connection.Create(&boxEntity)
}

// 박스 수정
func (db *database) UpdateBox(savebox dto.UpdateBoxRequest) {
	var boxEntity entity.SaveBox
	boxEntity.ID = savebox.ID
	boxEntity.BoxName = savebox.BoxName
	boxEntity.UserID = savebox.UserID
	boxEntity.OpenBoxDate, _ = time.Parse("2006-01-02", savebox.OpenDate)
	boxEntity.Status = "activate"
	boxEntity.IsOpened = entity.IsOpenedChange(boxEntity.OpenBoxDate)
	db.connection.Save(&boxEntity)
}

// 박스 전체 (어드민)
func (db *database) FindAllBox() []entity.SaveBox {
	var boxes []entity.SaveBox
	db.connection.Set("gorm:auto_preload", true).Find(&boxes)
	return boxes
}

// activate == true 상태인 박스 하나만 노출
func (db *database) ActivateBox(userId uint64) entity.SaveBox {
	var viewbox entity.SaveBox
	db.connection.Where("status = ? & user_id = ?", "activate", userId).First(&viewbox)
	return viewbox
}

// 비활성화 박스 전체 리스트
func (db *database) HistoryAllBox(userId uint64) []entity.SaveBox {
	var boxes []entity.SaveBox
	db.connection.Where("status = ? & user_id = ?", "history", userId).Find(&boxes)
	return boxes
}

// activate == false 박스 중 특정 박스 하나 (findbyid)
func (db *database) FindHistoryBoxById(id uint64) entity.SaveBox {
	var historybox entity.SaveBox
	db.connection.Where("status = ?", "history").First(&historybox, id)
	return historybox
}

// 박스 삭제
func (db *database) DeleteBox(savebox entity.SaveBox) {
	db.connection.Save(&savebox)
}

// ================================ Diary
// 일기 쓰기
func (db *database) Write(diary dto.CreateDiaryRequest) {
	var diaryEntity entity.Diary
	diaryEntity.EmojiID = diary.EmojiID
	diaryEntity.SaveBoxID = diary.BoxID
	diaryEntity.Content = diary.Content
	db.connection.Create(&diaryEntity)
}

// 일기 수정
func (db *database) UpdateDiary(diary dto.UpdateDiaryRequest) {
	var diaryEntity entity.Diary
	diaryEntity.ID = diary.ID
	diaryEntity.EmojiID = diary.EmojiID
	diaryEntity.SaveBoxID = diary.BoxID
	diaryEntity.Content = diary.Content
	db.connection.Save(&diaryEntity)
}

// 일기 전체 (박스별 리스트로 수정해야 함)
func (db *database) FindAllDiary() []entity.Diary {
	var diaries []entity.Diary
	db.connection.Set("gorm:auto_preload", true).Find(&diaries)
	return diaries
}

// 일기장 박스별 리스트
func (db *database) FindAllDiaryByBoxId(boxid uint64) []entity.Diary {
	var diaries []entity.Diary
	db.connection.Set("gorm:auto_preload", true).Where("save_box_id = ?", boxid).Find(&diaries)
	return diaries
}

// 박스 하나에 든 일기 카운트
func (db *database) CountDiaryInBox(boxId uint64) int64 {
	var diaries []entity.Diary
	return db.connection.Set("gorm:auto_preload", true).Where("save_box_id = ?", boxId).Find(&diaries).RowsAffected
}

// 일기장 보기 (1건)
func (db *database) FindDiaryById(id uint64) entity.Diary {
	var diary entity.Diary
	db.connection.Where("id = ?", id).First(&diary)
	return diary
}

// ================================ Emoji
// 이모지 리스트
func (db *database) FindAllEmoji() []entity.Emoji {
	var emoji []entity.Emoji
	db.connection.Set("gorm:auto_preload", true).Find(&emoji)
	return emoji
}

// 이모지 단건
func (db *database) EmojiOne(id uint64) string {
	var emoji entity.Emoji
	var emoticon string
	db.connection.Where("id = ?", id).Find(&emoji)
	emoticon = emoji.Emoticon
	return emoticon
}
