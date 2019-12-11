package zhenai

import (
	"crawler/engine"
	"log"
	"net/url"
	"regexp"
	"strconv"
)

func CityParser(contents []byte) engine.RequestResult {
	var requestResult engine.RequestResult

	//获取城市页的顶部的人的信息
	heads := parserHead(contents)
	requestResult.Requests = append(requestResult.Requests, heads.Requests...)
	requestResult.Items = append(requestResult.Items, heads.Items...)

	// 获取列表里面的人的信息,其中涉及到翻页问题
	userList := parserList(contents)
	requestResult.Requests = append(requestResult.Requests, userList.Requests...)
	requestResult.Items = append(requestResult.Items, userList.Items...)

	return requestResult
}

// 处理城市页面的head部分
var headRe = regexp.MustCompile(`<li class="f-photo-li"><a href="http://www.zhenai.com/n/registerInfo\?channelId=[0-9]+&amp;fromurl=(http%3A%2F%2Falbum.zhenai.com%2Fu%2F[0-9]+)"[^<]*><img src="(http://images.zastatic.com/app/seo/randomuser/[0-9\\._a-z]+)" alt class="u-avatar"></a> <div class="c-info"><span class="nickname">([^<]+)</span> <span class="age">([0-9]+)岁</span></div> <div class="introduce">([^<]+)</div></li>`)

func parserHead(contents []byte) engine.RequestResult {
	var requesrResult engine.RequestResult
	heads := headRe.FindAllSubmatch(contents, -1)
	for _, val := range heads {
		//decode url
		path, err := url.PathUnescape(string(val[1]))
		if err != nil {
			log.Printf("PathUnescape err; path:%s", path)
			continue
		}
		requesrResult.Requests = append(requesrResult.Requests, engine.Request{
			Url:     path,
			Handler: UserInfoParser,
		})
		age, err := strconv.Atoi(string(val[4]))

		requesrResult.Items = append(requesrResult.Items, struct {
			avatar    string
			name      string
			age       int
			introduce string
		}{
			avatar:    string(val[2]),
			name:      string(val[3]),
			age:       age,
			introduce: string(val[5]),
		})
	}
	return requesrResult
}

// 处理城市页面的用户列表部分，获取出用户列表里面的每个用户的姓名和请求地址
var userListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func parserList(contents []byte) engine.RequestResult {
	var requesrResult engine.RequestResult
	submatchs := userListRe.FindAllSubmatch(contents, -1)
	for _, submatch := range submatchs {
		requesrResult.Requests = append(requesrResult.Requests, engine.Request{
			Url:     string(submatch[1]),
			Handler: UserInfoParser,
		})

		requesrResult.Items = append(requesrResult.Items, string(submatch[2]))
	}

	return requesrResult
}
