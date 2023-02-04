package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//統計字符串的長度，按字節
	//golang 編碼統一為utf-8(ascii的字符(數字和字母)佔一個字符，中文字佔3的字節)
	str := "hello嗨"
	fmt.Println("str len=", len(str))

	//字符串遍歷，同時處理有中文的問題
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c \t", r[i])
	}
	fmt.Println()

	//字符串轉整數
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("轉換錯誤", err)
	} else {
		fmt.Println("轉換的結果是：", n)
	}

	//整數轉字符串
	str = strconv.Itoa(66)
	fmt.Printf("str=%v, str type=%T \n", str, str)

	//字符串轉`byte[]`
	var bytes = []byte("hello go")
	fmt.Println("bytes=", bytes)

	//`byte[]`轉字符串
	str = string([]byte{97, 98, 99})
	fmt.Println("str=", str)

	//10進制轉2, 8, 16進制
	str = strconv.FormatInt(123, 2)
	fmt.Println("123對應的二進制是：", str)

	//查找子串是否在指定的字符串中
	b := strings.Contains("seafood", "foo")
	fmt.Println("b=", b)

	//統計一個字符串有幾個指定的子串
	count := strings.Count("ceheese", "e")
	fmt.Println("count=", count)

	//不區分大小寫的字符串比較
	b = strings.EqualFold("abc", "ABC")
	fmt.Println("b=", b)

	//返回子串在字符串第一次出現的index值，如果沒有，返回`-1`
	index := strings.Index("NLT_abc", "abc")
	fmt.Println("index=", index)

	//返回子串在字符串最後一次出現的index值，如果沒有，返回`-1`
	index = strings.LastIndex("go golang", "go")
	fmt.Println("index=", index)

	//將指定的子串替換成另一子串，可以指定替換幾個，如果`n=-1`表示全部替換
	str = strings.Replace("go go hello", "go", "golang", 1)
	fmt.Println("str=", str)

	//按照指定的某個字符為分割標識，將一個字符串拆分成字符串數組
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Println("strArr=", strArr)

	//將字符串的字母進行大小寫的轉換
	str = "goLang Hello"
	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	fmt.Println("str=", str)
	fmt.Println("lower=", lower)
	fmt.Println("upper=", upper)

	//將字符串左右兩邊的空格去掉
	str = strings.TrimSpace("    hello      ")
	fmt.Printf("str=%s \n", str)

	//將字符串左右兩邊指定的字符去掉
	str = strings.Trim("! hello!", " !") //講左右兩邊的" "和!都去掉
	fmt.Printf("str=%s \n", str)

	//將字符串左邊指定的字符去掉
	str = strings.TrimLeft("! hello!", " !")
	fmt.Printf("str=%s \n", str)

	//將字符串右邊指定的字符去掉
	str = strings.TrimRight("! hello!", " !")
	fmt.Printf("str=%s \n", str)

	//判斷字符串是否已指定的字符串開頭
	b = strings.HasPrefix("ftp://192.168.0.1", "ftp")
	fmt.Println("b=", b)

	//判斷字符串是否已指定的字符串結束
	b = strings.HasSuffix("abc.jpg", "abc")
	fmt.Println("b=", b)
}
