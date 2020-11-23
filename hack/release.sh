#!/usr/bin/env bash

readonly RELEASE_DIR=${RELEASE_DIR:-release}
readonly DISCOVERY_VERSION=${DISCOVERY_VERSION:-"v0.19.0"}

function fail_test() {
  echo "$1"
  exit 1
}

if [ -d "${RELEASE_DIR}" ]; then
  echo "Removing ${RELEASE_DIR} directory"
  rm -r "${RELEASE_DIR}"
fi

mkdir "${RELEASE_DIR}"
readonly ARTIFACT="${RELEASE_DIR}/duckeye.yaml"

rm "${ARTIFACT}"
rm -r cmd/duckeye/kodata/*

cd ui && npm run build && npm run movebuild && cd - || exit 1

wget https://github.com/knative-sandbox/discovery/releases/download/"${DISCOVERY_VERSION}"/discovery-core.yaml -O - >>"${ARTIFACT}" || fail_test "Failed to download discovery-core.yaml"
wget https://github.com/knative-sandbox/discovery/releases/download/"${DISCOVERY_VERSION}"/knative-duck-types.yaml -O - >>"${ARTIFACT}" || fail_test "Failed to download discovery-duck-types.yaml"
ko resolve --strict -f config >>"${ARTIFACT}" || fail_test "Failed to resolve duckeye manifests"
