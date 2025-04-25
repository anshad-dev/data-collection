package models

type Lender struct {
	LenderName                   string `json:"lender_name" bson:"lender_name"`
	WhiteLabel                   string `json:"white_label" bson:"white_label"`
	RestrictedStates             string `json:"restricted_states" bson:"restricted_states"`
	MinCreditScore               *int   `json:"min_credit_score,omitempty" bson:"min_credit_score,omitempty"`
	Consolidation                *int   `json:"consolidation,omitempty" bson:"consolidation,omitempty"`
	SoleProp                     *int   `json:"sole_prop,omitempty" bson:"sole_prop,omitempty"`
	HomeBasedBusiness            *int   `json:"home_based_business,omitempty" bson:"home_based_business,omitempty"`
	DoesNotDo                    string `json:"does_not_do" bson:"does_not_do"`
	NonProfit                    string `json:"non_profit" bson:"non_profit"`
	DailyWeekly                  *int   `json:"daily_weekly,omitempty" bson:"daily_weekly,omitempty"`
	MaxNegativeDays              *int   `json:"max_negative_days,omitempty" bson:"max_negative_days,omitempty"`
	MaxAdvance                   *int   `json:"max_advance,omitempty" bson:"max_advance,omitempty"`
	Nsf                          *int   `json:"nsf,omitempty" bson:"nsf,omitempty"`
	MinBusinessTime              *int   `json:"min_business_time,omitempty" bson:"min_business_time,omitempty"`
	MinAmount                    *int   `json:"min_amount,omitempty" bson:"min_amount,omitempty"`
	MinDeposits                  *int   `json:"min_deposits,omitempty" bson:"min_deposits,omitempty"`
	MaxPosition                  *int   `json:"max_position,omitempty" bson:"max_position,omitempty"`
	MaxTerm                      *int   `json:"max_term,omitempty" bson:"max_term,omitempty"`
	Holdback                     string `json:"holdback" bson:"holdback"`
	IsBankruptcy                 bool   `json:"is_bankruptcy,omitempty" bson:"is_bankruptcy,omitempty"`
	BankruptcyMonths             *int   `json:"bankruptcy_months,omitempty" bson:"bankruptcy_months,omitempty"`
	BankruptcyCaseStatus         string `json:"bankruptcy_case_status" bson:"bankruptcy_case_status"`
	BankruptcyCloseStatus        string `json:"bankruptcy_close_status" bson:"bankruptcy_close_status"`
	IsTaxLien                    bool   `json:"is_tax_lien,omitempty" bson:"is_tax_lien,omitempty"`
	TaxLienMonths                *int   `json:"tax_lien_months,omitempty" bson:"tax_lien_months,omitempty"`
	TaxLienStatus                string `json:"tax_lien_status" bson:"tax_lien_status"`
	TaxLienSatisfiedHowLongAgo   string `json:"tax_lien_satisfied_how_long_ago" bson:"tax_lien_satisfied_how_long_ago"`
	IsCriminalHistory            bool   `json:"is_criminal_history,omitempty" bson:"is_criminal_history,omitempty"`
	IsDefaulted                  bool   `json:"is_defaulted,omitempty" bson:"is_defaulted,omitempty"`
	DefaultMonths                *int   `json:"default_months,omitempty" bson:"default_months,omitempty"`
	IsDefaultSatisfied           bool   `json:"is_default_satisfied,omitempty" bson:"is_default_satisfied,omitempty"`
	DefaultedSatisfiedHowLongAgo *int   `json:"defaulted_satisfied_how_long_ago,omitempty" bson:"defaulted_satisfied_how_long_ago,omitempty"`
	ProductType                  string `json:"product_type" bson:"product_type"`
	MinAvgDailyBalance           *int   `json:"min_avg_daily_balance,omitempty" bson:"min_avg_daily_balance,omitempty"`
}
