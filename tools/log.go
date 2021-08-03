package tools

import (
	"fmt"
	"log"
	"os"
)

//LogInit 初始化log
func LogInit() {
	logFile, logErr := os.OpenFile("./logfile.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
