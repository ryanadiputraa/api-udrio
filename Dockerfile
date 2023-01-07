FROM golang:1.17-alpine
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./ 
COPY go.sum ./
RUN go mod download

COPY . .
COPY .udrio.yml ./

RUN go build -o /api-udrio

EXPOSE 8080

CMD ["/api-udrio"]
