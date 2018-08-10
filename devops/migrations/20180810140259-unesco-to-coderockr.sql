
-- +migrate Up
update organizations set email = 'admin@coderockr.com', video = 'https://www.youtube.com/watch?v=fNbcPla2hVY', name = 'Coderockr', phone = '47 3227-6359', about = 'Somos a Coderockr', password = '$2a$10$MtTaCVIsnjCDolZ0R/veXO0TvQ3TrvWEWLaHU2w7fuAq2MmnIezJ2', slug = 'coderockr', street = 'Rua Henrique Meyer', number = '40', complement = 'Sala 1',neighborhood = 'Centro', city = 'Joinville' , state = 'SC', zipcode = '89201-405', facebook = 'https://www.facebook.com/Coderockr/', website = 'http:/codeorckr.com', instagram = null, whatsapp = null where email = 'admin@unesco.org';
-- +migrate Down
select 1;