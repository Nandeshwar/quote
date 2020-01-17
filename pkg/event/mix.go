package event

func MixEvents() []*EventDetail {
	var allEvents []*EventDetail

	NeemKaroliBabaDisAppearance := &EventDetail{
		Day:   11,
		Month: 9,
		Year:  1973,
		Title: "Neem Karoli Baba - Disappearance day",
		Info: `    Neem Karoli Baba or Neeb Karori Baba - known to his followers as Maharaj-ji - was a Hindu guru, mystic and devotee of the Hindu lord Hanuman.
    Full name: Lakshmi Narayan Sharma. Born: Akbarpur, India
    Parents: Durga Prasad Sharma
    Children: Aneg Singh Sharma, Girija Bhatele, Dharma Narayan Sharma
`,
		URL: "https://en.wikipedia.org/wiki/Neem_Karoli_Baba",
	}
	allEvents = append(allEvents, NeemKaroliBabaDisAppearance)

	return allEvents
}
