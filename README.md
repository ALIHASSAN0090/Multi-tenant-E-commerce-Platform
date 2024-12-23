# Multi-tenant-E-commerce-Platform 
A scalable e-commerce platform that allows clients to create their own online stores. Users can browse different stores, search for items, and make purchases. The application is built using Go, with a focus on clean architecture and efficient command-line management using Cobra.

## Features

- **Multi-Store Support**: Clients can create and manage their own stores.
- **User Management**: Users can sign up, browse stores, and purchase items.
- **Admin Panel**: Admins can manage users, stores, and orders.
- **Role-Based Access Control (RBAC)**: Secure routes based on user roles (Admin, Client, User).
- **Command-Line Interface**: Manage the application using Cobra commands.
- **Scalable Architecture**: Designed for scalability and maintainability.

## Architecture

- **Router**: Handles HTTP requests and routes them to the appropriate controllers.
- **Controllers**: Manage the business logic for different parts of the application.
- **Services**: Provide reusable functionalities like authentication and validation.
- **Database**: PostgreSQL is used for data storage, with schema defined in `db_migrator/schema.sql`.

## **Schema Overview**
![Sample Image](db_migratorions/migration_files/schema.png)  
*(This image represents the PostgreSQL schema structure used for the project.)*

## Role-Based Access Control (RBAC)

The application uses RBAC to manage access to different routes based on user roles:

- **Admin**: Full access to manage users, stores, and orders.
- **Client**: Access to manage their own store and view orders.
- **User**: Access to browse stores and make purchases.

### Implementing RBAC

RBAC is implemented using middleware that checks the user's role before allowing access to specific routes. This ensures that only authorized users can perform certain actions.

## Getting Started

### Prerequisites

- Go 1.18 or later
- PostgreSQL
- Make

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/ALIHASSAN0090/Multi-tenant-E-commerce-Platform.git
   cd Multi-tenant-E-commerce-Platform
   ```

2. **Set up the database**:
   - Ensure PostgreSQL is running.
   - Create a database and run the schema from `db_migrator/schema.sql`.

3. **Build the application**:
   ```bash
   make build
   ```

4. **Run the application**:
   ```bash
   make api
   ```

### Usage

- **Health Check**: Verify the application is running at `/health`.

## Future Plans

- **Docker**: Containerize the application for easier deployment.
- **Kubernetes**: Orchestrate the application using Kubernetes for scalability.
- **AWS Deployment**: Deploy the application on AWS for production use.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## Contact

For questions or support,  contact [me](mailto:alihassankhan285@gmail.com).
