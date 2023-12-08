package collector

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type TaiwanDailyPriceCollector struct {
	*colly.Collector
}

func NewTaiwanDailyPriceCollector(domain string) *TaiwanDailyPriceCollector {
	return &TaiwanDailyPriceCollector{
		Collector: colly.NewCollector(
			colly.AllowedDomains(domain),
		),
	}
}

type StockDailyValueObj struct {
	Date         time.Time // 日期
	Volume       int64     // 成交股數
	Amount       int64     // 成交金額
	ClosingPrice float64   // 收盤價
}

// GetMonthlyPrices 拿取台灣個股的整個月份的StockDailyValueObj
func (c *TaiwanDailyPriceCollector) GetMonthlyPrices(stockNo string, year, month uint) (data []StockDailyValueObj, err error) {
	var apiResp struct {
		Data [][]string `json:"data"`
	}
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "application/json")
	})

	c.OnResponse(func(r *colly.Response) {
		if err = json.Unmarshal(r.Body, &apiResp); err != nil {
			return
		}
		data = make([]StockDailyValueObj, len(apiResp.Data))
		for idx := range apiResp.Data {
			var (
				date         time.Time
				volume       int64
				amount       int64
				closingPrice float64
			)
			//apiResp.Data[idx][0] 日期: 112/12/01
			// 分割年、月、日
			parts := strings.Split(apiResp.Data[idx][0], "/")
			if len(parts) != 3 {
				err = errors.New("日期格式不正確")
				return
			}

			// 將民國年轉換為西元年
			var y int
			if y, err = strconv.Atoi(parts[0]); err != nil {
				return
			}
			// 民國年轉西元年
			y += 1911
			// 重組為西元年的日期字串
			adDate := fmt.Sprintf("%d-%s-%s", y, parts[1], parts[2])

			// 使用 time.Parse 解析日期
			if date, err = time.Parse("2006-01-02", adDate); err != nil {
				return
			}
			// 加載 UTC+8 時區（台灣時區）
			var location *time.Location
			if location, err = time.LoadLocation("Asia/Taipei"); err != nil {
				return
			}
			utc8Date := date.In(location)
			data[idx].Date = utc8Date

			//apiResp.Data[idx][1] 成交股數: 28,797,795
			volumeStr := strings.Replace(apiResp.Data[idx][1], ",", "", -1)
			if volume, err = strconv.ParseInt(volumeStr, 10, 64); err != nil {
				return
			}
			data[idx].Volume = volume

			//apiResp.Data[idx][2] 成交金額: 16,600,169,332
			amountStr := strings.Replace(apiResp.Data[idx][2], ",", "", -1)
			if amount, err = strconv.ParseInt(amountStr, 10, 64); err != nil {
				return
			}
			data[idx].Amount = amount

			//apiResp.Data[idx][6] 收盤價: 579.00
			if closingPrice, err = strconv.ParseFloat(apiResp.Data[idx][6], 64); err != nil {
				return
			}
			data[idx].ClosingPrice = closingPrice

		}
	})

	// 解析基礎 URL
	baseUrl, err := url.Parse("https://www.twse.com.tw/rwd/zh/afterTrading/STOCK_DAY")
	if err != nil {
		return nil, fmt.Errorf("[TaiwanDailyPriceCollector] url.Parse err: %w", err)
	}
	// 創建並設置查詢參數
	params := url.Values{}
	params.Add("date", fmt.Sprintf("%d%d01", year, month))
	params.Add("stockNo", stockNo)
	params.Add("response", "json")

	// 將查詢參數添加到 URL
	baseUrl.RawQuery = params.Encode()

	if err = c.Visit(baseUrl.String()); err != nil {
		return nil, fmt.Errorf("[TaiwanDailyPriceCollector] c.Visit err: %w", err)
	}

	return data, nil
}
