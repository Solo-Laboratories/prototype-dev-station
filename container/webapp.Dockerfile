# First stage: Build the application
FROM golang:1.23.2 AS build-stage

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY web-app .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dev-station .

# Second stage: Create the final image
FROM scratch

# Copy the built application from the first stage
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-stage /app/dev-station /dev-station
COPY --from=build-stage /app/index.html /
COPY --from=build-stage /app/values-files /values-files
COPY --from=build-stage /app/manifest-files /manifest-files

# Set the entrypoint to the executable
ENTRYPOINT ["/dev-station"]
