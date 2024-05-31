CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    est_delivery_time INT,
    total_distance FLOAT,
    total_deliveryfee BIGINT,
    total_purchase BIGINT,
    total_order BIGINT,
    status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES auths(id) ON DELETE CASCADE
    );