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

func newCategoriesSaving(db *gorm.DB, opts ...gen.DOOption) categoriesSaving {
	_categoriesSaving := categoriesSaving{}

	_categoriesSaving.categoriesSavingDo.UseDB(db, opts...)
	_categoriesSaving.categoriesSavingDo.UseModel(&model.CategoriesSaving{})

	tableName := _categoriesSaving.categoriesSavingDo.TableName()
	_categoriesSaving.ALL = field.NewAsterisk(tableName)
	_categoriesSaving.CategoryID = field.NewInt32(tableName, "category_id")
	_categoriesSaving.UserID = field.NewInt32(tableName, "user_id")
	_categoriesSaving.Name = field.NewString(tableName, "name")
	_categoriesSaving.Description = field.NewString(tableName, "description")

	_categoriesSaving.fillFieldMap()

	return _categoriesSaving
}

type categoriesSaving struct {
	categoriesSavingDo

	ALL         field.Asterisk
	CategoryID  field.Int32
	UserID      field.Int32
	Name        field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (c categoriesSaving) Table(newTableName string) *categoriesSaving {
	c.categoriesSavingDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c categoriesSaving) As(alias string) *categoriesSaving {
	c.categoriesSavingDo.DO = *(c.categoriesSavingDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *categoriesSaving) updateTableName(table string) *categoriesSaving {
	c.ALL = field.NewAsterisk(table)
	c.CategoryID = field.NewInt32(table, "category_id")
	c.UserID = field.NewInt32(table, "user_id")
	c.Name = field.NewString(table, "name")
	c.Description = field.NewString(table, "description")

	c.fillFieldMap()

	return c
}

func (c *categoriesSaving) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *categoriesSaving) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["category_id"] = c.CategoryID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["name"] = c.Name
	c.fieldMap["description"] = c.Description
}

func (c categoriesSaving) clone(db *gorm.DB) categoriesSaving {
	c.categoriesSavingDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c categoriesSaving) replaceDB(db *gorm.DB) categoriesSaving {
	c.categoriesSavingDo.ReplaceDB(db)
	return c
}

type categoriesSavingDo struct{ gen.DO }

func (c categoriesSavingDo) Debug() *categoriesSavingDo {
	return c.withDO(c.DO.Debug())
}

func (c categoriesSavingDo) WithContext(ctx context.Context) *categoriesSavingDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categoriesSavingDo) ReadDB() *categoriesSavingDo {
	return c.Clauses(dbresolver.Read)
}

func (c categoriesSavingDo) WriteDB() *categoriesSavingDo {
	return c.Clauses(dbresolver.Write)
}

func (c categoriesSavingDo) Session(config *gorm.Session) *categoriesSavingDo {
	return c.withDO(c.DO.Session(config))
}

func (c categoriesSavingDo) Clauses(conds ...clause.Expression) *categoriesSavingDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categoriesSavingDo) Returning(value interface{}, columns ...string) *categoriesSavingDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categoriesSavingDo) Not(conds ...gen.Condition) *categoriesSavingDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categoriesSavingDo) Or(conds ...gen.Condition) *categoriesSavingDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categoriesSavingDo) Select(conds ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categoriesSavingDo) Where(conds ...gen.Condition) *categoriesSavingDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categoriesSavingDo) Order(conds ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categoriesSavingDo) Distinct(cols ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categoriesSavingDo) Omit(cols ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categoriesSavingDo) Join(table schema.Tabler, on ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categoriesSavingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categoriesSavingDo) RightJoin(table schema.Tabler, on ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categoriesSavingDo) Group(cols ...field.Expr) *categoriesSavingDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categoriesSavingDo) Having(conds ...gen.Condition) *categoriesSavingDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categoriesSavingDo) Limit(limit int) *categoriesSavingDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categoriesSavingDo) Offset(offset int) *categoriesSavingDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categoriesSavingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *categoriesSavingDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categoriesSavingDo) Unscoped() *categoriesSavingDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categoriesSavingDo) Create(values ...*model.CategoriesSaving) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categoriesSavingDo) CreateInBatches(values []*model.CategoriesSaving, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categoriesSavingDo) Save(values ...*model.CategoriesSaving) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categoriesSavingDo) First() (*model.CategoriesSaving, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesSaving), nil
	}
}

func (c categoriesSavingDo) Take() (*model.CategoriesSaving, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesSaving), nil
	}
}

func (c categoriesSavingDo) Last() (*model.CategoriesSaving, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesSaving), nil
	}
}

func (c categoriesSavingDo) Find() ([]*model.CategoriesSaving, error) {
	result, err := c.DO.Find()
	return result.([]*model.CategoriesSaving), err
}

func (c categoriesSavingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CategoriesSaving, err error) {
	buf := make([]*model.CategoriesSaving, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categoriesSavingDo) FindInBatches(result *[]*model.CategoriesSaving, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categoriesSavingDo) Attrs(attrs ...field.AssignExpr) *categoriesSavingDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categoriesSavingDo) Assign(attrs ...field.AssignExpr) *categoriesSavingDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categoriesSavingDo) Joins(fields ...field.RelationField) *categoriesSavingDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categoriesSavingDo) Preload(fields ...field.RelationField) *categoriesSavingDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categoriesSavingDo) FirstOrInit() (*model.CategoriesSaving, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesSaving), nil
	}
}

func (c categoriesSavingDo) FirstOrCreate() (*model.CategoriesSaving, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoriesSaving), nil
	}
}

func (c categoriesSavingDo) FindByPage(offset int, limit int) (result []*model.CategoriesSaving, count int64, err error) {
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

func (c categoriesSavingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categoriesSavingDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categoriesSavingDo) Delete(models ...*model.CategoriesSaving) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categoriesSavingDo) withDO(do gen.Dao) *categoriesSavingDo {
	c.DO = *do.(*gen.DO)
	return c
}
