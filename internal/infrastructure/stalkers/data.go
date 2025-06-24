package stalkers

var FirstNamesMale = []string{
	"Artyom", "Yuri", "Viktor", "Dmitri", "Alexei", "Mikhail", "Sergei", "Oleg", "Nikolai", "Andrei", "Ivan", "Grigori", "Vasili", "Pavel", "Stepan", "Bogdan",
	"Igor", "Anatoly", "Roman", "Leonid", "Taras", "Petro", "Oleksandr", "Mykola", "Volodymyr", "Denys", "Kyrylo", "Vadym", "Yaroslav", "Borys", "Ihor", "Stanislav",
	"Maxym", "Valeriy", "Vitaliy", "Artur", "Semen", "Georgiy", "Oleksiy", "Ruslan", "Hryhoriy", "Illia", "Arkadiy", "Evgen", "Fedor", "Gennadiy", "Ostap", "Zinoviy",
	"Andriy", "Marko",
}

var FirstNamesFemale = []string{
	"Anastasia", "Irina", "Elena", "Natalia", "Tatiana", "Svetlana", "Galina", "Maria", "Yulia", "Daria", "Oksana", "Olga", "Larysa", "Inna", "Iryna", "Kateryna",
	"Halyna", "Nadiya", "Taisiia", "Viktoria", "Liliya", "Nina", "Zinaida", "Marta", "Alla", "Valentina", "Yana", "Lyudmila", "Tamara", "Olesya", "Veronika", "Marichka",
	"Bogdana", "Lesya", "Oriana", "Roksana", "Lubov", "Evgeniya", "Zoya", "Milana", "Karina", "Sofiya", "Solomiya", "Olena", "Kseniya", "Polina", "Daryna", "Marianna",
	"Alyona", "Zlata",
}

var LastNamesMale = []string{
	"Petrov", "Sidorov", "Ivanov", "Morozov", "Pavlenko", "Shevchenko", "Bondarenko", "Tkachenko", "Kovalenko", "Melnyk", "Boyko", "Kravchenko", "Mazurenko", "Polishchuk",
	"Tymoshenko", "Samoylenko", "Volkov", "Tarasenko", "Kozlov", "Zinchenko", "Horobets", "Myronenko", "Rudenko", "Savchenko", "Pushkarev", "Babich", "Andriienko", "Filatov",
	"Gerasimov", "Korol", "Lebedev", "Maksymenko", "Nazarov", "Ponomarenko", "Radchenko", "Semenov", "Titov", "Ustinov", "Vasilenko", "Yermakov", "Zhuravlev", "Fedorenko",
	"Dorofeev", "Yakovlev", "Gavrilov", "Panchenko", "Reznikov", "Karpov", "Chumak", "Doronin",
}

var LastNamesFemale = []string{
	"Petrova", "Sidorova", "Ivanova", "Morozova", "Pavlenko", "Shevchenko", "Bondarenko", "Tkachenko", "Kovalenko", "Melnyka", "Boyko", "Kravchenko", "Mazurenko",
	"Polishchuka", "Tymoshenko", "Samoylenko", "Volkova", "Tarasenko", "Kozlova", "Zinchenko", "Horobetsa", "Myronenko", "Rudenko", "Savchenko", "Pushkareva", "Babicha",
	"Andriienko", "Filatova", "Gerasimova", "Korola", "Lebedeva", "Maksymenko", "Nazarova", "Ponomarenko", "Radchenko", "Semenova", "Titova", "Ustinova", "Vasilenko",
	"Yermakova", "Zhuravleva", "Fedorenko", "Dorofeeva", "Yakovleva", "Gavrilova", "Panchenko", "Reznikova", "Karpova", "Chumaka", "Doronina",
}

var GeneralMonikers = []string{
	"Ghost", "Viper", "Torch", "Fang", "Shadow", "Scarab", "Raven", "Frost", "Ash", "Grim", "Nomad", "Howl", "Echo", "Dust", "Blight", "Snare", "Drift",
	"Spook", "Moth", "Spire", "Gnarl", "Rust", "Stray", "Scar", "Witch", "Thorn", "Flint", "Soot", "Shiv", "Bark", "Brisk", "Snag", "Tusk", "Crack",
	"Fume", "Spike", "Crush", "Choke", "Scorn", "Gloom", "Flick", "Pox", "Splice", "Warp", "Glint", "Grub", "Fleck", "Wretch", "Whisk", "Slag", "Smudge",
	"Bulldog", "Ironclad", "Major", "Grim", "Lockjaw", "Torch", "Dragnet", "Vigil", "Crimson", "Warhound", "Sentinel", "Anvil", "Grizzly",
	"Breaker", "Steel", "Razor", "Ghosthound", "Marshal", "Rampart", "Bunker", "Grit", "Howler", "Smokescreen", "Krepost", "Sickle",
}

var MonolithMonikers = []string{
	"Voice of the Light", "Whisper of the Core", "Obelisk", "Seraphim", "Purity", "Echo of the Flame", "Bearer of Silence", "Disciple of the Glow",
	"Radiance", "Choir of Dust", "Watcher of the Singularity", "Hand of Truth", "Pilgrim of the Pulse", "Eye of the Signal", "Herald of the Path",
	"Covenant", "Brother of the Beacon", "Mother of Rebirth", "Lumen", "Heir of the Emission", "Call of the Monolith", "Saint of the Void",
	"Father of Clarity", "Ascendant", "Oathbound", "Speaker of the Cycle", "Shard of the Will", "Sentinel of the Divide", "Sister of the Spiral",
	"Lightbearer", "Fragment of the Core", "Preacher of Resonance", "Gaze of the Tower", "Pulse of Unity", "Warden of the Silence", "Flame of Convergence",
	"Ritekeeper", "Anointed", "Vessel of Truth", "Echo of the Singularity", "Blessed of the Zone", "The Bound One", "Scripture",
	"Choir of Light", "Oracle of the Heart", "Librarian of the Signal", "Fragment of the Mind", "Voice of the Source", "The Consecrated", "Shield of the Faith",
	"Crown of Emission", "Seeker of the Spark", "Path of Purity", "Watcher of the Glow", "The Pale Flame", "Embodiment", "Song of the Core",
	"Ash of the Prophet", "Bearer of Unity", "Witness of the Sky", "Fang of the Truth", "Serpent of the Light", "Monolith's Chosen", "Mask of Harmony",
	"Wanderer of the Dream", "Relic of the Zone", "The Undying Voice", "Bond of the Pillar", "Devotion", "Child of the Emission", "Wisdom of the Shard",
	"Gospel", "Star of the Divide", "Sigil of the Will", "Soul of the Spiral", "Scribe of the Pulse", "Veil of the Beacon", "Beaconborn",
	"Prophet of the Flame", "The Enlightened", "Speaker of the Tower", "Breath of Monolith", "The Quiet Psalm", "Sanctum", "Sentinel of Clarity",
	"Flesh of the Light", "Avatar of the Path", "The Reclaimed", "Omen of Convergence", "Caller of the Dawn", "The Risen Word", "Cradle of Illumination",
	"Shard of the Spiral", "The Absolute", "Memory of the Signal", "Essence of the Mind", "Martyr of the Zone", "Silence of the Prophet", "The Unshaken",
	"Womb of the Truth", "The Calling", "Light of the Ascension", "Transcendence", "Oblation", "Conviction", "Beacon", "Aegis", "Divinity",
	"Resonance", "Sanctuary", "Penance", "Ascension", "Incantation", "Revelation", "Consecration", "Faithbound", "Rapture", "Epiphany",
	"Silhouette", "Mandate", "Providence", "Genesis", "Anathema", "Convergence", "Atonement", "Dogma", "Sentience", "Oracle",
	"Sublime", "Chant", "Absolution", "Serenity", "Shrine", "Testament", "Monolithian", "Sacrament", "Vestige", "Vigil", "Abeyance",
	"Salvation", "Litany", "Obscurity", "Covenant", "Reliquary", "Ritual", "Illumination", "Pillar", "Credence", "Scripture", "Exaltation",
	"Halo", "Sanction", "Invocation",
}

var ForeignFullNames = []string{
	"Michał \"Gnat\" Domański",
	"Václav \"Stín\" Novák",
	"Ion \"Fier\" Stan",
	"Milan \"Tarča\" Petrović",
	"Zoran \"Bljesak\" Kovačević",
	"Anna \"Wilk\" Kowalczyk",
	"Eliška \"Sněžka\" Svobodová",
	"Elena \"Umbra\" Dobre",
	"Milica \"Zmija\" Jovanović",
	"Sara \"Zora\" Hadžić",
	"Tomasz \"Piorun\" Zieliński",
	"Adam \"Blaze\" Smith",
	"Kristijan \"Oblak\" Radić",
	"Bogdan \"Aether\" Ionescu",
	"Katarína \"Blesk\" Horváthová",
	"Emily \"Shade\" Johnson",
	"Joanna \"Róża\" Szymańska",
	"Dr. Jan \"Spektr\" Kubiš",
	"Dr. Mateusz \"Analizator\" Wójcik",
	"Dr. Emil \"Kvark\" Popa",
	"Dr. Ivana \"Beta\" Kostadinović",
	"Dr. Victoria \"Node\" Morgan",
	"Radek \"Kosa\" Bartosz",
	"Željko \"Rđa\" Nikolić",
	"Florin \"Cârlig\" Enache",
	"Agnieszka \"Szkło\" Maj",
	"Mirjana \"Lisica\" Tomić",
	"Brother \"Acolyte\" Marius",
	"Servant \"Silent Veil\"",
	"Sister \"Crystal Flame\"",
	"Witness \"Nebula\"",
	"Ondřej \"Průnik\" Malý",
	"Martin \"Reflektor\" Németh",
	"Gabriela \"Mir\" Andreeva",
	"Maja \"Haze\" Nowicka",
	"Krzysztof \"Cień\" Mazur",
	"Andrei \"Nomad\" Popescu",
	"David \"Ghost\" Walker",
	"Lucia \"Echo\" Dumitrescu",
	"Marta \"Zagadka\" Pawlak",
	"Kate \"Static\" Bell",
}

var ScientificTitles = []string{"Dr.", "Prof.", "Lab Tech"}

var MilitaryTitles = []string{"Pvt.", "Cpl.", "Sgt.", "Lt.", "Cpt.", "Maj.", "Col."}
