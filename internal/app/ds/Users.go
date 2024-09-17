package ds

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Login       string `gorm:"type:varchar(50)"`
	Password    string `gorm:"type:varchar(50)"`
	IsModerator bool   `gorm:"type:boolean"`
}
