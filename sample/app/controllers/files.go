package controllers

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"strconv"

	"github.com/iassic/revel-modz/modules/user-files"
	"github.com/revel/revel"
)

type Files struct {
	User
}

func (c Files) FilesQuery() revel.Result {
	user := c.userConnected()

	dsInfos, err := userfiles.GetUserFileInfos(c.Txn, user.UserId)
	checkERROR(err)

	if dsInfos == nil || len(dsInfos) == 0 {
		revel.WARN.Println("No files on server")
		return c.RenderText("NOFILES")
	}

	return c.RenderJson(dsInfos)
}

func (c Files) FileUpload(filedata []byte) revel.Result {
	user := c.userConnected()

	reader := c.Request.Request.Body
	body, err := ioutil.ReadAll(reader)
	checkERROR(err)

	var user_file userfiles.UserFileWire
	err = json.Unmarshal(body, &user_file)
	checkERROR(err)

	var udt userfiles.UserFileInfo
	udt.UserId = user.UserId
	udt.Name = user_file.Name
	udt.Path = user_file.Path
	udt.Folder = user_file.Folder
	udt.Size = user_file.Size
	content := []byte(user_file.Content)

	err = userfiles.AddUserFile(c.Txn, &udt, content)
	checkERROR(err)

	// TODO if err != nil, should return message to client stating such
	return c.RenderText("added")
}

func (c Files) FileContent(data_id string) revel.Result {
	user := c.userConnected()

	dsId, err := strconv.ParseInt(data_id, 10, 64)
	checkERROR(err)

	uds, err := userfiles.GetUserFileById(c.Txn, user.UserId, dsId)
	checkERROR(err)

	content := string(uds.Content)
	return c.RenderText(content)
}
