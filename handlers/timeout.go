package handlers

import (
	"context"
	"time"
)


func GetTimeoutContext(parent context.Context, duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, duration)
}


func GetDefaultTimeoutContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, time.Second * 30)
}
