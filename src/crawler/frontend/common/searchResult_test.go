package common

import (
	"crawler/engine"
	"crawler/frontend/model"
	dataModel "crawler/model"
	"os"
	"testing"
)

func TestSearchResult(t *testing.T) {
	htmlTemplate := CreateSearchResultTemplateView("../view/searchResult.html")

	out, err := os.Create("searchResultTemplate_test.html")
	if err != nil {
		panic(err)
	}
	searchResult := model.SearchResult{
		Keyword: "",
		PageInfo: struct {
			Total      int
			PrevPage   int
			TargetPage int
			NextPage   int
			PageSize   int
		}{
			100,
			0,
			1,
			2,
			10,
		},
		Items: nil,
	}
	itemInfo := engine.Item{
		Url:  "https://album.zhenai.com/u/1662184411",
		Type: "",
		Id:   "1662184411",
		DoType: dataModel.User{
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

			UserObjInfo: dataModel.UserObj{
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
	for i := 0; i < 10; i++ {
		searchResult.Items = append(searchResult.Items, itemInfo)
	}

	err = htmlTemplate.Render(out, searchResult)
	if err != nil {
		panic(err)
	}
}
