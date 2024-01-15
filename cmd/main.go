package main

import (
	"fmt"

	"github.com/zulfianfreza/test-sharing-vision-backend/app"
)

func main() {
	_ = app.NewDB()
	fmt.Println("hello")
}
