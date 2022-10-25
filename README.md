# Conduit Connector for Tinybird
[Conduit](https://conduit.io) for Tinybird.

## General
The Tinybird plugin is one of [Conduit](https://github.com/ConduitIO/conduit) plugins.
It currently provides only destination Tinybird connector, allowing for using it as a destination in a Conduit pipeline.

Tinybird provides [an API endpoint](https://www.tinybird.co/docs/api-reference/datasource-api.html#post-v0-events-title) for ingesting events, like the CDC events genrated by Conduit. It supports `ndjson` format, so we can push events in bulk.

Requests to Tinybird must be authenticated using the Admin token in the `Authorization: Bearer <admin_token>` header.

## How to build?
Run `make build` to build the connector.

## Testing
Run `make test` to run all the unit tests. Run `make test-integration` to run the integration tests.

The Docker compose file at `test/docker-compose.yml` can be used to run the required resource locally.

## Source

Not implemented

## Destination
A destination connector pushes data from upstream resources to an external resource via Conduit.

### Configuration

| name          | description                             | required | default value             |
|---------------|-----------------------------------------|----------|---------------------------|
| `host`        | Tinybird host url.                      | false    | `https://api.tinybird.co` |
| `token`       | Tinybird admin token, you can find it [here](https://ui.tinybird.co/f7069c89-9c4b-4305-90b7-0e64589884d9/tokens) | true | |
| `datasource`  | Name of the new data source in Tinybird | true     |                           |

## Known Issues & Limitations
* Only Destination for now, no support for Source conenctor.

## Planned work
- [ ] Item A
- [ ] Item B
