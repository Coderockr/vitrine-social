# vitrine-social

# Instalação

```
    git clone git@github.com:Coderockr/vitrine-social.git $GOPATH/src/Coderockr/vitrine-social;

    cd $GOPATH/src/Coderockr/vitrine-social;

    make install;

    make serve;
```

# Migrations

## Criar uma migration

    sql-migrate new -config=./devops/dbconfig.yml -env=production default-categories
