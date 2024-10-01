package ds

type OrderText struct {
	OrderID  int            `gorm:"primaryKey"`
	TextID   int            `gorm:"primaryKey"`
	Order    EncOrDecOrder  `gorm:"primaryKey;foreignKey:OrderID"`
	Text     TextToEncOrDec `gorm:"primaryKey;foreignKey:TextID"`
	Position int
}
