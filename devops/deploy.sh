#!/bin/sh

# Stop server service
ssh -i ./devops/vitrinesocial.pem -t $DEPLOY_USER@$DEPLOY_HOST 'sudo systemctl stop caddy vitrine-social'

# Upload new Caddy config file and docker-compose
scp -i ./devops/vitrinesocial.pem devops/Caddyfile $DEPLOY_USER@$DEPLOY_HOST:~/

# Upload new compiled file
scp -i ./devops/vitrinesocial.pem server/vitrine-social $DEPLOY_USER@$DEPLOY_HOST:~/vitrine-social/

# Start serve service
ssh -i ./devops/vitrinesocial.pem -t $DEPLOY_USER@$DEPLOY_HOST 'sudo systemctl start caddy vitrine-social'

# Run Migrations
sql-migrate up -config=devops/dbconfig.yml -env=production

# Update sitemap
ssh -i ./devops/vitrinesocial.pem -t $DEPLOY_USER@$DEPLOY_HOST 'cd ./vitrine-social; ./vitrine-social sitemap-generate'
