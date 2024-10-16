package ds

type TextToEncOrDec struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Enc     bool   `json:"enc"`
	Text    string `json:"text"`
	Img     string `json:"img"`
	ByteLen int    `json:"byte_len"`
	Status  bool   `json:"status"`
}
