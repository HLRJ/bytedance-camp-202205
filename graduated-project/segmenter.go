package main

import (
	"fmt"
	"strings"
)

// 只为下列字符建立索引，返回false，如标点符号一类的，返回true
func isIgnoredChar(char rune) bool {
	if (char >= 0x4E00 && char <= 0x9FA5) || // 汉字
		(char >= 'A' && char <= 'Z') || // A-Z
		(char >= 'a' && char <= 'z') || // a-z
		(char >= '0' && char <= '9') { // 0-9
		return false
	}
	return true

}

//首先将文本过滤，例如将标点符号过滤
//输入string  返回string
func Text_filter(str string) string {
	var filter_text strings.Builder // 创建一个返回值类型，这种拼接字符串，效率更高
	chars := []rune(str)
	for _, ch := range chars {
		if isIgnoredChar(ch) {
			continue
		} else {
			// 将字符串拼接在一起
			filter_text.WriteString(string(ch))
		}
	}

	return filter_text.String()
}

// 将上面过滤后的字符进行取词
func Segment_word(str string, n int) []string {
	nWordSet := make([]string, 0, 32)
	// 如不转成rune，直接遍历string，一个汉字占两个索引，不方便取词
	chars := []rune(str)
	//每隔n个
	for i := 0; i < len(chars); {
		//每隔n个取切片加入string切片
		nWordSet = append(nWordSet, string(chars[i:i+n]))
		i = i + n
		if i+n >= len(chars) { //最后一组，可能不足n个元素，比如我们取三个元素，最后一组只有一个或者两个，直接取到末尾即可
			nWordSet = append(nWordSet, string(chars[i:]))
			break
		}

	}
	return nWordSet
}
func main() {
	s := "你好nigh,dfawdfaw,$$$$$"
	fmt.Println("========================原始字符串==========================")
	fmt.Println(s)
	fmt.Println("=========================过滤=========================")
	fmt.Println(Text_filter(s))
	fmt.Println("===================n元分词=====================")
	fmt.Println(Segment_word(Text_filter(s), 3))

}
