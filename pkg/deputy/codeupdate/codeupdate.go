package codeupdate

type ScriptOptions struct {
	DecafBase64 string `json:"strDecaf,omitempty"`
	DeXMLBase64 string `json:"strDexml,omitempty"`
	Label       string `json:"strLabel,omitempty"`
	ScriptType  uint   `json:"intScriptType"`
	TriggerORM  uint   `json:"strTriggerOrm,omitempty"`
}

type ReportOptions struct {
	Name             string `json:"strName"`
	IsDev            uint   `json:"blnInDev"`
	DateFilter       uint   `json:"intDateFilter"`
	CompanyFilter    uint   `json:"blnCompanyFilter"`
	Category         uint   `json:"intCategory"`
	HTMLBase64       string `json:"strBody"`
	JavascriptBase64 string `json:"strJsCode"`
	DecafBase64      string `json:"strDecaf,omitempty"`
	DeXMLBase64      string `json:"strDexml,omitempty"`
	CSSBase64        string `json:"strCss"`
	WidgetDexml      string `json:"strWidgetDexml,omitempty"`
}

type CustomAppOptions struct {
	Name             string `json:"strName"`
	HTMLBase64       string `json:"strBody"`
	JavascriptBase64 string `json:"strJsCode"`
	DecafBase64      string `json:"strDecaf,omitempty"`
	DeXMLBase64      string `json:"strDexml,omitempty"`
	CSSBase64        string `json:"strCss"`
}

const (
	HelperSupportScript  uint = 1  // 1 - Helper Support Script
	ExecuteOnLogin            = 2  // 2 - Execute on Login
	ExecuteAtCron             = 3  // 3 - Execute at Cron
	RecordValidation          = 4  // 4 - Record Validation*
	PeriodPayCalculation      = 6  // 6 - Period Pay Calculation
	RecordAfterSave           = 7  // 7 - Record After Save*
	ExecuteViaApi             = 8  // 8 - Executable via SOAP/REST api
	PayrollExport             = 9  // 9 - Payroll Export Script
	ApplicationLaunch         = 10 // 10 - Application Launch*
	ShiftPayCalculation       = 11 // 11 - Shift Pay Calculation
	KioskLaunch               = 12 // 12 - Kiosk Launch
	Notification              = 13 // 13 - Notification
	Widget                    = 14 // 14 - Widget*
	RecordAfterDelete         = 15 // 15 - Record After Delete*
)
