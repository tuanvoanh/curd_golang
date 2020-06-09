package configs

import (
	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB
var err error

// ConnectDatabase ...
func ConnectDatabase(dbType string, dbURI string) (*gorm.DB, error) {
	DB, err = gorm.Open(dbType, dbURI)
	return DB, err
}
