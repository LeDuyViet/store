#!/bin/sh

echo "machine gitlab.volio.vn login ${VOLIO_GIT_DEPLOY_USERNAME} password ${VOLIO_GIT_DEPLOY_PWD}" > ~/.netrc

# git config --global http.sslverify false
go env -w GOPRIVATE=gitlab.volio.vn/*
go env -w GOINSECURE='gitlab.volio.vn'
git config --global --add url."https://${VOLIO_GIT_DEPLOY_USERNAME}:${VOLIO_GIT_DEPLOY_PWD}@gitlab.volio.vn".insteadOf "https://gitlab.volio.vn"