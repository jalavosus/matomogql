// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jalavosus/matomogql/graph/scalars"
	"github.com/shopspring/decimal"
)

type AggregateContinentInfo struct {
	Continent  string `json:"continent"`
	NumVisits  int    `json:"nb_visits"`
	PrettyName string `json:"prettyName"`
}

type AggregateCountryInfo struct {
	Country    string  `json:"country"`
	NumVisits  int     `json:"nb_visits"`
	Flag       *string `json:"flag,omitempty"`
	PrettyName string  `json:"prettyName"`
}

type AggregateDeviceInfo struct {
	Type    string             `json:"type"`
	Count   int                `json:"count"`
	Icon    *string            `json:"icon,omitempty"`
	Devices []*ShortDeviceInfo `json:"devices,omitempty"`
}

type BrowserInfo struct {
	Family            string  `json:"family"`
	FamilyDescription string  `json:"familyDescription"`
	Browser           string  `json:"browser"`
	Name              string  `json:"name"`
	Icon              *string `json:"icon,omitempty"`
	Code              string  `json:"code"`
	Version           string  `json:"version"`
}

type CampaignInfo struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Keyword   string `json:"keyword"`
	Medium    string `json:"medium"`
	Name      string `json:"name"`
	Source    string `json:"source"`
	Group     string `json:"group"`
	Placement string `json:"placement"`
}

type ConvertedVisitsOptions struct {
	Date graphql.Omittable[*DateRangeOptions] `json:"date,omitempty"`
}

type DateRangeOptions struct {
	Period    SegmentPeriod              `json:"period"`
	StartDate string                     `json:"startDate"`
	EndDate   graphql.Omittable[*string] `json:"endDate,omitempty"`
}

type DeviceInfo struct {
	Type                   string  `json:"type"`
	TypeIcon               *string `json:"typeIcon,omitempty"`
	Brand                  string  `json:"brand"`
	Model                  string  `json:"model"`
	OperatingSystem        string  `json:"operatingSystem"`
	OperatingSystemName    string  `json:"operatingSystemName"`
	OperatingSystemIcon    *string `json:"operatingSystemIcon,omitempty"`
	OperatingSystemCode    string  `json:"operatingSystemCode"`
	OperatingSystemVersion string  `json:"operatingSystemVersion"`
	Resolution             *string `json:"resolution,omitempty"`
}

type EcommerceGoal struct {
	IDSite                int     `json:"idsite"`
	Label                 string  `json:"label"`
	Revenue               int     `json:"revenue"`
	Quantity              int     `json:"quantity"`
	Orders                int     `json:"orders"`
	AveragePrice          float64 `json:"avg_price"`
	AverageQuantity       float64 `json:"avg_quantity"`
	NumVisits             int     `json:"nb_visits"`
	NumActions            int     `json:"nb_actions"`
	ConversionRatePercent string  `json:"conversion_rate"`
	//  Parsed from conversionRatePercent
	ConversionRate            float64  `json:"-"`
	Segment                   string   `json:"segment"`
	SumDailyNumUniqueVisitors int      `json:"sum_daily_nb_uniq_visitors"`
	ConvertedVisits           []*Visit `json:"convertedVisits,omitempty"`
}

type GetEcommerceGoalsOptions struct {
	Date *DateRangeOptions `json:"date"`
}

type GetGoalsOptions struct {
	OrderByName graphql.Omittable[*bool] `json:"orderByName,omitempty"`
}

type Goal struct {
	IDSite              int              `json:"idsite"`
	IDGoal              int              `json:"idgoal"`
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	MatchAttribute      string           `json:"match_attribute"`
	Pattern             *string          `json:"pattern,omitempty"`
	PatternType         *string          `json:"pattern_type,omitempty"`
	CaseSensitive       *scalars.BoolInt `json:"case_sensitive"`
	AllowMultiple       scalars.BoolInt  `json:"allow_multiple,omitempty"`
	Revenue             scalars.BoolInt  `json:"revenue"`
	Deleted             scalars.BoolInt  `json:"deleted"`
	EventValueAsRevenue int              `json:"event_value_as_revenue,omitempty"`
	ConvertedVisits     []*Visit         `json:"convertedVisits,omitempty"`
}

type LastVisitsOpts struct {
	Date     graphql.Omittable[*DateRangeOptions] `json:"date,omitempty"`
	Segments graphql.Omittable[[]string]          `json:"segments,omitempty"`
	Limit    graphql.Omittable[*int]              `json:"limit,omitempty"`
}

type Location struct {
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	CountryFlag   *string `json:"countryFlag,omitempty"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"regionCode"`
	City          string  `json:"city"`
	Location      string  `json:"location"`
	Latitude      string  `json:"latitude"`
	Longitude     string  `json:"longitude"`
}

type OrderByOptions struct {
	Timestamp graphql.Omittable[*OrderBy] `json:"timestamp,omitempty"`
}

type Query struct {
}

type ReferrerInfo struct {
	Type              *string `json:"type,omitempty"`
	Name              *string `json:"name,omitempty"`
	Keyword           *string `json:"keyword,omitempty"`
	KeywordPosition   *string `json:"keywordPosition,omitempty"`
	URL               *string `json:"url,omitempty"`
	SearchEngineURL   *string `json:"searchEngineUrl,omitempty"`
	SearchEngineIcon  *string `json:"searchEngineIcon,omitempty"`
	SocialNetworkURL  *string `json:"socialNetworkUrl,omitempty"`
	SocialNetworkIcon *string `json:"socialNetworkIcon,omitempty"`
}

type ShortDeviceInfo struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Site struct {
	IDSite                       int                `json:"idsite"`
	Name                         string             `json:"name"`
	MainURL                      string             `json:"main_url"`
	TsCreated                    string             `json:"ts_created"`
	Ecommerce                    int                `json:"ecommerce"`
	Sitesearch                   int                `json:"sitesearch"`
	SitesearchKeywordParameters  string             `json:"sitesearch_keyword_parameters"`
	SitesearchCategoryParameters string             `json:"sitesearch_category_parameters"`
	Timezone                     string             `json:"timezone"`
	TimezoneName                 string             `json:"timezone_name"`
	Currency                     string             `json:"currency"`
	CurrencyName                 string             `json:"currency_name"`
	KeepURLFragment              scalars.BoolInt    `json:"keep_url_fragment"`
	ExcludeUnknownUrls           scalars.BoolInt    `json:"excludeUnknownUrls"`
	ExcludedIPs                  scalars.StringList `json:"excluded_ips"`
	ExcludedParameters           scalars.StringList `json:"excluded_parameters"`
	ExcludedUserAgents           scalars.StringList `json:"excluded_user_agents"`
	ExcludedReferrers            scalars.StringList `json:"excluded_referrers"`
	Group                        string             `json:"group"`
	Type                         string             `json:"type"`
	Goals                        []*Goal            `json:"goals,omitempty"`
	LastVisits                   []*Visit           `json:"lastVisits,omitempty"`
}

type Visit struct {
	IDSite                         int                   `json:"idSite"`
	SiteName                       string                `json:"siteName"`
	SiteCurrency                   string                `json:"siteCurrency"`
	SiteCurrencySymbol             string                `json:"siteCurrencySymbol"`
	IDVisit                        int                   `json:"idVisit"`
	VisitIP                        string                `json:"visitIp"`
	VisitorID                      string                `json:"visitorId"`
	Fingerprint                    string                `json:"fingerprint"`
	VisitorProfile                 *VisitorProfile       `json:"visitorProfile,omitempty"`
	VisitServerHour                string                `json:"visitServerHour"`
	GoalConversions                int                   `json:"goalConversions"`
	ActionDetails                  []*VisitActionDetails `json:"actionDetails,omitempty"`
	ServerDate                     string                `json:"serverDate"`
	ServerDatePretty               string                `json:"serverDatePretty"`
	ServerTimestamp                int                   `json:"serverTimestamp"`
	ServerTimePretty               string                `json:"serverTimePretty"`
	FirstActionTimestamp           int                   `json:"firstActionTimestamp"`
	LastActionTimestamp            int                   `json:"lastActionTimestamp"`
	LastActionDateTime             string                `json:"lastActionDateTime"`
	ServerDatePrettyFirstAction    string                `json:"serverDatePrettyFirstAction"`
	ServerTimePrettyFirstAction    string                `json:"serverTimePrettyFirstAction"`
	UserID                         *string               `json:"userId,omitempty"`
	VisitorType                    string                `json:"visitorType"`
	VisitorTypeIcon                *string               `json:"visitorTypeIcon,omitempty"`
	VisitConverted                 int                   `json:"visitConverted"`
	VisitConvertedIcon             *string               `json:"visitConvertedIcon,omitempty"`
	VisitCount                     *int                  `json:"visitCount,omitempty"`
	VisitEcommerceStatus           *string               `json:"visitEcommerceStatus,omitempty"`
	VisitEcommerceStatusIcon       *string               `json:"visitEcommerceStatusIcon,omitempty"`
	DaysSinceFirstVisit            int                   `json:"daysSinceFirstVisit"`
	SecondsSinceFirstVisit         int                   `json:"secondsSinceFirstVisit"`
	DaysSinceLastEcommerceOrder    int                   `json:"daysSinceLastEcommerceOrder"`
	SecondsSinceLastEcommerceOrder *int                  `json:"secondsSinceLastEcommerceOrder,omitempty"`
	VisitDuration                  int                   `json:"visitDuration"`
	VisitDurationPretty            string                `json:"visitDurationPretty"`
	Searches                       int                   `json:"searches"`
	Actions                        int                   `json:"actions"`
	Interactions                   int                   `json:"interactions"`
	LanguageCode                   string                `json:"languageCode"`
	Language                       string                `json:"language"`
	DeviceInfo                     *DeviceInfo           `json:"deviceInfo,omitempty"`
	DeviceType                     string                `json:"deviceType"`
	DeviceTypeIcon                 *string               `json:"deviceTypeIcon,omitempty"`
	DeviceBrand                    string                `json:"deviceBrand"`
	DeviceModel                    string                `json:"deviceModel"`
	OperatingSystem                string                `json:"operatingSystem"`
	OperatingSystemName            string                `json:"operatingSystemName"`
	OperatingSystemIcon            *string               `json:"operatingSystemIcon,omitempty"`
	OperatingSystemCode            string                `json:"operatingSystemCode"`
	OperatingSystemVersion         string                `json:"operatingSystemVersion"`
	Resolution                     *string               `json:"resolution,omitempty"`
	BrowserInfo                    *BrowserInfo          `json:"browserInfo,omitempty"`
	BrowserFamily                  string                `json:"browserFamily"`
	BrowserFamilyDescription       string                `json:"browserFamilyDescription"`
	Browser                        string                `json:"browser"`
	BrowserName                    string                `json:"browserName"`
	BrowserIcon                    *string               `json:"browserIcon,omitempty"`
	BrowserCode                    string                `json:"browserCode"`
	BrowserVersion                 string                `json:"browserVersion"`
	Events                         int                   `json:"events"`
	LocationInfo                   *Location             `json:"locationInfo,omitempty"`
	Continent                      string                `json:"continent"`
	ContinentCode                  string                `json:"continentCode"`
	Country                        string                `json:"country"`
	CountryCode                    string                `json:"countryCode"`
	CountryFlag                    *string               `json:"countryFlag,omitempty"`
	Region                         string                `json:"region"`
	RegionCode                     string                `json:"regionCode"`
	City                           string                `json:"city"`
	Location                       string                `json:"location"`
	Latitude                       string                `json:"latitude"`
	Longitude                      string                `json:"longitude"`
	VisitLocalTime                 string                `json:"visitLocalTime"`
	VisitLocalHour                 string                `json:"visitLocalHour"`
	DaysSinceLastVisit             int                   `json:"daysSinceLastVisit"`
	SecondsSinceLastVisit          int                   `json:"secondsSinceLastVisit"`
	Plugins                        *string               `json:"plugins,omitempty"`
	AdClickID                      string                `json:"adClickId"`
	AdProviderID                   string                `json:"adProviderId"`
	AdProviderName                 string                `json:"adProviderName"`
	FormConversions                int                   `json:"formConversions"`
	SessionReplayURL               *string               `json:"sessionReplayUrl,omitempty"`
	CampaignInfo                   *CampaignInfo         `json:"campaignInfo,omitempty"`
	CampaignID                     string                `json:"campaignId"`
	CampaignContent                string                `json:"campaignContent"`
	CampaignKeyword                string                `json:"campaignKeyword"`
	CampaignMedium                 string                `json:"campaignMedium"`
	CampaignName                   string                `json:"campaignName"`
	CampaignSource                 string                `json:"campaignSource"`
	CampaignGroup                  string                `json:"campaignGroup"`
	CampaignPlacement              string                `json:"campaignPlacement"`
}

type VisitActionDetails struct {
	Type             string           `json:"type"`
	URL              string           `json:"url"`
	Title            string           `json:"title"`
	Subtitle         string           `json:"subtitle"`
	PageTitle        string           `json:"pageTitle"`
	PageIDAction     int              `json:"pageIdAction"`
	IDPageView       string           `json:"idpageview"`
	ServerTimePretty string           `json:"serverTimePretty"`
	PageID           int              `json:"pageId"`
	TimeSpent        int              `json:"timeSpent"`
	TimeSpentPretty  string           `json:"timeSpentPretty"`
	PageViewPosition int              `json:"pageViewPosition"`
	Timestamp        int              `json:"timestamp"`
	GoalPageID       *int             `json:"goalPageId,omitempty"`
	Revenue          *decimal.Decimal `json:"revenue,omitempty"`
	RevenueSubTotal  *decimal.Decimal `json:"revenueSubTotal,omitempty"`
	EventCategory    *string          `json:"eventCategory,omitempty"`
	EventAction      *string          `json:"eventAction,omitempty"`
	EventName        *string          `json:"eventName,omitempty"`
	GoalName         *string          `json:"goalName,omitempty"`
	Referrer         *ReferrerInfo    `json:"referrer,omitempty"`
	ReferrerType     *string          `json:"referrerType,omitempty"`
	ReferrerName     *string          `json:"referrerName,omitempty"`
	ReferrerKeyword  *string          `json:"referrerKeyword,omitempty"`
}

type VisitorFirstLastVisit struct {
	Date            int    `json:"date"`
	PrettyDate      string `json:"prettyDate"`
	DaysAgo         int    `json:"daysAgo"`
	ReferrerType    string `json:"referrerType"`
	ReferrerURL     string `json:"referrerUrl"`
	RefferalSummary string `json:"refferalSummary"`
}

type VisitorProfile struct {
	VisitorID  string                    `json:"visitorId"`
	FirstVisit *VisitorFirstLastVisit    `json:"firstVisit"`
	LastVisit  *VisitorFirstLastVisit    `json:"lastVisit"`
	LastVisits []*Visit                  `json:"lastVisits,omitempty"`
	Devices    []*AggregateDeviceInfo    `json:"devices,omitempty"`
	Countries  []*AggregateCountryInfo   `json:"countries,omitempty"`
	Continents []*AggregateContinentInfo `json:"continents,omitempty"`
}

type OrderBy string

const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

var AllOrderBy = []OrderBy{
	OrderByAsc,
	OrderByDesc,
}

func (e OrderBy) IsValid() bool {
	switch e {
	case OrderByAsc, OrderByDesc:
		return true
	}
	return false
}

func (e OrderBy) String() string {
	return string(e)
}

func (e *OrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderBy", str)
	}
	return nil
}

func (e OrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SegmentPeriod string

const (
	SegmentPeriodDay   SegmentPeriod = "Day"
	SegmentPeriodWeek  SegmentPeriod = "Week"
	SegmentPeriodMonth SegmentPeriod = "Month"
	SegmentPeriodRange SegmentPeriod = "Range"
)

var AllSegmentPeriod = []SegmentPeriod{
	SegmentPeriodDay,
	SegmentPeriodWeek,
	SegmentPeriodMonth,
	SegmentPeriodRange,
}

func (e SegmentPeriod) IsValid() bool {
	switch e {
	case SegmentPeriodDay, SegmentPeriodWeek, SegmentPeriodMonth, SegmentPeriodRange:
		return true
	}
	return false
}

func (e SegmentPeriod) String() string {
	return string(e)
}

func (e *SegmentPeriod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SegmentPeriod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SegmentPeriod", str)
	}
	return nil
}

func (e SegmentPeriod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
