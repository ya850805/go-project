package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1. 先創建一個原始陣列
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //藍子

	//2. 輸出看看原始的陣列
	for _, row := range chessMap {
		for _, v := range row {
			fmt.Printf("%d \t", v)
		}
		fmt.Println()
	}

	//3. 轉成稀疏陣列
	//(1). 遍歷chessMap，如果我們發現有一個元素值不為0，創建一個node結構體
	//(2). 將node放到切片中即可
	var sparseArr []ValNode

	//標準的一個稀疏陣列應該還有一個紀錄元素的二維陣列的規模(行、列、默認值)
	sparseArr = append(sparseArr, ValNode{
		row: 11,
		col: 11,
		val: 0,
	})

	for i, row := range chessMap {
		for j, v := range row {
			if v != 0 {
				node := ValNode{
					row: i,
					col: j,
					val: v,
				}
				sparseArr = append(sparseArr, node)
			}
		}
	}

	//輸出並保存稀疏陣列
	filePath := "/Users/jason/Desktop/chess.data"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Open file error=", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Println("當前的稀疏陣列是::::")
	for i, valNode := range sparseArr {
		writer.WriteString(fmt.Sprintf("%d %d %d\n", valNode.row, valNode.col, valNode.val))
		fmt.Printf("%d: %d %d %d \n", i, valNode.row, valNode.col, valNode.val)
	}

	writer.Flush()

	//////////////////////////////////////////////////////////////////////////////////////

	//如何恢復原始陣列
	file, err = os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Open file error=", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	//跳過第一行
	_, err = reader.ReadString('\n')

	if err != nil {
		fmt.Println("read first line error=", err)
	}

	//firstRow = strings.Trim(firstRow, "\n")
	//firstRowArr := strings.Split(firstRow, " ")
	//totalRow, _ := strconv.Atoi(firstRowArr[0])
	//totalCol, _ := strconv.Atoi(firstRowArr[1])

	var chessMap2 [11][11]int
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.Trim(str, "\n")
		arr := strings.Split(str, " ")
		row, _ := strconv.Atoi(arr[0])
		col, _ := strconv.Atoi(arr[1])
		val, _ := strconv.Atoi(arr[2])
		chessMap2[row][col] = val
	}

	//打印chessMap2，查看是不是恢復
	fmt.Println("恢復後稀疏陣列是::::")
	for _, rows := range chessMap2 {
		for _, val := range rows {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}
