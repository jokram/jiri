#!/bin/bash
export GIT_COMMIT_ID=$(git rev-parse --short)
export GIT_COMMIT_DATE=$(git log -n1 --format=%ci)