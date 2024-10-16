package ds

import "time"

type EncOrDecOrder struct {
	Id          int       `gorm:"primaryKey" json:"id"`
	Status      int       `json:"status"`
	DateCreate  time.Time `json:"date_create"`
	DateUpdate  time.Time `json:"date_update"`
	DateFinish  time.Time `json:"date_finish"`
	CreatorID   *int      `json:"creator_id"`
	ModeratorID *int      `json:"moderator_id"`
	Creator     Users     `gorm:"foreignKey:CreatorID"`
	Moderator   Users     `gorm:"foreignKey:ModeratorID"`
	Priority    int       `json:"priority"`
}
