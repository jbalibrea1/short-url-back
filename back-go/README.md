# GO API - Short URL Service

This is a URL shortening service built with Go, Gin, and MongoDB.

## Prerequisites

- **Go** version 1.22.4+ is required.
- **MongoDB**: You must have a running instance of MongoDB, either locally or remotely.
  - Set the MongoDB connection URI in your environment variables (see below).

## Usage

Set the environment variables in a `.env` file like the `.env.example` file.

```bash
go mod download
go run ./cmd/api/main.go
```
