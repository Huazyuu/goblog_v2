package diverseType

import "encoding/json"

type ImageType int

const (
	Local  ImageType = 1 // 本地
	QiNiu  ImageType = 2 // 七牛云
	Remote ImageType = 3
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	case Remote:
		str = "远程图床"
	default:
		str = "本地"
	}
	return str
}
