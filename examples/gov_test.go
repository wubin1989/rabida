package examples

import (
	"context"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	. "github.com/JohnnyTing/rabida"
	"github.com/JohnnyTing/rabida/config"
	"github.com/chromedp/chromedp"
	"testing"
)

func TestGovCrawl(t *testing.T) {
	t.Run("sousou.gov", func(t *testing.T) {
		conf := config.LoadFromEnv()
		fmt.Printf("%+v\n", conf)
		rabi := NewRabida(conf)
		job := Job{
			Link: "http://sousuo.gov.cn/s.htm?t=zhengcelibrary&q=&timetype=&mintime=&maxtime=&sort=&sortType=&searchfield=&pcodeJiguan=&bmfl=&childtype=&subchildtype=&tsbq=&pubtimeyear=&puborg=&pcodeYear=&pcodeNum=&filetype=&p=&n=&orpro=&inpro=",
			CssSelector: CssSelector{
				Scope: `div.middle_result_con > div.dys_middle_result_content_item`,
				Attrs: map[string]CssSelector{
					"link": {
						Css:  "a",
						Attr: "href",
					},
					"title": {
						Css: "h5",
					},
					"date": {
						Css: ".dysMiddleResultConItemRelevant>span:last-child",
					},
				},
			},
		}
		err := rabi.Crawl(context.Background(), job, func(ret []interface{}, nextPageUrl string, currentPageNo int) bool {
			for _, item := range ret {
				fmt.Println(gabs.Wrap(item).StringIndent("", "  "))
			}
			if currentPageNo >= job.Limit {
				return true
			}
			return false
		}, nil, []chromedp.Action{
			chromedp.EmulateViewport(1777, 903, chromedp.EmulateLandscape),
		})
		if err != nil {
			t.Errorf("%+v", err)
		}

	})
}
