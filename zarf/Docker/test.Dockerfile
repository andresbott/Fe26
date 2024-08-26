FROM fe26-builder:latest

WORKDIR /project
COPY . .

RUN make verify