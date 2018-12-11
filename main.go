package main

import (
	"fmt"
	"time"
)

func main() {
    for {
      t := time.Now()
      fmt.Printf("%s, sleep 10s...\n", t.Format("2006-01-02 15:04:05"))
      time.Sleep(time.Duration(10) * time.Second)
    }
}
