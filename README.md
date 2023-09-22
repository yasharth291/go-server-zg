# Go Server Application

A simple Go server application that allows clients to send messages to the server, and the server responds with acknowledgments.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Concurrent Handling](#concurrent-handling)
- [Testing](#testing)
- [Error Handling](#error-handling)

## Introduction

This Go server application demonstrates a basic client-server interaction where clients can send messages to the server, and the server responds with acknowledgments. The server handles incoming client connections concurrently using goroutines and ensures thread safety with a mutex.

## Features

- Clients can connect and send messages to the server.
- The server responds with acknowledgments for received messages.
- Concurrent handling of multiple client connections using goroutines.
- Synchronization with a mutex to prevent data races.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed (version 1.23)

## Getting Started

Follow these steps to get the project up and running.

### Installation

Clone the repository:

```bash
git clone https://github.com/yasharth291/go-server.git
cd go-server
```

### Usage

```bash
go run main.go
```

### Project Structure

```bash
go-server/
  ├── server        # Server implementation
      ├─server.go
      ├─client_handler.go
  ├── README.md        # Project documentation (you are here)
  ├── main.go # entry file
```

### Concurrent Handling

Incoming client connections are handled concurrently using goroutines. Each client connection is processed in a separate goroutine, ensuring that multiple clients can connect and communicate with the server concurrently without conflicts. A mutex is used to prevent data races when handling client connections and messages.

### Testing

Unit tests are included to validate the functionality of the server. You can run the tests using the following command:

```bash
go test
```
### Error Handling

The application handles errors gracefully, including network-related errors. Error messages are logged to provide visibility into any issues that may occur during execution.








