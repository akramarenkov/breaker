package closing_test

import (
	"fmt"

	"github.com/akramarenkov/breaker/closing"
)

func ExampleClosing() {
	clg := closing.New()

	go func() {
		clg.Close()
	}()

	_, opened := <-clg.IsClosed()

	fmt.Println(opened)
	// Output:
	// false
}
