CREATE TABLE nendoroid (
    id SERIAL PRIMARY KEY,
);

CREATE TABLE IF NOT EXISTS languages
(
  language_code CHAR(2) PRIMARY KEY, 
  language_name VARCHAR(255) NOT NULL
);

INSERT INTO languages (language_code, language_name)
VALUES ('en', 'English'), ('ja', 'Japanese'), ('zh', 'Chinese');

CREATE TABLE IF NOT EXISTS nendoroid_name
(
  figure_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (figure_id, language_code),
  FOREIGN KEY (figure_id) REFERENCES figure (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);

CREATE TABLE IF NOT EXISTS nendoroid_description
(
  figure_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (figure_id, language_code),
  FOREIGN KEY (figure_id) REFERENCES figure (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
)

CREATE TABLE IF NOT EXISTS nendoroid_details
(
  figure_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  details JSONB,
  PRIMARY KEY (figure_id, language_code),
  FOREIGN KEY (figure_id) REFERENCES figure (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
);

CREATE TABLE IF NOT EXISTS nendoroid_link
(
  figure_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (figure_id, language_code),
  FOREIGN KEY (figure_id) REFERENCES figure (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
)

CREATE TABLE IF NOT EXISTS nendoroid_blog_link
(
  figure_id INT NOT NULL,
  language_code CHAR(2) NOT NULL,
  text VARCHAR(255) NOT NULL,
  PRIMARY KEY (figure_id, language_code),
  FOREIGN KEY (figure_id) REFERENCES figure (id),
  FOREIGN KEY (language_code) REFERENCES languages (language_code)
)

/*
// Input Data Structs
type OldFigure struct {
	English  Data `json:"en"`
	Japanese Data `json:"ja"`
	Chinese  Data `json:"zh"`
}

type Data struct {
	ItemNumber  string    `json:"itemNumber"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ItemLink    string    `json:"itemLink"`
	BlogLink    string    `json:"blogLink"`
	Details     []Details `json:"details"`
}

type Details struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
*/
