FROM golang:bookworm
ENV TZ="America/Chicago"
RUN date
WORKDIR /app
COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /test-server-go
EXPOSE 9090
ENTRYPOINT [ "/test-server-go" ]
