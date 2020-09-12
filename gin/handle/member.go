package handle

import (
	"github.com/GetStartWithGoLang/gin/models"
)

type MemberData struct {
	data []models.Member
}

func NewMember() *MemberData {
	return &MemberData{}
}

func (m *MemberData) AllData() []models.Member {
	return m.data
}

// func (m *MemberData)
