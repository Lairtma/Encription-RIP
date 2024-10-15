package ds

type OrderText struct {
	Id       int            `gorm:"primaryKey"`
	OrderID  int            `gorm:"unique"`
	TextID   int            `gorm:"unique"`
	Order    EncOrDecOrder  `gorm:"primaryKey;foreignKey:OrderID"`
	Text     TextToEncOrDec `gorm:"primaryKey;foreignKey:TextID"`
	Position int
	EncType  string
	Res      string
}
