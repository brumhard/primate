VERSION 0.6

chart-test:
    FROM quay.io/helmpack/chart-testing
    WORKDIR /src
    COPY --dir deployments/primate .
    RUN ct lint --all --validate-maintainers=false --chart-dirs .

chart-release:
    FROM quay.io/helmpack/chart-releaser
    WORKDIR /src
    COPY --dir deployments/primate .
    RUN cr package ./primate --package-path .cr-release-packages
    SAVE ARTIFACT .cr-release-packages

chart-index-update:
    FROM quay.io/helmpack/chart-releaser
    WORKDIR /src
    COPY --dir +chart-release/.cr-release-packages .
    COPY --dir .git .
    RUN --secret GITHUB_TOKEN \
        cr index \
            --owner brumhard \
            --git-repo primate \
            --package-path .cr-release-packages \
            --pages-branch gh-pages \
            --push $GITHUB_TOKEN


