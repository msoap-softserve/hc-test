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
    json, err := hchat.Parse("@chris you around?")
    json, err := hchat.Parse("Good morning! (megusta) (coffee)")
    json, err := hchat.Parse("Olympics are starting soon; http://www.nbcolympics.com")
    json, err := hchat.Parse("@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016")
}
```
