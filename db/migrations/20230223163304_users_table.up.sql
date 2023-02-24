CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(100) UNIQUE,
  username VARCHAR(50) UNIQUE,
  hashed_password VARCHAR(200)
);