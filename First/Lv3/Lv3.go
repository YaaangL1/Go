package Lv3

import (
	"fmt"
	"math/rand"
	"time"
)

// BubbleSort 冒泡排序
func BubbleSort() {
	//生成随机数数组
	var (
		randomArray [100]int
	)
	rand.Seed(time.Now().Unix()) //使用时间生成真随机数
	//使用随机数补全数组
	for i := 0; i < 100; i++ {
		//生成10000以内的随机数
		randomArray[i] = rand.Intn(10000)
	}
	fmt.Println("before:", randomArray)
	//排序
	for i := 0; i < len(randomArray)-1; i++ {
		for j := 0; j < len(randomArray)-1-i; j++ {
			if randomArray[j] > randomArray[j+1] {
				var temp = randomArray[j+1]
				randomArray[j+1] = randomArray[j]
				randomArray[j] = temp
			}
		}
	}
	fmt.Println("after:", randomArray)
}

// SelectionSort 选择排序
func SelectionSort() {
	//生成随机数数组
	var (
		randomArray [100]int
	)
	rand.Seed(time.Now().Unix()) //使用时间生成真随机数
	//使用随机数补全数组
	for i := 0; i < 100; i++ {
		//生成10000以内的随机数
		randomArray[i] = rand.Intn(10000)
	}
	fmt.Println("before:", randomArray)
	//排序
	for i := 0; i < len(randomArray)-1; i++ {
		var (
			minIndex = i
			temp     int
		)
		//未排序数中找出最小值记录下标
		for j := i + 1; j < len(randomArray); j++ {
			if randomArray[minIndex] > randomArray[j] {
				minIndex = j
			}
		}
		//排序
		temp = randomArray[i]
		randomArray[i] = randomArray[minIndex]
		randomArray[minIndex] = temp
	}
	fmt.Println("after:", randomArray)
}

// InsertionSort 插入排序
func InsertionSort() {
	//生成随机数数组
	var (
		randomArray [100]int
	)
	rand.Seed(time.Now().Unix()) //使用时间生成真随机数
	//使用随机数补全数组
	for i := 0; i < 100; i++ {
		//生成10000以内的随机数
		randomArray[i] = rand.Intn(10000)
	}
	fmt.Println("before:", randomArray)
	//排序
	for i := range randomArray {
		preIndex := i - 1                                      //指向要排序的位置的前一位
		current := randomArray[i]                              //记录要排序的数的值
		for preIndex >= 0 && randomArray[preIndex] > current { //依次比较、移位赋值
			randomArray[preIndex+1] = randomArray[preIndex]
			preIndex -= 1
		}
		randomArray[preIndex+1] = current //最后的位置赋值为需要排序的值
	}
	fmt.Println("after:", randomArray)
}
