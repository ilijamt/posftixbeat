package beater

import (
	"fmt"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/elastic/beats/filebeat/input"
	"github.com/elastic/beats/filebeat/publish"
	"github.com/elastic/beats/filebeat/registrar"

	cfg "github.com/ilijamt/postfixbeat/config"
	"github.com/ilijamt/postfixbeat/crawler"
	"github.com/ilijamt/postfixbeat/spooler"
)

type Postfixbeat struct {
	config *cfg.Config
	done   chan struct{}
}

// Creates beater
func New() *Postfixbeat {
	return &Postfixbeat{
		done: make(chan struct{}),
	}
}

func (pb *Postfixbeat) Config(b *beat.Beat) error {

	// Load beater config
	err := b.RawConfig.Unpack(&pb.config)
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	// Check if optional config_dir is set to fetch additional prospector config files
	pb.config.FetchConfigs()

	return nil
}

func (pb *Postfixbeat) Setup(b *beat.Beat) error {
	return nil
}

func (pb *Postfixbeat) Run(b *beat.Beat) error {
	logp.Info("postfixbeat is running! Hit CTRL-C to stop it.")

	var err error
	var config cfg.PostfixbeatConfig = pb.config.Postfixbeat

	// Setup registrar to persist state
	registrar, err := registrar.New(config.RegistryFile)
	if err != nil {
		logp.Err("Could not init registrar: %v", err)
		return err
	}

	// Channel from harvesters to spooler
	publisherChan := make(chan []*input.FileEvent, 1)

	// Publishes event to output
	publisher := publish.New(config.PublishAsync,
		publisherChan, registrar.Channel, b.Publisher.Connect())

	// Init and Start spooler: Harvesters dump events into the spooler.
	spooler, err := spooler.New(config, publisherChan)
	if err != nil {
		logp.Err("Could not init spooler: %v", err)
		return err
	}

	crawler, err := crawler.New(spooler, config.Prospectors)
	if err != nil {
		logp.Err("Could not init crawler: %v", err)
		return err
	}

	// The order of starting and stopping is important. Stopping is inverted to the starting order.
	// The current order is: registrar, publisher, spooler, crawler
	// That means, crawler is stopped first.

	// Start the registrar
	err = registrar.Start()
	if err != nil {
		logp.Err("Could not start registrar: %v", err)
	}
	// Stopping registrar will write last state
	defer registrar.Stop()

	// Start publisher
	publisher.Start()
	// Stopping publisher (might potentially drop items)
	defer publisher.Stop()

	// Starting spooler
	spooler.Start()
	// Stopping spooler will flush items
	defer spooler.Stop()

	err = crawler.Start(registrar.GetStates())
	if err != nil {
		return err
	}
	// Stop crawler -> stop prospectors -> stop harvesters
	defer crawler.Stop()

	// Blocks progressing. As soon as channel is closed, all defer statements come into play
	<-pb.done

	return nil
}

func (bt *Postfixbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Postfixbeat) Stop() {
	logp.Info("Stopping postfixbeat")
	close(bt.done)
}
