package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i:=0;i<5;i++{
		timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		//timestr := []rune(timestamp)
		//fmt.Println(timestr)
		time.Sleep(1)
		fmt.Println(timestamp[15:])
	}
}
