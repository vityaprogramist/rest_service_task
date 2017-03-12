-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION public.start_rent(user_login VARCHAR, film_id INT) RETURNS VOID AS $$ DECLARE   u_id BIGINT; BEGIN   SELECT id INTO u_id FROM users.user WHERE login = user_login;   IF u_id IS NOT NULL THEN     IF EXISTS(SELECT * FROM film.film WHERE id = film_id) THEN       INSERT INTO rental.rental(film_id, user_id) VALUES(film_id, u_id);     ELSE       RAISE EXCEPTION 'Не получилось аредовать фильм. Попробуйте еще раз.';     END IF;   ELSE     RAISE EXCEPTION 'Этот пользователь не может арендовать этот фильм.';   END IF;    EXCEPTION     WHEN UNIQUE_VIOLATION THEN       RAISE NOTICE 'Пользователь уже арендовал этот фильм.'; END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS public.start_rent(VARCHAR, INT);
