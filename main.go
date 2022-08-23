package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	log.Printf("Starting coffee shop service")
	http.HandleFunc("/serve-customer/", ServeCustomer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func ServeCustomer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	numCustomers, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/serve-customer/"))
	if err != nil || numCustomers == 0 {
		numCustomers = 1
	}

	wg := sync.WaitGroup{}
	count := 0
	for i := 0; i < numCustomers; i++ {
		wg.Add(1)
		go MakeCoffee(&wg)
		count++
	}
	wg.Wait()

	timeTaken := time.Since(start)
	log.Printf("Took %s to serve coffee to %v customer(s)", timeTaken, count)
}

func MakeCoffee(wg *sync.WaitGroup) {
	defer wg.Done()

	newWg := sync.WaitGroup{}
	newWg.Add(3)
	go PayForCoffee(&newWg)
	go MakeEspresso(&newWg)
	go SteamMilk(&newWg)
	newWg.Wait()
}

func PayForCoffee(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	log.Printf("Coffee paid for 💰")
}

func MakeEspresso(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	log.Printf("Espresso made ☕️")
}

func SteamMilk(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	log.Printf("Milk steamed 🥛")
}
