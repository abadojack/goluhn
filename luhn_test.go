package goluhn

import "testing"

func Test_IsValid(t *testing.T) {
	tests := map[string]bool{
		"378282246310005":  true,
		"371449635398431":  true,
		"5610591081018250": true,
		"30569309025904":   true,
		"38520000023237":   true,
		"6011111111111117": true,
		"6011000990139424": true,
		"3530111333300000": true,
		"3566002020360505": true,
		"5555555555554444": true,
		"5105105105105100": true,
		"4111111111111111": true,
		"4012888888881881": true,
		"4222222222222":    true,
		"76009244561":      true,
		"5019717010103742": true,
		"6331101999990016": true,
	}

	for key, want := range tests {
		got := IsValid(key)
		if got != want {
			t.Fatalf("%s, want:%t, got:%t\n", key, want, got)
		}
	}
}

func Test_CardType(t *testing.T) {
	tests := map[string]string{
		"8800000000000000": "unionpay",

		"5019000000000000": "dankort",

		"4026000000000000": "electron",
		"4175000000000000": "electron",
		"4405000000000000": "electron",
		"4508000000000000": "electron",
		"4844000000000000": "electron",
		"4913000000000000": "electron",
		"4917000000000000": "electron",

		"5018000000000000": "maestro",
		"5020000000000000": "maestro",
		"5038000000000000": "maestro",
		"5612000000000000": "maestro",
		"5893000000000000": "maestro",
		"6304000000000000": "maestro",
		"6759000000000000": "maestro",
		"6761000000000000": "maestro",
		"6762000000000000": "maestro",
		"6763000000000000": "maestro",
		"0604000000000000": "maestro",
		"6390000000000000": "maestro",

		"3528000000000000": "jcb",
		"3589000000000000": "jcb",
		"3529000000000000": "jcb",

		"6360000000000000": "interpayment",

		"4916338506082832": "visa",
		"4556015886206505": "visa",
		"4539048040151731": "visa",
		"4024007198964305": "visa",
		"4716175187624512": "visa",

		"5280934283171080": "mastercard",
		"5456060454627409": "mastercard",
		"5331113404316994": "mastercard",
		"5259474113320034": "mastercard",
		"5442179619690834": "mastercard",

		"6011894492395579": "discover",
		"6011388644154687": "discover",
		"6011880085013612": "discover",
		"6011652795433988": "discover",
		"6011375973328347": "discover",

		"345936346788903": "amex",
		"377669501013152": "amex",
		"373083634595479": "amex",
		"370710819865268": "amex",
		"371095063560404": "amex",
	}
	for key, want := range tests {
		got := CardType(key)
		if got != want {
			t.Fatal(key + " want: " + want + ", got: " + got)
		}
	}

}
