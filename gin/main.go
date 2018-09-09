package main

import (
	routes "gin-app/routes"

	_ "github.com/lib/pq"
)

func main() {
	r := routes.NewRouter()

	r.Run(":8080")
}
