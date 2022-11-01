package main

import (
	"fmt"
	"time"
)




func main() {
	var num int
	//var count int
	fmt.Println("請問最大的數字為？")
	fmt.Scanf("%d", &num)
	//fmt.Println("請問要抽出幾個？")
	//fmt.Scanf("%d", &count)
	time.Sleep(5000)
	SetMaxAmountOfBalls(num)
	check := 0
	for {
		fmt.Scanf("%d", &check)
		Shuffle(MaxAmountOfBallsSlice)
		fmt.Println(MaxAmountOfBallsSlice)
		if check != 0{
			break
		}	
	}
	Shuffle(MaxAmountOfBallsSlice)
	
	/*for {
		fmt.Println("請問抽獎嗎（抽請輸入：0)")
		fmt.Scanf("%d", &num)
		if num == 0{
			fmt.Println(len(MaxAmountOfBallsSlice))
			MaxAmountOfBallsSlice = GetOneNumber(MaxAmountOfBallsSlice)
			fmt.Println("剩餘",MaxAmountOfBallsSlice)
		}else{
			return
		}
	}*/

}