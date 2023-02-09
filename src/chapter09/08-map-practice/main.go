package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	users[name]["pwd"] = "888888"
}

func main() {
	//*使用`map[string]map[string]string`的`map`類型。*
	//*key: 表示用戶名，是唯一的，不可以重複。*
	//*如果某個用戶名存在，就將其密碼修改成"888888"，如果不存在就增加這個用戶信息(包括暱稱nickname和密碼pwd)。*
	//*編寫一個函數`modifyUser(users map[string]map[string]string, name string)`完成上述功能。*

	users := make(map[string]map[string]string)

	var name string
	var nickname string
	var pwd string
	for {
		fmt.Print("請輸入用戶名：")
		fmt.Scanln(&name)
		fmt.Print("請輸入暱稱：")
		fmt.Scanln(&nickname)
		fmt.Print("請輸入密碼：")
		fmt.Scanln(&pwd)

		if users[name] != nil {
			modifyUser(users, name)
		} else {
			user := make(map[string]string, 2)
			user["nickname"] = nickname
			user["pwd"] = pwd

			users[name] = user
		}

		fmt.Println("所有用戶為：", users)
	}
}
