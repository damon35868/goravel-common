package common

import (
	"log/slog"
	"os"
)

/**
 * @description: JSON终端日志
 * @return {*}
 */
var Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
