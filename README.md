# Invoice Management System

An invoice management system to manage customer data, invoices, and items.

## Table of Contents
- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Invoice Management System is a simple application to manage customer data, create invoices, and track items.

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go (version 1.20.x)
- MySQL (version 5.7.x)

### Installation

build the application:

```bash
cd esb-invoice
go build

```
## Usage

```bash
./esb-invoice
cd cmd
go run main.go
```

## Project Structure

```
esb-invoice
├── cmd
│   ├── main.go
├── docs
├── internal
│   ├── app
│   │   ├── config
│   │   │   ├── db
│   │   │   │   ├── gormmysql.go
│   │   ├── controller
│   │   │   ├── customer_controller.go
│   │   │   ├── invoice_controller.go
│   │   │   ├── item_controller.go
│   │   │   ├── type_controller.go
│   │   ├── handler
│   │   │   ├── filters
│   │   │   │   ├── customer_filter.go
│   │   │   │   ├── invoice_filter.go
│   │   │   │   ├── item_filter.go
│   │   │   ├── helpers
│   │   │   │   ├── globalhelper.go
│   │   │   │   ├── returnjson.go
│   │   │   │   ├── validation.go
│   │   │   ├── request
│   │   │   │   ├── customer_request.go
│   │   │   │   ├── invoice_request.go
│   │   │   │   ├── item_request.go
│   │   │   │   ├── type_request.go
│   │   │   ├── response
│   │   │   │   ├── customer_response.go
│   │   │   │   ├── global_response.go
│   │   │   │   ├── invoice_response.go
│   │   │   │   ├── item_response.go
│   │   │   │   ├── type_response.go
│   │   ├── middleware
│   │   ├── router
│   │   │   ├── router.go
│   ├── domain
│   │   ├── model
│   │   │   ├── customer.go
│   │   │   ├── invoice.go
│   │   │   ├── invoiceitem.go
│   │   │   ├── item.go
│   │   │   ├── type.go
│   │   ├── repository
│   │   │   ├── gorm
│   │   │   │   ├── customer.go
│   │   │   │   ├── invoice.go
│   │   │   │   ├── invoiceitem.go
│   │   │   │   ├── item.go
│   │   │   │   ├── type.go
│   │   │   ├── repositories.go
│   │   ├── service
│   │   │   ├── customer_service.go
│   │   │   ├── invoice_service.go
│   │   │   ├── invoiceitem_service.go
│   │   │   ├── type_service.
├── migrations
│   ├── esb-invoice.sql
├── seeders
│   ├── seeder.go
├── testing
│   ├── customer_test.go
│   ├── invoice_test.go
│   ├── item_test.go
│   ├── type_test.go
├── .env
├── .env_example
├── .gitignore
├── go.mod
├── go.sum
├── README.md
```

## Configuration


