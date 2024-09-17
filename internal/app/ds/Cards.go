package ds

type Card struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(40)"`
	First_img   string `gorm:"type:varchar(150)"`
	Second_img  string `gorm:"type:varchar(150)"`
	Encrypting  bool   `gorm:"type:boolean"`
	Description string `gorm:"type:varchar(1600)"`
	Used        bool   `gorm:"type:boolean"`
}
