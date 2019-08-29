FROM golang:latest

WORKDIR /app
COPY . .

# get all the dependencies
RUN go mod download
# build the application
RUN go build -o main .

EXPOSE 3000

# run the server
CMD ["./main"]