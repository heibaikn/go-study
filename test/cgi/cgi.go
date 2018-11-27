package cgi

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Test(arr []int) []int {
	//fmt.Println("test")
	fmt.Printf("arr %+v %p \n", arr, arr)
	//fmt.Println(arr[0])
	return []int{}
}

//var Arr1 [100]string

func CreateArr(arr []int) []int {
	//new(Arr1)

	for k := range arr {
		arr[k] = k
	}
	newArr, err := Random(arr[:], len(arr)-1)
	if err != nil {
		return nil
	}
	return newArr

	//fmt.Println("arr",arr)
}
func Random(arr []int, length int) ([]int, error) {
	if len(arr) <= 0 {
		return arr, errors.New("the length of the parameter arr should not be less than 0")
	}
	if length <= 0 || len(arr) <= length {
		return arr, errors.New("the size of the parameter length illegal")
	}
	//fmt.Println(arr)
	rand.Seed(time.Now().Unix())
	for i := len(arr) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		arr[i], arr[num] = arr[num], arr[i]
	}
	return arr, nil
}

// 冒泡
func BubbleSort(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	for i := 0; i < len(newArr); i++ {
		for j := 1; j < len(newArr)-i; j++ {
			if newArr[j] < newArr[j-1] {
				//交换
				newArr[j], newArr[j-1] = newArr[j-1], newArr[j]
			}
		}
	}
	return newArr
}

// 选择
func SelectSort(arr []int) []int {

	newArr := make([]int, len(arr))
	copy(newArr, arr)
	length := len(newArr)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ {
			if newArr[j] > newArr[maxIndex] {
				maxIndex = j
			}
		}
		newArr[length-i-1], newArr[maxIndex] = newArr[maxIndex], newArr[length-i-1]
	}
	return newArr
}

// 插入
func InsertSort(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	for i := 1; i < len(newArr); i++ {
		if newArr[i] < newArr[i-1] {
			j := i - 1
			temp := newArr[i]
			for j >= 0 && newArr[j] > temp {
				newArr[j+1] = newArr[j]
				j--
			}
			newArr[j+1] = temp
		}
	}
	return newArr
}

// 快速
func QuickSort(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	recursionSort(newArr, 0, len(newArr)-1)
	//fmt.Println()
	return newArr
}
func recursionSort(nums []int, left int, right int) {
	fmt.Println(nums, left, right)
	if left < right {
		pivot := partition(nums, left, right)
		fmt.Println(pivot)
		recursionSort(nums, left, pivot-1)
		recursionSort(nums, pivot+1, right)
	}

}
func partition(nums []int, left int, right int) int {
	for left < right {
		for left < right && nums[left] <= nums[right] {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		for left < right && nums[left] <= nums[right] {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}

	return left
}

// 基数排序
//求数据的最大位
func maxbit(values []int) (ret int) {
	ret = 1 //保存最大的位数
	var key int = 10
	iLen := len(values)
	for i := 0; i < iLen; i++ {
		for values[i] >= key {
			key = key * 10
			ret++
		}
	}
	return
}

func Radixsort(values []int) []int {
	key := maxbit(values)
	iLen := len(values)
	tmp := make([]int, iLen, iLen)
	count := new([10]int)
	radix := 1
	var i, j, k int
	for i = 0; i < key; i++ { //进行key次排序
		for j = 0; j < 10; j++ {
			count[j] = 0
		}
		for j = 0; j < iLen; j++ {
			k = (values[j] / radix) % 10
			count[k]++
		}

		for j = 1; j < 10; j++ { //将tmp中的为准依次分配给每个桶
			count[j] = count[j-1] + count[j]
		}
		for j = iLen - 1; j >= 0; j-- {
			k = (values[j] / radix) % 10
			tmp[count[k]-1] = values[j]
			count[k]--
		}
		for j = 0; j < iLen; j++ {
			values[j] = tmp[j]
		}
		radix = radix * 10
	}
	return values
}

// 桶排序
func BucketSort(arr []int) ([]int) {
	var newArr []int
	tmp := make([]int, 54)
	//fmt.Println(arr)
	//fmt.Println(tmp)
	for _, v := range arr {
		tmp[v]++
	}
	//fmt.Println(tmp)
	for k, v := range tmp {
		for v > 0 {
			newArr = append(newArr, k)
			v--
		}
	}
	//fmt.Println(newArr)
	return newArr
}
