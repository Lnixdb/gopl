package Context

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func TestCancelContext() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}

func timeout1(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout1")
	}
}

func timeout2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout2")
	}
}

func TestTimeoutContext() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)

	go timeout1(ctx)
	go timeout2(ctx)

	time.Sleep(1000 * time.Millisecond)
}

const keyID = "id"

func value1(ctx context.Context) {
	// do some work

	// we can get the value from the context by passing in the key
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	value2(ctx)
}

func value2(ctx context.Context) {
	// do some work

	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}

func TestValueContext() {
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), keyID, rand.Int())
	value1(ctx)
}
