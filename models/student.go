package models

type Student struct {
	Nama   string `gorm:"type:varchar(300)" json:"nama"`
	Nim    int64  `gorm:"type:int(9)" json:"nim"`
	Alamat string `gorm:"type:varchar(300)" json:"alamat"`
}
