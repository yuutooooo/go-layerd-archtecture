package db

import (
	"layerd-archtecture/domain/model"
	"time"

	"gorm.io/gorm"
)

// ユーザーのシードデータ
func SeedUsers(db *gorm.DB) error {
	users := []model.User{
		{Name: "テストユーザー1", Email: "test1@example.com", Password: "password"},
		{Name: "テストユーザー2", Email: "test2@example.com", Password: "password"},
	}
	return db.Create(&users).Error
}

// グループのシードデータ
func SeedGroups(db *gorm.DB) error {
	groups := []model.Group{
		{Name: "個人"},
		{Name: "家族"},
	}
	return db.Create(&groups).Error
}

// カテゴリーのシードデータ
func SeedCategories(db *gorm.DB) error {
	categories := []model.Category{
		{Name: "給与", Type: "income", Color: "#4CAF50"},
		{Name: "食費", Type: "expense", Color: "#F44336"},
		{Name: "交通費", Type: "expense", Color: "#2196F3"},
	}
	return db.Create(&categories).Error
}

// 収支データのシード
func SeedIncomeAndExpenditures(db *gorm.DB) error {
	var user model.User
	if err := db.First(&user).Error; err != nil {
		return err
	}

	var category model.Category
	if err := db.First(&category).Error; err != nil {
		return err
	}

	incomeAndExpenditures := []model.IncomeAndExpenditure{
		{Amount: 300000, Date: time.Now(), CategoryID: category.ID, UserID: user.ID, Description: "1月給与"},
		{Amount: -1500, Date: time.Now(), CategoryID: category.ID, UserID: user.ID, Description: "ランチ"},
	}
	return db.Create(&incomeAndExpenditures).Error
}

// 定期支払いデータのシード
func SeedRecurringPayments(db *gorm.DB) error {
	var user model.User
	if err := db.First(&user).Error; err != nil {
		return err
	}

	var category model.Category
	if err := db.First(&category).Error; err != nil {
		return err
	}

	recurringPayments := []model.RecurringPayment{
		{Amount: -5000, StartDate: time.Now(), Frequency: "monthly", CategoryID: category.ID, UserID: user.ID, Description: "サブスクリプション"},
	}
	return db.Create(&recurringPayments).Error
}

// メモのシードデータ
func SeedNotes(db *gorm.DB) error {
	var incomeAndExpenditure model.IncomeAndExpenditure
	if err := db.First(&incomeAndExpenditure).Error; err != nil {
		return err
	}

	notes := []model.Note{
		{Content: "現金支払い", IncomeAndExpenditureID: incomeAndExpenditure.ID},
	}
	return db.Create(&notes).Error
}

// タグのシードデータ
func SeedTags(db *gorm.DB) error {
	tags := []model.Tag{
		{Name: "固定費"},
		{Name: "変動費"},
		{Name: "臨時収入"},
	}
	return db.Create(&tags).Error
}
