FROM fe26-builder:latest

WORKDIR /project
COPY . .

RUN make package-ui
RUN goreleaser build --auto-snapshot --clean -f .goreleaser.yaml
RUN chmod -R 0777 dist/