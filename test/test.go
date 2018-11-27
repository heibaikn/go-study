package main

import (
	"fmt"
	"math/rand"
	"test/cgi"
	"test/playLandlord"
	"time"
)

//var roomInfo

func main() {

	card()

}

func card() {
	var arr [54]int
	// 生成随机数组
	t1 := time.Now()
	playerInfo, threeCards := playLandlord.DealCard(cgi.CreateArr(arr[:]))
	fmt.Printf("发牌 \n %+v \n", playerInfo)
	fmt.Println("底牌为",threeCards)
	fmt.Printf("create arr cost time%v \n\n\n", time.Since(t1))

	// 产生地主
	rand.Seed(time.Now().Unix())
	landlord := rand.Intn(3)
	fmt.Printf("第%v的玩家为地主 \n",landlord)
	playerInfo[landlord].HandCard = cgi.BucketSort(append(playerInfo[landlord].HandCard, threeCards...))
	playerInfo[landlord].LevelCard = playLandlord.GetCardLevel(playerInfo[landlord].HandCard[:])
	//fmt.Println()
	fmt.Printf("抢完地主后的值为 \n %v \n\n",playerInfo)

	// 判断大小 [1,2,3,4,5] 判断三家是否大过 3个3带2
	fmt.Println("有3带2")
	tableCards :=[]int{1,2,3,4,5}
	tableCards1 :=[]int{8,9,10,13,14}
	bl,_ := playLandlord.Compare3with2(tableCards,tableCards1)
	fmt.Println("[]int{8,9,10,13,14}是否大于{1,2,3,4,5}",bl)
	bl1,_ := playLandlord.Compare3with2(tableCards1,tableCards)
	fmt.Println("[]int{1,2,3,4,5}是否大于{8,9,10,13,14}",bl1)

	// 机器人 是否能大过  3个3带2
	fmt.Println("\n\n\n玩家0是否有3带2",playerInfo[0].HandCard[:])
	pushCards := playLandlord.GetBig3with2AI(tableCards,playerInfo[0].HandCard[:])
	fmt.Println("打出的牌",pushCards)

	// 机器人 打出手牌  不拆牌
	newHandCards :=playLandlord.DelHandCards(playerInfo[0].HandCard[:],pushCards)
	fmt.Println("机器人新手牌为",newHandCards)
	//fmt.Println(newHandCards)


	//Twith2,err :=playLandlord.Get3With2(playerInfo[0].HandCard[:])
	//slice.SliceDiff()
	//if err !=nil{
	//	return
	//}
	//fmt.Println()

}

func sort() {
	var arr [1000000]int
	// 生成随机数组
	t1 := time.Now()
	newArr := cgi.CreateArr(arr[:])
	fmt.Println("newArr", newArr[5])
	fmt.Println("create arr cost time ", time.Since(t1))
	//js 1k=> 5.28466796875ms 10k=>219.3310546875ms  100k=>26110.3701171875ms

	// 冒泡排序
	// 1k=> 890.833µs 10k=>104.723879ms  100k=>10.087987521s
	//t1 = time.Now()
	//bubbleArr := cgi.BubbleSort(newArr[:])
	//fmt.Println("bubbleArr",bubbleArr[5])
	//fmt.Println("冒泡排序 time ",time.Since(t1))

	// 选择排序
	// 1k=>1.028137ms 10k=>107.419876ms  100k=>9.945407391s
	//fmt.Println("newArr",newArr[5])
	//t1 = time.Now()
	//SelectArr := cgi.SelectSort(newArr[:])
	//fmt.Println("SelectArr",SelectArr[5])
	//fmt.Println("选择排序 time ",time.Since(t1))

	//	插入排序
	// 1k=> 156.285µs 10k=>16.065024ms 100k=> 1.504980468s
	//t1 = time.Now()
	//insertSort := cgi.InsertSort(newArr[:])
	//fmt.Println("insertSort",insertSort[5])
	//fmt.Println("插入排序 time ",time.Since(t1))

	//	快速排序
	// 1k=> 60.997µs 10k=>668.357µs  100k=> 8.658123ms
	//t1 = time.Now()
	//quickSort := cgi.QuickSort(newArr[:])
	//fmt.Println("quickSort",quickSort[5])
	//fmt.Println("快速排序 time ",time.Since(t1))

	////	桶排序
	//// 1k=> 36.274µs 10k=>335.531µs  100k=>3.621272ms
	//t1 = time.Now()
	//bucketSort := cgi.BucketSort(newArr[:])
	//fmt.Println("BucketSort",bucketSort[5])
	//fmt.Println("桶排序 time ",time.Since(t1))

}
