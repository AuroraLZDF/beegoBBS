package utils

import (
	"fmt"
	"math"
)

/**
* 分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
* page	当前页
* prepage	每页显示条数
* num 	总数据条数
* currentPath	url
 */
func Paginator(page, prepage int, nums int64, currentPath string) string {
	fmt.Println("page:", page, " | prepage:", prepage, " | num", nums, " currentPath | ", currentPath)
	var firstpage int //前一页地址
	var lastpage int  //后一页地址

	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}

	if page <= 0 {
		page = 1
	}

	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		fmt.Println(pages)
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	}

	str := Default(pages, firstpage, lastpage, page, currentPath)

	return str
}

/**
* 分页样式
* pages	总页数
* firstpage	首页
* lastpage	最后一页
* currentPage	当前页
* url	URL
 */
func Default(pages []int, firstpage int, lastpage int, currentPage int, url string) string {
	var html = "<nav aria-label='Page navigation'><ul class='pagination'>"
	if currentPage == firstpage {
		html = html + "<li class='disabled'><span aria-hidden='true'>&laquo;</span></li>"
	} else {
		html = html + "<li><a href='" + url + "?page=" + IntToString(firstpage) + "' aria-label='Previous'><span aria-hidden='true'>&laquo;</span></a></li>"
	}

	for _, page := range pages {
		if currentPage == page {
			html = html + "<li class='active'><span>" + IntToString(page) + "</span></li>"
		} else {
			html = html + "<li><a href='" + url + "?page=" + IntToString(page) + "'" + ">" + IntToString(page) + "</a></li>"
		}
	}

	if currentPage >= lastpage {
		html = html + "<li class='disabled'><span aria-hidden='true'>&raquo;</span></li>"
	} else {
		html = html + "<li><a href='" + url + "?page=" + IntToString(lastpage) + "' aria-label='Next'><span aria-hidden='true'>&raquo;</span></a></li>"
	}

	html = html + "</ul></nav>"

	return html
}
