package model

type User struct {
	Avatar         string `json:"avatar"`          // 头像
	NickName       string `json:"nickname"`        //昵称
	InnerMonologue string `json:"inner_monologue"` //内心独白
	MaritalStatus  string `json:"marital_status"`  //婚姻状况
	Age            int    `json:"age"`             //年龄
	Constellation  string `json:"constellation"`   //星座
	Height         int    `json:"height"`          //身高
	WorkPlace      string `json:"work_place"`      //工作地
	MonthlyIncome  string `json:"monthly_income"`  //月收入
	Job            string `json:"job"`             //职业
	Education      string `json:"education"`       //学历
	Nation         string `json:"nation"`          //民族
	Hometown       string `json:"hometown"`        //籍贯
	Smoking        string `json:"smoking"`         //吸烟否
	Drink          string `json:"drink"`           //喝酒否
	Car            string `json:"car"`             //购车否
	House          string `json:"house"`           //购房否
	HasChild       string `json:"has_child"`       //有孩子没有
	WantHaveChild  string `json:"want_have_child"` //想要孩子吗
	WhenMarital    string `json:"when_marital"`    //何时结婚
}
