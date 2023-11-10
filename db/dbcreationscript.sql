-- Create the database
CREATE DATABASE IF NOT EXISTS BUDGETANDTRACKING;
USE BUDGETANDTRACKING;

-- Create User table
CREATE TABLE User (
  UserID INT PRIMARY KEY AUTO_INCREMENT,
  Email VARCHAR(255) NOT NULL,
  Password VARCHAR(255) NOT NULL,
  RegistrationDate DATETIME NOT NULL
);

-- Create Account table
CREATE TABLE Account (
  AccountID INT PRIMARY KEY AUTO_INCREMENT,
  UserID INT,
  AccountName VARCHAR(50) NOT NULL,
  Balance DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (UserID) REFERENCES User(UserID)
);

-- Create Transaction table
CREATE TABLE Transaction (
  ID INT PRIMARY KEY AUTO_INCREMENT,
  tdate DATETIME NOT NULL,
  pdate DATETIME NOT NULL,
  from_account INT,
  to_account INT,
  context VARCHAR(255),
  tags VARCHAR(255),
  amount DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (from_account) REFERENCES Account(AccountID),
  FOREIGN KEY (to_account) REFERENCES Account(AccountID)
);

-- Create Reminder table
CREATE TABLE Reminder (
  ReminderID INT PRIMARY KEY AUTO_INCREMENT,
  UserID INT,
  AccountID INT,
  ReminderDate DATETIME NOT NULL,
  Status VARCHAR(20) NOT NULL,
  FOREIGN KEY (UserID) REFERENCES User(UserID),
  FOREIGN KEY (AccountID) REFERENCES Account(AccountID)
);
