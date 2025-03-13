package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_PANEER_PRICE float32 = 4

func main() {
	var chickenChannel = make(chan string)

	var paneerChannel = make(chan string)

	var websites = []string{"daraz.com", "bigmart.com", "pathao.com", "bhatbhateni.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkPaneerPrices(websites[i], paneerChannel)
	}
	sendMessage(chickenChannel, paneerChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Millisecond * 1000)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkPaneerPrices(website string, paneerChannel chan string) {
	for {
		time.Sleep(time.Millisecond * 1000)
		var pannerPrice = rand.Float32() * 20
		if pannerPrice <= MAX_CHICKEN_PRICE {
			paneerChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, paneerChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound the chicken at lowest price %s", website)
	case website := <-paneerChannel:
		fmt.Printf("\nFound the paneer at lowest price %s", website)

	}
}
