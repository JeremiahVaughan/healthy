sudo cp $HOME/healthy/healthy.service /etc/systemd/system/healthy.service
sudo systemctl enable healthy.service
sudo systemctl start healthy.service
sudo systemctl restart healthy.service
