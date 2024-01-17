// Contains a Breaker implementation that is used to break a goroutine and
// wait for it to complete.
package breaker

import "github.com/akramarenkov/breaker/closing"

// Used to break goroutine and wait it completion.
type Breaker struct {
	completer   *closing.Closing
	interrupter *closing.Closing
}

// Creates Breaker instance.
func New() *Breaker {
	brk := &Breaker{
		completer:   closing.New(),
		interrupter: closing.New(),
	}

	return brk
}

// Closes channel returned by IsBreaked() method and wait call of Complete() method.
//
// Method is thread-safe, it can be called many times.
//
// Also you can specify additional functions that will be called once,
// before waiting for completion.
func (brk *Breaker) Break(adds ...func()) {
	brk.interrupter.Close(adds...)
	<-brk.completer.IsClosed()
}

// Returns channel that will be closed by Break() method.
func (brk *Breaker) IsBreaked() <-chan struct{} {
	return brk.interrupter.IsClosed()
}

// Marks a goroutine as completed.
func (brk *Breaker) Complete() {
	brk.completer.Close(nil)
}
