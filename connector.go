package tinybird

import (
    destination "github.com/alejandromav/conduit-connector-tinybird/destination"
    sdk "github.com/conduitio/conduit-connector-sdk"
)

var Connector = sdk.Connector{
    NewSpecification: Specification,
    NewDestination:   destination.NewDestination,
}
