package bot

import (
	"fmt"

	"github.com/nii236/oanda"
)

func New() *Guy {
	return &Guy{}
}

type Guy struct {
}

func (g *Guy) OnTick(bids []*oanda.Bid, asks []*oanda.Ask) {
	fmt.Println("bids:", bids)
	fmt.Println("asks:", asks)
}
