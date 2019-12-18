package zhenai_test

import (
	"crawler/model"
	"crawler/parser/zhenai"
	"io/ioutil"
	"testing"
)

func TestUserInfoParser(t *testing.T) {
	contents, err := ioutil.ReadFile("userInfoParserTemplate.html")
	if err != nil {
		panic(err)
	}

	result := zhenai.UserInfoParser(contents, "")

	user := model.User{
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
	}

	// verify result
	resultUser := result.Items[0].(model.User)
	if resultUser.Avatar != user.Avatar {
		t.Errorf("expected %v, but got: %v \n", user.Avatar, resultUser.Avatar)
	}
	if resultUser.NickName != user.NickName {
		t.Errorf("expected %s, but got: %s \n", user.NickName, resultUser.NickName)
	}
	if resultUser.Sex != user.Sex {
		t.Errorf("expected %v, but got: %v \n", user.Sex, resultUser.Sex)
	}
	if resultUser.InnerMonologue != user.InnerMonologue {
		t.Errorf("expected %v, but got: %v \n", user.InnerMonologue, resultUser.InnerMonologue)
	}
	if resultUser.MaritalStatus != user.MaritalStatus {
		t.Errorf("expected %v, but got: %v \n", user.MaritalStatus, resultUser.MaritalStatus)
	}
	if resultUser.Age != user.Age {
		t.Errorf("expected %v, but got: %v \n", user.Age, resultUser.Age)
	}
	if resultUser.Figure != user.Figure {
		t.Errorf("expected %v, but got: %v \n", user.Figure, resultUser.Figure)
	}
	if resultUser.Constellation != user.Constellation {
		t.Errorf("expected %s, but got: %s \n", user.Constellation, resultUser.Constellation)
	}
	if resultUser.Height != user.Height {
		t.Errorf("expected %v, but got: %v \n", user.Height, resultUser.Height)
	}
	if resultUser.Weight != user.Weight {
		t.Errorf("expected %v, but got: %v \n", user.Weight, resultUser.Weight)
	}
	if resultUser.WorkPlace != user.WorkPlace {
		t.Errorf("expected %v, but got: %v \n", user.WorkPlace, resultUser.WorkPlace)
	}
	if resultUser.MonthlyIncome != user.MonthlyIncome {
		t.Errorf("expected %v, but got: %v \n", user.MonthlyIncome, resultUser.MonthlyIncome)
	}
	if resultUser.Job != user.Job {
		t.Errorf("expected %v, but got: %v \n", user.Job, resultUser.Job)
	}
	if resultUser.Education != user.Education {
		t.Errorf("expected %s, but got: %s \n", user.Education, resultUser.Education)
	}
	if resultUser.Nation != user.Nation {
		t.Errorf("expected %v, but got: %v \n", user.Nation, resultUser.Nation)
	}
	if resultUser.Hometown != user.Hometown {
		t.Errorf("expected %v, but got: %v \n", user.Hometown, resultUser.Hometown)
	}
	if resultUser.Smoking != user.Smoking {
		t.Errorf("expected %v, but got: %v \n", user.Smoking, resultUser.Smoking)
	}
	if resultUser.Drink != user.Drink {
		t.Errorf("expected %v, but got: %v \n", user.Drink, resultUser.Drink)
	}
	if resultUser.Car != user.Car {
		t.Errorf("expected %v, but got: %v \n", user.Car, resultUser.Car)
	}
	if resultUser.House != user.House {
		t.Errorf("expected %s, but got: %s \n", user.House, resultUser.House)
	}
	if resultUser.HasChild != user.HasChild {
		t.Errorf("expected %v, but got: %v \n", user.HasChild, resultUser.HasChild)
	}
	if resultUser.WantHaveChild != user.WantHaveChild {
		t.Errorf("expected %v, but got: %v \n", user.WantHaveChild, resultUser.WantHaveChild)
	}
	if resultUser.WhenMarital != user.WhenMarital {
		t.Errorf("expected %v, but got: %v \n", user.WhenMarital, resultUser.WhenMarital)
	}
	if resultUser.UserObjInfo.ObjAge != user.UserObjInfo.ObjAge {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjAge, resultUser.UserObjInfo.ObjAge)
	}
	if resultUser.UserObjInfo.ObjHeight != user.UserObjInfo.ObjHeight {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjHeight, resultUser.UserObjInfo.ObjHeight)
	}
	if resultUser.UserObjInfo.ObjWeight != user.UserObjInfo.ObjWeight {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjWeight, resultUser.UserObjInfo.ObjWeight)
	}
	if resultUser.UserObjInfo.ObjWorkPlace != user.UserObjInfo.ObjWorkPlace {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjWorkPlace, resultUser.UserObjInfo.ObjWorkPlace)
	}
	if resultUser.UserObjInfo.ObjSalary != user.UserObjInfo.ObjSalary {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjSalary, resultUser.UserObjInfo.ObjSalary)
	}
	if resultUser.UserObjInfo.ObjMaritalStatus != user.UserObjInfo.ObjMaritalStatus {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjMaritalStatus, resultUser.UserObjInfo.ObjMaritalStatus)
	}
	if resultUser.UserObjInfo.ObjFigure != user.UserObjInfo.ObjFigure {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjFigure, resultUser.UserObjInfo.ObjFigure)
	}

	if resultUser.UserObjInfo.ObjHasChild != user.UserObjInfo.ObjHasChild {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjHasChild, resultUser.UserObjInfo.ObjHasChild)
	}
	if resultUser.UserObjInfo.ObjWantHaveChild != user.UserObjInfo.ObjWantHaveChild {
		t.Errorf("expected %v, but got: %v \n", user.UserObjInfo.ObjWantHaveChild, resultUser.UserObjInfo.ObjWantHaveChild)
	}
}
