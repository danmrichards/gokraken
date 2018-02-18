package asset

func Valid(currency string) bool {
	for _, v := range validCurrencies {
		if v.String() == currency {
			return true
		}
	}
	return false
}

func Find(currency string) *Currency {
	for _, v := range validCurrencies {
		if v.String() == currency {
			return &v
		}
	}
	return nil
}
