package ratelimit

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Visitor is a struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type Visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Change the the map to hold values of the type visitor.
var (
	visitors = make(map[string]*Visitor)
	mu       sync.Mutex
)

// GetVisitor gets the visitor IP address.
func GetVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(1, 10)
		// Include the current time when creating a new visitor.
		visitors[ip] = &Visitor{limiter, time.Now()}
		return limiter
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	return v.limiter
}

// CleanupVisitors checks the map for visitors that haven't
// been seen for more than 2 minutes and delete the entries.
func CleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 2*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}
