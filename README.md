# sample-test-service

A minimal Go HTTP service used to test the DevLift deployment pipeline against a
**public** repository, where GitHub branch rulesets are enforced on the free plan.

Deliberately contains no configuration, credentials or private code — everything
here is safe to be public.

## Endpoints

| Path       | Purpose                     |
|------------|-----------------------------|
| `/healthz` | Liveness/readiness probe    |
| `/`        | Service name, env, timestamp|

## Run locally

    go run ./app          # listens on :8080 (override with PORT)

## Why this repo exists

DevLift opens a *service workflow PR* against a service repo during deployment.
This repo is the target for verifying that a blocked merge (required approvals
via a branch ruleset) is reported correctly — the deployment must fail and alert,
not silently report success.
