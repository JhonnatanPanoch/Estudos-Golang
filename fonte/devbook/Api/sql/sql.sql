CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS "usuarios";

CREATE TABLE "usuarios" (
	"id"	INTEGER NOT NULL UNIQUE,
	"nome"	varchar(50) NOT NULL,
	"nick"	varchar(50) NOT NULL UNIQUE,
	"email"	varchar(50) NOT NULL UNIQUE,
	"senha"	varchar(50) NOT NULL,
	"criadoEm"	datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id" AUTOINCREMENT)
)