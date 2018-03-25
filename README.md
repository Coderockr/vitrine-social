Vitrine Social [![Codacy Badge](https://api.codacy.com/project/badge/Grade/941498dc58244d23aa2cf09148509512)](https://www.codacy.com/app/lucassabreu/vitrine-social?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Coderockr/vitrine-social&amp;utm_campaign=Badge_Grade) [![Build Status](https://travis-ci.org/Coderockr/vitrine-social.svg?branch=master)](https://travis-ci.org/Coderockr/vitrine-social) [![codecov](https://codecov.io/gh/Coderockr/vitrine-social/branch/master/graph/badge.svg)](https://codecov.io/gh/Coderockr/vitrine-social)
===============

## Instalação Backend

```sh
git clone git@github.com:Coderockr/vitrine-social.git $GOPATH/src/Coderockr/vitrine-social;

cd $GOPATH/src/Coderockr/vitrine-social;

make install;

make serve;
```

### Migrations

#### Criar uma migration

```sh
sql-migrate new -config=./devops/dbconfig.yml -env=production default-categories
```

## Instalação Frontend

```sh
cd Frontend

yarn

yarn start
```

### Reicons

Mover ícones para assets/icons

```sh
yarn reicons
```

## Documentação API

Para acessar a versão mais recente da definição acesse: 

http://coderockr.com/vitrine-social/
