package main

import "github.com/LeMinh0706/SocialMediaFood-Backend/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run(":8070")

}
