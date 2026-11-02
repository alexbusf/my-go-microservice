-- +goose Up
-- Эта часть выполнится при команде "goose up"

-- Удаляем таблицу, если она осталась от прошлых экспериментов
DROP TABLE IF EXISTS items;

-- Создаем чистую таблицу
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Наполняем данными (Seeding)
INSERT INTO items (name) VALUES ('Laptop');
INSERT INTO items (name) VALUES ('Keyboard');
INSERT INTO items (name) VALUES ('Mouse');


-- +goose Down
-- Эта часть выполнится, если мы решим откатить миграцию (goose down)

DROP TABLE IF EXISTS items;