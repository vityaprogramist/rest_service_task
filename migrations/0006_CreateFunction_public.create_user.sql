-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION public.create_user(   first_name VARCHAR,   last_name VARCHAR,   login VARCHAR,   password_hash VARCHAR,   age INT,   phone BIGINT ) RETURNS VOID AS $$ DECLARE   cn text; BEGIN    INSERT INTO users.user(first_name, last_name, login, password_hash, age, phone)   VALUES ($1, $2, $3, $4, $5, $6);    EXCEPTION   WHEN UNIQUE_VIOLATION THEN     RAISE EXCEPTION 'Пользователь с имене "%" уже существует.', login;   WHEN CHECK_VIOLATION THEN     GET STACKED DIAGNOSTICS cn = CONSTRAINT_NAME;     IF cn = 'age_check' THEN       RAISE EXCEPTION 'Возраст не может быть меньше 0.';     ELSEIF cn = 'phone_check' THEN       RAISE EXCEPTION 'Введите корректный лелефон.';     END IF;   WHEN OTHERS THEN     RAISE 'unknown error';  END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS public.create_user(VARCHAR, VARCHAR, VARCHAR, VARCHAR, INT, BIGINT);
