-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

DELETE FROM rental.rental;
DELETE FROM film.film_genre;
DELETE FROM film.film;
DELETE FROM film.genre;
DELETE FROM users.user;

ALTER SEQUENCE film.film_id_seq RESTART WITH 1;
ALTER SEQUENCE film.genre_id_seq RESTART WITH 1;
ALTER SEQUENCE users.user_id_seq RESTART WITH 1;

INSERT INTO film.film(title, release_year) VALUES
  ('Побег из шоушенка',1994),
  ('Крёстный отец',1972),
  ('Крёстный отец 2',1974),
  ('Тёмный рыцарь',2008),
  ('12 рaзгневанных мужчин',1957),
  ('Список Шиндлера',1993),
  ('Криминальное чтиво',1994),
  ('Властелин колец: Возвращение короля',2003),
  ('Хороший, плохой, злой',1966),
  ('Бойцовский клуб',1999),
  ('Властелин колец: Братство кольца',2001),
  ('Звёздные войны. Эпизод 5: Империя наносит ответный удар',1980),
  ('Форрест Гамп',1994),
  ('Начало',2010),
  ('Властелин колец: Две крепости',2002),
  ('Пролетая над гнездом кукушки',1975),
  ('Славные парни',1990),
  ('Матрица',1999);

INSERT INTO film.genre(name) VALUES
  ('action'),
  ('drama'),
  ('comedy'),
  ('fantasy'),
  ('horror'),
  ('criminal'),
  ('adventure'),
  ('thriller'),
  ('historical'),
  ('western');

INSERT INTO film.film_genre(film_id, genre_id) VALUES
  (1,2),
  (1,6),
  (2,1),
  (2,2),
  (2,6),
  (3,1),
  (3,2),
  (3,6),
  (4,1),
  (5,2),
  (6,2),
  (6,9),
  (7,1),
  (7,6),
  (8,1),
  (8,4),
  (8,2),
  (8,7),
  (9,1),
  (9,10),
  (10,1),
  (10,8),
  (10,2),
  (11,1),
  (11,2),
  (11,4),
  (11,7),
  (12,1),
  (12,2),
  (12,4),
  (13,2),
  (14,1),
  (14,4),
  (14,7),
  (14,8),
  (15,1),
  (15,4),
  (15,2),
  (15,7),
  (16,2),
  (17,1),
  (17,2),
  (17,6),
  (17,7),
  (18,1),
  (18,4),
  (18,7);

INSERT INTO users.user(first_name, last_name, login, password_hash, age, phone) VALUES
  ('Василий','Петров','petrov_vasya','f6b54c979710647f4338489c07192d4ce9b40ba3',19,79991236677),
  ('Татьяна','Краснова','tanya_red','ece9b8a4b6fbfaf292ec450b14f14bb4e8ba1f8c',29,787639933803),
  ('Елена','Авдеева','elen','16fd5d06c7e8c692953b238fe1cae87dfb97af95',45,794347602709),
  ('Алена','Беляева','belka','bd68c638f745cccb5d386abb6b55dfd920f3d036',34,79958495714),
  ('Станислав','Зыков','stas_bigstar','3e4f1c1d275d04d212d4886535940e63e1f7e260',20,758553567994),
  ('Ашот','Иванов','ashotishe','c60c65072c40bf2386020834f92580a5ce70a032',21,79736886404),
  ('Иван','Захаров','xxx_ivan_xxx','45e00df6472750dd4047191a10b7d1568c6cb598',33,78857197520),
  ('Петр','Ильин','petrovich','8e58131057215bd00597da9b614d97f47a4b8a16',19,79938666829),
  ('Виктор','Субботин','subbota','3033afdf4a00f9fd49f01df5a16704ba80f117a6',23,797785433432);

INSERT INTO rental.rental(film_id, user_id) VALUES
  (1,1),
  (2,2),
  (3,3),
  (4,4),
  (5,5),
  (6,6),
  (7,7),
  (8,8),
  (9,9),
  (10,1),
  (11,2),
  (12,3),
  (13,4),
  (14,5),
  (15,6),
  (16,7),
  (17,8),
  (18,9),
  (1,2),
  (2,3),
  (3,4),
  (4,5),
  (5,6),
  (6,7),
  (7,8),
  (8,9),
  (9,1);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM rental.rental;
DELETE FROM film.film_genre;
DELETE FROM users.user;
DELETE FROM film.genre;
DELETE FROM film.film;
