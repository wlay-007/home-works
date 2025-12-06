package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numsCh := make(chan int)
	sqrCh := make(chan int)
	go generateNumbers(numsCh)
	go powNumbers(numsCh, sqrCh)

	for num := range sqrCh {
		fmt.Println(num)
	}

}

func generateNumbers(channel chan int) {
	nums := make([]int, 10)

	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(101)
	}
	fmt.Print(nums)
	for _, value := range nums {
		channel <- value
	}
	close(channel)
}

func powNumbers(channelNumbers chan int, channelPow chan int) {
	for item := range channelNumbers {
		channelPow <- item * item
	}
	close(channelPow)
}
