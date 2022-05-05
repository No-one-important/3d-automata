package main

import (
	"time"
)

// Clock stores elapsed times
type Clock struct {
	previousTime float64
	elapsed      float64
}

// Init set the first reference time
func (c *Clock) Init() {
	c.previousTime = float64(time.Now().UnixNano()) / 1e9
}

// Tic sets the clock's reference time to now
func (c *Clock) Tic() {
	time := float64(time.Now().UnixNano()) / 1e9
	c.elapsed = time - c.previousTime
	c.previousTime = time
}

// Toc retreive the time difference between now and the
// clock's reference time
func (c *Clock) Toc() float64 {
	time := float64(time.Now().UnixNano()) / 1e9
	return time - c.previousTime
}
