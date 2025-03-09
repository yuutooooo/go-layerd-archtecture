package model

import (
	"time"

	"gorm.io/gorm"
)

// User ユーザーモデル
type User struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         `gorm:"size:255;not null"`
	Email     string         `gorm:"size:255;not null;unique"`
	Password  string         `gorm:"size:255;not null"`
	Groups    []Group        `gorm:"many2many:user_groups;"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Group グループモデル
type Group struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         `gorm:"size:255;not null"`
	Users     []User         `gorm:"many2many:user_groups;"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Category カテゴリーモデル
type Category struct {
	ID          uint           `gorm:"primarykey"`
	Name        string         `gorm:"size:50;not null"`
	Type        string         `gorm:"size:20;not null"` // income or expense
	Description string         `gorm:"size:255"`
	Color       string         `gorm:"size:7"` // #RRGGBB形式
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// IncomeAndExpenditure 収支モデル
type IncomeAndExpenditure struct {
	ID          uint           `gorm:"primarykey"`
	Amount      int            `gorm:"not null"`
	Date        time.Time      `gorm:"not null"`
	CategoryID  uint           `gorm:"not null"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
	UserID      uint           `gorm:"not null"`
	User        User           `gorm:"foreignKey:UserID"`
	Description string         `gorm:"size:255"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// RecurringPayment 定期支払いモデル
type RecurringPayment struct {
	ID          uint      `gorm:"primarykey"`
	Amount      int       `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     *time.Time
	Frequency   string         `gorm:"size:20;not null"` // monthly, yearly等
	CategoryID  uint           `gorm:"not null"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
	UserID      uint           `gorm:"not null"`
	User        User           `gorm:"foreignKey:UserID"`
	Description string         `gorm:"size:255"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Note メモモデル
type Note struct {
	ID                     uint                 `gorm:"primarykey"`
	Content                string               `gorm:"type:text"`
	IncomeAndExpenditureID uint                 `gorm:"not null"`
	IncomeAndExpenditure   IncomeAndExpenditure `gorm:"foreignKey:IncomeAndExpenditureID"`
	CreatedAt              time.Time            `gorm:"not null"`
	UpdatedAt              time.Time            `gorm:"not null"`
	DeletedAt              gorm.DeletedAt       `gorm:"index"`
}

// Tag タグモデル
type Tag struct {
	ID                    uint                   `gorm:"primarykey"`
	Name                  string                 `gorm:"size:50;not null"`
	IncomeAndExpenditures []IncomeAndExpenditure `gorm:"many2many:income_and_expenditure_tags;"`
	CreatedAt             time.Time              `gorm:"not null"`
	UpdatedAt             time.Time              `gorm:"not null"`
	DeletedAt             gorm.DeletedAt         `gorm:"index"`
}
