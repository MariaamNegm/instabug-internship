FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /internship-2023
EXPOSE 9090
CMD [ "go","run","." ]