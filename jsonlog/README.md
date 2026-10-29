## Description

jsonlog builds a [zap](https://github.com/uber-go/zap) logger that writes one JSON object per line to stdout. The logger is safe to share between goroutines, reports the file and line of each log call, and never samples entries.

### Usage

```go
import (
	"github.com/devetek/go-core/jsonlog"
	"go.uber.org/zap"
)

func main() {
	logger := jsonlog.New(jsonlog.WithService("my-service"))
	defer logger.Sync()

	logger.Info("message received", zap.String("message_id", "42"))
}
```

Output:

```json
{"level":"info","time":"2026-10-29T20:00:00.123456789+07:00","caller":"app/main.go:12","msg":"message received","service":"my-service","message_id":"42"}
```

Options: `WithService(name)`, `WithLevel(zapcore.Level)` (default info), `WithOutput(zapcore.WriteSyncer)` (default stdout).
