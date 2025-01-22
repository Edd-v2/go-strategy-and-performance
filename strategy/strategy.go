package strategy

type Strategy interface {
	Execute(workload int) error
}
