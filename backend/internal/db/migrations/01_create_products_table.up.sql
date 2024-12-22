CREATE TABLE products (
    id UUID PRIMARY KEY,
    aliExpressID INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);