# TODO: move to charts directory, create one root Earthfile referencing it
VERSION 0.6
ARG OWNER=brumhard
ARG REPO=primate

# # release
# ## app
# - tag new app version (trigger)
# LATEST_CHART_VERSION=$(git describe --tags --abbrev=0 | rg 'primate-chart-(.*)' -r '$1' |sort -V |tail -1)
# LATEST_APP_VERSION=$(svu current)
# - release new app version like in e.g. kindacool -> goreleaser and build etc
# - release new chart patch version (with new chart tag)

# ## chart
# - tag new chart version (trigger)
# - release new chart version (like cr update)
# - update gh-pages index (like cr index)

chart-test:
    FROM quay.io/helmpack/chart-testing
    WORKDIR /src
    COPY --dir deployments/primate .
    RUN ct lint --all --validate-maintainers=false --chart-dirs .

# export GITHUB_ACCESS_TOKEN=""
# earthly --push --secret GITHUB_TOKEN=$GITHUB_ACCESS_TOKEN +chart-release --GIT_EMAIL=code@brumhard.com --GIT_USER=Tobias Brumhard

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



