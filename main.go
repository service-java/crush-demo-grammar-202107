package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("Hello, 世界")

	//initVar()

	//str()

	//zero()

	//shortVar()

	//iotaDemo()

	//strConvert()

	strMethod()
}

func initVar() {
	var i int = 10
	fmt.Println(i)

	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, ",f64 is", f64)
}

func str() {
	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("s1 is", s1, ",s2 is", s2, "s1+s2", s1+s2)
}

func zero() {
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)
}

func shortVar() {
	i := 10
	bf := false
	s1 := "Hello"
	print(i, bf, s1)
}

func iotaDemo() {
	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)
}

func strConvert() {
	i := 1
	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)

	f64 := 11222.211212
	i2f := float64(i)
	f2i := int(f64)
	fmt.Println(i2f, f2i)
}

func strMethod()  {
	s1 := "Ho121232312"
	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1,"H"))
	//在s1中查找字符串o
	fmt.Println(strings.Index(s1,"o"))
	//把s1全部转为大写
	fmt.Println(strings.ToUpper(s1))
}
