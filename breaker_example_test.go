package breaker_test

import (
	"fmt"

	"github.com/akramarenkov/breaker"
)

func ExampleBreaker() {
	brk := breaker.New()

	go func() {
		defer brk.Complete()

		_, opened := <-brk.IsBreaked()

		fmt.Println(opened)
	}()

	brk.Break()

	fmt.Println(brk.IsStopped())
	// Output:
	// false
	// true
}
