package main

import "fmt"
import "log"

func runTestRoutines() {
	defer un(trace("runTestRoutines"))
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime, ok := <-ch
		if !ok {
			fmt.Println("end")
			break
		}
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func generate(ch chan int) {
	for i := 2; i < 20; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}

	close(ch)
	log.Println("end of generate")
}

func filter(in, out chan int, prime int) {
	for {
		i, ok := <-in // Receive value of new variable 'i' from 'in'.
		if !ok {
			log.Printf("break of filter[%x]\n", &in)
			break
		}

		if i%prime != 0 {
			log.Printf("new filter : %d\n", i)
			out <- i // Send 'i' to channel 'out'.
		}
	}

	log.Printf("close of filter[%x]\n", &out)
	close(out)
}
