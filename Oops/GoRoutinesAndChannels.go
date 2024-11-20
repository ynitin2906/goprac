package oops

import "fmt"

func sumPart(nums []int, result chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	result <- sum
}

func RunGoRoutineAndChannel() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int)

	go sumPart(nums[:len(nums)/2], result)
	go sumPart(nums[len(nums)/2:], result)

	sum1 := <-result
	sum2 := <-result

	fmt.Println("sum1:", sum1, "sum2:", sum2)

}
