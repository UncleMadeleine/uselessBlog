package tools

import (
	"log"
	"strconv"
)

//StrToUInt string数据转换为uint
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		log.Print(e)
		return 0
	}
	return uint(i)
}

//StrToInt string 转int
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		log.Print(e)
		return 0
	}
	return i
}
