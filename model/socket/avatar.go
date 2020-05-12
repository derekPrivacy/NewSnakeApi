package model

type Avatar struct {
	ID int `sql:"AUTO_INCREMENT"`

	// for snake example
	PositionX float64
	PositionY float64

	BodyLength int
	Direction  string
}
