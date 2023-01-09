FROM golang:1.19

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./ 
COPY go.sum ./
RUN go mod download

COPY . .

#Setup hot-reload for dev stage
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
