// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/cdlavacudeg/go-budget-guardian/model"
)

func newCategoriesExpense(db *gorm.DB, opts ...gen.DOOption) categoriesExpense {
	_categoriesExpense := categoriesExpense{}

	_categoriesExpense.categoriesExpenseDo.UseDB(db, opts...)
	_categoriesExpense.categoriesExpenseDo.UseModel(&model.CategoriesExpense{})

	tableName := _categoriesExpense.categoriesExpenseDo.TableName()
	_categoriesExpense.ALL = field.NewAsterisk(tableName)
	_categoriesExpense.CategoryID = field.NewInt32(tableName, "category_id")
	_categoriesExpense.UserID = field.NewInt32(tableName, "user_id")
	_categoriesExpense.Name = field.NewString(tableName, "name")
	_categoriesExpense.Description = field.NewString(tableName, "description")

	_categoriesExpense.fillFieldMap()

	return _categoriesExpense
}

type categoriesExpense struct {
	categoriesExpenseDo

	ALL         field.Asterisk
	CategoryID  field.Int32
	UserID      field.Int32
	Name        field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (c categoriesExpense) Table(newTableName string) *categoriesExpense {
	c.categoriesExpenseDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c categoriesExpense) As(alias string) *categoriesExpense {
	c.categoriesExpenseDo.DO = *(c.categoriesExpenseDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *categoriesExpense) updateTableName(table string) *categoriesExpense {
	c.ALL = field.NewAsterisk(table)
	c.CategoryID = field.NewInt32(table, "category_id")
	c.UserID = field.NewInt32(table, "user_id")
	c.Name = field.NewString(table, "name")
	c.Description = field.NewString(table, "description")

	c.fillFieldMap()

	return c
}

func (c *categoriesExpense) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *categoriesExpense) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["category_id"] = c.CategoryID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["name"] = c.Name
	c.fieldMap["description"] = c.Description
}

func (c categoriesExpense) clone(db *gorm.DB) categoriesExpense {
	c.categoriesExpenseDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c categoriesExpense) replaceDB(db *gorm.DB) categoriesExpense {
	c.categoriesExpenseDo.ReplaceDB(db)
	return c
}

type categoriesExpenseDo struct{ gen.DO }

func (c categoriesExpenseDo) Debug() *categoriesExpenseDo {
	return c.withDO(c.DO.Debug())
}

func (c categoriesExpenseDo) WithContext(ctx context.Context) *categoriesExpenseDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categoriesExpenseDo) ReadDB() *categoriesExpenseDo {
	return c.Clauses(dbresolver.Read)
}

func (c categoriesExpenseDo) WriteDB() *categoriesExpenseDo {
	return c.Clauses(dbresolver.Write)
}

func (c categoriesExpenseDo) Session(config *gorm.Session) *categoriesExpenseDo {
	return c.withDO(c.DO.Session(config))
}

func (c categoriesExpenseDo) Clauses(conds ...clause.Expression) *categoriesExpenseDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categoriesExpenseDo) Returning(value interface{}, columns ...string) *categoriesExpenseDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categoriesExpenseDo) Not(conds ...gen.Condition) *categoriesExpenseDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categoriesExpenseDo) Or(conds ...gen.Condition) *categoriesExpenseDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categoriesExpenseDo) Select(conds ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categoriesExpenseDo) Where(conds ...gen.Condition) *categoriesExpenseDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categoriesExpenseDo) Order(conds ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categoriesExpenseDo) Distinct(cols ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categoriesExpenseDo) Omit(cols ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categoriesExpenseDo) Join(table schema.Tabler, on ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categoriesExpenseDo) LeftJoin(table schema.Tabler, on ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categoriesExpenseDo) RightJoin(table schema.Tabler, on ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categoriesExpenseDo) Group(cols ...field.Expr) *categoriesExpenseDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categoriesExpenseDo) Having(conds ...gen.Condition) *categoriesExpenseDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categoriesExpenseDo) Limit(limit int) *categoriesExpenseDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categoriesExpenseDo) Offset(offset int) *categoriesExpenseDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categoriesExpenseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *categoriesExpenseDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categoriesExpenseDo) Unscoped() *categoriesExpenseDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categoriesExpenseDo) Create(values ...*model.CategoriesExpense) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categoriesExpenseDo) CreateInBatches(values []*model.CategoriesExpense, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categoriesExpenseDo) Save(values ...*model.CategoriesExpense) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categoriesExpenseDo) First() (*model.CategoriesExpense, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesExpense), nil
	}
}

func (c categoriesExpenseDo) Take() (*model.CategoriesExpense, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesExpense), nil
	}
}

func (c categoriesExpenseDo) Last() (*model.CategoriesExpense, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesExpense), nil
	}
}

func (c categoriesExpenseDo) Find() ([]*model.CategoriesExpense, error) {
	result, err := c.DO.Find()
	return result.([]*model.CategoriesExpense), err
}

func (c categoriesExpenseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CategoriesExpense, err error) {
	buf := make([]*model.CategoriesExpense, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categoriesExpenseDo) FindInBatches(result *[]*model.CategoriesExpense, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categoriesExpenseDo) Attrs(attrs ...field.AssignExpr) *categoriesExpenseDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categoriesExpenseDo) Assign(attrs ...field.AssignExpr) *categoriesExpenseDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categoriesExpenseDo) Joins(fields ...field.RelationField) *categoriesExpenseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categoriesExpenseDo) Preload(fields ...field.RelationField) *categoriesExpenseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categoriesExpenseDo) FirstOrInit() (*model.CategoriesExpense, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesExpense), nil
	}
}

func (c categoriesExpenseDo) FirstOrCreate() (*model.CategoriesExpense, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesExpense), nil
	}
}

func (c categoriesExpenseDo) FindByPage(offset int, limit int) (result []*model.CategoriesExpense, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c categoriesExpenseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categoriesExpenseDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categoriesExpenseDo) Delete(models ...*model.CategoriesExpense) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categoriesExpenseDo) withDO(do gen.Dao) *categoriesExpenseDo {
	c.DO = *do.(*gen.DO)
	return c
}
