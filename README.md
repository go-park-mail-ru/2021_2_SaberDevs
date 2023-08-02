# :star: 2021_2_SaberDevs :star:
Backend-репозиторий команды SaberDevs

Project: vc.ru <br/>
[Репозиторий фронтенда](https://github.com/frontend-park-mail-ru/2021_2_SaberDevs)

Ссылка на деплой:<br/>
http://87.228.2.178:8080/

## Команда
:diamonds: [Турчин Денис](https://github.com/Denactive) -> Frontend <br/>
:gem: [Любский Юрий](https://github.com/yurij-lyubskij) -> Backend <br/>
:diamonds: [Очеретная Светлана](https://github.com/Svetlanlka) -> Frontend <br/>
:gem: [Аристов Алексей](https://github.com/MollenAR) -> Backend <br/>

## Менторы
:diamonds: [Антон Елагин](https://github.com/AntonElagin) <br/>
:gem: [Екатерина Придиус](https://github.com/pringleskate)

## Файлы в репозитории
- `configs/config.json` - настройки для сервера
- `createdb.go` - скрипт для инициализации БД
- `build/` - содержит Dockerfile'ы отдельных микросервисов
- `run.sh` - содержит скрипт для запуска отдельных микросервисов одной командой в терминале без использования Docker

## Запуск
С запуском все очень сложно. По мере добавления новых функций, скрипт инициализации БД не обновлялся, поэтому неизвестно, как должна выглядеть финальная версия БД.

Самая свежая рабочая версия скрипта для базы находится в ветке [fkey](https://raw.githubusercontent.com/go-park-mail-ru/2021_2_SaberDevs/fkey/createdb.go). Только надо создать файл `configs/config.json`.

Создаются таблицы:
```bash
alexeydb=# \dt
            List of relations
 Schema |     Name      | Type  | Owner  
--------+---------------+-------+--------
 public | articles      | table | alexey
 public | author        | table | alexey
 public | categories    | table | alexey
 public | tags          | table | alexey
 public | tags_articles | table | alexey
(5 rows)
```
Тест:
```bash
curl -X POST http://10.0.2.15:8081/api/v1/user/login -H 'Content-Type: application/json;charset=UTF-8' -d '{"login": "DenisTest", "email": "", "password": "DenisTest1"}'
```

### История добавления функций по веткам Git
За своевременным подливом веток с функциями в master никто не сследил. Поэтому имеем что имеем.

Формат: дата последнего коммита, название ветки, ссылка на ветку, наличие скриптов createdb.go x и comments.sql, автор функции.
- Nov 13 dbv9 https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/dbv9 createdb.go x
- Nov 21 imges https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/imges createdb.go x
- Nov 21 config https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/config createdb.go x
- Nov 22 microserv https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/microserv createdb.go  Yura
- Nov 23 fullrex https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/fullrex createdb.go x
- Nov 23 fulltext https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/fulltext createdb.go x Yura
- Nov 23 tosearch https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/tosearch createdb.go x MollenAR
- Nov 24 microservices https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/microservices createdb.go comments MollenAR
- Nov 25 micro https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/micro createdb.go comments Yura
- Nov 25 fkey https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/fkey   createdb.go comments
- Nov 25 metrics https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/metrics createdb.go comments Yura
- Dec 14 newmetrics https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/newmetrics x comments Yura
- Dec 14 tls https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/tls x comments MollenAR
- Dec 18 comments https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/comments x comments
- Dec 18 newlikes https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/newlikes x comments Yura
- Dec 19 logs https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/logs x comments Yura
- Dec 19 tests https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/tests x comments Yura
- Dec 21 dev https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/dev  x comments
- Dec 23 logger https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/logger x comments Yura
- Apr 22 (2022) docker https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/docker x comments MollenAR
- Коммит, который добавляет этот README в master. Основывается на ветке [docker](https://github.com/go-park-mail-ru/2021_2_SaberDevs/tree/docker).

### Подробная инструкция по запуску
Результат зависит от итогового выбранного функционала. По идее, рабочей версией является [fkey](https://raw.githubusercontent.com/go-park-mail-ru/2021_2_SaberDevs/fkey/createdb.go). Во всех остальных надо пробовать запускать, проверять, чтобы обращения из кода к БД соответстовало схеме БД.

Запускать сервер, СУБД Tarantool, СУБД Postgres можно на выбор: в докере, в терминале, на разных машинах, в т.ч. виртуальных. Запустить СУБД можно на Windows. Запустить golang-сервер под Windows не представляется возможным из-за C++ библиотеки *libwebp*. Можно попробовать скомпилировать ее из исходников под Windows, но у меня не получилось - непонятно, как подключить ее к одному из go-пакетов через cgo.

Инструкция ниже включает сборку нужных файлов из разных веток.

#### Установка библиотек
```bash
sudo apt-get update
sudo apt-get install libssl-dev
# если бы не эта либа, сервер был бы кроссплатформенным.
# Она есть на Windows. Можно даже легко скомпилировать в MSVS, но как затолкать ее в cgo решительно не ясно https://developers.google.com/speed/webp/docs/using?hl=ru
sudo apt-get install libwebp-dev 
```

#### Установка Postgres и настройка пользователей
```bash
sudo apt-get install postgresql postgresql-contrib
sudo -i -u postgres # смена текущего юзера на юзера postgres
psql # далее до выхода \q все внутри psql
create database postgres; # мне вывел, что уже есть
grant all privileges on database postgres to postgres; -->

create database alexeydb;
create USER alexey with password 'alexey';
grant all privileges on database alexeydb to alexey;

\q
Ctrl+D # выход обратно в оригинального пользователя
# подключение к базе от нового пользователя
psql -U alexey -d alexeydb
# вывод таблиц
\dt
```
Если возникает ошибка *error peer authentication failed*, надо изменить настройку
```bash
sudo nano /etc/postgresql/12/main/pg_hba.conf # может быть в другом месте. Воспользуйся командой locate pg_hba.conf
# меняем local all postgres peer на  local all postgres md5
```

#### Установка docker
```bash
# https://docs.docker.com/engine/install/ubuntu/
# все команды по ссылке выше
sudo apt  install docker-compose

curl -O -J -L https://github.com/docker/compose/releases/download/v2.11.2/docker-compose-linux-x86_64
chmod +x docker-compose-linux-x86_64
sudo cp ./docker-compose-linux-x86_64 /usr/bin/docker-compose
rm docker-compose-linux-x86_64
docker-compose version

sudo docker login  # с судо! обязательно

# делаем так, чтобы работало без sudo | https://docs.docker.com/engine/install/linux-postinstall/
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
# на виртуалке может потребоваться перезагрузка: If you’re running Linux in a virtual machine, it may be necessary to restart the virtual machine for changes to take effect.
```

#### Установка Go
```bash
# https://go.dev/doc/install
sudo wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
rm go1.20.4.linux-amd64.tar.gz
sudo nano /etc/profile # добавляем в самый конец export PATH=$PATH:/usr/local/go/bin | export PATH=$PATH:/usr/local/go/bin     # если по какой-то причине видишь, go: command not found - выполняй эту команду
# перезапуск
sudo visudo # найти строчку Defaults    secure_path = и в конец добавить :/usr/local/go/bin # https://stackoverflow.com/questions/71899341/go-command-not-found-error-when-using-sudo
go version
sudo go version # тест на то, что команда работает в sudo
```


#### Инициализируем таблицы в базе
```bash
git checkout fkey
mkdir config/
nano config/config.json # прописываем конфиг
go run create_db
git checkout docker
```

проверка:
```bash
psql -U alexey -d alexeydb -a -f comments.sql
psql -U alexey -d alexeydb -a -f likes.sql
psql -U alexey -d alexeydb
\dt
\q
```
вывод:
```bash
            List of relations
 Schema |     Name      | Type  | Owner  
--------+---------------+-------+--------
 public | articles      | table | alexey
 public | author        | table | alexey
 public | categories    | table | alexey
 public | comments      | table | alexey
 public | articles_likes| table | alexey
 public | tags          | table | alexey
 public | tags_articles | table | alexey
(7 rows)
```

#### Подготовка к запуску сервера
Нужно настроить конфиг и выбрать режим: HTTP / HTTPS.
```bash
# Объявляем модуль Go
rm go.mod go.sum
go mod init github.com/go-park-mail-ru/2021_2_SaberDevs
# Устанавливаем пакеты
go get ./... 
# мотаем в конец. Комментируем (//)
# либо startTLS, если нужен HTTP,
# либо e.Logger.Fatal(e.Start(address)), если нужен HTTPS.
# Во втором случае нужно указать аргументы - абсолютный путь к сертификатам: "/etc/ssl/sabernews.crt", "/etc/ssl/sabernews.key".
nano cmd/sybernews/server.go
# Настраиваем подключение к базе
# Указываем куда и под каким пользователем подключаться 
nano config/config.json
# Указываем директорию, куда будут помещены логи
# где-то в начале: "outputPaths": ["stdout", "logs"] - логи будут в локальной папке проекта в подпапке logs/
nano internal/dbWrappers.go
# запускаем
go run cmd/main.go
```
Если все ок, видим это
```bash
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on 127.0.0.1:8081
```
Останавливаем.


#### Запуск
1. В разных терминалах:

```bash
docker run -dp 3302:3302 tarantool # запустить тарантул. Я запуска в докере, но можно запустить локально со скриптом build/app.lua
go run cmd/main.go
# http server started on 127.0.0.1:8081
```

```bash
go run ./internal/comment/server/comment_app.go ./internal/comment/server/comment_server.go
# starting comment server at :8077
```

```bash
go run ./internal/user/server/user_app.go ./internal/user/server/user_server.go
# starting user server at :8078
```

```bash
go run ./internal/article/server/article_server.go
# starting server at :8079
```

2. Либо в двух
```bash
docker run -dp 3302:3302 tarantool # запустить тарантул. Я запуска в докере, но можно запустить локально со скриптом build/app.lua
go run cmd/main.go
# http server started on 127.0.0.1:8081
./run.sh # использовать новый, который ... & ... & ...
# starting comment server at :8077
# starting user server at :8078
# starting server at :8079
```

3. Либо в докере 
```bash
cd newDocker
docker-compose.up
```

Все это можно хостить на виртуальной машине. Внутри VM можно запускать как в докере, так и просто. Тогда нужно будет настроить в VirtualBox Nat-сеть (не NAT !), и подключить к гостевой ОС с такой настройкой проброса портов.
```md
127.0.0.1:8081 -> 10.0.2.15:8081
```
При этом в cmd/main.go нужно указать адрес 10.0.2.15


#### Тест 
(подставить сюда правильный ip)
```bash
curl -X POST http://10.0.2.15:8081/api/v1/user/login -H 'Content-Type: application/json;charset=UTF-8' -d '{"login": "DenisTest", "email": "", "password": "DenisTest1"}'
```


#### Запуск фронтенда
```bash
npm install
# закоментировать строчку # import ... .scss
nano public/main.js
npm run debug-no-webpack
```

