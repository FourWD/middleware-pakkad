package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type User struct {
	ID       string `json:"id" query:"id" db:"id" gorm:"type:varchar(18);primary_key;"`
	Username string

	orm.GormModel

	LoginDate     time.Time `json:"login_date" query:"login_date" db:"login_date" gorm:"default:null;type:datetime;"`
	LogoutDate    time.Time `json:"logout_date" query:"logout_date" db:"logout_date" gorm:"default:null;type:datetime;"`
	CurrentStatus string
}
