# setup project and deps
FROM golang:1.23-bullseye AS init

WORKDIR /go/url2anki/

COPY go.mod* go.sum* ./
RUN go mod download

COPY . ./

FROM init as vet
RUN go vet ./...

# run tests
FROM init as test
RUN go test -coverprofile c.out -v ./...

# build binary
FROM init as build
ARG LDFLAGS

RUN CGO_ENABLED=0 go build -ldflags="${LDFLAGS}"

# runtime image
FROM scratch
# Copy our static executable.
COPY --from=build /go/url2anki/url2anki /go/bin/url2anki
# Run the binary.
ENTRYPOINT ["/go/bin/url2anki"]