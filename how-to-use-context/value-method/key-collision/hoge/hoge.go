package hoge

import (
	"context"
	"fmt"
)

func SetValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "a", "c")
}

func GetValue(ctx context.Context) {
	val, ok := ctx.Value("a").(string)
	fmt.Println(val, ok)
}
