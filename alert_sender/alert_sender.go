// Holds logic to send alert based on log type's configuration
package alertsender

import (
	l "log-alerting-system/log_instance"
	t "log-alerting-system/log_types"
	nc "log-alerting-system/notification_channels"

)

type LogData struct {
	sendNextLogAfter int64
	endtime          int64
	configWidowQueue []l.LogInstance // A queue like array to store logs that are going to be sent
}

var WarnLogData LogData = LogData{
	sendNextLogAfter: -1,
	endtime:          -1,
	configWidowQueue: []l.LogInstance{},
}

var ErrorLogData LogData = LogData{
	sendNextLogAfter: -1,
	endtime:          -1,
	configWidowQueue: []l.LogInstance{},
}

func (logData *LogData) Scan(log *l.LogInstance) {
	if len(logData.configWidowQueue) == 0 {
		logData.endtime = log.Timestamp + log.LogType.MeasurementWindowInSeconds()
	}

	if log.Timestamp >= logData.sendNextLogAfter {
		if logData.endtime >= log.Timestamp {
			logData.send(log, false)
		} else {
			// remove the ones outside window
			newStartTime := log.Timestamp - log.LogType.MeasurementWindowInSeconds()
			for _, l := range logData.configWidowQueue {
				if l.Timestamp < newStartTime {
					logData.configWidowQueue = logData.configWidowQueue[1:]
				}
			}
			logData.send(log, true)
		}
	}
}

// can be called async: go logData.send()
func (logData *LogData) send(log *l.LogInstance, updateEndTime bool) {
	logData.configWidowQueue = append(logData.configWidowQueue, *log)
	if len(logData.configWidowQueue) == log.LogType.Threshold {
		for _, c := range t.WarnLogType.NotificationChannels {
			nc.SendAlertOnChannel(c, logData.configWidowQueue, t.WarnLogType.NotifiableUsers)
		}
		if(updateEndTime) {
			logData.endtime = logData.configWidowQueue[0].Timestamp + log.LogType.MeasurementWindowInSeconds()
		}
		logData.configWidowQueue = []l.LogInstance{}
		logData.sendNextLogAfter = log.Timestamp + log.LogType.WaitTimeInSeconds()
	}
}
