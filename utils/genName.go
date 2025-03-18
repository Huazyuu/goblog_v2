package utils

import (
	"math/rand"
)

type randomType int

const (
	AttributeCharacter randomType = iota // 属性+角色 例如：傲娇的夜刀神十香
	CharacterAction                      // 角色+动作 例如：御坂美琴释放超电磁炮
)

func GenerateName() string {
	switch randomType(rand.Intn(2)) {
	case AttributeCharacter:
		return AnimeAttribute[rand.Intn(AnimeAttributeCount)] + AnimeCharacter[rand.Intn(AnimeCharacterCount)]
	case CharacterAction:
		return AnimeCharacter[rand.Intn(AnimeCharacterCount)] + AnimeAction[rand.Intn(AnimeActionCount)]
	default:
		return AnimeCharacter[rand.Intn(AnimeCharacterCount)] + AnimeAction[rand.Intn(AnimeActionCount)]
	}
}

var AnimeAttribute = []string{
	"傲娇的", "三无的", "病娇的", "天然呆的", "中二的", "黑化的", "毒舌的", "电波的", "冒失的", "腹黑的",
	"元气的", "高冷的", "废柴的", "兽耳的", "异色瞳的", "魔法少女的", "机甲驾驶的", "吸血鬼的", "天使的", "恶魔的",
	"女仆装的", "和风系的", "赛博朋克的", "异世界转生的", "灵力觉醒的", "契约召唤的", "超能力者的", "时间穿越的",
	"人造人形の", "魔王转世の", "圣剑继承の", "龙族血统の", "幽灵附体の", "妖狐化身の", "异能觉醒の",
}
var AnimeAttributeCount = len(AnimeAttribute)

var AnimeAction = []string{
	"展开绝对领域", "发动无限剑制", "释放星爆气流斩", "召唤使魔", "开启写轮眼", "吟唱禁忌咒文",
	"进入zone状态", "展开AT力场", "使用二刀流", "发动Geass", "展开替身攻击", "释放超电磁炮",
	"开启赫眼", "发动炼金术", "召唤Persona", "进行精灵同化", "启动零时迷子", "展开量子接续",
	"使用王之财宝", "发动固有结界", "开启轮回眼", "召唤神威灵装", "进行英灵召唤", "启动暴走模式",
	"展开水晶结界", "发动二重存在", "进入里世界", "使用言灵之力", "觉醒魔眼", "启动λ-Driver",
}
var AnimeActionCount = len(AnimeAction)

var AnimeCharacter = []string{
	"夜刀神十香", "御坂美琴", "立华奏", "夏娜", "凉宫春日", "晓美焰",
	"五河琴里", "时崎狂三", "雪之下雪乃", "椎名真白", "加藤惠",
	"绫波丽", "明日香", "雷姆", "拉姆", "爱蜜莉雅", "珂朵莉",
	"空银子", "四宫辉夜", "藤原千花", "战场原黑仪", "牧濑红莉栖",
	"阿尔托莉雅", "远坂凛", "间桐樱", "两仪式", "卫宫士郎",
	"鲁路修", "金木研", "艾伦·耶格尔", "利威尔", "炭治郎",
	"宇智波佐助", "漩涡鸣人", "坂田银时", "阿虚", "上条当麻",
	"一方通行", "折原临也", "夜神月", "L·Lawliet", "比企谷八幡",
	"桐谷和人", "司波达也", "琦玉老师", "五条悟", "虎杖悠仁",
	"神楽", "鹿目圆", "晓美焰", "巴麻美", "佐仓杏子",
	"立花泷", "宫水三叶", "千反田爱瑠", "折木奉太郎",
	"雪ノ下雪乃", "由比滨结衣", "一色彩羽",
	"初音未来", "镜音铃", "巡音流歌", "洛天依",
	"时崎狂三", "鸢一折纸", "诱宵美九",
	"霞之丘诗羽", "英梨梨", "安艺伦也",
	"夜斗", "雪音", "壹岐日和",
	"夏目贵志", "猫咪老师",
	"富樫勇太", "小鸟游六花",
	"高木同学", "西片",
	"千反田える", "折木奉太郎",
}
var AnimeCharacterCount = len(AnimeCharacter)
