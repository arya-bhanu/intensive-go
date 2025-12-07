package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrCircuitBreakerOpen = errors.New("circuit breaker is open")
	ErrTooManyRequests    = errors.New("too many requests in half-open state")
)

func main() {
	// Example usage of the circuit breaker
	fmt.Println("Circuit Breaker Pattern Example")

	// Create a circuit breaker configuration
	config := Config{
		MaxRequests: 3,
		Interval:    time.Minute,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(m Metrics) bool {
			return m.ConsecutiveFailures >= 3
		},
		OnStateChange: func(name string, from State, to State) {
			fmt.Printf("Circuit breaker %s: %s -> %s\n", name, from, to)
		},
	}

	cb := NewCircuitBreaker(config)

	// Simulate some operations
	ctx := context.Background()

	// Successful operation
	result, err := cb.Call(ctx, func() (interface{}, error) {
		return "success", nil
	})
	fmt.Printf("Result: %v, Error: %v\n", result, err)

	// Failing operation
	result, err = cb.Call(ctx, func() (interface{}, error) {
		return nil, errors.New("simulated failure")
	})
	fmt.Printf("Result: %v, Error: %v\n", result, err)

	// Print current state and metrics
	fmt.Printf("Current state: %v\n", cb.GetState())
	fmt.Printf("Current metrics: %+v\n", cb.GetMetrics())
}

type CircuitBreaker interface {
	Call(ctx context.Context, operation func() (interface{}, error)) (interface{}, error)
	GetState() State
	GetMetrics() Metrics
}

type State int

func (s State) String() string {
	switch s {
	case StateOpen:
		return "Open"
	case StateClosed:
		return "Closed"
	case StateHalfOpen:
		return "Half-Open"
	default:
		return "Unknown"
	}

}

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

type Metrics struct {
	Requests            int64
	Successes           int64
	Failures            int64
	ConsecutiveFailures int64
	LastFailureTime     time.Time
}
type Config struct {
	MaxRequests   uint32
	Interval      time.Duration
	Timeout       time.Duration
	ReadyToTrip   func(Metrics) bool
	OnStateChange func(name string, from State, to State)
}

type circuitBreakerImpl struct {
	name             string
	config           Config
	state            State
	metrics          Metrics
	lastStateChange  time.Time
	halfOpenRequests uint32
	mutex            sync.RWMutex
}

func (c *circuitBreakerImpl) GetMetrics() Metrics {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.metrics
}

func (c *circuitBreakerImpl) GetState() State {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.state
}

func (c *circuitBreakerImpl) Call(ctx context.Context, operation func() (interface{}, error)) (interface{}, error) {
	type returnCall struct {
		res interface{}
		err error
	}
	var mainCallFuncChan = make(chan returnCall)

	go func() {
		var resOp interface{}
		var errOp error
		if err := c.canExecute(); err != nil {
			resOp, errOp = nil, err
		} else {
			resOp, errOp = operation()
			if errOp != nil || resOp == nil {
				c.recordFailure()
			} else {
				c.recordSuccess()
			}
		}
		mainCallFuncChan <- returnCall{
			res: resOp,
			err: errOp,
		}
	}()
	select {
	case <-ctx.Done():
		return nil, errors.New("context cancelled")
	case result := <-mainCallFuncChan:

		res, err := result.res, result.err

		return res, err
	}

}

func (c *circuitBreakerImpl) setState(newState State) {
	if c.config.OnStateChange != nil {
		c.config.OnStateChange(c.name, c.GetState(), newState)
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.lastStateChange = time.Now()
	if newState != StateHalfOpen {
		c.halfOpenRequests = 0
	}
	c.state = newState
}
func (c *circuitBreakerImpl) canExecute() error {
	state := c.GetState()
	switch state {
	case StateClosed:
		return nil
	case StateOpen:
		if !c.isReady() {
			return ErrCircuitBreakerOpen
		}
		c.setState(StateHalfOpen)
		fallthrough
	case StateHalfOpen:
		if !c.isHalfOpenAllowRequest() {
			return ErrTooManyRequests
		}
	}
	return nil
}

func (c *circuitBreakerImpl) recordSuccess() {
	c.mutex.Lock()
	c.metrics.Successes++
	c.metrics.Requests++
	c.metrics.ConsecutiveFailures = 0
	c.mutex.Unlock()
	if c.GetState() == StateHalfOpen {
		c.setState(StateClosed)
	}
}

func (c *circuitBreakerImpl) recordFailure() {
	c.mutex.Lock()
	c.metrics.Failures++
	c.metrics.Requests++
	c.metrics.ConsecutiveFailures++
	c.metrics.LastFailureTime = time.Now()
	c.mutex.Unlock()

	if c.GetState() == StateHalfOpen {
		c.setState(StateOpen)
		return
	}

	if c.shouldTrip() {
		c.setState(StateOpen)
	}
}

func (c *circuitBreakerImpl) shouldTrip() bool {
	return c.config.ReadyToTrip(c.GetMetrics())
}

func (c *circuitBreakerImpl) isReady() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	timeout := c.config.Timeout
	lastStateChange := c.lastStateChange
	totalTime := lastStateChange.Add(timeout)
	now := time.Now()
	return now.After(totalTime)
}

func (c *circuitBreakerImpl) isHalfOpenAllowRequest() bool {

	if c.halfOpenRequests > c.config.MaxRequests {
		c.mutex.Lock()
		c.halfOpenRequests = 0
		c.mutex.Unlock()
		c.setState(StateClosed)
		return false
	}
	c.mutex.Lock()
	c.halfOpenRequests++
	c.mutex.Unlock()
	return true
}

func NewCircuitBreaker(config Config) CircuitBreaker {
	if config.MaxRequests == 0 {
		config.MaxRequests = 1
	}
	if config.Interval == 0 {
		config.Interval = time.Minute
	}
	if config.Timeout == 0 {
		config.Timeout = time.Second * 30
	}

	if config.ReadyToTrip == nil {
		config.ReadyToTrip = func(m Metrics) bool {
			return m.ConsecutiveFailures == 5
		}
	}
	return &circuitBreakerImpl{
		config:          config,
		name:            "circuit-breaker",
		state:           StateClosed,
		lastStateChange: time.Now(),
	}
}
