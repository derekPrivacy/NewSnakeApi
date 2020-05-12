package model

type Match struct {
	ID             int `sql:"AUTO_INCREMENT"`
	FirstPlayerID  string
	SecondPlayerID string
	Result         string
}
