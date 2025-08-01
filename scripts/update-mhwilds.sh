#!/bin/bash

# Absolute paths
POOGIE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MHW_REPO="$POOGIE_DIR/../mhdb-wilds-data"
MHW_SOURCE="$MHW_REPO/output/merged"
MHW_DEST="$POOGIE_DIR/data/mhwilds"

echo "🔄 Pulling latest data from mhdb-wilds-data..."
cd "$MHW_REPO" || exit 1
git pull origin main || exit 1

echo "📦 Copying merged data into PoogieBot..."
mkdir -p "$MHW_DEST"
cp -r "$MHW_SOURCE/"* "$MHW_DEST/"

echo "✅ Data update complete!"
