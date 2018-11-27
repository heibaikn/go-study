package playLandlord

import (
	"errors"
	"fmt"
	"reflect"
	"test/cgi"
	"test/global"
)

//type RoomInfo struct {
//	Player1 Player
//	Player2 Player
//	Player3 Player
//}
type Player struct {
	Name      int
	HandCard  []int
	LevelCard []int
}

func DealCard(arr []int) ([]Player, []int) {
	var roomInfo []Player
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	fmt.Println("洗牌 \n", arr)
	for i := 0; i < 3; i++ {
		var player Player
		player.Name = i
		player.HandCard = cgi.BucketSort(newArr[i*17 : (i+1)*17])
		player.LevelCard = GetCardLevel(player.HandCard[:])
		roomInfo = append(roomInfo, player)
	}
	return roomInfo, newArr[51:]
	//fmt.Println(roomInfo)

}

func GetCardLevel(arr []int) []int {
	tmp := make([]int, 54)
	level := make([]int, 15)
	for _, v := range arr {
		tmp[v]++
	}
	// 1-13 正常牌
	for i := 0; i < 13; i++ {
		s := tmp[i*4 : (i+1)*4]
		time := 0
		for _, v := range s {
			if v == 1 {
				time++
			}
		}

		level[i] = time
	}
	// 小王
	if tmp[52] == 1 {
		level[13] = 1
	}
	//大王
	if tmp[53] == 1 {
		level[14] = 1
	}
	return level
}

func getArrayValLength(arr []int) int {
	i := 0
	for _, v := range arr {
		if v != 0 {
			i++
		}
	}
	return i
}

func Get3With2LevelArr(arr []int) (map[int][]int, error) {
	level := GetCardLevel(arr)
	fmt.Println(level)
	//var Twith2 map[int]int
	Twith2 := make(map[int][]int)
	var l3, l2 []int

	//l3,l2 := make([]int,13),make([]int,13)
	for k, v := range level {
		if v == 3 {
			l3 = append(l3, k)
		}
	}
	if len(l3) == 0 {
		fmt.Println("无3张连牌")
		return nil, errors.New("无3张连牌")
	}
	Twith2[0] = l3

	for k, v := range level {
		if v == 2 {
			l2 = append(l2, k)
		}
	}
	if len(l2) == 0 {
		fmt.Println("无2张连牌")
		return nil, errors.New("无2张连牌")
	}
	Twith2[1] = l2

	fmt.Println(Twith2, len(Twith2))
	return Twith2, nil
	//for ; ;  {
	//
	//}
}

func Is3With2(arr []int) int {
	i := -1
	for k, v := range arr {
		if v == 3 {
			i = k
		}
	}
	return i
}

func Compare3with2(dist []int, arr []int) (bool, error) {
	distArr := GetCardLevel(dist)
	currentArr := GetCardLevel(arr)
	l := getArrayValLength(currentArr)
	if l != 2 {
		return false, errors.New("牌型错误")
	}
	if Is3With2(currentArr) > Is3With2(distArr) {
		return true, nil
	}
	fmt.Println(distArr, currentArr)
	return false, errors.New("无法大过")
}

func GetBig3with2AI(dist []int, arr []int)[]int {

	Twith2, err := Get3With2LevelArr(arr)
	if err != nil {
		return nil
	}
	// 找到大于值
	var currentIdx int
	distArr := GetCardLevel(dist)
	disIdx := Is3With2(distArr)
	fmt.Println("上家出牌", dist, Is3With2(distArr))
	for _, v := range Twith2[0] {
		//fmt.Println(v)
		if v > disIdx {
			currentIdx = v
			break
		}
	}
	if currentIdx == 0 {
		fmt.Println("无大牌")
		return nil
	}
	fmt.Println(currentIdx, Twith2[1][0], distArr)
	return mapLevel2Value3With2(arr, currentIdx, Twith2[1][0])

}

func mapLevel2Value3With2(dist []int, l3 int, l2 int)[]int {
	fmt.Println(dist, l3, l2, (l3+0)*4-1, (l3+1)*4-1)
	var arr []int
	//append(arr, dist[])
	for _, v := range dist {
		if (v > l3*4-1) && (v < (l3+1)*4) {
			arr = append(arr, v)
		}
		if (v > l2*4-1) && (v < (l2+1)*4) {
			arr = append(arr, v)
		}
	}
	return arr
}

func DelHandCards(dist []int, input []int) []int {
	var arr []int
	for _, v := range dist {
		a, _ := global.Contain(input, v)
		if !a {
			arr = append(arr, v)
		}
	}
	return arr

}

func struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//
//type RoomInfo struct {
//	Play1 Play
//	Play2 Play
//	Play3 Play
//}
//type Play struct {
//	Name int
//	HandCard []int
//	LevelCard string
//}
//func DealCard(arr []int)  {
//	var roomInfo RoomInfo
//
//	roomInfo.Play1.HandCard = arr[:17]
//	roomInfo.Play2.HandCard = arr[17:34]
//	roomInfo.Play3.HandCard = arr[34:51]
//	//roomInfo.Bottom = arr[51:]
//	//bottom :=arr[51:]
//
//	data := struct2Map(roomInfo)
//	k := reflect.TypeOf(roomInfo)
//value := reflect.ValueOf(roomInfo)
//	//fmt.Println(value)
//	for i := 0; i < k.NumField(); i++ {
//		p := k.Field(i).Name
//		fmt.Println(p)
//
//	}
//	fmt.Printf("%+v \n",data)
//	for k,v:=range data {
//		fmt.Printf("%+v,%v \n",v,k)
//		v.LevelCard = "123"
//		//v.LevelCard =[]int
//	}
//}
//
//func struct2Map(obj interface{}) map[string]interface{} {
//	t := reflect.TypeOf(obj)
//	v := reflect.ValueOf(obj)
//
//	var data = make(map[string]interface{})
//	for i := 0; i < t.NumField(); i++ {
//		data[t.Field(i).Name] = v.Field(i).Interface()
//	}
//	return data
//}
