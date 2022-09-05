package main

import (
	"bufio"
	"log"
	s "log-alerting-system/alert_sender"
	l "log-alerting-system/log_instance"
	t "log-alerting-system/log_types"
	"os"
)

func main() {

	file, err := os.Open("logs_timestamp.txt") // take this from command line
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner capacity is 64k characters
	for scanner.Scan() {
		logString := scanner.Text()
		log, err := l.StringToLog(logString)

		if err == nil {
			switch log.LogType {
			case &t.ErrorLogType:
				s.ErrorLogData.Scan(log)
				break
			case &t.WarnLogType:
				s.WarnLogData.Scan(log)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
