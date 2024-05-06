package global_

import (
	"github.com/peifengll/SpaceRepetition/pkg/helper/fsrs"
	"time"
)

const pubName = "space_repeat"

var Today time.Time
var Schedule string

func SetToday(now time.Time) {
	Today = now

}
func SetSchedule(now time.Time) {
	Schedule = now.Format("2006-01-02")
}

// 需要复习的公告

func GetAnnouncementToReadKey() string {
	return pubName + "_announcement"
}

func GetReviewOpNumberKey(userid, sc string) string {
	return pubName + "_" + userid + "_" + sc + "_rate"
}

func GetFsrsParmsKey(userid string) string {
	return pubName + userid + "_fsrs"
}

// 复习记录相关
func GetReviewOpNumKey(userid string, i fsrs.Rating) string {
	suffix := ""
	switch i {
	case fsrs.Good:
		suffix = "_Good_Num"
	case fsrs.Easy:
		suffix = "_Easy_Num"
	case fsrs.Again:
		suffix = "_Again_Num"
	case fsrs.Hard:
		suffix = "_Hard_Num"
	}
	return pubName + "_" + userid + suffix
}

// 今天复习的卡片集合
func GetReviewCardsSetKey(userid, sc string) string {
	return pubName + "_" + userid + "_" + sc + "_cards"
}

// 今天有多少个人在复习卡片，要进行统计的
func GetReviewersTodayKey(sc string) string {
	return pubName + "_" + sc + "_reviewers"
}
