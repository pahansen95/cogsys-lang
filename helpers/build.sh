#!/usr/bin/env bash

set -xeEou pipefail

: "${CI_PROJECT_DIR:?"CI_PROJECT_DIR is not set."}"

cd "${CI_PROJECT_DIR}/src"
go install -v -buildvcs=false
