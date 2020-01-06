package main

import (
	"crawler_distributed/config"
	"crawler_distributed/rpc"
	"crawler_distributed/workerRpc"
	"log"
	"testing"
	"time"
)

func TestWorkService(t *testing.T) {
	const host = ":1235"
	//启服务
	go rpc.ServeRpc(workerRpc.WorkService{}, host)
	time.Sleep(time.Second)

	//启client
	client, err := rpc.NewRpcClient(host)
	if err != nil {
		panic(err)
	}

	//client调用
	var result workerRpc.RequestResult
	err = client.Call(config.WorkerService, workerRpc.Request{
		Url: "https://album.zhenai.com/u/1662184411",
		Handle: workerRpc.SerializedParser{
			Name: config.UserInfoParser,
			Args: nil,
		},
	}, &result)
	if err != nil {
		panic(err)
	} else {
		requestResult := workerRpc.DeserializeRequestResult(result)
		log.Printf("%v", requestResult)
	}

	//验证数据
	//item := engine.Item{
	//	Url:  "https://album.zhenai.com/u/1662184411",
	//	Type: "",
	//	Id:   "1662184411",
	//	DoType: model.User{
	//		Avatar:         "https://photo.zastatic.com/images/photo/415547/1662184411/10203260896941091.jpg",
	//		NickName:       "",
	//		Sex:            "",
	//		InnerMonologue: "将一段记忆在风雨中渐渐淡忘，在时光的转角处抹去这段忧伤，让岁月不再诉说心酸的过往，让风月从此不再黯然神伤。",
	//		MaritalStatus:  "未婚",
	//		Age:            27,
	//		Figure:         "苗条",
	//		Constellation:  "天蝎座(10.23-11.21)",
	//		Height:         160,
	//		Weight:         41,
	//		WorkPlace:      "阿坝",
	//		MonthlyIncome:  "5001-8000元",
	//		Job:            "",
	//		Education:      "大学本科",
	//		Nation:         "羌族",
	//		Hometown:       "四川阿坝",
	//		Smoking:        "不吸烟",
	//		Car:            "未买车",
	//		House:          "",
	//		Drink:          "不喝酒",
	//		WantHaveChild:  "视情况而定",
	//		WhenMarital:    "时机成熟就结婚",
	//		HasChild:       "没有小孩",
	//
	//		UserObjInfo: model.UserObj{
	//			ObjAge:           "27-30岁",
	//			ObjHeight:        "160cm",
	//			ObjWeight:        "",
	//			ObjWorkPlace:     "四川阿坝茂县",
	//			ObjSalary:        "月薪:5千以上",
	//			ObjMaritalStatus: "未婚",
	//			ObjFigure:        "",
	//			ObjHasChild:      "没有小孩",
	//			ObjWantHaveChild: "想要孩子",
	//		},
	//	},
	//}

}
