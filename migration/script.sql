create table genre
(
    id   serial
        primary key,
    name varchar(100) not null
);

alter table genre
    owner to postgres;

create table country
(
    id   serial
        primary key,
    name varchar(100) not null
);

alter table country
    owner to postgres;

create table person
(
    id   serial
        primary key,
    name varchar(200) not null
);

alter table person
    owner to postgres;

create table post
(
    id   serial
        primary key,
    name varchar(200) not null
);

alter table post
    owner to postgres;

create table movie
(
    id           serial
        primary key,
    name         varchar(200) not null,
    budget       integer      not null,
    box_office   integer      not null,
    release      date,
    duration     time,
    synopsis     text,
    rating       real,
    rating_num   integer,
    trailer_link text,
    poster       text
);

alter table movie
    owner to postgres;

create table movie_genre
(
    id       serial
        primary key,
    movie_id integer
        references movie
            on update cascade on delete restrict,
    genre_id integer
        references genre
            on update cascade on delete restrict
);

alter table movie_genre
    owner to postgres;

create table movie_country
(
    id         serial
        primary key,
    movie_id   integer
        references movie
            on update cascade on delete restrict,
    country_id integer
        references country
            on update cascade on delete restrict
);

alter table movie_country
    owner to postgres;

create table person_post
(
    id        serial
        primary key,
    person_id integer
        references person
            on update cascade on delete restrict,
    post_id   integer
        references post
            on update cascade on delete restrict
);

alter table person_post
    owner to postgres;

create table person_post_movie
(
    id             serial
        primary key,
    movie_id       integer
        references movie
            on update cascade on delete restrict,
    person_post_id integer
        references person_post
            on update cascade on delete restrict
);

alter table person_post_movie
    owner to postgres;

create table usr
(
    id                      serial
        primary key,
    username                varchar(100),
    password                varchar(100),
    account_non_expired     boolean default true,
    account_non_locked      boolean default true,
    credentials_non_expired boolean default true,
    enabled                 boolean default true,
    avatar                  varchar(256)
);

alter table usr
    owner to postgres;

create table authority
(
    id   serial
        primary key,
    name varchar(100)
        unique
);

alter table authority
    owner to postgres;

create table usr_authorities
(
    id           serial
        primary key,
    usr_id       integer
        references usr
            on update cascade on delete restrict,
    authority_id integer
        references authority
            on update cascade on delete restrict
);

alter table usr_authorities
    owner to postgres;

create table usr_rating_movie
(
    id       serial
        primary key,
    usr_id   integer
        references usr
            on update cascade on delete restrict,
    movie_id integer
        references movie
            on update cascade on delete restrict,
    rating   integer
);

INSERT INTO public.authority (id, name) VALUES (1, 'ROLE_USER');

INSERT INTO public.movie (id, name, budget, box_office, release, duration, synopsis, rating, rating_num, trailer_link, poster) VALUES (5, 'The Lord of the Rings: The Return of the Kings', 94000000, 377027325, '2003-12-01', '03:21:00', 'Aragorn is revealed as the heir to the ancient kings as he, Gandalf and the other members of the broken fellowship struggle to save Gondor.', 6.733333, 3, 'https://www.youtube.com/embed/r5X-hFf6Bwo', 'D:\movie-library-images\movie-posters\Lord-of-rings3-poster.jpg');
INSERT INTO public.movie (id, name, budget, box_office, release, duration, synopsis, rating, rating_num, trailer_link, poster) VALUES (4, 'The Godfather', 6000000, 243862778, '1972-03-14', '02:55:00', 'In late summer 1945, guests are gathered for the wedding reception of Don Vito Corleone''s daughter Connie and Carlo Rizzi. Vito, the head of the Corleone Mafia family, is known to friends and associates as "Godfather." He and Tom Hagen, the Corleone family lawyer and Vito''s adopted son, are hearing requests for favors because, according to Italian tradition, "no Sicilian can refuse a request on his daughter''s wedding day." One of the men who asks the Don for a favor is Amerigo Bonasera, a mortician and acquaintance of the Don, whose daughter was brutally beaten by two young men for refusing their advances; the men received minimal punishment. The Don is disappointed in Bonasera, who''d avoided most contact with the Don due to Corleone''s criminal dealings. The Don''s wife is godmother to Bonasera''s shamed daughter, a relationship the Don uses to extract new loyalty from the undertaker. The Don agrees to have his men punish the young men responsible in return for future service if necessary.', 8.885715, 7, 'https://www.youtube.com/embed/sY1S34973zA', 'D:\movie-library-images\movie-posters\Godfather-poster.jpg');
INSERT INTO public.movie (id, name, budget, box_office, release, duration, synopsis, rating, rating_num, trailer_link, poster) VALUES (3, 'Pulp Fiction', 8000000, 213900000, '1994-05-21', '02:34:00', 'Vincent Vega and Jules Winnfield are hitmen with a penchant for philosophical discussions. In this ultra-hip, multi-strand crime movie, their storyline is interwoven with those of their boss, gangster Marsellus Wallace; hist actress wife, Mia; struggling boxer Butch Coolidge; master fixer Winston Wolfe and a nervous pair of armed robbers, "Pumpkin" and "Honey Bunny".', 8.4, 5, 'https://www.youtube.com/embed/s7EdQ4FqbhY', 'D:\movie-library-images\movie-posters\Pulp-Fiction-poster.jpg');
INSERT INTO public.movie (id, name, budget, box_office, release, duration, synopsis, rating, rating_num, trailer_link, poster) VALUES (1, 'Tenet', 205000000, 300000000, '2020-09-03', '02:30:00', 'Armed with only one word—Tenet—and fighting for the survival of the entire world, the Protagonist journeys through a twilight world of international espionage on a mission that will unfold in something beyond real time.

Not time travel. Inversion.', 8.6, 5, 'https://www.youtube.com/embed/L3pk_TBkihU', 'D:\movie-library-images\movie-posters\Tenet-poster.jpg');
INSERT INTO public.movie (id, name, budget, box_office, release, duration, synopsis, rating, rating_num, trailer_link, poster) VALUES (2, 'Lock, Stock and Two Smoking Barrels', 1350000, 25000000, '1998-08-23', '01:47:00', 'Hoping to make a bundle in a high-stakes poker game, four shiftless lad from London''s East End instead find themselves swimming in debt.', 9.083333, 6, 'https://www.youtube.com/embed/h6hZkvrFIj0', 'D:\movie-library-images\movie-posters\Lock-Stock-2-Smoking-Barrels-poster.jpg');

INSERT INTO public.usr (id, username, password, account_non_expired, account_non_locked, credentials_non_expired, enabled, avatar) VALUES (1, 'testuser', '$2a$10$8n5lLmZL4R6q.VxL5Hm6muLStqTKMgR9cCureiipxzbKIVsakLPIK', true, true, true, true, 'D:\movie-library-images\users-avatar\testuser-avatar.jpeg');

INSERT INTO public.post (id, name) VALUES (1, 'director');
INSERT INTO public.post (id, name) VALUES (2, 'writer');
INSERT INTO public.post (id, name) VALUES (3, 'producer');
INSERT INTO public.post (id, name) VALUES (4, 'operator');
INSERT INTO public.post (id, name) VALUES (5, 'compositor');
INSERT INTO public.post (id, name) VALUES (6, 'actor');

INSERT INTO public.genre (id, name) VALUES (1, 'Fantasy');
INSERT INTO public.genre (id, name) VALUES (2, 'Action');
INSERT INTO public.genre (id, name) VALUES (3, 'Comedy');
INSERT INTO public.genre (id, name) VALUES (4, 'Criminal');
INSERT INTO public.genre (id, name) VALUES (5, 'Thriller');
INSERT INTO public.genre (id, name) VALUES (6, 'Drama');
INSERT INTO public.genre (id, name) VALUES (7, 'Adventure');

INSERT INTO public.person (id, name) VALUES (1, 'Christopher Nolan');
INSERT INTO public.person (id, name) VALUES (2, 'Ivo Felt');
INSERT INTO public.person (id, name) VALUES (3, 'Thomas Hayslip');
INSERT INTO public.person (id, name) VALUES (4, 'Emma Thomas');
INSERT INTO public.person (id, name) VALUES (5, 'Hoyte Van Hoytema');
INSERT INTO public.person (id, name) VALUES (6, 'Ludwig Goransson');
INSERT INTO public.person (id, name) VALUES (7, 'John David Washington');
INSERT INTO public.person (id, name) VALUES (8, 'Robert Pattinson');
INSERT INTO public.person (id, name) VALUES (9, 'Elizabeth Debicki');
INSERT INTO public.person (id, name) VALUES (10, 'Guy Ritchie');
INSERT INTO public.person (id, name) VALUES (11, 'Matthew Vaughn');
INSERT INTO public.person (id, name) VALUES (12, 'David A. Hughes');
INSERT INTO public.person (id, name) VALUES (13, 'John Murphy');
INSERT INTO public.person (id, name) VALUES (14, 'Jason Flemyng');
INSERT INTO public.person (id, name) VALUES (15, 'Dexter Fletcher');
INSERT INTO public.person (id, name) VALUES (16, 'Nick Moran');
INSERT INTO public.person (id, name) VALUES (17, 'Jason Statham');
INSERT INTO public.person (id, name) VALUES (18, 'Tim Maurice-Jones');
INSERT INTO public.person (id, name) VALUES (19, 'Quentin Tarantino');
INSERT INTO public.person (id, name) VALUES (20, 'Roger Avary');
INSERT INTO public.person (id, name) VALUES (21, 'Lawrence Bender');
INSERT INTO public.person (id, name) VALUES (22, 'Andzej Sekuta');
INSERT INTO public.person (id, name) VALUES (23, 'John Travolta');
INSERT INTO public.person (id, name) VALUES (24, 'Samuel L. Jackson');
INSERT INTO public.person (id, name) VALUES (25, 'Uma Thurman');
INSERT INTO public.person (id, name) VALUES (26, 'Harvey Keitel');
INSERT INTO public.person (id, name) VALUES (27, 'Tim Roth');
INSERT INTO public.person (id, name) VALUES (28, 'Francis Ford Coppola');
INSERT INTO public.person (id, name) VALUES (29, 'Albert S. Ruddy');
INSERT INTO public.person (id, name) VALUES (30, 'Mario Puzo');
INSERT INTO public.person (id, name) VALUES (31, 'Gordon Willis');
INSERT INTO public.person (id, name) VALUES (32, 'Nino Rota');
INSERT INTO public.person (id, name) VALUES (33, 'Marlon Brando');
INSERT INTO public.person (id, name) VALUES (34, 'Al Pacino');
INSERT INTO public.person (id, name) VALUES (35, 'James Caan');
INSERT INTO public.person (id, name) VALUES (36, 'Richard Castellano');
INSERT INTO public.person (id, name) VALUES (37, 'Robert Duvall');
INSERT INTO public.person (id, name) VALUES (38, 'Peter Jackson');
INSERT INTO public.person (id, name) VALUES (39, 'Barrie M. Osborne');
INSERT INTO public.person (id, name) VALUES (40, 'Fran Waish');
INSERT INTO public.person (id, name) VALUES (41, 'Philippa Boyens');
INSERT INTO public.person (id, name) VALUES (42, 'Andrew Lesnie');
INSERT INTO public.person (id, name) VALUES (43, 'Howard Shore');
INSERT INTO public.person (id, name) VALUES (44, 'Elijah Wood');
INSERT INTO public.person (id, name) VALUES (45, 'Ian McKellen');
INSERT INTO public.person (id, name) VALUES (46, 'Liv Tyler');
INSERT INTO public.person (id, name) VALUES (47, 'Viggo Mortensen');
INSERT INTO public.person (id, name) VALUES (48, 'Sean Astin');

INSERT INTO public.country (id, name) VALUES (1, 'Great Britain');
INSERT INTO public.country (id, name) VALUES (2, 'USA');
INSERT INTO public.country (id, name) VALUES (3, 'New Zealand');

INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (1, 1, 1);
INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (2, 1, 2);
INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (3, 2, 1);
INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (4, 3, 2);
INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (5, 4, 2);
INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (6, 4, 2);
INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (7, 5, 2);
INSERT INTO public.movie_country (id, movie_id, country_id) VALUES (8, 5, 3);

INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (1, 1, 1);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (2, 1, 2);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (3, 2, 3);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (4, 2, 4);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (5, 3, 5);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (6, 3, 3);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (7, 3, 4);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (8, 4, 6);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (9, 4, 4);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (10, 4, 6);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (11, 4, 4);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (12, 5, 1);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (13, 5, 7);
INSERT INTO public.movie_genre (id, movie_id, genre_id) VALUES (14, 5, 6);

INSERT INTO public.person_post (id, person_id, post_id) VALUES (1, 1, 1);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (2, 1, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (3, 2, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (4, 3, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (5, 1, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (6, 4, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (7, 5, 4);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (8, 6, 5);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (9, 7, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (10, 8, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (11, 9, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (12, 10, 1);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (13, 10, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (14, 11, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (15, 18, 4);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (16, 12, 5);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (17, 13, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (18, 14, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (19, 15, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (20, 17, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (21, 19, 1);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (22, 19, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (23, 20, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (24, 21, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (25, 22, 4);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (26, 23, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (27, 24, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (28, 25, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (29, 26, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (30, 27, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (31, 28, 1);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (32, 30, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (33, 28, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (34, 29, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (35, 31, 4);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (36, 32, 5);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (37, 33, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (38, 34, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (39, 35, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (40, 36, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (41, 37, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (42, 28, 1);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (43, 30, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (44, 28, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (45, 29, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (46, 31, 4);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (47, 32, 5);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (48, 33, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (49, 34, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (50, 35, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (51, 36, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (52, 37, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (53, 38, 1);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (54, 40, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (55, 41, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (56, 38, 2);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (57, 38, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (58, 39, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (59, 40, 3);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (60, 42, 4);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (61, 43, 5);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (62, 44, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (63, 45, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (64, 46, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (65, 47, 6);
INSERT INTO public.person_post (id, person_id, post_id) VALUES (66, 48, 6);

INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (1, 1, 1);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (2, 1, 2);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (3, 1, 3);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (4, 1, 4);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (5, 1, 5);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (6, 1, 6);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (7, 1, 7);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (8, 1, 8);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (9, 1, 9);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (10, 1, 10);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (11, 1, 11);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (12, 2, 12);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (13, 2, 13);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (14, 2, 14);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (15, 2, 15);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (16, 2, 16);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (17, 2, 17);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (18, 2, 18);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (19, 2, 19);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (20, 2, 20);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (21, 3, 21);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (22, 3, 22);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (23, 3, 23);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (24, 3, 24);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (25, 3, 25);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (26, 3, 26);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (27, 3, 27);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (28, 3, 28);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (29, 3, 29);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (30, 3, 30);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (31, 4, 31);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (32, 4, 32);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (33, 4, 33);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (34, 4, 34);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (35, 4, 35);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (36, 4, 36);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (37, 4, 37);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (38, 4, 38);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (39, 4, 39);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (40, 4, 40);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (41, 4, 41);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (42, 5, 42);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (43, 5, 43);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (44, 5, 44);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (45, 5, 45);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (46, 5, 46);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (47, 5, 47);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (48, 5, 48);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (49, 5, 49);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (50, 5, 50);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (51, 5, 51);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (52, 5, 52);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (53, 5, 53);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (54, 5, 54);
INSERT INTO public.person_post_movie (id, movie_id, person_post_id) VALUES (55, 5, 55);

INSERT INTO public.usr_authorities (id, usr_id, authority_id) VALUES (1, 1, 1);

INSERT INTO public.usr_rating_movie (id, usr_id, movie_id, rating) VALUES (1, 1, 5, 5);
INSERT INTO public.usr_rating_movie (id, usr_id, movie_id, rating) VALUES (3, 1, 4, 7);
INSERT INTO public.usr_rating_movie (id, usr_id, movie_id, rating) VALUES (4, 1, 3, 6);
INSERT INTO public.usr_rating_movie (id, usr_id, movie_id, rating) VALUES (2, 1, 2, 7);