package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // logFile(受け取ったファイル名)に対して読み書き、作成、追記、パーミッションの設定をする
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)   // システムからの標準出力をlogfileに書き込むように設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // logfileの書式設定
	log.SetOutput(multiLogFile)
}
