package db

import (
	"fmt"
)

type kvs map[string]string

// DbrLogger handles the logging for DBR on PG
var DbrLogger = &PgEventReceiver{}

// PgEventReceiver is a sentinel EventReceiver.
// Use it if the caller doesn't supply one.
type PgEventReceiver struct{}

// Event receives a simple notification when various events occur.
func (n *PgEventReceiver) Event(eventName string) {
	fmt.Println("DBR Event:", eventName)
}

// EventKv receives a notification when various events occur along with
// optional key/value data.
func (n *PgEventReceiver) EventKv(eventName string, kvs map[string]string) {
	fmt.Println("DBR EventKv:", eventName, kvs)
}

// EventErr receives a notification of an error if one occurs.
func (n *PgEventReceiver) EventErr(eventName string, err error) error {
	fmt.Println("DBR EventErr:", eventName, err)
	return err
}

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data.
func (n *PgEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	fmt.Println("DBR EventErrKv:", eventName, err, kvs)
	return err
}

// Timing receives the time an event took to happen.
func (n *PgEventReceiver) Timing(eventName string, nanoseconds int64) {
	fmt.Println("DBR Timing:", eventName, nanoseconds)
}

// TimingKv receives the time an event took to happen along with optional key/value data.
func (n *PgEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	fmt.Println("DBR TimingKv:", eventName, nanoseconds, kvs)
}
