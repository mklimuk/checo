package lichess

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Perf struct {
	Games int `json:"games"`
	Prog int `json:"prog"`
	Rating int `json:"rating"`
	RD int `json:"rd"`
}

type Account struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Title string `json:"title"`
	Online bool `json:"online"`
	Playing bool `json:"playing"`
	Streaming bool `json:"streaming"`
	CreatedAt uint64 `json:"createdAt"`
	SeenAt uint64 `json:"seenAt"`
	Profile struct {
		Bio string `json:"bio"`
		Country string `json:"country"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Links string `json:"links"`
		Location string `json:"location"`
	} `json:"profile"`
	NbFollowers int `json:"nbFollowers"`
	NbFollowing int `json:"nbFollowing"`
	CompletionRate int `json:"completionRate"`
	Language string `json:"language"`
	Count struct {
		AI int `json:"ai"`
		All int `json:"all"`
		Bookmark int `json:"bookmark"`
		Draw int `json:"draw"`
		DrawH int `json:"drawH"`
		Import int `json:"import"`
		Loss int `json:"loss"`
		LossH int `json:"lossH"`
		Me int `json:"me"`
		Playing int `json:"playing"`
		Rated int `json:"rated"`
		Win int `json:"win"`
		WinH int `json:"winH"`
	} `json:"count"`
	Perfs struct {
		Blitz    Perf `json:"blitz"`
		Bullet   Perf `json:"bullet"`
		Chess960 Perf `json:"chess960"`
		Puzzle   Perf `json:"puzzle"`
	} `json:"perfs"`
	Patron bool `json:"patron"`
	Disabled bool `json:"disabled"`
	Engine bool `json:"engine"`
	Booster bool `json:"booster"`
	PlayTime struct {
		Total int `json:"total"`
		TV int `json:"tv"`
	} `json:"playTime"`
}

func (c *Client) GetAccount() (*Account, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/account", c.baseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("could not build request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform http request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http response status: %d", res.StatusCode)
	}
	dec := json.NewDecoder(res.Body)
	defer res.Body.Close()
	var acc Account
	err = dec.Decode(&acc)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}
	return &acc, nil
}