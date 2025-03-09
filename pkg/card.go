package pokemontcgv2

import (
	"encoding/json"
	"fmt"

	"github.com/tcghouse/pokemon-tcg-sdk-go-v2/pkg/request"
)

// A PokemonCard represents a Pokemon card and its data.
type PokemonCard struct {
	ID           string   `json:"id" db:"id"`
	Name         string   `json:"name" db:"name"`
	Supertype    string   `json:"supertype" db:"supertype"`
	Subtypes     []string `json:"subtypes" db:"subtypes"`
	Level        string   `json:"level" db:"level"`
	Hp           string   `json:"hp" db:"hp"`
	Types        []string `json:"types" db:"types"`
	EvolvesFrom  string   `json:"evolvesFrom" db:"evolvesFrom"`
	EvolvesTo    []string `json:"evolvesTo" db:"evolvesTo"`
	Rules        []string `json:"rules" db:"rules"`
	AncientTrait *struct {
		Name string `json:"name" db:"name"`
		Text string `json:"text" db:"text"`
	} `json:"ancientTrait" db:"ancientTrait"`
	Abilities []struct {
		Name string `json:"name" db:"name"`
		Text string `json:"text" db:"text"`
		Type string `json:"type" db:"type"`
	} `json:"abilities" db:"abilities"`
	Attacks []struct {
		Name                string   `json:"name" db:"name"`
		Cost                []string `json:"cost" db:"cost"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost" db:"convertedEnergyCost"`
		Damage              string   `json:"damage" db:"damage"`
		Text                string   `json:"text" db:"text"`
	} `json:"attacks" db:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type" db:"type"`
		Value string `json:"value" db:"value"`
	} `json:"weaknesses" db:"weaknesses"`
	Resistances []struct {
		Type  string `json:"type" db:"type"`
		Value string `json:"value" db:"value"`
	} `json:"resistances" db:"resistances"`
	RetreatCost          []string `json:"retreatCost" db:"retreatCost"`
	ConvertedRetreatCost int      `json:"convertedRetreatCost" db:"convertedRetreatCost"`
	Set                  struct {
		ID           string `json:"id" db:"id"`
		Name         string `json:"name" db:"name"`
		Series       string `json:"series" db:"series"`
		PrintedTotal int    `json:"printedTotal" db:"printedTotal"`
		Total        int    `json:"total" db:"total"`
		Legalities   struct {
			Unlimited string `json:"unlimited" db:"unlimited"`
			Standard  string `json:"standard" db:"standard"`
			Expanded  string `json:"expanded" db:"expanded"`
		} `json:"legalities" db:"legalities"`
		PtcgoCode   string `json:"ptcgoCode" db:"ptcgoCode"`
		ReleaseDate string `json:"releaseDate" db:"releaseDate"`
		UpdatedAt   string `json:"updatedAt" db:"updatedAt"`
		Images      struct {
			Symbol string `json:"symbol" db:"symbol"`
			Logo   string `json:"logo" db:"logo"`
		} `json:"images" db:"images"`
	} `json:"set" db:"set"`
	Number                 string `json:"number" db:"number"`
	Artist                 string `json:"artist" db:"artist"`
	Rarity                 string `json:"rarity" db:"rarity"`
	FlavorText             string `json:"flavorText" db:"flavorText"`
	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers" db:"nationalPokedexNumbers"`
	Legalities             struct {
		Unlimited string `json:"unlimited" db:"unlimited"`
		Standard  string `json:"standard" db:"standard"`
		Expanded  string `json:"expanded" db:"expanded"`
	} `json:"legalities" db:"legalities"`
	RegulationMark string `json:"regulationMark"`
	Images         struct {
		Small string `json:"small" db:"small"`
		Large string `json:"large" db:"large"`
	} `json:"images" db:"images"`
	TCGPlayer *struct {
		URL       string `json:"url" db:"url"`
		UpdatedAt string `json:"updatedAt" db:"updatedAt"`
		Prices    struct {
			Holofoil *struct {
				Low    float64 `json:"low" db:"low"`
				Mid    float64 `json:"mid" db:"mid"`
				High   float64 `json:"high" db:"high"`
				Market float64 `json:"market" db:"market"`
			} `json:"holofoil,omitempty" db:"holofoil,omitempty"`
			ReverseHolofoil *struct {
				Low    float64 `json:"low" db:"low"`
				Mid    float64 `json:"mid" db:"mid"`
				High   float64 `json:"high" db:"high"`
				Market float64 `json:"market" db:"market"`
			} `json:"reverseHolofoil,omitempty" db:"reverseHolofoil,omitempty"`
			Normal *struct {
				Low    float64 `json:"low" db:"low"`
				Mid    float64 `json:"mid" db:"mid"`
				High   float64 `json:"high" db:"high"`
				Market float64 `json:"market" db:"market"`
			} `json:"normal,omitempty" db:"normal,omitempty"`
		} `json:"prices" db:"prices"`
	} `json:"tcgplayer,omitempty" db:"tcgplayer,omitempty"`
	CardMarket *struct {
		URL       string `json:"url" db:"url"`
		UpdatedAt string `json:"updatedAt" db:"updatedAt"`
		Prices    struct {
			AverageSellPrice *float64 `json:"averageSellPrice" db:"averageSellPrice"`
			LowPrice         *float64 `json:"lowPrice" db:"lowPrice"`
			TrendPrice       *float64 `json:"trendPrice" db:"trendPrice"`
			GermanProLow     *float64 `json:"germanProLow" db:"germanProLow"`
			SuggestedPrice   *float64 `json:"suggestedPrice" db:"suggestedPrice"`
			ReverseHoloSell  *float64 `json:"reverseHoloSell" db:"reverseHoloSell"`
			ReverseHoloLow   *float64 `json:"reverseHoloLow" db:"reverseHoloLow"`
			ReverseHoloTrend *float64 `json:"reverseHoloTrend" db:"reverseHoloTrend"`
			LowPriceExPlus   *float64 `json:"lowPriceExPlus" db:"lowPriceExPlus"`
			Avg1             *float64 `json:"avg1" db:"avg1"`
			Avg7             *float64 `json:"avg7" db:"avg7"`
			Avg30            *float64 `json:"avg30" db:"avg30"`
			ReverseHoloAvg1  *float64 `json:"reverseHoloAvg1" db:"reverseHoloAvg1"`
			ReverseHoloAvg7  *float64 `json:"reverseHoloAvg7" db:"reverseHoloAvg7"`
			ReverseHoloAvg30 *float64 `json:"reverseHoloAvg30" db:"reverseHoloAvg30"`
		} `json:"prices" db:"prices"`
	} `json:"cardmarket,omitempty" db:"cardmarket,omitempty"`
}

// GetCards allows you to search and filter for cards using given options.
// Docs: https://docs.pokemontcg.io/#api_v2cards_list
func (c *apiClient) GetCards(o ...request.Option) ([]*PokemonCard, error) {
	r := request.New(endpointCards, o...)

	u, err := r.GetURL()
	if err != nil {
		return nil, err
	}

	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cardList := struct {
		Data []*PokemonCard `json:"data"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&cardList); err != nil {
		return nil, err
	}

	return cardList.Data, nil
}

// GetCardByID returns a single pokemon card.
// Docs: https://docs.pokemontcg.io/#api_v2cards_get
func (c *apiClient) GetCardByID(id string) (*PokemonCard, error) {
	r := request.New(fmt.Sprintf(endpointCard, id))
	u, err := r.GetURL()
	if err != nil {
		return nil, err
	}

	cards := make(map[string]PokemonCard)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&cards); err != nil {
		return nil, err
	}

	card, ok := cards["data"]
	if !ok {
		return nil, err
	}

	return &card, nil
}
