/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/14 1:32 下午
@Desc
*/
package grpc

import (
	pb "codeup.aliyun.com/qimao/go-freebook/bookadv/web/grpc/proto/adv"
	"codeup.aliyun.com/qimao/go-freebook/bookadv/web/pkg/grpc/client"
	"context"
	"fmt"
	"testing"
)

func TestGrpc(t *testing.T) {

	fmt.Println("discover")
	grpcClient, _, e := client.NewService(context.Background(), client.Config{
		//Address:  "localhost:8080",
		//Insecure: true,
		Address:  "grpc-adv.wtzw.com:443",
		Insecure: false,
	})
	t.Log(e)
	reqData := pb.DiscoverRequest{
		Platform:   "android",
		AppVersion: "50000",
		//Channel:       "unknown12",
		//Gender:        "0",
		//ApplicationId: "com.kmxs.reader",
		//SysVer:        "9",
		//TeenyMode:     "0",
		//Brand:         "vivo",
		//Model:         "V1813BA",
		//NetEnv:        "1",
		//NoPermiss:     "0",
		//DeviceId:      "20190712141352c5590aeffbbec19d7b474f3958f971c8018c311326c075ea",
		//VipExpireAt:   "0",
		//Source:        "1",
	}
	rs, err := grpcClient.GetDiscovers2(&reqData)
	t.Log(rs, err)
}
