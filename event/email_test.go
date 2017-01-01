// +build !integration

package event

import (
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/stretchr/testify/assert"
)

func TestMailEventToMapStr(t *testing.T) {
	event := MailEvent{}
	mapStr := event.ToMapStr()
	_, found := mapStr["fields"]
	assert.False(t, found)
}

func TestMailEventToMapStrJSON(t *testing.T) {

	type io struct {
		Event         MailEvent
		ExpectedItems common.MapStr
	}

	now := time.Now()
	size := 100
	num := 0.14

	tests := []io{
		{
			Event: MailEvent{
				Timestamp:        now,
				QueueID:          "QUEUE_ID",
				MailFrom:         "test@test.com",
				Size:             &size,
				QueuedOn:         &now,
				Delay:            &num,
				UpdateOn:         &now,
				TimeBeforeQmgr:   &num,
				TimeInQmgr:       &num,
				TimeConnSetup:    &num,
				TimeTransmission: &num,
			},
			ExpectedItems: common.MapStr{
				"@timestamp":        common.Time(now),
				"queue_id":          "QUEUE_ID",
				"mail_from":         "test@test.com",
				"size":              100,
				"queued_on":         common.Time(now),
				"delay":             num,
				"time_before_qmgr":  num,
				"time_in_qmgr":      num,
				"time_conn_setup":   num,
				"time_transmission": num,
				"update_on":         common.Time(now),
			},
		},
	}

	for _, test := range tests {
		result := test.Event.ToMapStr()
		t.Log("Executing test:", test)
		for k, v := range test.ExpectedItems {
			assert.Equal(t, v, result[k])
		}
	}

}
