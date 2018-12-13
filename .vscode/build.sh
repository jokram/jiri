#!/bin/bash
GIT_COMMIT_ID=$(git rev-parse --short HEAD)
GIT_COMMIT_DATE=$(git log -n1 --format=%ci)
go install -v -ldflags "-X 'fuchsia.googlesource.com/jiri/version.GitCommit=$GIT_COMMIT_ID' -X 'fuchsia.googlesource.com/jiri/version.BuildTime=$GIT_COMMIT_DATE'" -a fuchsia.googlesource.com/jiri/cmd/jiri