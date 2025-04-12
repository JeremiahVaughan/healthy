set -e
GOOS=linux GOARCH=arm64 go build -o ./deploy/app
rsync -avz --delete -e ssh ./deploy/ ./ui pi3:$HOME/healthy
ssh pi3 "$HOME/healthy/remote-deploy.sh"

