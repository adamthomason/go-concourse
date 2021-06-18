package plans

type PlanItem interface {
	EnableItem()
}

type Plan []PlanItem
