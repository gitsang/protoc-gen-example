package main

import (
	"encoding/json"
	"io"
	"log/slog"
)

type JSONIndentHandlers struct {
	*slog.JSONHandler
}

func NewIdentJSONHandlers(w io.Writer, opts *slog.HandlerOptions) *JSONIndentHandlers {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &JSONIndentHandlers{
		JSONHandler: slog.NewJSONHandler(NewJSONIndentWriter(w), opts),
	}
}

type JSONIndentWriter struct {
	w io.Writer
}

func NewJSONIndentWriter(w io.Writer) *JSONIndentWriter {
	return &JSONIndentWriter{
		w: w,
	}
}

func (jiw *JSONIndentWriter) Write(p []byte) (n int, err error) {
	var alias map[string]any
	if err := json.Unmarshal(p, &alias); err != nil {
		return jiw.w.Write(p)
	}

	type LogStruct struct {
		Level  string `json:"level,omitempty"`
		Msg    string `json:"msg,omitempty"`
		Time   string `json:"time,omitempty"`
		Source any    `json:"source,omitempty"`
		Logs   any    `json:"logs,omitempty"`
	}
	logStruct := LogStruct{}
	logStruct.Level = alias["level"].(string)
	delete(alias, "level")
	logStruct.Msg = alias["msg"].(string)
	delete(alias, "msg")
	logStruct.Time = alias["time"].(string)
	delete(alias, "time")
	logStruct.Source = alias["source"]
	delete(alias, "source")
	logStruct.Logs = alias

	p, err = json.MarshalIndent(logStruct, "", "  ")
	if err != nil {
		return
	}
	n, err = jiw.w.Write(append(p, '\n'))
	if err != nil {
		return
	}
	return
}
