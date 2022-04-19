CREATE TABLE types (
    id integer auto_increment NOT NULL,
	description varchar(100) NOT NULL,
	CONSTRAINT types_pk PRIMARY KEY (id)
);

insert into types (id, description) values
(1, "Monstruo")
,(2, "Magica")
,(3, "Trampa");

CREATE TABLE subtypes (
    id integer auto_increment NOT NULL,
	description varchar(100) NOT NULL,
	type_id integer NOT NULL,
	CONSTRAINT subtypes_pk PRIMARY KEY (id),
	CONSTRAINT subtypes_FK FOREIGN KEY (type_id) REFERENCES types(id)
);

insert into subtypes (id, description, type_id) values
(1, "Monstruos Normales", 1)
,(2, "Monstruos de Efecto", 1)
,(3, "Monstruos de Ritual", 1)
,(4, "Cartas Magicas de Juego Rapido", 2)
,(5, "Cartas de Trampa de Contraefecto", 3);

CREATE TABLE images (
    id integer auto_increment NOT NULL,
	description varchar(100) NOT NULL,
	url varchar(250) NOT NULL,
	CONSTRAINT images_pk PRIMARY KEY (id)
);

insert into images (id, description, url) values
(1, "Mago Oscuro", "http://test.com")
,(2, "Jinzo", "http://test.com")
,(3, "Exodia, el prohibido", "http://test.com")
,(4, "Olla de la Codicia", "http://test.com")
,(5, "Dragon Negro de Ojos Rojos", "http://test.com");

CREATE TABLE cards (
    id integer auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	first_edition tinyint(4) DEFAULT 0 NULL,
	serie_code varchar(20) NOT NULL,
	subtype_id integer NOT NULL,
	atk integer NULL,
	def integer NULL,
	stars integer NULL,
	description varchar(250) NULL,
	price float NOT NULL,
	image_id integer NOT NULL,
	date_created date NOT NULL,
	CONSTRAINT cards_pk PRIMARY KEY (id),
	CONSTRAINT cards_fk FOREIGN KEY (subtype_id) REFERENCES subtypes(id),
	CONSTRAINT cards_fk_1 FOREIGN KEY (image_id) REFERENCES images(id)
);

insert into cards (id, name, first_edition, serie_code, subtype_id, atk, def, stars, description, price, image_id, date_created) values
(1, "Mago Oscuro", 1, "46986414", 1, 2500, 2100, 7, "El más grande de los magos en cuanto al ataque y la defensa", 100, 1, "2002-03-08")
,(2, "Jinzo", 1, "77585513", 2, 2400, 1500, 6, "No se pueden activar Cartas de Trampa, ni sus efectos en el Campo. Niega todos los efectos de Trampas en el Campo", 100, 2, "2002-10-20")
,(3, "Exodia, el Prohibido", 1, "33396948", 2, 1000, 1000, 3, "Si tienes Pierna Derecha del Prohibido, Pierna Izquierda del Prohibido, Brazo Derecho del Prohibido y Brazo Izquierdo del Prohibido, además de esta carta en tu mano, ganas el Duelo", 100, 3, "2002-03-08");
