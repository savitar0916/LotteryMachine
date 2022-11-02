package main

import (
	"fmt"
	"math/rand"
	"time"
)
// https://github.com/savitar0916/LotteryMachine
// 附上github連結可以看到修改的過程
// Lottery := new (Lottery)
// Lottery.Init(100000)
// Lottery.SetMaxAmountOfBalls()
// result,status := Lottery.GetOneNumber()

type Lottery struct{
	MaxAmountOfBalls int
}
//總共要幾顆球
func (l *Lottery) Init(Balls int) {
	l.MaxAmountOfBalls = Balls
}
// 封裝
var maxAmountOfBallsSlice []int
// 顯示目前的slice裡有哪些球
func (l *Lottery)ShowmaxAmountOfBallsSlice(){
	fmt.Println(maxAmountOfBallsSlice)
}
//依據球數丟入球
func (l *Lottery)SetMaxAmountOfBalls(){
	for i := 1; i <= l.MaxAmountOfBalls; i++ {
		maxAmountOfBallsSlice = append(maxAmountOfBallsSlice, i)
	}
}

//洗球返回洗完後的slice
func (l *Lottery)Shuffle()[]int{
	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

	for i := len(maxAmountOfBallsSlice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		maxAmountOfBallsSlice[i], maxAmountOfBallsSlice[j] = maxAmountOfBallsSlice[j], maxAmountOfBallsSlice[i]
	}
	return maxAmountOfBallsSlice
}
//抽球
func (l *Lottery) GetOneNumber()(int,string){
	if len(maxAmountOfBallsSlice) == 1 {
		return maxAmountOfBallsSlice[0],"no more any num"
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(maxAmountOfBallsSlice))
	result := maxAmountOfBallsSlice[num]
	//直接用append的方式去將其賦予新的值 把中間的result給刪除
	maxAmountOfBallsSlice = append(maxAmountOfBallsSlice[:num],maxAmountOfBallsSlice[num+1:]... )
	//用for去比對我抽到的球是哪一個 比對到的時候把它取代掉 然後將slice長度減掉（舊的方法
	/*for i := 0;i<=len(maxAmountOfBallsSlice)-1;i++{
		if maxAmountOfBallsSlice[num] == maxAmountOfBallsSlice[i]{
			maxAmountOfBallsSlice[num] = maxAmountOfBallsSlice[len(maxAmountOfBallsSlice)-1]
		}
	}
	maxAmountOfBallsSlice = maxAmountOfBallsSlice[:len(maxAmountOfBallsSlice)-1]
	*/
	return result,""
}
