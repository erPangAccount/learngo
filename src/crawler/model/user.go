package model

type User struct {
	Avatar         string `json:"avatar"`          // 头像
	NickName       string `json:"nickname"`        //昵称
	Sex            string `json:"sex"`             //性别
	InnerMonologue string `json:"inner_monologue"` //内心独白
	MaritalStatus  string `json:"marital_status"`  //婚姻状况
	Age            int    `json:"age"`             //年龄
	Figure         string `json:"figure"`          //身材
	Constellation  string `json:"constellation"`   //星座
	Height         int    `json:"height"`          //身高
	Weight         int    `json:"weight"`          //体重
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

	UserObjInfo UserObj `json:"user_obj_info"`
}

type UserObj struct {
	ObjAge           string `json:"obj_age"`             //年龄
	ObjHeight        string `json:"obj_height"`          //身高
	ObjWeight        string `json:"obj_weight"`          //体重
	ObjWorkPlace     string `json:"obj_work_place"`      //工作地
	ObjSalary        string `json:"obj_salary"`          //薪资
	ObjMaritalStatus string `json:"obj_marital_status"`  //婚姻状况
	ObjFigure        string `json:"obj_figure"`          //身材
	ObjHasChild      string `json:"obj_has_child"`       //有孩子没有
	ObjWantHaveChild string `json:"obj_want_have_child"` //想要孩子吗
}
