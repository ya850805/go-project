package main

import "fmt"

// golang字符串類型使用
func main() {
	var address = "台北市大安區"
	fmt.Println("address=", address)

	//報錯：Cannot assign to str[0]，字符串一但賦值了就不能修改了
	//var str = "hello"
	//str[0] = 'a'

	str1 := "abc\nabc"
	fmt.Println(str1)

	//**反引號(\`)，以字符串的原生形式輸出，包括換行和特殊字符，可以實現防止攻擊、輸出源代碼等效果**
	str2 := `
		hello\n\t\\
			123
		456
	`
	fmt.Println(str2)

	//字符串拼接
	str3 := "hello" + " world"
	fmt.Println("str3=", str3)
	str3 += " haha"
	fmt.Println("str3=", str3)

	//當一個拼接的操作很長時，可以分行寫，需要將+號保留在上一行
	str4 := "hello" + " world" + "hello" + " world" + "hello" + " world" +
		"hello" + " world" +
		"hello" + " world"
	fmt.Println("str4=", str4)
}
