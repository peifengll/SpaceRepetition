package server

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/global_"
	"github.com/peifengll/SpaceRepetition/internal/model"
	"github.com/peifengll/SpaceRepetition/internal/query"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

type Cleaner struct {
	rds *redis.Client
	que *query.Query
}

func NewCleaner(rds *redis.Client,
	que *query.Query) *Cleaner {
	return &Cleaner{
		rds: rds,
		que: que,
	}
}

var clear_signal = make(chan struct{})

var yesterday time.Time
var yesterdaySchedule string

func (c *Cleaner) Start(ctx context.Context) error {
	yesterday = time.Now()
	if yesterday.Hour() < 4 {
		yesterday = yesterday.AddDate(0, 0, -1)
	}
	global_.SetToday(yesterday)
	global_.SetSchedule(yesterday)
	go func() {
		for {
			select {
			case <-clear_signal:
				break
			default:
				c.checkDate()
			}
		}
	}()
	return nil
}

func subDay(date1, date2 time.Time) int {
	// 将时分秒截断
	date1 = date1.Truncate(24 * time.Hour)
	date2 = date2.Truncate(24 * time.Hour)

	// 计算日期差异
	difference := date1.Sub(date2)
	daysDifference := int(difference.Hours() / 24)
	return daysDifference
}

func (c *Cleaner) checkDate() {
	// 判断等不等第二天的4
	now := time.Now()
	// 就可以执行了
	if now.Hour() >= 4 && subDay(now, yesterday) == 1 {
		yesterdaySchedule = global_.Schedule
		//已经是明天了
		global_.SetSchedule(now)
		// 完成一次清除和数据库保存
		c.scheduleClearDataAtFourAM()
		yesterday = now
		global_.SetToday(now)
	} else {
		time.Sleep(time.Second * 5)
	}
}

func (c *Cleaner) scheduleClearDataAtFourAM() {
	// 获取数据库，清楚数据库
	ctx := context.Background()
	userids, err := c.rds.SMembers(ctx, global_.GetReviewersTodayKey(yesterdaySchedule)).Result()
	if err != nil {
		log.Println(err)
		return
	}
	c.rds.Del(ctx, global_.GetReviewersTodayKey(yesterdaySchedule))
	// 从现在开始 处理今天有复习的每个人的数据
	for _, userid := range userids {
		// 各个复习情况的数量
		infos, err := c.rds.HGetAll(ctx, global_.GetReviewOpNumberKey(userid, yesterdaySchedule)).Result()
		if err != nil {
			log.Println(err)
			return
		}
		// 获取复习数目
		cardNum, err := c.rds.SCard(ctx, global_.GetReviewCardsSetKey(userid, yesterdaySchedule)).Result()
		if err != nil {
			log.Println(err)
			return
		}
		//
		po := &model.DayReviewStatistic{
			CardNum:    cardNum,
			RecordDate: yesterday,
			UserID:     userid,
		}

		if val, ok := infos["Good"]; ok {
			num, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return
			}
			po.GoodNum = num
		}
		if val, ok := infos["Again"]; ok {
			num, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return
			}
			po.AgainNum = num
		}
		if val, ok := infos["Hard"]; ok {
			num, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return
			}
			po.HardNum = num
		}
		if val, ok := infos["Easy"]; ok {
			num, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return
			}
			po.EasyNum = num
		}

		err = c.que.DayReviewStatistic.Create(po)
		if err != nil {
			log.Println(err)
			return
		}
		// 清除redis
		c.rds.Del(ctx, global_.GetReviewersTodayKey(yesterdaySchedule))
		c.rds.Del(ctx, global_.GetReviewCardsSetKey(userid, yesterdaySchedule))
	}
}

func (c *Cleaner) Stop(ctx context.Context) error {
	log.Println("cleaner stop...")
	clear_signal <- struct{}{}
	return nil
}
