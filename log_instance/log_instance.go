// Provides an instance of log as a struct
package loginstance

import (
	"errors"
	t "log-alerting-system/log_types"
	"strings"
	"time"
)

type LogInstance struct {
	LogType   *t.LogType
	Timestamp int64
	Message   string
}

func StringToLog(logString string) (*LogInstance, error) {
	words := strings.Fields(logString)
	timestamp, tErr := time.Parse("02-01-2006 15:04:05", words[1]+" "+words[2]) // DD-MM-YYYY HH:MM:SS
	// epoch := timestamp.Unix()
	var logType *t.LogType
	var logInstance *LogInstance
	errorMessage := ""

	if tErr != nil {
		errorMessage += " Log timestamp invalid "
	}

	switch words[0] {
	case "WARN":
		logType = &t.WarnLogType
		break
	case "ERROR":
		logType = &t.ErrorLogType
		break
	default:
		logType = nil
		errorMessage += " Logtype not found "
	}

	if logType != nil && tErr == nil {
		logInstance = &LogInstance{
			Timestamp: timestamp.Unix(),
			Message:   words[3],
			LogType:   logType,
		}
		return logInstance, nil
	} else {
		return nil, errors.New(errorMessage)
	}
}

func (log *LogInstance) ToString() string {
	logTypeString := ""
	switch log.LogType {
	case &t.WarnLogType:
		logTypeString = "WARN"
		break
	case &t.ErrorLogType:
		logTypeString = "ERROR"
		break
	}
	timeStamString:= time.Unix(log.Timestamp, 0).Format("02-01-2006 15:04:05")
	
	return logTypeString + " " + timeStamString + " " + log.Message
}
