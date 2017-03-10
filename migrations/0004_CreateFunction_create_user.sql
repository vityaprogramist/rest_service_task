-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION create_user(name VARCHAR, login VARCHAR, password_hash VARCHAR, age INTEGER, phone BIGINT) RETURNS VOID AS $$ DECLARE cn text; BEGIN INSERT INTO users.users(name, login, password_hash, age, phone) VALUES(name, login, password_hash, age, phone); EXCEPTION WHEN UNIQUE_VIOLATION THEN RAISE 'Пользователь с имене "%" уже существует.', login; WHEN CHECK_VIOLATION THEN GET STACKED DIAGNOSTICS cn = CONSTRAINT_NAME; IF cn LIKE 'users_age_check' THEN RAISE NOTICE 'Возраст не может быть меньше 0.'; ELSEIF cn LIKE 'users_phone_check' THEN RAISE NOTICE 'Введите корректный лелефон.'; END IF; WHEN OTHERS THEN RAISE 'unknown error'; END;$$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS create_user(VARCHAR, VARCHAR, VARCHAR, INTEGER, BIGINT);
