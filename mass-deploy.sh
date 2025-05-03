set -e

../nats/deploy.sh $1
../maint/deploy.sh $1
./deploy.sh $1
../jobby/deploy.sh $1

../terminaltype/deploy.sh $1

../soulshare/mass-deploy.sh $1

../haproxy/deploy.sh $1

