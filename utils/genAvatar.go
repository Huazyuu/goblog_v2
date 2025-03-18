package utils

import (
	"fmt"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"unicode/utf8"
)

func GenerateNameAvatar() {
	dir := "uploads/chat_avatar"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(filepath.Clean(dir), 0755)
		fmt.Printf("目录 %s 已创建\n", dir)
	}
	fmt.Printf("目录 %s 已存在\n", dir)
	for _, s := range AnimeAttribute {
		DrawImage(string([]rune(s)[0]), dir)
	}
	for _, s := range AnimeCharacter {
		DrawImage(string([]rune(s)[0]), dir)
	}
}

func DrawImage(name string, dir string) {
	fontFile, err := os.ReadFile("uploads/font/STHUPO.TTF")
	font, err := freetype.ParseFont(fontFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	options := &letteravatar.Options{
		Font: font,
	}
	// 绘制文字
	firstLetter, _ := utf8.DecodeRuneInString(name)
	img, err := letteravatar.Draw(140, firstLetter, options)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 保存
	filePath := path.Join(dir, name+".png")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err)
		return
	}
}
