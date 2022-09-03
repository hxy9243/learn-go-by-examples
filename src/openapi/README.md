OpenAPI Go Web App Example
====

OpenAPI is the standard for documenting RESTful API for web applications and services. A well-documented
API not only makes life easier for end-user to understand and consume, but also helps greatly with
development, and testing.

OpenAPI compatible toolchains also helps generating code for your server / client applications,
so as to keep good consistency between your openAPI definition and implementation.

This example demos an example Go webapp in [Gorilla](https://github.com/gorilla/mux) framework
as well as a cli command line in Python.

See more at:

- OpenAPI: https://www.openapis.org/
- OpenAPI Generator: https://openapi-generator.tech/docs/generators/go-server/
- OpenAPI Generator Github: https://github.com/OpenAPITools/openapi-generator
- Go Gorilla Framework: https://github.com/gorilla/mux

# Project Layout

- [api](api): the OpenAPI 3.0 definition in YAML, including [api.yaml](api/api.yaml) as the
    main definitions of all the REST API paths, and other files describing data models.
- [server](server): the server implementation, including generationg files and implementation.
    - [app](server/app): the main function and entry of the application.
    - [gen](server/gen): the generated code from openapi-generator.
        - [doc](server/gen/doc): the HTML documentation for the API.
        - [openapi](server/gen/openapi): the generated Go source code for openAPI.
    - [service](server/service): the service implementation.


# Quick Start

## Build

Run go mod to build the output Go binary in `build/` directory.

```
make
```

## Run server

After building, run server with

`./build/server [address]`

where address defaults to `0.0.0.0:8000`.

You can query the endpoints with curl or other HTTP testing tools.
Currently only the `/books` API is actually implemented with some mock data.

# Development

## Dependencies

To contribute to this project, you'll need the following toolchains:

- openAPI generator, see: https://openapi-generator.tech/docs/installation

## Documentation

Run `make doc` to generate HTML document in [API doc directory](server/gen/doc).

## Code Generation

After modifying the openAPI definition in [API directory](openapi/api), run
`make gen` to regenerate the openapi Go source code.

The results will be updated in [server/gen/openapi](server/gen/openapi).
