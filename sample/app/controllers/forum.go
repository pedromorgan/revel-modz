package controllers

import (
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/modules/forum"
	"github.com/iassic/revel-modz/sample/app/routes"
)

func (c App) Forum(msg_pos, count int) revel.Result {
	if count == 0 {
		count = 20
	}
	topics, err := forum.GetTopicList(c.Txn, msg_pos, count)
	if err != nil {
		revel.ERROR.Println("Getting forum topic list: ", err)
	}

	return c.Render(topics)
}

func (c App) ForumTopic(topic_id, msg_id int) revel.Result {
	revel.INFO.Println("Forum: ", topic_id, msg_id)
	if msg_id != 0 {
		// enable the scroll to message
	}

	detail, err := forum.GetAllMessagesByTopicId(c.Txn, int64(topic_id))
	if err != nil {
		revel.ERROR.Println("Getting forum topic detail: ", topic_id, err)
	}

	return c.Render(detail)
}

func (c App) ForumMessage(topic_id, msg_id int) revel.Result {
	return c.Redirect(routes.App.ForumTopic(topic_id, msg_id))
}

func (c User) ForumTopicPost(subject, content string, tags []string) revel.Result {
	u := c.userConnected()
	revel.INFO.Println("Forum Topic POST: ", u.UserName, subject, content)

	err := forum.AddNewTopic(c.Txn, u.UserName, subject, content)
	if err != nil {
		revel.ERROR.Println("Posting forum topic: ", err)
		return c.RenderText("Error")
	}

	// Add topic tags?

	return c.RenderText("Success")
}

func (c User) ForumMessagePost(content string, topicId int64) revel.Result {
	u := c.userConnected()
	revel.INFO.Println("Forum Message POST: ", u.UserName, topicId, "\n", content)

	err := forum.AddNewMessage(c.Txn, u.UserName, content, topicId)
	if err != nil {
		revel.ERROR.Println("Posting forum message: ", err)
		return c.RenderText("Error")
	}

	return c.RenderText("Success")
}
