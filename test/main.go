package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

//
//type I interface {
//	Get() int
//	Set(int)
//}
//type SS struct {
//	Age int
//}
//func (s SS) Get() int {
//	return s.Age
//}
//func (s SS) Set(age int) {
//	s.Age = age
//}
//func f(i I) {
//	i.Set(10)
//	fmt.Println(i.Get())
//}
//func main(){
//	ss := SS{}
//	f(&ss) //ponter
//	f(ss)  //value
//}


type item struct {
	Name string
}

func (i item) String() string {
	return fmt.Sprintf("item name: %v", i.Name)
}

type person struct {
	Name string
	Sex  string
}

func (p person) String() string {
	return fmt.Sprintf("person name: %v sex: %v", p.Name, p.Sex)
}

func Parse(i interface{}) interface{} {
	//fmt.Println(i)
	t :=fmt.Sprintf("%T",i)
	fmt.Println(t)
	switch i.(type) {
	case string:
		return &item{
			Name: i.(string),
		}
	case []string:
		data := i.([]string)
		length := len(data)
		if length == 2 {
			return &person{
				Name: data[0],
				Sex:  data[1],
			}
		} else {
			return nil
		}
	case int:
		var numbers map[string]int
		//numbers := make(map[string]int)
		numbers["one"] = 1  //赋值
		numbers["ten"] = 10 //赋值
		numbers["three"] = 3
		//a.a="1"
		//a["a"]="a"
		//a["b"]="b"
		//a=append(a, "1")
		//a=append(a, "2")
		return numbers
	default:
		panic(errors.New("type match miss"))
	}
	return nil
}

type Rule struct {
	value interface{}
	name string
	reg string
	msg string
}

func Check(aa []Rule)string  {
	for _, num := range aa {
		fmt.Println("value %v ,%#T /n",num.value,num.value)

		switch num.value.(type) {
		case string:
			m, _ := regexp.MatchString(num.reg, num.value.(string))
			fmt.Print(m)
			if !m {
				return num.msg
			}
		case int:
			fmt.Println(1>0)
		default:
			fmt.Printf("default")
		}

	}
	return ""
}
func main() {
	a :="2"
	//b :=2
	int,err:=strconv.Atoi(a)
	fmt.Println(int,err)
	fmt.Print(a)

	//m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname"));
	//var aa []Rule
	//aa=append(aa, Rule{name: "Alice", value: 30,reg:">10",msg:"123"})
	//aa=append(aa, Rule{name: "Alice", value: 30,reg:">11",msg:"123"})
	//aa=append(aa, Rule{name: "Alice", value: "1234",reg:"^[0-9]+$",msg:"123"})
	//fmt.Print(Check(aa))
	//make()
	//rule := {
	//value: 1,
	//	key: "id",
	//	//reg:
	//}

}