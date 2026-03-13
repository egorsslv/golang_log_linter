package loggers

import (
	"context"
	"log/slog"
	"os"
)

func log() {
	// Test Upper case:
	slog.Debug("Test") // want "first symbol of log msg shouldn't be in upper case"

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Test english case:
	l.Info("тест на английский язык") // want "log msg must be in english only"

	// Test specifical symbols and emojis case
	l.DebugContext(context.Background(), "zasoV1!2_-") // want "log msg shouldn't have specifical symbols and emojis"

	// Test sensitive data case
	const apiKey = "dpsdpqrwoiwqir2o1r21rpsfod"
	l.WarnContext(context.Background(), "check apikey"+apiKey) // want "args shouldn't be had sensitive data: apiKey"

	//Test linter checking args case
	l.Error("some msg", "тест") // want "log msg must be in english only"

	l.Error("some msg", "creds") // // want "args shouldn't be had sensitive data: creds"

	l.InfoContext(context.Background(), "hi", "user", "😅") // want "log msg shouldn't have specifical symbols and emojis"

}
