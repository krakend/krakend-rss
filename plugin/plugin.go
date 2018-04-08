package main

import (
	"fmt"
	"io"

	"github.com/devopsfaith/krakend-rss"
)

var Registrable registrable

type registrable int

func (r *registrable) RegisterDecoder(setter func(name string, dec func(bool) func(io.Reader, *map[string]interface{}) error) error) error {
	fmt.Println("registrable", r, "from plugin 'krakend_rss' is registering its decoders at", setter)

	setter(rss.Name, rss.DecoderFactory)

	return nil
}

func init() {
	fmt.Println("plugin 'krakend_rss' loaded!")
}

func main() {

}
