# Kong API Plugin using Go

This project demonstrates how to write a custom Kong API plugin using Go, build it with Docker, and run it with Kong using Docker Compose. The plugin adds a custom header (`X-Custom-Header`) to requests processed by Kong.

## Prerequisites

Before getting started, ensure you have the following installed on your system:

- Docker and Docker Compose
- Go (version 1.18 or later)
- Basic knowledge of Go and Docker

## Setup Instructions

### Step 1: Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/yourusername/kong-api-plugin-go.git
cd kong-api-plugin-go
```

### Step 2: Build the Plugin

Build the plugin using the following command:

```bash
chmod +x start.sh
./start.sh
```

