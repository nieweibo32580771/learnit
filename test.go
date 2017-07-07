package main

import "fmt" // alias3
// func main() {
// 	str := "我在这里"
// 	l := len(str)
// 	fmt.Println(l)
// 	// for i := 0; i <= l; i++ {
// 	// 	jie := str[i]
// 	// 	fmt.Println(i, jie)
// 	// }
// 	fmt.Println(4 >> 2)  //1
// 	fmt.Print(4 >> 2)    //1
// 	fmt.Println(4 << 2)  //16
// 	fmt.Print(4 << 2)    //16
// 	fmt.Print(^3)        //-4
// 	fmt.Print(^-5)       //4
// 	fmt.Println(^0)      //-1
// 	fmt.Println(1 ^ 2)   //30
// 	fmt.Println(11 ^ 21) //30
// 	const a = 12
// 	fmt.Printf("%d", a) //30
// 	const b = "我在这里你在哪里"
// 	fmt.Printf("%s", b)

// 	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
// 	modify(array)                  // 传递给一个函数，并试图在函数体内修改这个数组内容
// 	fmt.Println("In main(), array values:", array)

// }

// func modify(array [10]int) {
// 	array[0] = 10 // 试图修改数组的第一个元素
// 	fmt.Println("In modify(), array values:", array)
// }
func main() {
	// =====================================================================
	// 数组
	// =====================================================================
	// arr0 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(arr0)
	// arr1 := make([]int, 3)
	// fmt.Println(arr1)
	// arr2 := []int{11, 22, 33, 44}

	// // copy(arr0, arr2)
	// // fmt.Println(arr0)
	// // fmt.Println(arr2

	// copy(arr2, arr0)
	// fmt.Println(arr0)
	// fmt.Println(arr2)

	// fmt.Println(append(arr0, 1)) //arr0的赋值品
	// fmt.Println(arr0)            //arro0
	// =====================================================================
	// 数组切片
	// =====================================================================
	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := []int{5, 4, 3}

	// copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	// fmt.Println(slice2)
	// fmt.Println(slice1)

	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	// fmt.Println(slice2)
	// fmt.Println(slice1)

	// slice4 := make([]int, 5)
	// fmt.Println(slice4)
	// fmt.Println(append(slice4, slice2...)) //arr0的赋值品
	// fmt.Println(slice4)
	// =====================================================================
	// 字典
	// =====================================================================

	// var m1 map[int]string //声明  值是nil 无法赋值
	// fmt.Println(m1)

	// m1 = map[int]string{1: "1234"} //赋值
	// m1[2] = "123456"               //赋值
	// fmt.Println(m1)
	// fmt.Println(len(m1))

	// m3 := map[int]string{1: "12345"} //创建
	// fmt.Println(m3)

	// mymap := make(map[string]int, 3) //创建
	// fmt.Println(mymap)

	// mymap["1"] = 123 //赋值
	// fmt.Println(mymap)

	// m2 := map[string]int{"a": 1}
	// fmt.Println(m2)

	// m2["b"] = 34
	// fmt.Println(m2)

	// delete(m2, "b") //删除
	// fmt.Print(m2)

	//===============================================================
	//fun函数
	//===================================================================
	x := 1
	y := &x
	fmt.Println(y)
	fmt.Println(*y)

	myfunc(2, 3, 4)

	// .struct无序和有序
	type s struct {
		name, sex string
		age       int
	}
	i := s{
		"nie",
		"nan",
		12,
	}
	fmt.Println(i)
	fmt.Println(i.name)
	j := s{
		age: 78,
	}
	fmt.Println(j.age) //78
	var d func()
	var f = func() {
		fmt.Println("123")
	}
	f() //123
	d = f
	d()              //123
	g()              //nieshi
	fmt.Println(g()) //nieshi  1
	fmt.Println(h()) //2
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.printf("%c", x)
		}
	}

}

// func myfunc(args ...int) {

// 	for _, i := range args {
// 		fmt.Println(i)
// 	}
// }
// func MyPrintf(args ...interface{}) {
// 	for _, arg := range args {
// 		switch arg.(type) {
// 		case int:
// 			fmt.Println(arg, "is an int value")
// 		case string:
// 			fmt.Println(arg, "is an string value")
// 		case bool:
// 			fmt.Println(arg, "is an bool value")
// 		}
// 	}
// }

// func myfunc(x, y int) int {
// 	return x + y
// }
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func g() (result int) {
	defer func() {
		fmt.Println("nieshi")
		result++
	}()
	return 1
}
func h() (result int) {
	// defer func() {
	// 	result++
	// }()
	return
}
