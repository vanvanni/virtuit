package out

import (
	"context"
	"log/slog"
	"os"
	"runtime/debug"
	"time"
)

type TimeOnlyHandler struct {
	slog.Handler
}

func (h *TimeOnlyHandler) Handle(ctx context.Context, r slog.Record) error {
	r.Time = r.Time.Truncate(time.Second)
	return h.Handler.Handle(ctx, r)
}

var handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: false,
	Level:     slog.LevelDebug,
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			t := a.Value.Time()
			a.Value = slog.StringValue(t.Format("15:04:05"))
		}
		return a
	},
})

var buildInfo, _ = debug.ReadBuildInfo()
var logger = slog.New(handler)

var Logger = logger.With(
	slog.Group("program",
		slog.Int("pid", os.Getpid()),
		slog.String("go_version", buildInfo.GoVersion),
	),
)
