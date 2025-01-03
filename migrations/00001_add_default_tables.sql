-- +goose Up
CREATE TABLE Users (
    ID SERIAL PRIMARY KEY,
    Username VARCHAR(50) NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL,
    IsAdmin BOOLEAN DEFAULT FALSE,
    IsBanned BOOLEAN DEFAULT FALSE,
    BannedReason VARCHAR(255),

    UNIQUE (Username)
);

CREATE INDEX idx_user_username ON Users (Username);

CREATE TABLE Products (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Price DECIMAL(10, 2) NOT NULL,
    ProcessName VARCHAR(255) NOT NULL,

    UNIQUE (Name)
);

CREATE INDEX idx_product_name ON Products (Name);

CREATE TABLE Logs (
    ID SERIAL PRIMARY KEY,
    Username VARCHAR(50) NOT NULL,
    Action VARCHAR(50) NOT NULL,
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Promocodes (
    ID SERIAL PRIMARY KEY,
    Code VARCHAR(50) NOT NULL,
    Activations DECIMAL(10, 2) NOT NULL,

    UNIQUE (Code)
);

-- +goose Down
DROP INDEX IF EXISTS idx_user_username;
DROP TABLE IF EXISTS Users;
DROP INDEX IF EXISTS idx_product_name;
DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Logs;
DROP TABLE IF EXISTS Promocodes;