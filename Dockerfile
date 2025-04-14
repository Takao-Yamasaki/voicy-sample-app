FROM golang:1.23-alpine
WORKDIR /app
# FIXME: コメントアウト解除
# COPY go.mod go.sum main.go ./
COPY . ./
RUN go mod download \
  && go build -o main /app/main.go
USER 1001
CMD [ "/app/main" ] 
