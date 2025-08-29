package main

import (
	"trainnig-api-poc/api"
)

func main() {
	r := api.SetupRouter()
	r.Run(":8080")
}
