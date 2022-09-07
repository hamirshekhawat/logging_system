// Different logtypes configurations are added here
package logtypes

import "time"

import u "log-alerting-system/user"

// LogType is a struct to allow creation of LogTypes on the fly if required.
type LogType struct {
	LogtypeName          string
	Threshold            int
	WaitTime             string
	MeasurmentWindow     string
	NotificationChannels []string
	NotifiableUsers      []u.User
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
	NotificationChannels: []string{"EMAIL"},
	NotifiableUsers: []u.User{{
		Name:    "user 1",
		EmailId: "",
		PnId:    "",
		SmsId:   "",
	}, {
		Name:    "user 2",
		EmailId: "",
		PnId:    "",
		SmsId:   "",
	}},
}

var ErrorLogType LogType = LogType{
	LogtypeName:          "ERROR",
	Threshold:            2,
	WaitTime:             "2m",
	MeasurmentWindow:     "30s",
	NotificationChannels: []string{"PN", "SMS"},
	NotifiableUsers: []u.User{{
		Name:    "user 3",
		EmailId: "",
		PnId:    "",
		SmsId:   "",
	}, {
		Name:    "user 4",
		EmailId: "",
		PnId:    "",
		SmsId:   "",
	}},
}
