package shopware

// RuleCondition see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Rule/Aggregate/RuleCondition/RuleConditionEntity.php
type RuleCondition struct {
	Entity
	Typ      string           `json:"type"`
	RuleID   string           `json:"ruleId"`
	ParentID string           `json:"parentId"`
	Value    []string         `json:"value"`
	Rule     *Rule            `json:"rule"`
	Children []*RuleCondition `json:"children"`
	Parent   []*RuleCondition `json:"parent"`
	Position int              `json:"position"`
}
