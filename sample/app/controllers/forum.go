package controllers

import (
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/modules/forum"
	"github.com/iassic/revel-modz/sample/app/routes"
)

func (c App) Forum() revel.Result {
	return c.Render()
}

func (c App) ForumTopic(topic_id, msg_id int) revel.Result {
	revel.INFO.Println("Forum: ", topic_id, msg_id)
	if msg_id != 0 {
		// enable the scroll to message
	}

	topics, err := forum.GetTopicList(c.Txn)
	if err != nil {
		revel.ERROR.Println("Getting forum topic list: ", err)
	}

	return c.Render(topics)
}

func (c App) ForumMessage(topic_id, msg_id int) revel.Result {
	return c.Redirect(routes.App.ForumTopic(topic_id, msg_id))
}

func (c User) ForumPost(author, subject, content string, tags []string) revel.Result {
	revel.INFO.Println("Forum POST: ", author, subject, content, tags)
	return c.Render()
}
