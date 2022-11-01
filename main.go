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
	Lottery := Lottery{num}
	Lottery.SetMaxAmountOfBalls()
	/*check := 0
	/*for {
		if check != 0{
			break
		}
		fmt.Scanf("%d", &check)
		Lottery.Shuffle()
		Lottery.ShowmaxAmountOfBallsSlice()
	}*/

	
	for {
		//fmt.Println("請問要抽獎嗎（抽請輸入：0，不抽請輸入：1)")
		result,status := Lottery.GetOneNumber()
		fmt.Println("抽到",result)
		fmt.Print("抽獎機內目前剩餘")
		Lottery.ShowmaxAmountOfBallsSlice()
		if status != ""{
			fmt.Println("抽完囉")
			break
		}
	}

}