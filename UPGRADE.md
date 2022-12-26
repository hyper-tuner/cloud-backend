# Upgrade guide

> **Warning**
> Upgrade one **minor** version at a time! Example: `x.1.x` to `x.2.x` to `x.3.x` and so on.

## From v1.1.x to v1.2.x

This version adds tags and verified author badge.

1. import new schema first
2. get the new binary or Dockerfile and run `cloud-backend`, make sure everything works as it was before (don't forget to backup `pb_data` data)
3. upgrade frontend to `v1.2.x`

## From v1.0.1 to v1.1.x

This version adds stargazers.

1. import new schema first
2. get the new binary or Dockerfile and run `cloud-backend`, make sure everything works as it was before (don't forget to backup `pb_data` data)
3. upgrade frontend to `v1.1.x`

## From v1.0.0 to v1.0.1

This version adds to new custom endpoints for fetching tune and ini file.

1. get the new binary or Dockerfile and run `cloud-backend`, make sure everything works as it was before (don't forget to backup `pb_data` data)
2. upgrade frontend to `v1.0.1`
3. import new schema
