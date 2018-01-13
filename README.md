# vitrine-social

# Instalação

```
    go get github.com/Coderockr/vitrine-social;

    cd $GOPATH/Coderockr/vitrine-social;

    make install;

    make serve;
```

# Migrations

## Criar uma migration

    sql-migrate new -config=./devops/dbconfig.yml -env=production default-categories
