package main

import (
	"context"
	"time"
)

func main() {
	conf := LogConfig{
		Path:       "/logs/application.log",
		Level:      "info",
		Formatter:  "json",
		MaxSize:    10,
		MaxBackups: 3,
	}
	if err := GetLogger(conf); err != nil {
		panic(err)
	}

	for {
		Infof(context.Background(), "Sample log entry with user_id at %s", time.Now().Format(time.RFC3339))
		time.Sleep(1 * time.Second)
	}
}
