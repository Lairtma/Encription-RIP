package ds

type OrderText struct {
	Id       int            `gorm:"primaryKey" json:"id"`
	OrderID  int            `gorm:"unique" json:"order_id"`
	TextID   int            `gorm:"unique" json:"text_id"`
	Order    EncOrDecOrder  `gorm:"primaryKey;foreignKey:OrderID"`
	Text     TextToEncOrDec `gorm:"primaryKey;foreignKey:TextID"`
	Position int            `json:"position"`
	EncType  string         `json:"encType"`
	Res      string         `json:"res"`
}
