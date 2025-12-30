//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/config/env"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

// Source creates and returns a new SPIFFE X509Source for workload API
// communication. It establishes a connection to the SPIFFE workload API using
// the default endpoint socket with a configurable timeout to prevent indefinite
// blocking on socket issues.
//
// The timeout can be configured using the SPIKE_SPIFFE_SOURCE_TIMEOUT
// environment variable (default: 30s).
//
// The function will terminate the program with exit code 1 if the source
// creation fails or times out.
//
// Returns:
//   - *workloadapi.X509Source: A new X509Source for SPIFFE workload API
//     communication
func Source() *workloadapi.X509Source {
	const fName = "Source"

	ctx, cancel := context.WithTimeout(
		context.Background(),
		env.SPIFFESourceTimeoutVal(),
	)
	defer cancel()

	source, _, err := spiffe.Source(ctx, spiffe.EndpointSocket())
	if err != nil {
		failErr := sdkErrors.ErrSPIFFEUnableToFetchX509Source.Wrap(err)
		log.FatalErr(fName, *failErr)
	}
	return source
}
