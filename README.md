# :star: 2021_2_SaberDevs :star:
Backend-репозиторий команды SaberDevs

Project: vc.ru <br/>
[Репозиторий фронтенда](https://github.com/frontend-park-mail-ru/2021_2_SaberDevs)

## Команда
:diamonds: [Турчин Денис](https://github.com/Denactive) -> Frontend <br/>
:gem: [Любский Юрий](https://github.com/yurij-lyubskij) -> Backend <br/>
:diamonds: [Очеретная Светлана](https://github.com/Svetlanlka) -> Frontend <br/>
:gem: [Аристов Алексей](https://github.com/MollenAR) -> Backend <br/>

## Менторы
:diamonds: [Антон Елагин](https://github.com/AntonElagin) <br/>
:gem: [Екатерина Придиус](https://github.com/pringleskate)

## API
**Ошибки** <br/>
общий вид ошибки с её описание, могут приходить на любой запрос
````
{
    "status": не 200
    "error": "описание ошибки"
}
````

**Логин:** POST /api/v1/user/login

запрос:
````
{
    "email": "emal@emal.com"
} 
````
ответ:
````
{
    "status": 200
    "body" : 
    {
        "ID": 123
        "name": "name"
        "avatar": jpg?? пока нету
    }
}
````

**Логаут:** POST /api/v1/user/logout

запрос:
````
{
    "email": "emal@emal.com"
    "password": "password_name"
} 
````
ответ:
````
{
    "status": 200
    "GoodbuyMsg": "Goodbuy, mollen@exp.ru!"
    "body":
}
````

**Регистрация:** POST /api/v1/user/register

запрос:
````
{
    "email": "emal@emal.com"
    "password": "password_name"
} 
````
ответ:
````
{
    "status": 200
    "body":
    {
        "ID": 123
        "name": "name"
        "avatar": jpg??
    }
}
````
**Статьи(не готовы!!):** GET /api/v1/aricles/

запрос:
````
{
    "body":
    {
        "email": "emal@emal.com"
        "password": "password_name"
    }
} 
````
ответ:
````
{
    "status": 200
    "body":
    {
        "ID": 123
        "name": "name"
        "avatar": jpg??
    }
}
````
