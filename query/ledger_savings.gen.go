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

func newLedgerSaving(db *gorm.DB, opts ...gen.DOOption) ledgerSaving {
	_ledgerSaving := ledgerSaving{}

	_ledgerSaving.ledgerSavingDo.UseDB(db, opts...)
	_ledgerSaving.ledgerSavingDo.UseModel(&model.LedgerSaving{})

	tableName := _ledgerSaving.ledgerSavingDo.TableName()
	_ledgerSaving.ALL = field.NewAsterisk(tableName)
	_ledgerSaving.IDSaving = field.NewInt32(tableName, "id_saving")
	_ledgerSaving.AccountID = field.NewInt32(tableName, "account_id")
	_ledgerSaving.GroupID = field.NewInt32(tableName, "group_id")
	_ledgerSaving.SavingCategoryID = field.NewInt32(tableName, "saving_category_id")
	_ledgerSaving.Description = field.NewString(tableName, "description")
	_ledgerSaving.Amount = field.NewFloat64(tableName, "amount")
	_ledgerSaving.Date = field.NewTime(tableName, "date")

	_ledgerSaving.fillFieldMap()

	return _ledgerSaving
}

type ledgerSaving struct {
	ledgerSavingDo

	ALL              field.Asterisk
	IDSaving         field.Int32
	AccountID        field.Int32
	GroupID          field.Int32
	SavingCategoryID field.Int32
	Description      field.String
	Amount           field.Float64
	Date             field.Time

	fieldMap map[string]field.Expr
}

func (l ledgerSaving) Table(newTableName string) *ledgerSaving {
	l.ledgerSavingDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l ledgerSaving) As(alias string) *ledgerSaving {
	l.ledgerSavingDo.DO = *(l.ledgerSavingDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *ledgerSaving) updateTableName(table string) *ledgerSaving {
	l.ALL = field.NewAsterisk(table)
	l.IDSaving = field.NewInt32(table, "id_saving")
	l.AccountID = field.NewInt32(table, "account_id")
	l.GroupID = field.NewInt32(table, "group_id")
	l.SavingCategoryID = field.NewInt32(table, "saving_category_id")
	l.Description = field.NewString(table, "description")
	l.Amount = field.NewFloat64(table, "amount")
	l.Date = field.NewTime(table, "date")

	l.fillFieldMap()

	return l
}

func (l *ledgerSaving) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *ledgerSaving) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 7)
	l.fieldMap["id_saving"] = l.IDSaving
	l.fieldMap["account_id"] = l.AccountID
	l.fieldMap["group_id"] = l.GroupID
	l.fieldMap["saving_category_id"] = l.SavingCategoryID
	l.fieldMap["description"] = l.Description
	l.fieldMap["amount"] = l.Amount
	l.fieldMap["date"] = l.Date
}

func (l ledgerSaving) clone(db *gorm.DB) ledgerSaving {
	l.ledgerSavingDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l ledgerSaving) replaceDB(db *gorm.DB) ledgerSaving {
	l.ledgerSavingDo.ReplaceDB(db)
	return l
}

type ledgerSavingDo struct{ gen.DO }

func (l ledgerSavingDo) Debug() *ledgerSavingDo {
	return l.withDO(l.DO.Debug())
}

func (l ledgerSavingDo) WithContext(ctx context.Context) *ledgerSavingDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l ledgerSavingDo) ReadDB() *ledgerSavingDo {
	return l.Clauses(dbresolver.Read)
}

func (l ledgerSavingDo) WriteDB() *ledgerSavingDo {
	return l.Clauses(dbresolver.Write)
}

func (l ledgerSavingDo) Session(config *gorm.Session) *ledgerSavingDo {
	return l.withDO(l.DO.Session(config))
}

func (l ledgerSavingDo) Clauses(conds ...clause.Expression) *ledgerSavingDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l ledgerSavingDo) Returning(value interface{}, columns ...string) *ledgerSavingDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l ledgerSavingDo) Not(conds ...gen.Condition) *ledgerSavingDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l ledgerSavingDo) Or(conds ...gen.Condition) *ledgerSavingDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l ledgerSavingDo) Select(conds ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l ledgerSavingDo) Where(conds ...gen.Condition) *ledgerSavingDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l ledgerSavingDo) Order(conds ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l ledgerSavingDo) Distinct(cols ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l ledgerSavingDo) Omit(cols ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l ledgerSavingDo) Join(table schema.Tabler, on ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l ledgerSavingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l ledgerSavingDo) RightJoin(table schema.Tabler, on ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l ledgerSavingDo) Group(cols ...field.Expr) *ledgerSavingDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l ledgerSavingDo) Having(conds ...gen.Condition) *ledgerSavingDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l ledgerSavingDo) Limit(limit int) *ledgerSavingDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l ledgerSavingDo) Offset(offset int) *ledgerSavingDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l ledgerSavingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *ledgerSavingDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l ledgerSavingDo) Unscoped() *ledgerSavingDo {
	return l.withDO(l.DO.Unscoped())
}

func (l ledgerSavingDo) Create(values ...*model.LedgerSaving) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l ledgerSavingDo) CreateInBatches(values []*model.LedgerSaving, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l ledgerSavingDo) Save(values ...*model.LedgerSaving) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l ledgerSavingDo) First() (*model.LedgerSaving, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LedgerSaving), nil
	}
}

func (l ledgerSavingDo) Take() (*model.LedgerSaving, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LedgerSaving), nil
	}
}

func (l ledgerSavingDo) Last() (*model.LedgerSaving, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LedgerSaving), nil
	}
}

func (l ledgerSavingDo) Find() ([]*model.LedgerSaving, error) {
	result, err := l.DO.Find()
	return result.([]*model.LedgerSaving), err
}

func (l ledgerSavingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LedgerSaving, err error) {
	buf := make([]*model.LedgerSaving, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l ledgerSavingDo) FindInBatches(result *[]*model.LedgerSaving, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l ledgerSavingDo) Attrs(attrs ...field.AssignExpr) *ledgerSavingDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l ledgerSavingDo) Assign(attrs ...field.AssignExpr) *ledgerSavingDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l ledgerSavingDo) Joins(fields ...field.RelationField) *ledgerSavingDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l ledgerSavingDo) Preload(fields ...field.RelationField) *ledgerSavingDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l ledgerSavingDo) FirstOrInit() (*model.LedgerSaving, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LedgerSaving), nil
	}
}

func (l ledgerSavingDo) FirstOrCreate() (*model.LedgerSaving, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LedgerSaving), nil
	}
}

func (l ledgerSavingDo) FindByPage(offset int, limit int) (result []*model.LedgerSaving, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l ledgerSavingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l ledgerSavingDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l ledgerSavingDo) Delete(models ...*model.LedgerSaving) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *ledgerSavingDo) withDO(do gen.Dao) *ledgerSavingDo {
	l.DO = *do.(*gen.DO)
	return l
}
