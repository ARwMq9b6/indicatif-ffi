package main

import (
	"time"

	"github.com/ARwMq9b6/indicatif-ffi/indicatif-ffi-go-cgo"
)

func main() {
	var downloaded uint64
	var total_size uint64 = 231231231

	style := indicatif.NewProgressStyleWithDefaultBar()
	style.SetTemplate("{spinner:.green} [{elapsed_precise}] [{bar:40.cyan/blue}] {bytes}/{total_bytes} ({eta})")
	style.SetProgressChars("#>-")

	pb := indicatif.NewProgressBar(total_size)
	pb.SetStyle(style)

	for downloaded < total_size {
		if downloaded += 223211; downloaded > total_size {
			downloaded = total_size
		}
		pb.SetPosition(downloaded)
		time.Sleep(12 * time.Millisecond)
	}

	pb.FinishWithMessage("downloaded")
}
