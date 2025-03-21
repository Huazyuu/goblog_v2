package req

type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
}

type DateRequest struct {
	Date DateType `json:"date" form:"date"`
}

type DateType int

const (
	OneWeek = 1 + iota
	OneMonth
	TwoMonth
	ThreeMonth
	HalfYear
	OneYear
)

func (d DateType) String() string {
	var str string
	switch d {
	case OneWeek:
		str = "一周"
	case OneMonth:
		str = "一个月"
	case TwoMonth:
		str = "两个月"
	case ThreeMonth:
		str = "三个月"
	case HalfYear:
		str = "半年"
	case OneYear:
		str = "一年"
	}
	return str
}
