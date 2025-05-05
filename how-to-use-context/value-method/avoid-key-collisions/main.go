package main

import (
	"context"

	"github.com/ryichk/go-practice/how-to-use-context/value-method/avoid-key-collisions/fuga"
	"github.com/ryichk/go-practice/how-to-use-context/value-method/avoid-key-collisions/hoge"
)

func main() {
	ctx := context.Background()
	ctx = hoge.SetValue(ctx)
	ctx = fuga.SetValue(ctx)

	hoge.GetValue(ctx)
	fuga.GetValue(ctx)
}