package main

import (
	"fmt"
	"reflect"
)

func detectType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("Тип: int, значение: %d\n", val)
	case string:
		fmt.Printf("Тип: string, значение: %q\n", val)
	case bool:
		fmt.Printf("Тип: bool, значение: %t\n", val)
	default:
		// Проверим, канал ли это
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Chan {
			fmt.Printf("Тип: chan, значение: %v\n", val)
		} else {
			fmt.Printf("Неизвестный тип: %T\n", val)
		}
	}
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	ch := make(chan int)
	detectType(ch)
	detectType(3.14)
}
