#!/usr/bin/env bash

readonly RELEASE_DIR=${RELEASE_DIR:-release}

if [ -d "${RELEASE_DIR}" ]; then
  echo "Removing ${RELEASE_DIR} directory"
  rm -r "${RELEASE_DIR}"
fi

mkdir "${RELEASE_DIR}"
readonly ARTIFACT="${RELEASE_DIR}/duckeye.yaml"

ko resolve --strict -f config >"${ARTIFACT}"
