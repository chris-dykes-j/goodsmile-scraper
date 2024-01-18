DROP TABLE IF EXISTS nendoroid, nendoroid_data;

CREATE TABLE nendoroid
(
    id SERIAL PRIMARY KEY,
    item_number INT
);

CREATE TABLE IF NOT EXISTS languages
(
  language_code CHAR(2) PRIMARY KEY, 
  language_name VARCHAR(255) NOT NULL
);

INSERT INTO languages (language_code, language_name)
VALUES ('en', 'English'), ('ja', 'Japanese'), ('zh', 'Chinese')
ON CONFLICT (language_code) DO NOTHING;

CREATE TABLE IF NOT EXISTS nendoroid_data
(
  nendoroid_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  item_link TEXT NOT NULL,
  blog_link TEXT NOT NULL,
  details JSONB NOT NULL,
  PRIMARY KEY (nendoroid_id, language_code),
  FOREIGN KEY (nendoroid_id) REFERENCES nendoroid (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);
