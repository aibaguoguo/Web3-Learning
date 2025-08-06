package main

// 指针 题目1：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func plusTen(num *int) {
	*num += 10
}

// 指针 题目2：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func multiplyByTwo(nums []int) {
	for i, _ := range nums {
		nums[i] *= 2
	}
}
