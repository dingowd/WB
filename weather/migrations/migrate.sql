CREATE TABLE IF NOT EXISTS cities
(
    city_id serial primary key,
    name varchar(50) not null,
    lat real  not null,
    lon real  not null,
    country varchar(2)  not null,
    state varchar(255) not null,
    UNIQUE (name, country, state)
);

create table if not exists weather
(
    id serial primary key,
    city_id int references cities (city_id),
    date varchar(10) not null,
    temp real not null,
    detail JSONB not null
);

create table if not exists favor
(
    id serial primary key,
    user_name text unique not null,
    favor text[] not null default '{}'::text[]
);

insert into cities(name, lat, lon, country, state) values
    ('Москва', 55.7504461, 37.6174943, 'RU', 'Moscow'),
    ('Марсель', 43.2961743, 5.3699525, 'FR', 'Provence-Alpes-Côte dAzur'),
    ('городской округ Казань', 55.7823547, 49.1242266, 'RU', 'Tatarstan'),
    ('Самара', 53.198627, 50.113987, 'RU', 'Samara Oblast'),
    ('Салехард', 66.5375387, 66.6157469, 'RU', 'Yamalo-Nenets Autonomous Okrug'),
    ('Алматы', 43.2363924, 76.9457275, 'KZ', ''),
    ('Нижний Новгород', 56.3264816, 44.0051395, 'RU', 'Nizhny Novgorod Oblast'),
    ('Владивосток', 43.1150678, 131.8855768, 'RU', 'Primorsky Krai'),
    ('Тюмень', 57.153534, 65.542274, 'RU', 'Tyumen Oblast'),
    ('Воронеж', 51.6605982, 39.2005858, 'RU', 'Voronezh Oblast'),
    ('Лондон', 51.5073219, -0.1276474, 'GB', 'England'),
    ('Томск', 56.488712, 84.952324, 'RU', 'Tomsk Oblast'),
    ('Нью-Йорк', 40.7127281, -74.0060152, 'US', 'New York'),
    ('Чебоксары', 56.1307195, 47.2449597, 'RU', 'Chuvashia'),
    ('Париж', 48.8588897, 2.3200410217200766, 'FR', 'Ile-de-France'),
    ('Белгород', 50.5955595, 36.5873394, 'RU', 'Belgorod Oblast'),
    ('городской округ Йошкар-Ола', 56.6328248, 47.8972462, 'RU', 'Mari El Republic'),
    ('Санкт-Петербург', 59.938732, 30.316229, 'RU', 'Saint Petersburg'),
    ('Уфа', 54.7261409, 55.947499, 'RU', 'Bashkortostan'),
    ('Челябинск', 55.1598408, 61.4025547, 'RU', 'Chelyabinsk Oblast');