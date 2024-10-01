package ds

type Users struct {
	Id          int    `gorm:"primaryKey"`
	Login       string `gorm:"type:varchar(255)"`
	Password    string `gorm:"type:varchar(255)"`
	IsModerator bool   `gorm:"type:boolean"`
	FIO         string `gorm:"type:varchar(255)"`
}
