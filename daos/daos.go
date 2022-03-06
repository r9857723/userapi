package daos

type User struct {
	Id int64 `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id"`
	Account  string    `gorm:"type:varchar(50) NOT NULL;" json:"account"`
	Password  string    `gorm:"type:varchar(50) NOT NULL;" json:"password"`
}



