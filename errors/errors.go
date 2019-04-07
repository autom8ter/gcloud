package errors

import (
	"github.com/hashicorp/go-multierror"
)

func Append(base error, toAppend ...error) error {
	return multierror.Append(base, toAppend...)
}

