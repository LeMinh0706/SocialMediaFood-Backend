package main

import (
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		return
	}
	fmt.Println("Run:", config.DBDriver)
}
