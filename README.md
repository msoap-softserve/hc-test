Library for parse chat messages
-------------------------------

Install
=======

    go get -u github.com/msoap-softserve/hc-test

Usage
=====

```Go
import (
    hchat github.com/msoap-softserve/hc-test
)

func main () {
    chatParser := hchat.New()
    json, err := chatParser.Parse("@chris you around?")
    json, err := chatParser.Parse("Good morning! (megusta) (coffee)")
    json, err := chatParser.Parse("Olympics are starting soon; http://www.nbcolympics.com")
    json, err := chatParser.Parse("@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016")
}
```
