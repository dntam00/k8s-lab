package main

import (
	"time"
)

func main() {
	conf := LogConfig{
		Path:       "/logs/application.log",
		Level:      "info",
		Formatter:  "json", // or "json"
		MaxSize:    10,
		MaxBackups: 3,
	}
	if err := GetLogger(conf); err != nil {
		panic(err)
	}

	for {
		InfofNw("Sample log entry at %s", time.Now().Format(time.RFC3339))
		time.Sleep(1 * time.Second)
	}
}
