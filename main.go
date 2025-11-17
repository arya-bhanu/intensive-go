package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context Management Challenge")
	fmt.Println("Implement the context manager methods!")

	cm := NewContextManager()

	ctx, cancel := cm.CreateCancellableContext(context.Background())
	defer cancel()

	ctx = cm.AddValue(ctx, "user", "alice")
	ctx = cm.AddValue(ctx, "requestID", "12345")

	fmt.Println("Context created with values!")
}

type simpleContextManager struct{}

type ContextManager interface {
	// Create a cancellable context from a parent context
	CreateCancellableContext(parent context.Context) (context.Context, context.CancelFunc)

	// Create a context with timeout
	CreateTimeoutContext(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)

	// Add a value to context
	AddValue(parent context.Context, key, value interface{}) context.Context

	// Get a value from context
	GetValue(ctx context.Context, key interface{}) (interface{}, bool)

	// Execute a task with context cancellation support
	ExecuteWithContext(ctx context.Context, task func() error) error

	// Wait for a duration or until context is cancelled
	WaitForCompletion(ctx context.Context, duration time.Duration) error
}

func (s *simpleContextManager) CreateCancellableContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(parent)
}

func (s *simpleContextManager) CreateTimeoutContext(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {

}

func (s *simpleContextManager) AddValue(parent context.Context, key, value interface{}) context.Context {
	return context.WithValue(parent, key, value)
}

func (s *simpleContextManager) GetValue(ctx context.Context, key interface{}) (interface{}, bool) {
	val := ctx.Value(key)
	if val == nil {
		return nil, false
	}
	return val, true
}

func (s *simpleContextManager) ExecuteWithContext(ctx context.Context, task func() error) error {

}

func (s *simpleContextManager) WaitForCompletion(ctx context.Context, duration time.Duration) error {

}

// Simulate work that can be cancelled
func SimulateWork(ctx context.Context, workDuration time.Duration, description string) error {

}

// Process multiple items with context awareness
func ProcessItems(ctx context.Context, items []string) ([]string, error) {

}

func NewContextManager() *simpleContextManager {
	return &simpleContextManager{}
}
