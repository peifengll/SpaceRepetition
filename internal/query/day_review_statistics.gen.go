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

	"github.com/peifengll/SpaceRepetition/internal/model"
)

func newDayReviewStatistic(db *gorm.DB, opts ...gen.DOOption) dayReviewStatistic {
	_dayReviewStatistic := dayReviewStatistic{}

	_dayReviewStatistic.dayReviewStatisticDo.UseDB(db, opts...)
	_dayReviewStatistic.dayReviewStatisticDo.UseModel(&model.DayReviewStatistic{})

	tableName := _dayReviewStatistic.dayReviewStatisticDo.TableName()
	_dayReviewStatistic.ALL = field.NewAsterisk(tableName)
	_dayReviewStatistic.ID = field.NewInt64(tableName, "id")
	_dayReviewStatistic.CardNum = field.NewInt64(tableName, "card_num")
	_dayReviewStatistic.EasyNum = field.NewInt64(tableName, "easy_num")
	_dayReviewStatistic.GoodNum = field.NewInt64(tableName, "good_num")
	_dayReviewStatistic.AgainNum = field.NewInt64(tableName, "again_num")
	_dayReviewStatistic.HardNum = field.NewInt64(tableName, "hard_num")
	_dayReviewStatistic.RecordDate = field.NewTime(tableName, "record_date")

	_dayReviewStatistic.fillFieldMap()

	return _dayReviewStatistic
}

// dayReviewStatistic 用户每天的复习情况记录
type dayReviewStatistic struct {
	dayReviewStatisticDo

	ALL        field.Asterisk
	ID         field.Int64
	CardNum    field.Int64 // 当天复习了多少张卡片
	EasyNum    field.Int64 // 复习时选择easy的次数
	GoodNum    field.Int64 // 复习时选择good的次数
	AgainNum   field.Int64 // 复习时选择again的次数
	HardNum    field.Int64 // 复习时选择hard的次数
	RecordDate field.Time  // 是哪一天

	fieldMap map[string]field.Expr
}

func (d dayReviewStatistic) Table(newTableName string) *dayReviewStatistic {
	d.dayReviewStatisticDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dayReviewStatistic) As(alias string) *dayReviewStatistic {
	d.dayReviewStatisticDo.DO = *(d.dayReviewStatisticDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dayReviewStatistic) updateTableName(table string) *dayReviewStatistic {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.CardNum = field.NewInt64(table, "card_num")
	d.EasyNum = field.NewInt64(table, "easy_num")
	d.GoodNum = field.NewInt64(table, "good_num")
	d.AgainNum = field.NewInt64(table, "again_num")
	d.HardNum = field.NewInt64(table, "hard_num")
	d.RecordDate = field.NewTime(table, "record_date")

	d.fillFieldMap()

	return d
}

func (d *dayReviewStatistic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dayReviewStatistic) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 7)
	d.fieldMap["id"] = d.ID
	d.fieldMap["card_num"] = d.CardNum
	d.fieldMap["easy_num"] = d.EasyNum
	d.fieldMap["good_num"] = d.GoodNum
	d.fieldMap["again_num"] = d.AgainNum
	d.fieldMap["hard_num"] = d.HardNum
	d.fieldMap["record_date"] = d.RecordDate
}

func (d dayReviewStatistic) clone(db *gorm.DB) dayReviewStatistic {
	d.dayReviewStatisticDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dayReviewStatistic) replaceDB(db *gorm.DB) dayReviewStatistic {
	d.dayReviewStatisticDo.ReplaceDB(db)
	return d
}

type dayReviewStatisticDo struct{ gen.DO }

type IDayReviewStatisticDo interface {
	gen.SubQuery
	Debug() IDayReviewStatisticDo
	WithContext(ctx context.Context) IDayReviewStatisticDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDayReviewStatisticDo
	WriteDB() IDayReviewStatisticDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDayReviewStatisticDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDayReviewStatisticDo
	Not(conds ...gen.Condition) IDayReviewStatisticDo
	Or(conds ...gen.Condition) IDayReviewStatisticDo
	Select(conds ...field.Expr) IDayReviewStatisticDo
	Where(conds ...gen.Condition) IDayReviewStatisticDo
	Order(conds ...field.Expr) IDayReviewStatisticDo
	Distinct(cols ...field.Expr) IDayReviewStatisticDo
	Omit(cols ...field.Expr) IDayReviewStatisticDo
	Join(table schema.Tabler, on ...field.Expr) IDayReviewStatisticDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDayReviewStatisticDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDayReviewStatisticDo
	Group(cols ...field.Expr) IDayReviewStatisticDo
	Having(conds ...gen.Condition) IDayReviewStatisticDo
	Limit(limit int) IDayReviewStatisticDo
	Offset(offset int) IDayReviewStatisticDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDayReviewStatisticDo
	Unscoped() IDayReviewStatisticDo
	Create(values ...*model.DayReviewStatistic) error
	CreateInBatches(values []*model.DayReviewStatistic, batchSize int) error
	Save(values ...*model.DayReviewStatistic) error
	First() (*model.DayReviewStatistic, error)
	Take() (*model.DayReviewStatistic, error)
	Last() (*model.DayReviewStatistic, error)
	Find() ([]*model.DayReviewStatistic, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DayReviewStatistic, err error)
	FindInBatches(result *[]*model.DayReviewStatistic, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DayReviewStatistic) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDayReviewStatisticDo
	Assign(attrs ...field.AssignExpr) IDayReviewStatisticDo
	Joins(fields ...field.RelationField) IDayReviewStatisticDo
	Preload(fields ...field.RelationField) IDayReviewStatisticDo
	FirstOrInit() (*model.DayReviewStatistic, error)
	FirstOrCreate() (*model.DayReviewStatistic, error)
	FindByPage(offset int, limit int) (result []*model.DayReviewStatistic, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDayReviewStatisticDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dayReviewStatisticDo) Debug() IDayReviewStatisticDo {
	return d.withDO(d.DO.Debug())
}

func (d dayReviewStatisticDo) WithContext(ctx context.Context) IDayReviewStatisticDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dayReviewStatisticDo) ReadDB() IDayReviewStatisticDo {
	return d.Clauses(dbresolver.Read)
}

func (d dayReviewStatisticDo) WriteDB() IDayReviewStatisticDo {
	return d.Clauses(dbresolver.Write)
}

func (d dayReviewStatisticDo) Session(config *gorm.Session) IDayReviewStatisticDo {
	return d.withDO(d.DO.Session(config))
}

func (d dayReviewStatisticDo) Clauses(conds ...clause.Expression) IDayReviewStatisticDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dayReviewStatisticDo) Returning(value interface{}, columns ...string) IDayReviewStatisticDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dayReviewStatisticDo) Not(conds ...gen.Condition) IDayReviewStatisticDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dayReviewStatisticDo) Or(conds ...gen.Condition) IDayReviewStatisticDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dayReviewStatisticDo) Select(conds ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dayReviewStatisticDo) Where(conds ...gen.Condition) IDayReviewStatisticDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dayReviewStatisticDo) Order(conds ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dayReviewStatisticDo) Distinct(cols ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dayReviewStatisticDo) Omit(cols ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dayReviewStatisticDo) Join(table schema.Tabler, on ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dayReviewStatisticDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dayReviewStatisticDo) RightJoin(table schema.Tabler, on ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dayReviewStatisticDo) Group(cols ...field.Expr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dayReviewStatisticDo) Having(conds ...gen.Condition) IDayReviewStatisticDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dayReviewStatisticDo) Limit(limit int) IDayReviewStatisticDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dayReviewStatisticDo) Offset(offset int) IDayReviewStatisticDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dayReviewStatisticDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDayReviewStatisticDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dayReviewStatisticDo) Unscoped() IDayReviewStatisticDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dayReviewStatisticDo) Create(values ...*model.DayReviewStatistic) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dayReviewStatisticDo) CreateInBatches(values []*model.DayReviewStatistic, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dayReviewStatisticDo) Save(values ...*model.DayReviewStatistic) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dayReviewStatisticDo) First() (*model.DayReviewStatistic, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DayReviewStatistic), nil
	}
}

func (d dayReviewStatisticDo) Take() (*model.DayReviewStatistic, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DayReviewStatistic), nil
	}
}

func (d dayReviewStatisticDo) Last() (*model.DayReviewStatistic, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DayReviewStatistic), nil
	}
}

func (d dayReviewStatisticDo) Find() ([]*model.DayReviewStatistic, error) {
	result, err := d.DO.Find()
	return result.([]*model.DayReviewStatistic), err
}

func (d dayReviewStatisticDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DayReviewStatistic, err error) {
	buf := make([]*model.DayReviewStatistic, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dayReviewStatisticDo) FindInBatches(result *[]*model.DayReviewStatistic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dayReviewStatisticDo) Attrs(attrs ...field.AssignExpr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dayReviewStatisticDo) Assign(attrs ...field.AssignExpr) IDayReviewStatisticDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dayReviewStatisticDo) Joins(fields ...field.RelationField) IDayReviewStatisticDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dayReviewStatisticDo) Preload(fields ...field.RelationField) IDayReviewStatisticDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dayReviewStatisticDo) FirstOrInit() (*model.DayReviewStatistic, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DayReviewStatistic), nil
	}
}

func (d dayReviewStatisticDo) FirstOrCreate() (*model.DayReviewStatistic, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DayReviewStatistic), nil
	}
}

func (d dayReviewStatisticDo) FindByPage(offset int, limit int) (result []*model.DayReviewStatistic, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dayReviewStatisticDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dayReviewStatisticDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dayReviewStatisticDo) Delete(models ...*model.DayReviewStatistic) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dayReviewStatisticDo) withDO(do gen.Dao) *dayReviewStatisticDo {
	d.DO = *do.(*gen.DO)
	return d
}
