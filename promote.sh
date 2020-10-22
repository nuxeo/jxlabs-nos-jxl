#!/bin/bash

echo "promoting the new version ${VERSION} to downstream repositories"

jx step create pr regex \
    --regex 'version: (.*)' \
    --version ${VERSION} \
    --files docker/gcr.io/build-jx-prod/jxlabs-nos-master/jxlabs-nos-jxl.yml \
    --files charts/jxlabs-nos/jxl-boot.yml \
    --repo https://github.com/nuxeo/jxlabs-nos-versions.git
