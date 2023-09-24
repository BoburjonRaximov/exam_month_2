package main

import (
	"context"
	"fmt"
	"new_project/api"
	"new_project/api/handler"
	"new_project/config"
	"new_project/pkg/logger"
	"new_project/storage/memory"
)

func main() {

	cfg := config.LoadP()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := memory.NewStorage(context.Background(), cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	h := handler.NewHandler(strg, *config.Load(), log)

	r := api.NewServer(h)
	r.Run(":8080")
}
