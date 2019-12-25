package persist

import (
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
	"testing"
)

func TestSave(t *testing.T) {
	user := engine.Item{
		Url:  "https://album.zhenai.com/u/1662184411",
		Type: "",
		Id:   "1662184411",
		DoType: model.User{
			Avatar:         "https://photo.zastatic.com/images/photo/415547/1662184411/10203260896941091.jpg",
			NickName:       "",
			Sex:            "",
			InnerMonologue: "将一段记忆在风雨中渐渐淡忘，在时光的转角处抹去这段忧伤，让岁月不再诉说心酸的过往，让风月从此不再黯然神伤。",
			MaritalStatus:  "未婚",
			Age:            27,
			Figure:         "苗条",
			Constellation:  "天蝎座(10.23-11.21)",
			Height:         160,
			Weight:         41,
			WorkPlace:      "阿坝",
			MonthlyIncome:  "5001-8000元",
			Job:            "",
			Education:      "大学本科",
			Nation:         "羌族",
			Hometown:       "四川阿坝",
			Smoking:        "不吸烟",
			Car:            "未买车",
			House:          "",
			Drink:          "不喝酒",
			WantHaveChild:  "视情况而定",
			WhenMarital:    "时机成熟就结婚",
			HasChild:       "没有小孩",

			UserObjInfo: model.UserObj{
				ObjAge:           "27-30岁",
				ObjHeight:        "160cm",
				ObjWeight:        "",
				ObjWorkPlace:     "四川阿坝茂县",
				ObjSalary:        "月薪:5千以上",
				ObjMaritalStatus: "未婚",
				ObjFigure:        "",
				ObjHasChild:      "没有小孩",
				ObjWantHaveChild: "想要孩子",
			},
		},
	}

	id, err := save(user)
	if err != nil && id != "exists" {
		panic(err)
	}

	// TODO 开启elasticsearch后再试
	client, err := elastic.NewClient(
		elastic.SetURL(engine.ElasticHost),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	if user.Type == "" {
		user.Type = "index"
	}

	result, err := client.Get().Index("test").Type(user.Type).Id(user.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var getData engine.Item
	err = json.Unmarshal(result.Source, &getData)
	if err != nil {
		panic(err)
	}
	getData.DoType, err = model.FromJsonObj(getData.DoType)
	if err != nil {
		panic(err)
	}

	if getData != user {
		t.Errorf("Got %+v, expected %+v", getData, user)
	}
	t.Logf("Got %+v", getData)
}
