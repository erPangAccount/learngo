package persist

import "log"

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item server : got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}