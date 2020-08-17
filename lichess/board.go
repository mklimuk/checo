package lichess

type EventType string

const (
	EventTypeChallenge EventType = "challenge"
	EventTypeGameStart EventType = "gameStart"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	Rating int `json:"rating"`
	Patron bool `json:"patron"`
	Online bool `json:"online"`
	Lag int `json:"lag"`
	Provisional bool `json:"provisional"`
}

type ChallengeEvent struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Challenger User `json:"challenger"`
	DestUser User `json:"destUser"`
	Variant struct {
		Key string `json:"key"`
		Name string `json:"name"`
		Short string `json:"short"`
	} `json:"variant"`
	Rated bool `json:"rated"`
	TimeControl struct {
		Type string `json:"type"`
		Limit int `json:"limit"`
		Increment int `json:"increment"`
		Show string `json:"show"`
	} `json:"timeControl"`
	Color string `json:"color"`
	Perf struct {
		Icon string `json:"icon"`
		Name string `json:"name"`
	} `json:"perf"`
}

type GameEvent struct {
	ID string `json:"id"`
}

type Event struct {
	Type EventType `json:"type"`
	Challenge *ChallengeEvent `json:"challenge"`
	GameStart *GameEvent `json:"gameStart"`
}

/*
[
  {
    "type": "gameFull",
    "id": "5IrD6Gzz",
    "rated": true,
    "variant": {
      "key": "standard",
      "name": "Standard",
      "short": "Std"
    },
    "clock": {
      "initial": 1200000,
      "increment": 10000
    },
    "speed": "classical",
    "perf": {
      "name": "Classical"
    },
    "createdAt": 1523825103562,
    "white": {
      "id": "lovlas",
      "name": "lovlas",
      "provisional": false,
      "rating": 2500,
      "title": "IM"
    },
    "black": {
      "id": "leela",
      "name": "leela",
      "rating": 2390,
      "title": null
    },
    "initialFen": "startpos",
    "state": {
      "type": "gameState",
      "moves": "e2e4 c7c5 f2f4 d7d6 g1f3 b8c6 f1c4 g8f6 d2d3 g7g6 e1g1 f8g7",
      "wtime": 7598040,
      "btime": 8395220,
      "winc": 10000,
      "binc": 10000,
      "status": "started"
    }
  },
  {
    "type": "gameState",
    "moves": "e2e4 c7c5 f2f4 d7d6 g1f3 b8c6 f1c4 g8f6 d2d3 g7g6 e1g1 f8g7 b1c3",
    "wtime": 7598040,
    "btime": 8395220,
    "winc": 10000,
    "binc": 10000,
    "status": "started"
  },
  {
    "type": "chatLine",
    "username": "thibault",
    "text": "Good luck, have fun",
    "room": "player"
  },
  {
    "type": "chatLine",
    "username": "lovlas",
    "text": "!eval",
    "room": "spectator"
  },
  {
    "type": "gameState",
    "moves": "e2e4 c7c5 f2f4 d7d6 g1f3 b8c6 f1c4 g8f6 d2d3 g7g6 e1g1 f8g7 b1c3",
    "wtime": 7598040,
    "btime": 8395220,
    "winc": 10000,
    "binc": 10000,
    "status": "resign",
    "winner": "black"
  }
]
 */

func (c *Client) StreamEvents(chan<- *Event) error {

}

func (c *Client) Seek() error {

}

func (c *Client) GameEvents(chan<- *GameEvent) error {

}