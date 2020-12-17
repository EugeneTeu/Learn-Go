FROM golang:1.15.6-alpine3.12 AS build
WORKDIR /
COPY . .
RUN go get github.com/gorilla/mux
RUN go build -o out/example .
FROM scratch AS bin
copy --from=build /out/example /

