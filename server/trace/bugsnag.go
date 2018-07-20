package trace

import (
	"os"

	bugsnag "github.com/bugsnag/bugsnag-go"
)

// NewBugsnag return bugsnag notifier
func NewBugsnag() *bugsnag.Notifier {
	errorHandlerConfig := bugsnag.Configuration{
		APIKey:          os.Getenv("BUGSNAG_KEY"),
		ProjectPackages: []string{"main", "github.com/Coderockr/vitrine-social/*"},
	}

	notifier := bugsnag.New(errorHandlerConfig)
	notifier.AutoNotify()

	return notifier
}
