# Use the official golang latest image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code to the container
COPY . .

# Make script can be executable in the container
RUN chmod +x deploy.sh

# # Build the Go application and output the binary file as "vcg-academy-backend"
RUN go build -o go-grpc-server -buildvcs=false

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 9999


# Set the container's entrypoint to the compiled binary file
CMD ["bash","deploy.sh"]
