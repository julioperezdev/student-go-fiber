package model

type Student struct {
	ID   int    `gorm:"column:id" json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
