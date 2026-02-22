package main

import (
	"fmt"
	"go_task1/internal/greeting"
)

func main() {
	msg := greeting.Hello("World")
	fmt.Println(msg)
}
