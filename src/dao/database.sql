--Database: super_hero
CREATE DATABASE super_hero;

-- Table: heroes
CREATE TABLE heroes 
(	uuid SERIAL PRIMARY KEY,
	id INTEGER,
   	name VARCHAR(50) UNIQUE NOT NULL, 
	fullname VARCHAR(255) UNIQUE NOT NULL, 
	intelligence INTEGER NOT NULL, 
	power INTEGER NOT NULL,
	occupation INTEGER NOT NULL,
	image VARCHAR(255) NOT NULL, 
	groups VARCHAR(255) NOT NULL,
	relatives VARCHAR(255) NOT NULL,
	numrelatives INTEGER
);