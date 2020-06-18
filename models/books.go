package models

import (
	"github.com/jinzhu/gorm"
)

func GetAllBook(DB *gorm.DB, b *[]Book) (err error) {
	if err = DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBook(DB *gorm.DB, b *Book) (err error) {
	if err = DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(DB *gorm.DB, b *Book, id string) (err error) {
	if err := DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(DB *gorm.DB, b *Book, id string) (err error) {
	DB.Save(b)
	return nil
}

func DeleteBook(DB *gorm.DB, b *Book, id string) (err error) {
	DB.Where("id = ?", id).Delete(b)
	return nil
}
