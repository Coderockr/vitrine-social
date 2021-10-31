Vitrine Social [Codacy Badge](https://www.codacy.com/app/lucassabreu/vitrine-social?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Coderockr/vitrine-social&amp;utm_campaign=Badge_Grade) [![Build Status](https://travis-ci.org/Coderockr/vitrine-social.svg?branch=master)](https://travis-ci.org/Coderockr/vitrine-social) [![codecov](https://codecov.io/gh/Coderockr/vitrine-social/branch/master/graph/badge.svg)](https://codecov.io/gh/Coderockr/vitrine-social)
===============
[Waffle.io - Columns and their card count](https://waffle.io/Coderockr/vitrine-social)

## Issues e Progresso

O controle das tarefas e do progresso das mesmas estão sendo feitas no Waffle. Clique aqui para acompanhar: https://waffle.io/Coderockr/vitrine-social

## Instalação Backend (Go)

Estamos utilizando [Go Modules](https://github.com/golang/go/wiki/Modules) nesse projeto, por isso a pasta do projeto precisa ficar fora do seu `GOPATH`, ou terá que adicionar a ENV `GO111MODULE` como `on` em seu ambiente para que o projeto funcione dentro do `GOPATH`.

Recomendamos manter o projeto fora do seu `GOPATH`, assim o `go` não vai gerar um módulo sem necessidade na raiz do projeto, ou afetar outros projetos `go` em seu ambiente que ainda não estejam utilizando `Go Modules`.

Resumo da ópera, para começar a trabalhar basta rodar os seguintes comandos:

```sh
git clone git@github.com:Coderockr/vitrine-social.git /not/your/go/path/vitrine-social;

make setup # executar na primeira vez para instalar todas as dependencias e ferramentas

make migrations # isso pode falhar por causa do warmup do postgres

make serve # agora esta rodando :)
```

#### Instalando ambiente de desenvolvimento backend no Docker

Na primeira vez que for utilizar o projeto execute o comando:
```sh
make setup-on-docker
```

Após o comando concluir e nas próximas vezes precisa apenas:
```sh
make migrations-on-docker
make serve-on-docker
```

O terminal estará bloqueado durante a execução do backend

### Domínios e Subdomínios locais

Incluir os seguintes domínios no seu `/etc/hosts` deve agilizar o setup do seu projeto:

```sh
127.0.0.1 api.vitrinesocial.test # usar porta 8000 (golang)
127.0.0.1 images.vitrinesocial.test # usar porta 7000 (images-server)
127.0.0.1 minio.vitrinesocial.test # usar porta 9000 (minio)
127.0.0.1 vitrinesocial.test # usar porta 3000 (frontend)
```

## Instalação Frontend (React)

```sh
cd frontend

yarn

yarn start
```


### Reicons

Mover ícones para assets/icons

```sh
yarn reicons
```

## Comandos Auxiliares (dia-a-dia)

Estamos mantendo todos os comandos auxiliares (criar migration, rodar migrations, regerar docs, etc) dentro do `Makefile` na raiz do projeto.

Para ver quais são os comandos disponívels execute: `make help` e todos serão listados.

## Documentação API

Para acessar a versão mais recente da definição acesse:

http://coderockr.com/vitrine-social/

> [Como atualizar a documentação?](./CONTRIBUTING.md#atualize-a-documentação)

## [Contribuindo](./CONTRIBUTING.md)

Leia o nosso [CONTRIBUTING.md](./CONTRIBUTING.md) para aprender sobre o nosso processo de desenvolvimento, como propor bugfixes e melhorias, e como encontrar issues para atuar.
