#!/bin/bash

# Version management script for local development
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
VERSION_FILE="$ROOT_DIR/VERSION"

get_current_version() {
    if [ -f "$VERSION_FILE" ]; then
        cat "$VERSION_FILE"
    else
        echo "v0.0.0"
    fi
}

increment_patch() {
    local current=$(get_current_version)
    local major=$(echo $current | cut -d. -f1 | sed 's/v//')
    local minor=$(echo $current | cut -d. -f2)
    local patch=$(echo $current | cut -d. -f3)

    patch=$((patch + 1))
    local new_version="v${major}.${minor}.${patch}"
    echo "$new_version" > "$VERSION_FILE"
    echo "$new_version"
}

increment_minor() {
    local current=$(get_current_version)
    local major=$(echo $current | cut -d. -f1 | sed 's/v//')
    local minor=$(echo $current | cut -d. -f2)

    minor=$((minor + 1))
    local new_version="v${major}.${minor}.0"
    echo "$new_version" > "$VERSION_FILE"
    echo "$new_version"
}

increment_major() {
    local current=$(get_current_version)
    local major=$(echo $current | cut -d. -f1 | sed 's/v//')

    major=$((major + 1))
    local new_version="v${major}.0.0"
    echo "$new_version" > "$VERSION_FILE"
    echo "$new_version"
}

case "${1:-}" in
    "current"|"get")
        get_current_version
        ;;
    "patch")
        increment_patch
        ;;
    "minor")
        increment_minor
        ;;
    "major")
        increment_major
        ;;
    *)
        echo "Usage: $0 {current|patch|minor|major}"
        echo ""
        echo "Commands:"
        echo "  current  - Show current version"
        echo "  patch    - Increment patch version (x.x.X)"
        echo "  minor    - Increment minor version (x.X.0)"
        echo "  major    - Increment major version (X.0.0)"
        echo ""
        echo "Current version: $(get_current_version)"
        exit 1
        ;;
esac