CREATE DATABASE IF NOT EXISTS eCombombDB;

USE eCombombDB;

CREATE TABLE
    users (
        id INT NOT NULL AUTO_INCREMENT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        profile_image LONGTEXT NULL,
        PRIMARY KEY (id),
        UNIQUE (username)
    );

CREATE TABLE
    products (
        id INT NOT NULL AUTO_INCREMENT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL,
        name VARCHAR(255) NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        description LONGTEXT,
        user_id INT NOT NULL,
        product_image LONGTEXT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    carts (
        id INT NOT NULL AUTO_INCREMENT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL,
        user_id INT NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    cart_items (
        id INT NOT NULL AUTO_INCREMENT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL,
        cart_id INT NOT NULL,
        product_id INT NOT NULL,
        quantity INT NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (cart_id) REFERENCES carts (id),
        FOREIGN KEY (product_id) REFERENCES products (id)
    );

CREATE TABLE
    orders (
        id INT NOT NULL AUTO_INCREMENT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL,
        PRIMARY KEY (id),
        user_id INT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    order_items (
        id INT NOT NULL AUTO_INCREMENT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL,
        order_id INT NOT NULL,
        product_id INT NOT NULL,
        quantity INT NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (order_id) REFERENCES orders (id),
        FOREIGN KEY (product_id) REFERENCES products (id)
    );