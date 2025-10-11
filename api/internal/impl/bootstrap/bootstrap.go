package bootstrap

import (
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

// PutShardContributionRequest

// PostBootstrapVerifyRequest

// Contribute

// Verify

func Contribute(
	source *workloadapi.X509Source,
) error {
	if source == nil {
		return errors.New("nil X509Source")
	}

	return nil
}

func Verify(
	source *workloadapi.X509Source,
) error {
	if source == nil {
		return errors.New("nil X509Source")
	}

	return nil
}
