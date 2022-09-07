<!-- GETTING STARTED -->
## Getting Started

Solution for [Explorex problem](problem.txt)
Written in Go 1.19

### Prerequisites

Install Go from https://go.dev/doc/install



<!-- USAGE EXAMPLES -->
## Usage

Open terminal in the project directory and do:
```sh
  go run main.go <input_filename>
  ```

If no filename is provided, [logs_problem.txt](logs_problem.txt) is used as input.

### Input file logs template

LOG_TYPE DD-MM-YYYY HH:MM:SS  log_message

### Output format

sending <notification_channel> to users: [{user_name} ...] Logs: [LOG_TYPE DD-MM-YYYY HH:MM:SS  log_message ...]



<!-- ARCHITECTURE -->
## Code Design

### Log Type

This is the type of log, for example: WARN, ERROR etc.
It is represented as a struct:
```go
type LogType struct {
	LogtypeName          string
	Threshold            int
	WaitTime             string
	MeasurmentWindow     string
	NotificationChannels []string
	NotifiableUsers      []User
}
```
The file [log_types.go](log_types/log_types.go) defines this struct, helper methods and also defines different logtypes. 
To define a log type one can add to this file for example, a warn type can be:
```go
var WarnLogType LogType = LogType{
	LogtypeName:          "WARN",
	Threshold:            5,
	WaitTime:             "5m",
	MeasurmentWindow:     "1m",
	NotificationChannels: []string{"EMAIL"}, // notification channel can be an interface
	NotifiableUsers: []u.User{{
		Name:    "user 1",
		EmailId: "",
		PnId:    "",
		SmsId:   "",
	}, {
		Name:    "user 2",
		EmailId: "user@email.co",
		PnId:    "",
		SmsId:   "",
	}}, // user can also be a struct
}
```


### Log instance

Every log that is parsed and an object of LogInstance struct is created:
```go
type LogInstance struct {
    LogType   *t.LogType
    Timestamp int64
    Message   string
}

func StringToLog(logString string) (*LogInstance, error) {...}
```


### Alert sending logic

Alert sending logic is handled in this file. 

All logs that need to be sent are stored in objects of a struct LogsData defined in [alert_sender.go](alert_sender.go)
```go
type LogsData struct {
    sendNextLogAfter int64
	endtime          int64
	configWidowQueue []l.LogInstance
}
```
It uses configWidowQueue to store logs that can be valid to send, depending of log type's configuration. For every LogType, a LogsData object should be created like for above WarnLogType, the log data would look like:
```go
var WarnLogsData LogsData = LogsData{
    sendNextLogAfter: -1,
    endtime:          -1,
    configWidowQueue: []l.LogInstance{},
}
```
When a warning log is encountered, it should be added to the window if:
1. It is within the window time frame. Window timeframe's end is defined as ```endTime```.
2. If it's timestamp is after ```sendNextLogAfter```. This is so that the system has to wait for WaitTime to get over.

To add to LogsData, use 
```go
  LogsData.Scan(og *l.LogInstance)
```

Once the size of the window reaches threshold, logs are sent using 
```go
  func (LogsData *LogsData) send(log *l.LogInstance, updateEndTime bool) {...}
```
This function also handles cleanup.

### Notification Channel
NotificationChannel is an interface to provide common place for preprocessing before sending message
```go
type NotificationChannel interface {
	SendNotification()
}
```

Different notification channels are structs implementing NotificationChannel. Example:
```go
type SMS struct {
	Logs  []l.LogInstance
	Users []u.User
}

func (s SMS) SendNotification() {...}
```

Logs are added to resepective channel and sent over the channel using 
```go
func SendAlertOnChannel(channel string, logs []l.LogInstance, users []u.User) {...}
```

## Design patterns used

Creational pattern (Factory Method is StringToLog) is used to create instances of logs and store them in their respective windows LogsData. LogsData then handles the sending logic. Every notification channel is basicaly an implementation of NotificationChannel interface and should abide by it's signature. It is ensured by SendLogs(nc NotificationChannel) function. 

## Improvements

- [ ] Logic to automate creation of LogsData of different LogTypes. This doesn't need to be hardcoded.
- [ ] Mechanism to create a default log type when an unknown log type is encountered in the logs input.
- [ ] Code generation tool that takes a config file and adds LogTypes to log_types.go 

