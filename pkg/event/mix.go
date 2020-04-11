package event

import "time"

func MixEvents() []*EventDetail {
	var allEvents []*EventDetail

	NeemKaroliBabaDisAppearance := &EventDetail{
		Day:   11,
		Month: 9,
		Year:  1973,
		Title: "Neem Karoli Baba - Disappearance day",
		Info: `He was born around 1900, in village Akbarpur, Firozabad district, Uttar Pradesh, India, in a Brahmin family of Durga Prasad Sharma. He was named Lakshman Das Sharma. After being married by his parents aged 11 he left home to become a wandering sadhu.
Neem Karoli Baba or Neeb Karori Baba - known to his followers as Maharaj-ji - was a Hindu guru, mystic and devotee of the Hindu lord Hanuman. Wikipedia
Born: Akbarpur, India.
Died: September 11, 1973, Vrindavan, India.
Full name: Lakshmi Narayan Sharma.
Guru: Hanuman.
Parents: Durga Prasad Sharma.
Children: Aneg Singh Sharma, Girija Bhatele, Dharma Narayan Sharma.
`,
		URL:          "https://en.wikipedia.org/wiki/Neem_Karoli_Baba",
		CreationDate: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
	}
	allEvents = append(allEvents, NeemKaroliBabaDisAppearance)

	ramanandSagarAppearance := &EventDetail{
		Day:   29,
		Month: 12,
		Year:  1917,
		Title: "Ramanand Sagar - Appearance day",
		Info: ` Ramanand Sagar was an Indian film director. He is most famous for making the Ramayan television series, a 78-part TV adaptation of the ancient Hindu epic of the same name, starring Arun Govil as Lord Ram and Deepika Chikhalia as Sita. This TV serial was then widely watched and liked across the country.
Born: December 29, 1917, Punjab, Pakistan.
Died: December 12, 2005, Mumbai, India.
Grandchildren: Meenakshi Sagar, Preeti Sagar, Amrit Sagar, Akash Chopra, Shakti Sagar, Namita Sagar, Jyoti Sagar.
Books: Bleeding Partition: A Novel.
Great grandchild: Sakshi Chopra.
`,
		URL:          "https://www.google.com/search?q=ramanand+sagar&rlz=1C5GCEA_enUS869US869&oq=ramanand+sagar&aqs=chrome..69i57j0l4j69i65j69i60l2.3223j0j7&sourceid=chrome&ie=UTF-8",
		CreationDate: time.Date(2019, 2, 1, 11, 28, 0, 0, time.Local),
	}
	allEvents = append(allEvents, ramanandSagarAppearance)

	ramanandSagarDisAppearance := &EventDetail{
		Day:   12,
		Month: 12,
		Year:  2005,
		Title: "Ramanand Sagar - Disappearance day",
		Info: ` Ramanand Sagar was an Indian film director. He is most famous for making the Ramayan television series, a 78-part TV adaptation of the ancient Hindu epic of the same name, starring Arun Govil as Lord Ram and Deepika Chikhalia as Sita. This TV serial was then widely watched and liked across the country.
Born: December 29, 1917, Punjab, Pakistan.
Died: December 12, 2005, Mumbai, India.
Grandchildren: Meenakshi Sagar, Preeti Sagar, Amrit Sagar, Akash Chopra, Shakti Sagar, Namita Sagar, Jyoti Sagar.
Books: Bleeding Partition: A Novel.
Great grandchild: Sakshi Chopra.
`,
		URL:          "https://www.google.com/search?q=ramanand+sagar&rlz=1C5GCEA_enUS869US869&oq=ramanand+sagar&aqs=chrome..69i57j0l4j69i65j69i60l2.3223j0j7&sourceid=chrome&ie=UTF-8",
		CreationDate: time.Date(2019, 2, 1, 11, 28, 0, 0, time.Local),
	}
	allEvents = append(allEvents, ramanandSagarDisAppearance)

	arunGovilAppearance := &EventDetail{
		Day:   30,
		Month: 10,
		Year:  1958,
		Title: "Arun Govil - Ram in Ramayan of Ramanand Sagar - appearance day",
		Info: ` Arun Govil played role of Ram in Ramayan of Ramanand Sagar.
Arun Govil. Arun Govil is an Indian actor and producer. ... He was then cast as Lord Rama in Sagar's TV series Ramayan (1986), for which he won the Uptron Award in the Best Actor in a Leading Role category in 1988. He reprised his role as Rama in Sagar's Luv Kush.
Parents: Shri Chandra Prakash Govil
Film: Paheli, Sawan Ko Aane Do, Saanch Ko A...
Sibling: Vijay Govil
TV show: Ramayan, Vikram Aur Betaal, Luv Kush
`,
		URL:          "https://en.wikipedia.org/wiki/Arun_Govil",
		CreationDate: time.Date(2020, 4, 11, 14, 26, 0, 0, time.Local),
	}
	allEvents = append(allEvents, arunGovilAppearance)

	DeepikaChikhalia := &EventDetail{
		Day:   29,
		Month: 4,
		Year:  1965,
		Title: "Deepika Chikhalia - Sita in Ramayan of Ramanand Sagar - appearance day",
		Info: `Dipika Topiwala is an Indian actress who rose to fame playing Devi Sita in Ramanand Sagar's hit television serial Ramayan and was known for acting in other Indian historical TV serials. Wikipedia
Born: April 29, 1965 (age 54 years), Mumbai, India
Height: 4′ 10″
Spouse: Hemant Topiwala
Party: Bharatiya Janata Party
Children: Juhi Topiwala, Nidhi Topiwala
`,
		URL:          "https://en.wikipedia.org/wiki/Deepika_Chikhalia",
		CreationDate: time.Date(2020, 4, 11, 14, 26, 0, 0, time.Local),
	}
	allEvents = append(allEvents, DeepikaChikhalia)

	sunilLahri := &EventDetail{
		Day:   9,
		Month: 1,
		Year:  1961,
		Title: "Sunil Lahri - Laxman in Ramayan of Ramanand Sagar - appearance day",
		Info: `Sunil Lahri is an Indian actor. He is most famous for appearing in the television works of Ramanand Sagar, beginning with his most famous role as Lakshman in Ramayan. Before Ramayan, he appeared in some stories of Vikram aur Betaal and Dada-Dadi Ki Kahaniyan
`,
		URL:          "https://en.wikipedia.org/wiki/Sunil_Lahri",
		CreationDate: time.Date(2020, 4, 11, 14, 26, 0, 0, time.Local),
	}
	allEvents = append(allEvents, sunilLahri)

	sanjayJog := &EventDetail{
		Day:   24,
		Month: 9,
		Year:  1955,
		Title: "Sanjay Jog - Bharat in Ramayan of Ramanand Sagar - appearance day",
		Info: `Sanjay Jog played role of bharat ji in  Ramanand Sagar's Ramayan.
`,
		URL:          "https://maitrimanthan.wordpress.com/2011/10/10/sanjay-jog/",
		CreationDate: time.Date(2020, 4, 11, 14, 26, 0, 0, time.Local),
	}
	allEvents = append(allEvents, sanjayJog)

	DaraSinghAppearance := &EventDetail{
		Day:   19,
		Month: 11,
		Year:  1928,
		Title: "DaraSinghAppearance - Hanuman ji in Ramayan of Ramanand Sagar - appearance day",
		Info: `2010
Born	Deedar Singh Randhawa
19 November 1928
Dharmuchak,
`,
		URL:          "https://en.wikipedia.org/wiki/Dara_Singh",
		CreationDate: time.Date(2020, 4, 11, 14, 26, 0, 0, time.Local),
	}
	allEvents = append(allEvents, DaraSinghAppearance)

	DaraSinghDisAppearance := &EventDetail{
		Day:   12,
		Month: 7,
		Year:  2012,
		Title: "DaraSinghDisAppearance - Hanuman ji in Ramayan of Ramanand Sagar - disappearance day",
		Info: `Died	12 July 2012 (aged 83)
`,
		URL:          "https://en.wikipedia.org/wiki/Dara_Singh",
		CreationDate: time.Date(2020, 4, 11, 14, 26, 0, 0, time.Local),
	}
	allEvents = append(allEvents, DaraSinghDisAppearance)

	return allEvents
}
