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

func newKnowledge(db *gorm.DB, opts ...gen.DOOption) knowledge {
	_knowledge := knowledge{}

	_knowledge.knowledgeDo.UseDB(db, opts...)
	_knowledge.knowledgeDo.UseModel(&model.Knowledge{})

	tableName := _knowledge.knowledgeDo.TableName()
	_knowledge.ALL = field.NewAsterisk(tableName)
	_knowledge.ID = field.NewInt64(tableName, "id")
	_knowledge.CreatedAt = field.NewTime(tableName, "created_at")
	_knowledge.UpdatedAt = field.NewTime(tableName, "updated_at")
	_knowledge.DeletedAt = field.NewField(tableName, "deleted_at")
	_knowledge.Font = field.NewString(tableName, "font")
	_knowledge.Originfont = field.NewString(tableName, "originfont")
	_knowledge.Back = field.NewString(tableName, "back")
	_knowledge.Onlearning = field.NewInt64(tableName, "onlearning")
	_knowledge.Typeof = field.NewInt64(tableName, "typeof")
	_knowledge.Deckid = field.NewInt64(tableName, "deckid")
	_knowledge.Skilled = field.NewFloat64(tableName, "skilled")
	_knowledge.Sectionid = field.NewInt64(tableName, "sectionid")
	_knowledge.UserID = field.NewString(tableName, "user_id")

	_knowledge.fillFieldMap()

	return _knowledge
}

type knowledge struct {
	knowledgeDo

	ALL        field.Asterisk
	ID         field.Int64
	CreatedAt  field.Time // 创建时间
	UpdatedAt  field.Time // 更新时间
	DeletedAt  field.Field
	Font       field.String
	Originfont field.String
	Back       field.String
	Onlearning field.Int64   // 0:未开始复习;1:复习中;2:复习周期长，记录为复习完成，不再加入复习调度
	Typeof     field.Int64   // 卡片类型
	Deckid     field.Int64   // 牌组id
	Skilled    field.Float64 // 熟练度
	Sectionid  field.Int64   // 章节id
	UserID     field.String  // 用户id

	fieldMap map[string]field.Expr
}

func (k knowledge) Table(newTableName string) *knowledge {
	k.knowledgeDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k knowledge) As(alias string) *knowledge {
	k.knowledgeDo.DO = *(k.knowledgeDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *knowledge) updateTableName(table string) *knowledge {
	k.ALL = field.NewAsterisk(table)
	k.ID = field.NewInt64(table, "id")
	k.CreatedAt = field.NewTime(table, "created_at")
	k.UpdatedAt = field.NewTime(table, "updated_at")
	k.DeletedAt = field.NewField(table, "deleted_at")
	k.Font = field.NewString(table, "font")
	k.Originfont = field.NewString(table, "originfont")
	k.Back = field.NewString(table, "back")
	k.Onlearning = field.NewInt64(table, "onlearning")
	k.Typeof = field.NewInt64(table, "typeof")
	k.Deckid = field.NewInt64(table, "deckid")
	k.Skilled = field.NewFloat64(table, "skilled")
	k.Sectionid = field.NewInt64(table, "sectionid")
	k.UserID = field.NewString(table, "user_id")

	k.fillFieldMap()

	return k
}

func (k *knowledge) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *knowledge) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 13)
	k.fieldMap["id"] = k.ID
	k.fieldMap["created_at"] = k.CreatedAt
	k.fieldMap["updated_at"] = k.UpdatedAt
	k.fieldMap["deleted_at"] = k.DeletedAt
	k.fieldMap["font"] = k.Font
	k.fieldMap["originfont"] = k.Originfont
	k.fieldMap["back"] = k.Back
	k.fieldMap["onlearning"] = k.Onlearning
	k.fieldMap["typeof"] = k.Typeof
	k.fieldMap["deckid"] = k.Deckid
	k.fieldMap["skilled"] = k.Skilled
	k.fieldMap["sectionid"] = k.Sectionid
	k.fieldMap["user_id"] = k.UserID
}

func (k knowledge) clone(db *gorm.DB) knowledge {
	k.knowledgeDo.ReplaceConnPool(db.Statement.ConnPool)
	return k
}

func (k knowledge) replaceDB(db *gorm.DB) knowledge {
	k.knowledgeDo.ReplaceDB(db)
	return k
}

type knowledgeDo struct{ gen.DO }

type IKnowledgeDo interface {
	gen.SubQuery
	Debug() IKnowledgeDo
	WithContext(ctx context.Context) IKnowledgeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IKnowledgeDo
	WriteDB() IKnowledgeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IKnowledgeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IKnowledgeDo
	Not(conds ...gen.Condition) IKnowledgeDo
	Or(conds ...gen.Condition) IKnowledgeDo
	Select(conds ...field.Expr) IKnowledgeDo
	Where(conds ...gen.Condition) IKnowledgeDo
	Order(conds ...field.Expr) IKnowledgeDo
	Distinct(cols ...field.Expr) IKnowledgeDo
	Omit(cols ...field.Expr) IKnowledgeDo
	Join(table schema.Tabler, on ...field.Expr) IKnowledgeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IKnowledgeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IKnowledgeDo
	Group(cols ...field.Expr) IKnowledgeDo
	Having(conds ...gen.Condition) IKnowledgeDo
	Limit(limit int) IKnowledgeDo
	Offset(offset int) IKnowledgeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IKnowledgeDo
	Unscoped() IKnowledgeDo
	Create(values ...*model.Knowledge) error
	CreateInBatches(values []*model.Knowledge, batchSize int) error
	Save(values ...*model.Knowledge) error
	First() (*model.Knowledge, error)
	Take() (*model.Knowledge, error)
	Last() (*model.Knowledge, error)
	Find() ([]*model.Knowledge, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Knowledge, err error)
	FindInBatches(result *[]*model.Knowledge, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Knowledge) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IKnowledgeDo
	Assign(attrs ...field.AssignExpr) IKnowledgeDo
	Joins(fields ...field.RelationField) IKnowledgeDo
	Preload(fields ...field.RelationField) IKnowledgeDo
	FirstOrInit() (*model.Knowledge, error)
	FirstOrCreate() (*model.Knowledge, error)
	FindByPage(offset int, limit int) (result []*model.Knowledge, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IKnowledgeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (k knowledgeDo) Debug() IKnowledgeDo {
	return k.withDO(k.DO.Debug())
}

func (k knowledgeDo) WithContext(ctx context.Context) IKnowledgeDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k knowledgeDo) ReadDB() IKnowledgeDo {
	return k.Clauses(dbresolver.Read)
}

func (k knowledgeDo) WriteDB() IKnowledgeDo {
	return k.Clauses(dbresolver.Write)
}

func (k knowledgeDo) Session(config *gorm.Session) IKnowledgeDo {
	return k.withDO(k.DO.Session(config))
}

func (k knowledgeDo) Clauses(conds ...clause.Expression) IKnowledgeDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k knowledgeDo) Returning(value interface{}, columns ...string) IKnowledgeDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k knowledgeDo) Not(conds ...gen.Condition) IKnowledgeDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k knowledgeDo) Or(conds ...gen.Condition) IKnowledgeDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k knowledgeDo) Select(conds ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k knowledgeDo) Where(conds ...gen.Condition) IKnowledgeDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k knowledgeDo) Order(conds ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k knowledgeDo) Distinct(cols ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k knowledgeDo) Omit(cols ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k knowledgeDo) Join(table schema.Tabler, on ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k knowledgeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k knowledgeDo) RightJoin(table schema.Tabler, on ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k knowledgeDo) Group(cols ...field.Expr) IKnowledgeDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k knowledgeDo) Having(conds ...gen.Condition) IKnowledgeDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k knowledgeDo) Limit(limit int) IKnowledgeDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k knowledgeDo) Offset(offset int) IKnowledgeDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k knowledgeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IKnowledgeDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k knowledgeDo) Unscoped() IKnowledgeDo {
	return k.withDO(k.DO.Unscoped())
}

func (k knowledgeDo) Create(values ...*model.Knowledge) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k knowledgeDo) CreateInBatches(values []*model.Knowledge, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k knowledgeDo) Save(values ...*model.Knowledge) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k knowledgeDo) First() (*model.Knowledge, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Knowledge), nil
	}
}

func (k knowledgeDo) Take() (*model.Knowledge, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Knowledge), nil
	}
}

func (k knowledgeDo) Last() (*model.Knowledge, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Knowledge), nil
	}
}

func (k knowledgeDo) Find() ([]*model.Knowledge, error) {
	result, err := k.DO.Find()
	return result.([]*model.Knowledge), err
}

func (k knowledgeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Knowledge, err error) {
	buf := make([]*model.Knowledge, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k knowledgeDo) FindInBatches(result *[]*model.Knowledge, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k knowledgeDo) Attrs(attrs ...field.AssignExpr) IKnowledgeDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k knowledgeDo) Assign(attrs ...field.AssignExpr) IKnowledgeDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k knowledgeDo) Joins(fields ...field.RelationField) IKnowledgeDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k knowledgeDo) Preload(fields ...field.RelationField) IKnowledgeDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k knowledgeDo) FirstOrInit() (*model.Knowledge, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Knowledge), nil
	}
}

func (k knowledgeDo) FirstOrCreate() (*model.Knowledge, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Knowledge), nil
	}
}

func (k knowledgeDo) FindByPage(offset int, limit int) (result []*model.Knowledge, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k knowledgeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k knowledgeDo) Scan(result interface{}) (err error) {
	return k.DO.Scan(result)
}

func (k knowledgeDo) Delete(models ...*model.Knowledge) (result gen.ResultInfo, err error) {
	return k.DO.Delete(models)
}

func (k *knowledgeDo) withDO(do gen.Dao) *knowledgeDo {
	k.DO = *do.(*gen.DO)
	return k
}
