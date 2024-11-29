package main

import (
	"log"
	"project/infra"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	log.Print(ctx)

	//r := routes.NewRoutes(*ctx)
	//
	//err = r.Run(":8080")
	//if  err != nil {
	//	log.Fatalf("failed to run server: %v", err)
	//}
}
