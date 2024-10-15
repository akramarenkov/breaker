package breaker_test

import (
	"fmt"

	"github.com/akramarenkov/breaker"
)

func ExampleBreaker() {
	breaker := breaker.New()

	go func() {
		defer breaker.Complete()

		_, opened := <-breaker.IsBreaked()

		fmt.Println(opened)
	}()

	breaker.Break()

	fmt.Println(breaker.IsStopped())
	// Output:
	// false
	// true
}
