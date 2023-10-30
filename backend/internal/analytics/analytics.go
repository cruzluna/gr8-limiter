package analytics

import (
	"sync"

	"github.com/posthog/posthog-go"
)

var (
	Client posthog.Client
	once   sync.Once
)

func Init(posthogKey string) error {
	var err error
	once.Do(func() {
		var client posthog.Client
		client, err = posthog.NewWithConfig(posthogKey, posthog.Config{Endpoint: "https://app.posthog.com"})
		Client = client
	})
	return err
}
