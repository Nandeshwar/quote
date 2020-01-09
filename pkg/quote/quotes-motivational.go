package quote

import (
	"math/rand"
	"time"
)

func QuoteMotivationalImage() string {
	quotes := []string{
		"image/competitionWithMySelf.jpg",
		"image/becomegood.jpg",
		"image/hope.jpg",
		"image/thanks-to-obstacles.jpg",
		"image/you-are-responsible.jpg",
		"image/lift-up.jpg",
		"image/forgive.jpg",
		"image/wish-good-for-others.jpg",
		"image/anyway-mother-teresa.jpeg",
		"image/Mother-Teresa-love-family.jpg",
		"image/Beautiful-Mother-Teresa-Quotes.jpg",
		"image/Mother-Teresa-children-of-god.jpg",
		"image/Mother-Teresa-with-love.jpg",
		"image/Mother-Teresa-Quotes-on-Love.jpg",
		"image/Mother-Teresa-Quote-on-Love-Life.jpg",
		"image/Mother-teresa-love-people.jpg",
		"image/Best-Mother-Teresa-with-love.jpg",
		"image/Mother-Teresa-make-others-happy.jpg",
		"image-motivational/mother-teresa-we-have-today.jpg",
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	ind := r.Intn(len(quotes))

	return quotes[ind]
}