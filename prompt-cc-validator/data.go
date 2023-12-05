package main

type vendor struct {
	ccType  string
	ccRegex string
}

func VendorRegex() {
	v1 := vendor{
		ccType:  "Amex Card",
		ccRegex: "^3[47][0-9]{13}$",
	}
	v2 := vendor{
		ccType:  "BCGlobal",
		ccRegex: "^(6541|6556)[0-9]{12}$",
	}
	v3 := vendor{
		ccType:  "Carte Blanche Card",
		ccRegex: "^389[0-9]{11}$",
	}
	v4 := vendor{
		ccType:  "Diners Club Card",
		ccRegex: "^3(?:0[0-5]|[68][0-9])[0-9]{11}$",
	}
	v5 := vendor{
		ccType:  "Discover Card",
		ccRegex: "^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$",
	}
	v6 := vendor{
		ccType:  "Insta Payment Card",
		ccRegex: "^63[7-9][0-9]{13}$",
	}
	v7 := vendor{
		ccType:  "KoreanLocalCard",
		ccRegex: "^9[0-9]{15}$",
	}
	v8 := vendor{
		ccType:  "Laser Card",
		ccRegex: "^(6304|6706|6709|6771)[0-9]{12,15}$",
	}
	v9 := vendor{
		ccType:  "Maestro Card",
		ccRegex: "^(5018|5020|5038|6304|6759|6761|6763)[0-9]{8,15}$",
	}
	v10 := vendor{
		ccType:  "Mastercard",
		ccRegex: "^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$",
	}
	v11 := vendor{
		ccType:  "Solo Card",
		ccRegex: "^(6334|6767)[0-9]{12}|(6334|6767)[0-9]{14}|(6334|6767)[0-9]{15}$",
	}
	v12 := vendor{
		ccType:  "Switch Card",
		ccRegex: "^(4903|4905|4911|4936|6333|6759)[0-9]{12}|(4903|4905|4911|4936|6333|6759)[0-9]{14}|(4903|4905|4911|4936|6333|6759)[0-9]{15}|564182[0-9]{10}|564182[0-9]{12}|564182[0-9]{13}|633110[0-9]{10}|633110[0-9]{12}|633110[0-9]{13}$",
	}
	v13 := vendor{
		ccType:  "Union Pay Card",
		ccRegex: "^(62[0-9]{14,17})$",
	}
	v14 := vendor{
		ccType:  "Visa Card",
		ccRegex: "^4[0-9]{12}(?:[0-9]{3})?$",
	}

	// Adds data structures to the map
	var vendorInfo = map[vendor]int{v1: 1, v2: 2, v3: 3, v4: 4, v5: 5, v6: 6, v7: 7, v8: 8, v9: 9, v10: 10, v11: 11, v12: 12, v13: 13, v14: 14}

	CCTest(vendorInfo)
}
