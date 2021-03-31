package intersystems

import (
	"context"
	"time"

	"github.com/caretdev/go-irisnative/src/connection"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type ApplicationError struct {
	Time     string
	Message  string
	PID      string
	Username string
}

type ApplicationErrors []ApplicationError

func (l ApplicationErrors) Frames() data.Frames {
	frame := data.NewFrame("",
		data.NewField("Time", nil, []time.Time{}),
		data.NewField("Message", nil, []string{}),
		data.NewField("Process ID", nil, []string{}),
		data.NewField("Username", nil, []string{}),
		data.NewField("level", nil, []string{}),
	).SetMeta(&data.FrameMeta{
		PreferredVisualization: "logs",
	})

	for _, ll := range l {
		t, _ := FromHorolog(ll.Time)
		frame.AppendRow(
			t,
			ll.Message,
			ll.PID,
			ll.Username,
			"err",
		)
	}

	return data.Frames{frame}
}

func GetApplicationErrors(ctx context.Context, client InterSystems, opts models.ListApplicationErrorsOptions) (ApplicationErrors, error) {
	var (
		errors = make([]ApplicationError, 0)
	)

	var conn connection.Connection
	var err error
	if conn, err = client.Connect(); err != nil {
		return errors, err
	}

	for date := ""; ; {
		if hasNext, _ := conn.GlobalPrev("ERRORS", &date); !hasNext {
			break
		}
		for i := ""; ; {
			if hasNext, _ := conn.GlobalPrev("ERRORS", &i, date); !hasNext {
				break
			}

			var ts string
			conn.GlobalGet("ERRORS", &ts, date, i, "*STACK", 0, "V", "$H")

			var message string
			conn.GlobalGet("ERRORS", &message, date, i, "*STACK", 0, "V", "Error")

			var pid string
			conn.GlobalGet("ERRORS", &pid, date, i, "*STACK", 0, "V", "$J")

			var user string
			conn.GlobalGet("ERRORS", &user, date, i, "*STACK", 0, "V", "$USERNAME")

			errors = append(errors, ApplicationError{
				ts,
				message,
				pid,
				user,
			})
		}
	}

	return errors, nil
}
