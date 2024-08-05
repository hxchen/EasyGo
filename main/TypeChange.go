package main

import (
	"fmt"
	"strconv"
)

func main() {
	int2float()
	string2int("123")
	int2string(1234)
	interface2string()
	writerTest()
	printValue(42)
	printValue("hello")
	printValue(3.14)
}

func int2float() {
	var sum int = 17
	var count int = 5
	var mean float32
	fmt.Printf("sum is %d\n", sum)
	fmt.Printf("count is %d\n", count)
	mean = float32(sum) / float32(count)
	fmt.Printf("mean is %f\n", mean)
}

func string2int(str string) {
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("string %s to int is %d\n", str, num)
	} else {
		fmt.Printf("string to int error, %s\n", err)
	}
}

func int2string(value int) {
	str := strconv.Itoa(value)
	fmt.Printf("integer %d to string is %s\n", value, str)
}

func interface2string() {
	var a interface{}
	a = "test"
	str, ok := a.(string)
	if ok {
		fmt.Printf("interface to string is %s\n", str)
	} else {
		fmt.Printf("interface to string error\n")
	}
}

// Writer 定义接口Writer
type Writer interface {
	Write([]byte) (int, error)
}

// StringWriter 定义结构体StringWriter
type StringWriter struct {
	str string
}

// Write 实现接口Writer的Write方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	if sw == nil {
		return 0, fmt.Errorf("nil StringWriter")
	}
	sw.str += string(data)
	return len(data), nil
}

func writerTest() {
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}

	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)

	// 修改 StringWriter 的字段
	sw.str = "Hello, World"

	// 打印 StringWriter 的字段值
	fmt.Println(sw.str)
}

func printValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("value is int %d\n", v)
	case string:
		fmt.Printf("value is string %s\n", v)
	default:
		fmt.Printf("value is unknown\n")
	}
}
