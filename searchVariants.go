package gosearcher

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"ac":  "https://www.google.ac/search?q=",
	"ad":  "https://www.google.ad/search?q=",
	"ae":  "https://www.google.ae/search?q=",
	"af":  "https://www.google.af/search?q=",
}

var yandexDomains = map[string]string{
	"com": "https://yandex.com/search/?text=",
	"ru":  "https://yandex.ru/search/?text=",
	"ua":  "https://yandex.ua/search/?text=",
	"kz":  "https://yandex.kz/search/?text=",
	"by":  "https://yandex.by/search/?text=",
	"tr":  "https://yandex.tr/search/?text=",
}
