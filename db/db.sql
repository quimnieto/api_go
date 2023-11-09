-- Si no existe la base de datos 'api_go_db', la crea
CREATE DATABASE IF NOT EXISTS api_go_db;
USE api_go_db;

-- Si no existe la tabla 'courses', la crea
CREATE TABLE IF NOT EXISTS courses (
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
