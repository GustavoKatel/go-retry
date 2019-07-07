package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestErrInit(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	assert.NotNil(Retry(
		ctx,
		nil,
		Init(func(ctx context.Context) (interface{}, error) {
			return nil, fmt.Errorf("Test")
		}),
	))
}

func TestErrInitSecondCall(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	started := false

	assert.NotNil(Retry(
		ctx,
		func(ctx context.Context, _ interface{}) error {
			return fmt.Errorf("Test")
		},
		Init(func(ctx context.Context) (interface{}, error) {
			if !started {
				started = true
				return nil, nil
			}
			return nil, fmt.Errorf("Test")
		}),
	))
}

func TestMaxAttemptsReach(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	var count int

	assert.NotNil(Retry(
		ctx,
		func(ctx context.Context, _ interface{}) error {
			count++
			return fmt.Errorf("Test")
		},
		MaxAttempts(2),
	))

	assert.Equal(2, count)
}

func TestErrCleanup(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	assert.NotNil(Retry(
		ctx,
		func(_ context.Context, _ interface{}) error {
			return nil
		},
		Clean(func(ctx context.Context, _ interface{}, _ error) error {
			return fmt.Errorf("Test")
		}),
	))
}

func TestCleanupCatchActorErr(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	err := errors.New("Test")

	assert.NotNil(Retry(
		ctx,
		func(_ context.Context, _ interface{}) error {
			return err
		},
		Clean(func(ctx context.Context, _ interface{}, e error) error {
			assert.Equal(err, e)
			return e
		}),
	))
}

func TestCtxCancel(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	count := 0

	assert.NotNil(Retry(
		ctx,
		func(ctx context.Context, _ interface{}) error {
			count++
			<-time.After(100 * time.Millisecond)
			return fmt.Errorf("Test")
		},
	))

	assert.Equal(10, count)
}

func TestValue(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	value := 5

	assert.Nil(Retry(
		ctx,
		func(_ context.Context, v interface{}) error {
			if &value != v {
				return fmt.Errorf("Value diff")
			}
			return nil
		},
		Init(func(ctx context.Context) (interface{}, error) {
			return &value, nil
		}),
		Clean(func(ctx context.Context, v interface{}, e error) error {
			assert.Nil(e)
			if &value != v {
				return fmt.Errorf("Value diff")
			}
			return e
		}),
	))
}
