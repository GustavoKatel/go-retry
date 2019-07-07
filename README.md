# go-retry

[![GoDoc](https://godoc.org/github.com/GustavoKatel/go-retry?status.svg)](https://godoc.org/github.com/GustavoKatel/go-retry)

Retry operation executor with init and clean up functions

## Example

```golang
ctx := context.Background()

value := 5

Retry(
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
        if &value != v {
            return fmt.Errorf("Value diff")
        }
        return e
    }),
)
```