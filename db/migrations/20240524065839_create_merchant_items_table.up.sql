CREATE TABLE merchant_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    product_category VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    merchant_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id) ON DELETE CASCADE
    );