
## API
**Ошибки** <br/>
общий вид ошибки с её описание, могут приходить на любой запрос
````
{
    "status": не 200
    "ErrorMsg": "описание ошибки"
}
````

**Логин:** POST /login

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
    "data" : 
    {
        "login": "login"
        "name": "name"
        "surname": "surname"
        "email": "email"
        "score": 123
    }
    "msg": "OK"
}
````

**Логаут:** POST /logout

запрос:
````
{}
````
ответ:
````
{
    "status": 200
    "goodbye": "Goodbuy, mollen@exp.ru!"
}
````

**Регистрация:** POST /register

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
    "data" : 
    {
        "login": "login"
        "name": "name"
        "surname": "surname"
        "email": "email"
        "score": 123
    }
    "msg": "OK"
}
````
**Статьи:** GET /getfeed/

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
