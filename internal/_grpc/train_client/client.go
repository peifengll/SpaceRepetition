package train_client

import (
	"context"
	"github.com/peifengll/SpaceRepetition/internal/_grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"time"
)

var serverAddr = "127.0.0.1:40000"
var mutex = &sync.Mutex{}

// 这个是单向的

// 拿文件地址去请求,
func TrainDataRequest(filepath string) (bool, error) {
	mutex.Lock()
	defer mutex.Unlock()
	// 连接gRPC服务器
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("connect error to: ", serverAddr)
		return false, err
	}
	defer conn.Close()

	// 实例化一个client对象，传入参数conn
	c := pb.NewTrainerClient(conn)

	// 初始化上下文，设置请求超时时间为1min
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	//延迟关闭请求会话
	defer cancel()

	r, err := c.TrainWeights(ctx, &pb.TrainRequest{RevlogPath: filepath})
	if err != nil {
		log.Println("can not greet to: ", serverAddr)
		return false, err
	} else {
		log.Println("response from server: ", r.GetRevlogPath())
		return true, nil
	}
}
