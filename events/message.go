package events

const (
	MsgUnknown = "Я вас не понимаю! Пишите понятнее!!!"
	MsgStart   = "Привет пользователь!" +
		"\nЯ бот 'Вкус и Точка'. Первый, настоящий, официальный! Я готов принять ваши заказы" +
		"\nДля открытия меню напишите /menu" +
		"\nДля заказа напишите /order"
	MsgMenu = "Меню ресторана:" +
		"\n1. Жаренная свинина в майонезе - 150 р" +
		"\n2. Картошка фри - 35 р" +
		"\n3. Напиток - 60 р"
	MsgHelp = "О, вам нужна помошь по командам? Отлично!" +
		"\n/start - запускает бота" +
		"\n/menu - показать меню ресторана" +
		"\n/help - получить справку" +
		"\n/order: <Заказ> - сделать заказ"
	MsgOrder     = "Мы пока не принимаем заказы"
	MsgSaveOrder = "Ваш заказ сохранён!"
	MsgHelpOrder = "Ошибка синтаксиса! Неуказан знак разделения : или указано больше :, чем 1!!! Введите /help для +" +
		"получения справки о синтаксисе"
	MsgErrorOrder = "Произошла ошибка сохранения! Обратитесь позже!!!"
)
