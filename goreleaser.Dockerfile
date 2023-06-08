FROM cgr.dev/chainguard/static:latest

COPY gitlab-runner-cleaner /usr/bin/gitlab-runner-cleaner

ENTRYPOINT ["/usr/bin/gitlab-runner-cleaner"]
