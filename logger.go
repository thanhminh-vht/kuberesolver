package kuberesolver

import (
	"log"
)

type LoggerPrintFuncType func(format string, args ...any)

var LoggerPrintFunc LoggerPrintFuncType = log.Printf
