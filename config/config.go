package config

import "errors"

const (
    KeyToken      string = "token"
    KeyDataSource string = "datasource"
    KeyHost       string = "host"

    defaultHost = "https://api.tinybird.co"
)

var (
    ErrEmptyConfig = errors.New("missing or empty config")
)

type Config struct {
    // Connection string connection to Tinybird

    // Tinybird admin token
    TinybirdToken string `validate:"required"`

    // Host, default to https://api.tinybird.co
    Host string

    // Landing data source name
    Datasource string `validate:"required"`
}

func Parse(cfg map[string]string) (Config, error) {
    // Raise error in passed config is empty
    err := checkEmpty(cfg)
    if err != nil {
        return Config{}, err
    }

    config := Config{
        TinybirdToken: cfg[KeyToken],
        Datasource:    cfg[KeyDataSource],
        Host:          defaultHost,
    }

    // Assign default host
    if cfg[KeyHost] != "" {
        config.Host = defaultHost
    }

    return config, nil
}

func checkEmpty(cfg map[string]string) error {
    if len(cfg) == 0 {
        return ErrEmptyConfig
    }
    return nil
}
