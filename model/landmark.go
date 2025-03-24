package model

type Landmark struct {
	Idx         int     `gorm:"column:idx;type:int(11);primary_key;AUTO_INCREMENT"`
	Name        string  `gorm:"column:name;type:varchar(255);NOT NULL"`
	Country     int     `gorm:"column:country;type:int(11);NOT NULL"`
	Detail      string  `gorm:"column:detail;type:varchar(2000);NOT NULL"`
	Url         string  `gorm:"column:url;type:varchar(1000);NOT NULL"`
	Countrydata Country `gorm:"foreignKey:Country"`
}

func (m *Landmark) TableName() string {
	return "landmark"
}
