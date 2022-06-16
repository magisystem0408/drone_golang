package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//logfileを読み書きで追加できるようにした。
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logfile err=%s", err.Error())
	}
	//実行したあとログファイルに書き込む
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//日付時間。ショートファイル
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)

}
