// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTotalPerCategoryGroupExpense = "total_per_category_group_expense"

// TotalPerCategoryGroupExpense mapped from table <total_per_category_group_expense>
type TotalPerCategoryGroupExpense struct {
	Year      int32   `gorm:"column:Year" json:"Year"`
	Month     int32   `gorm:"column:Month" json:"Month"`
	Usuario   int32   `gorm:"column:Usuario;not null" json:"Usuario"`
	Group_    string  `gorm:"column:Group;not null" json:"Group"`
	IDCat     int32   `gorm:"column:idCat;not null" json:"idCat"`
	Categoria string  `gorm:"column:Categoria;not null" json:"Categoria"`
	Total     float64 `gorm:"column:Total" json:"Total"`
}

// TableName TotalPerCategoryGroupExpense's table name
func (*TotalPerCategoryGroupExpense) TableName() string {
	return TableNameTotalPerCategoryGroupExpense
}