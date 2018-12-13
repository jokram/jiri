for /f "delims=" %%A in ('git rev-parse --short HEAD') do setx GIT_COMMIT_ID "%%A"
for /f "delims=" %%B in ('git log -n1 "--format=%%ci"') do setx GIT_COMMIT_DATE "%%B"