package logger

import "fmt"

type UnknownLogLevel string

func (e UnknownLogLevel) Error() string {
	return fmt.Sprintf("unknown log level: %s", string(e))
}
