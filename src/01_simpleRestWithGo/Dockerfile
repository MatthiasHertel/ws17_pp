# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:latest

# Create a directory inside the container to store all our application and then make it the working directory.

# Copy the 1_simpleRestWithGo directory (where the Dockerfile lives) into the container.
COPY ./src /go/src/github.com/MatthiasHertel/ws17_pp
COPY ./cmd/server /go/src/github.com/MatthiasHertel/ws17_pp/cmd/server
WORKDIR /go/src/github.com/MatthiasHertel/ws17_pp/cmd/server

RUN go get github.com/gorilla/mux

RUN go build -o /go/bin/server .

# Download and install any required third party dependencies into the container.
# RUN go-wrapper download
# RUN go-wrapper install

# Set the PORT environment variable inside the container
ENV PORT 8080

# Expose port 8080 to the host so we can access our application
EXPOSE 8080

# Now tell Docker what command to run when the container starts
CMD ["/go/bin/server"]
