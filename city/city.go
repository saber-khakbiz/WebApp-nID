package city

var codecity = map[string]string{
	"412": "بروجرد",
	"413": "بروجرد",
	"406": "خرم آباد",
	"407": "خرم آباد",
	"421": "دورود",
	"498": "بابلسر",
	"483": "چالوس",
	"227": "رامسر",
	"627": "کلاردشت",
	"262": "تالش",
}

func Detectcity(nID string) string {
	PnID := nID[:3]
	if _, ok := codecity[PnID]; ok {
		return codecity[PnID]
	}
	return "undefined"
}
