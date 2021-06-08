package log

import "fmt"

func typeof(datum interface{}, types ...string) bool {
	if len(types) > 0 {
		var answer bool = false
		var typeofdatum string = fmt.Sprintf("%T", datum)
		for _, t := range types {
			if t == typeofdatum {
				answer = true
				break
			}
		}
		return answer
	}
	fmt.Printf("%T\n", datum)
	return false
}

func ShowType(datum interface{}, types ...string) bool {
	return typeof(datum, types...)
}

func TypeOf(datum interface{}, types ...string) bool {
	return typeof(datum, types...)
}

func Type(datum interface{}, types ...string) bool {
	return typeof(datum, types...)
}

func sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func print(args ...interface{}) {
	fmt.Print(args...)
}

func printf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

func println(args ...interface{}) {
	fmt.Println(args...)
}

func Printf(str string, args ...interface{}) {
	printf(str, args...)
}

func Pf(str string, args ...interface{}) {
	printf(str, args...)
}

func Log(args ...interface{}) {
	Println(args...)
}

func L(args ...interface{}) {
	Println(args...)
}

func Print(args ...interface{}) {
	print(args...)
}

func Println(args ...interface{}) {
	println(args...)
}

func P(args ...interface{}) {
	Println(args...)
}

func Sprintf(str string, args ...interface{}) string {
	return sprintf(str, args...)
}

func Sp(str string, args ...interface{}) string {
	return sprintf(str, args...)
}
