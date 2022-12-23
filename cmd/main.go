package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/kmisha/fan-in-pattern-go/actions"
	"github.com/kmisha/fan-in-pattern-go/models"
	service "github.com/kmisha/fan-in-pattern-go/services/entity"
	"github.com/kmisha/fan-in-pattern-go/services/statistics"
)

func main() {
	ctx := context.Background()
	statisticsContext, statisticsCansel := context.WithCancel(ctx)
	var wg sync.WaitGroup

	// create services
	productService := service.NewEntityService(ctx)
	userService := service.NewEntityService(ctx)
	cartService := service.NewEntityService(ctx)
	orderService := service.NewEntityService(ctx)

	statisticsCh := statistics.MergeData(
		statisticsContext, productService.Updates, userService.Updates, cartService.Updates, orderService.Updates)

	// mock data
	bike, _ := models.NewProduct("Bike")
	user := models.NewUser("John", "Doe")
	createProduct := actions.CreateProductAction{Name: "Bike"}
	updateProduct := actions.UpdateProductAction{NewName: "Super Bike", Product: bike}
	createCart := actions.CreateCartAction{Owner: user, Product: bike}
	createOrder := actions.CreateOrderAction{Owner: user}

	// do some stuffs
	wg.Add(1)
	go func() {
		defer wg.Done()
		productService.Do(&createProduct)
		productService.Do(&updateProduct)
		userService.Do(&actions.CreateUserAction{Name: "Bob", Surename: "Smith"})
		cartService.Do(&createCart)
		orderService.Do(&createOrder)
	}()

	// wait when we
	go func() {
		wg.Wait()
		statisticsCansel()
	}()

	// do some statistics
	stat := 0
	for msg := range statisticsCh {
		fmt.Println(msg)
		stat += 1
	}

	fmt.Printf("During the main task we did %d actions\n", stat)
}
