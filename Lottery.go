package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MaxAmountOfBallsSlice []int

func SetMaxAmountOfBalls(Max int){

	i := 1
	for i <= Max{
		MaxAmountOfBallsSlice = append(MaxAmountOfBallsSlice, i)
		i ++
	}
	fmt.Println(MaxAmountOfBallsSlice)
	
}


func Shuffle(slice []int)[]int{
	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

	for i := len(slice) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func GetOneNumber(slice []int)[]int{
	if len(slice) == 1 {
		fmt.Println("抽完囉")
		return slice
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(slice))
	fmt.Println("抽到",slice[num])
	for i := 0;i<=len(slice)-1;i++{
		if slice[num] == slice[i]{
			slice[i] = slice[len(slice)-1]
		}
	}
	return slice[:len(slice)-1]
}
