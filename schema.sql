-- Using MySQL

CREATE DATABASE IF NOT EXISTS BEApp;
USE BEApp;

-- Create users table
CREATE TABLE users (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  creation DATETIME NOT NULL,
  email VARCHAR(255) NOT NULL,
  pass_hash VARCHAR(255) NOT NULL
);

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	tdate DATETIME NOT NULL, -- actual date of transaction
	pdate DATETIME NOT NULL, -- posted date of transaction
	from_account VARCHAR(100) NOT NULL, -- account paying (could be a personal bank account)
	to_account VARCHAR(100) NOT NULL, -- account earning (could be Amazon)
	amount DECIMAL(10, 2) NOT NULL,
	context TEXT NOT NULL,
	tags VARCHAR(100) NOT NULL
);

CREATE INDEX idx_from_account ON transactions (from_account);
CREATE INDEX idx_to_account ON transactions (to_account);
