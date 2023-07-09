-- Create the products table
CREATE TABLE products (
  product_id INT PRIMARY KEY,
  product_name VARCHAR(255),
  product_type VARCHAR(255),
  product_description VARCHAR(255),
  operator VARCHAR(255),
  payment_methods VARCHAR(255)
);

-- Create the customers table
CREATE TABLE customers (
  customer_id INT PRIMARY KEY,
  name VARCHAR(255),
  address VARCHAR(255),
  date_of_birth DATE,
  status_user VARCHAR(255),
  gender VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- Create the transactions table
CREATE TABLE transactions (
  transaction_id INT PRIMARY KEY,
  customer_id INT,
  transaction_date DATE,
  total_amount DECIMAL(10, 2),
  FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

-- Create the transaction_details table
CREATE TABLE transaction_details (
  transaction_detail_id INT PRIMARY KEY,
  transaction_id INT,
  product_id INT,
  quantity INT,
  price DECIMAL(10, 2),
  FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id),
  FOREIGN KEY (product_id) REFERENCES products(product_id)
);
