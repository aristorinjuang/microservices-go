package main

import (
	"time"

	"github.com/aristorinjuang/microservices-go/pkg"
	"github.com/aristorinjuang/microservices-go/user/internal"
)

func main() {
	for {
		pkg.Clear()
		internal.Run()
		time.Sleep(time.Hour)
	}
}
