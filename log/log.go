package log

import "fmt"

func typeof(x interface{}) {
	var typeof string = fmt.Sprintf("%T", x)
	// fmt.Printf("%T\n", x)
	P(typeof)
}

func ShowType(x interface{}) {
	typeof(x)
}

func Type(x interface{}) {
	typeof(x)
}

func TypeOf(x interface{}) {
	typeof(x)
}

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
