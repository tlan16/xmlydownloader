package xmlydownloader

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetAlbumInfo(t *testing.T) {
	albumID := 2780581 //https://www.ximalaya.com/youshengshu/2780581/
	ai, err := GetAlbumInfo(albumID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ai)
	assert.NotNil(t, ai)
	assert.Equal(t, "0", ai.Msg, "Expect Msg to be string 0")
	assert.Zero(t, ai.Ret, "Expect Reg to be number 0")
	assert.Equal(t, true, ai.CostTimeAlbumInfo > 0, "Expect CostTimeAlbumInfo to be a positive number")
	assert.Zero(t, ai.Ret, "Expect Reg to be number 0")
	// assert album
	assert.True(t, ai.Data.Album.AlbumID > 1, "Expect a valid Data.Album.AlbumID")
	assert.True(t, len(ai.Data.Album.Title) > 0, "Expect Data.Album.Title to be an non-empty string")
	assert.True(t, ai.Data.Album.IsPaid, "Expect paid album to be detected correctly indicated via Data.Album.IsPaid")
	assert.True(t, ai.Data.Album.TrackCount > 0, "Expect Data.Album.TrackCount to be a positive number")
	assert.Zero(t, ai.Data.Album.VipFreeType, "Expect vip free type to be detected correctly indicated via Data.Album.vipFreeType")
	assert.True(t, len(ai.Data.Album.PriceTypes) > 0, "Expect Data.Album.PriceTypes to be an non-empty array")
	for _, PriceType := range ai.Data.Album.PriceTypes {
		assert.True(t, len(PriceType.FreeTrackIds) > 0, "Expect Data.Album.PriceTypes.FreeTrackIds to be an non-empty string")
		assert.True(t, PriceType.FreeTrackCount > 0, "Expect Data.Album.PriceTypes.FreeTrackCount to be an positive number")
	}
	assert.Zero(t, ai.Data.Album.IsFinished, "Expect is album finished to be detected correctly indicated via Data.Album.IsFinished")
}

func TestGetAllTrackList(t *testing.T) {
	var trackInfoList []*TrackInfo
	albumID := 2780581 //https://www.ximalaya.com/youshengshu/2780581/
	tracks, err := GetTrackList(albumID, 1, false)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, tracks)
	assert.Equal(t, "0", tracks.Msg, "Expect Msg to be string 0")
	assert.Zero(t, tracks.Ret, "Expect Reg to be number 0")
	for _, v := range tracks.Data.List {
		trackInfoList = append(trackInfoList, v)
	}

	for i := 2; i <= tracks.Data.MaxPageID; i++ {
		tracks, err = GetTrackList(albumID, i, false)
		if err != nil {
			t.Fatal(err)
		}

		for _, v := range tracks.Data.List {
			trackInfoList = append(trackInfoList, v)
		}
	}
	for i, v := range trackInfoList {
		t.Log(i, v.Title)
		assert.True(t, v.TrackID > 0)
		assert.True(t, v.TrackRecordID > 0)
		assert.True(t, v.UID > 0)
		assert.Equal(t, v.IsFree, len(v.PlayURL64) > 0)
		assert.Equal(t, v.IsFree, len(v.PlayURL32) > 0)
		assert.Equal(t, v.IsFree, len(v.PlayPathAacv164) > 0)
		assert.Equal(t, v.IsFree, len(v.PlayPathAacv224) > 0)
		assert.True(t, v.Duration > 0)
		assert.True(t, v.AlbumID > 0)
		assert.True(t, v.IsPaid)
		assert.NotNil(t, v.IsFree)
		assert.False(t, v.IsVideo)
		assert.False(t, v.IsDraft)
		assert.False(t, v.IsRichAudio)
		assert.False(t, v.IsAuthorized)
		assert.True(t, v.Price > 0)
		assert.True(t, v.DiscountedPrice > 0)
		assert.True(t, v.PriceTypeID > 0)
		assert.True(t, v.SampleDuration > 0)
		assert.True(t, v.PriceTypeID > 0)
		assert.True(t, len(v.DisplayPrice) > 0)
		assert.True(t, len(v.DisplayDiscountedPrice) > 0)
		assert.True(t, v.Type >= 0)
		assert.True(t, v.RelatedID >= 0)
		assert.True(t, v.OrderNo > 0)
		assert.True(t, v.IsHoldCopyright)
		assert.Equal(t, 0, v.VipFirstStatus)
		assert.Equal(t, 0, v.PaidType)
		assert.False(t, v.IsSample)
		assert.True(t, v.ProcessState > 0)
		assert.True(t, v.CreatedAt > 0)
		assert.True(t, len(v.CoverSmall) > 0)
		assert.True(t, len(v.CoverMiddle) > 0)
		assert.True(t, len(v.CoverLarge) > 0)
		assert.True(t, len(v.Nickname) > 0)
		assert.True(t, len(v.SmallLogo) > 0)
		assert.True(t, v.UserSource > 0)
		assert.True(t, v.OpType > 0)
		assert.True(t, v.IsPublic)
		assert.True(t, v.Likes > 0)
		assert.True(t, v.Playtimes >= 0)
		assert.True(t, v.Comments >= 0)
		assert.True(t, v.Shares >= 0)
		assert.True(t, v.Status >= 0)
		assert.False(t, v.IsTrailer)
	}
}

func TestQRCode(t *testing.T) {
	qrCode, err := GetQRCode()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(qrCode.Img)
	log.Println(qrCode.QrID)
	//
	status, cookie, err := CheckQRCodeStatus("FEDECD84A3014713B396B4B6ED4F3483")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status, cookie)
}
