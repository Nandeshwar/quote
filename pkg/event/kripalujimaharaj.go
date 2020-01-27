package event

import "time"

func KripaluJiMaharajEvents() []*EventDetail {
	var allEvents []*EventDetail
	kripaluJiJagadGurutam := &EventDetail{
		Day:   14,
		Month: 1,
		Year:  1957,
		Title: "Kripalu Ji Maharaj - Jagadgurutam title day",
		Info: `    He was formally installed as the fifth Jagadguru (world teacher).
   He was 34 years old when given the title on 14 January 1957 by the Kashi Vidvat Parishat, a group of Hindu scholars.
   The Kashi Vidvat Parishat conferred on him the titles Bhaktiyog-Ras-Avtar and Jagadguruttama.
   Followers claim that he is the "fifth original Jagadguru" in the series of Jagadgurus after 
   Śrīpāda Śaṅkarācārya (A.D. 788-820),
   Śrīpāda Rāmānujācārya (1017-1137),
   Śrī Nimbārkācārya and, 
   Śrīpāda Madhvācārya (1239-1319)

    Jagadguru Kripalu Ji Maharaj appeared in a small village called Mangarh, near Allahabad, in India, on the auspicious night of Sharat Purnima in October 1922. 
    His mother, Bhagvati Devi, and father, Lalita Prasad, named Him Ram Kripalu at birth. 
    From the very first day, He delighted the hearts of everyone around Him with His sweet smile and serene look
`,
		URL:          "https://en.wikipedia.org/wiki/Kripalu_Maharaj",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, kripaluJiJagadGurutam)

	kripaluJiAppearance := &EventDetail{
		Day:   5,
		Month: 10,
		Year:  1922,
		Title: "Kripalu Ji Maharaj - appearance day",
		Info: `    Kripalu (Sanskrit: जगद्गुरु श्री कृपालु जी महाराज, IAST: Kṛpālu) (5 October 1922 – 15 November 2013) 
    He was a Hindu spiritual leader  from Allahabad (Prayag) - Mangarh, Pratapgarh, India.
    He was the founder of Jagadguru Kripalu Parishat (JKP), 
    a worldwide Hindu non-profit organization with five main ashrams; four in India and one in the United States.
    JKP Radha Madhav Dham is one of the largest Hindu Temple complexes in the Western Hemisphere, and the largest in North America.

    Jagadguru Kripalu Ji Maharaj appeared in a small village called Mangarh, near Allahabad, in India, on the auspicious night of Sharat Purnima in October 1922. 
    His mother, Bhagvati Devi, and father, Lalita Prasad, named Him Ram Kripalu at birth. 
    From the very first day, He delighted the hearts of everyone around Him with His sweet smile and serene look
     `,
		URL:          "https://en.wikipedia.org/wiki/Kripalu_Maharaj",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, kripaluJiAppearance)

	kripaluJiDisAppearance := EventDetail{
		Day:   15,
		Month: 11,
		Year:  2013,
		Title: "Kripalu Ji Maharaj - disappearance day",
		Info: `He disappeared 15 November 2013 (aged 91) New Delhi, Delhi, the dense forest near Sharbhang ashram in Chitrakoot.

   He was formally installed as the fifth Jagadguru (world teacher).
   He was 34 years old when given the title on 14 January 1957 by the Kashi Vidvat Parishat, a group of Hindu scholars.
   The Kashi Vidvat Parishat conferred on him the titles Bhaktiyog-Ras-Avtar and Jagadguruttama.
   Followers claim that he is the "fifth original Jagadguru" in the series of Jagadgurus after 
   Śrīpāda Śaṅkarācārya (A.D. 788-820),
   Śrīpāda Rāmānujācārya (1017-1137),
   Śrī Nimbārkācārya and, 
   Śrīpāda Madhvācārya (1239-1319)

    Jagadguru Kripalu Ji Maharaj appeared in a small village called Mangarh, near Allahabad, in India, on the auspicious night of Sharat Purnima in October 1922. 
    His mother, Bhagvati Devi, and father, Lalita Prasad, named Him Ram Kripalu at birth. 
    From the very first day, He delighted the hearts of everyone around Him with His sweet smile and serene look
    `,
		URL:          "https://en.wikipedia.org/wiki/Kripalu_Maharaj;https://www.jkyog.org/;http://jkp.org.in/life-story-jagadguruttam/",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, &kripaluJiDisAppearance)

	premMandirInauguration := EventDetail{
		Day:   15,
		Month: 2,
		Year:  2012,
		Title: "Kripalu Ji Maharaj-Prem Mandir - Inauguration day",
		Info: ` Prem Mandir is in Vrindavan, Mathura, India. The temple was opened to public on 17 February
    The temple structure was established by the fifth Jagadguru, Kripalu Maharaj.
    Figures of Shri Krishna and his followers depicting important events surrounding the Lord's existence cover the main temple.
    The foundation stone was laid by Jagadguru Shri Kripalu Ji Maharaj in the presence of thousand devotees on 14 January 2001. 
    It took approximately 1000 artists about 12 years to build the complex.
   
    `,
		URL:          "https://en.wikipedia.org/wiki/Prem_Mandir,_Vrindavan",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, &premMandirInauguration)

	bhaktiMandirInauguration := EventDetail{
		Day:   1,
		Month: 11,
		Year:  2005,
		Title: "Kripalu Ji Maharaj-Bhakti Mandir - Inauguration day",
		Info: ` The foundation stone of Bhakti Mandir was laid on 26 October 1996, and was inaugurated in November 2005.
    Note: I did not find exact day in November.
    Bhakti Mandir is a Hindu Temple located in Kunda, India. 
    This divine temple was established by the world's fifth original Jagadguru in November 2005. 
    It is maintained by Jagadguru Kripalu Parishat, a non-profit, charitable, educational and spiritual organisation.
   
    `,
		URL:          "https://en.wikipedia.org/wiki/Bhakti_Mandir_Mangarh",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, &bhaktiMandirInauguration)

	mukundanandJiAppearance := EventDetail{
		Day:   19,
		Month: 12,
		Year:  1960,
		Title: "Kripalu Ji Maharaj Senior disciple Mukundanand Ji - appearance day",
		Info: ` As a child, Swami Mukundananda spent long hours in meditation and contemplation.
    He graduated from the Indian Institute of Technology, Delhi with a degree in engineering and received a postgraduate degree from the Indian Institute of Management Calcutta.

    After that, he worked for some time with one of the India's most topmost industrial houses. 
    He left a career in business to join the order of Sannyas, dedicating his time to devotional pursuits and travelled throughout India as a Sanyasi.
    Under the guidance of guru Jagadguru Shree Kripaluji Maharaj who is lovingly called "Maharajji" by his devotees,
    Mukundananda studied the Vedic scriptures,Indian and Western philosophy, and Bhakti Yog. 
    Kripaluji Maharaj entrusted him with the key task of spreading Vedic knowledge around the globe
   
    `,
		URL:          "https://en.wikipedia.org/wiki/Mukundananda",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, &mukundanandJiAppearance)

	return allEvents
}
