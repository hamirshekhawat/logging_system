package main

import (
	"bufio"
	"log"
	s "log-alerting-system/alert_sender"
	l "log-alerting-system/log_instance"
	t "log-alerting-system/log_types"
	"os"
	"fmt"
)

func main() {
	var filename string
	if(len(os.Args) > 1) {
		filename = os.Args[1]
	} else {
		fmt.Println("warning: Please provide the logs filename as argument.\nUsing test log file for demonstartion.")
		filename = "logs_problem.txt"
	}
	
	file, err := os.Open(filename) // take this from command line
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner capacity is 64k characters
	for scanner.Scan() {
		logString := scanner.Text()
		_log, err := l.StringToLog(logString)

		if err == nil {
			switch _log.LogType {
			case &t.ErrorLogType:
				s.ErrorLogData.Scan(_log)
				break
			case &t.WarnLogType:
				s.WarnLogData.Scan(_log)
				break
			}
		} else {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
