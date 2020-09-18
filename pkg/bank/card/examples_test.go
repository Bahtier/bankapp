package card
import(
	"fmt"
	"bank/pkg/bank/types"	
)
func ExamplePaymentSources()  {
	card := []types.Card{
		types.Card{
			PAN: "3006",
			Balance: 32_000,
			Active:  true,			
		},

		types.Card{
			PAN: "4006",
			Balance: 122_000,
			Active:  true,			
		},

		types.Card{
			PAN: "5006",
			Balance: 42_000,
			Active:  false,			
		},
	}

	payCard := PaymentSources(card)
	
	for _, payCard := range payCard {
		fmt.Println(payCard.Number)
	}
	
	//Output: 
	//3006 
	//4006
}