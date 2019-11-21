package main

import "fmt"

/**
 思想：
	对于每个字符X
		查找最后出现的位置，当最后出现的位置不存在或者小于当前子串开始位置，不做任何操作
		否则 更新start = 最后出现的位置 + 1
		更新子串最大长度和X的最后出现位置
		返回子串最大长度
*/
func lengthOfNonRepeatingSubStr(s string) int {
	last := make(map[rune]int)
	subStart, subLength := 0, 0
	for i, ch := range []rune(s) {
		if value, ok := last[ch]; ok && value >= subStart {
			subStart = value + 1
		}

		if (i - subStart + 1) > subLength {
			subLength = i - subStart + 1
		}
		last[ch] = i
	}
	return subLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcadd"))     //4
	fmt.Println(lengthOfNonRepeatingSubStr("bbb"))           //1
	fmt.Println(lengthOfNonRepeatingSubStr("waeasfasasd"))   //4
	fmt.Println(lengthOfNonRepeatingSubStr("你好golang"))      //7
	fmt.Println(lengthOfNonRepeatingSubStr("阿斯蒂芬阿斯顿法国红酒看来")) //11
}
