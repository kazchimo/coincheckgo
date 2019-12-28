package coincheck

type CurrencyPair struct{ pairString string }

const withJpy = "_jpy"

var BtcJpy = CurrencyPair{"btc" + withJpy}
var EthJpy = CurrencyPair{"eth" + withJpy}
var XrpJpy = CurrencyPair{"xrp" + withJpy}
var EtcJpy = CurrencyPair{"etc" + withJpy}
var LskJpy = CurrencyPair{"lsk" + withJpy}
var FctJpy = CurrencyPair{"fct" + withJpy}
var XemJpy = CurrencyPair{"xem" + withJpy}
var LtcJpy = CurrencyPair{"ltc" + withJpy}
var BchJpy = CurrencyPair{"bch" + withJpy}
var MonaJpy = CurrencyPair{"mona" + withJpy}
