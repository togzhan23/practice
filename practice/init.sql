CREATE TABLE crop_products (
    product_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price_per_unit NUMERIC(10,2),
    categories TEXT[],
    certifications TEXT[],
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE crop_components (
    product_id INT,
    component_id INT,
    quantity NUMERIC(10,2),
    PRIMARY KEY (product_id, component_id),
    FOREIGN KEY (product_id) REFERENCES crop_products(product_id) ON DELETE CASCADE,
    FOREIGN KEY (component_id) REFERENCES crop_products(product_id) ON DELETE CASCADE
);

CREATE TABLE farm_supply_items (
    supply_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    unit TEXT,
    quantity NUMERIC(10,2),
    category TEXT,
    location TEXT,
    metadata JSONB,
    last_restocked TIMESTAMP
);

CREATE TABLE customer_orders (
    order_id SERIAL PRIMARY KEY,
    client_name TEXT NOT NULL,
    contact_info TEXT,
    status TEXT CHECK (status IN ('в обработке', 'доставлен', 'отменен')) DEFAULT 'в обработке',
    order_date DATE DEFAULT CURRENT_DATE,
    delivery_date DATE,
    total_amount NUMERIC(12,2)
);

CREATE TABLE order_items (
    order_id INT,
    product_id INT,
    quantity NUMERIC(10,2),
    price_per_unit NUMERIC(10,2),
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES customer_orders(order_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES crop_products(product_id)
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT CHECK (role IN ('admin', 'manager', 'worker')),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE activity_logs (
    log_id SERIAL PRIMARY KEY,
    user_id INT,
    action TEXT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
