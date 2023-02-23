CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(100),
  username VARCHAR(50),
  hashed_password VARCHAR(200)
);