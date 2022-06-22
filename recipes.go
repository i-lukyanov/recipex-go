package main

type res struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Picture     string   `json:"picture"`
	Recipe      string   `json:"recipe"`
}

type allRes []res

var Recipes = allRes{
	{
		ID:    "1",
		Title: "Диетическое овсяное печенье",
		Ingredients: []string{
			"Овсяные хлопья - 1 стакан",
			"Мука - 1/3 стакана",
			"Молотая корица - 2 ч.л.",
			"Разрыхлитель - 1 ч.л.",
			"Яблоко - 1 шт.",
			"Яичный белок - 2 шт.",
			"Ваниль - 0,5 ч.л. (по вкусу)",
			"Семена льна - 2 ст.л.",
		},
		Picture: "public/1.jpg",
		Recipe:  "Нам понадобится.Смешать сухие ингредиенты: овсяные хлопья, муку, семена льна, разрыхлитель и корицу.Белки отделить от желтков. Желток нам не потребуется. Яблоко натереть на мелкой тёрке и добавить к белкам.Смешать сухие ингредиенты с жидкими и очень хорошо перемешать. Оставить на 15 минут, чтобы хлопья немного набухли.Лепить печенье влажными руками. Скатать овсяное тесто в небольшой шарик и придавить его - получатся плоские печенюшки. Выложить на противень застеленный бумагой.Разогреть духовку до 220 градусов и выпечь печенье в течение 15 минут до золотистого цвета.",
	},
	{
		ID:    "2",
		Title: "Печенье из геркулеса",
		Ingredients: []string{
			"Овсяная мука - 0,5 стакана",
			"Овсяные хлопья - 0,5 стакана",
			"Банан - 0,5 шт.",
			"Яблоко - 1 шт.",
			"Рафинированное растительное масло - 3 ст.л.",
			"Сахар - 1-2 ст.л.",
			"Фундук - 50 грамм",
			"Сушеная вишня - 0,5 стакана",
			"Молотая корица - 0,5 ч.л.",
			"Разрыхлитель - 1 ч.л.",
			"Соль - щепотка",
			"Ванилин - на кончике ножа",
		},
		Picture: "public/2.jpg",
		Recipe:  "Подготовьте необходимые продукты. Сушеную вишню вымойте и просушите.Орехи порубите ножом.В миске соедините все сухие ингредиенты, фундук и сушеную вишню.Яблоко очистите от кожуры, удалите сердцевину и натрите на мелкой терке.Половину банана разомните вилкой.Соедините яблочное и банановое пюре, добавьте растительное масло и хорошо перемешайте.Соедините фруктовое пюре с сухими ингредиентами. Тщательно перемешайте и оставьте на 10-15 минут.Духовку разогрейте до 180 градусов. Противень застелите пергаментом.Чайной ложкой выкладывайте тесто на противень на небольшом расстоянии друг от друга. Выпекайте печенье около 30 минут.Готовое овсяное печенье из геркулеса остудите на решетке.",
	},
	{
		ID:    "3",
		Title: "Печенье овсяное на кефире",
		Ingredients: []string{
			"Овсяные хлопья - 300 г",
			"Кефир - 300 мл",
			"Мед - 3 ст.л.",
			"Изюм - 200 г",
			"Ванильный сахар - 10 г",
			"Корица - 0,5 ч.л.",
		},
		Picture: "public/3.jpg",
		Recipe:  "Есть определенный секрет у этого печенья: его надо обязательно полностью остудить перед употреблением, иначе оно не будет хрустящим.Подготовим продукты по списку.Вольем кефир в овсяные хлопья. Хорошо смешаем кефир с хлопьями и оставим смесь на 40 минут настояться. За это время овсяные хлопья размягчатся.Вольем в смесь мед.Добавляем ванильный сахар и корицу.Если у вас изюм сильно сухой, следует его предварительно размочить, залив кипятком на 10 минут. Затем воду слить, изюм просушить.Выкладываем изюм в миску к остальным ингредиентам. Хорошо вымешиваем тесто.Выкладываем ложкой полученное тесто на противень, немного придавливаем его и придаем ему круглую форму. Отправляем печенье в разогретую до 180 градусов С духовку на 30-35 минут.Овсяное печенье на кефире готово!",
	},
	{
		ID:    "4",
		Title: "Овсяное печенье с корицей",
		Ingredients: []string{
			"Мука пшеничная – 170 г",
			"Мука овсяная – 80 г",
			"Разрыхлитель – 7 г",
			"Сахар тростниковый – 150 г",
			"Сливочное масло – 85 г",
			"Ванильный экстракт – 2 мл",
			"Изюм – 35 г",
			"Корица молотая – 0.5 ч.л.",
			"Соль – щепотка",
		},
		Picture: "public/4.jpg",
		Recipe:  "Для приготовления овсяного печенья с корицей необходимо подготовить ингредиенты по списку.Изюм залить кипятком на 5-6 минут, а затем измельчить в блендере. Соединить сливочное масло комнатной температуры, сахар, ванильный экстракт, изюм, овсяную муку и корицу.Пшеничную муку смешать с разрыхлителем и просеять в тесто. Хорошо перемешать. Соль растворить в тёплой воде (60 мл) и постепенно влить в тесто. Ещё раз хорошо перемешать.Сформировать небольшие шарики и выложить на противень с пергаментом. Выпекать в разогретой до 200 градусов духовке 15 минут. Остудить на решётке.",
	},
	{
		ID:    "5",
		Title: "Постное овсяное печенье",
		Ingredients: []string{
			"Овсяные хлопья - 1,5 стакана",
			"Бананы – 2 шт.",
			"Мёд – 1,5 ст.л.",
			"Грецкие орехи - 1/4 стакана",
			"Сода пищевая – 1/2 ч.л. (погасить соком лимона)",
			"Корица – 1/2 ч.л",
		},
		Picture: "public/5.jpg",
		Recipe:  "Бананы очистить от кожуры и размять вилкой. Подойдут очень спелые, желтые бананы. Овсяные хлопья перемолоть в блендере, не очень мелко. Добавить к банановому пюре. Добавить мед - если нет жидкого, тогда нужно его немного растопить. Перемешать ингредиенты до однородной массы. Добавить измельченные грецкие орехи, корицу. Соду погасить несколькими каплями лимонного сока. Все перемешать и оставить на 20 минут настояться. На противень с пергаментной бумагой выложить слепленные печеньки. Мокрыми руками формируются отлично. Размер небольшой.Выпекать постное овсяное печенье из овсяных хлопьев при 200 градусах 15-20 минут, до золотистого цвета и приятного аромата.",
	},
	{
		ID:    "6",
		Title: "Флэпджек (овсяно-яблочное печенье в год лошади)",
		Ingredients: []string{
			"Овсяные хлопья - 250 грамм;",
			"Сливочное масло - 100 грамм;",
			"Сахар или демерара - около 100 грамм;",
			"Мед или патока - около 100 грамм;",
			"Яблоко крупное - 1 штука;",
			"Изюм - 0,3 стакана;",
			"Орехи или пралине - 2 ст. ложки;",
			"Вода - 1-2 ст. ложки.",
		},
		Picture: "public/6.jpg",
		Recipe:  "Подготовьте ингредиенты: Растопите сливочное масло на среднем огне. Добавьте сахар и мед и перемешайте до однородного состояния. Яблоко нарежьте мелкими кубиками. Соедините и перемешайте овсяные хлопья, с изюмом, кусочками яблок, орехами или пралине и сладкой масляной массой. При необходимости добавьте немного воды.Уложите массу на смазанный или проложенный пекарской бумагой противень, разровняйте.Запекайте в духовке при 180 градусах 15-20 минут, затем разлинуйте и дайте полностью остыть. Остывший флэпджек нарежьте на квадраты. Можно хранить это овсяно-яблочное печенье до одной недели.",
	},
	{
		ID:    "7",
		Title: "Постное морковно-овсяное печенье с орехами и изюмом",
		Ingredients: []string{
			"Мука цельнозерновая – 100-150 г",
			"Овсяные хлопья – 100 г",
			"Морковь – 200 г",
			"Яблоко – 100-150 г (1 шт.)",
			"Банан (спелый) – 1 шт.",
			"Изюм – 150 г",
			"Грецкий орех - 50-100 г",
			"Яблочный сок – 50 мл",
			"Разрыхлитель – 1.5 ч.л.",
			"Корица молотая – 1 ч.л.",
			"Имбирь молотый – 0.5 ч.л.",
			"Ванильный сахар/ванилин – по вкусу",
			"Кунжут – 3 щепотки (по желанию)",
		},
		Picture: "public/7.jpg",
		Recipe:  "Разомните 1 крупный спелый банан до состояния пюре. Добавьте натертые на мелкой терке морковь и яблоко и тщательно все перемешайте. Поместите изюм, орехи и овсяные хлопья в чашу блендера и измельчите до состояния мелкой крошки.Добавьте получившуюся массу к смеси моркови, банана и яблока. Добавьте по вкусу специи. Влейте яблочный сок, и добавляя постепенно, вмешайте в смесь цельнозерновую муку и разрыхлитель. Сформируйте из теста небольшие лепешки и выложить на противень.Посыпать заготовки кунжутом. Подготовленное печенье поместите в духовку, разогретую до 175 градусов. Выпекайте печенье 25-30 минут.",
	},
	{
		ID:    "8",
		Title: "Шоколадно-овсяное печенье",
		Ingredients: []string{
			"Яйцо куриное (мелкое) - 1 шт.",
			"Мука пшеничная - 50 г",
			"Масло сливочное - 50 г",
			"Шоколад черный - 50 г",
			"Овсяные хлопья (геркулес) - 50 г",
			"Сахар - 30 г",
			"Какао-порошок - 1 ст.л.",
			"Корица - 0.5 ч.л.",
			"Сода - на кончике ножа",
		},
		Picture: "public/8.jpg",
		Recipe:  "Для приготовления печенья муку просеять, добавить соду, какао и корицу.Масло, шоколад и сахар растопить на водяной бане.Соединить масло, шоколад и сахарК остывшей шоколадной массе добавить яйцо. Хорошенько перемешать.Добавить яйцоСоединить с сухими ингредиентами.Смешать сухую смесь с жидкойЛожкой тщательно перемешать.Высыпать овсяные хлопья.И снова перемешать, тесто получается мягкое.Влажными руками разделить тесто на 6 частей, скатать шарики и приплюснуть, сформировав печенье.Выложить тесто на противеньВыпекать при температуре 180 градусов 8-10 минут.",
	},
	{
		ID:    "9",
		Title: "Овсяное печенье с бананом",
		Ingredients: []string{
			"Овсяные хлопья – 200 г",
			"Банан – 1 шт.",
			"Изюм/сухофрукты – 50-100 г",
			"Растительное масло – 2 ст.л.",
			"Разрыхлитель – 0.5 ч.л.",
			"Семена льна – 1 ст.л.",
			"Кунжут – 1 ст.л.",
			"Соль – 1 щепотка",
			"Специи – по вкусу",
			"Сахар/мед – по желанию",
		},
		Picture: "public/9.jpg",
		Recipe:  "Залейте сухофрукты кипятком и дайте настояться несколько минут. Банан разомните вилкой или измельчите в блендере. К получившемуся банановому пюре добавьте растительное масло. Тщательно все перемешайте, до получения однородной смеси. Отмерьте овсяные хлопья, добавьте щепотку соли. По желанию, добавьте 2 ст.л. семян или измельченных орехов, по вкусу. Тщательно все перемешайте.Соединить семена льна, овсяные хлопья и кунжутПостепенно, помешивая, добавьте в банановое пюре подготовленную смесь сухих компонентов. Тщательно все перемешайте.Добавьте сухофрукты. Добавьте разрыхлитель и специи по вкусу. Я добавляю немного мускатного ореха. Форму для запекания выстелите пекарской бумагой. Смажьте бумагу тонким слоем растительного масла. Сформируйте лепешки. Подготовленное печенье поместите в разогретую до 170 градусов духовку и выпекайте 25-30 минут, до золотистого цвета.",
	},
	{
		ID:    "10",
		Title: "Печенье овсяное с творогом",
		Ingredients: []string{
			"Овсяные хлопья - 1 стакан",
			"Творог - 100 г",
			"Сахар - 100 г",
			"Разрыхлитель теста - 1 ч.л.",
			"Яйцо куриное - 1 шт.",
			"Сливочное масло - 30 г",
			"Корица - 10 г",
		},
		Picture: "public/10.jpg",
		Recipe:  "Подготовим все ингредиенты по списку. Качество творога очень влияет на структуру печенья, лучше брать творог, а не творожную массу. Также на структуру печенья влияет и качество овсяных хлопьев, у меня хлопья - очень нежные и тонкие, поэтому печенье будет намного мягче и влажнее. В миску выкладываем овсяные хлопья, вбиваем яйцо, добавляем сахар. Всыпаем разрыхлитель теста и корицу молотую. В самом конце добавляем творог и размягченное сливочное масло (или маргарин). Хорошо миксером или руками вымешиваем массу.Руки смачиваем в воде и из теста формуем шарики размером с грецкий орех. Выкладываем шарики на противень, застеленный фольгой или бумагой для выпечки. Отправляем печенье в разогретую до 180 градусов С на 20-25 минут.Готовое печенье достаем из духовки и даем ему немного остыть на противне.",
	},
	{
		ID:    "11",
		Title: "Печенье Анзак",
		Ingredients: []string{
			"Мука пшеничная - 100 г",
			"Овсяные хлопья - 100 г",
			"Кокосовая стружка - 65 г",
			"Сахар - 100 г",
			"Мед - 1 ст.л.",
			"Сливочное масло - 100 г",
			"Сода - 1,5 ч.л.",
			"Вода - 2 ст.л.",
		},
		Picture: "public/11.jpg",
		Recipe:  "Итак, подготовим все необходимые ингредиенты для печенья Анзак.В миске смешаем сухие ингредиенты: муку, овсяные хлопья, кокосовую стружку.Добавим сахар и разрыхлитель для теста. Хорошо перемешаем все сухие ингредиенты.Сливочное масло и мед растопим в микроволновой печи и вольем в сухую массу.Добавим воду и теперь руками хорошо вымешиваем массу, чтобы получилась маслянистая крошка.Формируем шарики, а затем на противне приплюснем шарики руками, чтобы получились лепешки. Тесто будет крошиться, но не стоит добавлять лишней воды, просто плотно руками придавайте форму печенью. Отправляем противень с печеньем Анзак в разогретую до 170 градусов С духовку на 20 минут.Готовому печенью даем полностью остыть на противне и только потом перекладываем его в банку или коробку.",
	},
	{
		ID:    "12",
		Title: "Овсяное печенье по ГОСТу",
		Ingredients: []string{
			"Маргарин (масло) - 99 грамм;",
			"Сахар - 147 грамм;",
			"Изюм темный - 10 грамм;",
			"Корица - 4 грамма;",
			"Ванильный сахар - 4 грамма;",
			"Мука овсяная (хлопья) - 85 грамм;",
			"Мука пшеничная - 198 грамм;",
			"Сода - 2 грамма;",
			"Патока (мед) - 17грамм;",
			"Вода - 50 мл;",
			"Соль - 2 грамма.",
		},
		Picture: "public/12.jpg",
		Recipe:  "Если вы не любите слишком сладкое, то положите сахара поменьше, я беру ровно половину.Смешаем маргарин с сахаром, измельченным изюмом, корицей и ванилью.Добавим овсяную муку.Затем в горячую воду положим соль и добавим в тесто, потом патоку, муку с содой. Быстро замесим однородное теплое тесто.Раскатаем в пласт толщиной 1 см и вырежем печенье.Выпекаем в разогретой до 180 градусов духовке 10-12 минут не дожидаясь румяной корочки, чтобы не пересушить.",
	},
	{
		ID:    "13",
		Title: "Овсяное печенье без масла",
		Ingredients: []string{
			"Хлопья овсяные - 200 грамм;",
			"Мёд - 3 ст.л;",
			"Яйцо куриное - 2 шт;",
			"Сода - 0.5 ч.л;",
			"Мука пшеничная - 3 ст.л.",
		},
		Picture: "public/13.jpg",
		Recipe:  "Соединяем куриные яйца и мёд. Хорошо взбиваем кухонным венчиком.Добавляем соду. Хорошо перемешиваем.Овсяные хлопья измельчаем в блендере в мелкую крошку. Частями добавляем к яичной смеси. Перемешиваем венчиком.Вводим пшеничную муку. Хорошо перемешиваем. Должно получится густоватое тесто.Форму для выпекания застилаем пергаментной бумагой. Ложкой выкладываем тесто небольшими горками, на небольшом расстоянии друг от друга, так как печенье ещё вырастет. Отправляем в горячую духовку на 190-200 градусов на 10-15 минут.Овсяное печенье без масла готово.",
	},
	{
		ID:    "14",
		Title: "Фитнес-батончики",
		Ingredients: []string{
			"Овсяные хлопья - 200 г",
			"Банан – 1 шт.",
			"Растительное масло – 2 ст.л.",
			"Кунжут – 1 ст.л.",
			"Семена льна – 1 ст.л.",
			"Соль – 1 щепотка",
			"Разрыхлитель – 0.5 ч.л.",
			"Сухофрукты – 2-3 ст.л.",
			"Какао – 2 ч.л.",
			"Кофе натуральный молотый – 1 ч.л.",
			"Корица - по вкусу",
			"Мускатный орех – по вкусу",
			"Шоколадные капли – 1 ст.л. (по желанию)",
			"Белый /черный шоколад – 2-3 кусочка (для украшения)",
			"Мед/сахар – по необходимости",
		},
		Picture: "public/14.jpg",
		Recipe:  "Сухофрукты залейте кипятком и дайте настояться несколько минут.Овсяные хлопья, периодически помешивая, поджарьте в духовке 5–10 минут при температуре 180 градусов, до золотистого цвета и появления аромата.Соедините сухие компоненты. Добавьте специи по вкусу. Банан нарежьте небольшими кусочками, а затем разомните вилкой. К банановому пюре добавьте растительное масло и тщательно перемешайте.Добавьте сухие компоненты к жидким, разрыхлитель и 2-3 ст. л сухофруктов, орехов. Тщательно все перемешайте.Выложите смесь в форму, разровняйте и утрамбуйте ложкой или толкушкой для картофеля.Поместите подготовленную смесь в разогретую до 175 градусов духовку и выпекайте 15 минут до золотистого цвета, и охладите.",
	},
	{
		ID:    "15",
		Title: "Овсяное печенье с финиками и орехами",
		Ingredients: []string{
			"Овсяные хлопья - 120 граммов;",
			"Яйцо - 1 штука;",
			"Сахар - 30 граммов;",
			"Сливочное масло - 60 граммов;",
			"Ванильный сахар - 1 ч.л.;",
			"Разрыхлитель - 1 ч.л.;",
			"Финики - 8-10 штук;",
			"Измельченные грецкие орехи - 2 ст.л.",
		},
		Picture: "public/15.jpg",
		Recipe:  "Финики можно заменить сушеной клюквой или вишней, добавив печенью немного кислинки, или любыми другими сухофруктами на ваш вкус.Подготовьте необходимые продукты.Из фиников удалите косточки. Нарежьте финики мелкими кубиками.Яйцо взбейте с сахаром.Овсяные хлопья смешайте с ванильным сахаром и разрыхлителем.Сливочное масло растопите, охладите и добавьте к овсяным хлопьям. Перемешайте.Добавьте в овсяное тесто взбитое с сахаром яйцо, финики и грецкие орехи. Хорошо перемешайте.Руки смочите холодной водой и сформируйте из теста печенье. Противень застелите пергаментом, переложите печенье и выпекайте в разогретой до 180 градусов духовке 15-20 минут.Если хотите, чтобы печенье было мягким внутри, не держите его долго в духовке. Для более твердого печенья увеличьте время выпечки примерно на 10 минут.",
	},
	{
		ID:    "16",
		Title: "Морковное печенье с овсяными хлопьями",
		Ingredients: []string{
			"Мука пшеничная - 80 г",
			"Хлопья овсяные - 0.5 стакана",
			"Морковь - 100 г",
			"Орехи грецкие - 1 ст.л.",
			"Изюм - горсть",
			"Мёд - 1 ст.л.",
			"Разрыхлитель - 0.25 ч.л.",
			"Сахар - 1-2 ст.л.",
			"Масло оливковое - 30 мл",
		},
		Picture: "public/16.jpg",
		Recipe:  "Для приготовления печенья подготовить из списка необходимые продукты.Морковь следует почистить и натереть на мелкой тёрке. Добавить к моркови оливковое масло, промытый и обсушенный изюм, а также рубленые орехи. Добавить ложку мёда.Массу перемешать, добавить овсяные хлопья и перемешать ещё раз. Оставить на 30 минут. За это время хлопья размякнут.По истечении времени в эту массу добавить просеянную муку вместе с разрыхлителем и замесить тесто. Оно будет слегка липким.Тесто разделить на 16 одинаковых частей. Я делю так: сначала на 2 части, потом каждую - снова на 2 части, и так далее. Из каждого кусочка скатать шарик.В блюдце насыпать сахар и обвалять в нём каждый шарик.  Раскладываем печенье на пергамент.Духовку разогреть до 180 градусов и выпекать морковное печенье с овсяными хлопьями около 20 минут. Оно должно зарумяниться, но не передержите, чтоб оно оставалось мягким.",
	},
	{
		ID:    "17",
		Title: "Печенье из тыквы и овсяных хлопьев",
		Ingredients: []string{
			"Тыква - 200 г",
			"Мука цельнозерновая - 150 г",
			"Масло растительное - 100 мл",
			"Хлопья овсяные - 70 г",
			"Сахар - 80 г",
			"Разрыхлитель - 1 ч.л.",
		},
		Picture: "public/17.jpg",
		Recipe:  "Подготовим все компоненты для приготовления овсяного печенья с тыквой по списку.Отвариваем тыкву.Прожариваем хлопья 10 минут, перемешивая.Соединяем хлопья и тыкву, измельчаем блендером.Вводим сахар, муку, разрыхлитель.Вливаем масло, замешиваем тесто. Получается мягкая масса, структурная за счет хлопьев.Формируем печенье и выпекаем при 180 С 20-30 минут.Печенье из тыквы и овсяных хлопьев готово.",
	},
	{
		ID:    "18",
		Title: "Постное лимонное печенье с овсяными хлопьями",
		Ingredients: []string{
			"Овсяные хлопья - 50 грамм;",
			"Лимон - 1 шт;",
			"Мука пшеничная - 1.5 стакана;",
			"Сода - 0.5 ч.л;",
			"Растительное масло - 100 мл;",
			"Сахар - 100 грамм.",
		},
		Picture: "public/18.jpg",
		Recipe:  "Овсяные хлопья следует измельчить в комбайне с помощью насадки \"острый нож\".Лимон нарезать кусочками,вынуть косточки и вместе со шкуркой измельчить в комбайне.Переложить лимонную массу в миску, добавить к лимону сахар и растительное масло, перемешать и добавить измельчённые хлопья.Муку смешать с содой и постепенно добавить к основной массе.Замесить тесто.Затем его раскатать и с помощью формочки вырезать фигурки.Застелить противень пергаментом и выложить печенье рядами.Разогреть духовку до 200 градусов и запекать печенье 15 минут до румяного цвета.",
	},
	{
		ID:    "19",
		Title: "Хаврекакур (шведское овсяное печенье)",
		Ingredients: []string{
			"Овсяные хлопья (быстрого приготовления) - 500 мл",
			"Сахар - 1 стакан (200 мл)",
			"Мука пшеничная - 1 стакан (200 мл)",
			"Разрыхлитель - 1 ч.л.",
			"Ванильный сахар - 2 ч.л.",
			"Масло сливочное - 200 г",
		},
		Picture: "public/19.jpg",
		Recipe:  "Сливочное масло положите в миску и растопите в микроволновке либо на водяной/паровой бане. Включите духовку на 180 градусов.Добавьте в растопленное масло овсяные хлопья быстрого приготовления.Положите в миску пшеничную муку, сахарный песок, ванильный сахар и разрыхлитель.Перемешайте до получения однородной массы. Тесто будет рассыпчатым, напоминающим немного мокрый песок.Утрамбуйте тесто в миске. Примерно разделите на 4 части.Сформируйте на противне, застеленном пергаментной бумагой, 4 \"холмика-колбаски\" на длину обычного противня и плотно придавите их. Каждый \"холмик\" придавите сверху ладонью, т.е. приплюсните каждую из \"колбасок\".Подготовленное таким образом тесто поставьте в духовку, разогретую до 180 градусов, на 10-12 минут.Тесто должно начать подрумяниваться по краям. Выньте печенье из духовки, дайте ему остыть 10 минут. Остывшее, но еще теплое печенье нарежьте наискось на порционные кусочки-печенюшки.",
	},
}
