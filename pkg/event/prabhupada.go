package event

import "time"

func PrabhuyEvents() []*EventDetail {
	var allEvents []*EventDetail
	prabhupadaJiAppearance := &EventDetail{
		Day:   1,
		Month: 9,
		Year:  1896,
		Title: "A. C. Bhaktivedanta Swami Prabhupada - Appearance day",
		Info: `    Abhaya Charanaravinda Bhaktivedanta Svami (born Abhay Charan De; 1 September 1896 – 14 November 1977) was an Indian spiritual teacher 
    and the founder-preceptor of the International Society for Krishna Consciousness[2] (ISKCON), commonly known as the "Hare Krishna Movement".
    Members of the ISKCON movement view Bhaktivedānta Swāmi as a representative and messenger of Krishna Chaitanya.
`,
		URL:          "https://en.wikipedia.org/wiki/A._C._Bhaktivedanta_Swami_Prabhupada",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, prabhupadaJiAppearance)

	prabhupadaJiDisAppearance := &EventDetail{
		Day:   14,
		Month: 11,
		Year:  1977,
		Title: "A. C. Bhaktivedanta Swami Prabhupada - Disappearance day",
		Info: `    Bhaktivedanta Swami died on 14 November 1977 in Vrindavan, India, and his body was buried in Krishna Balaram Mandir in Vrindavan India.
`,
		URL:          "https://en.wikipedia.org/wiki/A._C._Bhaktivedanta_Swami_Prabhupada",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, prabhupadaJiDisAppearance)

	return allEvents
}
