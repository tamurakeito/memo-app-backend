# memo-app-backend

```
chmod +x ./ssh.sh ./deploy.sh ./gcp-script.sh
```

```
./ssh.sh
```

```
make deploy
```

※立ち上げたVMインスタンスで初回デプロイする方法
・まずは`make ssh`でインスタンスにSSH接続する。
```
sudo apt update
sudo apt install rsync
sudo apt install docker.io
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER
```
ここでログインし直す
```
sudo curl -L "https://github.com/docker/compose/releases/download/v2.20.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```