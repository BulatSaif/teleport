#!/bin/bash

set -euo pipefail

info() {
    printf "\e[1;32m%s\e[0m\n" "$1"
}
error() {
    printf "\e[1;31m%s\e[0m\n" "$1"
    exit 1
}

VERSION="$1"

if [ -z "$VERSION" ]; then
  error "Version parameter is required"
fi

GENERATOR="tfplugindocs"

if ! command -v "$GENERATOR" &> /dev/null
then
    error "tfplugindocs could not be found"
fi

TFDIR="$(pwd)"
DOCSDIR="$(pwd)/../../docs/pages/reference/terraform"
TMPDIR="$(mktemp -d)"

info "Generating provider's schema"

cd "$TMPDIR"
cat > main.tf << EOF
terraform {
  required_providers {
    teleport = {
      source  = "terraform.releases.teleport.dev/gravitational/teleport"
      version = "= $VERSION"
    }
  }
}
EOF

terraform init
terraform providers schema -json > schema.json

info "Rendering markdown files"

$GENERATOR generate \
  --providers-schema "$TMPDIR/schema.json" \
  --provider-name "terraform.releases.teleport.dev/gravitational/teleport" \
  --rendered-provider-name "teleport" \
  --rendered-website-dir="$TMPDIR/docs" \
  --website-source-dir="$TFDIR/templates" \
  --provider-dir "$TFDIR" \
  --examples-dir="$TFDIR/examples" \
  --website-temp-dir="$TMPDIR/temp"

info "Converting .md files to .mdx"

cd "$TMPDIR/docs"
find . -iname '*.md' -type f -exec sh -c 'i="$1"; mv "$i" "${i%.md}.mdx"' shell {} \;

info "Copying generated documentation into the teleport docs directory"

rm -rf "$DOCSDIR"
cp -r "$TMPDIR/docs" "$DOCSDIR"

info "TF documentation successfully generated"