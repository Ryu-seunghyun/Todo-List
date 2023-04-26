FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["sh", "-c", "./main" ,">" , "todo.log", "2>&1"]