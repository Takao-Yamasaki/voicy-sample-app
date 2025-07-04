FROM golang:1.24-alpine
WORKDIR /app
COPY . ./
RUN go mod download \
  && go build -o main .
USER 1001
CMD [ "/app/main" ]
