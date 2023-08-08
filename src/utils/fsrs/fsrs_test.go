package fsrs

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSimpleWorkWell(t *testing.T) {
	Example()
}
func Space() {
	fmt.Println("==========================================================")
	fmt.Println("==========================================================")
}
func Example() {
	// 这算法怎么用呢， 新建一张新的卡片， 对其进行复习， 将最终的难度 给出， 重新插入数据库 over
	// 其实就是这么简单
	create, err := os.Create("cc.json")
	if err != nil {
		return
	}
	defer create.Close()
	p := DefaultParam()
	card := NewCard()
	now := time.Date(2022, 11, 29, 12, 30, 0, 0, time.UTC)
	schedulingCards := p.Repeat(card, now)
	schedule, _ := json.MarshalIndent(schedulingCards, "", "\t")
	create.WriteString(string(schedule))
	fmt.Println(string(schedule))

	Space()

	card = schedulingCards[Good].Card
	now = card.Due
	schedulingCards = p.Repeat(card, now)

	schedule, _ = json.MarshalIndent(schedulingCards, "", "\t")
	create.WriteString(string(schedule))
	fmt.Println(string(schedule))

	Space()

	card = schedulingCards[Good].Card
	now = card.Due
	schedulingCards = p.Repeat(card, now)
	schedule, _ = json.Marshal(schedulingCards)
	create.WriteString(string(schedule))
	fmt.Println(string(schedule))

	Space()

	card = schedulingCards[Again].Card
	now = card.Due
	schedulingCards = p.Repeat(card, now)
	schedule, _ = json.MarshalIndent(schedulingCards, "", "\t")

	create.WriteString(string(schedule))
	fmt.Println(string(schedule))

	Space()

	card = schedulingCards[Good].Card
	now = card.Due
	schedulingCards = p.Repeat(card, now)
	schedule, _ = json.MarshalIndent(schedulingCards, "", "\t")
	create.WriteString(string(schedule))
	fmt.Println(string(schedule))
}
