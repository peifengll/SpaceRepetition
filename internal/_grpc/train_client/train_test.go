package train_client

import (
	"log"
	"sync"
	"testing"
)

func TestTrainDataRequest(t *testing.T) {
	result, err := TrainDataRequest("D:\\work\\code\\go\\SpaceRepetition\\storage\\test_data\\revlog3.csv")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
}

// 目前的这个 优化器不支持并发，新的地址会覆盖旧的地址，所有文件会被导入到一个地方
// 所以还是加了个锁，一次只运行一个
func TestSupervene(t *testing.T) {
	//开两个协程
	group := &sync.WaitGroup{}
	group.Add(2)
	gofun1 := func() {
		log.Println("go fun2 start:")
		result, err := TrainDataRequest("D:\\work\\code\\go\\SpaceRepetition\\storage\\test_data\\revlog3.csv")
		if err != nil {
			log.Println(err)
		}
		log.Println(result)
		group.Done()
	}
	gofun2 := func() {
		log.Println("go fun1 start:")
		result, err := TrainDataRequest("D:\\work\\code\\go\\SpaceRepetition\\storage\\test_data\\test_only_1000_data.csv")
		if err != nil {
			log.Println(err)
		}
		log.Println(result)
		group.Done()
	}
	go gofun1()
	go gofun2()
	group.Wait()

}
