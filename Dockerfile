FROM golang:1.19

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
