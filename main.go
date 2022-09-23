package main

import (
	"fmt"
	skip2 "github.com/skip2/go-qrcode"
	yeqown "github.com/yeqown/go-qrcode"
	yeqown2 "github.com/yeqown/go-qrcode/v2"
	yeqown2compressed "github.com/yeqown/go-qrcode/writer/compressed"
	yeqown2std "github.com/yeqown/go-qrcode/writer/standard"
	"image"
	"image/png"
	"io"
	"os"
	"strings"
)

func main() {
	// some json
	content := `{"first_name":"John","last_name":"Doe","field1":"RpilI4cbWXIPhfh50dzj9TeshBQT2--TGycXsXc34MY79g","field2":"4fRIl8NA3VzEtM7LzpBNiNHFBKd4ccHV4NkRksKqsRcAw","field3":"Yy-Q5rJpIq-v2Kow-NHTkssPz_Y7ykSFPjfUGT_EWgOB","field4":"KW223jlP9RssPBtRdXGPpiWiBJHbnzcJR_qSKA","field5":true}`
	fmt.Printf("Content length: %d\n", len(content))
	encodeWithYeqownCompression(content, "qr-yeqown-none.png", png.NoCompression)
	encodeWithYeqownCompression(content, "qr-yeqown-speed.png", png.BestSpeed)
	encodeWithYeqown(content, "qr-yeqown.png")
	encodeWithYeqownCompression(content, "qr-yeqown-best.png", png.BestCompression)
	encodeWithYeqownCurrent(content, "qr-yeqown-current.jpeg")
	encodeWithYeqownNewJpg(content, "qr-yeqown-v2.jpg")
	encodeWithYeqownNewPng(content, "qr-yeqown-v2.png")
	encodeWithYeqownNewPngCompressed(content, "qr-yeqown-v2-compressed.png")
	encodeWithYeqownNewPngCompressed1Px(content, "qr-yeqown-v2-compressed-1px.png")
	encodeWithSkip2(content, "qr-skip2-best-1px.png")
	encodeWithSkip2(jsonToAlphaNum(content), "qr-skip2-best-1px-alphanum.png")
}

// strip json to alpha num just for a test
func jsonToAlphaNum(content string) string {
	content = strings.Replace(content, "{", "$", -1)
	content = strings.Replace(content, "}", "%", -1)
	content = strings.Replace(content, "\n", " ", -1)
	content = strings.Replace(content, "\"", "*", -1)
	content = strings.Replace(content, ",", "/", -1)
	content = strings.Replace(content, "[", "/", -1)
	content = strings.Replace(content, "]", "/", -1)
	content = strings.ToUpper(content)
	return content
}

func encodeWithSkip2(content, name string) {
	err := skip2.WriteFile(content, skip2.High, 0, name)
	if err != nil {
		panic(err)
	}
	printFileSize(name)
}

type CustomPngEncoder struct {
	CompressionLevel png.CompressionLevel
}

func (j CustomPngEncoder) Encode(w io.Writer, img image.Image) error {
	pngEncoder := png.Encoder{CompressionLevel: j.CompressionLevel}
	return pngEncoder.Encode(w, img)
}

// Use default PNG compression level
func encodeWithYeqown(content, name string) {
	cfg := yeqown.DefaultConfig()
	cfg.EcLevel = yeqown.ErrorCorrectionQuart

	imgOpts := yeqown.WithBuiltinImageEncoder(yeqown.PNG_FORMAT)
	imgOpts2 := yeqown.WithQRWidth(4)
	imgOpts3 := yeqown.WithBorderWidth(4)
	qrc, err := yeqown.NewWithConfig(content, cfg, imgOpts, imgOpts2, imgOpts3)
	if err != nil {
		panic(err)
	}
	// save file
	if err := qrc.Save(name); err != nil {
		panic(err)
	}
	printFileSize(name)
}

func encodeWithYeqownCompression(content, name string, pngCompressionLevel png.CompressionLevel) {
	cfg := yeqown.DefaultConfig()
	cfg.EcLevel = yeqown.ErrorCorrectionQuart

	imageEncoder := CustomPngEncoder{pngCompressionLevel}
	imgOpts := yeqown.WithCustomImageEncoder(imageEncoder)
	imgOpts2 := yeqown.WithQRWidth(4)
	imgOpts3 := yeqown.WithBorderWidth(4)
	qrc, err := yeqown.NewWithConfig(content, cfg, imgOpts, imgOpts2, imgOpts3)
	if err != nil {
		panic(err)
	}
	// save file
	if err := qrc.Save(name); err != nil {
		panic(err)
	}
	printFileSize(name)
}

func encodeWithYeqownCurrent(content, name string) {
	cfg := yeqown.DefaultConfig()
	cfg.EcLevel = yeqown.ErrorCorrectionQuart

	imgOpts2 := yeqown.WithQRWidth(4)
	qrc, err := yeqown.NewWithConfig(content, cfg, imgOpts2)
	if err != nil {
		panic(err)
	}
	// save file
	if err := qrc.Save(name); err != nil {
		panic(err)
	}
	printFileSize(name)
}

func encodeWithYeqownNewJpg(content, name string) {
	qrc, err := yeqown2.New(content)
	if err != nil {
		panic(err)
	}
	imgOpt := yeqown2std.WithBuiltinImageEncoder(yeqown2std.JPEG_FORMAT)
	imgOpt2 := yeqown2std.WithQRWidth(4)
	imgOpt3 := yeqown2std.WithBorderWidth(4)
	w, err := yeqown2std.New(name, imgOpt, imgOpt2, imgOpt3)
	if err != nil {
		panic(err)
	}
	// save file
	if err := qrc.Save(w); err != nil {
		panic(err)
	}
	printFileSize(name)
}

func encodeWithYeqownNewPng(content, name string) {
	errCorrLvl := yeqown2.WithErrorCorrectionLevel(yeqown2.ErrorCorrectionQuart)
	qrc, err := yeqown2.NewWith(content, errCorrLvl)
	if err != nil {
		panic(err)
	}
	imgOpt := yeqown2std.WithBuiltinImageEncoder(yeqown2std.PNG_FORMAT)
	imgOpt2 := yeqown2std.WithQRWidth(4)
	imgOpt3 := yeqown2std.WithBorderWidth(4)
	w, err := yeqown2std.New(name, imgOpt, imgOpt2, imgOpt3)
	if err != nil {
		panic(err)
	}
	// save file
	if err := qrc.Save(w); err != nil {
		panic(err)
	}
	printFileSize(name)
}

func encodeWithYeqownNewPngCompressed(content, name string) {
	errCorrLvl := yeqown2.WithErrorCorrectionLevel(yeqown2.ErrorCorrectionQuart)
	qrc, err := yeqown2.NewWith(content, errCorrLvl)
	if err != nil {
		panic(err)
	}
	option := &yeqown2compressed.Option{
		Padding:   4,
		BlockSize: 4,
	}
	w, err := yeqown2compressed.New(name, option)
	if err != nil {
		panic(err)
	}
	// save file
	if err := qrc.Save(w); err != nil {
		panic(err)
	}
	printFileSize(name)
}

func encodeWithYeqownNewPngCompressed1Px(content, name string) {
	errCorrLvl := yeqown2.WithErrorCorrectionLevel(yeqown2.ErrorCorrectionQuart)
	qrc, err := yeqown2.NewWith(content, errCorrLvl)
	if err != nil {
		panic(err)
	}
	option := &yeqown2compressed.Option{
		Padding:   4,
		BlockSize: 1,
	}
	w, err := yeqown2compressed.New(name, option)
	if err != nil {
		panic(err)
	}
	// save file
	if err := qrc.Save(w); err != nil {
		panic(err)
	}
	printFileSize(name)
}

func printFileSize(name string) {
	stat, _ := os.Stat(name)
	fmt.Printf("%s: %v\n", name, stat.Size())
}
