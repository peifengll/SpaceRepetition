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

func newExportInfo(db *gorm.DB, opts ...gen.DOOption) exportInfo {
	_exportInfo := exportInfo{}

	_exportInfo.exportInfoDo.UseDB(db, opts...)
	_exportInfo.exportInfoDo.UseModel(&model.ExportInfo{})

	tableName := _exportInfo.exportInfoDo.TableName()
	_exportInfo.ALL = field.NewAsterisk(tableName)
	_exportInfo.ID = field.NewInt64(tableName, "id")
	_exportInfo.UserID = field.NewString(tableName, "user_id")
	_exportInfo.RevlogPath = field.NewString(tableName, "revlog_path")
	_exportInfo.TrainDataPath = field.NewString(tableName, "train_data_path")
	_exportInfo.DataStartTime = field.NewTime(tableName, "data_start_time")
	_exportInfo.DataEndTime = field.NewTime(tableName, "data_end_time")
	_exportInfo.ExportDate = field.NewTime(tableName, "export_date")
	_exportInfo.Columns_ = field.NewInt64(tableName, "columns")
	_exportInfo.State = field.NewInt64(tableName, "state")

	_exportInfo.fillFieldMap()

	return _exportInfo
}

// exportInfo 用户导出复习记录的情况（复习记录用于fsrs weight优化）
type exportInfo struct {
	exportInfoDo

	ALL           field.Asterisk
	ID            field.Int64
	UserID        field.String
	RevlogPath    field.String // 导出的revlog的路径
	TrainDataPath field.String // 训练结果的路径
	DataStartTime field.Time   // 复习数据的开始时间
	DataEndTime   field.Time   // 复习数据的结束时间
	ExportDate    field.Time   // 导出revlog的日期
	Columns_      field.Int64  // 导出的revlog有多少列
	State         field.Int64  // 0已完成，1等待中（不少等待就是完成就好）

	fieldMap map[string]field.Expr
}

func (e exportInfo) Table(newTableName string) *exportInfo {
	e.exportInfoDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e exportInfo) As(alias string) *exportInfo {
	e.exportInfoDo.DO = *(e.exportInfoDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *exportInfo) updateTableName(table string) *exportInfo {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.UserID = field.NewString(table, "user_id")
	e.RevlogPath = field.NewString(table, "revlog_path")
	e.TrainDataPath = field.NewString(table, "train_data_path")
	e.DataStartTime = field.NewTime(table, "data_start_time")
	e.DataEndTime = field.NewTime(table, "data_end_time")
	e.ExportDate = field.NewTime(table, "export_date")
	e.Columns_ = field.NewInt64(table, "columns")
	e.State = field.NewInt64(table, "state")

	e.fillFieldMap()

	return e
}

func (e *exportInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *exportInfo) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 9)
	e.fieldMap["id"] = e.ID
	e.fieldMap["user_id"] = e.UserID
	e.fieldMap["revlog_path"] = e.RevlogPath
	e.fieldMap["train_data_path"] = e.TrainDataPath
	e.fieldMap["data_start_time"] = e.DataStartTime
	e.fieldMap["data_end_time"] = e.DataEndTime
	e.fieldMap["export_date"] = e.ExportDate
	e.fieldMap["columns"] = e.Columns_
	e.fieldMap["state"] = e.State
}

func (e exportInfo) clone(db *gorm.DB) exportInfo {
	e.exportInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e exportInfo) replaceDB(db *gorm.DB) exportInfo {
	e.exportInfoDo.ReplaceDB(db)
	return e
}

type exportInfoDo struct{ gen.DO }

type IExportInfoDo interface {
	gen.SubQuery
	Debug() IExportInfoDo
	WithContext(ctx context.Context) IExportInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IExportInfoDo
	WriteDB() IExportInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IExportInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IExportInfoDo
	Not(conds ...gen.Condition) IExportInfoDo
	Or(conds ...gen.Condition) IExportInfoDo
	Select(conds ...field.Expr) IExportInfoDo
	Where(conds ...gen.Condition) IExportInfoDo
	Order(conds ...field.Expr) IExportInfoDo
	Distinct(cols ...field.Expr) IExportInfoDo
	Omit(cols ...field.Expr) IExportInfoDo
	Join(table schema.Tabler, on ...field.Expr) IExportInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IExportInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IExportInfoDo
	Group(cols ...field.Expr) IExportInfoDo
	Having(conds ...gen.Condition) IExportInfoDo
	Limit(limit int) IExportInfoDo
	Offset(offset int) IExportInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IExportInfoDo
	Unscoped() IExportInfoDo
	Create(values ...*model.ExportInfo) error
	CreateInBatches(values []*model.ExportInfo, batchSize int) error
	Save(values ...*model.ExportInfo) error
	First() (*model.ExportInfo, error)
	Take() (*model.ExportInfo, error)
	Last() (*model.ExportInfo, error)
	Find() ([]*model.ExportInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ExportInfo, err error)
	FindInBatches(result *[]*model.ExportInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ExportInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IExportInfoDo
	Assign(attrs ...field.AssignExpr) IExportInfoDo
	Joins(fields ...field.RelationField) IExportInfoDo
	Preload(fields ...field.RelationField) IExportInfoDo
	FirstOrInit() (*model.ExportInfo, error)
	FirstOrCreate() (*model.ExportInfo, error)
	FindByPage(offset int, limit int) (result []*model.ExportInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IExportInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e exportInfoDo) Debug() IExportInfoDo {
	return e.withDO(e.DO.Debug())
}

func (e exportInfoDo) WithContext(ctx context.Context) IExportInfoDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e exportInfoDo) ReadDB() IExportInfoDo {
	return e.Clauses(dbresolver.Read)
}

func (e exportInfoDo) WriteDB() IExportInfoDo {
	return e.Clauses(dbresolver.Write)
}

func (e exportInfoDo) Session(config *gorm.Session) IExportInfoDo {
	return e.withDO(e.DO.Session(config))
}

func (e exportInfoDo) Clauses(conds ...clause.Expression) IExportInfoDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e exportInfoDo) Returning(value interface{}, columns ...string) IExportInfoDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e exportInfoDo) Not(conds ...gen.Condition) IExportInfoDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e exportInfoDo) Or(conds ...gen.Condition) IExportInfoDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e exportInfoDo) Select(conds ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e exportInfoDo) Where(conds ...gen.Condition) IExportInfoDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e exportInfoDo) Order(conds ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e exportInfoDo) Distinct(cols ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e exportInfoDo) Omit(cols ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e exportInfoDo) Join(table schema.Tabler, on ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e exportInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e exportInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e exportInfoDo) Group(cols ...field.Expr) IExportInfoDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e exportInfoDo) Having(conds ...gen.Condition) IExportInfoDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e exportInfoDo) Limit(limit int) IExportInfoDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e exportInfoDo) Offset(offset int) IExportInfoDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e exportInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IExportInfoDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e exportInfoDo) Unscoped() IExportInfoDo {
	return e.withDO(e.DO.Unscoped())
}

func (e exportInfoDo) Create(values ...*model.ExportInfo) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e exportInfoDo) CreateInBatches(values []*model.ExportInfo, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e exportInfoDo) Save(values ...*model.ExportInfo) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e exportInfoDo) First() (*model.ExportInfo, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportInfo), nil
	}
}

func (e exportInfoDo) Take() (*model.ExportInfo, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportInfo), nil
	}
}

func (e exportInfoDo) Last() (*model.ExportInfo, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportInfo), nil
	}
}

func (e exportInfoDo) Find() ([]*model.ExportInfo, error) {
	result, err := e.DO.Find()
	return result.([]*model.ExportInfo), err
}

func (e exportInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ExportInfo, err error) {
	buf := make([]*model.ExportInfo, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e exportInfoDo) FindInBatches(result *[]*model.ExportInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e exportInfoDo) Attrs(attrs ...field.AssignExpr) IExportInfoDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e exportInfoDo) Assign(attrs ...field.AssignExpr) IExportInfoDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e exportInfoDo) Joins(fields ...field.RelationField) IExportInfoDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e exportInfoDo) Preload(fields ...field.RelationField) IExportInfoDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e exportInfoDo) FirstOrInit() (*model.ExportInfo, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportInfo), nil
	}
}

func (e exportInfoDo) FirstOrCreate() (*model.ExportInfo, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportInfo), nil
	}
}

func (e exportInfoDo) FindByPage(offset int, limit int) (result []*model.ExportInfo, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e exportInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e exportInfoDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e exportInfoDo) Delete(models ...*model.ExportInfo) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *exportInfoDo) withDO(do gen.Dao) *exportInfoDo {
	e.DO = *do.(*gen.DO)
	return e
}