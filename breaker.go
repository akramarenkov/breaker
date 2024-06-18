// Contains a Breaker implementation that is used to break a goroutine and
// wait for it to complete.
package breaker

import (
	"sync/atomic"

	"github.com/akramarenkov/breaker/closing"
)

// Used to break goroutine and wait it completion.
type Breaker struct {
	completer   *closing.Closing
	interrupter *closing.Closing
	stopper     *atomic.Bool
}

// Creates Breaker instance.
func New() *Breaker {
	brk := &Breaker{
		completer:   closing.New(),
		interrupter: closing.New(),
		stopper:     &atomic.Bool{},
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
	brk.stopper.Store(true)
	brk.interrupter.Close(adds...)

	<-brk.completer.IsClosed()
}

// Returns channel that will be closed by Break() method.
func (brk *Breaker) IsBreaked() <-chan struct{} {
	return brk.interrupter.IsClosed()
}

// Returns true if Break() method was called and false otherwise.
func (brk *Breaker) IsStopped() bool {
	return brk.stopper.Load()
}

// Marks a goroutine as completed.
func (brk *Breaker) Complete() {
	brk.completer.Close(nil)
}
