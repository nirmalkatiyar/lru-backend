# LRU Cache Manager

## Project Structure
    ├── api
    │   ├── websocket.go       # WebSocket implementation
    |   |── handler.go         # API handlers
    │   └── router.go          # API routes and middleware setup
    ├── api
    |   ├── cache.go           # Core cache logic
    ├── main.go                # Entry point for the application
    ├── go.mod                 # Go module dependencies
    ├── go.sum                 # Go module checksums
    └── README.md              # Project documentation


## Overview
LRUCacheManager is a Go-based API that implements a Least Recently Used (LRU) cache system. The cache supports operations to set, retrieve, and delete cache entries, and it includes a WebSocket endpoint for real-time updates.

## Features
1. LRU Cache: Automatically manages cache entries, evicting the least recently used items when the cache reaches its capacity.
2. RESTful API: Provides endpoints to get, set, and delete cache entries.
3. WebSocket Updates: Real-time notifications for cache changes.
4. CORS Support: Configured to handle cross-origin requests.

## Installation

These instructions will help you set up and run the project on your local machine.

### Prerequisites

Ensure you have [Go](https://golang.org/doc/install) installed on your system. You can check your Go version with:
    go version
    git clone <repository-url>
    cd <repository-directory>


### Install dependencies:
    go mod tidy

### Build  and run the project:
    go run main.go 

The server will start on the default port 8080.


