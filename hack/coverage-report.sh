#!/usr/bin/env bash

#    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
#  \\\\\ Copyright 2024-present SPIKE contributors.
# \\\\\\\ SPDX-License-Identifier: Apache-2.0

# This script generates an HTML coverage report for the SPIKE SDK Go codebase
# and publishes it to the documentation directory.
#
# Usage: ./hack/coverage-report.sh
#
# Output: $HOME/WORKSPACE/spike/docs/sdk/coverage.html
#
# The script performs the following steps:
#   1. Runs all tests with coverage profiling
#   2. Generates an HTML coverage report
#   3. Copies the report to the documentation directory
#   4. Displays a summary of the coverage results

set -euo pipefail

# Color codes for output
readonly RED='\033[0;31m'
readonly GREEN='\033[0;32m'
readonly YELLOW='\033[1;33m'
readonly BLUE='\033[0;34m'
readonly NC='\033[0m' # No Color

# Paths
readonly COVERAGE_OUT="/tmp/spike-sdk-go-coverage.out"
readonly COVERAGE_HTML="/tmp/spike-sdk-go-coverage.html"
readonly DOCS_DIR="${HOME}/WORKSPACE/spike/docs/sdk"
readonly TARGET_HTML="${DOCS_DIR}/coverage.html"

echo -e "${BLUE}╔════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║         SPIKE SDK Go - Coverage Report Generator              ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════════════╝${NC}"
echo ""

# Step 1: Run tests with coverage
echo -e "${YELLOW}📊 Running tests with coverage profiling...${NC}"
if go test -v -race -buildvcs -coverprofile="${COVERAGE_OUT}" ./...; then
    echo -e "${GREEN}✓ Tests completed successfully${NC}"
else
    echo -e "${RED}✗ Tests failed${NC}"
    exit 1
fi
echo ""

# Step 2: Generate HTML coverage report
echo -e "${YELLOW}🔨 Generating HTML coverage report...${NC}"
if go tool cover -html="${COVERAGE_OUT}" -o="${COVERAGE_HTML}"; then
    echo -e "${GREEN}✓ HTML report generated: ${COVERAGE_HTML}${NC}"
else
    echo -e "${RED}✗ Failed to generate HTML report${NC}"
    exit 1
fi
echo ""

# Step 3: Create docs directory if it doesn't exist
if [ ! -d "${DOCS_DIR}" ]; then
    echo -e "${YELLOW}📁 Creating documentation directory: ${DOCS_DIR}${NC}"
    mkdir -p "${DOCS_DIR}"
fi

# Step 4: Copy report to documentation directory
echo -e "${YELLOW}📤 Publishing report to documentation...${NC}"
if cp "${COVERAGE_HTML}" "${TARGET_HTML}"; then
    echo -e "${GREEN}✓ Coverage report published to: ${TARGET_HTML}${NC}"
else
    echo -e "${RED}✗ Failed to publish coverage report${NC}"
    exit 1
fi
echo ""

# Step 5: Display coverage summary
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}Coverage Summary:${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# Extract and display coverage percentages
go tool cover -func="${COVERAGE_OUT}" | tail -1

echo ""
echo -e "${GREEN}✓ Coverage report generation complete!${NC}"
echo -e "${BLUE}  View the report at: file://${TARGET_HTML}${NC}"
echo ""
