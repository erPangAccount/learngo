package zhenai

import (
	"crawler/engine"
	"fmt"
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

	return requestResult
}

// 处理城市页面的head部分
var headRe = regexp.MustCompile(`<li class="f-photo-li"><a href="http://www.zhenai.com/n/registerInfo\?channelId=[0-9]+&amp;fromurl=(http%3A%2F%2Falbum.zhenai.com%2Fu%2F[0-9]+)"[^<]*><img src="(http://images.zastatic.com/app/seo/randomuser/[0-9\\._a-z]+)" alt class="u-avatar"></a> <div class="c-info"><span class="nickname">([^<]+)</span> <span class="age">([0-9]+)岁</span></div> <div class="introduce">([^<]+)</div></li>`)

func parserHead(contents []byte) engine.RequestResult {
	var requesrResult engine.RequestResult
	heads := headRe.FindAllSubmatch(contents, -1)

	fmt.Println("---------------")
	for _, val := range heads {
		//decode url
		path, err := url.PathUnescape(string(val[1]))
		if err != nil {
			log.Printf("PathUnescape err; path:%s", path)
			continue
		}
		requesrResult.Requests = append(requesrResult.Requests, engine.Request{
			Url:     path,
			Handler: engine.NilRequestResultFunc,
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
		fmt.Printf("%s %s;%s;%s;%s\n", path, val[2], val[3], val[4], val[5])
	}
	return requesrResult
}
