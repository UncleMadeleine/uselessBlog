package tools

import (
	"io"
	"log"
	"os"
)

//ReadFile 读取文件内容的工具类
func ReadFile(filePath string) ([]byte, bool) {
	file, err := os.Open(filePath)
	var bytes []byte
	if err != nil {
		log.Print(err)
		return bytes, false
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		// 返回本次读取的字节数
		count, err := file.Read(buf)
		// 检测是否到了文件末尾
		if err == io.EOF {
			break
		} else if err != nil {
			log.Print(err)
			return bytes, false
		}

		// 取出本次读取的数据
		currBytes := buf[:count]

		// 将读取到的数据 追加到字节切片中
		bytes = append(bytes, currBytes...)
	}

	// 将字节切片转为字符串 最后打印出来文件内容
	// log.Print("文件内容：" + string(bytes))
	return bytes, true
}
