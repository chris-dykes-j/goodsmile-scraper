-- TODO Figure out the id crap. Then finish the insert statement.

INSERT INTO nendoroid (nendoroid_number) VALUES ($1);
INSERT INTO nendoroid_name (

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
