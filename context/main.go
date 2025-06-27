package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, name string) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println(name, "finished work")
	case <-ctx.Done():
		fmt.Println(name, "cancelled:", ctx.Err())
	}
}

func main() {
	// Example 1: context.WithCancel
	ctx1, cancel := context.WithCancel(context.Background())
	go doWork(ctx1, "worker1")
	time.Sleep(1 * time.Second)
	cancel() // cancel the context before work is done

	// Example 2: context.WithTimeout
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()
	go doWork(ctx2, "worker2")
	time.Sleep(2 * time.Second)

	// Example 3: context.WithValue
	ctx3 := context.WithValue(context.Background(), "userID", 42)
	printUserID(ctx3)
}

func printUserID(ctx context.Context) {
	if v := ctx.Value("userID"); v != nil {
		fmt.Println("userID from context:", v)
	}
}
