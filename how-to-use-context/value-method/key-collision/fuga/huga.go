package fuga

import (
	"context"
	"fmt"
)

func SetValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "a", "b")
}

func GetValue(ctx context.Context) {
	val, ok := ctx.Value("a").(string)
	fmt.Println(val, ok)
}
