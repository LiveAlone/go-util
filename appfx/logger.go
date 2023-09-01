package appfx

import (
	"go.uber.org/zap"
	"os"
)

var DEBUG = "-D"

func UtilsLogger() *zap.Logger {
	debug := false
	rmDebugArgs := make([]string, 0, len(os.Args))
	for _, arg := range os.Args {
		if arg == DEBUG {
			debug = true
			continue
		}
		rmDebugArgs = append(rmDebugArgs, arg)
	}
	os.Args = rmDebugArgs
	if debug {
		return zap.NewExample()
	} else {
		return zap.NewNop()
	}
}
