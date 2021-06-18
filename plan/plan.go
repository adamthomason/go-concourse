package plan

type PlanItem interface {
	EnableItem()
}

type Plan []PlanItem
