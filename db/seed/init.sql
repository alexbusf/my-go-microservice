-- db/init.sql

-- Удаляем таблицу, если она вдруг осталась от прошлых попыток
DROP TABLE IF EXISTS items;

-- Создаем чистую таблицу
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Наполняем данными
INSERT INTO items (name) VALUES ('Laptop');
INSERT INTO items (name) VALUES ('Keyboard');
INSERT INTO items (name) VALUES ('Mouse');