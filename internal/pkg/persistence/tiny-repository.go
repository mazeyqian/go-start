package persistence

import (
	"errors"
	"fmt"
	"log"

	"github.com/mazeyqian/go-gin-gee/internal/pkg/config"
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/tiny"
	"github.com/mazeyqian/go-gin-gee/pkg/helpers"
	"github.com/mazeyqian/gurl"
	"github.com/takuoki/clmconv"
)

type TinyRepository struct{}

var tinyRepository *TinyRepository

const cusConPrefix = "[Tiny]"

func GetTinyRepository() *TinyRepository {
	if tinyRepository == nil {
		tinyRepository = &TinyRepository{}
	}
	return tinyRepository
}

func (r *TinyRepository) SaveOriLink(OriLink string, addBaseUrl string, oneTime bool) (string, error) {
	var err error
	var tiny models.Tiny
	var linkForEncode string
	if addBaseUrl != "" {
		linkForEncode, err = gurl.SetHashParam(OriLink, "base_url", addBaseUrl)
		if err != nil {
			return "", err
		}
	} else {
		linkForEncode = OriLink
	}
	OriMd5 := helpers.ConvertStringToMD5Hash(linkForEncode)
	data, _ := r.QueryOriLinkByOriMd5(OriMd5)
	if data != nil {
		return data.TinyLink, nil
	}
	baseUrl := config.GetConfig().Data.BaseURL
	if addBaseUrl != "" {
		baseUrl = addBaseUrl
	}
	if baseUrl == "" {
		return "", errors.New("BASE_URL is required")
	}
	tiny.OriLink = OriLink
	tiny.OriMd5 = OriMd5
	err = Create(&tiny)
	if err != nil {
		return "", err
	}
	TinyId := tiny.ID
	// https://github.com/takuoki/clmconv
	converter := clmconv.New(clmconv.WithStartFromOne(), clmconv.WithLowercase())
	TinyKey := converter.Itoa(int(TinyId))
	TinyLink := fmt.Sprintf("%s/t/%s", baseUrl, TinyKey)
	// Compare
	specialLinks := config.GetConfig().Data.SpecialLinks
	if len(specialLinks) > 0 {
		for _, v := range specialLinks {
			if v.Key == TinyKey {
				log.Printf("%s Key(%s) is already in use", cusConPrefix, TinyKey)
				tiny.OriLink = v.Link
				tiny.OriMd5 = helpers.ConvertStringToMD5Hash(v.Link)
				tiny.TinyKey = TinyKey
				tiny.TinyLink = TinyLink
				err = Save(&tiny)
				if err != nil {
					return "", err
				}
				return r.SaveOriLink(OriLink, addBaseUrl, oneTime)
			}
		}
	}
	_, err = r.SaveTinyLink(TinyId, TinyLink, TinyKey, oneTime)
	if err != nil {
		return "", err
	}
	tiny.TinyKey = TinyKey
	tiny.TinyLink = TinyLink
	return tiny.TinyLink, err
}

func (r *TinyRepository) QueryOriLinkByTinyKey(TinyKey string) (string, error) {
	var tiny models.Tiny
	var err error
	specialLinks := config.GetConfig().Data.SpecialLinks
	if len(specialLinks) > 0 {
		for _, v := range specialLinks {
			if v.Key == TinyKey {
				log.Printf("%s Key(%s) is found in special links(%s)", cusConPrefix, TinyKey, v.Link)
				return v.Link, err
			}
		}
	}
	where := models.Tiny{}
	where.TinyKey = TinyKey
	notFound, err := First(&where, &tiny, []string{})
	log.Printf("%s Is this key NotFound in DB: %t", cusConPrefix, notFound)
	if err != nil {
		return "", errors.New("404 Link Not Found")
	}
	if tiny.OneTime && tiny.VisitCount > 0 {
		return "", errors.New("404 Link Expired")
	}
	go r.RecordVisitCountByTinyKey(TinyKey)
	return tiny.OriLink, err
}

func (r *TinyRepository) RecordVisitCountByTinyKey(TinyKey string) (bool, error) {
	var tiny models.Tiny
	var err error
	where := models.Tiny{}
	where.TinyKey = TinyKey
	notFound, err := First(&where, &tiny, []string{})
	if err != nil {
		return false, err
	}
	if notFound {
		return false, errors.New("link not found")
	}
	tiny.VisitCount = tiny.VisitCount + 1
	err = Updates(&where, &tiny)
	if err != nil {
		return false, err
	}
	log.Printf("%s Current Count: %d", cusConPrefix, tiny.VisitCount)
	return true, err
}

func (r *TinyRepository) QueryOriLinkByOriMd5(OriMd5 string) (*models.Tiny, error) {
	var tiny models.Tiny
	if OriMd5 == "" {
		return nil, errors.New("OriMd5 is required")
	}
	where := models.Tiny{}
	where.OriMd5 = OriMd5
	notFound, err := First(&where, &tiny, []string{})
	log.Printf("%s Is this link NotFound in DB: %t", cusConPrefix, notFound)
	if err != nil {
		return nil, err
	}
	return &tiny, err
}

func (r *TinyRepository) SaveTinyLink(TinyId uint64, TinyLink string, TinyKey string, oneTime bool) (bool, error) {
	var tiny models.Tiny
	var err error
	where := models.Tiny{}
	where.ID = TinyId
	tiny.TinyLink = TinyLink
	tiny.TinyKey = TinyKey
	tiny.OneTime = oneTime
	err = Updates(&where, &tiny)
	if err != nil {
		return false, err
	}
	return true, err
}
