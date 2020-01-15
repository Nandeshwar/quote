package quote

import (
	"math/rand"
	"quote/pkg/fileutil"
	"time"

	"github.com/logic-building/functional-go/fp"
	"github.com/sirupsen/logrus"
)

func AllMotivationalImage() (int, []string) {
	quotes := []string{
		"image/competitionWithMySelf.jpg",
		"image/becomegood.jpg",
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
		"image/Einstein-strong-people.jpg",
		"image/thankful-to-everything.jpg",
		"image/meditationgoogle.jpg",
		"image/renew-humanity.jpg",
		"image/nandeshwar-meditation.jpg",
	}

	imagesUnderDir1, err := fileutil.ListDir("./image-motivational")
	if err != nil {
		logrus.Errorf("Unable to read files from ./image-motivational=%v", err)
	}

	imagesUnderDir2, err := fileutil.ListDir("/image-motivational")
	if err != nil {
		logrus.Errorf("Unable to read files from /image-motivational=%v", err)
	}

	onlyJPGImages := fp.FilterStr(validJPG, imagesUnderDir1)
	quotes = append(quotes, onlyJPGImages...)

	onlyJPGImages = fp.FilterStr(validJPG, imagesUnderDir2)
	quotes = append(quotes, onlyJPGImages...)

	return len(quotes), quotes
}

func GetQuoteMotivationalImage(allImages []string) string {

	s2 := rand.NewSource(int64(time.Now().Nanosecond()))
	r2 := rand.New(s2)

	ind := r2.Intn(len(allImages))

	return allImages[ind]
}
