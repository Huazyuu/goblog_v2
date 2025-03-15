package redisService

import (
	"backend/global"
	"strconv"
)

type RedisCount interface {
	Set(id string) error
	SetCount(id string, num int) error
	Get(id string) int
	GetAll() map[string]int
	Clear()
}
type CountDB struct {
	Index string // 索引
}

func (c CountDB) Set(id string) error {
	err := global.Redis.HSet(c.Index, id, c.Get(id)+1).Err()
	return err
}

func (c CountDB) SetCount(id string, num int) error {
	err := global.Redis.HSet(c.Index, id, c.Get(id)+num).Err()
	return err
}

func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	return num
}

func (c CountDB) GetAll() map[string]int {
	countInfo := make(map[string]int)
	res := global.Redis.HGetAll(c.Index)
	maps := res.Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		countInfo[id] = num
	}
	return countInfo
}

func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
