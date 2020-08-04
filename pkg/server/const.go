package server

import "time"

const (
	shutdownWaitTimeout = 3 * time.Second
)

const (
	panicPrintStackSize = 4 << 10
	panicPrintStackAll  = true
)
