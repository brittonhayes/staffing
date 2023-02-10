# Staffing Service

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/staffing.svg)](https://pkg.go.dev/github.com/brittonhayes/staffing)
![Latest Release](https://img.shields.io/github/v/release/brittonhayes/staffing?label=latest%20release)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/staffing)](https://goreportcard.com/report/github.com/brittonhayes/staffing)

This is an example of a staffing service for a fictional consulting company. It is an implementation of [Domain Driven Design](https://www.amazon.com/Domain-Driven-Design-Tackling-Complexity-Software/dp/0321125215) and [Clean Architecture](https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164).

<!-- TODO CQRS and event-sourcing are in-progress -->
The service uses [CQRS](https://martinfowler.com/bliki/CQRS.html) and [Event Sourcing](https://martinfowler.com/eaaDev/EventSourcing.html) to implement the business logic. The service is implemented in [Go](https://golang.org/).

The structure of the service is based off [github.com/marcusolsson/goddd by Marcus Olsson](https://github.com/marcusolsson/goddd).

## Usage

```bash
# Clone the repository
gh repo clone brittonhayes/staffing

# Run the service with default settings
go run ./cmd/staffd/main.go
```

### Configuration

The service can be run with a few different customization options. These options can be set with environment variables or command line flags.

```bash
# OR run the service with in-memory storage and debug logging
go run ./cmd/staffd/main.go -inmem -debug
# OR run the service with BoltDB storage and debug logging
go run ./cmd/staffd/main.go -debug
```

## Docker

```bash
# Build the image
docker build -t ghcr.io/brittonhayes/staffing:latest .

# Run the image
docker run ghcr.io/brittonhayes/staffing:latest
```

## Development

```bash
# Clone the repository
gh repo clone brittonhayes/staffing
cd ./staffing

# Checkout to a new branch
git checkout -b feat/my-new-feature

# Commit your changes
git commit -am "feat:Add some feature"

# Create a PR
gh pr create -w
```

## Resources

### Articles

- http://www.citerus.se/go-ddd
- http://www.citerus.se/part-2-domain-driven-design-in-go
- http://www.citerus.se/part-3-domain-driven-design-in-go