package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ltinfo/responsedto"
	"net/http"
)

type CurrencyService interface {
	GetCurrencyDetails(name string) responsedto.CurrencyDetailsDTO
	GetAllCurrency() []responsedto.CurrencyDetailsDTO
}

type CurrencyServiceImpl struct {
}

func NewCurrencyService() CurrencyService {
	return &CurrencyServiceImpl{}
}

func (cs *CurrencyServiceImpl) GetAllCurrency() []responsedto.CurrencyDetailsDTO {
	val := cs.GetAllSymbols()
	return val
}

func (cs *CurrencyServiceImpl) GetCurrencyDetails(symbol string) responsedto.CurrencyDetailsDTO {
	symbolDetails := cs.GetSymbol(symbol)
	tickerDetails := cs.GetTicker(symbol)
	currencyDetails := cs.GetCurrency(symbolDetails["baseCurrency"])

	responseDTO := responsedto.CurrencyDetailsDTO{}
	responseDTO.Id = symbolDetails["baseCurrency"]
	responseDTO.FullName = currencyDetails["fullName"]
	responseDTO.FeeCurrency = symbolDetails["feeCurrency"]
	responseDTO.Ask = tickerDetails["ask"]
	responseDTO.Bid = tickerDetails["bid"]
	responseDTO.Last = tickerDetails["last"]
	responseDTO.Open = tickerDetails["open"]
	responseDTO.Low = tickerDetails["low"]
	responseDTO.High = tickerDetails["high"]

	return responseDTO
}

func (cs *CurrencyServiceImpl) GetAllSymbols() []responsedto.CurrencyDetailsDTO {
	url := "https://api.hitbtc.com/api/2/public/symbol"
	returnVal := cs.ClientCall("GET", url)
	var res []map[string]string
	json.Unmarshal(returnVal, &res)
	var respVal []responsedto.CurrencyDetailsDTO

	for i := 0; i < 20; i++ {

		tickerDetails := cs.GetTicker(res[i]["id"])
		currencyDetails := cs.GetCurrency(res[i]["baseCurrency"])

		responseDTO := responsedto.CurrencyDetailsDTO{}
		responseDTO.Id = res[i]["baseCurrency"]
		responseDTO.FullName = currencyDetails["fullName"]
		responseDTO.FeeCurrency = res[i]["feeCurrency"]
		responseDTO.Ask = tickerDetails["ask"]
		responseDTO.Bid = tickerDetails["bid"]
		responseDTO.Last = tickerDetails["last"]
		responseDTO.Open = tickerDetails["open"]
		responseDTO.Low = tickerDetails["low"]
		responseDTO.High = tickerDetails["high"]
		respVal = append(respVal, responseDTO)
	}
	return respVal
}

func (cs *CurrencyServiceImpl) GetSymbol(symbol string) map[string]string {
	url := "https://api.hitbtc.com/api/2/public/symbol/" + symbol
	returnVal := cs.ClientCall("GET", url)
	var res map[string]string
	json.Unmarshal(returnVal, &res)
	fmt.Println("res", res)
	return res

}

func (cs *CurrencyServiceImpl) GetTicker(symbol string) map[string]string {
	url := "https://api.hitbtc.com/api/3/public/ticker/" + symbol
	returnVal := cs.ClientCall("GET", url)
	var res map[string]string
	json.Unmarshal(returnVal, &res)
	return res
}

func (cs *CurrencyServiceImpl) GetCurrency(currency string) map[string]string {
	url := "https://api.hitbtc.com/api/2/public/currency/" + currency
	val := cs.ClientCall("GET", url)
	currencyDetails := make(map[string]string)
	json.Unmarshal(val, &currencyDetails)
	return currencyDetails
}

func (cs *CurrencyServiceImpl) ClientCall(method string, url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	var response []byte

	if err != nil {
		fmt.Println(err)
		return response
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return response
	}
	return body
}
