package _struct

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"testing"
)

//嵌套struct
type TextualClickTracker struct {
	TextualClickQuery
	IdfaMd5      string `json:"idfa_md5"`
	ImeiMd5      string `json:"imei_md5"`
	OaidMd5      string `json:"oaid_md5"`
	MacMd5       string `json:"mac_md5"`
	AndroididMd5 string `json:"androidid_md5"`
}

type TextualClickQuery struct {
	Aid       string `form:"aid" json:"aid" binding:"required"`
	Cid       string `form:"cid" json:"cid" binding:"required"`
	Os        int    `form:"os" json:"os" binding:"oneof=0 1 2 3"`
	Brand     string `form:"brand" json:"brand"`
	Model     string `form:"model" json:"model"`
	Osvs      string `form:"osvs" json:"osvs"`
	Idfa      string `form:"idfa" json:"idfa"`
	Androidid string `form:"androidid" json:"androidid"`
	Imei      string `form:"imei" json:"imei"`
	Oaid      string `form:"oaid" json:"oaid"`
	Uuid      string `form:"uuid" json:"uuid"`
	Mac       string `form:"mac" json:"mac"`
	Ua        string `form:"ua" json:"ua"`
	Ts        int64  `form:"ts" json:"ts" binding:"required"`
}

func TestStruct(t *testing.T) {
	query := TextualClickQuery{
		Aid: "aid",
		Cid: "cid",
	}

	var tracker TextualClickTracker
	tracker.Aid = query.Aid
	tracker.Cid = query.Aid
	tracker.ImeiMd5 = "3iejss"
	js, _ := json.Marshal(tracker)
	t.Logf("%s", js)
}

func (tracker *TextualClickTracker) init() {
	if len(tracker.Idfa) > 0 {
		tracker.IdfaMd5 = MD5(tracker.Idfa)
	}
}

//func TestJsonToStruct() {
//
//}
func MD5(value string) string {
	h := md5.New()
	_, _ = h.Write([]byte(value))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

type PersionA struct {
	Age  int
	Name string
}

func TestStructGO(t *testing.T) {
	e2 := new(PersionA)
	t.Log(e2)
}
