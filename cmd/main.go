package main

import (
	aho "go-search/pkg/ahocorasick"
	"go-search/pkg/binary"
)

func main() {
	patterns := []string{"eat", "drink"}

	text := `
him through the Enchanted Forest, where the trees seemed to whisper ancient secrets as the wind rustled their leaves. He encountered friendly woodland creatures, like the playful squirrels and the graceful deer, and felt a sense of harmony with nature.
	
	As he ventured deeper into the forest, the atmosphere turned eerie, and the shadows seemed to dance ominously. But Ethan pressed on, undeterred by fear, guided by an unshakable determination to uncover the truth about the legends he had heard.
	
	In a clearing among the dense trees, he stumbled upon an ethereal waterfall cascading down from the cliff's edge. The water sparkled like a thousand diamonds, and the sight filled him with awe. Little did he know that the mystical waterfall held a secret that would change his life forever.
	
	Approaching the waterfall, he noticed a glimmering object embedded in the rock. Brushing away the moss, he revealed a glowing medallion with ancient symbols etched onto its surface. The moment he touched it, a surge of energy coursed through his veins, and a voice echoed in his mind, guiding him to the next part of his quest.
	
	With newfound confidence, Ethan continued his journey, crossing vast plains, scaling treacherous peaks, and navigating through treacherous swamps. He faced challenges and encountered benevolent beings and cunning adversaries along the way.
	
	His travels took him to forgotten ruins and abandoned temples, where he deciphered ancient riddles and discovered ancient relics of great power. Each artifact he unearthed brought him closer to understanding the prophecies that foretold a cataclysmic event threatening Aldoria's existence.
	
	As his reputation as a daring adventurer grew, so did his circle of friends. He met fellow travelers and formed a fellowship united by a common purpose. Together, they faced the malevolent forces that sought to plunge the world into darkness.
	
	Through hardships, losses, and triumphs, Ethan and his companions became legends in their own right. They stood shoulder to shoulder, confronting the darkness and protecting the innocent. Their bond forged in the fires of adversity was unbreakable.
	
	Finally, on the eve of the prophesied cataclysm, Ethan and his fellowship faced the ultimate challenge – a confrontation with the ancient entity that sought to bring chaos to Aldoria. In a battle of epic proportions, they fought with valor and determination, utilizing the ancient artifacts' powers.
	
	In the end, it was Ethan's resilience and indomitable spirit that tipped the scales in their favor. With a final, powerful blow, he banished the malevolent entity back to the depths of oblivion, ensuring Aldoria's safety.
	
	The tale of Ethan and his grand adventure spread far and wide, becoming a legend for generations to come. His bravery and selflessness inspired countless souls to follow their dreams and to believe in the power of hope and unity.
	
	And so, in the distant lands of Aldoria, where lush green meadows meet towering mountains and serene seas, the legacy of Ethan lives on, forever etched in the hearts of the people, reminding them that within each one of them lies the potential to be a hero.
	
	Please note that this text is a fictional story created for illustrative purposes, and any resemblance to real places, events, or individuals is purely coincidental. If you have any specific topic or theme in mind for the text, feel free to let me know, and I can generate a different story accordingly.
	In the distant lands of Aldoria, a lush green paradise nestled between the majestic mountains and the glistening sea, a tale of heroism and adventure begins. The quaint village of Evergreen, with its thatched-roof cottages and colorful gardens, was home to a young and curious lad named Ethan.

	Ethan had always dreamt of exploring the uncharted territories beyond the village's borders. He would spend countless hours listening to the stories told by the elderly villagers about the mysterious creatures that roamed the dark forests and the hidden treasures that lay dormant in ancient ruins.
	
	One sunny morning, while the village bustled with life, Ethan decided that it was time to embark on his grand adventure. With a satchel slung over his shoulder and a walking stick in hand, he set off into the unknown, bidding farewell to his tearful parents and waving goodbye to his friends.
	
	The first leg of his journey led him through the Enchanted Forest, where the trees seemed to whisper ancient secrets as the wind rustled their leaves. He encountered friendly woodland creatures, like the playful squirrels and the graceful deer, and felt a sense of harmony with nature.
	
	As he ventured deeper into the forest, the atmosphere turned eerie, and the shadows seemed to dance ominously. But Ethan pressed on, undeterred by fear, guided by an unshakable determination to uncover the truth about the legends he had heard.
	
	In a clearing among the dense trees, he stumbled upon an ethereal waterfall cascading down from the cliff's edge. The water sparkled like a thousand diamonds, and the sight filled him with awe. Little did he know that the mystical waterfall held a secret that would change his life forever.
	
	Approaching the waterfall, he noticed a glimmering object embedded in the rock. Brushing away the moss, he revealed a glowing medallion with ancient symbols etched onto its surface. The moment he touched it, a surge of energy coursed through his veins, and a voice echoed in his mind, guiding him to the next part of his quest.
	
	With newfound confidence, Ethan continued his journey, crossing vast plains, scaling treacherous peaks, and navigating through treacherous swamps. He faced challenges and encountered benevolent beings and cunning adversaries along the way.
	
	His travels took him to forgotten ruins and abandoned temples, where he deciphered ancient riddles and discovered ancient relics of great power. Each artifact he unearthed brought him closer to understanding the prophecies that foretold a cataclysmic event threatening Aldoria's existence.
	
	As his reputation as a daring adventurer grew, so did his circle of friends. He met fellow travelers and formed a fellowship united by a common purpose. Together, they faced the malevolent forces that sought to plunge the world into darkness.
	
	Through hardships, losses, and triumphs, Ethan and his companions became legends in their own right. They stood shoulder to shoulder, confronting the darkness and protecting the innocent. Their bond forged in the fires of adversity was unbreakable.
	
	Finally, on the eve of the prophesied cataclysm, Ethan and his fellowship faced the ultimate challenge – a confrontation with the ancient entity that sought to bring chaos to Aldoria. In a battle of epic proportions, they fought with valor and determination, utilizing the ancient artifacts' powers.
	
	In the end, it was Ethan's resilience and indomitable spirit that tipped the scales in their favor. With a final, powerful blow, he banished the malevolent entity back to the depths of oblivion, ensuring Aldoria's safety.
	
	The tale of Ethan and his grand adventure spread far and wide, becoming a legend for generations to come. His bravery and selflessness inspired countless souls to follow their dreams and to believe in the power of hope and unity.
	
	And so, in the distant lands of Aldoria, where lush green meadows meet towering mountains and serene seas, the legacy of Ethan lives on, forever etched in the hearts of the people, reminding them that within each one of them lies the potential to be a hero.
	
	Please note that this text is a fictional story created for illustrative purposes, and any resemblance to real places, events, or individuals is purely coincidental. If you have any specific topic or theme in mind for the text, feel free to let me know, and I can generate a different story accordingly.
 his journey, crossing vast plains, scaling treacherous peaks, and navigating through treacherous swamps. He faced challenges and encountered benevolent beings and cunning adversaries along the way.
	
	His travels took him to forgotten ruins and abandoned temples, where he deciphered ancient riddles and discovered ancient relics of great power. Each artifact he unearthed brought him closer to understanding the prophecies that foretold a cataclysmic event threatening Aldoria's existence.
`
	aho.Search(text, patterns)
	binary.Search(text, patterns)
}
