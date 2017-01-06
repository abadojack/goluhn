package goluhn

import "testing"

func Test_IsValid(t *testing.T) {
	tests := map[string]bool{
		//AMEXCO
		"378282246310005": true,
		"371449635398431": true,
		//MASTERCARD
		"5555555555554444": true,
		"5105105105105100": true,
		//VISA
		"4111111111111111": true,
		"4012888888881881": true,
		//DINERS CLUB
		"38520000023237": true,
		"30569309025904": true,
		//DISCOVER
		"6011111111111117": true,
		"6011000990139424": true,
		//JCB
		"3530111333300000": true,
		"3566002020360505": true,
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
		//unionpay
		//dankort
		"5019717010103742": "dankort",
		//electron
		//"4245190000000311":"electron",
		//maestro
		"6759649826438453": "maestro",
		//"6799990100000000019":"maestro",
		//jcb
		"3530111333300000": "jcb",
		"3566002020360505": "jcb",
		"3543836016493632": "jcb",
		"3571263025835254": "jcb",
		"3557355638346998": "jcb",

		//interpayment
		//visa
		"4111111111111111": "visa",
		"4012888888881881": "visa",
		"4024007193707949": "visa",
		"4929817677470574": "visa",
		"4024007115673344": "visa",
		"4716532348013938": "visa",
		"4539384536961089": "visa",
		"4024007146787212": "visa",
		"4525688362572927": "visa",
		"4485402097325322": "visa",
		"4929408497555453": "visa",
		"4532286166828495": "visa",
		"4916995775692":    "visa",
		"4716923038761":    "visa",
		"4716238107442":    "visa",
		"4916022529070":    "visa",
		"4532515324620":    "visa",

		//mastercard
		"5555555555554444": "mastercard",
		"5105105105105100": "mastercard",
		"5160824924259722": "mastercard",
		"5429031374131834": "mastercard",
		"5223067257706407": "mastercard",
		"5441992156326104": "mastercard",
		"5284944423251672": "mastercard",
		"5375844422632254": "mastercard",
		"5403679621191959": "mastercard",
		"5411491612412158": "mastercard",
		"5501365628088221": "mastercard",
		"5574175648758857": "mastercard",
		//discover
		"6011111111111117": "discover",
		"6011000990139424": "discover",
		"6011842766013857": "discover",
		"6011936194350590": "discover",
		"6011935193692788": "discover",
		//amex
		"378282246310005": "amex",
		"371449635398431": "amex",
		"349228845945887": "amex",
		"370794788668820": "amex",
		"376485316211225": "amex",
		"346068046967428": "amex",
		"372725584902193": "amex",
	}
	for key, want := range tests {
		got := CardType(key)
		if got != want {
			t.Fatal(key + " want: " + want + ", got: " + got)
		}
	}

}
