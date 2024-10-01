package ds

type TextToEncOrDec struct {
	Id      int `gorm:"primaryKey"`
	Enc     bool
	Text    string
	Img     string
	ByteLen int
	Status  bool
}
