package ds

import (
	"time"
)

type Cart struct {
	Id          uint `gorm:"primaryKey"`
	Status      uint
	CreateDate  time.Time
	SendDate    time.Time
	EndDate     time.Time
	CreatorId   uint
	ModeratorId uint
	Creator     User `gorm:"foreignKey:CreatorId"`
	Moderator   User `gorm:"foreignKey:ModeratorId"`
}
