package breaker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBreaker(t *testing.T) {
	testBreaker(t)

	counter := 0

	add := func() {
		counter++
	}

	testBreaker(t, add, nil)
	require.Equal(t, 1, counter)
}

func testBreaker(t *testing.T, adds ...func()) {
	breaker := New()

	select {
	case <-breaker.IsBreaked():
		require.FailNow(t, "must not be breaked")
	default:
	}

	require.False(t, breaker.IsStopped())

	go func() {
		defer breaker.Complete()

		<-breaker.IsBreaked()
	}()

	breaker.Break(adds...)

	select {
	case <-breaker.IsBreaked():
	default:
		require.FailNow(t, "must be breaked")
	}

	require.True(t, breaker.IsStopped())

	breaker.Break(adds...)

	select {
	case <-breaker.IsBreaked():
	default:
		require.FailNow(t, "must be breaked")
	}

	require.True(t, breaker.IsStopped())
}
