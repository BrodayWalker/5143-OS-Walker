package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["trump"] = Trump
}

// Trump ...
func Trump(args []string) {
	// Open trump sound file
	soundf, err := os.Open("assets/trump.wav")
	if err != nil {
		fmt.Println("Error opening sound file: ", err)
	}
	defer soundf.Close()

	// Read the text file into a string
	content, err := ioutil.ReadFile("assets/trump")
	if err != nil {
		fmt.Println("Error opening text file: ", err)
	}
	trumpTxt := string(content)
	fmt.Println(trumpTxt)

	// Create the audio streamer
	streamer, format, err := wav.Decode(soundf)
	if err != nil {
		fmt.Println("Error decoding wav file: ", err)
	}
	defer streamer.Close()
	// Create a speaker to play the audio
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	// Play the sound
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done

}
