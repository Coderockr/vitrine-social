# Contribuindo para a Vitrine Social

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

## Fluxo das Issues / Como Achar Issues para Atuar

Hoje estamos usando as labels e processo do [Coderockr Way](https://github.com/Coderockr/coderockr-way-github-setup), elas são diferentes das labels padrões do GitHub, então para evitar confusão na hora de contribuir com uma tarefa, considere as seguintes labels:

 * `Stage: Backlog`: são issues que acabaram de ser criadas, ou que ainda não foram assumidas por ninguêm, caso esteja querendo ajudar implementando uma issue, [basta pegar uma com essa label](https://github.com/Coderockr/vitrine-social/labels/Stage%3A%20Backlog)
 * `Stage: Analysis` e `Stage: In progress`: são issues que já estão em execução, alguêm esta [analisando](https://github.com/Coderockr/vitrine-social/labels/Stage%3A%20Analysis) ou [codificando](https://github.com/Coderockr/vitrine-social/labels/Stage%3A%20In%20progress) ela.
 * `Stage: Review`: quando é aberto um pull request e o mesmo está pronto para ser avaliado pelos outros membros, a issue estará com esa label. Caso queria ajudar na revisão dos códigos [aqui é o lugar](https://github.com/Coderockr/vitrine-social/labels/Stage%3A%20Review).
 * `Stage: Testing`: depois do merge as issues são [movidas para essa etapa](https://github.com/Coderockr/vitrine-social/labels/Stage%3A%20Testing), um dos membros irá testar a alteração para identificar se ainda falta algo, ou se pode ser concluída.
 
Esse fluxo fica mais claro se olhar em nosso [Waffle](https://waffle.io/Coderockr/vitrine-social).

Também classificamos as issues por `Category`, sendo que uma issue pode estar em mais de uma categoria:
 
 * `Category: Frontend`: para issues que precisão ser feitas no frontend, normalmente exigindo alterações nos arquivos da pasta [`frontend`](https://github.com/Coderockr/vitrine-social/tree/master/frontend) e exigindo conhecimento de JavaScript e React.
 * `Category: Backend`: quando a issue precisa que algo no backend seja feito, alterando os arquivos da pasta [`server`](https://github.com/Coderockr/vitrine-social/tree/master/server), aqui é necessário ter conhecimento de Go e provavelmente de PostgreSQL.
 
Além disso tentamos informar qual a dificuldade da tarefa (`Level: Easy`, `Level: Medium` e `Level: Hard`), o tipo (`Type: Bug`, `Type: Improvement` e `Type: New Feature`) e a prioridade (`Priority: Lowest`, `Priority: Low`, `Priority: Medium`, `Priority: High` e `Priority: Highest`).

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
