package event

import (
	"time"

	"github.com/elastic/beats/libbeat/common"
)

type MailEvent struct {
	Timestamp time.Time
	QueueID   string

	// From Details
	MailFrom   string
	DomainFrom string
	Size       *int
	QueuedOn   *time.Time

	// To Details
	MailTo                     string
	DomainTo                   string
	Relay                      string
	Delay                      *float64
	TimeBeforeQmgr             *float64
	TimeInQmgr                 *float64
	TimeConnSetup              *float64
	TimeTransmission           *float64
	DeliveryStatusNotification string
	Status                     string
	StatusMessage              string
	UpdateOn                   *time.Time
}

func (me *MailEvent) ToMapStr() common.MapStr {
	event := common.MapStr{
		// generic
		"@timestamp": common.Time(me.Timestamp),
		"queue_id":   me.QueueID,

		// from
		"mail_from":   me.MailFrom,
		"domain_from": me.DomainFrom,

		// to
		"mail_to":        me.MailTo,
		"domain_to":      me.DomainTo,
		"relay":          me.Relay,
		"dsn":            me.DeliveryStatusNotification,
		"status":         me.Status,
		"status_message": me.StatusMessage,
	}

	for k := range event {
		if event[k] == nil || event[k] == "" {
			delete(event, k)
		}
	}

	if me.Size != nil {
		event["size"] = *me.Size
	}

	if me.QueuedOn != nil {
		event["queued_on"] = common.Time(*me.QueuedOn)
	}

	if me.Delay != nil {
		event["delay"] = *me.Delay
	}

	if me.TimeBeforeQmgr != nil {
		event["time_before_qmgr"] = *me.TimeBeforeQmgr
	}

	if me.TimeInQmgr != nil {
		event["time_in_qmgr"] = *me.TimeInQmgr
	}

	if me.TimeConnSetup != nil {
		event["time_conn_setup"] = *me.TimeConnSetup
	}

	if me.TimeTransmission != nil {
		event["time_transmission"] = *me.TimeTransmission
	}

	if me.UpdateOn != nil {
		event["update_on"] = common.Time(*me.UpdateOn)
	}

	return event
}
