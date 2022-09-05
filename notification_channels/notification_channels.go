package notificationchannels

import (
	"fmt"
	l "log-alerting-system/log_instance"
	u "log-alerting-system/user"
)

// NotificationChannel is an interface to provide common place for preprocessing before sending message
type NotificationChannel interface {
	SendNotification()
}

// Different notification channels are struct with the same method signature as NotificationChannel
type SMS struct {
	Logs  []l.LogInstance
	Users []u.User
}

// Any message processing unique to this channel can be done here.
func (s SMS) SendNotification() {
	fmt.Print("sending sms to users: ")
	var logStrings []string
	for _, l := range s.Logs {
		logStrings = append(logStrings, l.ToString())
	}
	fmt.Print(s.Users)
	fmt.Print(" Logs:")
	fmt.Print(logStrings)
	fmt.Println()

}

type EMAIL struct {
	Logs  []l.LogInstance
	Users []u.User
}

func (s EMAIL) SendNotification() {
	fmt.Print("sending email to users: ")
	var logStrings []string
	for _, l := range s.Logs {
		logStrings = append(logStrings, l.ToString())
	}
	fmt.Print(s.Users)
	fmt.Print(" Logs:")
	fmt.Print(logStrings)
	fmt.Println()

}

type PN struct {
	Logs  []l.LogInstance
	Users []u.User
}

func (s PN) SendNotification() {
	fmt.Print("sending pn to users: ")
	var logStrings []string
	for _, l := range s.Logs {
		logStrings = append(logStrings, l.ToString())
	}
	fmt.Print(s.Users)
	fmt.Print(" Logs:")
	fmt.Print(logStrings)
	fmt.Println()

}

func SendLogs(nc NotificationChannel) {
	// Any common preprocessing can be done here before sending message
	nc.SendNotification()
}

func SendAlertOnChannel(channel string, logs []l.LogInstance, users []u.User) {
	switch channel {
	case "EMAIL":
		emailChannel := EMAIL{
			Logs:  logs,
			Users: users,
		}
		SendLogs(emailChannel)
		break
	case "SMS":
		smsChannel := SMS{
			Logs:  logs,
			Users: users,
		}
		SendLogs(smsChannel)
		break
	case "PN":
		pnChannel := PN{
			Logs:  logs,
			Users: users,
		}
		SendLogs(pnChannel)
		break
	}
}
