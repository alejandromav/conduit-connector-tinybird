package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"

	tinybird "github.com/alejandromav/conduit-connector-tinybird"
)

func main() {
	sdk.Serve(tinybird.Connector)
}
