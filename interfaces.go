package retry

import "context"

// Initiator init func to be called once before every attempt
type Initiator func(ctx context.Context) (interface{}, error)

// Actor func to be called on every attempt
type Actor func(ctx context.Context, value interface{}) error

// Cleaner func to be called once after every attempt
// lastActorErr contains the last error emited by Actor, if any. Nil otherwise
type Cleaner func(ctx context.Context, value interface{}, lastActorErr error) error
