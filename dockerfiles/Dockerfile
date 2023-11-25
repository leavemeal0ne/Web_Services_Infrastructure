FROM golang:1.20 AS server_builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go mod tidy

RUN  CGO_ENABLED=0 go build -o ./bin/main cmd/main.go


FROM scratch AS runner

COPY --from=server_builder /app/bin/main .
COPY /configs ./configs
EXPOSE 8080

CMD ["./main"]