package zhenai

import (
	"crawler/engine"
	"crawler/parser/zhenai/type"
	"log"
	"regexp"
)

var userAvatar = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url\((https://photo.zastatic\.com/images/photo/[/0-9a-z.]+)[^>]+></div>`)

//var userPhotoRe = regexp.MustCompile(`<img src="(https://photo.zastatic.com/images/photo/[/0-9a-z.]+)[^"]+"[^>]*>`)
var innerMonologueRe = regexp.MustCompile(`<div class="m-title" data-v-8b1eac0c>内心独白</div> <div class="m-content-box m-des" data-v-8b1eac0c><span data-v-8b1eac0c>([^<]+)</span></div>`)
var someInfoRe = regexp.MustCompile(`<div class="m-title" data-v-8b1eac0c>个人资料</div> <div class="m-content-box" data-v-8b1eac0c><div class="purple-btns" data-v-8b1eac0c><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div></div> <div class="pink-btns" data-v-8b1eac0c><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div><div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div></div></div>`)

func UserInfoParser(contents []byte) engine.RequestResult {
	var requestResult engine.RequestResult
	userInfo := zhenai_type.User{}
	//用户头像
	avatars := userAvatar.FindSubmatch(contents)
	userInfo.Avatar = string(avatars[1])

	//用户相册 网页上看着有，但是查看网页源代码就没了，爬下来的也没有
	//photos := userPhotoRe.FindAllSubmatch(contents, -1)
	//log.Println(photos)
	//for _, photo := range photos {
	//	log.Printf("%s  ", photo[1])
	//}

	//内心独白
	innerMonologue := innerMonologueRe.FindSubmatch(contents)
	userInfo.InnerMonologue = string(innerMonologue[1])

	//个人信息
	someInfos := someInfoRe.FindSubmatch(contents)
	for _, val := range someInfos {
		log.Printf("%s; ", val)
	}

	return requestResult
}
