package ds

type CartCard struct {
	CartId   uint `gorm:"primaryKey"`
	CardId   uint `gorm:"primaryKey"`
	Cart     Cart `gorm:"foreignKey:CartId"`
	Card     Card `gorm:"foreignKey:CardId"`
	Position uint
	Text     string `gorm:"type:varchar(150)"`
	Result   string `gorm:"type:varchar(400)"`
}
