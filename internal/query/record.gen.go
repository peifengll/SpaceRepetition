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

func newRecord(db *gorm.DB, opts ...gen.DOOption) record {
	_record := record{}

	_record.recordDo.UseDB(db, opts...)
	_record.recordDo.UseModel(&model.Record{})

	tableName := _record.recordDo.TableName()
	_record.ALL = field.NewAsterisk(tableName)
	_record.ID = field.NewInt64(tableName, "id")
	_record.CreatedAt = field.NewTime(tableName, "created_at")
	_record.UpdatedAt = field.NewTime(tableName, "updated_at")
	_record.DeletedAt = field.NewField(tableName, "deleted_at")
	_record.KnowledgeID = field.NewInt64(tableName, "knowledge_id")
	_record.Due = field.NewTime(tableName, "Due")
	_record.Stability = field.NewFloat64(tableName, "Stability")
	_record.Difficulty = field.NewFloat64(tableName, "Difficulty")
	_record.ElapsedDays = field.NewInt64(tableName, "ElapsedDays")
	_record.ScheduledDays = field.NewInt64(tableName, "ScheduledDays")
	_record.Reps = field.NewInt64(tableName, "Reps")
	_record.Lapses = field.NewInt64(tableName, "Lapses")
	_record.State = field.NewInt64(tableName, "State")
	_record.On = field.NewInt64(tableName, "On")
	_record.Lastop = field.NewInt64(tableName, "lastop")
	_record.LastReview = field.NewTime(tableName, "LastReview")
	_record.UserID = field.NewString(tableName, "user_id")
	_record.Rate = field.NewInt64(tableName, "rate")

	_record.fillFieldMap()

	return _record
}

type record struct {
	recordDo

	ALL           field.Asterisk
	ID            field.Int64
	CreatedAt     field.Time // 创建时间
	UpdatedAt     field.Time // 更新时间
	DeletedAt     field.Field
	KnowledgeID   field.Int64   // 这张卡片跟哪一个知识点有关
	Due           field.Time    // 到期时间，也就是该复习的日子
	Stability     field.Float64 // 稳定性
	Difficulty    field.Float64 // 难度
	ElapsedDays   field.Int64   // 已过天数
	ScheduledDays field.Int64   // 几乎天数
	Reps          field.Int64   // 重复次数
	Lapses        field.Int64   // 失误次数
	State         field.Int64
	On            field.Int64  // 是否被使用
	Lastop        field.Int64  // 上一次的选择
	LastReview    field.Time   // 最后复习时间
	UserID        field.String // 属于哪一个用户
	Rate          field.Int64  // 复习的时候进行的评分

	fieldMap map[string]field.Expr
}

func (r record) Table(newTableName string) *record {
	r.recordDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r record) As(alias string) *record {
	r.recordDo.DO = *(r.recordDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *record) updateTableName(table string) *record {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.KnowledgeID = field.NewInt64(table, "knowledge_id")
	r.Due = field.NewTime(table, "Due")
	r.Stability = field.NewFloat64(table, "Stability")
	r.Difficulty = field.NewFloat64(table, "Difficulty")
	r.ElapsedDays = field.NewInt64(table, "ElapsedDays")
	r.ScheduledDays = field.NewInt64(table, "ScheduledDays")
	r.Reps = field.NewInt64(table, "Reps")
	r.Lapses = field.NewInt64(table, "Lapses")
	r.State = field.NewInt64(table, "State")
	r.On = field.NewInt64(table, "On")
	r.Lastop = field.NewInt64(table, "lastop")
	r.LastReview = field.NewTime(table, "LastReview")
	r.UserID = field.NewString(table, "user_id")
	r.Rate = field.NewInt64(table, "rate")

	r.fillFieldMap()

	return r
}

func (r *record) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *record) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 18)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["knowledge_id"] = r.KnowledgeID
	r.fieldMap["Due"] = r.Due
	r.fieldMap["Stability"] = r.Stability
	r.fieldMap["Difficulty"] = r.Difficulty
	r.fieldMap["ElapsedDays"] = r.ElapsedDays
	r.fieldMap["ScheduledDays"] = r.ScheduledDays
	r.fieldMap["Reps"] = r.Reps
	r.fieldMap["Lapses"] = r.Lapses
	r.fieldMap["State"] = r.State
	r.fieldMap["On"] = r.On
	r.fieldMap["lastop"] = r.Lastop
	r.fieldMap["LastReview"] = r.LastReview
	r.fieldMap["user_id"] = r.UserID
	r.fieldMap["rate"] = r.Rate
}

func (r record) clone(db *gorm.DB) record {
	r.recordDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r record) replaceDB(db *gorm.DB) record {
	r.recordDo.ReplaceDB(db)
	return r
}

type recordDo struct{ gen.DO }

type IRecordDo interface {
	gen.SubQuery
	Debug() IRecordDo
	WithContext(ctx context.Context) IRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRecordDo
	WriteDB() IRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRecordDo
	Not(conds ...gen.Condition) IRecordDo
	Or(conds ...gen.Condition) IRecordDo
	Select(conds ...field.Expr) IRecordDo
	Where(conds ...gen.Condition) IRecordDo
	Order(conds ...field.Expr) IRecordDo
	Distinct(cols ...field.Expr) IRecordDo
	Omit(cols ...field.Expr) IRecordDo
	Join(table schema.Tabler, on ...field.Expr) IRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRecordDo
	Group(cols ...field.Expr) IRecordDo
	Having(conds ...gen.Condition) IRecordDo
	Limit(limit int) IRecordDo
	Offset(offset int) IRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRecordDo
	Unscoped() IRecordDo
	Create(values ...*model.Record) error
	CreateInBatches(values []*model.Record, batchSize int) error
	Save(values ...*model.Record) error
	First() (*model.Record, error)
	Take() (*model.Record, error)
	Last() (*model.Record, error)
	Find() ([]*model.Record, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Record, err error)
	FindInBatches(result *[]*model.Record, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Record) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRecordDo
	Assign(attrs ...field.AssignExpr) IRecordDo
	Joins(fields ...field.RelationField) IRecordDo
	Preload(fields ...field.RelationField) IRecordDo
	FirstOrInit() (*model.Record, error)
	FirstOrCreate() (*model.Record, error)
	FindByPage(offset int, limit int) (result []*model.Record, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r recordDo) Debug() IRecordDo {
	return r.withDO(r.DO.Debug())
}

func (r recordDo) WithContext(ctx context.Context) IRecordDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r recordDo) ReadDB() IRecordDo {
	return r.Clauses(dbresolver.Read)
}

func (r recordDo) WriteDB() IRecordDo {
	return r.Clauses(dbresolver.Write)
}

func (r recordDo) Session(config *gorm.Session) IRecordDo {
	return r.withDO(r.DO.Session(config))
}

func (r recordDo) Clauses(conds ...clause.Expression) IRecordDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r recordDo) Returning(value interface{}, columns ...string) IRecordDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r recordDo) Not(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r recordDo) Or(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r recordDo) Select(conds ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r recordDo) Where(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r recordDo) Order(conds ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r recordDo) Distinct(cols ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r recordDo) Omit(cols ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r recordDo) Join(table schema.Tabler, on ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r recordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRecordDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r recordDo) RightJoin(table schema.Tabler, on ...field.Expr) IRecordDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r recordDo) Group(cols ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r recordDo) Having(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r recordDo) Limit(limit int) IRecordDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r recordDo) Offset(offset int) IRecordDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r recordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRecordDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r recordDo) Unscoped() IRecordDo {
	return r.withDO(r.DO.Unscoped())
}

func (r recordDo) Create(values ...*model.Record) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r recordDo) CreateInBatches(values []*model.Record, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r recordDo) Save(values ...*model.Record) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r recordDo) First() (*model.Record, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Record), nil
	}
}

func (r recordDo) Take() (*model.Record, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Record), nil
	}
}

func (r recordDo) Last() (*model.Record, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Record), nil
	}
}

func (r recordDo) Find() ([]*model.Record, error) {
	result, err := r.DO.Find()
	return result.([]*model.Record), err
}

func (r recordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Record, err error) {
	buf := make([]*model.Record, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r recordDo) FindInBatches(result *[]*model.Record, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r recordDo) Attrs(attrs ...field.AssignExpr) IRecordDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r recordDo) Assign(attrs ...field.AssignExpr) IRecordDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r recordDo) Joins(fields ...field.RelationField) IRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r recordDo) Preload(fields ...field.RelationField) IRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r recordDo) FirstOrInit() (*model.Record, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Record), nil
	}
}

func (r recordDo) FirstOrCreate() (*model.Record, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Record), nil
	}
}

func (r recordDo) FindByPage(offset int, limit int) (result []*model.Record, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r recordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r recordDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r recordDo) Delete(models ...*model.Record) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *recordDo) withDO(do gen.Dao) *recordDo {
	r.DO = *do.(*gen.DO)
	return r
}
