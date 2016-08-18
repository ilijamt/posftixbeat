package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/ilijamt/postfixbeat/config"
)

type Postfixbeat struct {
	beatConfig *config.Config
	done       chan struct{}
	period     time.Duration
	client     publisher.Client
}

// Creates beater
func New() *Postfixbeat {
	return &Postfixbeat{
		done: make(chan struct{}),
	}
}

/// *** Beater interface methods ***///

func (bt *Postfixbeat) Config(b *beat.Beat) error {

	// Load beater beatConfig
	err := b.RawConfig.Unpack(&bt.beatConfig)
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	return nil
}

func (bt *Postfixbeat) Setup(b *beat.Beat) error {

	return nil
}

func (bt *Postfixbeat) Run(b *beat.Beat) error {
	logp.Info("postfixbeat is running! Hit CTRL-C to stop it.")
	return nil
}

func (bt *Postfixbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Postfixbeat) Stop() {
	close(bt.done)
}
