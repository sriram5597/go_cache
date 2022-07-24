package main

import (
	"fmt"
	"log"
	"net/http"

	goCache "github.com/sriram5597/go_cache/src/cache"
	router "github.com/sriram5597/go_cache/src/api/router"
)


func main() {
	fmt.Println("Go Cache")
	log.Println("Initializting cache...")
	goCache.NewCache()
	log.Println("Cache is initialized...")

	log.Println("Server running at 8000...")
	err := http.ListenAndServe(":8000", router.GetRouter())
	log.Println("Server started...")
	log.Fatal(err)
}
