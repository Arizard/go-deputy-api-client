package payengine

type PayRule struct {
	DisplayName string
	Rank int
	Conditions []PayRuleCondition
}

type PayRuleCondition struct {
	DisplayName string
	Kind string
	Options map[string]interface{}
}
