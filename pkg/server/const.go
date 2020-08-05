package server

import "time"

const (
	httpPort            = ":8443"
	grpcPort            = ":8444"
	shutdownWaitTimeout = 3 * time.Second
)

const (
	panicPrintStackSize = 4 << 10
	panicPrintStackAll  = true
)
