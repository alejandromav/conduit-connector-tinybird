package tinybird

import (
    sdk "github.com/conduitio/conduit-connector-sdk"
)

// Specification returns the connector's specification.
func Specification() sdk.Specification {
    return sdk.Specification{
        Name:        "tinybird",
        Summary:     "Tinybird Conduit Connector",
        Description: "Tinybird Conduit Connector",
        Version:     "v0.1.0-SNAPSHOT",
        Author:      "Alejandro Martin <alejandromav@tinybird.co>",
    }
}
