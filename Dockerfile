FROM golang:alpine

LABEL maintainer="@goLangRestApi <eugeneteu@gmail.com>"

WORKDIR /

COPY go.mod .

COPY go.sum .

RUN go mod download 

COPY . .

RUN chmod +x ./wait-for.sh

ENV PORT ":8000"

EXPOSE 8000 8000

RUN go build 

CMD ["./wait-for.sh" , "mysql:3306" , "--timeout=300" , "--" , "./m"]