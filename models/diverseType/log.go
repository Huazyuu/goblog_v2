package diverseType

type LogLevel int

const (
	Info    LogLevel = 1
	Warning LogLevel = 2
	Error   LogLevel = 3
)

func (level LogLevel) String() string {
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

// func (level LogLevel) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(level.String())
// }

type LogType int

const (
	LoginType   LogType = 1
	ActionType  LogType = 2
	RuntimeType LogType = 3
)

func (t LogType) String() string {
	switch t {
	case LoginType:
		return "loginType"
	case ActionType:
		return "actionType"
	case RuntimeType:
		return "runtimeType"
	}
	return ""
}

// func (t LogType) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(t.String())
// }
