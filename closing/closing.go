// Contains a Closing implementation that is used to closes a channel once.
package closing

import (
	"sync"
)

// Closes the channel, but only once.
type Closing struct {
	channel chan struct{}
	once    *sync.Once
}

// Creates Closing instance.
func New() *Closing {
	cls := &Closing{
		channel: make(chan struct{}),
		once:    &sync.Once{},
	}

	return cls
}

// Closes channel once.
//
// Method is thread-safe.
//
// Also you can specify additional functions that will be called once.
func (cls *Closing) Close(adds ...func()) {
	closer := func() {
		cls.close()

		for _, add := range adds {
			if add != nil {
				add()
			}
		}
	}

	cls.once.Do(closer)
}

func (cls *Closing) close() {
	close(cls.channel)
}

// Returns channel that will be closed by Close() method.
func (cls *Closing) IsClosed() <-chan struct{} {
	return cls.channel
}
