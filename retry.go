package retry

import "context"

// Retry executes "actor" until it returns nil
func Retry(ctx context.Context, actor Actor, opts ...Option) error {
	var options Options
	options.SetDefaults()

	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return err
		}
	}

	attempts := options.MaxAttempts

	var lastErr error

	init := options.Init
	cleanup := options.Clean

	v, err := init(ctx)
	if err != nil {
		return err
	}

	for {
		if ctx.Err() != nil {
			break
		}

		if options.MaxAttempts > 0 && attempts <= 0 {
			lastErr = ErrMaxAttemptsReached
			break
		}

		attempts--
		actorErr := actor(ctx, v)

		if actorErr != nil {
			if err := cleanup(ctx, v, actorErr); err != nil {
				return err
			}
		} else {
			break
		}

		v, err = init(ctx)
		if err != nil {
			return err
		}
	}

	if err := cleanup(ctx, v, nil); err != nil {
		return err
	}

	if lastErr != nil {
		return lastErr
	}

	return ctx.Err()
}
