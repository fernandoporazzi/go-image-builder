package main

import(
  "bytes"
  "image"
  "image/color"
  "image/draw"
  "image/jpeg"
  "log"
  "net/http"
  "strconv"
)

func createImage(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  width, _ := strconv.Atoi(query.Get("width"))
  height, _ := strconv.Atoi(query.Get("height"))
  hexa := query.Get("background")

  red, green, blue := HexToRGB(hexa)

  img := image.NewRGBA( image.Rect( 0, 0, width, height ) )

  // set rgba without transparency
  color := color.RGBA{red, green, blue, 255}

  draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)

  var i image.Image = img

  putImage(w, &i)
}

func putImage(w http.ResponseWriter, i *image.Image) {
  buffer := new(bytes.Buffer)

  if err := jpeg.Encode(buffer, *i, nil); err != nil {
    log.Println("Unable to encode image")
  }

  w.Header().Set("Content-Type", "image/jpeg")
  w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

  if _, err := w.Write(buffer.Bytes()); err != nil {
    log.Println("unable to write image.")
  }
}