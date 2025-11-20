//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

// Close releases any resources held by the API instance.
//
// It ensures proper cleanup of the underlying X509Source. This method should
// be called when the API instance is no longer needed, typically during
// application shutdown or cleanup.
//
// Returns:
//   - *sdkErrors.SDKError: nil if successful or source is nil,
//     ErrSPIFFEFailedToCreateX509Source if closure fails
//
// Example:
//
//	api, err := NewAPI(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer func() {
//	    if err := api.Close(); err != nil {
//	        log.Printf("Failed to close API: %v", err)
//	    }
//	}()
func (a *API) Close() *sdkErrors.SDKError {
	if a.source == nil {
		return nil
	}
	return spiffe.CloseSource(a.source)
}
