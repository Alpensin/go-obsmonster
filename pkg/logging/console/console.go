package console

const LoggerNameField = "🚒 logger"

// Logger - common interface that our libraries must meet.
// Only JSON formatted message to stdout expected from them
type Logger interface {
	Debug(msg string, args ...Arg)
	Info(msg string, args ...Arg)
	Warn(msg string, args ...Arg)
	Error(msg string, args ...Arg)
}

// Arg - expected values. Different loggers provide own functions to log specific types
// It makes harder to unify interface between them. So we make own absraction to specify arguments we logging
type Arg struct {
	Key   string
	Value any
}

// NewArg - create NewArg to log data
func NewArg(key string, value any) Arg {
	return Arg{
		Key:   key,
		Value: value,
	}
}
