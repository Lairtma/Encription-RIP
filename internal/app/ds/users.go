package ds

type Users struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Login       string `gorm:"type:varchar(255)" json:"login"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	IsModerator bool   `gorm:"type:boolean" json:"is_moderator"`
	FIO         string `gorm:"type:varchar(255)" json:"fio"`
}
