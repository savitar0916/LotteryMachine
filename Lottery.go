package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Lottery struct{
	MaxAmountOfBalls int
}
// 封裝
var maxAmountOfBallsSlice []int

func (l *Lottery)ShowmaxAmountOfBallsSlice(){
	fmt.Println(maxAmountOfBallsSlice)
}

func (l *Lottery)SetMaxAmountOfBalls(){
	for i := 1; i <= l.MaxAmountOfBalls; i++ {
		maxAmountOfBallsSlice = append(maxAmountOfBallsSlice, i)
	}
}


func (l *Lottery)Shuffle()[]int{
	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

	for i := len(maxAmountOfBallsSlice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		maxAmountOfBallsSlice[i], maxAmountOfBallsSlice[j] = maxAmountOfBallsSlice[j], maxAmountOfBallsSlice[i]
	}
	return maxAmountOfBallsSlice
}

func (l *Lottery) GetOneNumber()(int,string){
	if len(maxAmountOfBallsSlice) == 1 {
		return maxAmountOfBallsSlice[0],"no more any num"
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(maxAmountOfBallsSlice))
	result := maxAmountOfBallsSlice[num]

	for i := 0;i<=len(maxAmountOfBallsSlice)-1;i++{
		if maxAmountOfBallsSlice[num] == maxAmountOfBallsSlice[i]{
			maxAmountOfBallsSlice[num] = maxAmountOfBallsSlice[len(maxAmountOfBallsSlice)-1]
		}
	}
	maxAmountOfBallsSlice = maxAmountOfBallsSlice[:len(maxAmountOfBallsSlice)-1]
	
	return result,""
}
