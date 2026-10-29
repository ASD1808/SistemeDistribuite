package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2

	delta = t4.Sub(t1) - t3.Sub(t2)

	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	initialOffset := 100 * time.Millisecond
	serverClock := &Clock{offset: 0}
	clientClock := &Clock{offset: initialOffset}

	fmt.Println("O singură cerere")
	fmt.Printf("Offset inițial: %v\n", clientClock.offset)

	theta, delta := measure(clientClock, serverClock)

	clientClock.offset += theta

	fmt.Printf("theta: %v, delta: %v\n", theta, delta)
	fmt.Printf("Offset: %v\n\n", clientClock.offset)

	fmt.Println("10 cereri succesive")

	clientClock.offset = initialOffset

	var minDelta time.Duration = 1<<63 - 1
	var bestTheta time.Duration

	for i := 0; i < 10; i++ {
		t, d := measure(clientClock, serverClock)
		if d < minDelta {
			minDelta = d
			bestTheta = t
		}
	}
	clientClock.offset += bestTheta

	fmt.Printf("theta: %v, delta minim: %v\n", bestTheta, minDelta)
	fmt.Printf("Offset: %v\n", clientClock.offset)
}
