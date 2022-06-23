package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/consumer"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/country"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"

	"github.com/stretchr/testify/assert"
)

func mockListCountriesFunc() ([]country.Country, error) {
	countries := []country.Country{
		{Code: "AX", Name: "Åland Islands"}, {Code: "AF", Name: "Afghanistan"}, {Code: "AL", Name: "Albania"}, {Code: "DZ", Name: "Algeria"}, {Code: "AS", Name: "American Samoa"}, {Code: "AD", Name: "Andorra"}, {Code: "AO", Name: "Angola"}, {Code: "AI", Name: "Anguilla"}, {Code: "AQ", Name: "Antarctica"}, {Code: "AG", Name: "Antigua & Barbuda"}, {Code: "AR", Name: "Argentina"}, {Code: "AM", Name: "Armenia"}, {Code: "AW", Name: "Aruba"}, {Code: "AU", Name: "Australia"}, {Code: "AT", Name: "Austria"}, {Code: "AZ", Name: "Azerbaijan"}, {Code: "BS", Name: "Bahamas"}, {Code: "BH", Name: "Bahrain"}, {Code: "BD", Name: "Bangladesh"}, {Code: "BB", Name: "Barbados"}, {Code: "BY", Name: "Belarus"},
		{Code: "BE", Name: "Belgium"}, {Code: "BZ", Name: "Belize"}, {Code: "BJ", Name: "Benin"}, {Code: "BM", Name: "Bermuda"}, {Code: "BT", Name: "Bhutan"}, {Code: "BO", Name: "Bolivia"}, {Code: "BQ", Name: "Bonaire, Sint Eustatius and Saba"}, {Code: "BA", Name: "Bosnia & Herzegovina"}, {Code: "BW", Name: "Botswana"}, {Code: "BV", Name: "Bouvet Island"}, {Code: "BR", Name: "Brazil"}, {Code: "IO", Name: "British Indian Ocean Territory"}, {Code: "BN", Name: "Brunei-Darussalam"}, {Code: "BG", Name: "Bulgaria"}, {Code: "BF", Name: "Burkina Faso"}, {Code: "BI", Name: "Burundi"}, {Code: "CI", Name: "Côte D'Ivoire (Ivory Coast)"}, {Code: "KH", Name: "Cambodia"}, {Code: "CM", Name: "Cameroon"}, {Code: "CA", Name: "Canada"}, {Code: "CV", Name: "Cape Verde"},
		{Code: "KY", Name: "Cayman Islands"}, {Code: "CF", Name: "Central African Republic"}, {Code: "TD", Name: "Chad"}, {Code: "CL", Name: "Chile"}, {Code: "CN", Name: "China"}, {Code: "CX", Name: "Chrismas Island"}, {Code: "CC", Name: "Cocos (Keeling) Islands"}, {Code: "MF", Name: "Collectivity of Saint Martin"}, {Code: "CO", Name: "Colombia"}, {Code: "KM", Name: "Comoros"}, {Code: "CG", Name: "Congo"}, {Code: "CD", Name: "Congo, Democratic Republic"}, {Code: "CK", Name: "Cook Islands"}, {Code: "CR", Name: "Costa Rica"}, {Code: "HR", Name: "Croatia"}, {Code: "CU", Name: "Cuba"}, {Code: "CW", Name: "Curaçao"}, {Code: "CY", Name: "Cyprus"}, {Code: "CZ", Name: "Czech Republic"}, {Code: "DK", Name: "Denmark"}, {Code: "DJ", Name: "Djibouti"},
		{Code: "DM", Name: "Dominica"}, {Code: "DO", Name: "Dominican Republic"}, {Code: "EC", Name: "Ecuador"}, {Code: "EG", Name: "Egypt"}, {Code: "SV", Name: "El Salvador"}, {Code: "GQ", Name: "Equatorial Guinea"}, {Code: "ER", Name: "Eritrea"}, {Code: "EE", Name: "Estonia"}, {Code: "ET", Name: "Ethiopia"}, {Code: "FK", Name: "Falkand Islands (Malvinas)"}, {Code: "FO", Name: "Faroe Islands"}, {Code: "FJ", Name: "Fiji"}, {Code: "FI", Name: "Finland"}, {Code: "FR", Name: "France"}, {Code: "GF", Name: "French Guiana"}, {Code: "PF", Name: "French Polynesia"}, {Code: "TF", Name: "French Southern Territories"}, {Code: "GA", Name: "Gabon"}, {Code: "GM", Name: "Gambia"}, {Code: "GE", Name: "Georgia"}, {Code: "DE", Name: "Germany"},
		{Code: "GH", Name: "Ghana"}, {Code: "GI", Name: "Gibraltar"}, {Code: "GR", Name: "Greece"}, {Code: "GL", Name: "Greenland"}, {Code: "GD", Name: "Grenada"}, {Code: "GP", Name: "Guadeloupe"}, {Code: "GU", Name: "Guam"}, {Code: "GT", Name: "Guatemala"}, {Code: "GN", Name: "Guinea"}, {Code: "GW", Name: "Guinea-Bissau"}, {Code: "GY", Name: "Guyana"}, {Code: "HT", Name: "Haiti"}, {Code: "HM", Name: "Heard Island And McDonald Islands"}, {Code: "HN", Name: "Honduras"}, {Code: "HK", Name: "Hong Kong"}, {Code: "HU", Name: "Hungary"}, {Code: "IS", Name: "Iceland"}, {Code: "IN", Name: "India"}, {Code: "ID", Name: "Indonesia"}, {Code: "IR", Name: "Iran"}, {Code: "IQ", Name: "Iraq"}, {Code: "IE", Name: "Ireland"}, {Code: "IL", Name: "Israel"},
		{Code: "IT", Name: "Italy"}, {Code: "JM", Name: "Jamaica"}, {Code: "JP", Name: "Japan"}, {Code: "JO", Name: "Jordan"}, {Code: "KZ", Name: "Kazakhstan"}, {Code: "KE", Name: "Kenya"}, {Code: "KI", Name: "Kiribati"}, {Code: "KW", Name: "Kuwait"}, {Code: "KG", Name: "Kyrgyzstan"}, {Code: "LA", Name: "Laos"}, {Code: "LV", Name: "Latvia"}, {Code: "LB", Name: "Lebanon"}, {Code: "LS", Name: "Lesotho"}, {Code: "LR", Name: "Liberia"}, {Code: "LY", Name: "Libyan Arab Jamahiriya"}, {Code: "LI", Name: "Liechtenstein"}, {Code: "LT", Name: "Lithuania"}, {Code: "LU", Name: "Luxembourg"}, {Code: "MO", Name: "Macao"}, {Code: "MK", Name: "Macedonia  (former YR)"}, {Code: "MG", Name: "Madagascar"}, {Code: "MW", Name: "Malawi"}, {Code: "MY", Name: "Malaysia"},
		{Code: "MV", Name: "Maldives (Maladive Ilands)"}, {Code: "ML", Name: "Mali"}, {Code: "MT", Name: "Malta"}, {Code: "MH", Name: "Marshall Islands"}, {Code: "MQ", Name: "Martinique"}, {Code: "MR", Name: "Mauritania"}, {Code: "MU", Name: "Mauritius"}, {Code: "YT", Name: "Mayotte"}, {Code: "MX", Name: "Mexico"}, {Code: "FM", Name: "Micronesia  (Federa States of)"}, {Code: "MD", Name: "Moldova, Republic of"}, {Code: "MC", Name: "Monaco"}, {Code: "MN", Name: "Mongolia"}, {Code: "ME", Name: "Montenegro"}, {Code: "MS", Name: "Montserrat"}, {Code: "MA", Name: "Morocco"}, {Code: "MZ", Name: "Mozambique"}, {Code: "MM", Name: "Myanmar (Burma)"}, {Code: "NA", Name: "Namibia"}, {Code: "NR", Name: "Nauru"}, {Code: "NP", Name: "Nepal"},
		{Code: "NL", Name: "Netherlands"}, {Code: "AN", Name: "Netherlands Antilles"}, {Code: "NC", Name: "New Caledonia"}, {Code: "NZ", Name: "New Zealand"}, {Code: "NI", Name: "Nicaragua"}, {Code: "NE", Name: "Niger"}, {Code: "NG", Name: "Nigeria"}, {Code: "NU", Name: "Niue"}, {Code: "NF", Name: "Norfolk Island"}, {Code: "KP", Name: "North Korea"}, {Code: "MP", Name: "Northern Mariana Islands"}, {Code: "NO", Name: "Norway"}, {Code: "OM", Name: "Oman"}, {Code: "PK", Name: "Pakistan"}, {Code: "PW", Name: "Palau"}, {Code: "PS", Name: "Palestinian Territory, Occupied"}, {Code: "PA", Name: "Panama"}, {Code: "PG", Name: "Papua-New Guinea"}, {Code: "PY", Name: "Paraguay"}, {Code: "PE", Name: "Peru"}, {Code: "PH", Name: "Philippines"},
		{Code: "PN", Name: "Pitcairn"}, {Code: "PL", Name: "Poland"}, {Code: "PT", Name: "Portugal"}, {Code: "PR", Name: "Puerto Rico"}, {Code: "QA", Name: "Quatar"}, {Code: "RE", Name: "Reunion"}, {Code: "RO", Name: "Romania"}, {Code: "RU", Name: "Russian Federation"}, {Code: "RW", Name: "Rwanda"}, {Code: "BL", Name: "Saint Barthélemy"}, {Code: "SH", Name: "Saint Helena, Ascension and Tristan da Cunha"}, {Code: "PM", Name: "Saint Pierre And Micquelon"}, {Code: "SM", Name: "San Marino"}, {Code: "ST", Name: "Sao Tome & Principe"}, {Code: "SA", Name: "Saudi Arabia"}, {Code: "SN", Name: "Senegal"}, {Code: "RS", Name: "Serbia"}, {Code: "SC", Name: "Seychelles"}, {Code: "SL", Name: "Sierra Leone"}, {Code: "SG", Name: "Singapore"}, {Code: "SX", Name: "Sint-Maarten"},
		{Code: "SK", Name: "Slovakia"}, {Code: "SI", Name: "Slovenia"}, {Code: "SB", Name: "Solomon Islands"}, {Code: "SO", Name: "Somalia"}, {Code: "ZA", Name: "South Africa"}, {Code: "GS", Name: "South Georgia & South Sandwich Islands"}, {Code: "KR", Name: "South Korea"}, {Code: "ES", Name: "Spain"}, {Code: "LK", Name: "Sri Lanka"}, {Code: "KN", Name: "St. Kitts & Nevis"}, {Code: "LC", Name: "St. Lucia"}, {Code: "VC", Name: "St. Vincent & Grenadines"}, {Code: "SD", Name: "Sudan"}, {Code: "SR", Name: "Suriname"}, {Code: "SJ", Name: "Svalbard And Jan Mayen"}, {Code: "SZ", Name: "Swaziland"}, {Code: "SE", Name: "Sweden"}, {Code: "CH", Name: "Switzerland"}, {Code: "SY", Name: "Syrian Arab Republic"}, {Code: "TW", Name: "Taiwan"}, {Code: "TJ", Name: "Tajikistan"},
		{Code: "TZ", Name: "Tanzania"}, {Code: "TH", Name: "Thailand"}, {Code: "TL", Name: "Timor-Leste"}, {Code: "TG", Name: "Togo"}, {Code: "TK", Name: "Tokelau"}, {Code: "TO", Name: "Tonga"}, {Code: "TT", Name: "Trinidad & Tobago"}, {Code: "TN", Name: "Tunisia"}, {Code: "TR", Name: "Turkey"}, {Code: "TM", Name: "Turkmenistan"}, {Code: "TC", Name: "Turks And Caicos Islands"}, {Code: "TV", Name: "Tuvalu"}, {Code: "UG", Name: "Uganda"}, {Code: "UA", Name: "Ukraine"}, {Code: "AE", Name: "United Arab Emirates"}, {Code: "GB", Name: "United Kingdom"}, {Code: "US", Name: "United States"}, {Code: "UM", Name: "United States Minor Outlying Islands"}, {Code: "UY", Name: "Uruguay"}, {Code: "UZ", Name: "Uzbekistan"}, {Code: "VU", Name: "Vanuatu"},
		{Code: "VA", Name: "Vatican City"}, {Code: "VE", Name: "Venezuela"}, {Code: "VN", Name: "Vietnam"}, {Code: "VG", Name: "Virgin Islands, British"}, {Code: "VI", Name: "Virgin Islands, U.S."}, {Code: "WF", Name: "Wallis And Futuna"}, {Code: "EH", Name: "Western Sahara"}, {Code: "WS", Name: "Western Samoa"}, {Code: "YE", Name: "Yemen"}, {Code: "ZM", Name: "Zambia"}, {Code: "ZW", Name: "Zimbabwe"},
	}

	return countries, nil
}

func mockListCountriesThrowFunc() ([]country.Country, error) {
	return nil, errs.ErrRepoMockAction
}

func TestHandleListCountries(t *testing.T) {

	type Params struct {
		Limit  *string
		Offset *string
	}

	limit := "5"
	offset := "240"

	otherLimit := "10"
	otherOffset := "5"

	testCases := []struct {
		Name                        string
		CountryConsumerListFunction func() ([]country.Country, error)
		MarshalFunction             func(v interface{}) ([]byte, error)
		WriteFunction               func(http.ResponseWriter, []byte) (int, error)
		AtoiFunction                func(s string) (int, error)
		Params                      *Params
		ExpectedCountries           []country.Country
		ExpectedStatusCode          int
	}{
		{
			Name:                        "Success handle list countries",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                atoi,
			Params:                      nil,
			ExpectedCountries: []country.Country{
				{Code: "AX", Name: "Åland Islands"}, {Code: "AF", Name: "Afghanistan"}, {Code: "AL", Name: "Albania"}, {Code: "DZ", Name: "Algeria"}, {Code: "AS", Name: "American Samoa"}, {Code: "AD", Name: "Andorra"}, {Code: "AO", Name: "Angola"}, {Code: "AI", Name: "Anguilla"}, {Code: "AQ", Name: "Antarctica"}, {Code: "AG", Name: "Antigua & Barbuda"}, {Code: "AR", Name: "Argentina"}, {Code: "AM", Name: "Armenia"}, {Code: "AW", Name: "Aruba"}, {Code: "AU", Name: "Australia"}, {Code: "AT", Name: "Austria"}, {Code: "AZ", Name: "Azerbaijan"}, {Code: "BS", Name: "Bahamas"}, {Code: "BH", Name: "Bahrain"}, {Code: "BD", Name: "Bangladesh"}, {Code: "BB", Name: "Barbados"}, {Code: "BY", Name: "Belarus"},
				{Code: "BE", Name: "Belgium"}, {Code: "BZ", Name: "Belize"}, {Code: "BJ", Name: "Benin"}, {Code: "BM", Name: "Bermuda"}, {Code: "BT", Name: "Bhutan"}, {Code: "BO", Name: "Bolivia"}, {Code: "BQ", Name: "Bonaire, Sint Eustatius and Saba"}, {Code: "BA", Name: "Bosnia & Herzegovina"}, {Code: "BW", Name: "Botswana"}, {Code: "BV", Name: "Bouvet Island"}, {Code: "BR", Name: "Brazil"}, {Code: "IO", Name: "British Indian Ocean Territory"}, {Code: "BN", Name: "Brunei-Darussalam"}, {Code: "BG", Name: "Bulgaria"}, {Code: "BF", Name: "Burkina Faso"}, {Code: "BI", Name: "Burundi"}, {Code: "CI", Name: "Côte D'Ivoire (Ivory Coast)"}, {Code: "KH", Name: "Cambodia"}, {Code: "CM", Name: "Cameroon"}, {Code: "CA", Name: "Canada"}, {Code: "CV", Name: "Cape Verde"},
				{Code: "KY", Name: "Cayman Islands"}, {Code: "CF", Name: "Central African Republic"}, {Code: "TD", Name: "Chad"}, {Code: "CL", Name: "Chile"}, {Code: "CN", Name: "China"}, {Code: "CX", Name: "Chrismas Island"}, {Code: "CC", Name: "Cocos (Keeling) Islands"}, {Code: "MF", Name: "Collectivity of Saint Martin"}, {Code: "CO", Name: "Colombia"}, {Code: "KM", Name: "Comoros"}, {Code: "CG", Name: "Congo"}, {Code: "CD", Name: "Congo, Democratic Republic"}, {Code: "CK", Name: "Cook Islands"}, {Code: "CR", Name: "Costa Rica"}, {Code: "HR", Name: "Croatia"}, {Code: "CU", Name: "Cuba"}, {Code: "CW", Name: "Curaçao"}, {Code: "CY", Name: "Cyprus"}, {Code: "CZ", Name: "Czech Republic"}, {Code: "DK", Name: "Denmark"}, {Code: "DJ", Name: "Djibouti"},
				{Code: "DM", Name: "Dominica"}, {Code: "DO", Name: "Dominican Republic"}, {Code: "EC", Name: "Ecuador"}, {Code: "EG", Name: "Egypt"}, {Code: "SV", Name: "El Salvador"}, {Code: "GQ", Name: "Equatorial Guinea"}, {Code: "ER", Name: "Eritrea"}, {Code: "EE", Name: "Estonia"}, {Code: "ET", Name: "Ethiopia"}, {Code: "FK", Name: "Falkand Islands (Malvinas)"}, {Code: "FO", Name: "Faroe Islands"}, {Code: "FJ", Name: "Fiji"}, {Code: "FI", Name: "Finland"}, {Code: "FR", Name: "France"}, {Code: "GF", Name: "French Guiana"}, {Code: "PF", Name: "French Polynesia"}, {Code: "TF", Name: "French Southern Territories"}, {Code: "GA", Name: "Gabon"}, {Code: "GM", Name: "Gambia"}, {Code: "GE", Name: "Georgia"}, {Code: "DE", Name: "Germany"},
				{Code: "GH", Name: "Ghana"}, {Code: "GI", Name: "Gibraltar"}, {Code: "GR", Name: "Greece"}, {Code: "GL", Name: "Greenland"}, {Code: "GD", Name: "Grenada"}, {Code: "GP", Name: "Guadeloupe"}, {Code: "GU", Name: "Guam"}, {Code: "GT", Name: "Guatemala"}, {Code: "GN", Name: "Guinea"}, {Code: "GW", Name: "Guinea-Bissau"}, {Code: "GY", Name: "Guyana"}, {Code: "HT", Name: "Haiti"}, {Code: "HM", Name: "Heard Island And McDonald Islands"}, {Code: "HN", Name: "Honduras"}, {Code: "HK", Name: "Hong Kong"}, {Code: "HU", Name: "Hungary"}, {Code: "IS", Name: "Iceland"}, {Code: "IN", Name: "India"}, {Code: "ID", Name: "Indonesia"}, {Code: "IR", Name: "Iran"}, {Code: "IQ", Name: "Iraq"}, {Code: "IE", Name: "Ireland"}, {Code: "IL", Name: "Israel"},
				{Code: "IT", Name: "Italy"}, {Code: "JM", Name: "Jamaica"}, {Code: "JP", Name: "Japan"}, {Code: "JO", Name: "Jordan"}, {Code: "KZ", Name: "Kazakhstan"}, {Code: "KE", Name: "Kenya"}, {Code: "KI", Name: "Kiribati"}, {Code: "KW", Name: "Kuwait"}, {Code: "KG", Name: "Kyrgyzstan"}, {Code: "LA", Name: "Laos"}, {Code: "LV", Name: "Latvia"}, {Code: "LB", Name: "Lebanon"}, {Code: "LS", Name: "Lesotho"}, {Code: "LR", Name: "Liberia"}, {Code: "LY", Name: "Libyan Arab Jamahiriya"}, {Code: "LI", Name: "Liechtenstein"}, {Code: "LT", Name: "Lithuania"}, {Code: "LU", Name: "Luxembourg"}, {Code: "MO", Name: "Macao"}, {Code: "MK", Name: "Macedonia  (former YR)"}, {Code: "MG", Name: "Madagascar"}, {Code: "MW", Name: "Malawi"}, {Code: "MY", Name: "Malaysia"},
				{Code: "MV", Name: "Maldives (Maladive Ilands)"}, {Code: "ML", Name: "Mali"}, {Code: "MT", Name: "Malta"}, {Code: "MH", Name: "Marshall Islands"}, {Code: "MQ", Name: "Martinique"}, {Code: "MR", Name: "Mauritania"}, {Code: "MU", Name: "Mauritius"}, {Code: "YT", Name: "Mayotte"}, {Code: "MX", Name: "Mexico"}, {Code: "FM", Name: "Micronesia  (Federa States of)"}, {Code: "MD", Name: "Moldova, Republic of"}, {Code: "MC", Name: "Monaco"}, {Code: "MN", Name: "Mongolia"}, {Code: "ME", Name: "Montenegro"}, {Code: "MS", Name: "Montserrat"}, {Code: "MA", Name: "Morocco"}, {Code: "MZ", Name: "Mozambique"}, {Code: "MM", Name: "Myanmar (Burma)"}, {Code: "NA", Name: "Namibia"}, {Code: "NR", Name: "Nauru"}, {Code: "NP", Name: "Nepal"},
				{Code: "NL", Name: "Netherlands"}, {Code: "AN", Name: "Netherlands Antilles"}, {Code: "NC", Name: "New Caledonia"}, {Code: "NZ", Name: "New Zealand"}, {Code: "NI", Name: "Nicaragua"}, {Code: "NE", Name: "Niger"}, {Code: "NG", Name: "Nigeria"}, {Code: "NU", Name: "Niue"}, {Code: "NF", Name: "Norfolk Island"}, {Code: "KP", Name: "North Korea"}, {Code: "MP", Name: "Northern Mariana Islands"}, {Code: "NO", Name: "Norway"}, {Code: "OM", Name: "Oman"}, {Code: "PK", Name: "Pakistan"}, {Code: "PW", Name: "Palau"}, {Code: "PS", Name: "Palestinian Territory, Occupied"}, {Code: "PA", Name: "Panama"}, {Code: "PG", Name: "Papua-New Guinea"}, {Code: "PY", Name: "Paraguay"}, {Code: "PE", Name: "Peru"}, {Code: "PH", Name: "Philippines"},
				{Code: "PN", Name: "Pitcairn"}, {Code: "PL", Name: "Poland"}, {Code: "PT", Name: "Portugal"}, {Code: "PR", Name: "Puerto Rico"}, {Code: "QA", Name: "Quatar"}, {Code: "RE", Name: "Reunion"}, {Code: "RO", Name: "Romania"}, {Code: "RU", Name: "Russian Federation"}, {Code: "RW", Name: "Rwanda"}, {Code: "BL", Name: "Saint Barthélemy"}, {Code: "SH", Name: "Saint Helena, Ascension and Tristan da Cunha"}, {Code: "PM", Name: "Saint Pierre And Micquelon"}, {Code: "SM", Name: "San Marino"}, {Code: "ST", Name: "Sao Tome & Principe"}, {Code: "SA", Name: "Saudi Arabia"}, {Code: "SN", Name: "Senegal"}, {Code: "RS", Name: "Serbia"}, {Code: "SC", Name: "Seychelles"}, {Code: "SL", Name: "Sierra Leone"}, {Code: "SG", Name: "Singapore"}, {Code: "SX", Name: "Sint-Maarten"},
				{Code: "SK", Name: "Slovakia"}, {Code: "SI", Name: "Slovenia"}, {Code: "SB", Name: "Solomon Islands"}, {Code: "SO", Name: "Somalia"}, {Code: "ZA", Name: "South Africa"}, {Code: "GS", Name: "South Georgia & South Sandwich Islands"}, {Code: "KR", Name: "South Korea"}, {Code: "ES", Name: "Spain"}, {Code: "LK", Name: "Sri Lanka"}, {Code: "KN", Name: "St. Kitts & Nevis"}, {Code: "LC", Name: "St. Lucia"}, {Code: "VC", Name: "St. Vincent & Grenadines"}, {Code: "SD", Name: "Sudan"}, {Code: "SR", Name: "Suriname"}, {Code: "SJ", Name: "Svalbard And Jan Mayen"}, {Code: "SZ", Name: "Swaziland"}, {Code: "SE", Name: "Sweden"}, {Code: "CH", Name: "Switzerland"}, {Code: "SY", Name: "Syrian Arab Republic"}, {Code: "TW", Name: "Taiwan"}, {Code: "TJ", Name: "Tajikistan"},
				{Code: "TZ", Name: "Tanzania"}, {Code: "TH", Name: "Thailand"}, {Code: "TL", Name: "Timor-Leste"}, {Code: "TG", Name: "Togo"}, {Code: "TK", Name: "Tokelau"}, {Code: "TO", Name: "Tonga"}, {Code: "TT", Name: "Trinidad & Tobago"}, {Code: "TN", Name: "Tunisia"}, {Code: "TR", Name: "Turkey"}, {Code: "TM", Name: "Turkmenistan"}, {Code: "TC", Name: "Turks And Caicos Islands"}, {Code: "TV", Name: "Tuvalu"}, {Code: "UG", Name: "Uganda"}, {Code: "UA", Name: "Ukraine"}, {Code: "AE", Name: "United Arab Emirates"}, {Code: "GB", Name: "United Kingdom"}, {Code: "US", Name: "United States"}, {Code: "UM", Name: "United States Minor Outlying Islands"}, {Code: "UY", Name: "Uruguay"}, {Code: "UZ", Name: "Uzbekistan"}, {Code: "VU", Name: "Vanuatu"},
				{Code: "VA", Name: "Vatican City"}, {Code: "VE", Name: "Venezuela"}, {Code: "VN", Name: "Vietnam"}, {Code: "VG", Name: "Virgin Islands, British"}, {Code: "VI", Name: "Virgin Islands, U.S."}, {Code: "WF", Name: "Wallis And Futuna"}, {Code: "EH", Name: "Western Sahara"}, {Code: "WS", Name: "Western Samoa"}, {Code: "YE", Name: "Yemen"}, {Code: "ZM", Name: "Zambia"}, {Code: "ZW", Name: "Zimbabwe"},
			},
			ExpectedStatusCode: 200,
		}, {
			Name:                        "Throwing error list continents function",
			CountryConsumerListFunction: mockListCountriesThrowFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                atoi,
			Params:                      nil,
			ExpectedStatusCode:          500,
		}, {
			Name:                        "Throwing error marshal function",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             fakeMarshal,
			WriteFunction:               write,
			AtoiFunction:                atoi,
			Params:                      nil,
			ExpectedStatusCode:          500,
		}, {
			Name:                        "Throwing error on write function",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               fakeWrite,
			AtoiFunction:                atoi,
			Params:                      nil,
			ExpectedStatusCode:          500,
		}, {
			Name:                        "Success handle list countries with limit",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                atoi,
			Params: &Params{
				Limit:  &limit,
				Offset: nil,
			},
			ExpectedCountries: []country.Country{
				{Code: "AX", Name: "Åland Islands"}, {Code: "AF", Name: "Afghanistan"},
				{Code: "AL", Name: "Albania"}, {Code: "DZ", Name: "Algeria"},
				{Code: "AS", Name: "American Samoa"},
			},
			ExpectedStatusCode: 200,
		}, {
			Name:                        "Throwing error on atoi function using limit",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                fakeAtoi,
			Params: &Params{
				Limit:  &limit,
				Offset: nil,
			},
			ExpectedCountries:  nil,
			ExpectedStatusCode: 500,
		}, {
			Name:                        "Success handle list countries with offset",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                atoi,
			Params: &Params{
				Limit:  nil,
				Offset: &offset,
			},
			ExpectedCountries: []country.Country{
				{Code: "WF", Name: "Wallis And Futuna"}, {Code: "EH", Name: "Western Sahara"},
				{Code: "WS", Name: "Western Samoa"}, {Code: "YE", Name: "Yemen"},
				{Code: "ZM", Name: "Zambia"}, {Code: "ZW", Name: "Zimbabwe"},
			},
			ExpectedStatusCode: 200,
		}, {
			Name:                        "Throwing error on atoi function using offset",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                fakeAtoi,
			Params: &Params{
				Limit:  nil,
				Offset: &offset,
			},
			ExpectedCountries:  nil,
			ExpectedStatusCode: 500,
		}, {
			Name:                        "Success handle list countries with limit and offset",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                atoi,
			Params: &Params{
				Limit:  &otherLimit,
				Offset: &otherOffset,
			},
			ExpectedCountries: []country.Country{
				{Code: "AD", Name: "Andorra"}, {Code: "AO", Name: "Angola"},
				{Code: "AI", Name: "Anguilla"}, {Code: "AQ", Name: "Antarctica"},
				{Code: "AG", Name: "Antigua & Barbuda"}, {Code: "AR", Name: "Argentina"},
				{Code: "AM", Name: "Armenia"}, {Code: "AW", Name: "Aruba"},
				{Code: "AU", Name: "Australia"}, {Code: "AT", Name: "Austria"},
			},
			ExpectedStatusCode: 200,
		}, {
			Name:                        "Throwing error on atoi function using limit and offset",
			CountryConsumerListFunction: mockListCountriesFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			AtoiFunction:                fakeAtoi,
			Params: &Params{
				Limit:  &otherLimit,
				Offset: &otherOffset,
			},
			ExpectedCountries:  nil,
			ExpectedStatusCode: 500,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		consumer.SetCountryConsumer(consumer.MockCountryConsumer{
			ListFunc: tc.CountryConsumerListFunction,
		})
		defer consumer.SetCountryConsumer(nil)

		jsonMarshal = tc.MarshalFunction
		defer restoreMarshal(jsonMarshal)

		write = tc.WriteFunction
		defer restoreWrite(write)

		atoi = tc.AtoiFunction
		defer restoreAtoi(atoi)

		req, err := http.NewRequest(http.MethodGet, "/countries", nil)

		if tc.Params != nil {
			values := req.URL.Query()

			if tc.Params.Limit != nil {
				values.Add("limit", *tc.Params.Limit)
			}

			if tc.Params.Offset != nil {
				values.Add("offset", *tc.Params.Offset)
			}

			req.URL.RawQuery = values.Encode()
		}

		assert.NoError(t, err)
		res := httptest.NewRecorder()

		HandleListCountries(res, req)

		assert.Equal(t, tc.ExpectedStatusCode, res.Code)

		if res.Code == 200 {
			countries := []country.Country{}
			err = json.Unmarshal(res.Body.Bytes(), &countries)
			assert.NoError(t, err)

			assert.Equal(t, tc.ExpectedCountries, countries)
		}
	}
}
