package monster

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	data, err1 := json.Marshal(m)
	if err1 != nil {
		fmt.Println("序列化monster錯誤，err=", err1)
		return false
	}

	path := "/Users/jason/Desktop/monster.txt"
	file, err2 := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

	defer file.Close()

	if err2 != nil {
		fmt.Println("寫入文件錯誤，err=", err2)
		return false
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()

	return true
}

func (m *Monster) ReStore() bool {
	path := "/Users/jason/Desktop/monster.txt"
	file, err1 := os.OpenFile(path, os.O_RDONLY, 0666)

	defer file.Close()

	if err1 != nil {
		fmt.Println("讀取文件錯誤，err=", err1)
		return false
	}

	reader := bufio.NewReader(file)
	str, _ := reader.ReadString('\n')

	err2 := json.Unmarshal([]byte(str), m)
	if err2 != nil {
		fmt.Println("反序列化monster錯誤，err=", err2)
		return false
	}

	return true
}
