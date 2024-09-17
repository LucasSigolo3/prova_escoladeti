package model

type Computador struct {
	Id             int    `gorm:"type:int;primary_key"`
	Name           string `gorm:"type:varchar(255)"`
	Cor            string `gorm:"type:varchar(255)"`
	DataFabricacao int    `gorm:"type:smallint"`
}

func (Computador) TableName() string {
	return "computador"
}
