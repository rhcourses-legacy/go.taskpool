package addelements

func AddElements(list []int) int { // Zu viele Klammern
	result := 0
	for _, el := range list { // Leerzeichen vor _ hat gefehlt.
		result += el // Ergebnis der Zuweisung wurde nicht benutzt.
	}
	return result
}
