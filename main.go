package main

import (
	"os"
	"os/signal"
	"pingrobot/workerpool"
	"sync"
	"syscall"
	"time"
)

func main() {
	url := []string{
		"https://www.reddit.com/",
		"https://www.youtube.com/",
		"https://www.twitch.tv/",
	}

	createNewMem := workerpool.СreatePool(url)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	done := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				time.Sleep(time.Second * 7)
				workerpool.CheckCode(createNewMem)
			}
		}
	}()

	<-signalChan
	close(done)
	wg.Wait()
}
