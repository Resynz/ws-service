/**
 * @Author: Resynz
 * @Date: 2021/7/20 11:02
 */
package util

import (
	"encoding/hex"
	"strconv"
)

func ComputedNumberByString(str string) int {
	ss := hex.EncodeToString([]byte(str))
	var i int
	for _, v := range ss {
		ii, _ := strconv.Atoi(string(v))
		i += ii
	}
	return i
}
