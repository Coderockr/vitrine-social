
-- +migrate Up
update organizations set street = 'rua exemplar', number = '40', neighborhood = 'centro', city = 'joinville', state = 'sc', zipcode = '89230800' where street = '';

-- +migrate Down
