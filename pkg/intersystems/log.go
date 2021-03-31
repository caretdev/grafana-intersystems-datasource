package intersystems

import (
	"context"
	"regexp"
	"strconv"
	"time"

	"github.com/caretdev/go-irisnative/src/connection"
	"github.com/caretdev/go-irisnative/src/iris"

	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var levels = []string{
	"info",
	"alert",
	"warn",
	"err",
}

type LogLine struct {
	Time  string
	Pid  string
	Level string
	Source  string
	Text  string
}

type LogLines []LogLine

func parseTime(date string) (time.Time, error) {
	return time.Parse("01/02/06-15:04:05.000", regexp.MustCompile(`:(\d+)$`).ReplaceAllString(date, ".$1"))
}

func (l LogLines) Frames() data.Frames {
	frame := data.NewFrame("", 
		data.NewField("time", nil, []time.Time{}),
		data.NewField("message", nil, []string{}),
		data.NewField("source", nil, []string{}),
		data.NewField("pid", nil, []string{}),
		data.NewField("level", nil, []string{}),
	).SetMeta(&data.FrameMeta{
		PreferredVisualization: "logs",
	})

	for _, ll := range l {
		t, _ := parseTime(ll.Time)
		frame.AppendRow(
			t,
			ll.Text,
			ll.Source,
			ll.Pid,
			ll.Level,
		)
	}

	return data.Frames{frame}
}

func GetLog(ctx context.Context, client InterSystems, opts models.ListLogOptions) (LogLines, error) {
	var (
		log = make([]LogLine, 0)
	)
	if opts.File == "" {
		return log, nil
	}
	var conn connection.Connection
	var err error
	if conn, err = client.ConnectSYS(); err != nil {
		return log, err
	}
	defer conn.Disconnect()

	re := regexp.MustCompile(`([\d\/\-:]+) \((\d+)\) (\d) \[([^\]]+)\] (.*)`)
	var fs iris.Oref
	conn.ClassMethod("%Stream.FileCharacter", "%New", &fs)
	var status string
	logFile := "/usr/irissys/mgr/" + opts.File
	conn.Method(fs, "LinkToFile", &status, logFile)
	var atEnd bool
	conn.PropertyGet(fs, "AtEnd", &atEnd)
	for ; !atEnd; conn.PropertyGet(fs, "AtEnd", &atEnd) {
		var line string
		conn.Method(fs, "ReadLine", &line)
		vals := re.FindStringSubmatch(line)
		if vals == nil {
			continue
		}
		time, pid, level, source, text := vals[1], vals[2], vals[3], vals[4], vals[5]
		levelInt, _ := strconv.Atoi(level)
		logLine := LogLine{time, pid, levels[levelInt], source, text}
		log = append([]LogLine{logLine}, log...)
	}

	return log, nil
}
