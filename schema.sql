-- Using MySQL

CREATE DATABASE IF NOT EXISTS BEApp;
USE BEApp;

CREATE TABLE IF NOT EXISTS transactions (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	tdate DATETIME NOT NULL,
	pdate DATETIME NOT NULL,
	from_account VARCHAR(100) NOT NULL, -- account paying (could be a personal bank account)
	to_account VARCHAR(100) NOT NULL, -- account earning (could be Amazon)
	context TEXT NOT NULL,
	tags VARCHAR(100) NOT NULL
);

CREATE INDEX idx_from_account ON transactions (from_account);
CREATE INDEX idx_to_account ON transactions (to_account);
