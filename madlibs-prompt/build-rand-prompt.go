package madlibsprompt

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func getRandFromList(itemList []string) (string, error) {

	//get random item from list
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(itemList))))
	if err != nil {
		return "", fmt.Errorf("error generating rand index: %s", err)
	}

	randIndex := int(nBig.Int64())
	randItem := itemList[randIndex]

	return randItem, nil
}

func BuildRandPrompt() (string, error) {
	prompt := ""
	// What kind of artwork
	artwork, err := getRandFromList(basicObjects)
	if err != nil {
		return "", fmt.Errorf("unable to get random artwork: %s", err)
	}
	// Tailor options to artwork's type
	if artwork == "painting" {
		object, err := imaginePainting()
		if err != nil {
			return "", err
		}
		prompt = object

	} else if artwork == "sculpture" {
		object, err := imagineSculpture()
		if err != nil {
			return "", err
		}

		prompt = fmt.Sprintf("Photograph of museum piece for art catalog: %s The object is by itself and centered on a plain background.", object)

	} else if artwork == "object d'art" {
		object, err := imagineObject()
		if err != nil {
			return "", err
		}

		prompt = fmt.Sprintf("Photograph of museum piece for art catalog: %s The object is by itself and centered on a plain background.", object)

	} else if artwork == "photography" {
		object, err := imaginePhoto()
		if err != nil {
			return "", err
		}

		prompt = object
	}

	print(prompt)
	return prompt, nil
}

func imaginePainting() (string, error) {
	// Artwork's characteristic
	mood, err := getRandFromList(genreMood)
	if err != nil {
		return "", err
	}
	// Painting type
	paint_type, err := getRandFromList(paint_type)
	if err != nil {
		return "", err
	}
	// Medium
	paint_medium, err := getRandFromList(paint_media)
	if err != nil {
		return "", err
	}
	// Theme
	paint_theme, err := getRandFromList(paint_themes)
	if err != nil {
		return "", err
	}
	// Art movement
	paint_movement, err := getRandFromList(paint_movements)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s on %s %s from the %s movement", mood, paint_type, paint_medium, paint_theme, paint_movement), nil
}

func imagineSculpture() (string, error) {
	// Characteristic
	mood, err := getRandFromList(genreMood)
	if err != nil {
		return "", err
	}
	// Material
	material, err := getRandFromList(sculpt_material)
	if err != nil {
		return "", err
	}
	// Sculpture type
	sculpt_type, err := getRandFromList(sculpt_types)
	if err != nil {
		return "", err
	}
	// Movement
	art_movement, err := getRandFromList(sculpt_movements)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s from the %s movement", mood, material, sculpt_type, art_movement), nil

}

func imagineObject() (string, error) {
	// Characteristic
	mood, err := getRandFromList(genreMood)
	if err != nil {
		return "", err
	}
	// Object
	object, err := getRandFromList(objects)
	if err != nil {
		return "", err
	}
	// Movement
	movement, err := getRandFromList(design_movements)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s from the %s movement.", mood, object, movement), nil
}

func imaginePhoto() (string, error) {
	// Artwork's characteristic
	mood, err := getRandFromList(genreMood)
	if err != nil {
		return "", err
	}
	//Select era
	timePeriod, err := getRandFromList(photoPeriod)
	if err != nil {
		return "", err
	}
	//Rest will come from if statement
	photoType := ""
	photoMedia := ""
	photoMovement := ""
	//Build photo
	if timePeriod == "early photography" {
		photoType, err = getRandFromList(earlyPhotoTypes)
		if err != nil {
			return "", err
		}
		photoMedia, err = getRandFromList(earlyPhotoMedia)
		if err != nil {
			return "", err
		}
		photoMovement, err = getRandFromList(earlyPhotoMovements)
		if err != nil {
			return "", err
		}
	} else if timePeriod == "turn of the century" {
		photoType, err = getRandFromList(turnCenturyTypes)
		if err != nil {
			return "", err
		}
		photoMedia, err = getRandFromList(turnCenturyMedia)
		if err != nil {
			return "", err
		}
		photoMovement, err = getRandFromList(turnCenturyMovements)
		if err != nil {
			return "", err
		}
	} else if timePeriod == "first half 20th century" {
		photoType, err = getRandFromList(fHalf20Types)
		if err != nil {
			return "", err
		}
		photoMedia, err = getRandFromList(fHalf20Media)
		if err != nil {
			return "", err
		}
		photoMovement, err = getRandFromList(fHalf20Movements)
		if err != nil {
			return "", err
		}
	} else if timePeriod == "second half 20th century" {
		photoType, err = getRandFromList(secHalf20Types)
		if err != nil {
			return "", err
		}
		photoMedia, err = getRandFromList(secHalf20Media)
		if err != nil {
			return "", err
		}
		photoMovement, err = getRandFromList(secHalf20Movements)
		if err != nil {
			return "", err
		}
	} else {
		photoType, err = getRandFromList(presentPhotoTypes)
		if err != nil {
			return "", err
		}
		photoMedia, err = getRandFromList(presentPhotoMedia)
		if err != nil {
			return "", err
		}
		photoMovement, err = getRandFromList(presentPhotoMovements)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s %s on %s from the %s movement.", mood, photoType, photoMedia, photoMovement), nil
}

func ObjectHistory() (string, error) {

	// subgenre, err := getRandFromList(subgenres)
	// if err != nil {
	// 	return "", err
	// }

	adj1, err := getRandFromList(generalMood)
	if err != nil {
		return "", err
	}

	adj2, err := getRandFromList(generalMood)
	if err != nil {
		return "", err
	}

	theme, err := getRandFromList(themes)
	if err != nil {
		return "", err
	}

	storyPrompt := fmt.Sprintf("This particular object carries an uncanny and original story of %s and %s. Its history has always raised questions about %s.", adj1, adj2, theme)

	return storyPrompt, nil
}

func ObjectAnecdote() (string, error) {
	protagonist, err := getRandFromList(protagonists)
	if err != nil {
		return "", err
	}

	fate, err := getRandFromList(fates)
	if err != nil {
		return "", err
	}

	ending, err := getRandFromList(endings)
	if err != nil {
		return "", err
	}

	anecdotePrompt := fmt.Sprintf("Directly related to this object, tell the story of a %s who due to interacting with this artwork %s. This uncanny story has a %s.", protagonist, fate, ending)

	return anecdotePrompt, nil
}

func GetArtistInfo() (string, error) {

	artistTypes := []string{"man", "man", "man", "woman", "woman", "woman", "collective", "unknown"}
	artistAges := []string{"young", "adult", "middle-aged", "mature"}
	// artistOrigins := []string{"Europe", "North America", "South America", "Africa", "Asia", "Oceania", "Another dimension"}

	artistType, err := getRandFromList(artistTypes)
	if err != nil {
		return "", err
	}
	artistAge, err := getRandFromList(artistAges)
	if err != nil {
		return "", err
	}

	// artistOrigin, err := getRandFromList(artistOrigins)
	// if err != nil {
	// 	return "", err
	// }

	artistInfo := ""

	if artistType == "man" || artistType == "woman" {
		artistInfo = fmt.Sprintf("The artist is an imaginary remarkable %s %s.", artistAge, artistType)
	} else if artistType == "collective" {
		artistInfo = fmt.Sprintf("This piece is attributed to an imaginary art %s", artistType)
	} else {
		artistInfo = fmt.Sprintf("The author of this piece is %s", artistType)
	}

	return artistInfo, nil
}
