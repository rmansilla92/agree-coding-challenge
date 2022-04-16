create table types (
    id integer not null auto,
    description varchar(100) not null,
    primary key (id)
)

insert into types (description) values
("Monstruo")
,("Magica")
,("Trampa") 

create table subtypes (
    id integer not null auto,
    description varchar(100) not null,
    type_id integer not null
    primary key (id)
    constraint subtypes_fk_1 foreign key (type_id) references types(id)
)

insert into subtypes (description, type_id) values
("Monstruos Normales", 1)
,("Monstruos de Efecto", 1)
,("Monstruos de Ritual", 1)
,("Cartas Magicas de Juego Rapdio", 2)
,("Cartas de Trampa de Contraefecto", 3)

create table cards (
    id integer not null auto,
    name varchar(100) not null,
    first_edition tinyint(4) default 0,
    serie_code varchar(20) not null,
    subtype_id integer not null, 
    atk integer default null,
    def integer default null,
    stars integer default null,
    description varchar(250) default null,
    price integer not null,
    image varchar(100) not null,
    date_created date not null,
    PRIMARY KEY (id)
    CONSTRAINT card_fk_1 FOREIGN KEY (subtype_id) REFERENCES subtypes(id) 
)

insert into cards (name, first_edition, serie_code, subtype_id, atk, def, stars, description, price, image, date_created) values
("Mago Oscuro", 1, "45086414", 1, 2500, 2100, 7, "El mas grande de los magos en relacion con el ataque y defensa", 100, "HIY-3006", "2022-04-16")
