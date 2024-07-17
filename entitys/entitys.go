package entitys

import "time"

type TodoLists struct {
	ID   int       `gorm:"column:id;primary_key"`
	Name string    `gorm:"column:name"`
	Done bool      `gorm:"column:done"`
	Time time.Time `gorm:"columT:time"`
}
