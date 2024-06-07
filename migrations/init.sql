CREATE TABLE IF NOT EXISTS submissions(
  id UUID PRIMARY KEY,
  body TEXT,
  author TEXT,
  type INT,
  created_at TIMESTAMP
);

CREATE INDEX submissions_type
ON submissions (type); 
