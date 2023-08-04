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

func newAccount(db *gorm.DB, opts ...gen.DOOption) account {
	_account := account{}

	_account.accountDo.UseDB(db, opts...)
	_account.accountDo.UseModel(&model.Account{})

	tableName := _account.accountDo.TableName()
	_account.ALL = field.NewAsterisk(tableName)
	_account.AccountID = field.NewInt32(tableName, "account_id")
	_account.UserID = field.NewInt32(tableName, "user_id")
	_account.Name = field.NewString(tableName, "name")
	_account.Description = field.NewString(tableName, "description")

	_account.fillFieldMap()

	return _account
}

type account struct {
	accountDo

	ALL         field.Asterisk
	AccountID   field.Int32
	UserID      field.Int32
	Name        field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (a account) Table(newTableName string) *account {
	a.accountDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a account) As(alias string) *account {
	a.accountDo.DO = *(a.accountDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *account) updateTableName(table string) *account {
	a.ALL = field.NewAsterisk(table)
	a.AccountID = field.NewInt32(table, "account_id")
	a.UserID = field.NewInt32(table, "user_id")
	a.Name = field.NewString(table, "name")
	a.Description = field.NewString(table, "description")

	a.fillFieldMap()

	return a
}

func (a *account) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *account) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["account_id"] = a.AccountID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["name"] = a.Name
	a.fieldMap["description"] = a.Description
}

func (a account) clone(db *gorm.DB) account {
	a.accountDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a account) replaceDB(db *gorm.DB) account {
	a.accountDo.ReplaceDB(db)
	return a
}

type accountDo struct{ gen.DO }

func (a accountDo) Debug() *accountDo {
	return a.withDO(a.DO.Debug())
}

func (a accountDo) WithContext(ctx context.Context) *accountDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountDo) ReadDB() *accountDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountDo) WriteDB() *accountDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountDo) Session(config *gorm.Session) *accountDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountDo) Clauses(conds ...clause.Expression) *accountDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountDo) Returning(value interface{}, columns ...string) *accountDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountDo) Not(conds ...gen.Condition) *accountDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountDo) Or(conds ...gen.Condition) *accountDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountDo) Select(conds ...field.Expr) *accountDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountDo) Where(conds ...gen.Condition) *accountDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountDo) Order(conds ...field.Expr) *accountDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountDo) Distinct(cols ...field.Expr) *accountDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountDo) Omit(cols ...field.Expr) *accountDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountDo) Join(table schema.Tabler, on ...field.Expr) *accountDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *accountDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountDo) RightJoin(table schema.Tabler, on ...field.Expr) *accountDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountDo) Group(cols ...field.Expr) *accountDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountDo) Having(conds ...gen.Condition) *accountDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountDo) Limit(limit int) *accountDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountDo) Offset(offset int) *accountDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *accountDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountDo) Unscoped() *accountDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountDo) Create(values ...*model.Account) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountDo) CreateInBatches(values []*model.Account, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountDo) Save(values ...*model.Account) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountDo) First() (*model.Account, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) Take() (*model.Account, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) Last() (*model.Account, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) Find() ([]*model.Account, error) {
	result, err := a.DO.Find()
	return result.([]*model.Account), err
}

func (a accountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Account, err error) {
	buf := make([]*model.Account, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountDo) FindInBatches(result *[]*model.Account, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountDo) Attrs(attrs ...field.AssignExpr) *accountDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountDo) Assign(attrs ...field.AssignExpr) *accountDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountDo) Joins(fields ...field.RelationField) *accountDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountDo) Preload(fields ...field.RelationField) *accountDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountDo) FirstOrInit() (*model.Account, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) FirstOrCreate() (*model.Account, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) FindByPage(offset int, limit int) (result []*model.Account, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a accountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountDo) Delete(models ...*model.Account) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountDo) withDO(do gen.Dao) *accountDo {
	a.DO = *do.(*gen.DO)
	return a
}