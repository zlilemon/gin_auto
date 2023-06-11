package store

import (
	"github.com/gin-gonic/gin"
	"github.com/zlilemon/gin_auto/pkg/comm"
	"github.com/zlilemon/gin_auto/pkg/log"
	"strings"
)

type IService interface {
	SGetStoreInfoList(c *gin.Context) (err error)
}

type Service struct {
	repo IRepository
}

var StoreService = new(Service)

func (s *Service) SGetStoreInfoList(c *gin.Context, req GetStoreInfoListReq, resp *comm.Result) (err error) {
	log.Infof("SGetStoreInfoList - ")

	//var storeInfoModel []*StoreInfo
	storeInfoModel, err := StoreRepository.GetStoreInfoList(c, req.City)
	if err != nil {
		log.Errorf("failed to SGetStoreInfoList - %+v", err)
		return
	}

	storeRespList := make([]GetStoreInfoListResp, 0)
	for _, v := range storeInfoModel {
		newItem := GetStoreInfoListResp{}
		newItem.ShopId = v.StoreId
		newItem.ShopName = v.StoreName
		newItem.Location = v.Address
		newItem.Longitude = v.Longitude
		newItem.Latitude = v.Latitude
		newItem.ShopPic = v.HomePicUrl

		log.Infof("shop_id : %s", newItem.ShopId)
		storeRespList = append(storeRespList, newItem)
	}

	resp.Data = storeRespList

	log.Infof("SGetStoreInfoList done")

	return
}

func (s *Service) SGetStoreInfo(c *gin.Context, req StoreDetailReq, resp *comm.Result) (err error) {
	log.Infof("SGetStoreInfo - ")

	shopDetailResp := StoreDetailResp{}
	//seatOrderDetail := make([]SeatInfo, 0)
	//picDetail := make([]string, 0)
	//detailItem := StoreDetailResp{}

	// 查询店铺具体信息
	storeInfoModel, err := StoreRepository.GetStoreInfo(c, req.StoreId)
	if err != nil {
		log.Errorf("failed to SGetStoreInfo - %+v", err)
		return
	}

	v := storeInfoModel[0]
	shopDetailResp.ShopId = v.StoreId
	shopDetailResp.ShopName = v.StoreName
	shopDetailResp.Location = v.Address
	shopDetailResp.Longitude = v.Longitude
	shopDetailResp.Latitude = v.Latitude
	shopDetailResp.ShopPic = strings.Split(v.HomePicUrl, ";")
	shopDetailResp.ShopDesc = v.Introduction
	//店铺详细图片，列表
	picDetail := strings.Split(v.IntroductionPicUrl, ";")
	shopDetailResp.ShopDescPics = picDetail
	shopDetailResp.ShopSeatPics = v.SeatPicUrl

	/*
		detailItem.ShopId = newItem.ShopId
		detailItem.ShopName = newItem.ShopName
		detailItem.Location = newItem.Location
		detailItem.Longitude = newItem.Longitude
		detailItem.Latitude = newItem.Latitude
		detailItem.ShopPic = newItem.ShopPic
		detailItem.ShopDescPics = newItem.ShopDescPics
		detailItem.ShopSeatPics = newItem.ShopSeatPics
	*/
	// 获取该店铺的座位信息
	seatInfoModel, err := StoreRepository.GetSeatInfoByStore(c, req.StoreId)
	if err != nil {
		log.Errorf("failed to GetSeatInfoByStore - %+v", err)
		return
	}
	//storeDetailResp := make([]StoreDetailResp, 0)
	seatInfoTotal := SeatInfo{}
	normalSeat := make([]SeatInfoItem, 0)
	vipSeat := make([]SeatInfoItem, 0)
	doubleSeat := make([]SeatInfoItem, 0)

	for _, v := range seatInfoModel {
		// 根据店铺id 和座位id，获取对应的预定信息
		storeID := v.StoreId
		seatID := v.SeatId
		seatName := v.SeatName
		seatType := v.SeatType

		newItem := SeatInfoItem{}
		bookTimeList := make([]BookTime, 0)

		newItem.SeatId = seatID
		newItem.SeatName = seatName

		bookOrderStatusModel, err := StoreRepository.GetBookOrderStatusBySeatId(c, storeID, seatID)
		if err != nil {
			log.Errorf("failed to GetBookOrderStatusBySeatId - %+v", err)
			return err
		}
		for _, bv := range bookOrderStatusModel {
			//seatInfoResp := make([]SeatInfo, 0)

			bookTime := BookTime{}

			bookTime.BookBeginTime = bv.BookBeginTime
			bookTime.BookEndTime = bv.BookEndTime

			bookTimeList = append(bookTimeList, bookTime)
		}

		newItem.BookTimeList = bookTimeList
		if seatType == "normal" {
			normalSeat = append(normalSeat, newItem)
		} else if seatType == "vip" {
			vipSeat = append(vipSeat, newItem)
		} else if seatType == "double" {
			doubleSeat = append(doubleSeat, newItem)
		}
	}
	seatInfoTotal.Normal = normalSeat
	seatInfoTotal.Vip = vipSeat
	seatInfoTotal.Double = doubleSeat

	//seatOrderDetail = append(seatOrderDetail, seatInfoTotal)

	shopDetailResp.ShopSeats = seatInfoTotal

	// 获取店铺座位的价格
	priceInfo := PriceInfo{}
	normalPriceItem := PriceItem{}
	vipPriceItem := PriceItem{}
	doublePriceItem := PriceItem{}

	priceModel, err := StoreRepository.GetStorePrice(c, req.StoreId)
	if err != nil {
		log.Errorf("failed to GetStorePrice - %+v", err)
		return
	}
	for _, v := range priceModel {
		seatType := v.SeatType
		priceType := v.PriceType
		if seatType == "normal" {
			if priceType == "minute" {
				normalPriceItem.Minute.OrigPrice = v.OriPrice
				normalPriceItem.Minute.RealPrice = v.RealPrice
			} else if priceType == "day" {
				normalPriceItem.Day.OrigPrice = v.OriPrice
				normalPriceItem.Day.RealPrice = v.RealPrice
			} else if priceType == "month" {
				normalPriceItem.Month.OrigPrice = v.OriPrice
				normalPriceItem.Month.RealPrice = v.RealPrice
			}
		} else if seatType == "vip" {
			if priceType == "minute" {
				vipPriceItem.Minute.OrigPrice = v.OriPrice
				vipPriceItem.Minute.RealPrice = v.RealPrice
			} else if priceType == "day" {
				vipPriceItem.Day.OrigPrice = v.OriPrice
				vipPriceItem.Day.RealPrice = v.RealPrice
			} else if priceType == "month" {
				vipPriceItem.Month.OrigPrice = v.OriPrice
				vipPriceItem.Month.RealPrice = v.RealPrice
			}
		} else if seatType == "double" {
			if priceType == "minute" {
				doublePriceItem.Minute.OrigPrice = v.OriPrice
				doublePriceItem.Minute.RealPrice = v.RealPrice
			} else if priceType == "day" {
				doublePriceItem.Day.OrigPrice = v.OriPrice
				doublePriceItem.Day.RealPrice = v.RealPrice
			} else if priceType == "month" {
				doublePriceItem.Month.OrigPrice = v.OriPrice
				doublePriceItem.Month.RealPrice = v.RealPrice
			}
		}
	}
	priceInfo.Normal = normalPriceItem
	priceInfo.Vip = vipPriceItem
	priceInfo.Double = doublePriceItem

	shopDetailResp.Price = priceInfo

	resp.Data = shopDetailResp
	log.Infof("SGetStoreInfo done")

	return
}

func (s *Service) SGetStoreNotice(c *gin.Context, resp *comm.Result) (err error) {
	log.Infof("SGetStoreNotice - ")

	//storeNoticeModel := StoreNoticeModel{}
	//seatOrderDetail := make([]SeatInfo, 0)
	//picDetail := make([]string, 0)
	//detailItem := StoreDetailResp{}

	// 查询店铺公告信息
	storeNoticeModel, err := StoreRepository.GetStoreNotice(c)
	if err != nil {
		log.Errorf("failed to SGetStoreInfo - %+v", err)
		return
	}

	// 查询常见问题
	questionModel, err := StoreRepository.GetStoreQuestion(c)
	if err != nil {
		log.Errorf("failed to SGetStoreInfo - %+v", err)
		return
	}

	storeNoticeResp := StoreNoticeResp{}
	storeNoticeResp.WifiName = storeNoticeModel[0].WifiName
	storeNoticeResp.WifiPasswd = storeNoticeModel[0].WifiPasswd
	storeNoticeResp.CustomerPhone = storeNoticeModel[0].CustomerPhone
	storeNoticeResp.QuestionList = questionModel

	resp.Data = storeNoticeResp

	log.Infof("SGetStoreNotice Success - ")
	return
}

func (s *Service) SGetBulletinList(c *gin.Context, resp *comm.Result) (err error) {
	log.Infof("SGetBulletinList - ")

	//storeNoticeModel := StoreNoticeModel{}
	//seatOrderDetail := make([]SeatInfo, 0)
	//picDetail := make([]string, 0)
	//detailItem := StoreDetailResp{}

	// 查询店铺公告信息
	bulletinSummaryModel, err := StoreRepository.GetBulletinList(c)
	if err != nil {
		log.Errorf("failed to SGetBulletinList - %+v", err)
		return
	}

	bulletinSummaryResp := BulletinSummaryResp{}
	bulletinSummaryResp.BulletinList = bulletinSummaryModel

	resp.Data = bulletinSummaryResp

	log.Infof("SGetBulletinList Success - ")
	return
}

func (s *Service) SGetBulletinDetail(c *gin.Context, req BulletinDetailReq, resp *comm.Result) (err error) {
	log.Infof("SGetBulletinDetail - ")

	//storeNoticeModel := StoreNoticeModel{}
	//seatOrderDetail := make([]SeatInfo, 0)
	//picDetail := make([]string, 0)
	//detailItem := StoreDetailResp{}

	// 查询店铺公告信息
	bulletinDetailModel, err := StoreRepository.GetBulletinDetail(c, req)
	if err != nil {
		log.Errorf("failed to SGetBulletinDetail - %+v", err)
		return
	}

	bulletinDetailResp := BulletinDetailResp{}
	bulletinDetailResp.Id = bulletinDetailModel[0].Id
	bulletinDetailResp.Title = bulletinDetailModel[0].Title
	bulletinDetailResp.PublishTimeUnixTime = bulletinDetailModel[0].PublishTimeUnixTime
	bulletinDetailResp.PublishDetail = bulletinDetailModel[0].PublishDetail
	bulletinDetailResp.Status = bulletinDetailModel[0].Status

	resp.Data = bulletinDetailResp

	log.Infof("SGetBulletinDetail Success - ")
	return
}
