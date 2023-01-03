DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS linked;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS infaq;
DROP TABLE IF EXISTS donate;
CREATE TABLE category (
  slug TEXT PRIMARY KEY,
  name TEXT
);
CREATE TABLE post (
  slug TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  cdate TEXT,
  content TEXT,
  image TEXT,
  views INTEGER DEFAULT(0)
);
CREATE TABLE linked(
  category,
  post,
  PRIMARY KEY(category,post),
  FOREIGN KEY(category) REFERENCES category(slug),
  FOREIGN KEY(post) REFERENCES post(slug)
);
CREATE TABLE comments (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  content TEXT,
  photo TEXT,
  post TEXT,
  cdate INT,
  FOREIGN KEY(post) REFERENCES post(slug)
);
CREATE TABLE infaq (
  post TEXT PRIMARY KEY,
  total INTEGER DEFAULT(0),
  now INTEGER DEFAULT(0),
  FOREIGN KEY(post) REFERENCES post(slug)
);
CREATE TABLE donate (
  post TEXT PRIMARY KEY,
  cdate TEXT,
  name TEXT,
  amount INTEGER DEFAULT(0),
  FOREIGN KEY(post) REFERENCES post(slug)
);
