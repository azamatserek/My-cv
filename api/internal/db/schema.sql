CREATE TABLE IF NOT EXISTS profile (
  id SERIAL PRIMARY KEY,
  full_name TEXT NOT NULL,
  title TEXT,
  email TEXT,
  phone TEXT,
  location TEXT,
  about TEXT,
  avatar_url TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS experience (
  id SERIAL PRIMARY KEY,
  company TEXT NOT NULL,
  role TEXT NOT NULL,
  start_date DATE NOT NULL,
  end_date DATE,
  location TEXT,
  description TEXT,
  highlights TEXT[],
  order_index INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS education (
  id SERIAL PRIMARY KEY,
  institution TEXT NOT NULL,
  degree TEXT,
  field TEXT,
  start_year INT,
  end_year INT,
  details TEXT,
  order_index INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS publication (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  venue TEXT,
  year INT,
  authors TEXT[],
  doi TEXT,
  url TEXT,
  abstract TEXT,
  order_index INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS achievement (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  issuer TEXT,
  date DATE,
  description TEXT,
  url TEXT,
  order_index INT DEFAULT 0
);