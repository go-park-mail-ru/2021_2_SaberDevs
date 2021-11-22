package data

import (
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
)

var TestData = [...]amodels.ArticleData{
	{"keyboard-254582_1920.jpg", []string{"IT-News", "Study"}, "Apple полгода не выпускала исправление уязвимости нулевого дня, хотя эксплойт для неё находился в открытом доступе",
		"Как сообщает портал SecurityLab, хакеры использовали политические новостные сайты Гонконга для заражения компьютеров на macOS бэкдором путём использования двух уязвимостей, одна из которых была ранее неизвестна. Атаки на устройства начались по меньшей мере в апреле этого года. " +
			"Тогда в сети появился эксплойт для уязвимости нулевого дня, однако Apple исправила проблему только в сентябре. Первая из двух уязвимостей CVE-2021-1789 — уязвимость удаленного выполнения кода в WebKit. Вторая CVE-2021-30869 — это уязвимость локального повышения привилегий в " +
			"компоненте ядра XNU, которую исправили в сентябре. С их помощью злоумышленники получали на macOS привилегии суперпользователя и загружали на устройство шпионские программы MACMA или OSX.CDDS. Загруженные программы могут записывать аудио, делать скриншоты, загружать и выгружать " +
			"файлы, записывать нажатия кнопок на клавиатуре, создавать отпечаток устройства для его идентификации и выполнять команды терминала.\n" +
			"Саму атаку обнаружили специалисты Threat Analysis Group (TAG) компании Google, которые и сообщили Apple о необходимости исправить уязвимость. Они также указали, что злоумышленники используют другую связку уязвимостей для атаки на iOS-устройства, но её пока не удалось распознать. " +
			"Специалисты TAG рассказали, что за атаками хакеров стояла хорошо финансируемая группа, вероятно, работающая на правительство и имеющая доступ к собственной команде инженеров программного обеспечения.	Несмотря на то, что эксплойт для CVE-2021-30869 находится в открытом доступе с " +
			"апреля, Apple очень долго не выпускала обновление для исправления уязвимости. О наличии проблемы сообщала не только Google, но и исследователи из Pangu Lab. Кроме того, эксплоит был представлен на Mobile Security Conference (MOSEC) в июле. " +
			"Согласно заметкам релиза сентябрьского обновления с исправлением уязвимостей, упомянутому ранее, полный список устройств, на которых закрыли CVE-2021-30869, включает iPhone 5s, iPhone 6, iPhone 6 Plus, iPad Air, iPad mini 2, iPad mini 3, iPod touch (6-го поколения) " +
			"под управлением iOS 12.5.5 и Mac с обновлением безопасности 2021-006 Catalina. Проблема исправлена в iOS 12.5.5, iOS 14.4 и iPad OS 14.4 и mac OS Big Sur 11.2. В обновлении также вышли патчи для уязвимостей нулевого дня, впервые обнаруженных в феврале " +
			"(CVE-2021-1870, CVE-2021-1871, CVE-2021-1872), марте (CVE-2021-1879), мае (CVE-2021-30663, CVE-2021-30665, CVE-2021-30713, CVE-2021-30666) и июне (CVE-2021-30761 и CVE-2021-30762).", "#", "Григорий", "user.jpg",
		"#", 97, 1001,
	},
	{"startup-593343_1920.jpg", []string{"IT-News", "Study"}, "Российская компания вышла в фина л международного конкурса Entrepreneurship World Cup (EWC) с призовым фондом $1 млн",
		"Компании из РФ нередко участвуют в международных конкурсах, занимая призовые места. В частности, платформа Botkin.AI резидента Фонда «Сколково» (Группа ВЭБ.РФ) «Интеллоджик» прошла в финал международного конкурса Entrepreneurship World Cup (EWC). Финал состоится в декабре 2021 в Саудовской Аравии." +
			"Компания стала единственной организацией из РФ, которая прошла финальный отбор. К слову, в этом году на EWC подано свыше 100 тыс заявок из 200 стран. Теперь 25 финалистов отправятся в Эр-Рияд для презентации своего продукта на заключительном этапе конкурса. Что касается российской платформы, " +
			"то она дает возможность при помощи ИИ анализировать медицинские снимки, выявляя патологические изменения." +
			"Подробности о проекте\n" +
			"Платформа Botkin.AI легко интегрируется с медицинским программным обеспечением и оборудованием в любых клиниках. Это первая российская разработка в этой области, получившая международный сертификат CE Mark. Использование платформы Botkin.AI позволяет повысить эффективность медицинской диагностики, " +
			"снизить нагрузку на врачей-рентгенологов и риск пропуска патологий даже на ранних стадиях, а также обеспечить 100%-й контроль качества описания диагностических изображений. Сейчас решение компании успешно используется в различных регионах России, а также зарубежных странах. Ежегодно Botkin.AI реализует " +
			"Еще немного подробностей\n" +
			"Сергей Сорокин, генеральный директор Botkin.AI: «Мы благодарны организаторамGlobal Entrepreneurship Week за возможность представить наш проект в финале. И, безусловно, гордимся нашей фантастической командой, которая сделала платформу Botkin.AI – проект, который был выбран среди сотни тысяч других компаний со всего мира». " +
			"Entrepreneurship World Cup – крупнейшая конкурсная программа поддержки предпринимателей. EWS направлен на помощь компаниям, находящимся на разных стадиях развития. С момента запуска в 2019 году проект помог более 300 000 участникам из 200 стран, предоставив финансовую и профессиональную помощь. " +
			"Юлия Щеглова, проектный менеджер Фонда «Сколково»: «“Интеллоджик” достоин представлять нашу страну в конкурсной программе международного уровня. У компании накопился солидный опыт в разработке и 3 патента на технологии искусственного интеллекта. Качество платформы Botkin AI было проверено врачами в 25 проектах по всему миру». ",
		"#", "Григорий", "user.jpg", "#", 97, 1002},

	{"usb-1284227_1920.jpg", []string{"IT-News", "Study"}, "Abcbot — новый ботнет, нацеленный на Linux",
		"Исследователи кибербезопасности из Netlab Qihoo 360 раскрыли подробности о новом растущем ботнете «Abcbot», который распространяется подобно червям, заражая системы Linux с последующим запуском распределенных DDoS-атак. " +
			"Первая версия вредоноса датируется июлем 2021 года, но 30 октября были замечены его новые разновидности, заточенные под атаки на слабозащищенные серверы Linux с уязвимостью нулевого дня. Все это говорит о том, что вредонос постоянно совершенствуется. " +
			"В основу выводов Netlab также лег отчет Trend Micro от начала сентября, в котором описывались атаки с использованием криптоджекинга, нацеленные на Huawei Cloud. Эти вторжения, помимо прочего, выделились тем, что вирусные скрипты оболочки, в частности, " +
			"отключали процесс мониторинга и сканирования серверов на проблемы безопасности, а также сбрасывали пароли пользователей от сервиса Elastic Cloud. " +
			"Теперь, как сообщает китайская компания Qihoo 360, эти скрипты используются для распространения Abcbot. Всего на сегодня было зафиксировано шесть версий этого ботнета. " +
			"После установки на захваченную систему вредонос запускает выполнение серии шагов, в результате которых зараженное устройство переопределяется в веб-сервер. Далее, помимо передачи системной информации C&C-серверу, происходит распространение вируса на " +
			"другие устройства путем сканирования открытых портов. При этом он также автоматически обновляется, когда операторы вносят доработки.\n", "#", "Григорий", "user.jpg",
		"#", 97, 1003,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 1004,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 1005,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 1006,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 1007,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 14,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 1008,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 1009,
	},
	{"computer.png", []string{"IT-News", "Study"}, "7 Skills of Highly Effective Programmers",
		"Our team was inspired by the seven skills of highly effective", "#", "Григорий", "user.jpg",
		"#", 97, 1010,
	},
}

var TestUsers = []amodels.Author{
	{1, "mollenTEST1", "mollenTEST1", "mollenTEST1", "user.jpg", "", "mollenTEST1", "mollenTEST1", 123456},
	{2, "dar", "darivush", "pavlov", "user.jpg", "", "dar@exp.ru", "123", 13553},
	{3, "viphania", "viphania", "pavlova", "user.jpg", "", "viphania@exp.ru", "123", 120},
	{4, "DenisTest", "DenisTest1", "DenisTest1", "user.jpg", "", "DenisTest1@exp.ru", "DenisTest1", 120},
}

var End = amodels.Preview{"end", "123456", "endOfFeed.png", []string{"IT-News", "Study"},
	"А всё, а раньше надо было", "А всё, а раньше надо было", "SaberDevs", amodels.Author{}, "#", 0, 0}

var CategoriesList = []string{
	"Маркетинг",
	"Личный опыт",
	"Вопросы",
	"SaberDevs",
	"Сервисы",
	"Будущее",
	"Финансы",
	"Дизайн",
	"Соцсети",
	"Техника",
	"Карьера",
	"Медиа",
	"Истории",
	"Торговля",
	"SEO",
	"Офлайн",
	"Видео",
	"Офис",
	"Офтоп",
	"Транспорт",
	"Право",
	"Трибуна",
	"Еда",
	"Конкурсы",
	"Приёмная",
	"Миграция",
}
