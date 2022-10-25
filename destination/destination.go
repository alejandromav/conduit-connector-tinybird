package destination

import (
    "context"
    "fmt"

    config "github.com/alejandromav/conduit-connector-tinybird/config"
    sdk "github.com/conduitio/conduit-connector-sdk"
)

type Destination struct {
    sdk.UnimplementedDestination

    config config.Config
}

func NewDestination() sdk.Destination {
    return &Destination{}
}

// Parameters is a map of named Parameters that describe how to configure
// the Destination.
func (d *Destination) Parameters() map[string]sdk.Parameter {
    return map[string]sdk.Parameter{
        config.KeyHost: {
            Description: "Tinybird host url",
        },
        config.KeyToken: {
            Required:    true,
            Description: "Tinybird admin token",
        },
        config.KeyDataSource: {
            Required:    true,
            Description: "Tinybird landing data source name",
        },
    }
}

// Configure is the first function to be called in a connector. It provides the
// connector with the configuration that needs to be validated and stored.
// In case the configuration is not valid it should return an error.
// Testing if your connector can reach the configured data source should be
// done in Open, not in Configure.
func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {
    sdk.Logger(ctx).Info().Msg("Configuring a Destination connector...")

    config, err := config.Parse(cfg)
    if err != nil {
        return err
    }

    d.config = config
    return nil
}

// Open is called after Configure to signal the plugin it can prepare to
// start writing records. If needed, the plugin should open connections in
// this function.
func (d *Destination) Open(ctx context.Context) error {
    sdk.Logger(ctx).Info().Msg("Connector open.")

    // TODO: Init tinybird client
    return nil
}

// Write writes len(r) records from r to the destination right away without
// caching. It should return the number of records written from r
// (0 <= n <= len(r)) and any error encountered that caused the write to
// stop early. Write must return a non-nil error if it returns n < len(r).
func (d *Destination) Write(ctx context.Context, records []sdk.Record) (int, error) {
    sdk.Logger(ctx).Info().Msg(fmt.Sprintf("Writing %d records...", len(records)))

    // TODO: Build ndjson payload

    // TODO: Send the bulk request

    // TODO: Handle response

    return len(records), nil
}

// Teardown signals to the plugin that all records were written and there
// will be no more calls to any other function. After Teardown returns, the
// plugin should be ready for a graceful shutdown.
func (d *Destination) Teardown(ctx context.Context) error {
    return nil // No close routine needed
}
