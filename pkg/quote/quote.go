package quote

import (
	"math/rand"
	"time"
)

func QuoteForTheDay() string {
	quotes := []string{
		"Time spent in love is never waste",
		"Enjoy every moment",
		"Wherever there is love, there is God",
		"The real way to loving is giving not demanding",
		"No one is greater or smaller than other. Everyone in this world is unique. Love everyone",
		"The person who has heart full of love, has always something to give",
		"Hell has three gates lust, anger, greed",
		"Be happy with nothing and you will be happy with everything",
		"Detachment is not that you should own nothing, but that nothing should own you",
		"Devotion has the power to burn down any karma",
		"Love is the greatest power on earth",
		"When you wish good for others, good things come back to you. This is the Law of Nature",
		"If you can win over your mind, you can win over the whole world",
		"Darkness cannot drive out darkness, only light can do that. Hate cannot drive out hate. only love can do that",
		"Silence says so mcuh. Just listen",
		"The greatest gift human can give to himself and others are tolerance and forgiveness",
		"The practice of devotion involves replacing desires for the world with the desires for God",
		"The wealth of divine love is the only true wealth. Every other form of wealth simply enhances one's pride",
		"Speak only when you feel your words are better than the silence",
		"For our spiritual growth, negative people are often placed in our path, so we may learn selfless love, forgiveness & surrender to the will of God",
		"The happiest people are givers not takers",
		"Why do we close our eyes when we pray, cry, kiss or dream? Because the most beautiful things in life are not seen, but felt by the heart",
		"If you have to choose between being kind and being right choose being kind and you will always be right",
		"Silence & Smile are two powerful tools.Smile is the way to solve many problems & Silence is the way to avoid many problems",
		"Don't get upset with people and situations, because both are powerless without your reaction",
		"Most of the people are in lack of knowledge.Don't hate them.Love them and understand that they are under influence of ignorance. Always do righteously.",
		"Every way and means that leads our mind to God is Devotion",
		"The Only Purpose of Our Human Life is to Attain God",
		"Meditation. Because some questions can't be answered by Google!",
		"This is the nature of existence - if you do the right things, the right things will happen to you",
		"Devotion is the only way to be liberated from material attachment. It is only then that we become free from lust, anger and greed",
		"I belong to no religion. My religion is love. Every heart is my temple",
		"Don't focus too much on the defects, be aware of them, but our endeavor should be towards positive",
		"To purify the mind, we must engage in devotion to the lord, When our mind is purified, out attitude and our experience will change towards the outer world",
		"The reason that we are in a state of suffering and we are enveloped in the darkness of material existence, is our forgetfulness of God",
		"If you can establish your relationship with God, that ultimate satisfaction that you have been searching for since innumerable lifetimes, will eventually be attained",
		"The Joy of the mind is the measure of its strength",
		"When you come to a point where you have no need to impress anybody, your freedom will begin",
		"Ritualistic worship, chanting and meditation are done with the body, voice and the mind: they excel each other in the ascending order",
		"Uttering the sacred word, either in a loud or low tone is preferable to chants in praise of the Supreme. Mental contemplation is superior to both",
		"When one learns to turn the mind away from material allurements and renounces the desires of the senses, such a person comes in touch with the inner bliss of the soul",
		"When we decide that God is ours and the whole world is His, then our consciousness transforms from seeking self-enjoyment to serving the Lord with everything that we have",
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	ind := r.Intn(len(quotes))

	return quotes[ind]
}

func QuoteForTheDayImage() string {
	quotes := []string{
		"image/competitionWithMySelf.jpg",
		"image/pleasegod.jpg",
		"image/alwaysdogood.jpg",
		"image/becomegood.jpg",
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	ind := r.Intn(len(quotes))

	return quotes[ind]
}
