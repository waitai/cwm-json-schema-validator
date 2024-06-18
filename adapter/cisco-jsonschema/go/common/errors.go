package common

import (
	"go.temporal.io/sdk/temporal"
)

func NewError(msg string, ec ErrorCode) error {
	return temporal.NewApplicationError(msg, ec.String())
}
