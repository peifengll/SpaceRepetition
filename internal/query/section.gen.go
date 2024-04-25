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

func newSection(db *gorm.DB, opts ...gen.DOOption) section {
	_section := section{}

	_section.sectionDo.UseDB(db, opts...)
	_section.sectionDo.UseModel(&model.Section{})

	tableName := _section.sectionDo.TableName()
	_section.ALL = field.NewAsterisk(tableName)
	_section.ID = field.NewInt64(tableName, "id")
	_section.CreatedAt = field.NewTime(tableName, "created_at")
	_section.UpdatedAt = field.NewTime(tableName, "updated_at")
	_section.DeletedAt = field.NewField(tableName, "deleted_at")
	_section.Deckid = field.NewInt64(tableName, "deckid")
	_section.Name = field.NewString(tableName, "name")
	_section.UserID = field.NewString(tableName, "user_id")
	_section.CardNum = field.NewInt64(tableName, "card_num")

	_section.fillFieldMap()

	return _section
}

type section struct {
	sectionDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time // 创建时间
	UpdatedAt field.Time // 更新时间
	DeletedAt field.Field
	Deckid    field.Int64
	Name      field.String
	UserID    field.String // 属于哪一个用户
	CardNum   field.Int64  // 章节下卡片的数量

	fieldMap map[string]field.Expr
}

func (s section) Table(newTableName string) *section {
	s.sectionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s section) As(alias string) *section {
	s.sectionDo.DO = *(s.sectionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *section) updateTableName(table string) *section {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Deckid = field.NewInt64(table, "deckid")
	s.Name = field.NewString(table, "name")
	s.UserID = field.NewString(table, "user_id")
	s.CardNum = field.NewInt64(table, "card_num")

	s.fillFieldMap()

	return s
}

func (s *section) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *section) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["deckid"] = s.Deckid
	s.fieldMap["name"] = s.Name
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["card_num"] = s.CardNum
}

func (s section) clone(db *gorm.DB) section {
	s.sectionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s section) replaceDB(db *gorm.DB) section {
	s.sectionDo.ReplaceDB(db)
	return s
}

type sectionDo struct{ gen.DO }

type ISectionDo interface {
	gen.SubQuery
	Debug() ISectionDo
	WithContext(ctx context.Context) ISectionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISectionDo
	WriteDB() ISectionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISectionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISectionDo
	Not(conds ...gen.Condition) ISectionDo
	Or(conds ...gen.Condition) ISectionDo
	Select(conds ...field.Expr) ISectionDo
	Where(conds ...gen.Condition) ISectionDo
	Order(conds ...field.Expr) ISectionDo
	Distinct(cols ...field.Expr) ISectionDo
	Omit(cols ...field.Expr) ISectionDo
	Join(table schema.Tabler, on ...field.Expr) ISectionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISectionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISectionDo
	Group(cols ...field.Expr) ISectionDo
	Having(conds ...gen.Condition) ISectionDo
	Limit(limit int) ISectionDo
	Offset(offset int) ISectionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISectionDo
	Unscoped() ISectionDo
	Create(values ...*model.Section) error
	CreateInBatches(values []*model.Section, batchSize int) error
	Save(values ...*model.Section) error
	First() (*model.Section, error)
	Take() (*model.Section, error)
	Last() (*model.Section, error)
	Find() ([]*model.Section, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Section, err error)
	FindInBatches(result *[]*model.Section, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Section) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISectionDo
	Assign(attrs ...field.AssignExpr) ISectionDo
	Joins(fields ...field.RelationField) ISectionDo
	Preload(fields ...field.RelationField) ISectionDo
	FirstOrInit() (*model.Section, error)
	FirstOrCreate() (*model.Section, error)
	FindByPage(offset int, limit int) (result []*model.Section, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISectionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sectionDo) Debug() ISectionDo {
	return s.withDO(s.DO.Debug())
}

func (s sectionDo) WithContext(ctx context.Context) ISectionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sectionDo) ReadDB() ISectionDo {
	return s.Clauses(dbresolver.Read)
}

func (s sectionDo) WriteDB() ISectionDo {
	return s.Clauses(dbresolver.Write)
}

func (s sectionDo) Session(config *gorm.Session) ISectionDo {
	return s.withDO(s.DO.Session(config))
}

func (s sectionDo) Clauses(conds ...clause.Expression) ISectionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sectionDo) Returning(value interface{}, columns ...string) ISectionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sectionDo) Not(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sectionDo) Or(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sectionDo) Select(conds ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sectionDo) Where(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sectionDo) Order(conds ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sectionDo) Distinct(cols ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sectionDo) Omit(cols ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sectionDo) Join(table schema.Tabler, on ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISectionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sectionDo) RightJoin(table schema.Tabler, on ...field.Expr) ISectionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sectionDo) Group(cols ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sectionDo) Having(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sectionDo) Limit(limit int) ISectionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sectionDo) Offset(offset int) ISectionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISectionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sectionDo) Unscoped() ISectionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sectionDo) Create(values ...*model.Section) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sectionDo) CreateInBatches(values []*model.Section, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sectionDo) Save(values ...*model.Section) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sectionDo) First() (*model.Section, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Section), nil
	}
}

func (s sectionDo) Take() (*model.Section, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Section), nil
	}
}

func (s sectionDo) Last() (*model.Section, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Section), nil
	}
}

func (s sectionDo) Find() ([]*model.Section, error) {
	result, err := s.DO.Find()
	return result.([]*model.Section), err
}

func (s sectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Section, err error) {
	buf := make([]*model.Section, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sectionDo) FindInBatches(result *[]*model.Section, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sectionDo) Attrs(attrs ...field.AssignExpr) ISectionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sectionDo) Assign(attrs ...field.AssignExpr) ISectionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sectionDo) Joins(fields ...field.RelationField) ISectionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sectionDo) Preload(fields ...field.RelationField) ISectionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sectionDo) FirstOrInit() (*model.Section, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Section), nil
	}
}

func (s sectionDo) FirstOrCreate() (*model.Section, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Section), nil
	}
}

func (s sectionDo) FindByPage(offset int, limit int) (result []*model.Section, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sectionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sectionDo) Delete(models ...*model.Section) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sectionDo) withDO(do gen.Dao) *sectionDo {
	s.DO = *do.(*gen.DO)
	return s
}
