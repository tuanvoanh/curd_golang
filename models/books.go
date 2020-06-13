package models

import (
	"API_BASIC_LEARN/configs"
)

func GetAllBook(b *[]Book) (err error) {
	if err = configs.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBook(b *Book) (err error) {
	if err = configs.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Book, id string) (err error) {
	if err := configs.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Book, id string) (err error) {
	configs.DB.Save(b)
	return nil
}

func DeleteBook(b *Book, id string) (err error) {
	configs.DB.Where("id = ?", id).Delete(b)
	return nil
}
