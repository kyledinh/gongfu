package util

import (
	"testing"
)

func Test_SatoshiConversion(t *testing.T) {
	var btc float64
	var sat int64
	var err error

	if sat, err = Btc_Sat(1.0); sat != 100000000 {
		t.Error("FAILED to SatoshiConversionTest 1.0 BTC : %v", sat)
	}
	if sat, err = Btc_Sat(-1.0); sat != -100000000 {
		t.Error("FAILED to SatoshiConversionTest -1.0 BTC : %v", sat)
	}
	if sat, err = Btc_Sat(0.00000001); sat != 1 {
		t.Error("FAILED to SatoshiConversionTest 1 Satoshi : %v", sat)
	}
	if sat, err = Btc_Sat(0.0000001); sat != 10 {
		t.Error("FAILED to SatoshiConversionTest 10 Satoshi : %v", sat)
	}
	if sat, err = Btc_Sat(0.317383); sat != 31738300 {
		t.Error("FAILED to SatoshiConversionTest Satoshi : %v", sat)
	}
	if sat, err = Btc_Sat(1234567890.0); sat != 123456789000000000 {
		t.Error("FAILED to SatoshiConversionTest Satoshi : %v", sat)
	}

	if btc, err = Sat_Btc(1); btc != 0.00000001 {
		t.Error("FAILED to SatoshiConversionTest 1 Satoshi : %v", btc)
	}
	if btc, err = Sat_Btc(-1); btc != -0.00000001 {
		t.Error("FAILED to SatoshiConversionTest -1 Satoshi : %v", btc)
	}
	if btc, err = Sat_Btc(1234567890); btc != 12.3456789 {
		t.Error("FAILED to SatoshiConversionTest large Satoshi : %v", btc)
	}

	if err != nil {
		t.Error("FAILED to SatoshiConversionTest: %s", err)
	}

}
