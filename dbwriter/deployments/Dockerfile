FROM golang

WORKDIR /app/dbwriter

COPY . .

COPY ./wait-for-postgres.sh .

#download psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh
RUN chmod +x wait-for-it.sh

RUN chmod -R 755 ./temp

RUN go mod download
RUN go build -o ./main ./cmd/main.go
