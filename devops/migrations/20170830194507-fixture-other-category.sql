
-- +migrate Up
insert into needs (category_id, organization_id, title, description, required_qtd, reached_qtd, due_date, status, unity) 
    values (5, 1, 'Voluntário', 'Ajudar a levar comida as áreas de risco', 100, 0, '2018-01-01', 'ACTIVE', 'semanas');
-- +migrate Down
