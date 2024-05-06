package etask

import (
	"context"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

type Stack struct {
	element []v1.ExportTask
}

// NewStack 创建一个新栈
func NewStack() *Stack {
	return &Stack{}
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	if len(s.element) == 0 {
		return true
	} else {
		return false
	}
}

// GetStackLength 求栈的长度
func (s *Stack) GetStackLength() int {
	return len(s.element)
}

// Push 进栈操作
func (s *Stack) Push(value v1.ExportTask) {
	s.element = append(s.element, value)
}

// Top 获取栈顶元素
func (s *Stack) Top() (v1.ExportTask, bool) {
	if s.IsEmpty() {
		return v1.ExportTask{}, false
	} else {
		return s.element[s.GetStackLength()-1], true
	}
}

// Pop 出栈操作
func (s *Stack) Pop() (v1.ExportTask, bool) {
	v, ok := s.Top()
	if s.IsEmpty() {
		return v1.ExportTask{}, false
	} else {
		s.element = s.element[:s.GetStackLength()-1]
	}
	return v, ok
}

var TM *TaskManager

type TaskManager struct {
	rds      *redis.Client
	mapLock  *sync.RWMutex
	taskList *Stack
	record   service.RecordService
}

func NewTaskManager(rd *redis.Client, recordService service.RecordService) *TaskManager {
	TM = &TaskManager{
		rds:      rd,
		mapLock:  &sync.RWMutex{},
		taskList: NewStack(),
		record:   recordService,
	}
	return TM
}

var signa = make(chan struct{})

func (r *TaskManager) Start(ctx context.Context) error {
	go func() {
		tk := time.NewTicker(time.Millisecond * 50)
		for {
			select {
			case <-signa:
				tk.Stop()
			case <-tk.C:
				continue
			default:
				for {
					// 检查队列里边有没有
					if r.taskList.IsEmpty() {
						break
					}
					taskinfo, ok := r.taskList.Pop()
					if ok {
						// 可以进行操作了
						r.record.ExportInfoRecordOp(taskinfo)
					}
				}
			}
		}
	}()
	return nil
}

func (r *TaskManager) Stop(ctx context.Context) error {
	signa <- struct{}{}
	log.Println("task manager abort")
	return nil
}

func (r *TaskManager) AddTask(tsk v1.ExportTask) (bool, error) {
	r.mapLock.Lock()
	defer r.mapLock.Unlock()
	// 要去创建一个信息了 ，然后得到id 然后加如队列
	if r.taskList.GetStackLength() < 100 {
		info := &model.ExportInfo{
			UserID:        tsk.UserId,
			DataStartTime: tsk.Start,
			DataEndTime:   tsk.End,
			State:         1, // 等待中
			ExportDate:    time.Now(),
		}
		err := r.record.CreateExportInfo(info)
		if err != nil {
			return false, err
		}
		tsk.ExportId = info.ID
		r.taskList.Push(tsk)
		return true, err
	} else {
		// 满了
		return false, nil
	}
}
