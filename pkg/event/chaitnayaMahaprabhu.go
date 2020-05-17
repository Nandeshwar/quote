package event

import "time"

func ChaitnaMahaprabhuEvents() []*EventDetail {
	var allEvents []*EventDetail
	chaitnyaMahaprbhuAppearance := &EventDetail{
		Day:   18,
		Month: 2,
		Year:  1486,
		Title: "Chaitnya Mahaprbhu - Appearance day",
		Info: `    Shri Krishna Chaitanya Mahaprabhu, honorific: "Mahāprabhu", was an Indian guru considered by his followers to be the Supreme Personality of Godhead and the chief proponent of the Achintya Bheda Abheda Vedanta school and the Gaudiya Vaishnavism tradition within Hinduism. Wikipedia
Born: February 18, 1486, Nabadwip, India
Died: June 14, 1534, Puri, India
Full name: Vishvambhar Mishra
Spouse: Lakshmipriya (m. ?–1505)
Guru: Swami Isvara Puri (mantra guru); Swami Kesava Bharati (sannyas guru)
Books: Siksastaka, The Vedânta Sûtras of Bâdarâyaṇa
`,
		URL:          "https://en.wikipedia.org/wiki/Chaitanya_Mahaprabhu",
		CreationDate: time.Date(2020, 5, 17, 12, 40, 0, 0, time.Local),
	}
	allEvents = append(allEvents, chaitnyaMahaprbhuAppearance)

	chaitnyaMahaprbhuDisAppearance := &EventDetail{
		Day:   14,
		Month: 6,
		Year:  1533,
		Title: "Chaitnya Mahaprbhu - Disappearance day",
		Info: `    Shri Krishna Chaitanya Mahaprabhu, honorific: "Mahāprabhu", was an Indian guru considered by his followers to be the Supreme Personality of Godhead and the chief proponent of the Achintya Bheda Abheda Vedanta school and the Gaudiya Vaishnavism tradition within Hinduism. Wikipedia
Born: February 18, 1486, Nabadwip, India
Died: June 14, 1534, Puri, India
Full name: Vishvambhar Mishra
Spouse: Lakshmipriya (m. ?–1505)
Guru: Swami Isvara Puri (mantra guru); Swami Kesava Bharati (sannyas guru)
Books: Siksastaka, The Vedânta Sûtras of Bâdarâyaṇa
`,
		URL:          "https://en.wikipedia.org/wiki/Chaitanya_Mahaprabhu",
		CreationDate: time.Date(2020, 5, 17, 12, 40, 0, 0, time.Local),
	}
	allEvents = append(allEvents, chaitnyaMahaprbhuDisAppearance)

	return allEvents
}
