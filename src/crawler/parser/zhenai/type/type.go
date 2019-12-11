package zhenai_type

type User struct {
	Avatar         string `json:"avatar"`          // 头像
	NickName       string `json:"nickname"`        //昵称
	InnerMonologue string `json:"inner_monologue"` //内心独白
	MaritalStatus  string `json:"marital_status"`  //婚姻状况

}
