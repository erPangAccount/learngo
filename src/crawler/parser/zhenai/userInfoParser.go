package zhenai

import (
	"crawler/engine"
	"crawler/model"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var userAvatar = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url\((https://photo.zastatic\.com/images/photo/[/0-9a-z.]+)[^>]+></div>`)
var baseInfoRe = regexp.MustCompile(`<div class="des f-cl"[^>]+>([^<]+)</div>`)

//var userPhotoRe = regexp.MustCompile(`<img src="(https://photo.zastatic.com/images/photo/[/0-9a-z.]+)[^"]+"[^>]*>`)
var innerMonologueRe = regexp.MustCompile(`<div class="m-title" data-v-8b1eac0c>内心独白</div> <div class="m-content-box m-des" data-v-8b1eac0c><span data-v-8b1eac0c>([^<]+)</span></div>`)
var ageRe = regexp.MustCompile(`([\d]+)岁`)
var heightRe = regexp.MustCompile(`([\d]+)cm`)

//以下正则匹配不出来是因为这个下面的数据格式不一致，每个人的数据都可能不同
var someInfoRe = regexp.MustCompile(`<div class="m-btn (purple|pink)"[^>]+>([^<]+)</div>`)
var constellationRe = regexp.MustCompile(`(.*座.*)`) //星座
var weightRe = regexp.MustCompile(`([\d]+)kg`)      //体重
var nationRe = regexp.MustCompile(`(.+族)`)          //民族
var hometownRe = regexp.MustCompile(`籍贯:(.+)`)      //籍贯
var smokingRe = regexp.MustCompile(`(.*烟.*)`)       //吸烟否
var drinkRe = regexp.MustCompile(`(.*酒.*)`)         //喝酒否
var carRe = regexp.MustCompile(`(.*车.*)`)           //车
var houseRe = regexp.MustCompile(`(.*房.*)`)         //房
var hasChildRe = regexp.MustCompile(`(.*有.*孩.*)`)   //是否有孩子
var wantChildRe = regexp.MustCompile(`是否想要孩子:(.+)`) //是否想要孩子
var whenMaritalRe = regexp.MustCompile(`何时结婚:(.+)`) //何时结婚
var figureRe = regexp.MustCompile(`体型:(.+)`)        //身材

// 择偶标准
var objSomeInfoRe = regexp.MustCompile(`<div class="m-btn"[^>]+>([^<]+)</div>`)
var objAgeRe = regexp.MustCompile(`([\d\-]+岁)`)            //年龄
var objHeightRe = regexp.MustCompile(`([\d\-]+cm)`)        //身高
var objWeightRe = regexp.MustCompile(`([\d\-]+kg)`)        //体重
var objWorkPlaceRe = regexp.MustCompile(`工作地:(.+)`)        //工作地
var objSalaryRe = regexp.MustCompile(`(.*薪.*)`)            //薪资
var objMaritalStatusRe = regexp.MustCompile(`(已婚|未婚|离异)`)  //婚姻状况
var objFigureRe = regexp.MustCompile(`体型:(.+)`)            //身材
var objHasChildRe = regexp.MustCompile(`(.*有.*孩.*)`)       //有孩子没有
var objWantHaveChildRe = regexp.MustCompile(`是否想要孩子:(.+)`) //想要孩子吗

func UserInfoParser(contents []byte, userNickName string) engine.RequestResult {
	var requestResult engine.RequestResult
	userInfo := model.User{}
	//用户头像
	userInfo.Avatar = extractString(contents, userAvatar)

	//用户名
	userInfo.NickName = userNickName

	//用户基本信息
	baseInfo := extractString(contents, baseInfoRe)
	baseInfos := strings.Split(baseInfo, " | ")
	if len(baseInfos) > 1 {
		//年龄
		ageInfo := extractString([]byte(baseInfo), ageRe)
		if age, err := strconv.Atoi(ageInfo); err == nil {
			userInfo.Age = age
		}
		//身高
		heightInfo := extractString([]byte(baseInfo), heightRe)
		if height, err := strconv.Atoi(heightInfo); err == nil {
			userInfo.Height = height
		}

		//工作地
		userInfo.WorkPlace = baseInfos[0]

		//学历
		userInfo.Education = baseInfos[2]

		//婚姻状态
		userInfo.MaritalStatus = baseInfos[3]

		//月收入
		userInfo.MonthlyIncome = baseInfos[5]
	}

	//内心独白
	userInfo.InnerMonologue = extractString(contents, innerMonologueRe)

	//匹配用户其他信息,由于统一匹配不出来，只能先按照每个去匹配，然后分解成一个字符串再次进行匹配。

	//个人信息栏
	someInfos_1 := someInfoRe.FindAllSubmatch(contents, -1)
	if len(someInfos_1) > 2 {
		for _, val := range someInfos_1 {
			if len(val) > 2 {
				//星座
				constellation := extractString(val[2], constellationRe)
				if len(constellation) > 0 {
					userInfo.Constellation = constellation
				}

				//体重
				weightInfo := extractString(val[2], weightRe)
				if weight, err := strconv.Atoi(weightInfo); err == nil {
					userInfo.Weight = weight
				}

				//民族
				nationInfo := extractString(val[2], nationRe)
				if len(nationInfo) > 0 {
					userInfo.Nation = nationInfo
				}

				//籍贯
				hometownInfo := extractString(val[2], hometownRe)
				if len(hometownInfo) > 0 {
					userInfo.Hometown = hometownInfo
				}

				//吸烟否
				smokingInfo := extractString(val[2], smokingRe)
				if len(smokingInfo) > 0 {
					userInfo.Smoking = smokingInfo
				}

				//喝酒否
				drinknfo := extractString(val[2], drinkRe)
				if len(drinknfo) > 0 {
					userInfo.Drink = drinknfo
				}

				//购车否
				carInfo := extractString(val[2], carRe)
				if len(carInfo) > 0 {
					userInfo.Car = carInfo
				}

				//购房否
				houseInfo := extractString(val[2], houseRe)
				if len(houseInfo) > 0 {
					userInfo.House = houseInfo
				}

				//有孩子吗
				hasChildInfo := extractString(val[2], hasChildRe)
				if len(hasChildInfo) > 0 {
					userInfo.HasChild = hasChildInfo
				}

				//想要孩子吗
				wantChildInfo := extractString(val[2], wantChildRe)
				if len(wantChildInfo) > 0 {
					userInfo.WantHaveChild = wantChildInfo
				}

				//合适结婚
				whenMaritalInfo := extractString(val[2], whenMaritalRe)
				if len(whenMaritalInfo) > 0 {
					userInfo.WhenMarital = whenMaritalInfo
				}

				//合适结婚
				figureInfo := extractString(val[2], figureRe)
				if len(figureInfo) > 0 {
					userInfo.Figure = figureInfo
				}
			}
		}
	}

	//择偶标准
	someInfos_2 := objSomeInfoRe.FindAllSubmatch(contents, -1)
	if len(someInfos_2) > 2 {
		for _, val := range someInfos_2 {
			log.Printf("%s", val)
			if len(val) > 1 {
				//年龄
				objAgeInfo := extractString(val[1], objAgeRe)
				if len(objAgeInfo) > 0 {
					userInfo.UserObjInfo.ObjAge = objAgeInfo
				}

				//体重
				objWeightInfo := extractString(val[1], objWeightRe)
				if len(objWeightInfo) > 0 {
					userInfo.UserObjInfo.ObjWeight = objWeightInfo
				}

				//身高
				objHeightInfo := extractString(val[1], objHeightRe)
				if len(objHeightInfo) > 0 {
					userInfo.UserObjInfo.ObjHeight = objHeightInfo
				}

				//工作地
				objWorkPlaceInfo := extractString(val[1], objWorkPlaceRe)
				if len(objWorkPlaceInfo) > 0 {
					userInfo.UserObjInfo.ObjWorkPlace = objWorkPlaceInfo
				}

				//薪资
				objSalaryInfo := extractString(val[1], objSalaryRe)
				if len(objSalaryInfo) > 0 {
					userInfo.UserObjInfo.ObjSalary = objSalaryInfo
				}

				//婚姻状况
				objMaritalStatusInfo := extractString(val[1], objMaritalStatusRe)
				if len(objMaritalStatusInfo) > 0 {
					userInfo.UserObjInfo.ObjMaritalStatus = objMaritalStatusInfo
				}

				//体型
				objFigureInfo := extractString(val[1], objFigureRe)
				if len(objFigureInfo) > 0 {
					userInfo.UserObjInfo.ObjFigure = objFigureInfo
				}

				//有孩子吗
				objHasChildInfo := extractString(val[1], objHasChildRe)
				if len(objHasChildInfo) > 0 {
					userInfo.UserObjInfo.ObjHasChild = objHasChildInfo
				}

				//想要孩子吗？
				objWantHaveChildInfo := extractString(val[1], objWantHaveChildRe)
				if len(objWantHaveChildInfo) > 0 {
					userInfo.UserObjInfo.ObjWantHaveChild = objWantHaveChildInfo
				}
			}
		}
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
