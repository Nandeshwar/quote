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

	return allEvents
}
