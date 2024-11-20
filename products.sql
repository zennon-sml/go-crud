-- Create the products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,    -- Auto-incrementing primary key
    name VARCHAR(255) NOT NULL,  -- Product name, cannot be null
    price DECIMAL(10, 2) NOT NULL -- Product price with 2 decimal places
);

-- Insert some sample data into the products table
INSERT INTO products (name, price) VALUES
    ('Laptop', 1200.50),
    ('Smartphone', 799.99),
    ('Headphones', 150.00),
    ('Monitor', 300.25),
    ('Keyboard', 45.99);

-- Query the table to verify the data
SELECT * FROM products;
