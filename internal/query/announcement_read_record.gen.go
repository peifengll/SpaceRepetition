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

func newAnnouncementReadRecord(db *gorm.DB, opts ...gen.DOOption) announcementReadRecord {
	_announcementReadRecord := announcementReadRecord{}

	_announcementReadRecord.announcementReadRecordDo.UseDB(db, opts...)
	_announcementReadRecord.announcementReadRecordDo.UseModel(&model.AnnouncementReadRecord{})

	tableName := _announcementReadRecord.announcementReadRecordDo.TableName()
	_announcementReadRecord.ALL = field.NewAsterisk(tableName)
	_announcementReadRecord.ID = field.NewInt64(tableName, "id")
	_announcementReadRecord.UserID = field.NewString(tableName, "user_id")
	_announcementReadRecord.AnnouncementID = field.NewInt64(tableName, "announcement_id")

	_announcementReadRecord.fillFieldMap()

	return _announcementReadRecord
}

// announcementReadRecord 记录用户是否阅读某一条公告
type announcementReadRecord struct {
	announcementReadRecordDo

	ALL            field.Asterisk
	ID             field.Int64
	UserID         field.String // 用户的userid
	AnnouncementID field.Int64  // 公告id

	fieldMap map[string]field.Expr
}

func (a announcementReadRecord) Table(newTableName string) *announcementReadRecord {
	a.announcementReadRecordDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a announcementReadRecord) As(alias string) *announcementReadRecord {
	a.announcementReadRecordDo.DO = *(a.announcementReadRecordDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *announcementReadRecord) updateTableName(table string) *announcementReadRecord {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UserID = field.NewString(table, "user_id")
	a.AnnouncementID = field.NewInt64(table, "announcement_id")

	a.fillFieldMap()

	return a
}

func (a *announcementReadRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *announcementReadRecord) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 3)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["announcement_id"] = a.AnnouncementID
}

func (a announcementReadRecord) clone(db *gorm.DB) announcementReadRecord {
	a.announcementReadRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a announcementReadRecord) replaceDB(db *gorm.DB) announcementReadRecord {
	a.announcementReadRecordDo.ReplaceDB(db)
	return a
}

type announcementReadRecordDo struct{ gen.DO }

type IAnnouncementReadRecordDo interface {
	gen.SubQuery
	Debug() IAnnouncementReadRecordDo
	WithContext(ctx context.Context) IAnnouncementReadRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAnnouncementReadRecordDo
	WriteDB() IAnnouncementReadRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAnnouncementReadRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAnnouncementReadRecordDo
	Not(conds ...gen.Condition) IAnnouncementReadRecordDo
	Or(conds ...gen.Condition) IAnnouncementReadRecordDo
	Select(conds ...field.Expr) IAnnouncementReadRecordDo
	Where(conds ...gen.Condition) IAnnouncementReadRecordDo
	Order(conds ...field.Expr) IAnnouncementReadRecordDo
	Distinct(cols ...field.Expr) IAnnouncementReadRecordDo
	Omit(cols ...field.Expr) IAnnouncementReadRecordDo
	Join(table schema.Tabler, on ...field.Expr) IAnnouncementReadRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAnnouncementReadRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAnnouncementReadRecordDo
	Group(cols ...field.Expr) IAnnouncementReadRecordDo
	Having(conds ...gen.Condition) IAnnouncementReadRecordDo
	Limit(limit int) IAnnouncementReadRecordDo
	Offset(offset int) IAnnouncementReadRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAnnouncementReadRecordDo
	Unscoped() IAnnouncementReadRecordDo
	Create(values ...*model.AnnouncementReadRecord) error
	CreateInBatches(values []*model.AnnouncementReadRecord, batchSize int) error
	Save(values ...*model.AnnouncementReadRecord) error
	First() (*model.AnnouncementReadRecord, error)
	Take() (*model.AnnouncementReadRecord, error)
	Last() (*model.AnnouncementReadRecord, error)
	Find() ([]*model.AnnouncementReadRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AnnouncementReadRecord, err error)
	FindInBatches(result *[]*model.AnnouncementReadRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AnnouncementReadRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAnnouncementReadRecordDo
	Assign(attrs ...field.AssignExpr) IAnnouncementReadRecordDo
	Joins(fields ...field.RelationField) IAnnouncementReadRecordDo
	Preload(fields ...field.RelationField) IAnnouncementReadRecordDo
	FirstOrInit() (*model.AnnouncementReadRecord, error)
	FirstOrCreate() (*model.AnnouncementReadRecord, error)
	FindByPage(offset int, limit int) (result []*model.AnnouncementReadRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAnnouncementReadRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a announcementReadRecordDo) Debug() IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Debug())
}

func (a announcementReadRecordDo) WithContext(ctx context.Context) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a announcementReadRecordDo) ReadDB() IAnnouncementReadRecordDo {
	return a.Clauses(dbresolver.Read)
}

func (a announcementReadRecordDo) WriteDB() IAnnouncementReadRecordDo {
	return a.Clauses(dbresolver.Write)
}

func (a announcementReadRecordDo) Session(config *gorm.Session) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Session(config))
}

func (a announcementReadRecordDo) Clauses(conds ...clause.Expression) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a announcementReadRecordDo) Returning(value interface{}, columns ...string) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a announcementReadRecordDo) Not(conds ...gen.Condition) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a announcementReadRecordDo) Or(conds ...gen.Condition) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a announcementReadRecordDo) Select(conds ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a announcementReadRecordDo) Where(conds ...gen.Condition) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a announcementReadRecordDo) Order(conds ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a announcementReadRecordDo) Distinct(cols ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a announcementReadRecordDo) Omit(cols ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a announcementReadRecordDo) Join(table schema.Tabler, on ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a announcementReadRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a announcementReadRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a announcementReadRecordDo) Group(cols ...field.Expr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a announcementReadRecordDo) Having(conds ...gen.Condition) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a announcementReadRecordDo) Limit(limit int) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a announcementReadRecordDo) Offset(offset int) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a announcementReadRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a announcementReadRecordDo) Unscoped() IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Unscoped())
}

func (a announcementReadRecordDo) Create(values ...*model.AnnouncementReadRecord) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a announcementReadRecordDo) CreateInBatches(values []*model.AnnouncementReadRecord, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a announcementReadRecordDo) Save(values ...*model.AnnouncementReadRecord) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a announcementReadRecordDo) First() (*model.AnnouncementReadRecord, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnouncementReadRecord), nil
	}
}

func (a announcementReadRecordDo) Take() (*model.AnnouncementReadRecord, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnouncementReadRecord), nil
	}
}

func (a announcementReadRecordDo) Last() (*model.AnnouncementReadRecord, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnouncementReadRecord), nil
	}
}

func (a announcementReadRecordDo) Find() ([]*model.AnnouncementReadRecord, error) {
	result, err := a.DO.Find()
	return result.([]*model.AnnouncementReadRecord), err
}

func (a announcementReadRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AnnouncementReadRecord, err error) {
	buf := make([]*model.AnnouncementReadRecord, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a announcementReadRecordDo) FindInBatches(result *[]*model.AnnouncementReadRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a announcementReadRecordDo) Attrs(attrs ...field.AssignExpr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a announcementReadRecordDo) Assign(attrs ...field.AssignExpr) IAnnouncementReadRecordDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a announcementReadRecordDo) Joins(fields ...field.RelationField) IAnnouncementReadRecordDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a announcementReadRecordDo) Preload(fields ...field.RelationField) IAnnouncementReadRecordDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a announcementReadRecordDo) FirstOrInit() (*model.AnnouncementReadRecord, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnouncementReadRecord), nil
	}
}

func (a announcementReadRecordDo) FirstOrCreate() (*model.AnnouncementReadRecord, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnouncementReadRecord), nil
	}
}

func (a announcementReadRecordDo) FindByPage(offset int, limit int) (result []*model.AnnouncementReadRecord, count int64, err error) {
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

func (a announcementReadRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a announcementReadRecordDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a announcementReadRecordDo) Delete(models ...*model.AnnouncementReadRecord) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *announcementReadRecordDo) withDO(do gen.Dao) *announcementReadRecordDo {
	a.DO = *do.(*gen.DO)
	return a
}
