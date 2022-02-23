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
);


DROP TABLE IF EXISTS "seguidores";
CREATE TABLE "seguidores" (
	"usuario_id"	INTEGER NOT NULL,
	"seguidor_id"	INTEGER NOT NULL,
	FOREIGN KEY("seguidor_id") REFERENCES "usuarios"("id") ON DELETE CASCADE,
	FOREIGN KEY("usuario_id") REFERENCES "usuarios"("id") ON DELETE CASCADE,
	PRIMARY KEY("usuario_id","seguidor_id")
);

CREATE TABLE "publicacoes" (
	"id"	INTEGER NOT NULL,
	"titulo"	varchar(50) NOT NULL,
	"conteudo"	varchar(300) NOT NULL,
	"autor_id"	INTEGER NOT NULL,
	"curtidas"	INTEGER NOT NULL DEFAULT 0,
	"criadaEm"	datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY("autor_id") REFERENCES "usuarios"("id") ON DELETE CASCADE,
	PRIMARY KEY("id" AUTOINCREMENT)
);