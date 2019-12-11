package main

type AggregateStorer interface {
	GetAggregate(buzzword string) (BuzzwordCounts, error)
	GetAllAggregates() ([]BuzzwordCounts, error)
	UpdateOrSetAggregate(ag *BuzzwordCounts) error
}
