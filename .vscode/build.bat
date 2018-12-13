@for /f "delims=" %%A in ('git rev-parse --short HEAD') do @set "GIT_COMMIT_ID=%%A"
@for /f "delims=" %%B in ('git log -n1 "--format=%%ci"') do @set "GIT_COMMIT_DATE=%%B"
@call go install -v -ldflags "-X 'fuchsia.googlesource.com/jiri/version.GitCommit=%GIT_COMMIT_ID%' -X 'fuchsia.googlesource.com/jiri/version.BuildTime=%GIT_COMMIT_DATE%'" -a fuchsia.googlesource.com/jiri/cmd/jiri