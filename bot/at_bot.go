package bot

import (
	"strings"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

func (b *MinecraftBot) atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	res := message.ETLInput(data.Content) // 去掉@结构和清除前后空格
	if strings.HasPrefix(res, "/") {      // 去掉/
		res = strings.Replace(res, "/", "", 1)
	}

	switch res {
	case "":

	}
	return nil
}
