
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
    "password: "123"
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
**Статьи(не готовы!!):** GET /api/v1/getfeed/

запрос:
````
{
   "from":0
   
   "to":2
} 
````
ответ:
````
{
	"status":200,
	"body": {
		"from":"0",
		"to":"2",
		"chunk":[
		{"id":"1","previewUrl":"static/img/computer.png","title":"7 Skills of Highly Effective Programmers","text":"Our team was inspired by the seven skills of 			highly effective","authorUrl":"#","authorName":"Григорий","authorAvatar":"static/img/photo-elon-musk.jpg","commentsUrl":"#",
		"comments":97,"likes":1001,"tags":["IT-News","Study"]},
		{"id":"2","previewUrl":"static/img/computer.png","title":"7 Skills of Highly Effective Programmers","text":"Our team was inspired by the seven skills of 			highly effective","authorUrl":"#","authorName":"Григорий","authorAvatar":"static/img/photo-elon-musk.jpg","commentsUrl":"#",
		"comments":97,"likes":1002,"tags":["IT-News","Study"]}
		]
		}
}

````
