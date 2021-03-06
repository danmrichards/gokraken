//go:generate stringer -type Currency
package asset

// Code generated by asset
// DO NOT EDIT

const (
	KFEE Currency = iota
	XDAO
	XNMC
	XXDG
	DASH
	EOS
	XXBT
	XXLM
	XREP
	BCH
	ZJPY
	XXMR
	ZCAD
	ZGBP
	XETH
	XMLN
	GNO
	XICN
	XLTC
	ZUSD
	USDT
	XETC
	XXRP
	XXVN
	XZEC
	ZEUR
	ZKRW
)

type Currency int

var validCurrencies = []Currency{
	KFEE,
	XDAO,
	XNMC,
	XXDG,
	DASH,
	EOS,
	XXBT,
	XXLM,
	XREP,
	BCH,
	ZJPY,
	XXMR,
	ZCAD,
	ZGBP,
	XETH,
	XMLN,
	GNO,
	XICN,
	XLTC,
	ZUSD,
	USDT,
	XETC,
	XXRP,
	XXVN,
	XZEC,
	ZEUR,
	ZKRW,
}
