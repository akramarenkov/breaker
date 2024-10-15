package closing_test

import (
	"fmt"

	"github.com/akramarenkov/breaker/closing"
)

func ExampleClosing() {
	closing := closing.New()

	go func() {
		closing.Close()
	}()

	_, opened := <-closing.IsClosed()

	fmt.Println(opened)
	// Output:
	// false
}
