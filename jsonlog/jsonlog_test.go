package jsonlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) Sync() error { return nil }

func (b *syncBuffer) entries(t *testing.T) []map[string]any {
	t.Helper()
	b.mu.Lock()
	defer b.mu.Unlock()
	var entries []map[string]any
	for _, line := range strings.Split(strings.TrimSpace(b.buf.String()), "\n") {
		if line == "" {
			continue
		}
		entry := map[string]any{}
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			t.Fatalf("line is not JSON: %q: %v", line, err)
		}
		entries = append(entries, entry)
	}
	return entries
}

func TestNewWritesJSONWithFields(t *testing.T) {
	out := &syncBuffer{}
	logger := New(WithOutput(out), WithService("svc"))

	logger.Info("message received", zap.String("message_id", "42"), zap.Int("attempt", 2))
	logger.Error("request failed", zap.Error(errors.New("boom")))

	entries := out.entries(t)
	if len(entries) != 2 {
		t.Fatalf("got %d entries, want 2", len(entries))
	}
	info := entries[0]
	for key, want := range map[string]any{"level": "info", "msg": "message received", "service": "svc", "message_id": "42", "attempt": float64(2)} {
		if info[key] != want {
			t.Fatalf("%s = %v, want %v (entry %v)", key, info[key], want, info)
		}
	}
	if caller, _ := info["caller"].(string); !strings.Contains(caller, "jsonlog_test.go") {
		t.Fatalf("caller = %q, want the test file", caller)
	}
	if _, ok := info["time"].(string); !ok {
		t.Fatalf("time missing: %v", info)
	}
	if entries[1]["level"] != "error" || entries[1]["error"] != "boom" {
		t.Fatalf("error entry = %v", entries[1])
	}
	if _, ok := entries[1]["stacktrace"]; ok {
		t.Fatal("error entries should not carry a stack trace")
	}
}

func TestNewRespectsLevel(t *testing.T) {
	out := &syncBuffer{}
	logger := New(WithOutput(out), WithLevel(zapcore.WarnLevel))
	logger.Info("hidden")
	logger.Warn("shown")
	entries := out.entries(t)
	if len(entries) != 1 || entries[0]["msg"] != "shown" {
		t.Fatalf("entries = %v", entries)
	}
}

func TestNewIsSafeForConcurrentUse(t *testing.T) {
	out := &syncBuffer{}
	logger := New(WithOutput(out))
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger.Info("message", zap.Int("n", i))
		}(i)
	}
	wg.Wait()

	seen := map[float64]bool{}
	for _, entry := range out.entries(t) {
		seen[entry["n"].(float64)] = true
	}
	if len(seen) != 50 {
		t.Fatalf("got %d distinct entries, want 50", len(seen))
	}
}
