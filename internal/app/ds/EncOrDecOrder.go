package ds

import "time"

type EncOrDecOrder struct {
	Id          int `gorm:"primaryKey"`
	Status      int
	DateCreate  time.Time
	DateUpdate  time.Time
	DateFinish  time.Time
	CreatorID   int
	ModeratorID int
	Creator     Users `gorm:"foreignKey:CreatorID"`
	Moderator   Users `gorm:"foreignKey:ModeratorID"`
	Priority    int
}
