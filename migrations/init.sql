CREATE TABLE IF NOT EXISTS submissions(
  id UUID PRIMARY KEY,
  body TEXT,
  author TEXT,
  type INT,
  created_at TIMESTAMP
);

CREATE INDEX submissions_type
ON submissions (type); 

ALTER TABLE submissions
ADD CONSTRAINT unique_body_author
UNIQUE (body,author);
