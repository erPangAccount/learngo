package rpcdemo

import "github.com/pkg/errors"

type DemoServer struct {
}

type Args struct {
	A, B int
}

func (DemoServer) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("div err")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
