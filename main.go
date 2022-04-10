package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	_ "image/png"
	"io"
	"log"
	"os"
)

// Основные переменные глобального уровня
const (
	width  = 1080
	height = 720
)

var (
	img     *ebiten.Image
	X, Y    = 540.0, 360.0
	speedX  = 3.0
	speedY  = 3.0
	vectorX = speedX
	vectorY = speedY
	size    = 100
)

// Функция для отгрузки файлов
func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("texture.png")
	if err != nil {
		log.Fatal(err)
	}
}

// Музыка
func run_sound() {
	f, _ := os.Open("Leonz_-_Among_Us_Trap_Remix_Among_Drip_Theme_Song_Original_72243941.mp3")
	d, _ := mp3.NewDecoder(f)
	c, _ := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	p := c.NewPlayer()
	_, l := io.Copy(p, d)
	_ = l
}

// Логика игры
func logic() {
	X = X + vectorX
	Y = Y + vectorY
	if int(X) <= 0 {
		vectorX = -vectorX
	}
	if int(Y) <= 0 {
		vectorY = -vectorY
	}
	if int(X) >= width-size {
		vectorX = -vectorX
	}
	if int(Y) >= height-size {
		vectorY = -vectorY
	}
}

type Game struct{}

func (g *Game) Update() error {
	logic()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Вывод ФПС
	fps := ebiten.CurrentFPS()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", int(fps)))
	// Даем координаты и выводим картинку
	src := &ebiten.DrawImageOptions{}
	src.GeoM.Translate(X, Y)
	screen.DrawImage(img, src)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("PizzusDVD")

	go run_sound()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
