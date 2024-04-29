# Restaurant Management System

This is a restaurant management system built with Go and MongoDB, using the Gin framework for handling HTTP requests.

## Features

- User management: Sign up, login, and retrieve user information.
- Food management: CRUD operations for managing food items.
- Menu management: CRUD operations for managing menus.
- Table management: CRUD operations for managing tables.
- Order management: CRUD operations for managing orders.
- Order item management: CRUD operations for managing order items.
- Invoice management: CRUD operations for managing invoices.

## Installation

1. Clone this repository:

    ```
    git clone https://github.com/MISHRA-TUSHAR/restaurant-management-backend.git
    ```

2. Navigate to the project directory:

    ```
    cd restaurant-management-backend
    ```

4. Install dependencies:

    ```
    go mod tidy
    ```

5. Set up your MongoDB database and update the connection string in the code.

6. Run the application:

    ```
    go run main.go
    ```

## Usage

- The application runs on port 8000 by default. You can access the API endpoints using tools like cURL or Postman.
- Refer to the routes in the routes directory.
