package weightconv

func PToK(p Pounds) Kilograms	{ return Kilograms(p * 0.4535924) }
func KToP(k Kilograms) Pounds	{ return Pounds(k * 2.204623) }
