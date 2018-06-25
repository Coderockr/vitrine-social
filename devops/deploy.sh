#!/bin/sh

# Stop server service
ssh -i ./devops/vitrinesocial.pem -t $DEPLOY_USER@$DEPLOY_HOST 'sudo systemctl stop vitrine-social'

# Upload new compiled file
scp -i ./devops/vitrinesocial.pem server/vitrine-social $DEPLOY_USER@$DEPLOY_HOST:~/vitrine-social/

# Start serve service
ssh -i ./devops/vitrinesocial.pem -t $DEPLOY_USER@$DEPLOY_HOST 'sudo systemctl start vitrine-social'

# Run Migrations
go get github.com/rubenv/sql-migrate/...
sql-migrate up -config=devops/dbconfig.yml -env=production
