package log

import "fmt"

func Log(x interface{}) {
	fmt.Println(x)
}

func L(x interface{}) {
	fmt.Println(x)
}

func Pf(x interface{}) {
	fmt.Print(x)
}

func Print(x interface{}) {
	fmt.Print(x)
}

func Println(x interface{}) {
	fmt.Println(x)
}

func P(x interface{}) {
	fmt.Println(x)
}

func Sp(x string) string {
	return fmt.Sprintf(string(x))
}
