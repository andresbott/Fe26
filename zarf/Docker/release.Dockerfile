FROM fe26-builder:latest

WORKDIR /project
COPY . .

RUN --mount=type=secret,id=GITHUB_TOKEN \
    GITHUB_TOKEN=$(cat /run/secrets/GITHUB_TOKEN) && \
    export GITHUB_TOKEN=$GITHUB_TOKEN && \
    goreleaser --clean -f .goreleaser.yaml

