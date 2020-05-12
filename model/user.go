package model

type GoUser struct {
	ID       int    `sql:"AUTO_INCREMENT"`
	UserName string `gorm:"unique;not null"`
	Password string
	//
	Level   int
	Ranking int
}
