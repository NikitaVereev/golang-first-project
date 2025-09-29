package auth

type MainRegistry struct {
	ID          uint   `gorm:"primaryKey;column:idmainregistry" json:"id"`
	Email       string `gorm:"column:email;size:45;uniqueIndex;not null" json:"email"`
	Phone       string `gorm:"column:phone;size:45;not null" json:"phone"`
	Password    string `gorm:"column:password;size:255;not null" json:"-"`
	Name        string `gorm:"column:name;size:45;not null" json:"name"`
	Position    string `gorm:"column:position;size:45" json:"position"`
	Filials     string `gorm:"column:fillals;size:45" json:"filials"`
	Brand       string `gorm:"column:brand;size:45;not null" json:"brand"`
	MainCabinet string `gorm:"column:maincabinet;size:45;uniqueIndex;not null" json:"main_cabinet"`
}

// MainUsers - для быстрой авторизации
type MainUsers struct {
	ID          uint   `gorm:"primaryKey;column:idmainusers" json:"id"`
	Login       string `gorm:"column:login;size:45;uniqueIndex;not null" json:"login"`
	MainCabinet string `gorm:"column:maincabinet;size:45;not null" json:"main_cabinet"`
}

// AuthUser - пользователь внутри кабинета
type AuthUser struct {
	ID          uint   `gorm:"primaryKey;column:idauth_user" json:"id"`
	Login       string `gorm:"column:login;size:45;not null" json:"login"`
	Email       string `gorm:"column:email;size:45;not null" json:"email"`
	Fio         string `gorm:"column:fio;size:45;not null" json:"fio"`
	Password    string `gorm:"column:password;size:255;not null" json:"-"`
	IDFather    string `gorm:"column:idfather;size:45" json:"id_father"`
	IDGroup     string `gorm:"column:idgroup;size:45" json:"id_group"`
	MainCabinet string `gorm:"column:maincabinet;size:45;not null" json:"main_cabinet"`
	Responsible string `gorm:"column:responsible;size:45" json:"responsible"`
	Role        string `gorm:"column:role;size:45;default:user" json:"role"`
}
