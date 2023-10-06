package tests

import (
	"fmt"
	"random-quote-picture/service"
	"testing"
)

func TestGetPictureUrl(t *testing.T) {
	grayscales := []string{"1", "", "0", "abc"}
	for _, grayscale := range grayscales {
		testName := fmt.Sprintf("grayscale-%s", grayscale)
		t.Run(testName, func(t *testing.T) {
			pictureUrl, err := service.GetPictureUrl(grayscale)
			if err != nil {
				t.Errorf("GetPictureUrl Failed, error: %s\n", err.Error())
			}
			if pictureUrl == "" {
				t.Errorf("GetPictureUrl Failed, pictureUrl is empty.\n")
			}
			t.Log("GetPictureUrl PASSED.")
		})
	}

}
