package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	value := "1604345905"
	integer, _ := strconv.ParseInt(value, 10, 0)
	date := time.Unix(integer, 0)
	fmt.Println(date)
}
