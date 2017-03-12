-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION public.auth_user(IN user_login VARCHAR, IN password_hash VARCHAR)   RETURNS TABLE(id BIGINT, login VARCHAR, first_name VARCHAR, last_name VARCHAR, age INT, phone BIGINT) AS $$ DECLARE   pass_hash VARCHAR; BEGIN   IF EXISTS(SELECT users.user.password_hash FROM users.user WHERE users.user.login = user_login) THEN     pass_hash := (SELECT users.user.password_hash FROM users.user WHERE users.user.login = user_login);     IF pass_hash = password_hash THEN       RETURN QUERY       SELECT         users.user.id AS id,         users.user.login AS login,         users.user.first_name AS first_name,         users.user.last_name AS last_name,         users.user.age AS age,         users.user.phone AS phone       FROM         users.user       WHERE         users.user.login LIKE user_login;     ELSE       RAISE 'Неверный логин или пароль.';     END IF;   ELSE     RAISE 'Неверный логин или пароль.';   END IF; END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS public.auth_user(VARCHAR, VARCHAR);
