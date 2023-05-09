package main

// ---------------------------
// yeh aggregation ke liye sahi hian
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main(){
// 	mych := make(chan int, 1024)
// 	wg := &sync.WaitGroup{}

// 	go hello1(mych, wg)
// 	wg.Add(1)

// 	go hello2(mych, wg)
// 	wg.Add(1)

// 	go hello3(mych, wg)
// 	wg.Add(1)

// 	wg.Wait()
// 	close(mych)

// 	for i := range mych{
// 		fmt.Println(i)
// 	}
// }

// func hello1(sendch chan int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	time.Sleep(100*time.Millisecond)
// 	sendch <- 1
// }

// func hello2(sendch chan int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	time.Sleep(1000*time.Millisecond)
// 	sendch <- 12
// }

// func hello3(sendch chan int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	time.Sleep(200*time.Millisecond)
// 	sendch <- 13
// }
// ---------------------------

// ---------------------------
// TODO
import (
	"context"
	"fmt"
	"log"
	"time"
)

type TaskFunc[T any] func(context.Context)(T, error)

type response[T any] struct {
	val T
	err error
}

type Task[T any] struct{
	respch chan response[T]
	ctx context.Context
	cancel context.CancelFunc
}

func (t *Task[T]) Await() (T, error){
	for {
		select {
		case <- t.ctx.Done():
			var val T
			return val, t.ctx.Err()  
		case resp := <- t.respch:
			return resp.val, resp.err
		}
	}
}

func (t *Task[T])Cancel(){
	t.cancel()
}

func SpawnWithTimeout[T any](fn TaskFunc[T], d time.Duration) *Task[T]{
	ctx, _ := context.WithTimeout(context.Background(), d)
	return spawn(ctx, fn)
}

func Spawn[T any](fn TaskFunc[T]) *Task[T]{
	ctx := context.Background()
	return spawn(ctx, fn)
}

func spawn[T any](ctx context.Context, fn TaskFunc[T])(*Task[T]){
	respch := make(chan response[T], 1024)

	// here we are wrapping our context with context with context.withcancel 
	ctx, cancel := context.WithCancel(ctx)

	// yeh background mein hoga abh
	go func(){
		val, err := fn(ctx);
		respch <- response[T]{
			val: val,
			err: err,
		}
	}()
	
	return &Task[T]{
		respch: respch,
		ctx: ctx,
		cancel: cancel,
	}
}

func main(){
	start := time.Now()
	t := SpawnWithTimeout(someApiCall, time.Millisecond*300)

	val, err := t.Await()
	if err != nil {
		log.Println("error awaiting", err) 
	}

	fmt.Println("Result: ", val)
	fmt.Println("Took: ", time.Since(start))
}

func someApiCall(ctx context.Context) (string, error){
	time.Sleep(time.Millisecond*300)
	return "working", nil
}

/*
TODO
1. fetchUserFromSlowDB iska decorator pattern hain voh dekhna hain (yeh sab abh ekh baari mein hee dekhke khatam hona chahiye)
2. 
*/