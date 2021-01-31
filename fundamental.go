package alphavantage

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
	LatestQuarter              *OptionalDate    `json:"LatestQuarter,string"`
	MarketCapitalization       *OptionalFloat64 `json:"MarketCapitalization,string"`
	EBITDA                     *OptionalFloat64 `json:"EBITDA,string"`
	PERatio                    *OptionalFloat64 `json:"PERatio,string"`
	PEGRatio                   *OptionalFloat64 `json:"PEGRatio,string"`
	BookValue                  *OptionalFloat64 `json:"BookValue,string"`
	DividendPerShare           *OptionalFloat64 `json:"DividendPerShare,string"`
	DividendYield              *OptionalFloat64 `json:"DividendYield,string"`
	EPS                        *OptionalFloat64 `json:"EPS,string"`
	RevenuePerShareTTM         *OptionalFloat64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               *OptionalFloat64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         *OptionalFloat64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          *OptionalFloat64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          *OptionalFloat64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 *OptionalFloat64 `json:"RevenueTTM,string"`
	GrossProfitTTM             *OptionalFloat64 `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              *OptionalFloat64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY *OptionalFloat64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  *OptionalFloat64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         *OptionalFloat64 `json:"AnalystTargetPrice,string"`
	TrailingPE                 *OptionalFloat64 `json:"TrailingPE,string"`
	ForwardPE                  *OptionalFloat64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       *OptionalFloat64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           *OptionalFloat64 `json:"PriceToBookRatio,string"`
	EVToRevenue                *OptionalFloat64 `json:"EVToRevenue,string"`
	EVToEBITDA                 *OptionalFloat64 `json:"EVToEBITDA,string"`
	Beta                       *OptionalFloat64 `json:"Beta,string"`
	High52Week                 *OptionalFloat64 `json:"52WeekHigh,string"`
	Low52Week                  *OptionalFloat64 `json:"52WeekLow,string"`
	MovingAverage50Day         *OptionalFloat64 `json:"50DayMovingAverage,string"`
	MovingAverage200Day        *OptionalFloat64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          *OptionalFloat64 `json:"SharesOutstanding,string"`
	SharesFloat                *OptionalFloat64 `json:"SharesFloat,string"`
	SharesShort                *OptionalFloat64 `json:"SharesShort,string"`
	SharesShortPriorMonth      *OptionalFloat64 `json:"SharesShortPriorMonth,string"`
	ShortRatio                 *OptionalFloat64 `json:"ShortRatio,string"`
	ShortPercentOutstanding    *OptionalFloat64 `json:"ShortPercentOutstanding,string"`
	ShortPercentFloat          *OptionalFloat64 `json:"ShortPercentFloat,string"`
	PercentInsiders            *OptionalFloat64 `json:"PercentInsiders,string"`
	PercentInstitutions        *OptionalFloat64 `json:"PercentInstitutions,string"`
	ForwardAnnualDividendRate  *OptionalFloat64 `json:"ForwardAnnualDividendRate,string"`
	ForwardAnnualDividendYield *OptionalFloat64 `json:"ForwardAnnualDividendYield,string"`
	PayoutRatio                *OptionalFloat64 `json:"PayoutRatio,string"`
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
