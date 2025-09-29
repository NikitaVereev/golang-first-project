package user

import "gorm.io/datatypes"

type User struct {
	Id          string          `json:"id"`
	Idfather    string          `json:"idfather"`
	Fio         string          `json:"fio"`
	Email       string          `json:"email"`
	Role        *string         `json:"role" example:"admin,administrator,manager"`
	Dtreg       string          `json:"dtreg"`
	IdFilial    string          `json:"idfilial"`
	Phone       string          `json:"phone"`
	Position    string          `json:"position"`
	Offline     string          `json:"offline"`
	Services    bool            `json:"services" validate:"required"`
	IdServices  datatypes.JSON  `json:"idservices" gorm:"type:jsonb"`
	Pushtockens *datatypes.JSON `json:"pushtockens" gorm:"type:jsonb"`
	Avatar      *string         `json:"avatar"`
	IdTgbon     *string         `json:"idtgbot"`
	Onlinerec   *bool           `json:"onlinerec"`
	TimeZone    *int            `json:"timezone"`
}
