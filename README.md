# playingcards

go package that implements types and functions to represent playing cards and decks of cards

## example

```go
package main

import (
	"fmt"

	"github.com/notwithering/playingcards"
)

func main() {
	deck := playingcards.NewDeck()
	deck.Shuffle()
	fmt.Println(deck)
}
```

```
🂩🃁🂳🂣🃆🂲🂵🃚🂽🃝🃒🃓🃍🃗🂭🂥🂡🂤🂫🂪🃃🃅🂷🂨🂦🂴🂢🂧🃇🃔🂹🂾🃊🃙🂮🃄🂻🃞🂸🂺🃑🃕🃉🃘🃎🃂🃛🃈🃋🂱🃖🂶
```
