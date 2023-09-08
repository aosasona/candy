package request

import "fmt"

var debug = false

func SetDebug(enable bool) {
	debug = enable
}

func EnableDebugMode() {
	debug = true
}

func DisableDebugMode() {
	debug = false
}

type debugOpts struct {
	URL     string
	Method  RequestMethod
	Headers map[string]string
	Body    any
}

func printDebug(opts debugOpts) {
	if debug {
		fmt.Printf(
			"\n\033[35mCANDY - DEBUG\033[0m\n\033[35mURL: %s\033[0m\nMethod: %s\nHeaders: %+v\nBody: %s\n",
			opts.URL,
			opts.Method,
			opts.Headers,
			opts.Body,
		)
	}
}
