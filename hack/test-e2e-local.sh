#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo "> E2E Tests"

# reduce flakiness in contended pipelines
export GOMEGA_DEFAULT_EVENTUALLY_TIMEOUT=5s
export GOMEGA_DEFAULT_EVENTUALLY_POLLING_INTERVAL=200ms
# if we're running low on resources, it might take longer for tested code to do something "wrong"
# poll for 5s to make sure, we're not missing any wrong action
export GOMEGA_DEFAULT_CONSISTENTLY_DURATION=5s
export GOMEGA_DEFAULT_CONSISTENTLY_POLLING_INTERVAL=200ms

ginkgo_flags=

shoot_names=(
  e2e-managedseed.garden
  e2e-hibernated.local
  e2e-unpriv.local
  e2e-wake-up.local
  e2e-migrate.local
  e2e-rotate.local
  e2e-default.local
)

# If running in prow, we want to generate a machine-readable output file under the location specified via $ARTIFACTS.
# This will add a JUnit view above the build log that shows an overview over successful and failed test cases.
if [ -n "${CI:-}" -a -n "${ARTIFACTS:-}" ]; then
  mkdir -p "$ARTIFACTS"
  ginkgo_flags="--output-dir=$ARTIFACTS --junit-report=junit.xml"

  # make shoot domains accessible to test
  for shoot in "${shoot_names[@]}" ; do
    printf "\n127.0.0.1 api.%s.external.local.gardener.cloud\n127.0.0.1 api.%s.internal.local.gardener.cloud\n" $shoot $shoot >>/etc/hosts
  done
  printf "\n127.0.0.1 gu-local--e2e-rotate.ingress.local.seed.local.gardener.cloud\n" >>/etc/hosts
  printf "\n127.0.0.1 api.e2e-managedseed.garden.external.local.gardener.cloud\n127.0.0.1 api.e2e-managedseed.garden.internal.local.gardener.cloud\n" >>/etc/hosts
else
  for shoot in "${shoot_names[@]}" ; do
    if ! grep -q "$(printf "\n127.0.0.1 api.%s.external.local.gardener.cloud\n127.0.0.1 api.%s.internal.local.gardener.cloud\n" $shoot $shoot)" /etc/hosts; then
      printf "To access shoot clusters and run e2e tests, you have to extend your /etc/hosts file.\nPlease refer to https://github.com/gardener/gardener/blob/master/docs/deployment/getting_started_locally.md#accessing-the-shoot-cluster\n"
    fi
  done
fi

for ((i = 2; i <= "$#"; i++)); do
  if [ "${!i}" = "--" ]; then
    break
  fi
done

GO111MODULE=on ginkgo run --timeout=1h $ginkgo_flags "${@:1:$((i - 1))}" --v --progress ./test/e2e/... "${@:$i}"
