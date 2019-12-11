package zhenai

import (
	"crawler/engine"
	"crawler/model"
	"log"
	"regexp"
	"strconv"
)

var userAvatar = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url\((https://photo.zastatic\.com/images/photo/[/0-9a-z.]+)[^>]+></div>`)

//var userPhotoRe = regexp.MustCompile(`<img src="(https://photo.zastatic.com/images/photo/[/0-9a-z.]+)[^"]+"[^>]*>`)
var innerMonologueRe = regexp.MustCompile(`<div class="m-title" data-v-8b1eac0c>内心独白</div> <div class="m-content-box m-des" data-v-8b1eac0c><span data-v-8b1eac0c>([^<]+)</span></div>`)

//以下正则匹配不出来是因为这个下面的数据格式不一致，每个人的数据都可能不同
var someInfoRe = regexp.MustCompile(`<div class="m-title" data-v-8b1eac0c>个人资料</div> <div class="m-content-box" data-v-8b1eac0c><div class="purple-btns" data-v-8b1eac0c><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div><div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div><div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div></div> <div class="pink-btns" data-v-8b1eac0c><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>体型:([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>是否想要孩子:([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>何时结婚:([^<]+)</div></div></div>`)

func UserInfoParser(contents []byte, userNickName string) engine.RequestResult {
	log.Printf("%s", contents)
	var requestResult engine.RequestResult
	userInfo := model.User{}
	//用户头像
	userInfo.Avatar = extractString(contents, userAvatar)

	//用户相册 网页上看着有，但是查看网页源代码就没了，爬下来的也没有
	//photos := userPhotoRe.FindAllSubmatch(contents, -1)
	//log.Println(photos)
	//for _, photo := range photos {
	//	log.Printf("%s  ", photo[1])
	//}
	userInfo.NickName = "Name: " + userNickName

	//内心独白
	userInfo.InnerMonologue = extractString(contents, innerMonologueRe)

	//个人信息
	someInfos := s.FindSubmatch(contents)
	log.Println(someInfos)
	if len(someInfos) > 1 {
		age, err := strconv.Atoi(string(someInfos[2]))
		if err != nil {
			userInfo.Age = age
		}

		height, err := strconv.Atoi(string(someInfos[4]))
		if err != nil {
			userInfo.Height = height
		}

		userInfo.MaritalStatus = string(someInfos[1])
		userInfo.Constellation = string(someInfos[3])
		userInfo.WorkPlace = string(someInfos[5])
		userInfo.MonthlyIncome = string(someInfos[6])
		userInfo.Job = string(someInfos[7])
		userInfo.Education = string(someInfos[8])
		userInfo.Nation = string(someInfos[9])
		userInfo.Hometown = string(someInfos[10])
		userInfo.Smoking = string(someInfos[11])
		userInfo.Drink = string(someInfos[12])
		userInfo.Car = string(someInfos[13])
		userInfo.House = string(someInfos[14])
		userInfo.HasChild = string(someInfos[15])
		userInfo.WantHaveChild = string(someInfos[16])
		userInfo.WhenMarital = string(someInfos[17])
	}

	requestResult.Items = append(requestResult.Items, userInfo)
	log.Printf("%v", requestResult)
	return requestResult
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	}
	return ""
}
