![SPIKE](assets/spike-banner-lg.png)


## SPIKE Go SDK

This library is a convenient Go library for working with [SPIKE](https://spike.ist/).

It leverages the [SPIFFE Workload API](https://github.com/spiffe/spiffe/blob/main/standards/SPIFFE_Workload_API.md), 
providing high level functionality that includes:
* Establishing mutually authenticated TLS (__mTLS__) between workloads powered by [SPIFFE](https://spiffe.io).
* Abstracting SPIKE REST API calls.

## Documentation

See the [Go Package](https://pkg.go.dev/github.com/spiffe/spike-sdk-go) 
documentation.

## Quick Start

Prerequisites:
1. Running [SPIRE](https://spiffe.io/spire/) or another SPIFFE Workload API
   implementation.
2. `SPIFFE_ENDPOINT_SOCKET` environment variable set to address of the Workload
   API (e.g. `unix:///tmp/agent.sock`). 

// TODO: add more details and an example here.

// TODO: add more documents COC, contributing, etc.