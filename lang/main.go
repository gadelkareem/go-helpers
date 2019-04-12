package lang

import "strings"

type Lang struct {
	Name       string
	NativeName string
}

func CodeToLang(iso string) *Lang {
	iso = strings.ToLower(iso[0:2])
	if l, ok := isoLangs[iso]; ok {
		return &l
	}
	return nil
}

func CodeToLangName(iso string) string {
	l := CodeToLang(iso)
	if l != nil {
		return l.Name
	}
	return ""
}

func LangNameToCode(s string) string {
	s = strings.ToLower(s)
	for c, l := range isoLangs {
		if s == strings.ToLower(l.Name) ||
			s == strings.ToLower(l.NativeName) {
			return c
		}
	}
	return ""
}

var isoLangs = map[string]Lang{
	"ab": {
		Name:       "Abkhaz",
		NativeName: "аҧсуа",
	},
	"aa": {
		Name:       "Afar",
		NativeName: "Afaraf",
	},
	"af": {
		Name:       "Afrikaans",
		NativeName: "Afrikaans",
	},
	"ak": {
		Name:       "Akan",
		NativeName: "Akan",
	},
	"sq": {
		Name:       "Albanian",
		NativeName: "Shqip",
	},
	"am": {
		Name:       "Amharic",
		NativeName: "አማርኛ",
	},
	"ar": {
		Name:       "Arabic",
		NativeName: "العربية",
	},
	"an": {
		Name:       "Aragonese",
		NativeName: "Aragonés",
	},
	"hy": {
		Name:       "Armenian",
		NativeName: "Հայերեն",
	},
	"as": {
		Name:       "Assamese",
		NativeName: "অসমীয়া",
	},
	"av": {
		Name:       "Avaric",
		NativeName: "авар мацӀ, магӀарул мацӀ",
	},
	"ae": {
		Name:       "Avestan",
		NativeName: "avesta",
	},
	"ay": {
		Name:       "Aymara",
		NativeName: "aymar aru",
	},
	"az": {
		Name:       "Azerbaijani",
		NativeName: "azərbaycan dili",
	},
	"bm": {
		Name:       "Bambara",
		NativeName: "bamanankan",
	},
	"ba": {
		Name:       "Bashkir",
		NativeName: "башҡорт теле",
	},
	"eu": {
		Name:       "Basque",
		NativeName: "euskara, euskera",
	},
	"be": {
		Name:       "Belarusian",
		NativeName: "Беларуская",
	},
	"bn": {
		Name:       "Bengali",
		NativeName: "বাংলা",
	},
	"bh": {
		Name:       "Bihari",
		NativeName: "भोजपुरी",
	},
	"bi": {
		Name:       "Bislama",
		NativeName: "Bislama",
	},
	"bs": {
		Name:       "Bosnian",
		NativeName: "bosanski jezik",
	},
	"br": {
		Name:       "Breton",
		NativeName: "brezhoneg",
	},
	"bg": {
		Name:       "Bulgarian",
		NativeName: "български език",
	},
	"my": {
		Name:       "Burmese",
		NativeName: "ဗမာစာ",
	},
	"ca": {
		Name:       "Catalan; Valencian",
		NativeName: "Català",
	},
	"ch": {
		Name:       "Chamorro",
		NativeName: "Chamoru",
	},
	"ce": {
		Name:       "Chechen",
		NativeName: "нохчийн мотт",
	},
	"ny": {
		Name:       "Chichewa; Chewa; Nyanja",
		NativeName: "chiCheŵa, chinyanja",
	},
	"zh": {
		Name:       "Chinese",
		NativeName: "中文 (Zhōngwén), 汉语, 漢語",
	},
	"cv": {
		Name:       "Chuvash",
		NativeName: "чӑваш чӗлхи",
	},
	"kw": {
		Name:       "Cornish",
		NativeName: "Kernewek",
	},
	"co": {
		Name:       "Corsican",
		NativeName: "corsu, lingua corsa",
	},
	"cr": {
		Name:       "Cree",
		NativeName: "ᓀᐦᐃᔭᐍᐏᐣ",
	},
	"hr": {
		Name:       "Croatian",
		NativeName: "hrvatski",
	},
	"cs": {
		Name:       "Czech",
		NativeName: "česky, čeština",
	},
	"da": {
		Name:       "Danish",
		NativeName: "dansk",
	},
	"dv": {
		Name:       "Divehi; Dhivehi; Maldivian;",
		NativeName: "ދިވެހި",
	},
	"nl": {
		Name:       "Dutch",
		NativeName: "Nederlands, Vlaams",
	},
	"en": {
		Name:       "English",
		NativeName: "English",
	},
	"eo": {
		Name:       "Esperanto",
		NativeName: "Esperanto",
	},
	"et": {
		Name:       "Estonian",
		NativeName: "eesti, eesti keel",
	},
	"ee": {
		Name:       "Ewe",
		NativeName: "Eʋegbe",
	},
	"fo": {
		Name:       "Faroese",
		NativeName: "føroyskt",
	},
	"fj": {
		Name:       "Fijian",
		NativeName: "vosa Vakaviti",
	},
	"fi": {
		Name:       "Finnish",
		NativeName: "suomi, suomen kieli",
	},
	"fr": {
		Name:       "French",
		NativeName: "français, langue française",
	},
	"ff": {
		Name:       "Fula; Fulah; Pulaar; Pular",
		NativeName: "Fulfulde, Pulaar, Pular",
	},
	"gl": {
		Name:       "Galician",
		NativeName: "Galego",
	},
	"ka": {
		Name:       "Georgian",
		NativeName: "ქართული",
	},
	"de": {
		Name:       "German",
		NativeName: "Deutsch",
	},
	"el": {
		Name:       "Greek, Modern",
		NativeName: "Ελληνικά",
	},
	"gn": {
		Name:       "Guaraní",
		NativeName: "Avañeẽ",
	},
	"gu": {
		Name:       "Gujarati",
		NativeName: "ગુજરાતી",
	},
	"ht": {
		Name:       "Haitian; Haitian Creole",
		NativeName: "Kreyòl ayisyen",
	},
	"ha": {
		Name:       "Hausa",
		NativeName: "Hausa, هَوُسَ",
	},
	"he": {
		Name:       "Hebrew (modern)",
		NativeName: "עברית",
	},
	"hz": {
		Name:       "Herero",
		NativeName: "Otjiherero",
	},
	"hi": {
		Name:       "Hindi",
		NativeName: "हिन्दी, हिंदी",
	},
	"ho": {
		Name:       "Hiri Motu",
		NativeName: "Hiri Motu",
	},
	"hu": {
		Name:       "Hungarian",
		NativeName: "Magyar",
	},
	"ia": {
		Name:       "Interlingua",
		NativeName: "Interlingua",
	},
	"id": {
		Name:       "Indonesian",
		NativeName: "Bahasa Indonesia",
	},
	"ie": {
		Name:       "Interlingue",
		NativeName: "Originally called Occidental; then Interlingue after WWII",
	},
	"ga": {
		Name:       "Irish",
		NativeName: "Gaeilge",
	},
	"ig": {
		Name:       "Igbo",
		NativeName: "Asụsụ Igbo",
	},
	"ik": {
		Name:       "Inupiaq",
		NativeName: "Iñupiaq, Iñupiatun",
	},
	"io": {
		Name:       "Ido",
		NativeName: "Ido",
	},
	"is": {
		Name:       "Icelandic",
		NativeName: "Íslenska",
	},
	"it": {
		Name:       "Italian",
		NativeName: "Italiano",
	},
	"iu": {
		Name:       "Inuktitut",
		NativeName: "ᐃᓄᒃᑎᑐᑦ",
	},
	"ja": {
		Name:       "Japanese",
		NativeName: "日本語 (にほんご／にっぽんご)",
	},
	"jv": {
		Name:       "Javanese",
		NativeName: "basa Jawa",
	},
	"kl": {
		Name:       "Kalaallisut, Greenlandic",
		NativeName: "kalaallisut, kalaallit oqaasii",
	},
	"kn": {
		Name:       "Kannada",
		NativeName: "ಕನ್ನಡ",
	},
	"kr": {
		Name:       "Kanuri",
		NativeName: "Kanuri",
	},
	"ks": {
		Name:       "Kashmiri",
		NativeName: "कश्मीरी, كشميري‎",
	},
	"kk": {
		Name:       "Kazakh",
		NativeName: "Қазақ тілі",
	},
	"km": {
		Name:       "Khmer",
		NativeName: "ភាសាខ្មែរ",
	},
	"ki": {
		Name:       "Kikuyu, Gikuyu",
		NativeName: "Gĩkũyũ",
	},
	"rw": {
		Name:       "Kinyarwanda",
		NativeName: "Ikinyarwanda",
	},
	"ky": {
		Name:       "Kirghiz, Kyrgyz",
		NativeName: "кыргыз тили",
	},
	"kv": {
		Name:       "Komi",
		NativeName: "коми кыв",
	},
	"kg": {
		Name:       "Kongo",
		NativeName: "KiKongo",
	},
	"ko": {
		Name:       "Korean",
		NativeName: "한국어 (韓國語), 조선말 (朝鮮語)",
	},
	"ku": {
		Name:       "Kurdish",
		NativeName: "Kurdî, كوردی‎",
	},
	"kj": {
		Name:       "Kwanyama, Kuanyama",
		NativeName: "Kuanyama",
	},
	"la": {
		Name:       "Latin",
		NativeName: "latine, lingua latina",
	},
	"lb": {
		Name:       "Luxembourgish, Letzeburgesch",
		NativeName: "Lëtzebuergesch",
	},
	"lg": {
		Name:       "Luganda",
		NativeName: "Luganda",
	},
	"li": {
		Name:       "Limburgish, Limburgan, Limburger",
		NativeName: "Limburgs",
	},
	"ln": {
		Name:       "Lingala",
		NativeName: "Lingála",
	},
	"lo": {
		Name:       "Lao",
		NativeName: "ພາສາລາວ",
	},
	"lt": {
		Name:       "Lithuanian",
		NativeName: "lietuvių kalba",
	},
	"lu": {
		Name:       "Luba-Katanga",
		NativeName: "",
	},
	"lv": {
		Name:       "Latvian",
		NativeName: "latviešu valoda",
	},
	"gv": {
		Name:       "Manx",
		NativeName: "Gaelg, Gailck",
	},
	"mk": {
		Name:       "Macedonian",
		NativeName: "македонски јазик",
	},
	"mg": {
		Name:       "Malagasy",
		NativeName: "Malagasy fiteny",
	},
	"ms": {
		Name:       "Malay",
		NativeName: "bahasa Melayu, بهاس ملايو‎",
	},
	"ml": {
		Name:       "Malayalam",
		NativeName: "മലയാളം",
	},
	"mt": {
		Name:       "Maltese",
		NativeName: "Malti",
	},
	"mi": {
		Name:       "Māori",
		NativeName: "te reo Māori",
	},
	"mr": {
		Name:       "Marathi (Marāṭhī)",
		NativeName: "मराठी",
	},
	"mh": {
		Name:       "Marshallese",
		NativeName: "Kajin M̧ajeļ",
	},
	"mn": {
		Name:       "Mongolian",
		NativeName: "монгол",
	},
	"na": {
		Name:       "Nauru",
		NativeName: "Ekakairũ Naoero",
	},
	"nv": {
		Name:       "Navajo, Navaho",
		NativeName: "Diné bizaad, Dinékʼehǰí",
	},
	"nb": {
		Name:       "Norwegian Bokmål",
		NativeName: "Norsk bokmål",
	},
	"nd": {
		Name:       "North Ndebele",
		NativeName: "isiNdebele",
	},
	"ne": {
		Name:       "Nepali",
		NativeName: "नेपाली",
	},
	"ng": {
		Name:       "Ndonga",
		NativeName: "Owambo",
	},
	"nn": {
		Name:       "Norwegian Nynorsk",
		NativeName: "Norsk nynorsk",
	},
	"no": {
		Name:       "Norwegian",
		NativeName: "Norsk",
	},
	"ii": {
		Name:       "Nuosu",
		NativeName: "ꆈꌠ꒿ Nuosuhxop",
	},
	"nr": {
		Name:       "South Ndebele",
		NativeName: "isiNdebele",
	},
	"oc": {
		Name:       "Occitan",
		NativeName: "Occitan",
	},
	"oj": {
		Name:       "Ojibwe, Ojibwa",
		NativeName: "ᐊᓂᔑᓈᐯᒧᐎᓐ",
	},
	"cu": {
		Name:       "Old Church Slavonic, Church Slavic, Church Slavonic, Old Bulgarian, Old Slavonic",
		NativeName: "ѩзыкъ словѣньскъ",
	},
	"om": {
		Name:       "Oromo",
		NativeName: "Afaan Oromoo",
	},
	"or": {
		Name:       "Oriya",
		NativeName: "ଓଡ଼ିଆ",
	},
	"os": {
		Name:       "Ossetian, Ossetic",
		NativeName: "ирон æвзаг",
	},
	"pa": {
		Name:       "Panjabi, Punjabi",
		NativeName: "ਪੰਜਾਬੀ, پنجابی‎",
	},
	"pi": {
		Name:       "Pāli",
		NativeName: "पाऴि",
	},
	"fa": {
		Name:       "Persian",
		NativeName: "فارسی",
	},
	"pl": {
		Name:       "Polish",
		NativeName: "polski",
	},
	"ps": {
		Name:       "Pashto, Pushto",
		NativeName: "پښتو",
	},
	"pt": {
		Name:       "Portuguese",
		NativeName: "Português",
	},
	"qu": {
		Name:       "Quechua",
		NativeName: "Runa Simi, Kichwa",
	},
	"rm": {
		Name:       "Romansh",
		NativeName: "rumantsch grischun",
	},
	"rn": {
		Name:       "Kirundi",
		NativeName: "kiRundi",
	},
	"ro": {
		Name:       "Romanian, Moldavian, Moldovan",
		NativeName: "română",
	},
	"ru": {
		Name:       "Russian",
		NativeName: "русский язык",
	},
	"sa": {
		Name:       "Sanskrit (Saṁskṛta)",
		NativeName: "संस्कृतम्",
	},
	"sc": {
		Name:       "Sardinian",
		NativeName: "sardu",
	},
	"sd": {
		Name:       "Sindhi",
		NativeName: "सिन्धी, سنڌي، سندھی‎",
	},
	"se": {
		Name:       "Northern Sami",
		NativeName: "Davvisámegiella",
	},
	"sm": {
		Name:       "Samoan",
		NativeName: "gagana faa Samoa",
	},
	"sg": {
		Name:       "Sango",
		NativeName: "yângâ tî sängö",
	},
	"sr": {
		Name:       "Serbian",
		NativeName: "српски језик",
	},
	"gd": {
		Name:       "Scottish Gaelic; Gaelic",
		NativeName: "Gàidhlig",
	},
	"sn": {
		Name:       "Shona",
		NativeName: "chiShona",
	},
	"si": {
		Name:       "Sinhala, Sinhalese",
		NativeName: "සිංහල",
	},
	"sk": {
		Name:       "Slovak",
		NativeName: "slovenčina",
	},
	"sl": {
		Name:       "Slovene",
		NativeName: "slovenščina",
	},
	"so": {
		Name:       "Somali",
		NativeName: "Soomaaliga, af Soomaali",
	},
	"st": {
		Name:       "Southern Sotho",
		NativeName: "Sesotho",
	},
	"es": {
		Name:       "Spanish; Castilian",
		NativeName: "español, castellano",
	},
	"su": {
		Name:       "Sundanese",
		NativeName: "Basa Sunda",
	},
	"sw": {
		Name:       "Swahili",
		NativeName: "Kiswahili",
	},
	"ss": {
		Name:       "Swati",
		NativeName: "SiSwati",
	},
	"sv": {
		Name:       "Swedish",
		NativeName: "svenska",
	},
	"ta": {
		Name:       "Tamil",
		NativeName: "தமிழ்",
	},
	"te": {
		Name:       "Telugu",
		NativeName: "తెలుగు",
	},
	"tg": {
		Name:       "Tajik",
		NativeName: "тоҷикӣ, toğikī, تاجیکی‎",
	},
	"th": {
		Name:       "Thai",
		NativeName: "ไทย",
	},
	"ti": {
		Name:       "Tigrinya",
		NativeName: "ትግርኛ",
	},
	"bo": {
		Name:       "Tibetan Standard, Tibetan, Central",
		NativeName: "བོད་ཡིག",
	},
	"tk": {
		Name:       "Turkmen",
		NativeName: "Türkmen, Түркмен",
	},
	"tl": {
		Name:       "Tagalog",
		NativeName: "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔",
	},
	"tn": {
		Name:       "Tswana",
		NativeName: "Setswana",
	},
	"to": {
		Name:       "Tonga (Tonga Islands)",
		NativeName: "faka Tonga",
	},
	"tr": {
		Name:       "Turkish",
		NativeName: "Türkçe",
	},
	"ts": {
		Name:       "Tsonga",
		NativeName: "Xitsonga",
	},
	"tt": {
		Name:       "Tatar",
		NativeName: "татарча, tatarça, تاتارچا‎",
	},
	"tw": {
		Name:       "Twi",
		NativeName: "Twi",
	},
	"ty": {
		Name:       "Tahitian",
		NativeName: "Reo Tahiti",
	},
	"ug": {
		Name:       "Uighur, Uyghur",
		NativeName: "Uyƣurqə, ئۇيغۇرچە‎",
	},
	"uk": {
		Name:       "Ukrainian",
		NativeName: "українська",
	},
	"ur": {
		Name:       "Urdu",
		NativeName: "اردو",
	},
	"uz": {
		Name:       "Uzbek",
		NativeName: "zbek, Ўзбек, أۇزبېك‎",
	},
	"ve": {
		Name:       "Venda",
		NativeName: "Tshivenḓa",
	},
	"vi": {
		Name:       "Vietnamese",
		NativeName: "Tiếng Việt",
	},
	"vo": {
		Name:       "Volapük",
		NativeName: "Volapük",
	},
	"wa": {
		Name:       "Walloon",
		NativeName: "Walon",
	},
	"cy": {
		Name:       "Welsh",
		NativeName: "Cymraeg",
	},
	"wo": {
		Name:       "Wolof",
		NativeName: "Wollof",
	},
	"fy": {
		Name:       "Western Frisian",
		NativeName: "Frysk",
	},
	"xh": {
		Name:       "Xhosa",
		NativeName: "isiXhosa",
	},
	"yi": {
		Name:       "Yiddish",
		NativeName: "ייִדיש",
	},
	"yo": {
		Name:       "Yoruba",
		NativeName: "Yorùbá",
	},
	"za": {
		Name:       "Zhuang, Chuang",
		NativeName: "Saɯ cueŋƅ, Saw cuengh",
	},
}
