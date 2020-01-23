package info

func getMiscInfo() []Info {
	infoList := []Info{
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
	}
	return infoList
}
