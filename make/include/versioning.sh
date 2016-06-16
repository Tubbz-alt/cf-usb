#!/bin/sh

set +o errexit +o nounset

test -n "${XTRACE}" && set -o xtrace

set -o errexit -o nounset

GIT_ROOT=${GIT_ROOT:-$(git rev-parse --show-toplevel)}
GIT_DESCRIBE=${GIT_DESCRIBE:-$(git describe --always --tags --long)}
GIT_BRANCH=${GIT_BRANCH:-$(git name-rev --name-only HEAD)}

GIT_TAG=${GIT_TAG:-$(echo ${GIT_DESCRIBE} | awk -F - '{ print $1 }' )}
GIT_COMMITS=${GIT_COMMITS:-$(echo ${GIT_DESCRIBE} | awk -F - '{ print $2 }' )}
GIT_SHA=${GIT_SHA:-$(echo ${GIT_DESCRIBE} | awk -F - '{ print $3 }' )}

ARTIFACT_NAME=${ARTIFACT_NAME:-$(basename $(git config --get remote.origin.url) .git | sed s/^hcf-//)}
ARTIFACT_VERSION=${GIT_TAG}+${GIT_COMMITS}.${GIT_SHA}.${GIT_BRANCH}

APP_VERSION=${ARTIFACT_NAME}-${ARTIFACT_VERSION}

set +o errexit +o nounset +o xtrace
