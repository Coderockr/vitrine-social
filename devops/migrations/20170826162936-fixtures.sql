
-- +migrate Up
insert into organizations values (1,'Unesco', 'http://placehold.it/200x200', 'Rua das Nações Unidas', '55 1223232', 'Descrição da Unesco', 'https://www.youtube.com/watch?v=PJC7zaZT-Dg', 'admin@unesco.org', '$2a$10$uKmMPOij7PhRMbWAQvLkUuWx2JyJtJzyCLXxxIno/RWocP9qP7Pji', 'unesco');
insert into needs values (1, 4, 1, 'Designer', 'Designer para novo site', 100, 0, null, 'ACTIVE', 'horas');
insert into needs values (2, 4, 1, 'Programador', 'Programador para novo site', 100, 0, null, 'ACTIVE', 'horas');
-- +migrate Down
