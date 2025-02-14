package controllers

import "time"

func MeasureExecutionTime(operation func()) time.Duration {
	start := time.Now()
	operation()
	return time.Since(start)
}
