create table bookmark
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR NOT NULL,
    uri        VARCHAR NOT NULL,
    category   VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

insert into bookmark (name, uri, category) values ('Google', 'https://google.com', 'education');
insert into bookmark (name, uri, category) values ('StackOverflow', 'https://stackoverflow.com', 'education');
insert into bookmark (name, uri, category) values ('Facebook', 'https://facebook.com', 'fakenews');
