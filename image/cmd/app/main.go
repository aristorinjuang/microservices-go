package main

import (
	"time"

	"github.com/aristorinjuang/microservices-go/image/internal"
	"github.com/aristorinjuang/microservices-go/pkg"
)

func main() {
	for {
		pkg.Clear()
		internal.Run()
		time.Sleep(time.Hour)
	}
}
