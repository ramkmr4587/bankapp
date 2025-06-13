-- seed.sql

-- Insert Users
INSERT INTO users (username, password) VALUES
  ('alice', '$2a$10$EXAMPLE_HASHED_PWD_ALICE'),
  ('bob',   '$2a$10$EXAMPLE_HASHED_PWD_BOB'),
  ('carol', '$2a$10$EXAMPLE_HASHED_PWD_CAROL');

-- Insert Accounts (balance precision DECIMAL(19,4))
INSERT INTO accounts (user_id, balance) VALUES
  (1, 1000.0000),
  (2,  500.5000),
  (3,    0.0000);

-- Seed Entries (ledger records aligned with accounts)
INSERT INTO entries (account_id, amount) VALUES
  (1, 1000.0000),  -- Alice initial deposit
  (2,  500.5000);  -- Bob initial deposit

-- Seed Transfers (example of a transaction)
INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES
  (1, 2, 200.0000),  -- Alice -> Bob
  (2, 3, 100.5000);  -- Bob -> Carol
