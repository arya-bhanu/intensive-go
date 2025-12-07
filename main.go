package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context Management Challenge")
	fmt.Println("Implement the context manager methods!")

	// Example of how the context manager should work:
	cm := NewContextManager()

	// Create a cancellable context
	ctx, cancel := cm.CreateCancellableContext(context.Background())
	defer cancel()

	// Add some values
	ctx = cm.AddValue(ctx, "user", "alice")
	ctx = cm.AddValue(ctx, "requestID", "12345")

	// Use the context
	fmt.Println("Context created with values!")
}

func SimulateWork(ctx context.Context, workDuration time.Duration, description string) error {
	fmt.Println(description)

	select {
	case <-ctx.Done():
		return context.Canceled
	case <-time.After(workDuration):
		return nil
	}
}
func ProcessItems(ctx context.Context, items []string) ([]string, error) {
	final := make(chan []string)
	go func() {
		var proceed []string
		for _, item := range items {
			select {
			case <-ctx.Done():
				final <- proceed
				return
			case <-time.After(30 * time.Millisecond):
				proceed = append(proceed, fmt.Sprintf("processed_%s", item))
			}

		}
		final <- proceed
	}()
	result := <-final
	if ctx.Err() != nil {
		return result, ctx.Err()
	}
	return result, nil
}

type ContextManager interface {
	CreateCancellableContext(parent context.Context) (context.Context, context.CancelFunc)
	CreateTimeoutContext(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
	AddValue(parent context.Context, key, value any) context.Context
	GetValue(parent context.Context, key any) (any, bool)
	ExecuteWithContext(ctx context.Context, op func() error) error
	WaitForCompletion(ctx context.Context, duration time.Duration) error
}

type simpleContextManager struct{}

func NewContextManager() ContextManager {
	return &simpleContextManager{}
}

func (s *simpleContextManager) CreateCancellableContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(parent)
}

func (s *simpleContextManager) CreateTimeoutContext(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

func (s *simpleContextManager) AddValue(parent context.Context, key, value any) context.Context {
	return context.WithValue(parent, key, value)
}

func (s *simpleContextManager) GetValue(parent context.Context, key any) (any, bool) {
	val := parent.Value(key)
	if val == nil {
		return nil, false
	}
	return val, true
}

func (s *simpleContextManager) ExecuteWithContext(ctx context.Context, task func() error) error {
	taskChannel := make(chan error)
	go func() {
		err := task()
		taskChannel <- err
	}()
	select {
	case <-ctx.Done():
		return context.Canceled
	case <-time.After(100 * time.Millisecond):
		return context.DeadlineExceeded
	case err := <-taskChannel:
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *simpleContextManager) WaitForCompletion(ctx context.Context, duration time.Duration) error {

	select {
	case <-ctx.Done():
		return context.Canceled
	case <-time.After(duration):
		return nil
	}

}
