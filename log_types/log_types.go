package logtypes

import "time"

// LogType is a struct and not an interface to allow creation of LogTypes on the fly if required.
type LogType struct {
	LogtypeName          string
	Threshold            int
	WaitTime             string
	MeasurmentWindow     string
	NotificationChannels []string
	NotifiableUsers      []string
}

func (logType *LogType) MeasurementWindowInSeconds() int64 {
	// logType.MeasurmentWindow
	t, _ := time.ParseDuration(logType.MeasurmentWindow)
	return int64(t.Seconds())
}

func (logType *LogType) WaitTimeInSeconds() int64 {
	t, _ := time.ParseDuration(logType.MeasurmentWindow)
	return int64(t.Seconds())
}

var WarnLogType LogType = LogType{
	LogtypeName:          "WARN",
	Threshold:            5,
	WaitTime:             "5m",
	MeasurmentWindow:     "1m",
	NotificationChannels: []string{"EMAIL"},          // notification channel can be an interface
	NotifiableUsers:      []string{"user1", "user2"}, // user can also be a struct
}

var ErrorLogType LogType = LogType{
	LogtypeName:          "ERROR",
	Threshold:            2,
	WaitTime:             "2m",
	MeasurmentWindow:     "30s",
	NotificationChannels: []string{"PN", "SMS"},      
	NotifiableUsers:      []string{"user3", "user4"}, 
}
