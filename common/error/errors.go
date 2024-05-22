package errors

import "github.com/rotisserie/eris"

func New(msg string) error {
	return eris.New(msg)
}

func Wrap(err error, msg string) error {
	return eris.Wrap(err, msg)
}

func ToString(err error) string {
	format := eris.NewDefaultStringFormat(eris.FormatOptions{
		InvertOutput: true,
		WithTrace:    true,
		InvertTrace:  true,
	})
	return eris.ToCustomString(err, format)
}
