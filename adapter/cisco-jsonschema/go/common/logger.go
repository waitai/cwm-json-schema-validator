package common

import (
	"context"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/log"
)

func GetLogger(ctx context.Context) log.Logger {
	return activity.GetLogger(ctx)
}
