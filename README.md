# Online Store Application with Echo Framework

![GO-ONLINE-STORE]

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Introduction

Welcome to the Online Store Application! This project utilizes the Echo framework for building a backend RESTful API that supports an online store where customers can interact with products, shopping carts, orders, and user accounts.

## Features

- **Customer Operations:**
  - View product list by category
  - Add products to shopping cart
  - View shopping cart contents
  - Delete products from shopping cart
  - Checkout and process payment transactions
  - Login and register customers

- **Product Management:**
  - CRUD operations for products
  - Product categorization and filtering

- **Order Management:**
  - Place orders
  - View order history
  - Update order status

- **Authentication and Authorization:**
  - User authentication (login/register)
  - JWT-based authentication and authorization

## Technologies Used

- **Programming Language:** Go (go 1.21.6) (Golang)
- **Framework:** Echo Framework (version 4)
- **Database:** MySQL
- **API Design:** RESTful API principles
- **Tools:** Git, Docker

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go programming language (version 1.21.6)
- Echo Framework (installation instructions [here](https://github.com/labstack/echo))
- MySQL database (version 8.3.0)
- Docker

### Installation

Clone the repository:

```bash
git clone https://github.com/sulfan-tech/go-online-store.git
cd go-online-store
```

Install dependencies:

```bash
go mod download || go mod tidy
```

### Configuration

1. Configure environment variables:
   - Copy `.env.example` to `.env` and configure database credentials, API keys, etc.

2. Initialize the database:
   - Run database migrations and seed initial data if applicable.

## Usage

1. Start the server:

```bash
go run cmd/api/main.go
```

2. Access the API endpoints at `http://localhost:port`.

## API Documentation

For detailed API documentation, refer to [API Documentation](https://sulfan.notion.site/create-an-online-store-application-API-d4aa504087334dc99740b357d9a8584e). Include detailed explanations of each endpoint, parameters, request bodies, and responses.

## Testing

Explain how to run automated tests for this system. For example:

```bash
go test ./...
```

## Deployment

Add additional notes about how to deploy this on a live system. For instance:

- Use Docker for containerization
- Deployment to cloud services like AWS, GCP, or Azure
- CI/CD pipeline setup

## Contributing

Provide guidelines for how others can contribute to and improve the project. For example:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

## License

This project is licensed under the [MIT License](LICENSE).


## Module Dependencies

Below are the dependencies listed in the `go.mod` file:

```
module go-online-store

go 1.21.6

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/validator/v10 v10.22.0
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/labstack/echo/v4 v4.12.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.24.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.10
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

## Checksum for `go.sum`

The checksums for the dependencies are listed in the `go.sum` file.
