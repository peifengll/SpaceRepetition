package global

import "github.com/peifengll/SpaceRepetition/pkg/helper/fsrs"

const pubName = "space_repeat"

// 需要复习的公告

func GetAnnouncementToReadKey() string {
	return pubName + "_announcement"
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
