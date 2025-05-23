package closing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClosing(t *testing.T) {
	testClosing(t)

	counter := 0

	add := func() {
		counter++
	}

	testClosing(t, add, nil)
	require.Equal(t, 1, counter)
}

func testClosing(t *testing.T, adds ...func()) {
	closing := New()

	select {
	case <-closing.IsClosed():
		require.FailNow(t, "must not be closed")
	default:
	}

	closing.Close(adds...)

	select {
	case <-closing.IsClosed():
	default:
		require.FailNow(t, "must be closed")
	}

	require.NotPanics(t, func() { closing.Close(adds...) })

	select {
	case <-closing.IsClosed():
	default:
		require.FailNow(t, "must be closed")
	}
}

func BenchmarkClosed(b *testing.B) {
	closing := New()

	closing.Close()

	for range b.N {
		<-closing.IsClosed()
	}
}

func BenchmarkUnclosed(b *testing.B) {
	closing := New()

	for range b.N {
		select {
		case <-closing.IsClosed():
		default:
		}
	}
}

func BenchmarkRace(b *testing.B) {
	closing := New()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				closing.Close()
			}
		},
	)
}

func BenchmarkRaceAdd(b *testing.B) {
	closing := New()

	counter := 0

	add := func() {
		counter++
	}

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				closing.Close(add, nil)
			}
		},
	)
}
