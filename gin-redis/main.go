package main

import (
	"fmt"
	"main/ginredis"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		ginredis.Set("first", "value")
	}

	fmt.Printf("Run set cost %d ms", (time.Now().UnixNano()-start)/1000000)

}
