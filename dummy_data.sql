-- Insert dummy data into the User table
INSERT INTO User (Email, Password, RegistrationDate)
VALUES
  ('user1@example.com', 'hashed_password_1', '2023-11-10 12:00:00'),
  ('user2@example.com', 'hashed_password_2', '2023-11-10 12:30:00'),
  ('user3@example.com', 'hashed_password_3', '2023-11-10 13:00:00'),
  ('user4@example.com', 'hashed_password_4', '2023-11-10 13:30:00'),
  ('user5@example.com', 'hashed_password_5', '2023-11-10 14:00:00'),
  ('user6@example.com', 'hashed_password_6', '2023-11-10 14:30:00'),
  ('user7@example.com', 'hashed_password_7', '2023-11-10 15:00:00'),
  ('user8@example.com', 'hashed_password_8', '2023-11-10 15:30:00'),
  ('user9@example.com', 'hashed_password_9', '2023-11-10 16:00:00'),
  ('user10@example.com', 'hashed_password_10', '2023-11-10 16:30:00'),
  ('user11@example.com', 'hashed_password_11', '2023-11-10 17:00:00'),
  ('user12@example.com', 'hashed_password_12', '2023-11-10 17:30:00'),
  ('user13@example.com', 'hashed_password_13', '2023-11-10 18:00:00'),
  ('user14@example.com', 'hashed_password_14', '2023-11-10 18:30:00'),
  ('user15@example.com', 'hashed_password_15', '2023-11-10 19:00:00'),
  ('user16@example.com', 'hashed_password_16', '2023-11-10 19:30:00'),
  ('user17@example.com', 'hashed_password_17', '2023-11-10 20:00:00'),
  ('user18@example.com', 'hashed_password_18', '2023-11-10 20:30:00'),
  ('user19@example.com', 'hashed_password_19', '2023-11-10 21:00:00'),
  ('user20@example.com', 'hashed_password_20', '2023-11-10 21:30:00');

-- Insert dummy data into the Account table
INSERT INTO Account (UserID, AccountName, Balance)
VALUES
  (1, 'Savings Account', 5000.00),
  (1, 'Checking Account', 2500.00),
  (2, 'Investment Account', 10000.00),
  (2, 'Credit Card', -500.00),
  (3, 'Emergency Fund', 3000.00),
  (3, 'Travel Fund', 2000.00),
  (4, 'Business Account', 15000.00),
  (4, 'Personal Loan', -2000.00),
  (5, 'Retirement Savings', 8000.00),
  (5, 'Mortgage', -1200.00),
  (6, 'Vacation Fund', 1000.00),
  (6, 'Car Loan', -3000.00),
  (7, 'Education Fund', 4000.00),
  (7, 'Utility Bills', -600.00),
  (8, 'Health Savings', 2000.00),
  (8, 'Shopping Budget', -300.00),
  (9, 'Gifts Fund', 500.00),
  (9, 'Rent', -1000.00),
  (10, 'Insurance Premiums', -700.00);

-- Insert dummy data into the Transaction table
INSERT INTO Transaction (tdate, pdate, from_account, to_account, context, tags, amount)
VALUES
  ('2023-11-10 13:00:00', '2023-11-10 13:30:00', 1, 2, 'Grocery shopping', 'Groceries', -100.00),
  ('2023-11-10 14:00:00', '2023-11-10 14:30:00', 2, 1, 'Paycheck deposit', 'Income', 1500.00),
  ('2023-11-10 15:00:00', '2023-11-10 15:30:00', 3, 1, 'Stock purchase', 'Investment', -500.00),
  ('2023-11-10 16:00:00', '2023-11-10 16:30:00', 4, 2, 'Credit card payment', 'Credit Card', -200.00),
  ('2023-11-10 17:00:00', '2023-11-10 17:30:00', 5, 1, 'Contribution to retirement savings', 'Retirement Savings', -500.00),
  ('2023-11-10 18:00:00', '2023-11-10 18:30:00', 6, 2, 'Car loan payment', 'Car Loan', -300.00),
  ('2023-11-10 19:00:00', '2023-11-10 19:30:00', 7, 1, 'Utility bill payment', 'Utility Bills', -100.00),
  ('2023-11-10 20:00:00', '2023-11-10 20:30:00', 8, 2, 'Shopping expense', 'Shopping Budget', -50.00),
  ('2023-11-10 21:00:00', '2023-11-10 21:30:00', 9, 1, 'Rent payment', 'Rent', -1000.00),
  ('2023-11-10 22:00:00', '2023-11-10 22:30:00', 10, 2, 'Insurance premium payment', 'Insurance Premiums', -200.00),
  ('2023-11-10 23:00:00', '2023-11-10 23:30:00', 11, 1, 'Vacation fund contribution', 'Vacation Fund', -100.00),
  ('2023-11-11 00:00:00', '2023-11-11 00:30:00', 12, 2, 'Education fund contribution', 'Education Fund', -200.00),
  ('2023-11-11 01:00:00', '2023-11-11 01:30:00', 13, 1, 'Business income', 'Quarterly Divended', -200.00);
