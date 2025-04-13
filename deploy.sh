set -e
GOOS=linux GOARCH=arm64 go build -o ./deploy/app
rsync -avz --delete -e ssh ./deploy/ ./ui deploy.target:$HOME/healthy
ssh deploy.target "$HOME/healthy/remote-deploy.sh"

