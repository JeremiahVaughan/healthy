set -e

cd ../nats
./deploy.sh $1
cd -

cd ../maint
./deploy.sh $1
cd -

./deploy.sh $1

cd ../jobby
./deploy.sh $1
cd -

cd ../terminaltype
./deploy.sh $1
cd -

cd ../haproxy
./deploy.sh $1
cd -

cd ../soulshare
./mass-deploy.sh $1


