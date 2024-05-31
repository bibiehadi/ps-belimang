CREATE TABLE ordered_items (
    id SERIAL PRIMARY KEY,
    ordered_merchants_id INT NOT NULL,
    item_id INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ordered_merchants_id) REFERENCES ordered_merchants(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES merchant_items(id) ON DELETE CASCADE
    );