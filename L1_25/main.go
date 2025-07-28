package main

import (
	"fmt"
	"time"
)

// Sleep блокирует выполнение на указанную длительность
func Sleep(d time.Duration) {
	<-time.After(d) // блокируется до наступления времени
}

func main() {
	fmt.Println("Start:", time.Now())
	Sleep(2 * time.Second)
	fmt.Println("End  :", time.Now())
}
