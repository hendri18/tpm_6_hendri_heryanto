package models

type Product struct {
	ID    uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"column:name"`
	Price uint64 `json:"price" gorm:"column:price"`
}
