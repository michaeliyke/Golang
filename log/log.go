package log

import "fmt"

func typeof(x interface{}, types ...string) bool {
	if len(types) > 0 {
		var answer bool = false
		var typeofx string = fmt.Sprintf("%T", x)
		for _, t := range types {
			if t == typeofx {
				answer = true
				break
			}
		}
		return answer
	}
	fmt.Printf("%T\n", x)
	return false
}

func ShowType(x interface{}, types ...string) bool {
	return typeof(x, types...)
}

func TypeOf(x interface{}, types ...string) bool {
	return typeof(x, types...)
}

func Type(x interface{}, types ...string) bool {
	return typeof(x, types...)
}

func sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func print(x ...interface{}) {
	fmt.Print(x...)
}

func printf(x string, args ...interface{}) {
	fmt.Printf(x, args...)
}

func println(x ...interface{}) {
	fmt.Println(x...)
}

func Log(x ...interface{}) {
	Println(x...)
}

func L(x ...interface{}) {
	Println(x...)
}

func Pf(str string, args ...interface{}) {
	printf(str, args...)
}

func Print(x ...interface{}) {
	print(x...)
}

func Println(x ...interface{}) {
	println(x...)
}

func P(x ...interface{}) {
	Println(x...)
}

func Sp(str string, args ...interface{}) string {
	return sprintf(str, args...)
}

func Sprintf(str string, args ...interface{}) string {
	return sprintf(str, args...)
}
