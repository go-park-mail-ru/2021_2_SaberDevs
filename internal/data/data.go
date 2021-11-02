package data

import (
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
)

var TestData = [...]amodels.Article{
	{"1", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1001,
	},
	{"2", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1002,
	},
	{"3", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1003,
	},
	{"4", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1004,
	},
	{"5", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1005,
	},
	{"6", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1006,
	},
	{"7", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1007,
	},
	{"8", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 14,
	},
	{"9", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1008,
	},
	{"10", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1009,
	},
	{"11", "static/img/computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "static/img/photo-elon-musk.jpg",
		"#", 97, 1010,
	},
	{"12", "static/img/endOfFeed.png", []string{"IT-News", "Study"},
		"А всё, а раньше надо было", "А всё, а раньше надо было", "#", "Tester-ender",
		"static/img/loader-1-HorizontalBalls.gif", "#", 0, 0,
	},
}

var TestUsers = []umodels.User{
	{"mollenTEST1", "mollenTEST1", "mollenTEST1", "mollenTEST1", "mollenTEST1", 123456},
	{"dar", "dar@exp.ru", "dar@exp.ru", "dar@exp.ru", "123", 13553},
	{"viphania", "viphania@exp.ru", "viphania@exp.ru", "viphania@exp.ru", "123", 120},
	{"DenisTest", "DenisTest1", "DenisTest1", "DenisTest1@exp.ru", "DenisTest1", 120},
}
curl -X 'POST' \
'http://89.208.197.247:8081/api/v1/user/login' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '{
"login": "string",
"password": "string"
}'

curl -X 'POST' \
  'http://89.208.197.247:8081/api/v1/user/profile/update' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  --cookie "session=382e740f-c0b7-40df-acb7-7177dafc8a18" \
  -d '{
  "password": "string",
  "firstName": "string",
  "lastName": "string"
}'