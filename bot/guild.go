package bot

import (
	"github.com/gucooing/BotQqbyMinecraftBds/logger"
	"github.com/tencent-connect/botgo/dto"
)

const (
	GuildCreateEvent = "GUILD_CREATE" // 机器人被加入到某个频道的事件
)

var guildId string

// 处理频道相关的事件
func (b *MinecraftBot) guildHandler(event *dto.WSPayload, data *dto.WSGuildData) error {
	if event.Type == GuildCreateEvent { // 当机器人加入频道时，获取频道的id
		guildId = data.ID
		logger.Debug("guildId = " + data.ID + " guildName = " + data.Name)
	}
	return nil
}
