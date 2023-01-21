package main

import "fmt"

//其他進制轉十進制轉換規則

//二進制轉十進制
//規則：從最低位開始(最右邊)，將每個為上的數提取出來，乘以2的(位數-1)次方，然後求和。
//案例：請將二進制1011轉成十進制的數。
//:point_right:1011 = 1 * 20 + 1 * 21 + 0 * 2 2 + 1 * 23 = 1 + 2 + 0 + 8 = 11

//八進制轉十進制
//規則：從最低位開始(最右邊)，將每個為上的數提取出來，乘以8的(位數-1)次方，然後求和。
//案例：請將八進制0123轉成十進制的數。
//:point_right:0123 = 3 * 80 + 2 * 81 + 1 * 82 = 3 + 16 + 64 = 83

//十六進制轉十進制
//規則：從最低位開始(最右邊)，將每個為上的數提取出來，乘以16的(位數-1)次方，然後求和。
//案例：請將十六進制0x34A轉成十進制的數。
//:point_right:0x34A = 10 * 160 + 4 * 161 + 3 * 162 = 10 + 64 + 768 = 842

////////////////////////////////////////////////////////////////////////////

//十進制轉其他進制

//十進制轉二進制
//規則：將該數不斷除以2，直到商為0為止，然後將每步得到的餘數倒過來，就是對應的二進制。
//案例：將56轉成二進制。
//:point_right:111000

//十進制轉八進制
//規則：將該數不斷除以8，直到商為0為止，然後將每步得到的餘數倒過來，就是對應的八進制。
//案例：將156轉成八進制。
//:point_right:0234

//十進制轉十六進制
//規則：將該數不斷除以16，直到商為0為止，然後將每步得到的餘數倒過來，就是對應的十六進制。
//案例：將356轉成十六進制。
//:point_right:0x164

////////////////////////////////////////////////////////////////////////////

//二進制轉其他進制

//二進制轉八進制
//規則：將二進制數每3位一組(從低位開始組合)，轉成對應的十進制數然後拼接即可。
//案例：將二進制11010101轉成八進制。
//:point_right:11010101 = 0325

//二進制轉十六進制
//規則：將二進制數每4位一組(從低位開始組合)，轉成對應的十進制數然後拼接即可。
//案例：將二進制11010101轉成十六進制。
//:point_right:11010101 = 0xD5

////////////////////////////////////////////////////////////////////////////

//其他進制轉二進制

//八進制轉二進制
//規則：將八進制的每1位，轉成對應的一個3位的二進制數即可。
//案例：將0237轉成二進制。
//:point_right:0237 = 10011111

//十六進制轉二進制
//規則：將十六進制的每1位，轉成對應的一個4位的二進制數即可。
//案例：將0x237轉成二進制。
//:point_right:0x237 = 1000110111

// golang中的進制
func main() {
	var i int = 5
	//二進制輸出
	fmt.Printf("%b \n", i)

	//**八進制**：0~7，滿8進1，以數字0開頭
	var j int = 011
	fmt.Println("j=", j)

	//**十六進制**：0~9及A~F，滿16進1，以0x或0X開頭表示，此處的A~F不區分大小寫。
	var k int = 0x11
	fmt.Println("k=", k)
}
