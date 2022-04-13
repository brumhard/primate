# TODO: move to charts directory, create one root Earthfile referencing it
VERSION 0.6
ARG OWNER=brumhard
ARG REPO=primate

chart-test:
    FROM quay.io/helmpack/chart-testing
    WORKDIR /src
    COPY --dir deployments/primate .
    RUN ct lint --all --validate-maintainers=false --chart-dirs .

chart-release:
    FROM quay.io/helmpack/chart-releaser
    WORKDIR /src

    COPY --dir deployments/primate .
    RUN cr package ./primate

    COPY --dir .git .
    # TODO: check if git config is already set
    ARG --required GIT_EMAIL
    ARG --required GIT_USER
    RUN git config user.email "$GIT_EMAIL" && \
        git config user.name "$GIT_USER"
    # fixes https://github.com/helm/chart-releaser/issues/124
    RUN git remote set-url origin https://github.com/$OWNER/$REPO

    RUN --push \
        --secret GITHUB_TOKEN \
        cr upload \
            --owner $OWNER \
            --git-repo $REPO \
            --release-name-template "{{ .Name }}-chart-{{ .Version }}" \
            --token $GITHUB_TOKEN \
            --skip-existing
    RUN --push \
        --secret GITHUB_TOKEN \
        mkdir .cr-index && \
        cr index \
            --owner brumhard \
            --git-repo primate \
            --release-name-template "{{ .Name }}-chart-{{ .Version }}" \
            --token $GITHUB_TOKEN \
            --pages-branch gh-pages \
            --push



