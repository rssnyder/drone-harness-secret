A plugin to .

# Usage

The following settings changes this plugin's behavior.

* NAME the name of the secret
* VALUE the content of the secret
* SECRET_MANAGER (optional) the secrets manager to use

For authentication and scope, use the following environment variables:

- HARNESS_PLATFORM_API_KEY: harness nextgen api key
- HARNESS_ACCOUNT_ID: harness account id
- HARNESS_PLATFORM_ORGANIZATION: organization id
- HARNESS_PLATFORM_PROJECT: project id

Below is an example `.drone.yml` that uses this plugin.

```yaml
kind: pipeline
name: default

steps:
- name: run rssnyder/drone-harness-secret plugin
  image: rssnyder/drone-harness-secret
  pull: if-not-exists
  settings:
    NAME: API_KEY
    VALUE: sdflsdjf03239hwip2du
```

# Building

Build the plugin binary:

```text
scripts/build.sh
```

Build the plugin image:

```text
docker build -t rssnyder/drone-harness-secret -f docker/Dockerfile .
```

# Testing

Execute the plugin from your current working directory:

```text
docker run --rm -e PLUGIN_PARAM1=foo -e PLUGIN_PARAM2=bar \
  -e DRONE_COMMIT_SHA=8f51ad7884c5eb69c11d260a31da7a745e6b78e2 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_BUILD_NUMBER=43 \
  -e DRONE_BUILD_STATUS=success \
  -w /drone/src \
  -v $(pwd):/drone/src \
  rssnyder/drone-harness-secret
```
