# Get latest golang docker image.
FROM golang:latest

# Create a directory inside the container to store our web-app and then make it working directory.
RUN mkdir -p /go/src/github.com/klar
WORKDIR /go/src/github.com/klar

# Copy the web-app directory into the container.
COPY . /go/src/github.com/klar

# Download and install third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install

# Set the PORT environment variable
ENV PORT 3000

# Expose port 8080 to the host so that outer-world can access your application
EXPOSE 3000

# Tell Docker what command to run when the container starts
CMD ["go-wrapper", "run"]