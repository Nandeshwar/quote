package info

import "time"

func getMyMeditaionInfo() []Info {
	infoList := []Info{
		{
			Title: "My Meditation experience - Think, chant, listen about Ram, Krishna",
			Info: ` 
Three thoughts generated strongly in my deep meditation today morning.
1. Talk, chant, listen, mediate on Ram and Krishna and their devotee so much that all positive, good souls gather together around me and help me to attend God.
Always think about him and do good things so that God will also want to enjoy my company.
If I think more about krishna and Ram, I will become like him one day.

2. I am a Gopi and helping Yasodha maiya and Nand baba since Krishna is not there. Whenever they cry for krishna, I tell them not to cry otherwise krishna will be unhappy.
I am trying to make them happy and smile every moment.

3. I am a Gopi and I invited krishna at my home and I am serving food to him. I am feeding him by my hand. I also went to many others gopi home where Krishna went. 

`,
			Link: []string{
				"",
			},
			CreationDate: time.Date(2019, 1, 30, 8, 11, 0, 0, time.Local),
		},
	}
	return infoList
}
