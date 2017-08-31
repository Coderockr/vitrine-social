-- +migrate Up
insert into needs_images (need_id, name, url) values(1, 'campanha alimentos', 'http://pibjc.org/novosite/wp-content/uploads/2016/11/campanha-alimentos.jpg');
insert into needs_images (need_id, name, url) values(1, 'outra campanha alimentos', 'http://pibjc.org/novosite/wp-content/uploads/2016/11/campanha-alimentos.jpg');
insert into organizations_images (organization_id, name, url) values(1, 'banner', 'http://info.geekie.com.br/wp-content/uploads/2015/04/Unesco.jpg');
insert into organizations_images (organization_id, name, url) values(1, 'SDG', 'http://en.unesco.org/sites/default/files/sdgs_poster_new1.png');
-- +migrate Down
