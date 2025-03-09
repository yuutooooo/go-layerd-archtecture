package main

import (
	"layerd-archtecture/domain/model"
	"layerd-archtecture/infrastructure/db"
	"log"
)

// マイグレーション時にはmain関数のコメントアウトを外しておく
// func main() {
	func migration() {
	d, err := db.InitDB()
	if err != nil {
		log.Fatalf("DB接続に失敗しました: %v", err)
	}
	log.Println("DB接続に成功しました")

	err = d.AutoMigrate(&model.User{}, &model.Group{}, &model.Category{}, &model.IncomeAndExpenditure{}, &model.RecurringPayment{}, &model.Note{}, &model.Tag{})
	if err != nil {
		log.Fatalf("マイグレーションに失敗しました: %v", err)
	}
	log.Println("マイグレーションが完了しました")

	log.Println("シードデータを挿入します")
	// 1. 基本データの挿入
	db.SeedUsers(d)      // 最初にユーザー
	db.SeedGroups(d)     // グループ
	db.SeedCategories(d) // カテゴリー

	// 2. 関連データの挿入
	db.SeedIncomeAndExpenditures(d) // 収支データ
	db.SeedRecurringPayments(d)     // 定期支払い
	db.SeedNotes(d)                 // メモ
	db.SeedTags(d)                  // タグ

	log.Println("シードデータの挿入が完了しました")
}
