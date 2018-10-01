Contribuindo para a Vitrine Social
==================================

Obrigado por se interessar em contribuir com o Vitrine Social, você pode contribuir em uma das seguintes formas:

- Notificando um Bug
- Discutindo sobre melhorias para o fonte
- Enviando correções
- Sugerindo melhorias e novas features

## Nós trabalhamos usando o GitHub

Nós usamos o GitHub para como nosso ponto focal, mantemos o fonte, controlamos as issues e sugestões de melhorias, assim como aceitamos pull requests.

## Usamos [Github Flow](https://guides.github.com/introduction/flow/index.html), Então todas as Mudanças Acontecem via Pull Requests

Pull requests são a melhor forma de sugerir mudanças no fonte (usamos o [Github Flow](https://guides.github.com/introduction/flow/index.html)). E seus pull requests são bem vindos.

1. Faça um fork do repositório e crie sua branch a partir da `master`.
2. Se adicionou código que precise de testes, crie novos testes.
3. Se mudou alguma API, [atualize a documentação](#atualize-a-documentação).
4. Garanta que os testes estão passando.
5. Tenha certeza que o seu código passa no [lint](#padrões-de-código).
6. Crie o pull request !

## Reporte Bugs pelo [Github:Issues](https://github.com/Coderockr/vitrine-social/issues)

Nós usamos o GitHub:Issues para controlar nossos bugs públicos. Registre um bug [abrindo uma nova issue](https://github.com/Coderockr/vitrine-social/issues/new?labels=Type%3A%20Bug,Stage%3A%20Backlog); é fácil.

## Padrões de Código

### Go (Backend)

Usamos as regras padrões da linguagem Go, para garantir pode executar o comando `make lint`

### JavaScript (Frontend)

Para os fontes em JavaScript estamos usamos ECMA6 com uma extensão do lint do Airbnb, para mais detalhes [clique aqui](https://github.com/Coderockr/vitrine-social/blob/master/frontend/.eslintrc).

Para garantir pode rodar `npm run lint` dentro da pasta `frontend`

## Atualize a Documentação

Sempre que forem feitas alterações nas APIs, a documentação em [`api.apib`](https://github.com/Coderockr/vitrine-social/blob/master/docs/api.apib) deve ser atualizada.

Após terminar de documentar execute o comando `make docs-build` para atualizar a documentação e faça um commit com o `api.apib` e `index.html` com eles.

## Referencias

Este documento foi adaptado do seguinte exemplo: [briandk/CONTRIBUTING.md](https://gist.github.com/briandk/3d2e8b3ec8daf5a27a62)
