package info

import "time"

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
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
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
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
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
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 4
		{
			Title: "Gita in 12 minutes",
			Info: `Nice explanation of Bhagwat gita in 12 minutes. 
`,
			Link: []string{
				"https://www.youtube.com/watch?v=jnifjBM9dpM",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 5
		{
			Title: "Hanuman Mahima 1 hour video songs",
			Info: `Hanumnan ji bhakti towards Ram and Krishna. A beautiful song.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=N-RSr4ecs9M&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=9",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 6
		{
			Title: "Garud puran in 56 min",
			Info: `A beautiful song - Garud puran
`,
			Link: []string{
				"https://www.youtube.com/watch?v=Vb8-7DLSJuE&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=35",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 7
		{
			Title: "Bhajan - Govind Chale aao, Gopal chale aao",
			Info: `A beautiful bhajan of Krishna. Heart touching song.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=AEpGYCwutcc&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=49",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
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
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 9
		{
			Title: "Bhajan: Meri vinti hai radha rani",
			Info: `A beautiful bhajan of Radha Rani.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=3mmMm45rJpA&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=58",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
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
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 11
		{
			Title: "Bhajan: Itni Sakti dena hame",
			Info: ` Beautiful bhajan : Itni Sakti hame dena bhagwan
`,
			Link: []string{
				"https://www.youtube.com/watch?v=3EnLaJKhO2A&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=107",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 12
		{
			Title: "Bhajan: Jahe vidhi rakhhe ram",
			Info: ` Heart touching bhajan : Sita ram, Sita ram kahiye - Jahe vidhi rakkhe ram
`,
			Link: []string{
				"https://www.youtube.com/watch?v=GnvQH8PH8sg&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=118",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 13
		{
			Title: "Bhajan: Jai Radha Madhav",
			Info: ` A beautiful Bhajan, Jai Radha Madhav, Jai Kunj Bihari
`,
			Link: []string{
				"https://www.youtube.com/watch?v=cPTqAyqBAfc&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=119",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 14
		{
			Title: "Bhajan: Ae ri Sakhhi Mangal gao ri: Rishi Nitya Nand",
			Info: ` A nice dance and bhajan - Ae ri Sakkhi Mangal gao ri
`,
			Link: []string{
				"https://www.youtube.com/watch?v=Ju4DR1A_vPY&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=120",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 15
		{
			Title: "Mission Genius mind - Soul after death",
			Info: ` A story about girl's spirit
`,
			Link: []string{
				"https://www.youtube.com/watch?v=nPRKafjtGGs&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=127",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 16
		{
			Title: "Bhajan: hum van ke bashi",
			Info: ` Nice bhajan from Ramanand Sagar Ramayan
`,
			Link: []string{
				"https://www.youtube.com/watch?v=wOkFqjRwQYY&list=PL4I6x06f1KCt0uxfKkb3BgJZXNAjM3a-6&index=134",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 17
		{
			Title: "Bhajan: Radha ke maan me bas gaye",
			Info: ` Nice bhajan by Mridul krishna Sastri
`,
			Link: []string{
				"https://www.youtube.com/watch?v=OtjFIe84d0E",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		// info 17
		{
			Title: "Bhajan: O Kanha ab to murli ki madhur suna do taan",
			Info: ` Heart touching bhajan
`,
			Link: []string{
				"https://www.youtube.com/watch?v=XP9rlhzJoxc",
			},
			CreationDate: time.Date(2019, 1, 26, 0, 0, 0, 0, time.Local),
		},

		{
			Title: "Pancha Kanya",
			Info: ` 
1. Ahilya - Wife of Gautam Rishi in Ramayan
2. Tara - Wife of Banar Raj Baali in Ramayan
3. Mandodari - Wife of Ravan in Ramayan
4. Kunti - Monther of pandav in Mahabharat
5. Dropati - Wife of Pandav
`,
			Link: []string{
				"https://en.wikipedia.org/wiki/Panchakanya",
			},
			CreationDate: time.Date(2019, 1, 27, 0, 0, 0, 0, time.Local),
		},

		{
			Title: "Ved vyas :  Who is his Father and who are his sons",
			Info: ` 
Father: Parashara. He is son of Shaki Maharishi and Shakti Maharishi is son of Vasisthha - Guru of Ram
Son: Vidura, Shuka (Shuka Dev or Shukdev), Pandu, Dhritarashtra

`,
			Link: []string{
				"https://en.wikipedia.org/wiki/Vyasa",
			},
			CreationDate: time.Date(2019, 1, 27, 0, 0, 0, 0, time.Local),
		},

		{
			Title: "Navada Bhakti: Sabri",
			Info: ` 
Father: Parashara. 1. प्रथम भगति संतन्ह कर संगा। 

The first step to devotion (Bhakti) is to keep company of the saints (Satsang).

2. दुसरि रति मम कथा प्रसंगा॥

The second step is to enjoy listening to legends/discourses pertaining to the Lord.

3. गुरु पद पंकज सेवा तीसरि भगति अमान।

Selfless service to the Guru’s lotus feet without any pride is the third step.

4. चौथि भगति मम गुन गन करइ कपट तजि गान॥

The fourth step is to earnestly sing praises of the Lord’s virtues with a heart clear of guile, deceit or hypocrisy.

5. मंत्र जाप मम दृढ़ बिस्वासा। पंचम भजन सो बेद प्रकासा॥

Chanting My Name with steadfast faith is the fifth step as the Vedas reveal.

6. छठ दम सील बिरति बहु करमा। निरत निरंतर सज्जन धरमा॥

The sixth, is to practice self-control, good character, detachment from manifold activities and always follow the duties as good religious person.

7. सातवँ सम मोहि मय जग देखा। मोतें संत अधिक करि लेखा॥

The seventh step is to perceive the world as God Himself and regard the saints higher than the Lord.

8. आठवँ जथालाभ संतोषा। सपनेहुं नहिं देखइ परदोषा॥

The eighth, is a state (which one arrives at when one travels the first seven steps) where there is no desire left, but the gift of perfect peace and contentment with whatever one has. (In this state) one does not see fault in others, even in a dream.

9. नवम सरल सब सन छलहीना। मम भरोस हिय हरष न दीना॥

In this state, one has full faith in the Lord, and becomes (child-like) simple with no hypocrisy or deceit. The devotee has strong faith in the Lord with neither exaltation or depression in any life circumstance (but becomes equanimous).

नव महुं एकउ जिन्ह कें होई। नारि पुरूष सचराचर कोई॥
सोइ अतिसय प्रिय भामिनी मोरें। सकल प्रकार भगति दृढ़ तोरें॥

Shri Ram adds that Shabri’s Bhakti is perfectly complete. Yet if anyone were to have taken even one step towards devotion, out of all nine, he/she would be very dear to the Lord.

`,
			Link: []string{
				"https://www.speakingtree.in/allslides/navadha-bhakti---the-nine-steps-of-bhakti-devotion/267393",
			},
			CreationDate: time.Date(2019, 1, 27, 0, 0, 0, 0, time.Local),
		},

		{
			Title: "Ram naam ka mahatav - importance of Ram Naam",
			Info: ` 
Chanting Ram Naam and attaching mind to him will give bliss. It's meditation, devotion.
Hanuman Ji, Tulsidas, Shankar Ji, Meera, Prabhu Paad ji, Kripalu Ji Maharaj, Narad Ji and many are examples.
Chanting Ram naam attracts all good soul.
Negative soul and negative energy will go away.
Hanuman Ji will be happy.
Shankar ji will be happy.
Narad Ji will be happy.

Hare Ram Hare Ram
Ram Ram Hare hare
Hare Krishna Hare Krishna
Krishna Krishna Hare Hare

1.
चित्रकूट सब दिन बसत प्रभु सिय लखन समेत|.
राम नाम जप जापकहि तुलसी अभिमत देत||. 

प्रस्तुत दोहे में तुलसीदासजी कहते हैं कि भगवान श्रीरामचंद्रजी श्रीसीताजी और श्रीलक्ष्मणजी के साथ चित्रकूट में हमेशा निवास करते हैं| राम-नाम का जप करने वाले को वे मनचाहा फल प्रदान करते हैं|.

2.
पय अहार फल खाइ जपु राम नाम षट मास|.
सकल सुमंगल सिद्धि सब करतल तुलसीदासजी||.

प्रस्तुत दोहे में तुलसीदासजी कहते हैं कि छ: मास तक केवल दुग्ध का आहार करके अथवा केवल फल खाकर राम-नाम का जप करो| ऐसा करने से हर प्रकार के सुमंगल ओर सब सिद्धियां करतलगत हो जाएंगी,  अर्थात अपने-आप ही मिल जाएंगी|.

3.
राम नाम मनिदीप धरु जीह देहरीं द्वार|.
तुलसी भीतर बाहेरहूं जौं चाहसि उजिआर||.

प्रस्तुत दोहे में तुलसीदासजी कहते हैं कि यदि तू अन्दर और बाहर दोनों तरफ ज्ञान का प्रकाश (लौकिक एवं पारमार्थिक ज्ञान) चाहता है तो मुखरूपी दरवाजे की दहलीज पर रामनामरूपी मणिदीप रख दे, अर्थात जीभ के द्वारा अखण्ड रूप से श्रीरामजी के नाम का जप करता रहे|.

4.
नामु राम को कलपतरू कलि कल्यान निवासु|.
जो सुमिरत भयो भांग तें तुलसी तुलसीदास||.

प्रस्तुत दोहे में तुलसीदासजी कहते हैं कि भगवान श्रीरामचंद्रजी का नाम इस कलियुग में कल्पवृक्ष अर्थात मनचाहा फल प्रदान करने वाला है और कल्याण का निवास अर्थात मुक्ति का घर है,
जिसका स्मरण करने से तुलसीदास भांग से अर्थात विषय मद से भरी और दूसरों को भी विषयमद उपजाने वाली साधुओं द्वारा त्याज्य स्थिति से बदलकर तुलसी के समान निर्दोष, भगवान का प्यारा, सबका आदरणीय और जगत को पावन करने वाला हो गया|.

5.
बिगरी जनम अनेक की सुधरै अबहीं आजु|.
होहि राम को नाम जपु तुलसी तजि कुसमाजु||.

प्रस्तुत दोहे में तुलसीदासजी कहते हैं कि तू कुसंगीत और चित्त के सभी बुरे विचारों का त्याग करके प्रभु श्रीराम में ध्यान लगा और उनके नाम ‘राम’ का जप कर| ऐसा करने से तेरी अनेकों जन्मों की बिगड़ी हुई स्थिति अभी सुधर सकती है|.

`,
			Link: []string{
				"https://spiritualworld.co.in/bhagat-tulsidas-ji-dohawali/ram-naam-jap-ki-nahtta/",
			},
			CreationDate: time.Date(2019, 1, 29, 8, 15, 0, 0, time.Local),
		},

		{
			Title: "Hanuman Chalisa Meaning in Hindi",
			Info: ` 
श्री गुरु चरण सरोज रज, निज मन मुकुरु सुधारि।
बरनऊं रघुवर बिमल जसु, जो दायकु फल चारि।.
अर्थ- श्री गुरु महाराज के चरण कमलों की धूलि से अपने मन रूपी दर्पण को पवित्र करके श्री रघुवीर के निर्मल यश का वर्णन करता हूं, जो चारों फल धर्म, अर्थ, काम और मोक्ष को देने वाला है।>.  
>.  
****. 
बुद्धिहीन तनु जानिके, सुमिरो पवन-कुमार।
बल बुद्धि विद्या देहु मोहिं, हरहु कलेश विकार।.
अर्थ- हे पवन कुमार! मैं आपको सुमिरन करता हूं। आप तो जानते ही हैं कि मेरा शरीर और बुद्धि निर्बल है। मुझे शारीरिक बल, सद्‍बुद्धि एवं ज्ञान दीजिए और मेरे दुखों व दोषों का नाश कार दीजिए।.

****.
जय हनुमान ज्ञान गुण सागर, जय कपीस तिहुं लोक उजागर॥1॥.
अर्थ- श्री हनुमान जी! आपकी जय हो। आपका ज्ञान और गुण अथाह है। हे कपीश्वर! आपकी जय हो! तीनों लोकों, स्वर्ग लोक, भूलोक और पाताल लोक में आपकी कीर्ति है।.

****.

राम दूत अतुलित बलधामा, अंजनी पुत्र पवन सुत नामा॥2॥.
अर्थ- हे पवनसुत अंजनी नंदन! आपके समान दूसरा बलवान नहीं है।.

****. 
महावीर विक्रम बजरंगी, कुमति निवार सुमति के संगी॥3॥.
अर्थ- हे महावीर बजरंग बली!आप विशेष पराक्रम वाले है। आप खराब बुद्धि को दूर करते है, और अच्छी बुद्धि वालों के साथी, सहायक है।.

****. 
कंचन बरन बिराज सुबेसा, कानन कुण्डल कुंचित केसा॥4॥.
अर्थ- आप सुनहले रंग, सुन्दर वस्त्रों, कानों में कुण्डल और घुंघराले बालों से सुशोभित हैं।.

****. 
हाथबज्र और ध्वजा विराजे, कांधे मूंज जनेऊ साजै॥5॥.
अर्थ- आपके हाथ में बज्र और ध्वजा है और कन्धे पर मूंज के जनेऊ की शोभा है।.

****.
शंकर सुवन केसरी नंदन, तेज प्रताप महा जग वंदन॥6॥.
अर्थ- शंकर के अवतार! हे केसरी नंदन आपके पराक्रम और महान यश की संसार भर में वन्दना होती है।.

****. 
विद्यावान गुणी अति चातुर, राम काज करिबे को आतुर॥7॥.
अर्थ- आप प्रकान्ड विद्या निधान है, गुणवान और अत्यन्त कार्य कुशल होकर श्री राम के काज करने के लिए आतुर रहते है।.

****. 
प्रभु चरित्र सुनिबे को रसिया, राम लखन सीता मन बसिया॥8॥.
अर्थ- आप श्री राम चरित सुनने में आनन्द रस लेते है। श्री राम, सीता और लखन आपके हृदय में बसे रहते है।.

****.
सूक्ष्म रूप धरि सियहिं दिखावा, बिकट रूप धरि लंक जरावा॥9॥.
अर्थ- आपने अपना बहुत छोटा रूप धारण करके सीता जी को दिखलाया और भयंकर रूप करके लंका को जलाया।.

****.
भीम रूप धरि असुर संहारे, रामचन्द्र के काज संवारे॥10॥.
अर्थ- आपने विकराल रूप धारण करके राक्षसों को मारा और श्री रामचन्द्र जी के उद्‍देश्यों को सफल कराया।.

****. 
लाय सजीवन लखन जियाये, श्री रघुवीर हरषि उर लाये॥11॥.
अर्थ- आपने संजीवनी बूटी लाकर लक्ष्मण जी को जिलाया जिससे श्री रघुवीर ने हर्षित होकर आपको हृदय से लगा लिया।.

****.
रघुपति कीन्हीं बहुत बड़ाई, तुम मम प्रिय भरत सम भाई॥12॥.

अर्थ- श्री रामचन्द्र ने आपकी बहुत प्रशंसा की और कहा कि तुम मेरे भरत जैसे प्यारे भाई हो।.


****.
सहस बदन तुम्हरो जस गावैं। अस कहि श्रीपति कंठ लगावैं॥13॥.
अर्थ- श्री राम ने आपको यह कहकर हृदय से लगा लिया की तुम्हारा यश हजार मुख से सराहनीय है।.

****. 
सनकादिक ब्रह्मादि मुनीसा,  नारद, सारद सहित अहीसा॥14॥.
अर्थ-  श्री सनक, श्री सनातन, श्री सनन्दन, श्री सनत्कुमार आदि मुनि ब्रह्मा आदि देवता नारद जी, सरस्वती जी, शेषनाग जी सब आपका गुण गान करते है।.

****.
जम कुबेर दिगपाल जहां ते, कबि कोबिद कहि सके कहां ते॥15॥.
अर्थ- यमराज, कुबेर आदि सब दिशाओं के रक्षक, कवि विद्वान, पंडित या कोई भी आपके यश का पूर्णतः वर्णन नहीं कर सकते।.

****. 
तुम उपकार सुग्रीवहि कीन्हा, राम मिलाय राजपद दीन्हा॥16॥.
अर्थ- आपने सुग्रीव जी को श्रीराम से मिलाकर उपकार किया, जिसके कारण वे राजा बने।.

****. 
तुम्हरो मंत्र विभीषण माना, लंकेस्वर भए सब जग जाना॥17॥.
अर्थ- आपके उपदेश का विभिषण जी ने पालन किया जिससे वे लंका के राजा बने, इसको सब संसार जानता है।.

****. 
जुग सहस्त्र जोजन पर भानू, लील्यो ताहि मधुर फल जानू॥18॥.
अर्थ- जो सूर्य इतने योजन दूरी पर है कि उस पर पहुंचने के लिए हजार युग लगे। दो हजार योजन की दूरी पर स्थित सूर्य को आपने एक मीठा फल समझकर निगल लिया।.

**** 
प्रभु मुद्रिका मेलि मुख माहि, जलधि लांघि गये अचरज नाहीं॥19॥.
अर्थ- आपने श्री रामचन्द्र जी की अंगूठी मुंह में रखकर समुद्र को लांघ लिया, इसमें कोई आश्चर्य नहीं है।.

****.
दुर्गम काज जगत के जेते, सुगम अनुग्रह तुम्हरे तेते॥20॥.
अर्थ- संसार में जितने भी कठिन से कठिन काम हो, वो आपकी कृपा से सहज हो जाते है।.

****.

राम दुआरे तुम रखवारे, होत न आज्ञा बिनु पैसा रे॥21॥.
अर्थ-
श्री रामचन्द्र जी के द्वार के आप रखवाले है, जिसमें आपकी आज्ञा बिना किसी को प्रवेश नहीं मिलता अर्थात् आपकी प्रसन्नता के बिना राम कृपा दुर्लभ है।.
****.

सब सुख लहै तुम्हारी सरना, तुम रक्षक काहू को डरना ॥22॥.
अर्थ-
जो भी आपकी शरण में आते है, उस सभी को आनन्द प्राप्त होता है, और जब आप रक्षक है, तो फिर किसी का डर नहीं रहता।.

****.
आपन तेज सम्हारो आपै, तीनों लोक हांक तें कांपै॥23॥.
अर्थ-
आपके सिवाय आपके वेग को कोई नहीं रोक सकता, आपकी गर्जना से तीनों लोक कांप जाते है।.

****.
भूत पिशाच निकट नहिं आवै, महावीर जब नाम सुनावै॥24॥.
अर्थ-
जहां महावीर हनुमान जी का नाम सुनाया जाता है, वहां भूत, पिशाच पास भी नहीं फटक सकते।.


****.
नासै रोग हरै सब पीरा, जपत निरंतर हनुमत बीरा ॥25॥.
अर्थ-
वीर हनुमान जी! आपका निरंतर जप करने से सब रोग चले जाते है और सब पीड़ा मिट जाती है।.

****.
संकट तें हनुमान छुड़ावै, मन क्रम बचन ध्यान जो लावै॥26॥.
अर्थ-
हे हनुमान जी! विचार करने में, कर्म करने में और बोलने में, जिनका ध्यान आपमें रहता है, उनको सब
संकटों से आप छुड़ाते है।.

****.
सब पर राम तपस्वी राजा, तिनके काज सकल तुम साजा॥27॥.
अर्थ-
तपस्वी राजा श्री रामचन्द्र जी सबसे श्रेष्ठ है, उनके सब कार्यों को आपने सहज में कर दिया।.

****.
और मनोरथ जो कोइ लावै, सोई अमित जीवन फल पावै॥28॥.
अर्थ-
जिस पर आपकी कृपा हो, वह कोई भी अभिलाषा करें तो उसे ऐसा फल मिलता है जिसकी जीवन में कोई सीमा नहीं होती।.

****.
चारों जुग परताप तुम्हारा, है परसिद्ध जगत उजियारा॥29॥.
अर्थ-
चारो युगों सतयुग, त्रेता, द्वापर तथा कलियुग में आपका यश फैला हुआ है, जगत में आपकी कीर्ति सर्वत्र प्रकाशमान है।.

****.
साधु सन्त के तुम रखवारे, असुर निकंदन राम दुलारे॥30॥.

अर्थ-
हे श्री राम के दुलारे! आप सज्जनों की रक्षा करते है और दुष्टों का नाश करते है।.

****.
अष्ट सिद्धि नौ निधि के दाता, अस बर दीन जानकी माता॥31॥.

अर्थ-
आपको माता श्री जानकी से ऐसा वरदान मिला हुआ है, जिससे आप किसी को भी आठों सिद्धियां और नौ निधियां दे सकते
है।
1.) अणिमा- जिससे साधक किसी को दिखाई नहीं पड़ता और कठिन से कठिन पदार्थ में प्रवेश कर जाता है।
2.) महिमा- जिसमें योगी अपने को बहुत बड़ा बना देता है।
3.) गरिमा- जिससे साधक अपने को चाहे जितना भारी बना लेता है।
4.) लघिमा- जिससे जितना चाहे उतना हल्का बन जाता है।
5.) प्राप्ति- जिससे इच्छित पदार्थ की प्राप्ति होती है।
6.) प्राकाम्य- जिससे इच्छा करने पर वह पृथ्वी में समा सकता है, आकाश में उड़ सकता है।
7.) ईशित्व- जिससे सब पर शासन का सामर्थ्य हो जाता है।
8.) वशित्व- जिससे दूसरों को वश में किया जाता है।

****.

राम रसायन तुम्हरे पासा, सदा रहो रघुपति के दासा॥32॥.

अर्थ-
आप निरंतर श्री रघुनाथ जी की शरण में रहते है, जिससे आपके पास बुढ़ापा और असाध्य रोगों के नाश के लिए राम नाम औषधि है।.
****.
तुम्हरे भजन राम को पावै, जनम जनम के दुख बिसरावै॥33॥.

अर्थ-
आपका भजन करने से श्री राम जी प्राप्त होते है और जन्म जन्मांतर के दुख दूर होते है।.
****.
अन्त काल रघुबर पुर जाई, जहां जन्म हरि भक्त कहाई॥34॥.

अर्थ-
अंत समय श्री रघुनाथ जी के धाम को जाते है और यदि फिर भी जन्म लेंगे तो भक्ति करेंगे और श्री राम भक्त कहलाएंगे।.
****.
और देवता चित न धरई, हनुमत सेई सर्व सुख करई॥35॥.

अर्थ-
हे हनुमान जी! आपकी सेवा करने से सब प्रकार के सुख मिलते है, फिर अन्य किसी देवता की आवश्यकता नहीं रहती।.
****.
संकट कटै मिटै सब पीरा, जो सुमिरै हनुमत बलबीरा॥36॥.

अर्थ-
हे वीर हनुमान जी! जो आपका सुमिरन करता रहता है, उसके सब संकट कट जाते है और सब पीड़ा मिट जाती है।.
****.
जय जय जय हनुमान गोसाईं, कृपा करहु गुरु देव की नाई॥37॥.

अर्थ-
हे स्वामी हनुमान जी! आपकी जय हो, जय हो, जय हो! आप मुझ पर कृपालु श्री गुरु जी के समान कृपा कीजिए।.
****.
जो सत बार पाठ कर कोई, छूटहि बंदि महा सुख होई॥38॥.

अर्थ-
जो कोई इस हनुमान चालीसा का सौ बार पाठ करेगा वह सब बंधनों से छूट जाएगा और उसे परमानन्द मिलेगा।.
****.
जो यह पढ़ै हनुमान चालीसा, होय सिद्धि साखी गौरीसा॥39॥.

अर्थ-
भगवान शंकर ने यह हनुमान चालीसा लिखवाया, इसलिए वे साक्षी है, कि जो इसे पढ़ेगा उसे निश्चय ही सफलता प्राप्त होगी।.
****.
तुलसीदास सदा हरि चेरा, कीजै नाथ हृदय मंह डेरा॥40॥.

अर्थ-
हे नाथ हनुमान जी! तुलसीदास सदा ही श्री राम का दास है। इसलिए आप उसके हृदय में निवास कीजिए।.
****.
पवन तनय संकट हरन, मंगल मूरति रूप। राम लखन सीता सहित, हृदय बसहु सूरभूप॥.

अर्थ-
हे संकट मोचन पवन कुमार! आप आनंद मंगलों के स्वरूप हैं। हे देवराज! आप श्री राम, सीता जी और लक्ष्मण सहित मेरे हृदय में निवास कीजिए।.

`,
			Link: []string{
				"https://www.jagranjunction.com/religious/meaning-of-hanuman-chalisa-in-hindi/",
			},
			CreationDate: time.Date(2019, 1, 29, 8, 15, 0, 0, time.Local),
		},

		{
			Title: "Ramayan Ramanand Sagar: Sita, Laxman love towards Ram",
			Info: `
Sita always wants to see Ram.
Sita always want to make Ram Happy.
Laxman Ji has not imagined himself without Ram.
`,
			Link: []string{
				"https://www.youtube.com/watch?v=AvXufbbrclY",
			},
			CreationDate: time.Date(2019, 2, 2, 1, 0, 0, 0, time.Local),
		},
	}
	return infoList
}
