package logStash

import "encoding/json"

type Level int

const (
	Info    Level = 1
	Warning Level = 2
	Error   Level = 3
)

func (level Level) String() string {
	switch level {
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	}
	return ""
}

func (level Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.String())
}
