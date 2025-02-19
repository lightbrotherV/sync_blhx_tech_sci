package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync_azur_lane/mysql"
	"time"
)

func main() {
	res, err := http.Get("https://wiki.biligame.com/blhx/%E8%88%B0%E9%98%9F%E7%A7%91%E6%8A%80")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Panicf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panic(err)
	}
	// 数据库初始化
	mysql.Linksql()
	// 查询列表
	var shipArr []*mysql.AzurLane
	now := time.Now().Format(time.DateTime)
	doc.Find("#CardSelectTr tr").Each(func(i int, tr *goquery.Selection) {
		if i < 2 {
			return
		}
		alRecord := &mysql.AzurLane{
			CreatedAt:   now,
			UpdatedAt:   now,
			CreatedById: 1,
			UpdatedById: 1,
		}
		alRecord.Code = removeWrap(tr.Find("td").Eq(0).Text())
		alRecord.Name = removeWrap(tr.Find("td").Eq(1).Find("a span").Eq(0).Text())
		alRecord.Camp = removeWrap(tr.Find("td").Eq(2).Text())
		alRecord.ShipType = removeWrap(tr.Find("td").Eq(3).Text())
		alRecord.TechPointGet, err = strconv.ParseInt(removeWrap(tr.Find("td").Eq(4).Text()), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		alRecord.TechPointStar, err = strconv.ParseInt(removeWrap(tr.Find("td").Eq(5).Text()), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		alRecord.TechPointLv120, err = strconv.ParseInt(removeWrap(tr.Find("td").Eq(6).Text()), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		alRecord.TechPointTotal, err = strconv.ParseInt(removeWrap(tr.Find("td").Eq(7).Text()), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		getAttStr := removeWrap(tr.Find("td").Eq(8).Text())
		// 切分获取舰船属性加成 字符串
		getAttArr := strings.Split(getAttStr, " ")
		alRecord.AttributeGetApplyShip = getAttArr[0]
		getAttDetail := strings.Split(getAttArr[1], "+")
		alRecord.AttributeNameGet = getAttDetail[0]
		alRecord.AttributeGet, err = strconv.ParseInt(getAttDetail[1], 10, 64)
		if err != nil {
			log.Panic(err)
		}

		lv120AttStr := removeWrap(tr.Find("td").Eq(9).Text())
		// 切分获取舰船属性加成 字符串
		lv120AttArr := strings.Split(lv120AttStr, " ")
		alRecord.AttributeLv120ApplyShip = lv120AttArr[0]
		getAttDetail = strings.Split(lv120AttArr[1], "+")
		alRecord.AttributeNameLv120 = getAttDetail[0]
		alRecord.AttributeLv120, err = strconv.ParseInt(getAttDetail[1], 10, 64)
		shipArr = append(shipArr, alRecord)
	})

	// 开始写入mysql
	// 存在就不管
	for _, row := range shipArr {
		tmpRecord, tmpErr := mysql.GetRecordByCode(row.Code)
		if tmpErr != nil {
			log.Panic(tmpErr)
		}
		// 不存在，就新建
		if tmpRecord.ID == 0 {
			tmpErr = mysql.InsertRecordByCode(row)
			if tmpErr != nil {
				log.Panic(tmpErr)
			}
		}
		// 科技点数据还没更新
		if tmpRecord.TechPointTotal == 0 && row.TechPointTotal == 0 {
			continue
		}

		// 更新科技点数据
		if tmpRecord.TechPointTotal == 0 && row.TechPointTotal > 0 {
			tmpErr = mysql.UpdateRecordByCode(row)
			if tmpErr != nil {
				log.Panic(tmpErr)
			}
		}
	}
}

func removeWrap(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}
