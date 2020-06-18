package configs

import (
	"github.com/jinzhu/gorm"
)

// DB ...
// var DB *gorm.DB
// var err error

// ConnectDatabase ...
func ConnectDatabase(dbType string, dbURI string) (*gorm.DB, error) {
	return gorm.Open(dbType, dbURI)
}
