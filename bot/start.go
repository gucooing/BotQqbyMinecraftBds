package bot

import (
	"context"
	"time"

	"github.com/gucooing/BotQqbyMinecraftBds/config"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

type MinecraftBot struct {
	Config   *config.Config
	botToken *token.Token
	api      openapi.OpenAPI
	ctx      context.Context
	ws       *dto.WebsocketAP
	me       *dto.User
}

func NewBot(conf *config.Config) *MinecraftBot {
	b := new(MinecraftBot)
	b.Config = conf

	botToken := token.BotToken(b.Config.AppID, b.Config.Token)
	b.botToken = botToken
	// api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second) // 使用NewSandboxOpenAPI创建沙箱环境的实例
	api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)
	b.api = api

	ctx := context.Background()
	b.ctx = ctx

	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		return nil
	}
	b.ws = ws
	me, err := api.Me(ctx)
	if err != nil {
		return nil
	}
	b.me = me

	return b
}

func (b *MinecraftBot) Run() error {
	var atMessage event.ATMessageEventHandler = b.atMessageEventHandler // @事件处理
	var guildEvent event.GuildEventHandler = b.guildHandler             // 频道事件处理

	intent := websocket.RegisterHandlers(atMessage, guildEvent) // 注册socket消息处理
	botgo.NewSessionManager().Start(b.ws, b.botToken, &intent)  // 启动socket监听
	return nil
}
