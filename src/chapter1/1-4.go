package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

//通过image的package生成bit-mapped图，并编码为GIF动图（利萨如图形）

//复合声明一个切片
var palette = []color.Color{color.White, color.Black}

//声明一个数组
//var palette = [2]color.Color{color.White, color.Black}

//包级别常量体
//声明常量方式1
//const whiteIndex = 0
//const blackIndex = 1
//声明常量方式2
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

//制作利萨如图形
func lissajous(out io.Writer) {
	//函数级别常量体
	//定义常量
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	//定义变量
	var (
		test1 = ""
		test2 = ""
		test3 = ""
	)
	fmt.Print(test1, test2, test3)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	//复合声明GIF
	//生成一个struct变量（类似对象）
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	//短声明 不声明类型
	//外层循环
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		//内存玄幻
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			//sin函数
			x := math.Sin(t)
			//sin函数
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		//通过.访问结构体属性
		//通过append添加至末尾
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//GIF结构体
type GIF struct {
	Image []*image.Paletted
	Delay []int
	LoopCount int
	Disposal []byte
	Config image.Config
	BackgroundIndex byte
}
