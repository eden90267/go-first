package bar

import "fmt"

func init()  { // 所有 package 預設都會執行
	fmt.Println("bar init")
}

func HelloWorld() string {
	return "Hello bar"
}