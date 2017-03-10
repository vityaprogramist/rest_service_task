-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO films.films(name, genre, release)
VALUES
  ('scream', '{"horror","comedy"}', '1990/02/20'),
  ('kuku', '{"comedy","drama"}', '1990/02/20'),
  ('kindzadza', '{"drama","fantasy"}', '1990/02/20'),
  ('lord of ring', '{"adventure","comedy","drama","action"}', '1990/02/20'),
  ('torque', '{"action"}', '1990/02/20'),
  ('avatar', '{"drama","fantasy"}', '1990/02/20'),
  ('comedy woman', '{"comedy"}', '1990/02/20'),
  ('transporter', '{"action"}', '1990/02/20'),
  ('mechanic', '{"action"}', '1990/02/20'),
  ('unstoppable', '{"action", "comedy"}', '1990/02/20'),
  ('unstoppable 2', '{"action", "comedy"}', '1990/02/20'),
  ('unstoppable 3', '{"action", "comedy"}', '1990/02/20');


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM films.films;

