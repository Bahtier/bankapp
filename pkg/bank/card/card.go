package card

import(
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSources {
	payCard := []types.PaymentSources{}

	for _, cards := range cards {
		if cards.Balance > 0 && cards.Active{
			payCard = append(payCard, types.PaymentSources{				
				cards.Balance,
				string(cards.PAN),
				},)									
			}	
		}	
		return payCard		
	}	

	
