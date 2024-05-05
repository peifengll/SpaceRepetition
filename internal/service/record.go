package service

import (
	"fmt"
	"github.com/gocarina/gocsv"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/global"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type RecordService interface {
	GetRecord(id int64) (*model.Record, error)
	ExportInfoRecord(ran *v1.ExportRevlogCsvReq, userid string) (bool, error)
	ExportInfoRecordOp(data v1.ExportTask) (bool, error)
	CreateExportInfo(info *model.ExportInfo) error
}

func NewRecordService(que *query.Query, service *Service, recordRepository repository.RecordRepository) RecordService {
	return &recordService{
		Service: service,
		query:   que,

		recordRepository: recordRepository,
	}
}

type recordService struct {
	*Service
	query            *query.Query
	recordRepository repository.RecordRepository
}

func (s *recordService) GetRecord(id int64) (*model.Record, error) {
	return s.recordRepository.FirstById(id)
}

// 根据card删除 记录
// todo  这中要排除报错为 没有那个数据的情况
func delRecordTx(kid int64, tx *query.QueryTx) error {
	_, err := tx.Record.Where(tx.Record.KnowledgeID.Eq(kid)).Delete()
	return err
}

// 拿到数据 之后 ，
func (s *recordService) ExportInfoRecord(ran *v1.ExportRevlogCsvReq, userid string) (bool, error) {
	infos, err := s.recordRepository.QueryDataByTimeRange(userid, ran.Start, ran.End)
	if err != nil {
		return false, err
	}
	// 拿userid 去取一个path

	nowTime := strconv.FormatInt(time.Now().Unix(), 10)
	filePath := global.ExportPrefix + "/" + userid + "/" + nowTime + ".csv"
	// 检查路径是否已经存在，如果不存在则创建所有必要的中间目录
	if err = os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		fmt.Println("创建目录失败：", err)
		return false, err
	}
	csvFile, err := os.Create(filePath)
	defer csvFile.Close()
	if err != nil {
		return false, err
	}
	err = gocsv.MarshalFile(&infos, csvFile)
	if err != nil {
		return false, err
	}
	// 这里文件也导出了，向exportinfo也整一个数据
	return true, nil
}

func (s *recordService) CreateExportInfo(info *model.ExportInfo) error {
	return s.query.ExportInfo.Create(info)
}

func (s *recordService) ExportInfoRecordOp(data v1.ExportTask) (bool, error) {
	infos, err := s.recordRepository.QueryDataByTimeRange(data.UserId, data.Start, data.End)
	if err != nil {
		return false, err
	}
	// 拿userid 去取一个path
	nowTime := strconv.FormatInt(time.Now().Unix(), 10)
	filePath := global.ExportPrefix + "/" + data.UserId + "/" + nowTime + ".csv"
	// 检查路径是否已经存在，如果不存在则创建所有必要的中间目录
	if err = os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		fmt.Println("创建目录失败：", err)
		return false, err
	}
	csvFile, err := os.Create(filePath)
	defer csvFile.Close()
	if err != nil {
		return false, err
	}
	err = gocsv.MarshalFile(&infos, csvFile)
	if err != nil {
		return false, err
	}
	// 这里文件也导出了，向exportinfo也整一个数据
	exp := s.query.ExportInfo
	_, err = exp.Where(exp.ID.Eq(data.ExportId)).Updates(map[string]any{
		"revlog_path": filePath,
		"columns":     len(infos),
		"state":       0,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
