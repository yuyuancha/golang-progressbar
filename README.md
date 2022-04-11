# golang-progressbar
練習在 golang 實現打印進度條的功能。

## 簡述

透過 `github.com/schollz/progressbar/v3` 套件包來實現在終端器上打印進度條的功能。

## 安裝 docker 環境及執行程式

* clone Github repository

```
get clone https://github.com/yuyuancha/golang-progressbar.git
```

* 透過 docker-compose 開啟 docker 容器

```
docker-compose up -d
```

* 執行 main.go

```
docker-compose exec app go run main.go
```

* 關閉 docker 容器

```
docker-compose down
```
