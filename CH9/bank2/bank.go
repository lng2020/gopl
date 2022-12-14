package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func Balancer() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
