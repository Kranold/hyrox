package logging

import (
	"log/slog"
	"os"
	"runtime"
)

func CreateLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	pc, file, line, _ := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	logger = logger.With(
		slog.String("func", funcname),
		slog.String("file", file),
		slog.Int("line", line),
	)

	return logger
}
