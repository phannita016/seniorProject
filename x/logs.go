package x

import "log/slog"

func Recover() {
	if err := recover(); err != nil {
		slog.Error("system-recover-running", slog.Any("msg", err))
	}
}
