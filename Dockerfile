FROM  golang:latest
WORKDIR /link_shortener_bot
COPY . .
RUN go mod download
RUN go build -o link_shortener_bot ./cmd...
CMD ./link_shortener_bot
