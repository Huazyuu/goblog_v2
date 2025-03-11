package utils

import "fmt"

func CombineIDs(id1, id2 uint) string {
	if id1 < id2 {
		return fmt.Sprintf("%d%d", id1, id2)
	}
	return fmt.Sprintf("%d%d", id2, id1)
}
