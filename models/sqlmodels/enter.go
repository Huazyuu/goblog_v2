package sqlmodels

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `gorm:"comment:创建时间" json:"created_at,select($any)" structs:"-"`  // 创建时间
	UpdatedAt time.Time `gorm:"comment:更新时间" json:"-" structs:"-"`                        // 更新时间
}

// structs:"-" 中的 structs 是结构体标签的键，
// 它表明这个标签是针对 go-structs 库使用的；- 是标签的值，
// 它是一个特殊的约定，表示在使用 go-structs 库进行相关操作时忽略该字段。
