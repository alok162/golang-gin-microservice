package users

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Num  int `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	Name string
	// Age   sql.NullInt64
	Email string `gorm:"type:varchar(100);unique_index"`
	// Address      string  `gorm:"index:addr"`      // create index with name `addr` for address
	// IgnoreMe     int     `gorm:"-"`               // ignore this field
}
