package user

import (
	"bubble/app/models/profile"
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	ID int `gorm:"primaryKey;autoIncrement;" xorm:"'id' int(11) pk autoincr notnull" json:"id"`

	*ThirdParty `gorm:"embedded"` // 第三方信息
	*Info       `gorm:"embedded"` // 	基本信息

	CreatedAt time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;" bson:"deleted_at"`

	// 关联
	Profile *profile.Profile // has one
}

func New(attr ...Attr) *User {
	user := &User{ThirdParty: NewThirdParty(), Info: NewInfo()}
	Attrs(attr).Apply(user)
	return user
}
