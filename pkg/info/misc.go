package info

func getMiscInfo() []Info {
	infoList := []Info{
		// info1
		{
			Title: "All Jagad guru",
			Info: `
Śrīpāda Śaṅkarācārya (Sanskrit: श्रीपाद शंकराचार्य) (A.D. 788–820) (also known as "Ādi Śaṅkara" (Sanskrit: आदि शंकर)), or Śaṅkara Bhagavatpāda (Sanskrit: शंकर भगवत्पाद): Founder of Advaita school of Vedānta.[6]
Śrīpāda Rāmānujācārya (Sanskrit: श्रीपाद रामानुजाचार्य) (1017–1137): Founder of Viśiṣṭādvaita school of Vedānta.[6]
Śrīpāda Madhvācārya (Sanskrit: श्रीपाद मध्वाचार्य) (A.D. 1239–1319) (also known as "Pūrna Prajña" (Sanskrit: पूर्णा प्रज्ञ) or "Ānanda Tīrtha Bhagavatpāda" (Sanskrit: आनन्द तीर्थ भगवत्पाद): Founder of the Dvaita school of Vedānta.[6]
Śrī Nimbārkācārya (Sanskrit: श्री निम्बार्काचार्य): Founder of Dvaitadvaita school of Vedānta.[6]
Śrī Vāllābhacārya (Sanskrit: श्री वल्लभाचार्य) (1479–1531): Founder of Shuddhadvaita school of Vedānta.[6]
Sri Kripalu ji Maharaj ( 1922 - 2013)
`,
			Link: []string{
				"https://en.wikipedia.org/wiki/Jagadguru",
				"https://en.wikipedia.org/wiki/Kripalu_Maharaj",
			},
		},

		// info 2
		{
			Title: "4 Vedas",
			Info: `
1. Rigveda: The Rigveda (Sanskrit: ऋग्वेद ṛgveda, from ṛc "praise" and veda "knowledge") is an ancient Indian collection of Vedic Sanskrit hymns. 
It is one of the four sacred canonical texts (śruti) of Hinduism known as the Vedas.

2. Yajurveda: The Yajurveda (Sanskrit: यजुर्वेद, yajurveda, from yajus meaning "worship",and veda meaning "knowledge") is the Veda primarily of prose mantras for worship rituals.
An ancient Vedic Sanskrit text, it is a compilation of ritual offering formulas that were said by a priest while an individual performed ritual actions such as those before the yajna fire.
Yajurveda is one of the four Vedas, and one of the scriptures of Hinduism. The exact century of Yajurveda's composition is unknown, 
and estimated by scholars to be around 1200 to 1000 BCE, contemporaneous with Samaveda and Atharvaveda.

3. Samaveda: The Samaveda (Sanskrit: सामवेद, sāmaveda, from sāman "song" and veda "knowledge"), is the Veda of melodies and chants.
It is an ancient Vedic Sanskrit text, and part of the scriptures of Hinduism. One of the four Vedas, it is a liturgical text which consists of 1,549 verses. 
All but 75 verses have been taken from the Rigveda. Three recensions of the Samaveda have survived, and variant manuscripts of the Veda have been found in various parts of India.[3][4]

4. Atharvaveda: The Atharva Veda (Sanskrit: अथर्ववेद, Atharvaveda from atharvāṇas and veda, meaning "knowledge") is the "knowledge storehouse of atharvāṇas, 
the procedures for everyday life". The text is the fourth Veda, but has been a late addition to the Vedic scriptures of Hinduism.
`,
			Link: []string{
				"https://en.wikipedia.org/wiki/Vedas",
				"https://en.wikipedia.org/wiki/Rigveda",
				"https://en.wikipedia.org/wiki/Yajurveda",
				"https://en.wikipedia.org/wiki/Samaveda",
				"https://en.wikipedia.org/wiki/Atharvaveda",
			},
		},

		// info 3
		{
			Title: "Wife of Guru Vashishtha - Arundhati",
			Info: ` Arundhati
In Rigvedic hymn 7.33.9, Vashishtha (Guru of Ram) is described as a scholar who moved across the Indus river to establish his school. 
He was married to Arundhati, and therefore he was also called Arundhati Nath, meaning the husband of Arundhati.
`,
			Link: []string{
				"https://en.wikipedia.org/wiki/Vasishtha",
			},
		},

		// info 4
		{
			Title: "Gita in 12 minutes",
			Info: `Nice explanation of Bhagwat gita in 12 minutes. 
`,
			Link: []string{
				"https://www.youtube.com/watch?v=jnifjBM9dpM",
			},
		},

		// info 5
		{
			Title: "Hanuman Mahima 1 hour video songs",
			Info: `Hanumnan ji bhakti towards Ram and Krishna. A beautiful song.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=N-RSr4ecs9M&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=9",
			},
		},

		// info 6
		{
			Title: "Garud puran in 56 min",
			Info: `A beautiful song - Garud puran
`,
			Link: []string{
				"https://www.youtube.com/watch?v=Vb8-7DLSJuE&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=35",
			},
		},

		// info 7
		{
			Title: "Bhajan - Govind Chale aao, Gopal chale aao",
			Info: `A beautiful bhajan of Krishna. Heart touching song.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=AEpGYCwutcc&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=49",
			},
		},

		// info 8
		{
			Title: "Bharat prem towards Ram",
			Info: `Ramayan - A beautiful video about Bharat Ji Prem towards Shree Ram Ji.
Very emotional video. True love. God is bound by love.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=Rc8v5SfXP5g&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=56",
			},
		},

		// info 9
		{
			Title: "Bhajan: Meri vinti hai radha rani",
			Info: `A beautiful bhajan of Radha Rani.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=3mmMm45rJpA&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=58",
			},
		},

		// info 10
		{
			Title: "Bhagwan Naam ki Mahatav- 12 years - Ayush Krishna Nayan Ji",
			Info: ` A beautiful story of Sushila, sudama told by 12 years old child. Speech given by this young kid is high level
of spiritual knowledge.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=jYkdBsUFMKs&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=101",
			},
		},

		// info 11
		{
			Title: "Bhajan: Itni Sakti dena hame",
			Info: ` Beautiful bhajan : Itni Sakti hame dena bhagwan
`,
			Link: []string{
				"https://www.youtube.com/watch?v=3EnLaJKhO2A&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=107",
			},
		},

		// info 12
		{
			Title: "Bhajan: Jahe vidhi rakhhe ram",
			Info: ` Heart touching bhajan : Sita ram, Sita ram kahiye - Jahe vidhi rakkhe ram
`,
			Link: []string{
				"https://www.youtube.com/watch?v=GnvQH8PH8sg&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=118",
			},
		},

		// info 13
		{
			Title: "Bhajan: Jai Radha Madhav",
			Info: ` A beautiful Bhajan, Jai Radha Madhav, Jai Kunj Bihari
`,
			Link: []string{
				"https://www.youtube.com/watch?v=cPTqAyqBAfc&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=119",
			},
		},

		// info 14
		{
			Title: "Bhajan: Ae ri Sakhhi Mangal gao ri: Rishi Nitya Nand",
			Info: ` A nice dance and bhajan - Ae ri Sakkhi Mangal gao ri
`,
			Link: []string{
				"https://www.youtube.com/watch?v=Ju4DR1A_vPY&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=120",
			},
		},

		// info 15
		{
			Title: "Mission Genius mind - Soul after death",
			Info: ` A story about girl's spirit
`,
			Link: []string{
				"https://www.youtube.com/watch?v=nPRKafjtGGs&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=127",
			},
		},

		// info 16
		{
			Title: "Bhajan: hum van ke bashi",
			Info: ` Nice bhajan from Ramanand Sagar Ramayan
`,
			Link: []string{
				"https://www.youtube.com/watch?v=wOkFqjRwQYY&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=134",
			},
		},

		// info 17
		{
			Title: "Bhajan: Radha ke maan me bas gaye",
			Info: ` Nice bhajan by Mridul krishna Sastri
`,
			Link: []string{
				"https://www.youtube.com/watch?v=OtjFIe84d0E",
			},
		},

		// info 17
		{
			Title: "Bhajan: O Kanha ab to murli ki madhur suna do taan",
			Info: ` Heart touching bhajan
`,
			Link: []string{
				"https://www.youtube.com/watch?v=XP9rlhzJoxc",
			},
		},
	}
	return infoList
}
