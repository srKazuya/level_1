package main

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func exitedByIf(wg *sync.WaitGroup) {
	wg.Add(1)

	done := true
	go func() {
		defer wg.Done()
		if done {
			fmt.Println("gorutine 1: exited by if")
			return
		}
	}()
}

func exitedByChannel(wg *sync.WaitGroup) {
	wg.Add(1)
	doneCh := make(chan bool)
	go func() {
		defer wg.Done()
		select {
		case <-doneCh:
			fmt.Println("gorutine 2: exited by channel")
			return
		}
	}()
	close(doneCh)
}

func exitedByContext(wg *sync.WaitGroup) {
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))

	go func() {
		defer wg.Done()
		defer cancel()
		select {
		case <-ctx.Done():
			fmt.Println("gorutine 3: exited by ctx")
		}
	}()
}
func exitedByTimer(wg *sync.WaitGroup) {
	wg.Add(1)

	timer := time.NewTimer(time.Duration(3 * time.Second))

	go func() {
		defer wg.Done()
		select {
		case <-timer.C:
			fmt.Println("gorutine 4: exited by timer")
		}
	}()
}
func exitedByTimeAfter(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		select {
		case <-time.After(time.Duration(4 * time.Second)):
			fmt.Println("gorutine 5: exited by timeAfter")
		}
	}()
}

func exitedByGoexit(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("gorutine 6: exited by runtime.Goexit")
		runtime.Goexit()
	}()
}

func exitedByOsSignal(wg *sync.WaitGroup) {
	wg.Add(1)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)

	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			stop()
			fmt.Println("gorutine 7: exited by ossignal")
		}
	}()

}

func exitedByPanic(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("gorutine 8: exited by panic")
			}
		}()
		panic("gorutine 8: exited by panic")
	}()
}

func main() {
	wg := &sync.WaitGroup{}

	exitedByIf(wg)
	exitedByChannel(wg)
	exitedByContext(wg)
	exitedByTimer(wg)
	exitedByTimeAfter(wg)
	exitedByGoexit(wg)
	exitedByOsSignal(wg)
	exitedByPanic(wg)

	wg.Wait()
}
