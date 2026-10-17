#!/bin/bash
# Fixes reporting parameters, enum names and request body wrappers.
set -euo pipefail
trap 'printf "spec patch failed at line %d\n" "$LINENO" >&2' ERR
go tool patch-spec
