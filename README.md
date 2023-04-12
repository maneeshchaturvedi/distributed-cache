# Distributed Cache

A simplified distributed cache system using consistent hashing and chain replication in Go. This project demonstrates the basic concepts and structure needed for a distributed cache system, but it may require further optimizations and features to be production-ready.

## Project Structure

- `cacheclient`: Contains the cache client implementation and the consistent hashing algorithm.
- `cacheprotocol`: Contains the cache node protocol and the chain replication implementation.
- `main.go`: The main entry point for the test harness.

## Features

- Consistent hashing for distributing keys among multiple cache nodes.
- Chain replication for data replication among nodes.
- Basic cache node operations: get, set, and delete.
- Test harness for measuring system performance under different loads.

## Usage

1. Clone the repository.

2. Build the project:

```sh
go build
