package retry

import "context"

type Options struct {
	MaxAttempts int
	Init        Initiator
	Clean       Cleaner
}

type Option func(opts *Options) error

func DefaultInitiator(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func DefaultCleaner(ctx context.Context, _ interface{}, _ error) error {
	return nil
}

func (opts *Options) SetDefaults() {
	opts.MaxAttempts = -1

	opts.Init = DefaultInitiator
	opts.Clean = DefaultCleaner
}

func MaxAttempts(max int) Option {
	return func(opts *Options) error {
		opts.MaxAttempts = max
		return nil
	}
}

func Init(init Initiator) Option {
	return func(opts *Options) error {
		opts.Init = init
		return nil
	}
}

func Clean(clean Cleaner) Option {
	return func(opts *Options) error {
		opts.Clean = clean
		return nil
	}
}
