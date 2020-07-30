package main

import (
	"fmt"
	"time"
)

//数据结构
type Color struct {
	model   int32 //显示模式，默认 1，高亮, 0 - 8对应[默认，高亮，下划线，闪烁，反白，不可见]，
	bgcolor int32 //背景颜色，默认39，不设，40-47对应[红，绿，黄，蓝，紫，湛，白]
	fgcolor int32 //前景颜色，默认33，黄色，30-37对应[红，绿，黄，蓝，紫，湛，白]
}

type Bar struct {
	color   Color  //进度条颜色，默认值[高亮显示，无背景色，前景为黄色]
	current int64  //当前进度
	total   int64  //总的进度
	percent int64  //进度百分比
	symbol  string //进度条符号，默认为>
	ratebar string //进度条
}

//设置进度条方式1
func (bar *Bar) SetColorBar(start, total int64) {
	bar.color = Color{1, 39, 33}
	bar.current = start
	bar.total = total
	bar.percent = bar.getPercent()

	if bar.symbol == "" {
		bar.symbol = ">"
	}

	for i := 0; i < int(bar.percent); i += 1 {
		bar.ratebar += bar.symbol
	}
}

func (bar *Bar) getPercent() int64 {
	return int64(float64(bar.current) * 100 / float64(bar.total))
}

//设置进度条方式2，带符号参数symbol
func (bar *Bar) SetColorBarWithSymbol(start, total int64, symbol string) {
	bar.symbol = symbol
	bar.SetColorBar(start, total)
}

//修改并显示进度条
func (bar *Bar) ModifyAndShowBar(current int64) {
	bar.current = current
	percent := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != percent && bar.percent%2 == 0 {
		bar.ratebar += bar.symbol
	}

	cl := bar.color
	fmt.Printf("\r\033[%d;%d;%dm[%-50s]\033[0m %3d%%  %8d/%d", cl.model, cl.bgcolor, cl.fgcolor, bar.ratebar, bar.percent, bar.current, bar.total)
}

//添加换行符并打印退出信息
func (bar *Bar) Finish() {
	fmt.Println()
	fmt.Println("Completed!")
}

func main() {
	var bar Bar

	bar.SetColorBar(0, 100) //设置进度条
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond) //替换此处为业务代码
		bar.ModifyAndShowBar(int64(i))     //修改并显示进度条
	}

	bar.Finish()
}
