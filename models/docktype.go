package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DockTypeModel struct{
	FacilityId int `gorm:"primary_key";"AUTO_INCREMENT"`
	Name string `gorm:"size:255"`
}

