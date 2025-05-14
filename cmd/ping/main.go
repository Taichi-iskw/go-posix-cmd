package main

import (
	"fmt"
	"os"
	"sync"
)


func main(){
	if len(os.Args) < 2{
		fmt.Println("Usage: ping <hostname>")
		os.Exit(1)
	}

	hosts := os.Args[1:]
	results := make(chan string, len(hosts))
	var wg sync.WaitGroup

	for _, host := range hosts {
		wg.Add(1)
		go func(h string){
			defer wg.Done()
			result := PingHost(h)
			results <- result
		}(host)
	}

	wg.Wait()
	close(results)

	for result := range results{
		fmt.Println(result)
	}
	
}