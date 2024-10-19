# Webshop

This is a webshop project with a **Go** backend and a **Vue.js** frontend. The backend uses a **MariaDB** database, which is run using Docker.

## Backend

### Prerequisites
- **Go** installed
- **Docker** and **Docker Compose** installed

### Running the Backend

1. Navigate to the `cmd/webshop` directory:
    ```bash
    cd cmd/webshop
    ```

2. Run the backend:
    ```bash
    go run main.go
    ```

### Setting Up MariaDB Database

1. Navigate to the backend's `docker` directory:
    ```bash
    cd backend/docker
    ```

2. Start MariaDB with Docker Compose in detached mode:
    ```bash
    docker-compose up -d
    ```

3. Ensure the MariaDB container is running:
    ```bash
    docker start backend_mariadb_1
    ```

4. Access the running container:
    ```bash
    docker exec -it backend_mariadb_1 bash
    ```

5. Connect to the MariaDB instance:
    ```bash
    mariadb -u admin -p
    ```

    Use the password you defined in the `docker-compose.yml` file.

### Example: Database Queries

Once connected to MariaDB, you can execute the following SQL queries:

- List all databases:
    ```sql
    SHOW DATABASES;
    ```

- Switch to the `webshop` database:
    ```sql
    USE webshop;
    ```

- Show the structure of the `product` table:
    ```sql
    DESCRIBE product;
    ```

- List all products:
    ```sql
    SELECT * FROM product;
    ```

## Frontend

### Prerequisites
- **Node.js** and **npm** installed

### Project Setup

1. Navigate to the `frontend` directory:
    ```bash
    cd frontend
    ```

2. Install the dependencies:
    ```bash
    npm install
    ```

### Running the Development Server

1. Start the development server with hot-reloading:
    ```bash
    npm run serve
    ```

2. Open [http://localhost:8080](http://localhost:8080) in your browser to view the frontend.

### Compiling for Production

1. Build the project for production:
    ```bash
    npm run build
    ```

    This will create a minified and optimized version of the project in the `dist` directory.

### Linting and Fixing Files

1. Lint and fix code issues:
    ```bash
    npm run lint
    ```

---
