package bot

import (
	"strings"

	"github.com/gucooing/BotQqbyMinecraftBds/server"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

func (b *MinecraftBot) atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	res := message.ETLInput(data.Content) // 去掉@结构和清除前后空格
	if strings.HasPrefix(res, "/") {      // 去掉/
		res = strings.Replace(res, "/", "", 1)
	}

	// 按空格分割字符串
	parts := strings.SplitN(res, " ", 2)

	// 输出分割后的结果
	switch parts[0] {
	case "绑定":
		if len(parts) != 2 {
			b.SendQQ(data, "参数非法")
			return nil
		}
		if b.FirstUser(data.Author.ID).Name != "" {
			b.SendQQ(data, "请勿重复绑定")
			return nil
		}
		b.AddUser(parts[1], data.Author.ID)
		server.Sender("whitelist add " + "\"" + parts[1] + "\"")
		b.SendQQ(data, "绑定玩家:"+"\""+parts[1]+"\""+" 成功")
		return nil
	case "解绑":
		user := b.DeleteUser(data.Author.ID)
		if user.Name != "" {
			server.Sender("whitelist remove " + user.Name)
			b.SendQQ(data, "解绑成功")
			return nil
		} else {
			b.SendQQ(data, "解绑失败")
			return nil
		}
	case "cmd":
		if len(parts) != 2 {
			b.SendQQ(data, "参数非法")
			return nil
		}
		// server.Sender("whitelist add " + parts[id+1])
		b.SendQQ(data, "不是管理员")
	default:
	}

	return nil
}

func (b *MinecraftBot) SendQQ(data *dto.WSATMessageData, msg string) {
	b.api.PostMessage(b.ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: msg})
}
