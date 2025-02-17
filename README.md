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

## Usage Example

```go 
package main

import (
	"fmt"

	spike "github.com/spiffe/spike-sdk-go/api"
)

func main() {
	api := spike.New() // Use the default Workload API Socket
	defer api.Close()  // Close the connection when done

	path := "/tenants/demo/db/creds"

	// Create a Secret
	err := api.PutSecret(path, map[string]string{
		"username": "SPIKE",
		"password": "SPIKE_Rocks",
	})
	if err != nil {
		fmt.Println("Error writing secret:", err.Error())
		return
	}

	// Read the Secret
	secret, err := api.GetSecret(path)
	if err != nil {
		fmt.Println("Error reading secret:", err.Error())
		return
	}

	if secret == nil {
		fmt.Println("Secret not found.")
		return
	}

	fmt.Println("Secret found:")

	data := secret.Data
	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}
}
```

## A Note on Security

We take **SPIKE**'s security seriously. If you believe you have
found a vulnerability, please responsibily disclose it to
[security@spike.ist](mailto:security@spike.ist).

See [SECURITY.md](SECURITY.md) for additional details.

## Community

Open Source is better together.

If you are a security enthusiast, [join SPIFFE's Slack Workspace][spiffe-slack]
and let us change the world together 🤘.

# Contributing

To contribute to **SPIKE**, [follow the contributing
guidelines](CONTRIBUTING.md) to get started.

Use GitHub issues to request features or file bugs.

## Communications

* [SPIFFE **Slack** is where the community hangs out][spiffe-slack].
* [Send comments and suggestions to
  **feedback@spike.ist**](mailto:feedback@spike.ist).

## License

[Mozilla Public License v2.0](LICENSE).

[spiffe-slack]: https://slack.spiffe.io/
[spiffe]: https://spiffe.io/
[spike]: https://spike.ist/
[quickstart]: https://spike.ist/#/quickstart
