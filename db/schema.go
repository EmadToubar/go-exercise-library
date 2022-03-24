package db

var schema = `

CREATE TABLE IF NOT EXISTS person (
    personID SERIAL PRIMARY KEY,
	first_name text,
    last_name text,
    email text
);


CREATE TABLE IF NOT EXISTS book (
	title text,
	isbn text PRIMARY KEY,
	status boolean
	);

CREATE TABLE IF NOT EXISTS person_book (
	loan_isbn text,
	loan_last_name text,
	FOREIGN KEY (loan_isbn) REFERENCES book(isbn) ON DELETE CASCADE,
	FOREIGN KEY (loan_last_name) REFERENCES person(last_name) ON DELETE CASCADE

			);
	
	`
