DROP TABLE IF EXISTS nendoroid, nendoroid_name, nendoroid_description, nendoroid_details, nendoroid_link, nendoroid_blog_link;

CREATE TABLE IF NOT EXISTS nendoroid
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

CREATE TABLE IF NOT EXISTS nendoroid_name
(
  nendoroid_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (nendoroid_id, language_code),
  FOREIGN KEY (nendoroid_id) REFERENCES nendoroid (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);

CREATE TABLE IF NOT EXISTS nendoroid_description
(
  nendoroid_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (nendoroid_id, language_code),
  FOREIGN KEY (nendoroid_id) REFERENCES nendoroid (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);

CREATE TABLE IF NOT EXISTS nendoroid_details
(
  nendoroid_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  details JSONB,
  PRIMARY KEY (nendoroid_id, language_code),
  FOREIGN KEY (nendoroid_id) REFERENCES nendoroid (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);

CREATE TABLE IF NOT EXISTS nendoroid_link
(
  nendoroid_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (nendoroid_id, language_code),
  FOREIGN KEY (nendoroid_id) REFERENCES nendoroid (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);

CREATE TABLE IF NOT EXISTS nendoroid_blog_link
(
  nendoroid_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (nendoroid_id, language_code),
  FOREIGN KEY (nendoroid_id) REFERENCES nendoroid (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);
