package main

import (
	"fmt"
	"flag"
	"log/slog"
	"os"
	"io"
)

var importLogs = flag.String("imp", "", "path of the file to import logs from")

type multilogger struct {
	writers []io.Writer
}

func (m *multilogger) Write(p []byte) (int , error){
	var lastErr error;
	for _, logger := range m.writers {
		_, err := logger.Write(p)
		if err != nil {
			lastErr = err;
		}
	}
	return len(p), lastErr
}

func main(){
	flag.Parse()
	file, err := os.OpenFile("./temp/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error to logs file: ", err)
	}
	defer file.Close()
	ml := &multilogger{
		writers: []io.Writer{
			file,
			os.Stdout,
		},
	}
	handler := slog.NewJSONHandler(ml, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	logger.Info("Application started")
	
	if *importLogs != "" {
		logger.Info("Importing logs")
		f, err := os.ReadFile(*importLogs)
		if err != nil {
			logger.Error("Error Importing logs from: ", f)
		}
		file.Write(f)
		logger.Info("Logs Imported")
	}
	logger.Info("Application Exited")
}
