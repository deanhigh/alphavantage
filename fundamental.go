package alphavantage

import (
	"strconv"
	"strings"
	"time"
)

// OptionalDate is the wrapper date for alphavantage dates in the api
type OptionalDate time.Time

// OptionalFloat64 is used to handle the fact that AV API can return None or a float64.
type OptionalFloat64 float64

// UnmarshalJSON Unmarshal optional types from the AV api which can include None or a float64 number.
func (j *OptionalFloat64) UnmarshalJSON(b []byte) error {
	// If None, return nil because AV api allows strings of None to be returned
	s := strings.Trim(string(b), "\"")
	if s == "None" {
		// Don't do anything and return
		return nil
	}
	// If not None, then try parse it.
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*j = OptionalFloat64(val)
	return nil
}

// UnmarshalJSON Unmarsal optional date from AV
func (j *OptionalDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "None" {
		// Don't do anything and return
		return nil
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = OptionalDate(t)
	return nil
}

// MarshalJSON marshal a jsondate
func (j OptionalDate) MarshalJSON() ([]byte, error) {
	t := time.Time(j)
	return t.MarshalJSON()
}

// GoString for printing your date to %v
func (j OptionalDate) GoString() string {
	t := time.Time(j)
	return t.Format("2006-01-02")
}

// FundamentalService wraps client for the service
type FundamentalService struct {
	client *Client
}

// CompanyOverview Struct unpacked from the alphavantage json
type CompanyOverview struct {
	Symbol                     string           `json:"Symbol"`
	AssetType                  string           `json:"AssetType"`
	Name                       string           `json:"Name"`
	Description                string           `json:"Description"`
	Exchange                   string           `json:"Exchange"`
	Currency                   string           `json:"Currency"`
	Country                    string           `json:"Country"`
	Sector                     string           `json:"Sector"`
	Industry                   string           `json:"Industry"`
	Address                    string           `json:"Address"`
	FullTimeEmployees          int              `json:"FullTimeEmployees,string"`
	FiscalYearEnd              string           `json:"FiscalYearEnd"`
	LatestQuarter              string           `json:"LatestQuarter"`
	MarketCapitalization       float64          `json:"MarketCapitalization,string"`
	EBITDA                     float64          `json:"EBITDA,string"`
	PERatio                    *OptionalFloat64 `json:"PERatio,string"`
	PEGRatio                   float64          `json:"PEGRatio,string"`
	BookValue                  float64          `json:"BookValue,string"`
	DividendPerShare           *OptionalFloat64 `json:"DividendPerShare,string"`
	DividendYield              float64          `json:"DividendYield,string"`
	EPS                        float64          `json:"EPS,string"`
	RevenuePerShareTTM         float64          `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64          `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64          `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64          `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64          `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 float64          `json:"RevenueTTM,string"`
	GrossProfitTTM             float64          `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64          `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64          `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64          `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float64          `json:"AnalystTargetPrice,string"`
	TrailingPE                 float64          `json:"TrailingPE,string"`
	ForwardPE                  float64          `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64          `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64          `json:"PriceToBookRatio,string"`
	EVToRevenue                float64          `json:"EVToRevenue,string"`
	EVToEBITDA                 float64          `json:"EVToEBITDA,string"`
	Beta                       float64          `json:"Beta,string"`
	High52Week                 float64          `json:"52WeekHigh,string"`
	Low52Week                  float64          `json:"52WeekLow,string"`
	MovingAverage50Day         float64          `json:"50DayMovingAverage,string"`
	MovingAverage200Day        float64          `json:"200DayMovingAverage,string"`
	SharesOutstanding          float64          `json:"SharesOutstanding,string"`
	SharesFloat                float64          `json:"SharesFloat,string"`
	SharesShort                float64          `json:"SharesShort,string"`
	SharesShortPriorMonth      float64          `json:"SharesShortPriorMonth,string"`
	ShortRatio                 float64          `json:"ShortRatio,string"`
	ShortPercentOutstanding    float64          `json:"ShortPercentOutstanding,string"`
	ShortPercentFloat          float64          `json:"ShortPercentFloat,string"`
	PercentInsiders            float64          `json:"PercentInsiders,string"`
	PercentInstitutions        float64          `json:"PercentInstitutions,string"`
	ForwardAnnualDividendRate  float64          `json:"ForwardAnnualDividendRate,string"`
	ForwardAnnualDividendYield float64          `json:"ForwardAnnualDividendYield,string"`
	PayoutRatio                float64          `json:"PayoutRatio,string"`
	DividendDate               *OptionalDate    `json:"DividendDate"`
	ExDividendDate             *OptionalDate    `json:"ExDividendDate"`
	LastSplitFactor            string           `json:"LastSplitFactor"`
	LastSplitDate              *OptionalDate    `json:"LastSplitDate"`
}

// NewFundamentalService creates a new fundamental service
func NewFundamentalService(client *Client) {
	c := FundamentalService{}
	c.client = client
}

// GetCompanyOverview Get company overview data for symbol
// https://www.alphavantage.co/documentation/#company-overview
func (fs *FundamentalService) GetCompanyOverview(symbol string) (*CompanyOverview, error) {
	co := CompanyOverview{}
	params := map[string]string{"function": "OVERVIEW", "symbol": symbol}

	query, err := fs.client.NewQuery(params)
	if err != nil {
		return nil, err
	}

	_, err = fs.client.Do(query, &co)

	if err != nil {
		return nil, err
	}

	return &co, err
}
