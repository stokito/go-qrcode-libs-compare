# Golang compare QR code generation libraries

Compare size of generated QR codes of three libraries:
* github.com/skip2/go-qrcode 
* github.com/yeqown/go-qrcode v1
* github.com/yeqown/go-qrcode/v2
* *TBD* https://github.com/makiuchi-d/gozxing

Results:
```
Content length: 274
qr-yeqown-none.png: 300085
qr-yeqown-speed.png: 5459
qr-yeqown.png: 3820
qr-yeqown-best.png: 3524
qr-yeqown-current.jpeg: 37153
qr-yeqown-v2.jpg: 35978
qr-yeqown-v2.png: 3846
qr-yeqown-v2-compressed.png: 1491
qr-yeqown-v2-compressed-1px.png: 1000
qr-skip2-best-1px.png: 1003
qr-skip2-best-1px-alphanum.png: 849
```

The `qr-skip2-best-1px.png` image is exactly the same as `qr-yeqown-v2-compressed-1px.png` but file size is still differs.

