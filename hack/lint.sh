#!/usr/bin/env bash

#    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
#  \\\\\ Copyright 2024-present SPIKE contributors.
# \\\\\\\ SPDX-License-Identifier: Apache-2.0

go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run -v
# '-v' flag is used for verbose output
