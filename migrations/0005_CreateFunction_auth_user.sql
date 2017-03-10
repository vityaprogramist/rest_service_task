-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION auth_user(IN user_login VARCHAR, IN password_hash VARCHAR)   RETURNS TABLE(id BIGINT, login VARCHAR, name VARCHAR, age INT, phone BIGINT) AS $$ DECLARE   pass_hash VARCHAR; BEGIN   IF EXISTS(SELECT users.users.password_hash FROM users.users WHERE users.users.login = user_login) THEN     pass_hash := (SELECT users.users.password_hash FROM users.users WHERE users.users.login = user_login);     IF pass_hash = password_hash THEN       RETURN QUERY         SELECT           users.users.id AS id,           users.users.login AS login,           users.users.name AS name,           users.users.age AS age,           users.users.phone AS phone         FROM           users.users         WHERE           users.users.login LIKE user_login;     ELSE       RAISE 'Неверный логин или пароль.';     END IF;   ELSE     RAISE 'Неверный логин или пароль.';   END IF; END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS auth_user(VARCHAR, VARCHAR);
