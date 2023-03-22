FROM golang:1.19

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./ 
COPY go.sum ./
RUN go mod download

COPY . .

# use this instead for hot reload on dev env
# RUN go install -mod=mod github.com/githubnemo/CompileDaemon
# ENTRYPOINT CompileDaemon --build="go build -o /api-udrio" --command="/api-udrio"

# build
RUN go build -o /api-udrio
EXPOSE 8080
CMD ["/api-udrio"]

