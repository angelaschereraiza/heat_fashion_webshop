# Webshop

## Backend
* cd backend/cmd/h
* go run main.go

### Connect to MariaDB
* cd backend/docker
* docker-compose up -d
* docker start backend_mariadb_1
* docker exec -it backend_mariadb_1 bash
* mariadb -u admin -p

#### Example how to use
* SHOW DATABASES;
* USE heatup_fashion;
* DESCRIBE product;
* SELECT * FROM product;

## Frontend

### Project setup
```
npm install
```

#### Compiles and hot-reloads for development
```
npm run serve
```

#### Compiles and minifies for production
```
npm run build
```

#### Lints and fixes files
```
npm run lint
```
