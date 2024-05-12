package sl

import (
	"golang.org/x/exp/slog"
)

func ErrorLog(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

//func (sl slog.Logger) Fatal(msg string, err error) {
//	sl.Error(msg, ErrorLog(err))
//	os.Exit(1)
//}
