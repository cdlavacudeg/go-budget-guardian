// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTotalPerCategoryIncome = "total_per_category_income"

// TotalPerCategoryIncome mapped from table <total_per_category_income>
type TotalPerCategoryIncome struct {
	Year      int32   `gorm:"column:Year" json:"Year"`
	Usuario   int32   `gorm:"column:Usuario;not null" json:"Usuario"`
	IDCat     int32   `gorm:"column:idCat;not null" json:"idCat"`
	Categoria string  `gorm:"column:Categoria;not null" json:"Categoria"`
	Total     float64 `gorm:"column:Total" json:"Total"`
}

// TableName TotalPerCategoryIncome's table name
func (*TotalPerCategoryIncome) TableName() string {
	return TableNameTotalPerCategoryIncome
}