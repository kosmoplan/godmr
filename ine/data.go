package ine

// Downloaded from ine.es and tweaked to include
// common mispellings I'm seeing all over the place. I
// plan to get something more clever in place but for
// now I'm happy to keep going like this.
var Data = []byte (`
CPRO,PRO,CMUN,DC,NOMBRE
01,Álava,001,4,Vitoria
01,Álava,001,4,Alegría-Dulantzi
01,Álava,002,9,Amurrio
01,Álava,049,3,Añana
01,Álava,003,5,Aramaio
01,Álava,006,6,Armiñón
01,Álava,037,6,Arraia-Maeztu
01,Álava,008,8,Arrazua-Ubarrundia
01,Álava,004,0,Artziniega
01,Álava,009,1,Asparrena
01,Álava,010,5,Ayala/Aiara
01,Álava,011,2,Baños de Ebro/Mañueta
01,Álava,013,3,Barrundia
01,Álava,014,8,Berantevilla
01,Álava,016,4,Bernedo
01,Álava,017,0,Campezo/Kanpezu
01,Álava,021,0,Elburgo/Burgelu
01,Álava,022,5,Elciego
01,Álava,023,1,Elvillar/Bilar
01,Álava,046,8,Erriberagoitia/Ribera Alta
01,Álava,056,5,Harana/Valle de Arana
01,Álava,901,5,Iruña Oka/Iruña de Oca
01,Álava,027,8,Iruraiz-Gauna
01,Álava,019,9,Kripan
01,Álava,020,3,Kuartango
01,Álava,028,4,Labastida/Bastida
01,Álava,030,1,Lagrán
01,Álava,031,8,Laguardia
01,Álava,032,3,Lanciego/Lantziego
01,Álava,902,0,Lantarón
01,Álava,033,9,Lapuebla de Labarca
01,Álava,036,0,Laudio/Llodio
01,Álava,058,7,Legutio
01,Álava,034,4,Leza
01,Álava,039,5,Moreda de Álava
01,Álava,041,6,Navaridas
01,Álava,042,1,Okondo
01,Álava,043,7,Oyón-Oion
01,Álava,044,2,Peñacerrada-Urizaharra
01,Álava,047,4,Ribera Baja/Erribera Beitia
01,Álava,051,3,Salvatierra/Agurain
01,Álava,052,8,Samaniego
01,Álava,053,4,San Millán/Donemiliaga
01,Álava,054,9,Urkabustaiz
01,Álava,055,2,Valdegovía/Gaubea
01,Álava,057,1,Villabuena de Álava/Eskuernaga
01,Álava,059,0,Vitoria-Gasteiz
01,Álava,060,4,Yécora/Iekora
01,Álava,061,1,Zalduondo
01,Álava,062,6,Zambrana
01,Álava,018,6,Zigoitia
01,Álava,063,2,Zuia
02,Albacete,001,9,Abengibre
02,Albacete,002,4,Alatoz
02,Albacete,003,0,Albacete
02,Albacete,004,5,Albatana
02,Albacete,005,8,Alborea
02,Albacete,006,1,Alcadozo
02,Albacete,007,7,Alcalá del Júcar
02,Albacete,008,3,Alcaraz
02,Albacete,009,6,Almansa
02,Albacete,010,0,Alpera
02,Albacete,011,7,Ayna
02,Albacete,012,2,Balazote
02,Albacete,014,3,"Ballestero, El"
02,Albacete,013,8,Balsa de Ves
02,Albacete,015,6,Barrax
02,Albacete,016,9,Bienservida
02,Albacete,017,5,Bogarra
02,Albacete,018,1,Bonete
02,Albacete,019,4,"Bonillo, El"
02,Albacete,020,8,Carcelén
02,Albacete,021,5,Casas de Juan Núñez
02,Albacete,022,0,Casas de Lázaro
02,Albacete,023,6,Casas de Ves
02,Albacete,024,1,Casas-Ibáñez
02,Albacete,025,4,Caudete
02,Albacete,026,7,Cenizate
02,Albacete,029,2,Chinchilla de Monte-Aragón
02,Albacete,027,3,Corral-Rubio
02,Albacete,028,9,Cotillas
02,Albacete,030,6,Elche de la Sierra
02,Albacete,031,3,Férez
02,Albacete,032,8,Fuensanta
02,Albacete,033,4,Fuente-Álamo
02,Albacete,034,9,Fuentealbilla
02,Albacete,035,2,"Gineta, La"
02,Albacete,036,5,Golosalvo
02,Albacete,037,1,Hellín
02,Albacete,038,7,"Herrera, La"
02,Albacete,039,0,Higueruela
02,Albacete,040,4,Hoya-Gonzalo
02,Albacete,041,1,Jorquera
02,Albacete,042,6,Letur
02,Albacete,043,2,Lezuza
02,Albacete,044,7,Liétor
02,Albacete,045,0,Madrigueras
02,Albacete,046,3,Mahora
02,Albacete,047,9,Masegoso
02,Albacete,048,5,Minaya
02,Albacete,049,8,Molinicos
02,Albacete,050,1,Montalvos
02,Albacete,051,8,Montealegre del Castillo
02,Albacete,052,3,Motilleja
02,Albacete,053,9,Munera
02,Albacete,054,4,Navas de Jorquera
02,Albacete,055,7,Nerpio
02,Albacete,056,0,Ontur
02,Albacete,057,6,Ossa de Montiel
02,Albacete,058,2,Paterna del Madera
02,Albacete,060,9,Peñas de San Pedro
02,Albacete,059,5,Peñascosa
02,Albacete,061,6,Pétrola
02,Albacete,062,1,Povedilla
02,Albacete,901,0,Pozo Cañada
02,Albacete,063,7,Pozohondo
02,Albacete,064,2,Pozo-Lorente
02,Albacete,065,5,Pozuelo
02,Albacete,066,8,"Recueja, La"
02,Albacete,067,4,Riópar
02,Albacete,068,0,Robledo
02,Albacete,069,3,"Roda, La"
02,Albacete,070,7,Salobre
02,Albacete,071,4,San Pedro
02,Albacete,072,9,Socovos
02,Albacete,073,5,Tarazona de la Mancha
02,Albacete,074,0,Tobarra
02,Albacete,075,3,Valdeganga
02,Albacete,076,6,Vianos
02,Albacete,077,2,Villa de Ves
02,Albacete,078,8,Villalgordo del Júcar
02,Albacete,079,1,Villamalea
02,Albacete,080,5,Villapalacios
02,Albacete,081,2,Villarrobledo
02,Albacete,082,7,Villatoya
02,Albacete,083,3,Villavaliente
02,Albacete,084,8,Villaverde de Guadalimar
02,Albacete,085,1,Viveros
02,Albacete,086,4,Yeste
03,Alicante,000,0,Alcoy
03,Alicante,000,0,Denia
03,Alicante,000,0,Elche
03,Alicante,000,0,Alicante
03,Alicante,001,5,Adsubia
03,Alicante,002,0,Agost
03,Alicante,003,6,Agres
03,Alicante,004,1,Aigües
03,Alicante,005,4,Albatera
03,Alicante,006,7,Alcalalí
03,Alicante,007,3,Alcocer de Planes
03,Alicante,008,9,Alcoleja
03,Alicante,009,2,Alcoy/Alcoi
03,Alicante,010,6,Alfafara
03,Alicante,011,3,"Alfàs del Pi, l'"
03,Alicante,012,8,Algorfa
03,Alicante,013,4,Algueña
03,Alicante,014,9,Alicante/Alacant
03,Alicante,015,2,Almoradí
03,Alicante,016,5,Almudaina
03,Alicante,017,1,"Alqueria d'Asnar, l'"
03,Alicante,018,7,Altea
03,Alicante,019,0,Aspe
03,Alicante,020,4,Balones
03,Alicante,021,1,Banyeres de Mariola
03,Alicante,022,6,Benasau
03,Alicante,023,2,Beneixama
03,Alicante,024,7,Benejúzar
03,Alicante,025,0,Benferri
03,Alicante,026,3,Beniarbeig
03,Alicante,027,9,Beniardá
03,Alicante,028,5,Beniarrés
03,Alicante,030,2,Benidoleig
03,Alicante,031,9,Benidorm
03,Alicante,032,4,Benifallim
03,Alicante,033,0,Benifato
03,Alicante,029,8,Benigembla
03,Alicante,034,5,Benijófar
03,Alicante,035,8,Benilloba
03,Alicante,036,1,Benillup
03,Alicante,037,7,Benimantell
03,Alicante,038,3,Benimarfull
03,Alicante,039,6,Benimassot
03,Alicante,040,0,Benimeli
03,Alicante,041,7,Benissa
03,Alicante,042,2,"Benitachell/Poble Nou de Benitatxell, el"
03,Alicante,043,8,Biar
03,Alicante,044,3,Bigastro
03,Alicante,045,6,Bolulla
03,Alicante,046,9,Busot
03,Alicante,049,4,Callosa de Segura
03,Alicante,048,1,Callosa d'En Sarrià
03,Alicante,047,5,Calp
03,Alicante,050,7,"Campello, el"
03,Alicante,051,4,"Campo de Mirra/Camp de Mirra, el"
03,Alicante,052,9,Cañada
03,Alicante,053,5,Castalla
03,Alicante,054,0,Castell de Castells
03,Alicante,075,9,"Castell de Guadalest, el"
03,Alicante,055,3,Catral
03,Alicante,056,6,Cocentaina
03,Alicante,057,2,Confrides
03,Alicante,058,8,Cox
03,Alicante,059,1,Crevillent
03,Alicante,061,2,Daya Nueva
03,Alicante,062,7,Daya Vieja
03,Alicante,063,3,Dénia
03,Alicante,064,8,Dolores
03,Alicante,065,1,Elche/Elx
03,Alicante,066,4,Elda
03,Alicante,067,0,Facheca
03,Alicante,068,6,Famorca
03,Alicante,069,9,Finestrat
03,Alicante,077,8,"Fondó de les Neus, el/Hondón de las Nieves"
03,Alicante,070,3,Formentera del Segura
03,Alicante,072,5,Gaianes
03,Alicante,071,0,Gata de Gorgos
03,Alicante,073,1,Gorga
03,Alicante,074,6,Granja de Rocamora
03,Alicante,076,2,Guardamar del Segura
03,Alicante,078,4,Hondón de los Frailes
03,Alicante,079,7,Ibi
03,Alicante,080,1,Jacarilla
03,Alicante,082,3,Jávea/Xàbia
03,Alicante,083,9,Jijona/Xixona
03,Alicante,085,7,Llíber
03,Alicante,084,4,"Lorcha/Orxa, l'"
03,Alicante,086,0,Millena
03,Alicante,088,2,Monforte del Cid
03,Alicante,089,5,Monóvar/Monòver
03,Alicante,903,7,"Montesinos, Los"
03,Alicante,091,6,Murla
03,Alicante,092,1,Muro de Alcoy
03,Alicante,090,9,Mutxamel
03,Alicante,093,7,Novelda
03,Alicante,094,2,"Nucia, la"
03,Alicante,095,5,Ondara
03,Alicante,096,8,Onil
03,Alicante,097,4,Orba
03,Alicante,099,3,Orihuela
03,Alicante,098,0,Orxeta
03,Alicante,100,7,Parcent
03,Alicante,101,4,Pedreguer
03,Alicante,102,9,Pego
03,Alicante,103,5,Penàguila
03,Alicante,104,0,Petrer
03,Alicante,902,1,Pilar de la Horadada
03,Alicante,105,3,"Pinós, el/Pinoso"
03,Alicante,106,6,Planes
03,Alicante,901,6,"Poblets, els"
03,Alicante,107,2,Polop
03,Alicante,060,5,Quatretondeta
03,Alicante,109,1,Rafal
03,Alicante,110,5,"Ràfol d'Almúnia, el"
03,Alicante,111,2,Redován
03,Alicante,112,7,Relleu
03,Alicante,113,3,Rojales
03,Alicante,114,8,"Romana, la"
03,Alicante,115,1,Sagra
03,Alicante,116,4,Salinas
03,Alicante,118,6,San Fulgencio
03,Alicante,904,2,San Isidro
03,Alicante,120,3,San Miguel de Salinas
03,Alicante,122,5,San Vicente del Raspeig/Sant Vicent del Raspeig
03,Alicante,117,0,Sanet y Negrals
03,Alicante,119,9,Sant Joan d'Alacant
03,Alicante,121,0,Santa Pola
03,Alicante,123,1,Sax
03,Alicante,124,6,Sella
03,Alicante,125,9,Senija
03,Alicante,127,8,Tàrbena
03,Alicante,128,4,Teulada
03,Alicante,129,7,Tibi
03,Alicante,130,1,Tollos
03,Alicante,131,8,Tormos
03,Alicante,132,3,"Torremanzanas/Torre de les Maçanes, la"
03,Alicante,133,9,Torrevieja
03,Alicante,134,4,"Vall d'Alcalà, la"
03,Alicante,136,0,Vall de Gallinera
03,Alicante,137,6,"Vall de Laguar, la"
03,Alicante,135,7,"Vall d'Ebo, la"
03,Alicante,138,2,"Verger, el"
03,Alicante,139,5,"Villajoyosa/Vila Joiosa, la"
03,Alicante,140,9,Villena
03,Alicante,081,8,Xaló
04,Almería,000,0,Almeria
04,Almería,001,0,Abla
04,Almería,002,5,Abrucena
04,Almería,003,1,Adra
04,Almería,004,6,Albánchez
04,Almería,005,9,Alboloduy
04,Almería,006,2,Albox
04,Almería,007,8,Alcolea
04,Almería,008,4,Alcóntar
04,Almería,009,7,Alcudia de Monteagud
04,Almería,010,1,Alhabia
04,Almería,011,8,Alhama de Almería
04,Almería,012,3,Alicún
04,Almería,013,9,Almería
04,Almería,014,4,Almócita
04,Almería,015,7,Alsodux
04,Almería,016,0,Antas
04,Almería,017,6,Arboleas
04,Almería,018,2,Armuña de Almanzora
04,Almería,019,5,Bacares
04,Almería,020,9,Bayárcal
04,Almería,021,6,Bayarque
04,Almería,022,1,Bédar
04,Almería,023,7,Beires
04,Almería,024,2,Benahadux
04,Almería,026,8,Benitagla
04,Almería,027,4,Benizalón
04,Almería,028,0,Bentarique
04,Almería,029,3,Berja
04,Almería,030,7,Canjáyar
04,Almería,031,4,Cantoria
04,Almería,032,9,Carboneras
04,Almería,033,5,Castro de Filabres
04,Almería,036,6,Chercos
04,Almería,037,2,Chirivel
04,Almería,034,0,Cóbdar
04,Almería,035,3,Cuevas del Almanzora
04,Almería,038,8,Dalías
04,Almería,902,6,"Ejido, El"
04,Almería,041,2,Enix
04,Almería,043,3,Felix
04,Almería,044,8,Fines
04,Almería,045,1,Fiñana
04,Almería,046,4,Fondón
04,Almería,047,0,Gádor
04,Almería,048,6,"Gallardos, Los"
04,Almería,049,9,Garrucha
04,Almería,050,2,Gérgal
04,Almería,051,9,Huécija
04,Almería,052,4,Huércal de Almería
04,Almería,053,0,Huércal-Overa
04,Almería,054,5,Illar
04,Almería,055,8,Instinción
04,Almería,056,1,Laroya
04,Almería,057,7,Láujar de Andarax
04,Almería,058,3,Líjar
04,Almería,059,6,Lubrín
04,Almería,060,0,Lucainena de las Torres
04,Almería,061,7,Lúcar
04,Almería,062,2,Macael
04,Almería,063,8,María
04,Almería,064,3,Mojácar
04,Almería,903,2,"Mojonera, La"
04,Almería,065,6,Nacimiento
04,Almería,066,9,Níjar
04,Almería,067,5,Ohanes
04,Almería,068,1,Olula de Castro
04,Almería,069,4,Olula del Río
04,Almería,070,8,Oria
04,Almería,071,5,Padules
04,Almería,072,0,Partaloa
04,Almería,073,6,Paterna del Río
04,Almería,074,1,Pechina
04,Almería,075,4,Pulpí
04,Almería,076,7,Purchena
04,Almería,077,3,Rágol
04,Almería,078,9,Rioja
04,Almería,079,2,Roquetas de Mar
04,Almería,080,6,Santa Cruz de Marchena
04,Almería,081,3,Santa Fe de Mondújar
04,Almería,082,8,Senés
04,Almería,083,4,Serón
04,Almería,084,9,Sierro
04,Almería,085,2,Somontín
04,Almería,086,5,Sorbas
04,Almería,087,1,Suflí
04,Almería,088,7,Tabernas
04,Almería,089,0,Taberno
04,Almería,090,4,Tahal
04,Almería,091,1,Terque
04,Almería,092,6,Tíjola
04,Almería,901,1,"Tres Villas, Las"
04,Almería,093,2,Turre
04,Almería,094,7,Turrillas
04,Almería,095,0,Uleila del Campo
04,Almería,096,3,Urrácal
04,Almería,097,9,Velefique
04,Almería,098,5,Vélez-Blanco
04,Almería,099,8,Vélez-Rubio
04,Almería,100,2,Vera
04,Almería,101,9,Viator
04,Almería,102,4,Vícar
04,Almería,103,0,Zurgena
05,Ávila,000,0,Avila
05,Ávila,001,3,Adanero
05,Ávila,002,8,"Adrada, La"
05,Ávila,005,2,Albornos
05,Ávila,007,1,Aldeanueva de Santa Cruz
05,Ávila,008,7,Aldeaseca
05,Ávila,010,4,"Aldehuela, La"
05,Ávila,012,6,Amavida
05,Ávila,013,2,"Arenal, El"
05,Ávila,014,7,Arenas de San Pedro
05,Ávila,015,0,Arevalillo
05,Ávila,016,3,Arévalo
05,Ávila,017,9,Aveinte
05,Ávila,018,5,Avellaneda
05,Ávila,019,8,Ávila
05,Ávila,021,9,"Barco de Ávila, El"
05,Ávila,022,4,"Barraco, El"
05,Ávila,023,0,Barromán
05,Ávila,024,5,Becedas
05,Ávila,025,8,Becedillas
05,Ávila,026,1,Bercial de Zapardiel
05,Ávila,027,7,"Berlanas, Las"
05,Ávila,029,6,Bernuy-Zapardiel
05,Ávila,030,0,Berrocalejo de Aragona
05,Ávila,033,8,Blascomillán
05,Ávila,034,3,Blasconuño de Matacabras
05,Ávila,035,6,Blascosancho
05,Ávila,036,9,"Bohodón, El"
05,Ávila,037,5,Bohoyo
05,Ávila,038,1,Bonilla de la Sierra
05,Ávila,039,4,Brabos
05,Ávila,040,8,Bularros
05,Ávila,041,5,Burgohondo
05,Ávila,042,0,Cabezas de Alambre
05,Ávila,043,6,Cabezas del Pozo
05,Ávila,044,1,Cabezas del Villar
05,Ávila,045,4,Cabizuela
05,Ávila,046,7,Canales
05,Ávila,047,3,Candeleda
05,Ávila,048,9,Cantiveros
05,Ávila,049,2,Cardeñosa
05,Ávila,051,2,"Carrera, La"
05,Ávila,052,7,Casas del Puerto
05,Ávila,053,3,Casasola
05,Ávila,054,8,Casavieja
05,Ávila,055,1,Casillas
05,Ávila,056,4,Castellanos de Zapardiel
05,Ávila,057,0,Cebreros
05,Ávila,058,6,Cepeda la Mora
05,Ávila,067,8,Chamartín
05,Ávila,059,9,Cillán
05,Ávila,060,3,Cisla
05,Ávila,061,0,"Colilla, La"
05,Ávila,062,5,Collado de Contreras
05,Ávila,063,1,Collado del Mirón
05,Ávila,064,6,Constanzana
05,Ávila,065,9,Crespos
05,Ávila,066,2,Cuevas del Valle
05,Ávila,903,5,Diego del Carpio
05,Ávila,069,7,Donjimeno
05,Ávila,070,1,Donvidas
05,Ávila,072,3,Espinosa de los Caballeros
05,Ávila,073,9,Flores de Ávila
05,Ávila,074,4,Fontiveros
05,Ávila,075,7,Fresnedilla
05,Ávila,076,0,"Fresno, El"
05,Ávila,077,6,Fuente el Saúz
05,Ávila,078,2,Fuentes de Año
05,Ávila,079,5,Gallegos de Altamiros
05,Ávila,080,9,Gallegos de Sobrinos
05,Ávila,081,6,Garganta del Villar
05,Ávila,082,1,Gavilanes
05,Ávila,083,7,Gemuño
05,Ávila,085,5,Gil García
05,Ávila,084,2,Gilbuena
05,Ávila,086,8,Gimialcón
05,Ávila,087,4,Gotarrendura
05,Ávila,088,0,Grandes y San Martín
05,Ávila,089,3,Guisando
05,Ávila,090,7,Gutierre-Muñoz
05,Ávila,092,9,Hernansancho
05,Ávila,093,5,Herradón de Pinares
05,Ávila,094,0,Herreros de Suso
05,Ávila,095,3,Higuera de las Dueñas
05,Ávila,096,6,"Hija de Dios, La"
05,Ávila,097,2,"Horcajada, La"
05,Ávila,099,1,Horcajo de las Torres
05,Ávila,100,5,"Hornillo, El"
05,Ávila,102,7,"Hoyo de Pinares, El"
05,Ávila,101,2,Hoyocasero
05,Ávila,103,3,Hoyorredondo
05,Ávila,106,4,Hoyos de Miguel Muñoz
05,Ávila,104,8,Hoyos del Collado
05,Ávila,105,1,Hoyos del Espino
05,Ávila,107,0,Hurtumpascual
05,Ávila,108,6,Junciana
05,Ávila,109,9,Langa
05,Ávila,110,3,Lanzahíta
05,Ávila,113,1,"Llanos de Tormes, Los"
05,Ávila,112,5,"Losar del Barco, El"
05,Ávila,114,6,Madrigal de las Altas Torres
05,Ávila,115,9,Maello
05,Ávila,116,2,Malpartida de Corneja
05,Ávila,117,8,Mamblas
05,Ávila,118,4,Mancera de Arriba
05,Ávila,119,7,Manjabálago
05,Ávila,120,1,Marlín
05,Ávila,121,8,Martiherrero
05,Ávila,122,3,Martínez
05,Ávila,123,9,Mediana de Voltoya
05,Ávila,124,4,Medinilla
05,Ávila,125,7,Mengamuñoz
05,Ávila,126,0,Mesegar de Corneja
05,Ávila,127,6,Mijares
05,Ávila,128,2,Mingorría
05,Ávila,129,5,"Mirón, El"
05,Ávila,130,9,Mironcillo
05,Ávila,131,6,Mirueña de los Infanzones
05,Ávila,132,1,Mombeltrán
05,Ávila,133,7,Monsalupe
05,Ávila,134,2,Moraleja de Matacabras
05,Ávila,135,5,Muñana
05,Ávila,136,8,Muñico
05,Ávila,138,0,Muñogalindo
05,Ávila,139,3,Muñogrande
05,Ávila,140,7,Muñomer del Peco
05,Ávila,141,4,Muñopepe
05,Ávila,142,9,Muñosancho
05,Ávila,143,5,Muñotello
05,Ávila,144,0,Narrillos del Álamo
05,Ávila,145,3,Narrillos del Rebollar
05,Ávila,149,1,Narros de Saldueña
05,Ávila,147,2,Narros del Castillo
05,Ávila,148,8,Narros del Puerto
05,Ávila,152,6,Nava de Arévalo
05,Ávila,153,2,Nava del Barco
05,Ávila,151,1,Navacepedilla de Corneja
05,Ávila,154,7,Navadijos
05,Ávila,155,0,Navaescurial
05,Ávila,156,3,Navahondilla
05,Ávila,157,9,Navalacruz
05,Ávila,158,5,Navalmoral
05,Ávila,159,8,Navalonguilla
05,Ávila,160,2,Navalosa
05,Ávila,161,9,Navalperal de Pinares
05,Ávila,162,4,Navalperal de Tormes
05,Ávila,163,0,Navaluenga
05,Ávila,164,5,Navaquesera
05,Ávila,165,8,Navarredonda de Gredos
05,Ávila,166,1,Navarredondilla
05,Ávila,167,7,Navarrevisca
05,Ávila,168,3,"Navas del Marqués, Las"
05,Ávila,169,6,Navatalgordo
05,Ávila,170,0,Navatejares
05,Ávila,171,7,Neila de San Miguel
05,Ávila,172,2,Niharra
05,Ávila,173,8,Ojos-Albos
05,Ávila,174,3,Orbita
05,Ávila,175,6,"Oso, El"
05,Ávila,176,9,Padiernos
05,Ávila,177,5,Pajares de Adaja
05,Ávila,178,1,Palacios de Goda
05,Ávila,179,4,Papatrigo
05,Ávila,180,8,"Parral, El"
05,Ávila,181,5,Pascualcobo
05,Ávila,182,0,Pedro Bernardo
05,Ávila,183,6,Pedro-Rodríguez
05,Ávila,184,1,Peguerinos
05,Ávila,185,4,Peñalba de Ávila
05,Ávila,186,7,Piedrahíta
05,Ávila,187,3,Piedralaves
05,Ávila,188,9,Poveda
05,Ávila,189,2,Poyales del Hoyo
05,Ávila,190,6,Pozanco
05,Ávila,191,3,Pradosegar
05,Ávila,192,8,Puerto Castilla
05,Ávila,193,4,Rasueros
05,Ávila,194,9,Riocabado
05,Ávila,195,2,Riofrío
05,Ávila,196,5,Rivilla de Barajas
05,Ávila,197,1,Salobral
05,Ávila,198,7,Salvadiós
05,Ávila,199,0,San Bartolomé de Béjar
05,Ávila,200,4,San Bartolomé de Corneja
05,Ávila,201,1,San Bartolomé de Pinares
05,Ávila,206,3,San Esteban de los Patos
05,Ávila,208,5,San Esteban de Zapardiel
05,Ávila,207,9,San Esteban del Valle
05,Ávila,209,8,San García de Ingelmos
05,Ávila,901,4,San Juan de Gredos
05,Ávila,210,2,San Juan de la Encinilla
05,Ávila,211,9,San Juan de la Nava
05,Ávila,212,4,San Juan del Molinillo
05,Ávila,213,0,San Juan del Olmo
05,Ávila,214,5,San Lorenzo de Tormes
05,Ávila,215,8,San Martín de la Vega del Alberche
05,Ávila,216,1,San Martín del Pimpollar
05,Ávila,217,7,San Miguel de Corneja
05,Ávila,218,3,San Miguel de Serrezuela
05,Ávila,219,6,San Pascual
05,Ávila,220,0,San Pedro del Arroyo
05,Ávila,231,5,San Vicente de Arévalo
05,Ávila,204,7,Sanchidrián
05,Ávila,205,0,Sanchorreja
05,Ávila,222,2,Santa Cruz de Pinares
05,Ávila,221,7,Santa Cruz del Valle
05,Ávila,226,9,Santa María de los Caballeros
05,Ávila,224,3,Santa María del Arroyo
05,Ávila,225,6,Santa María del Berrocal
05,Ávila,902,9,Santa María del Cubillo
05,Ávila,227,5,Santa María del Tiétar
05,Ávila,228,1,Santiago del Collado
05,Ávila,904,0,Santiago del Tormes
05,Ávila,229,4,Santo Domingo de las Posadas
05,Ávila,230,8,Santo Tomé de Zabarcos
05,Ávila,232,0,"Serrada, La"
05,Ávila,233,6,Serranillos
05,Ávila,234,1,Sigeres
05,Ávila,235,4,Sinlabajos
05,Ávila,236,7,Solana de Ávila
05,Ávila,237,3,Solana de Rioalmar
05,Ávila,238,9,Solosancho
05,Ávila,239,2,Sotalbo
05,Ávila,240,6,Sotillo de la Adrada
05,Ávila,241,3,"Tiemblo, El"
05,Ávila,242,8,Tiñosillos
05,Ávila,243,4,Tolbaños
05,Ávila,244,9,Tormellas
05,Ávila,245,2,Tornadizos de Ávila
05,Ávila,247,1,"Torre, La"
05,Ávila,246,5,Tórtoles
05,Ávila,249,0,Umbrías
05,Ávila,251,0,Vadillo de la Sierra
05,Ávila,252,5,Valdecasa
05,Ávila,253,1,Vega de Santa María
05,Ávila,254,6,Velayos
05,Ávila,256,2,Villaflor
05,Ávila,257,8,Villafranca de la Sierra
05,Ávila,905,3,Villanueva de Ávila
05,Ávila,258,4,Villanueva de Gómez
05,Ávila,259,7,Villanueva del Aceral
05,Ávila,260,1,Villanueva del Campillo
05,Ávila,261,8,Villar de Corneja
05,Ávila,262,3,Villarejo del Valle
05,Ávila,263,9,Villatoro
05,Ávila,264,4,Viñegra de Moraña
05,Ávila,265,7,Vita
05,Ávila,266,0,Zapardiel de la Cañada
05,Ávila,267,6,Zapardiel de la Ribera
06,Badajoz,001,6,Acedera
06,Badajoz,002,1,Aceuchal
06,Badajoz,003,7,Ahillones
06,Badajoz,004,2,Alange
06,Badajoz,005,5,"Albuera, La"
06,Badajoz,006,8,Alburquerque
06,Badajoz,007,4,Alconchel
06,Badajoz,008,0,Alconera
06,Badajoz,009,3,Aljucén
06,Badajoz,010,7,Almendral
06,Badajoz,011,4,Almendralejo
06,Badajoz,012,9,Arroyo de San Serván
06,Badajoz,013,5,Atalaya
06,Badajoz,014,0,Azuaga
06,Badajoz,015,3,Badajoz
06,Badajoz,016,6,Barcarrota
06,Badajoz,017,2,Baterno
06,Badajoz,018,8,Benquerencia de la Serena
06,Badajoz,019,1,Berlanga
06,Badajoz,020,5,Bienvenida
06,Badajoz,021,2,Bodonal de la Sierra
06,Badajoz,022,7,Burguillos del Cerro
06,Badajoz,023,3,Cabeza del Buey
06,Badajoz,024,8,Cabeza la Vaca
06,Badajoz,025,1,Calamonte
06,Badajoz,026,4,Calera de León
06,Badajoz,027,0,Calzadilla de los Barros
06,Badajoz,028,6,Campanario
06,Badajoz,029,9,Campillo de Llerena
06,Badajoz,030,3,Capilla
06,Badajoz,031,0,Carmonita
06,Badajoz,032,5,"Carrascalejo, El"
06,Badajoz,033,1,Casas de Don Pedro
06,Badajoz,034,6,Casas de Reina
06,Badajoz,035,9,Castilblanco
06,Badajoz,036,2,Castuera
06,Badajoz,042,3,Cheles
06,Badajoz,037,8,"Codosera, La"
06,Badajoz,038,4,Cordobilla de Lácara
06,Badajoz,039,7,"Coronada, La"
06,Badajoz,040,1,Corte de Peleas
06,Badajoz,041,8,Cristina
06,Badajoz,043,9,Don Álvaro
06,Badajoz,044,4,Don Benito
06,Badajoz,045,7,Entrín Bajo
06,Badajoz,046,0,Esparragalejo
06,Badajoz,047,6,Esparragosa de la Serena
06,Badajoz,048,2,Esparragosa de Lares
06,Badajoz,049,5,Feria
06,Badajoz,050,8,Fregenal de la Sierra
06,Badajoz,051,5,Fuenlabrada de los Montes
06,Badajoz,052,0,Fuente de Cantos
06,Badajoz,053,6,Fuente del Arco
06,Badajoz,054,1,Fuente del Maestre
06,Badajoz,055,4,Fuentes de León
06,Badajoz,056,7,Garbayuela
06,Badajoz,057,3,Garlitos
06,Badajoz,058,9,"Garrovilla, La"
06,Badajoz,059,2,Granja de Torrehermosa
06,Badajoz,060,6,Guareña
06,Badajoz,061,3,"Haba, La"
06,Badajoz,062,8,Helechosa de los Montes
06,Badajoz,063,4,Herrera del Duque
06,Badajoz,064,9,Higuera de la Serena
06,Badajoz,065,2,Higuera de Llerena
06,Badajoz,066,5,Higuera de Vargas
06,Badajoz,067,1,Higuera la Real
06,Badajoz,068,7,Hinojosa del Valle
06,Badajoz,069,0,Hornachos
06,Badajoz,070,4,Jerez de los Caballeros
06,Badajoz,071,1,"Lapa, La"
06,Badajoz,073,2,Llera
06,Badajoz,074,7,Llerena
06,Badajoz,072,6,Lobón
06,Badajoz,075,0,Magacela
06,Badajoz,076,3,Maguilla
06,Badajoz,077,9,Malcocinado
06,Badajoz,078,5,Malpartida de la Serena
06,Badajoz,079,8,Manchita
06,Badajoz,080,2,Medellín
06,Badajoz,081,9,Medina de las Torres
06,Badajoz,082,4,Mengabril
06,Badajoz,083,0,Mérida
06,Badajoz,084,5,Mirandilla
06,Badajoz,085,8,Monesterio
06,Badajoz,086,1,Montemolín
06,Badajoz,087,7,Monterrubio de la Serena
06,Badajoz,088,3,Montijo
06,Badajoz,089,6,"Morera, La"
06,Badajoz,090,0,"Nava de Santiago, La"
06,Badajoz,091,7,Navalvillar de Pela
06,Badajoz,092,2,Nogales
06,Badajoz,093,8,Oliva de la Frontera
06,Badajoz,094,3,Oliva de Mérida
06,Badajoz,095,6,Olivenza
06,Badajoz,096,9,Orellana de la Sierra
06,Badajoz,097,5,Orellana la Vieja
06,Badajoz,098,1,Palomas
06,Badajoz,099,4,"Parra, La"
06,Badajoz,100,8,Peñalsordo
06,Badajoz,101,5,Peraleda del Zaucejo
06,Badajoz,102,0,Puebla de Alcocer
06,Badajoz,103,6,Puebla de la Calzada
06,Badajoz,104,1,Puebla de la Reina
06,Badajoz,107,3,Puebla de Obando
06,Badajoz,108,9,Puebla de Sancho Pérez
06,Badajoz,105,4,Puebla del Maestre
06,Badajoz,106,7,Puebla del Prior
06,Badajoz,902,2,Pueblonuevo del Guadiana
06,Badajoz,109,2,Quintana de la Serena
06,Badajoz,110,6,Reina
06,Badajoz,111,3,Rena
06,Badajoz,112,8,Retamal de Llerena
06,Badajoz,113,4,Ribera del Fresno
06,Badajoz,114,9,Risco
06,Badajoz,115,2,"Roca de la Sierra, La"
06,Badajoz,116,5,Salvaleón
06,Badajoz,117,1,Salvatierra de los Barros
06,Badajoz,119,0,San Pedro de Mérida
06,Badajoz,123,2,San Vicente de Alcántara
06,Badajoz,118,7,Sancti-Spíritus
06,Badajoz,120,4,Santa Amalia
06,Badajoz,121,1,Santa Marta
06,Badajoz,122,6,"Santos de Maimona, Los"
06,Badajoz,124,7,Segura de León
06,Badajoz,125,0,Siruela
06,Badajoz,126,3,Solana de los Barros
06,Badajoz,127,9,Talarrubias
06,Badajoz,128,5,Talavera la Real
06,Badajoz,129,8,Táliga
06,Badajoz,130,2,Tamurejo
06,Badajoz,131,9,Torre de Miguel Sesmero
06,Badajoz,132,4,Torremayor
06,Badajoz,133,0,Torremejía
06,Badajoz,134,5,Trasierra
06,Badajoz,135,8,Trujillanos
06,Badajoz,136,1,Usagre
06,Badajoz,137,7,Valdecaballeros
06,Badajoz,901,7,Valdelacalzada
06,Badajoz,138,3,Valdetorres
06,Badajoz,139,6,Valencia de las Torres
06,Badajoz,140,0,Valencia del Mombuey
06,Badajoz,141,7,Valencia del Ventoso
06,Badajoz,146,9,Valle de la Serena
06,Badajoz,147,5,Valle de Matamoros
06,Badajoz,148,1,Valle de Santa Ana
06,Badajoz,142,2,Valverde de Burguillos
06,Badajoz,143,8,Valverde de Leganés
06,Badajoz,144,3,Valverde de Llerena
06,Badajoz,145,6,Valverde de Mérida
06,Badajoz,149,4,Villafranca de los Barros
06,Badajoz,150,7,Villagarcía de la Torre
06,Badajoz,151,4,Villagonzalo
06,Badajoz,152,9,Villalba de los Barros
06,Badajoz,153,5,Villanueva de la Serena
06,Badajoz,154,0,Villanueva del Fresno
06,Badajoz,156,6,Villar de Rena
06,Badajoz,155,3,Villar del Rey
06,Badajoz,157,2,Villarta de los Montes
06,Badajoz,158,8,Zafra
06,Badajoz,159,1,Zahínos
06,Badajoz,160,5,Zalamea de la Serena
06,Badajoz,162,7,"Zarza, La"
06,Badajoz,161,2,Zarza-Capilla
07,Islas Baleares,002,7,Palma de Mallorca
07,Islas Baleares,002,7,Alaior
07,Islas Baleares,001,2,Alaró
07,Islas Baleares,003,3,Alcúdia
07,Islas Baleares,004,8,Algaida
07,Islas Baleares,005,1,Andratx
07,Islas Baleares,901,3,Ariany
07,Islas Baleares,006,4,Artà
07,Islas Baleares,007,0,Banyalbufar
07,Islas Baleares,008,6,Binissalem
07,Islas Baleares,009,9,Búger
07,Islas Baleares,010,3,Bunyola
07,Islas Baleares,011,0,Calvià
07,Islas Baleares,012,5,Campanet
07,Islas Baleares,013,1,Campos
07,Islas Baleares,014,6,Capdepera
07,Islas Baleares,064,5,"Castell, Es"
07,Islas Baleares,015,9,Ciutadella de Menorca
07,Islas Baleares,016,2,Consell
07,Islas Baleares,017,8,Costitx
07,Islas Baleares,018,4,Deyá
07,Islas Baleares,026,0,Eivissa
07,Islas Baleares,019,7,Escorca
07,Islas Baleares,020,1,Esporles
07,Islas Baleares,021,8,Estellencs
07,Islas Baleares,022,3,Felanitx
07,Islas Baleares,023,9,Ferreries
07,Islas Baleares,024,4,Formentera
07,Islas Baleares,025,7,Fornalutx
07,Islas Baleares,027,6,Inca
07,Islas Baleares,028,2,Lloret de Vistalegre
07,Islas Baleares,029,5,Lloseta
07,Islas Baleares,030,9,Llubí
07,Islas Baleares,031,6,Llucmajor
07,Islas Baleares,033,7,Manacor
07,Islas Baleares,034,2,Mancor de la Vall
07,Islas Baleares,032,1,Maó
07,Islas Baleares,035,5,Maria de la Salut
07,Islas Baleares,036,8,Marratxí
07,Islas Baleares,037,4,"Mercadal, Es"
07,Islas Baleares,902,8,"Migjorn Gran, Es"
07,Islas Baleares,038,0,Montuïri
07,Islas Baleares,039,3,Muro
07,Islas Baleares,040,7,Palma
07,Islas Baleares,041,4,Petra
07,Islas Baleares,044,0,"Pobla, Sa"
07,Islas Baleares,042,9,Pollença
07,Islas Baleares,043,5,Porreres
07,Islas Baleares,045,3,Puigpunyent
07,Islas Baleares,059,8,"Salines, Ses"
07,Islas Baleares,046,6,Sant Antoni de Portmany
07,Islas Baleares,049,1,Sant Joan
07,Islas Baleares,050,4,Sant Joan de Labritja
07,Islas Baleares,048,8,Sant Josep de sa Talaia
07,Islas Baleares,051,1,Sant Llorenç des Cardassar
07,Islas Baleares,052,6,Sant Lluís
07,Islas Baleares,053,2,Santa Eugènia
07,Islas Baleares,054,7,Santa Eulalia del Río
07,Islas Baleares,055,0,Santa Margalida
07,Islas Baleares,056,3,Santa María del Camí
07,Islas Baleares,057,9,Santanyí
07,Islas Baleares,058,5,Selva
07,Islas Baleares,047,2,Sencelles
07,Islas Baleares,060,2,Sineu
07,Islas Baleares,061,9,Sóller
07,Islas Baleares,062,4,Son Servera
07,Islas Baleares,063,0,Valldemossa
07,Islas Baleares,065,8,Vilafranca de Bonany
08,Barcelona,000,0,Castellar del Valles
08,Barcelona,000,0,Hospitalet de Llobregat
08,Barcelona,001,8,Mataro
08,Barcelona,001,8,Abrera
08,Barcelona,002,3,Aguilar de Segarra
08,Barcelona,014,2,Aiguafreda
08,Barcelona,003,9,Alella
08,Barcelona,004,4,Alpens
08,Barcelona,005,7,"Ametlla del Vallès, L'"
08,Barcelona,006,0,Arenys de Mar
08,Barcelona,007,6,Arenys de Munt
08,Barcelona,008,2,Argençola
08,Barcelona,009,5,Argentona
08,Barcelona,010,9,Artés
08,Barcelona,011,6,Avià
08,Barcelona,012,1,Avinyó
08,Barcelona,013,7,Avinyonet del Penedès
08,Barcelona,015,5,Badalona
08,Barcelona,904,5,Badia del Vallès
08,Barcelona,016,8,Bagà
08,Barcelona,017,4,Balenyà
08,Barcelona,018,0,Balsareny
08,Barcelona,252,0,Barberà del Vallès
08,Barcelona,019,3,Barcelona
08,Barcelona,020,7,Begues
08,Barcelona,021,4,Bellprat
08,Barcelona,022,9,Berga
08,Barcelona,023,5,Bigues i Riells
08,Barcelona,024,0,Borredà
08,Barcelona,025,3,"Bruc, El"
08,Barcelona,026,6,"Brull, El"
08,Barcelona,027,2,"Cabanyes, Les"
08,Barcelona,028,8,Cabrera d'Anoia
08,Barcelona,029,1,Cabrera de Mar
08,Barcelona,030,5,Cabrils
08,Barcelona,031,2,Calaf
08,Barcelona,034,8,Calders
08,Barcelona,033,3,Caldes de Montbui
08,Barcelona,032,7,Caldes d'Estrac
08,Barcelona,035,1,Calella
08,Barcelona,037,0,Calldetenes
08,Barcelona,038,6,Callús
08,Barcelona,036,4,Calonge de Segarra
08,Barcelona,039,9,Campins
08,Barcelona,040,3,Canet de Mar
08,Barcelona,041,0,Canovelles
08,Barcelona,042,5,Cànoves i Samalús
08,Barcelona,043,1,Canyelles
08,Barcelona,044,6,Capellades
08,Barcelona,045,9,Capolat
08,Barcelona,046,2,Cardedeu
08,Barcelona,047,8,Cardona
08,Barcelona,048,4,Carme
08,Barcelona,049,7,Casserres
08,Barcelona,057,5,Castell de l'Areny
08,Barcelona,052,2,Castellar de n'Hug
08,Barcelona,050,0,Castellar del Riu
08,Barcelona,051,7,Castellar del Vallès
08,Barcelona,053,8,Castellbell i el Vilar
08,Barcelona,054,3,Castellbisbal
08,Barcelona,055,6,Castellcir
08,Barcelona,056,9,Castelldefels
08,Barcelona,058,1,Castellet i la Gornal
08,Barcelona,060,8,Castellfollit de Riubregós
08,Barcelona,059,4,Castellfollit del Boix
08,Barcelona,061,5,Castellgalí
08,Barcelona,062,0,Castellnou de Bages
08,Barcelona,063,6,Castellolí
08,Barcelona,064,1,Castellterçol
08,Barcelona,065,4,Castellví de la Marca
08,Barcelona,066,7,Castellví de Rosanes
08,Barcelona,067,3,Centelles
08,Barcelona,268,7,Cercs
08,Barcelona,266,5,Cerdanyola del Vallès
08,Barcelona,068,9,Cervelló
08,Barcelona,069,2,Collbató
08,Barcelona,070,6,Collsuspina
08,Barcelona,071,3,Copons
08,Barcelona,072,8,Corbera de Llobregat
08,Barcelona,073,4,Cornellà de Llobregat
08,Barcelona,074,9,Cubelles
08,Barcelona,075,2,Dosrius
08,Barcelona,076,5,Esparreguera
08,Barcelona,077,1,Esplugues de Llobregat
08,Barcelona,078,7,"Espunyola, L'"
08,Barcelona,079,0,"Estany, L'"
08,Barcelona,134,7,Figaró-Montmany
08,Barcelona,080,4,Fígols
08,Barcelona,082,6,Fogars de la Selva
08,Barcelona,081,1,Fogars de Montclús
08,Barcelona,083,2,Folgueroles
08,Barcelona,084,7,Fonollosa
08,Barcelona,085,0,Font-rubí
08,Barcelona,086,3,"Franqueses del Vallès, Les"
08,Barcelona,090,2,Gaià
08,Barcelona,087,9,Gallifa
08,Barcelona,088,5,"Garriga, La"
08,Barcelona,089,8,Gavà
08,Barcelona,091,9,Gelida
08,Barcelona,092,4,Gironella
08,Barcelona,093,0,Gisclareny
08,Barcelona,094,5,"Granada, La"
08,Barcelona,095,8,Granera
08,Barcelona,096,1,Granollers
08,Barcelona,097,7,Gualba
08,Barcelona,099,6,Guardiola de Berguedà
08,Barcelona,100,0,Gurb
08,Barcelona,101,7,"Hospitalet de Llobregat, L'"
08,Barcelona,162,9,"Hostalets de Pierola, Els"
08,Barcelona,102,2,Igualada
08,Barcelona,103,8,Jorba
08,Barcelona,104,3,"Llacuna, La"
08,Barcelona,105,6,"Llagosta, La"
08,Barcelona,107,5,Lliçà d'Amunt
08,Barcelona,108,1,Lliçà de Vall
08,Barcelona,106,9,Llinars del Vallès
08,Barcelona,109,4,Lluçà
08,Barcelona,110,8,Malgrat de Mar
08,Barcelona,111,5,Malla
08,Barcelona,112,0,Manlleu
08,Barcelona,113,6,Manresa
08,Barcelona,242,3,Marganell
08,Barcelona,114,1,Martorell
08,Barcelona,115,4,Martorelles
08,Barcelona,116,7,"Masies de Roda, Les"
08,Barcelona,117,3,"Masies de Voltregà, Les"
08,Barcelona,118,9,"Masnou, El"
08,Barcelona,119,2,Masquefa
08,Barcelona,120,6,Matadepera
08,Barcelona,121,3,Mataró
08,Barcelona,122,8,Mediona
08,Barcelona,138,5,Moià
08,Barcelona,123,4,Molins de Rei
08,Barcelona,124,9,Mollet del Vallès
08,Barcelona,128,7,Monistrol de Calders
08,Barcelona,127,1,Monistrol de Montserrat
08,Barcelona,125,2,Montcada i Reixac
08,Barcelona,130,4,Montclar
08,Barcelona,131,1,Montesquiu
08,Barcelona,126,5,Montgat
08,Barcelona,132,6,Montmajor
08,Barcelona,133,2,Montmaneu
08,Barcelona,135,0,Montmeló
08,Barcelona,136,3,Montornès del Vallès
08,Barcelona,137,9,Montseny
08,Barcelona,129,0,Muntanyola
08,Barcelona,139,8,Mura
08,Barcelona,140,2,Navarcles
08,Barcelona,141,9,Navàs
08,Barcelona,142,4,"Nou de Berguedà, La"
08,Barcelona,143,0,Òdena
08,Barcelona,145,8,Olèrdola
08,Barcelona,146,1,Olesa de Bonesvalls
08,Barcelona,147,7,Olesa de Montserrat
08,Barcelona,148,3,Olivella
08,Barcelona,149,6,Olost
08,Barcelona,144,5,Olvan
08,Barcelona,150,9,Orís
08,Barcelona,151,6,Oristà
08,Barcelona,152,1,Orpí
08,Barcelona,153,7,Òrrius
08,Barcelona,154,2,Pacs del Penedès
08,Barcelona,155,5,Palafolls
08,Barcelona,156,8,Palau-solità i Plegamans
08,Barcelona,157,4,Pallejà
08,Barcelona,905,8,"Palma de Cervelló, La"
08,Barcelona,158,0,"Papiol, El"
08,Barcelona,159,3,Parets del Vallès
08,Barcelona,160,7,Perafita
08,Barcelona,161,4,Piera
08,Barcelona,163,5,Pineda de Mar
08,Barcelona,164,0,"Pla del Penedès, El"
08,Barcelona,165,3,"Pobla de Claramunt, La"
08,Barcelona,166,6,"Pobla de Lillet, La"
08,Barcelona,167,2,Polinyà
08,Barcelona,182,5,"Pont de Vilomara i Rocafort, El"
08,Barcelona,168,8,Pontons
08,Barcelona,169,1,"Prat de Llobregat, El"
08,Barcelona,171,2,Prats de Lluçanès
08,Barcelona,170,5,"Prats de Rei, Els"
08,Barcelona,230,3,Premià de Dalt
08,Barcelona,172,7,Premià de Mar
08,Barcelona,174,8,Puigdàlber
08,Barcelona,175,1,Puig-reig
08,Barcelona,176,4,Pujalt
08,Barcelona,177,0,"Quar, La"
08,Barcelona,178,6,Rajadell
08,Barcelona,179,9,Rellinars
08,Barcelona,180,3,Ripollet
08,Barcelona,181,0,"Roca del Vallès, La"
08,Barcelona,183,1,Roda de Ter
08,Barcelona,184,6,Rubí
08,Barcelona,185,9,Rubió
08,Barcelona,901,9,Rupit i Pruit
08,Barcelona,187,8,Sabadell
08,Barcelona,188,4,Sagàs
08,Barcelona,190,1,Saldes
08,Barcelona,191,8,Sallent
08,Barcelona,194,4,Sant Adrià de Besòs
08,Barcelona,195,7,Sant Agustí de Lluçanès
08,Barcelona,196,0,Sant Andreu de la Barca
08,Barcelona,197,6,Sant Andreu de Llavaneres
08,Barcelona,198,2,Sant Antoni de Vilamajor
08,Barcelona,199,5,Sant Bartomeu del Grau
08,Barcelona,200,9,Sant Boi de Llobregat
08,Barcelona,201,6,Sant Boi de Lluçanès
08,Barcelona,203,7,Sant Cebrià de Vallalta
08,Barcelona,202,1,Sant Celoni
08,Barcelona,204,2,Sant Climent de Llobregat
08,Barcelona,205,5,Sant Cugat del Vallès
08,Barcelona,206,8,Sant Cugat Sesgarrigues
08,Barcelona,207,4,Sant Esteve de Palautordera
08,Barcelona,208,0,Sant Esteve Sesrovires
08,Barcelona,210,7,Sant Feliu de Codines
08,Barcelona,211,4,Sant Feliu de Llobregat
08,Barcelona,212,9,Sant Feliu Sasserra
08,Barcelona,209,3,Sant Fost de Campsentelles
08,Barcelona,213,5,Sant Fruitós de Bages
08,Barcelona,215,3,Sant Hipòlit de Voltregà
08,Barcelona,193,9,Sant Iscle de Vallalta
08,Barcelona,216,6,Sant Jaume de Frontanyà
08,Barcelona,218,8,Sant Joan de Vilatorrada
08,Barcelona,217,2,Sant Joan Despí
08,Barcelona,903,0,Sant Julià de Cerdanyola
08,Barcelona,220,5,Sant Julià de Vilatorta
08,Barcelona,221,2,Sant Just Desvern
08,Barcelona,222,7,Sant Llorenç d'Hortons
08,Barcelona,223,3,Sant Llorenç Savall
08,Barcelona,225,1,Sant Martí d'Albars
08,Barcelona,224,8,Sant Martí de Centelles
08,Barcelona,226,4,Sant Martí de Tous
08,Barcelona,227,0,Sant Martí Sarroca
08,Barcelona,228,6,Sant Martí Sesgueioles
08,Barcelona,229,9,Sant Mateu de Bages
08,Barcelona,231,0,Sant Pere de Ribes
08,Barcelona,232,5,Sant Pere de Riudebitlles
08,Barcelona,233,1,Sant Pere de Torelló
08,Barcelona,234,6,Sant Pere de Vilamajor
08,Barcelona,189,7,Sant Pere Sallavinera
08,Barcelona,235,9,Sant Pol de Mar
08,Barcelona,236,2,Sant Quintí de Mediona
08,Barcelona,237,8,Sant Quirze de Besora
08,Barcelona,238,4,Sant Quirze del Vallès
08,Barcelona,239,7,Sant Quirze Safaja
08,Barcelona,240,1,Sant Sadurní d'Anoia
08,Barcelona,241,8,Sant Sadurní d'Osormort
08,Barcelona,098,3,Sant Salvador de Guardiola
08,Barcelona,262,8,Sant Vicenç de Castellet
08,Barcelona,264,9,Sant Vicenç de Montalt
08,Barcelona,265,2,Sant Vicenç de Torelló
08,Barcelona,263,4,Sant Vicenç dels Horts
08,Barcelona,243,9,Santa Cecília de Voltregà
08,Barcelona,244,4,Santa Coloma de Cervelló
08,Barcelona,245,7,Santa Coloma de Gramenet
08,Barcelona,246,0,Santa Eugènia de Berga
08,Barcelona,247,6,Santa Eulàlia de Riuprimer
08,Barcelona,248,2,Santa Eulàlia de Ronçana
08,Barcelona,249,5,Santa Fe del Penedès
08,Barcelona,250,8,Santa Margarida de Montbui
08,Barcelona,251,5,Santa Margarida i els Monjos
08,Barcelona,253,6,Santa Maria de Besora
08,Barcelona,254,1,Santa Maria de Corcó
08,Barcelona,256,7,Santa Maria de Martorelles
08,Barcelona,255,4,Santa Maria de Merlès
08,Barcelona,257,3,Santa Maria de Miralles
08,Barcelona,259,2,Santa Maria de Palautordera
08,Barcelona,258,9,Santa Maria d'Oló
08,Barcelona,260,6,Santa Perpètua de Mogoda
08,Barcelona,261,3,Santa Susanna
08,Barcelona,192,3,Santpedor
08,Barcelona,267,1,Sentmenat
08,Barcelona,269,0,Seva
08,Barcelona,270,4,Sitges
08,Barcelona,271,1,Sobremunt
08,Barcelona,272,6,Sora
08,Barcelona,273,2,Subirats
08,Barcelona,274,7,Súria
08,Barcelona,276,3,Tagamanent
08,Barcelona,277,9,Talamanca
08,Barcelona,278,5,Taradell
08,Barcelona,275,0,Tavèrnoles
08,Barcelona,280,2,Tavertet
08,Barcelona,281,9,Teià
08,Barcelona,279,8,Terrassa
08,Barcelona,282,4,Tiana
08,Barcelona,283,0,Tona
08,Barcelona,284,5,Tordera
08,Barcelona,285,8,Torelló
08,Barcelona,286,1,"Torre de Claramunt, La"
08,Barcelona,287,7,Torrelavit
08,Barcelona,288,3,Torrelles de Foix
08,Barcelona,289,6,Torrelles de Llobregat
08,Barcelona,290,0,Ullastrell
08,Barcelona,291,7,Vacarisses
08,Barcelona,292,2,Vallbona d'Anoia
08,Barcelona,293,8,Vallcebre
08,Barcelona,294,3,Vallgorguina
08,Barcelona,295,6,Vallirana
08,Barcelona,296,9,Vallromanes
08,Barcelona,297,5,Veciana
08,Barcelona,298,1,Vic
08,Barcelona,299,4,Vilada
08,Barcelona,301,5,Viladecans
08,Barcelona,300,8,Viladecavalls
08,Barcelona,305,4,Vilafranca del Penedès
08,Barcelona,306,7,Vilalba Sasserra
08,Barcelona,303,6,Vilanova de Sau
08,Barcelona,302,0,Vilanova del Camí
08,Barcelona,902,4,Vilanova del Vallès
08,Barcelona,307,3,Vilanova i la Geltrú
08,Barcelona,214,0,Vilassar de Dalt
08,Barcelona,219,1,Vilassar de Mar
08,Barcelona,304,1,Vilobí del Penedès
08,Barcelona,308,9,Viver i Serrateix
09,Burgos,001,1,Abajas
09,Burgos,003,2,Adrada de Haza
09,Burgos,006,3,Aguas Cándidas
09,Burgos,007,9,Aguilar de Bureba
09,Burgos,009,8,Albillos
09,Burgos,010,2,Alcocero de Mola
09,Burgos,011,9,Alfoz de Bricia
09,Burgos,907,0,Alfoz de Quintanadueñas
09,Burgos,012,4,Alfoz de Santa Gadea
09,Burgos,013,0,Altable
09,Burgos,014,5,"Altos, Los"
09,Burgos,016,1,Ameyugo
09,Burgos,017,7,Anguix
09,Burgos,018,3,Aranda de Duero
09,Burgos,019,6,Arandilla
09,Burgos,020,0,Arauzo de Miel
09,Burgos,021,7,Arauzo de Salce
09,Burgos,022,2,Arauzo de Torre
09,Burgos,023,8,Arcos
09,Burgos,024,3,Arenillas de Riopisuerga
09,Burgos,025,6,Arija
09,Burgos,026,9,Arlanzón
09,Burgos,027,5,Arraya de Oca
09,Burgos,029,4,Atapuerca
09,Burgos,030,8,"Ausines, Los"
09,Burgos,032,0,Avellanosa de Muñó
09,Burgos,033,6,Bahabón de Esgueva
09,Burgos,034,1,"Balbases, Los"
09,Burgos,035,4,Baños de Valdearados
09,Burgos,036,7,Bañuelos de Bureba
09,Burgos,037,3,Barbadillo de Herreros
09,Burgos,038,9,Barbadillo del Mercado
09,Burgos,039,2,Barbadillo del Pez
09,Burgos,041,3,Barrio de Muñó
09,Burgos,043,4,"Barrios de Bureba, Los"
09,Burgos,044,9,Barrios de Colina
09,Burgos,045,2,Basconcillos del Tozo
09,Burgos,046,5,Bascuñana
09,Burgos,047,1,Belbimbre
09,Burgos,048,7,Belorado
09,Burgos,050,3,Berberana
09,Burgos,051,0,Berlangas de Roa
09,Burgos,052,5,Berzosa de Bureba
09,Burgos,054,6,Bozoó
09,Burgos,055,9,Brazacorta
09,Burgos,056,2,Briviesca
09,Burgos,057,8,Bugedo
09,Burgos,058,4,Buniel
09,Burgos,059,7,Burgos
09,Burgos,060,1,Busto de Bureba
09,Burgos,061,8,Cabañes de Esgueva
09,Burgos,062,3,Cabezón de la Sierra
09,Burgos,064,4,Caleruega
09,Burgos,065,7,Campillo de Aranda
09,Burgos,066,0,Campolara
09,Burgos,067,6,Canicosa de la Sierra
09,Burgos,068,2,Cantabrana
09,Burgos,070,9,Carazo
09,Burgos,071,6,Carcedo de Bureba
09,Burgos,072,1,Carcedo de Burgos
09,Burgos,073,7,Cardeñadijo
09,Burgos,074,2,Cardeñajimeno
09,Burgos,075,5,Cardeñuela Riopico
09,Burgos,076,8,Carrias
09,Burgos,077,4,Cascajares de Bureba
09,Burgos,078,0,Cascajares de la Sierra
09,Burgos,079,3,Castellanos de Castro
09,Burgos,083,5,Castil de Peones
09,Burgos,082,9,Castildelgado
09,Burgos,084,0,Castrillo de la Reina
09,Burgos,085,3,Castrillo de la Vega
09,Burgos,088,8,Castrillo de Riopisuerga
09,Burgos,086,6,Castrillo del Val
09,Burgos,090,5,Castrillo Matajudíos
09,Burgos,091,2,Castrojeriz
09,Burgos,063,9,Cavia
09,Burgos,093,3,Cayuela
09,Burgos,094,8,Cebrecos
09,Burgos,095,1,Celada del Camino
09,Burgos,098,6,Cerezo de Río Tirón
09,Burgos,100,3,Cerratón de Juarros
09,Burgos,101,0,Ciadoncha
09,Burgos,102,5,Cillaperlata
09,Burgos,103,1,Cilleruelo de Abajo
09,Burgos,104,6,Cilleruelo de Arriba
09,Burgos,105,9,Ciruelos de Cervera
09,Burgos,108,4,Cogollos
09,Burgos,109,7,Condado de Treviño
09,Burgos,110,1,Contreras
09,Burgos,112,3,Coruña del Conde
09,Burgos,113,9,Covarrubias
09,Burgos,114,4,Cubillo del Campo
09,Burgos,115,7,Cubo de Bureba
09,Burgos,117,6,"Cueva de Roa, La"
09,Burgos,119,5,Cuevas de San Clemente
09,Burgos,120,9,Encío
09,Burgos,122,1,Espinosa de Cervera
09,Burgos,124,2,Espinosa de los Monteros
09,Burgos,123,7,Espinosa del Camino
09,Burgos,125,5,Estépar
09,Burgos,127,4,Fontioso
09,Burgos,128,0,Frandovínez
09,Burgos,129,3,Fresneda de la Sierra Tirón
09,Burgos,130,7,Fresneña
09,Burgos,131,4,Fresnillo de las Dueñas
09,Burgos,132,9,Fresno de Río Tirón
09,Burgos,133,5,Fresno de Rodilla
09,Burgos,134,0,Frías
09,Burgos,135,3,Fuentebureba
09,Burgos,136,6,Fuentecén
09,Burgos,137,2,Fuentelcésped
09,Burgos,138,8,Fuentelisendo
09,Burgos,139,1,Fuentemolinos
09,Burgos,140,5,Fuentenebro
09,Burgos,141,2,Fuentespina
09,Burgos,143,3,Galbarros
09,Burgos,144,8,"Gallega, La"
09,Burgos,148,6,Grijalba
09,Burgos,149,9,Grisaleña
09,Burgos,151,9,Gumiel de Izán
09,Burgos,152,4,Gumiel de Mercado
09,Burgos,154,5,Hacinas
09,Burgos,155,8,Haza
09,Burgos,159,6,Hontanas
09,Burgos,160,0,Hontangas
09,Burgos,162,2,Hontoria de la Cantera
09,Burgos,164,3,Hontoria de Valdearados
09,Burgos,163,8,Hontoria del Pinar
09,Burgos,166,9,"Hormazas, Las"
09,Burgos,167,5,Hornillos del Camino
09,Burgos,168,1,"Horra, La"
09,Burgos,169,4,Hortigüela
09,Burgos,170,8,Hoyales de Roa
09,Burgos,172,0,Huérmeces
09,Burgos,173,6,Huerta de Arriba
09,Burgos,174,1,Huerta de Rey
09,Burgos,175,4,Humada
09,Burgos,176,7,Hurones
09,Burgos,177,3,Ibeas de Juarros
09,Burgos,178,9,Ibrillos
09,Burgos,179,2,Iglesiarrubia
09,Burgos,180,6,Iglesias
09,Burgos,181,3,Isar
09,Burgos,182,8,Itero del Castillo
09,Burgos,183,4,Jaramillo de la Fuente
09,Burgos,184,9,Jaramillo Quemado
09,Burgos,189,0,Junta de Traslaloma
09,Burgos,190,4,Junta de Villalba de Losa
09,Burgos,191,1,Jurisdicción de Lara
09,Burgos,192,6,Jurisdicción de San Zadornil
09,Burgos,194,7,Lerma
09,Burgos,195,0,Llano de Bureba
09,Burgos,196,3,Madrigal del Monte
09,Burgos,197,9,Madrigalejo del Monte
09,Burgos,198,5,Mahamud
09,Burgos,199,8,Mambrilla de Castrejón
09,Burgos,200,2,Mambrillas de Lara
09,Burgos,201,9,Mamolar
09,Burgos,202,4,Manciles
09,Burgos,206,1,Mazuela
09,Burgos,208,3,Mecerreyes
09,Burgos,209,6,Medina de Pomar
09,Burgos,211,7,Melgar de Fernamental
09,Burgos,213,8,Merindad de Cuesta-Urria
09,Burgos,214,3,Merindad de Montija
09,Burgos,906,4,Merindad de Río Ubierna
09,Burgos,215,6,Merindad de Sotoscueva
09,Burgos,216,9,Merindad de Valdeporres
09,Burgos,217,5,Merindad de Valdivielso
09,Burgos,218,1,Milagros
09,Burgos,219,4,Miranda de Ebro
09,Burgos,220,8,Miraveche
09,Burgos,221,5,Modúbar de la Emparedada
09,Burgos,223,6,Monasterio de la Sierra
09,Burgos,224,1,Monasterio de Rodilla
09,Burgos,225,4,Moncalvillo
09,Burgos,226,7,Monterrubio de la Demanda
09,Burgos,227,3,Montorio
09,Burgos,228,9,Moradillo de Roa
09,Burgos,229,2,Nava de Roa
09,Burgos,230,6,Navas de Bureba
09,Burgos,231,3,Nebreda
09,Burgos,232,8,Neila
09,Burgos,235,2,Olmedillo de Roa
09,Burgos,236,5,Olmillos de Muñó
09,Burgos,238,7,Oña
09,Burgos,239,0,Oquillas
09,Burgos,241,1,Orbaneja Riopico
09,Burgos,242,6,Padilla de Abajo
09,Burgos,243,2,Padilla de Arriba
09,Burgos,244,7,Padrones de Bureba
09,Burgos,246,3,Palacios de la Sierra
09,Burgos,247,9,Palacios de Riopisuerga
09,Burgos,248,5,Palazuelos de la Sierra
09,Burgos,249,8,Palazuelos de Muñó
09,Burgos,250,1,Pampliega
09,Burgos,251,8,Pancorbo
09,Burgos,253,9,Pardilla
09,Burgos,255,7,Partido de la Sierra en Tobalina
09,Burgos,256,0,Pedrosa de Duero
09,Burgos,259,5,Pedrosa de Río Úrbel
09,Burgos,257,6,Pedrosa del Páramo
09,Burgos,258,2,Pedrosa del Príncipe
09,Burgos,261,6,Peñaranda de Duero
09,Burgos,262,1,Peral de Arlanza
09,Burgos,265,5,Piérnigas
09,Burgos,266,8,Pineda de la Sierra
09,Burgos,267,4,Pineda Trasmonte
09,Burgos,268,0,Pinilla de los Barruecos
09,Burgos,269,3,Pinilla de los Moros
09,Burgos,270,7,Pinilla Trasmonte
09,Burgos,272,9,Poza de la Sal
09,Burgos,273,5,Prádanos de Bureba
09,Burgos,274,0,Pradoluengo
09,Burgos,275,3,Presencio
09,Burgos,276,6,"Puebla de Arganzón, La"
09,Burgos,277,2,Puentedura
09,Burgos,279,1,Quemada
09,Burgos,281,2,Quintana del Pidio
09,Burgos,280,5,Quintanabureba
09,Burgos,283,3,Quintanaélez
09,Burgos,287,0,Quintanaortuño
09,Burgos,288,6,Quintanapalla
09,Burgos,289,9,Quintanar de la Sierra
09,Burgos,292,5,Quintanavides
09,Burgos,294,6,Quintanilla de la Mata
09,Burgos,901,2,Quintanilla del Agua y Tordueles
09,Burgos,295,9,Quintanilla del Coco
09,Burgos,298,4,Quintanilla San García
09,Burgos,301,8,Quintanilla Vivar
09,Burgos,297,8,"Quintanillas, Las"
09,Burgos,302,3,Rabanera del Pinar
09,Burgos,303,9,Rábanos
09,Burgos,304,4,Rabé de las Calzadas
09,Burgos,306,0,Rebolledo de la Torre
09,Burgos,307,6,Redecilla del Camino
09,Burgos,308,2,Redecilla del Campo
09,Burgos,309,5,Regumiel de la Sierra
09,Burgos,310,9,Reinoso
09,Burgos,311,6,Retuerta
09,Burgos,314,2,Revilla del Campo
09,Burgos,316,8,Revilla Vallejera
09,Burgos,312,1,"Revilla y Ahedo, La"
09,Burgos,315,5,Revillarruz
09,Burgos,317,4,Rezmondo
09,Burgos,318,0,Riocavado de la Sierra
09,Burgos,321,4,Roa
09,Burgos,323,5,Rojas
09,Burgos,325,3,Royuela de Río Franco
09,Burgos,326,6,Rubena
09,Burgos,327,2,Rublacedo de Abajo
09,Burgos,328,8,Rucandio
09,Burgos,329,1,Salas de Bureba
09,Burgos,330,5,Salas de los Infantes
09,Burgos,332,7,Saldaña de Burgos
09,Burgos,334,8,Salinillas de Bureba
09,Burgos,335,1,San Adrián de Juarros
09,Burgos,337,0,San Juan del Monte
09,Burgos,338,6,San Mamés de Burgos
09,Burgos,339,9,San Martín de Rubiales
09,Burgos,340,3,San Millán de Lara
09,Burgos,360,8,San Vicente del Valle
09,Burgos,343,1,Santa Cecilia
09,Burgos,345,9,Santa Cruz de la Salceda
09,Burgos,346,2,Santa Cruz del Valle Urbión
09,Burgos,347,8,Santa Gadea del Cid
09,Burgos,348,4,Santa Inés
09,Burgos,350,0,Santa María del Campo
09,Burgos,351,7,Santa María del Invierno
09,Burgos,352,2,Santa María del Mercadillo
09,Burgos,353,8,Santa María Rivarredonda
09,Burgos,354,3,Santa Olalla de Bureba
09,Burgos,355,6,Santibáñez de Esgueva
09,Burgos,356,9,Santibáñez del Val
09,Burgos,358,1,Santo Domingo de Silos
09,Burgos,361,5,Sargentes de la Lora
09,Burgos,362,0,Sarracín
09,Burgos,363,6,Sasamón
09,Burgos,365,4,"Sequera de Haza, La"
09,Burgos,366,7,Solarana
09,Burgos,368,9,Sordillos
09,Burgos,369,2,Sotillo de la Ribera
09,Burgos,372,8,Sotragero
09,Burgos,373,4,Sotresgudo
09,Burgos,374,9,Susinos del Páramo
09,Burgos,375,2,Tamarón
09,Burgos,377,1,Tardajos
09,Burgos,378,7,Tejada
09,Burgos,380,4,Terradillos de Esgueva
09,Burgos,381,1,Tinieblas de la Sierra
09,Burgos,382,6,Tobar
09,Burgos,384,7,Tordómar
09,Burgos,386,3,Torrecilla del Monte
09,Burgos,387,9,Torregalindo
09,Burgos,388,5,Torrelara
09,Burgos,389,8,Torrepadre
09,Burgos,390,2,Torresandino
09,Burgos,391,9,Tórtoles de Esgueva
09,Burgos,392,4,Tosantos
09,Burgos,394,5,Trespaderne
09,Burgos,395,8,Tubilla del Agua
09,Burgos,396,1,Tubilla del Lago
09,Burgos,398,3,Úrbel del Castillo
09,Burgos,400,0,Vadocondes
09,Burgos,403,8,Valdeande
09,Burgos,405,6,Valdezate
09,Burgos,406,9,Valdorros
09,Burgos,408,1,Vallarta de Bureba
09,Burgos,904,8,Valle de las Navas
09,Burgos,908,6,Valle de Losa
09,Burgos,409,4,Valle de Manzanedo
09,Burgos,410,8,Valle de Mena
09,Burgos,411,5,Valle de Oca
09,Burgos,902,7,Valle de Santibáñez
09,Burgos,905,1,Valle de Sedano
09,Burgos,412,0,Valle de Tobalina
09,Burgos,413,6,Valle de Valdebezana
09,Burgos,414,1,Valle de Valdelaguna
09,Burgos,415,4,Valle de Valdelucio
09,Burgos,416,7,Valle de Zamanzas
09,Burgos,417,3,Vallejera
09,Burgos,418,9,Valles de Palenzuela
09,Burgos,419,2,Valluércanes
09,Burgos,407,5,Valmala
09,Burgos,422,8,"Vid de Bureba, La"
09,Burgos,421,3,"Vid y Barrios, La"
09,Burgos,423,4,Vileña
09,Burgos,427,1,Villadiego
09,Burgos,428,7,Villaescusa de Roa
09,Burgos,429,0,Villaescusa la Sombría
09,Burgos,430,4,Villaespasa
09,Burgos,431,1,Villafranca Montes de Oca
09,Burgos,432,6,Villafruela
09,Burgos,433,2,Villagalijo
09,Burgos,434,7,Villagonzalo Pedernales
09,Burgos,437,9,Villahoz
09,Burgos,438,5,Villalba de Duero
09,Burgos,439,8,Villalbilla de Burgos
09,Burgos,440,2,Villalbilla de Gumiel
09,Burgos,441,9,Villaldemiro
09,Burgos,442,4,Villalmanzo
09,Burgos,443,0,Villamayor de los Montes
09,Burgos,444,5,Villamayor de Treviño
09,Burgos,445,8,Villambistia
09,Burgos,446,1,Villamedianilla
09,Burgos,447,7,Villamiel de la Sierra
09,Burgos,448,3,Villangómez
09,Burgos,449,6,Villanueva de Argaño
09,Burgos,450,9,Villanueva de Carazo
09,Burgos,451,6,Villanueva de Gumiel
09,Burgos,454,2,Villanueva de Teba
09,Burgos,455,5,Villaquirán de la Puebla
09,Burgos,456,8,Villaquirán de los Infantes
09,Burgos,903,3,Villarcayo de Merindad de Castilla la Vieja
09,Burgos,458,0,Villariezo
09,Burgos,460,7,Villasandino
09,Burgos,463,5,Villasur de Herreros
09,Burgos,464,0,Villatuelda
09,Burgos,466,6,Villaverde del Monte
09,Burgos,467,2,Villaverde-Mogina
09,Burgos,471,2,Villayerno Morquillas
09,Burgos,472,7,Villazopeque
09,Burgos,473,3,Villegas
09,Burgos,476,4,Villoruebo
09,Burgos,424,9,Viloria de Rioja
09,Burgos,425,2,Vilviestre del Pinar
09,Burgos,478,6,Vizcaínos
09,Burgos,480,3,Zael
09,Burgos,482,5,Zarzosa de Río Pisuerga
09,Burgos,483,1,Zazuar
09,Burgos,485,9,Zuñeda
10,Cáceres,000,0,Caceres
10,Cáceres,001,5,Abadía
10,Cáceres,002,0,Abertura
10,Cáceres,003,6,Acebo
10,Cáceres,004,1,Acehúche
10,Cáceres,005,4,Aceituna
10,Cáceres,006,7,Ahigal
10,Cáceres,903,7,Alagón del Río
10,Cáceres,007,3,Albalá
10,Cáceres,008,9,Alcántara
10,Cáceres,009,2,Alcollarín
10,Cáceres,010,6,Alcuéscar
10,Cáceres,012,8,Aldea del Cano
10,Cáceres,013,4,"Aldea del Obispo, La"
10,Cáceres,011,3,Aldeacentenera
10,Cáceres,014,9,Aldeanueva de la Vera
10,Cáceres,015,2,Aldeanueva del Camino
10,Cáceres,016,5,Aldehuela de Jerte
10,Cáceres,017,1,Alía
10,Cáceres,018,7,Aliseda
10,Cáceres,019,0,Almaraz
10,Cáceres,020,4,Almoharín
10,Cáceres,021,1,Arroyo de la Luz
10,Cáceres,023,2,Arroyomolinos
10,Cáceres,022,6,Arroyomolinos de la Vera
10,Cáceres,024,7,Baños de Montemayor
10,Cáceres,025,0,Barrado
10,Cáceres,026,3,Belvís de Monroy
10,Cáceres,027,9,Benquerencia
10,Cáceres,028,5,Berrocalejo
10,Cáceres,029,8,Berzocana
10,Cáceres,030,2,Bohonal de Ibor
10,Cáceres,031,9,Botija
10,Cáceres,032,4,Brozas
10,Cáceres,033,0,Cabañas del Castillo
10,Cáceres,034,5,Cabezabellosa
10,Cáceres,035,8,Cabezuela del Valle
10,Cáceres,036,1,Cabrero
10,Cáceres,037,7,Cáceres
10,Cáceres,038,3,Cachorrilla
10,Cáceres,039,6,Cadalso
10,Cáceres,040,0,Calzadilla
10,Cáceres,041,7,Caminomorisco
10,Cáceres,042,2,Campillo de Deleitosa
10,Cáceres,043,8,Campo Lugar
10,Cáceres,044,3,Cañamero
10,Cáceres,045,6,Cañaveral
10,Cáceres,046,9,Carbajo
10,Cáceres,047,5,Carcaboso
10,Cáceres,048,1,Carrascalejo
10,Cáceres,049,4,Casar de Cáceres
10,Cáceres,050,7,Casar de Palomero
10,Cáceres,051,4,Casares de las Hurdes
10,Cáceres,052,9,Casas de Don Antonio
10,Cáceres,053,5,Casas de Don Gómez
10,Cáceres,056,6,Casas de Millán
10,Cáceres,057,2,Casas de Miravete
10,Cáceres,054,0,Casas del Castañar
10,Cáceres,055,3,Casas del Monte
10,Cáceres,058,8,Casatejada
10,Cáceres,059,1,Casillas de Coria
10,Cáceres,060,5,Castañar de Ibor
10,Cáceres,061,2,Ceclavín
10,Cáceres,062,7,Cedillo
10,Cáceres,063,3,Cerezo
10,Cáceres,064,8,Cilleros
10,Cáceres,065,1,Collado
10,Cáceres,066,4,Conquista de la Sierra
10,Cáceres,067,0,Coria
10,Cáceres,068,6,Cuacos de Yuste
10,Cáceres,069,9,"Cumbre, La"
10,Cáceres,070,3,Deleitosa
10,Cáceres,071,0,Descargamaría
10,Cáceres,072,5,Eljas
10,Cáceres,073,1,Escurial
10,Cáceres,075,9,Fresnedoso de Ibor
10,Cáceres,076,2,Galisteo
10,Cáceres,077,8,Garciaz
10,Cáceres,078,4,"Garganta, La"
10,Cáceres,079,7,Garganta la Olla
10,Cáceres,080,1,Gargantilla
10,Cáceres,081,8,Gargüera
10,Cáceres,082,3,Garrovillas de Alconétar
10,Cáceres,083,9,Garvín
10,Cáceres,084,4,Gata
10,Cáceres,085,7,"Gordo, El"
10,Cáceres,086,0,"Granja, La"
10,Cáceres,087,6,Guadalupe
10,Cáceres,088,2,Guijo de Coria
10,Cáceres,089,5,Guijo de Galisteo
10,Cáceres,090,9,Guijo de Granadilla
10,Cáceres,091,6,Guijo de Santa Bárbara
10,Cáceres,092,1,Herguijuela
10,Cáceres,093,7,Hernán-Pérez
10,Cáceres,094,2,Herrera de Alcántara
10,Cáceres,095,5,Herreruela
10,Cáceres,096,8,Hervás
10,Cáceres,097,4,Higuera
10,Cáceres,098,0,Hinojal
10,Cáceres,099,3,Holguera
10,Cáceres,100,7,Hoyos
10,Cáceres,101,4,Huélaga
10,Cáceres,102,9,Ibahernando
10,Cáceres,103,5,Jaraicejo
10,Cáceres,104,0,Jaraíz de la Vera
10,Cáceres,105,3,Jarandilla de la Vera
10,Cáceres,106,6,Jarilla
10,Cáceres,107,2,Jerte
10,Cáceres,108,8,Ladrillar
10,Cáceres,109,1,Logrosán
10,Cáceres,110,5,Losar de la Vera
10,Cáceres,111,2,Madrigal de la Vera
10,Cáceres,112,7,Madrigalejo
10,Cáceres,113,3,Madroñera
10,Cáceres,114,8,Majadas
10,Cáceres,115,1,Malpartida de Cáceres
10,Cáceres,116,4,Malpartida de Plasencia
10,Cáceres,117,0,Marchagaz
10,Cáceres,118,6,Mata de Alcántara
10,Cáceres,119,9,Membrío
10,Cáceres,120,3,Mesas de Ibor
10,Cáceres,121,0,Miajadas
10,Cáceres,122,5,Millanes
10,Cáceres,123,1,Mirabel
10,Cáceres,124,6,Mohedas de Granadilla
10,Cáceres,125,9,Monroy
10,Cáceres,126,2,Montánchez
10,Cáceres,127,8,Montehermoso
10,Cáceres,128,4,Moraleja
10,Cáceres,129,7,Morcillo
10,Cáceres,130,1,Navaconcejo
10,Cáceres,131,8,Navalmoral de la Mata
10,Cáceres,132,3,Navalvillar de Ibor
10,Cáceres,133,9,Navas del Madroño
10,Cáceres,134,4,Navezuelas
10,Cáceres,135,7,Nuñomoral
10,Cáceres,136,0,Oliva de Plasencia
10,Cáceres,137,6,Palomero
10,Cáceres,138,2,Pasarón de la Vera
10,Cáceres,139,5,Pedroso de Acim
10,Cáceres,140,9,Peraleda de la Mata
10,Cáceres,141,6,Peraleda de San Román
10,Cáceres,142,1,Perales del Puerto
10,Cáceres,143,7,Pescueza
10,Cáceres,144,2,"Pesga, La"
10,Cáceres,145,5,Piedras Albas
10,Cáceres,146,8,Pinofranqueado
10,Cáceres,147,4,Piornal
10,Cáceres,148,0,Plasencia
10,Cáceres,149,3,Plasenzuela
10,Cáceres,150,6,Portaje
10,Cáceres,151,3,Portezuelo
10,Cáceres,152,8,Pozuelo de Zarzón
10,Cáceres,153,4,Puerto de Santa Cruz
10,Cáceres,154,9,Rebollar
10,Cáceres,155,2,Riolobos
10,Cáceres,156,5,Robledillo de Gata
10,Cáceres,157,1,Robledillo de la Vera
10,Cáceres,158,7,Robledillo de Trujillo
10,Cáceres,159,0,Robledollano
10,Cáceres,160,4,Romangordo
10,Cáceres,901,6,Rosalejo
10,Cáceres,161,1,Ruanes
10,Cáceres,162,6,Salorino
10,Cáceres,163,2,Salvatierra de Santiago
10,Cáceres,164,7,San Martín de Trevejo
10,Cáceres,165,0,Santa Ana
10,Cáceres,166,3,Santa Cruz de la Sierra
10,Cáceres,167,9,Santa Cruz de Paniagua
10,Cáceres,168,5,Santa Marta de Magasca
10,Cáceres,169,8,Santiago de Alcántara
10,Cáceres,170,2,Santiago del Campo
10,Cáceres,171,9,Santibáñez el Alto
10,Cáceres,172,4,Santibáñez el Bajo
10,Cáceres,173,0,Saucedilla
10,Cáceres,174,5,Segura de Toro
10,Cáceres,175,8,Serradilla
10,Cáceres,176,1,Serrejón
10,Cáceres,177,7,Sierra de Fuentes
10,Cáceres,178,3,Talaván
10,Cáceres,179,6,Talaveruela de la Vera
10,Cáceres,180,0,Talayuela
10,Cáceres,181,7,Tejeda de Tiétar
10,Cáceres,182,2,Toril
10,Cáceres,183,8,Tornavacas
10,Cáceres,184,3,"Torno, El"
10,Cáceres,187,5,Torre de Don Miguel
10,Cáceres,188,1,Torre de Santa María
10,Cáceres,185,6,Torrecilla de los Ángeles
10,Cáceres,186,9,Torrecillas de la Tiesa
10,Cáceres,190,8,Torrejón el Rubio
10,Cáceres,189,4,Torrejoncillo
10,Cáceres,191,5,Torremenga
10,Cáceres,192,0,Torremocha
10,Cáceres,193,6,Torreorgaz
10,Cáceres,194,1,Torrequemada
10,Cáceres,195,4,Trujillo
10,Cáceres,196,7,Valdastillas
10,Cáceres,197,3,Valdecañas de Tajo
10,Cáceres,198,9,Valdefuentes
10,Cáceres,199,2,Valdehúncar
10,Cáceres,200,6,Valdelacasa de Tajo
10,Cáceres,201,3,Valdemorales
10,Cáceres,202,8,Valdeobispo
10,Cáceres,203,4,Valencia de Alcántara
10,Cáceres,204,9,Valverde de la Vera
10,Cáceres,205,2,Valverde del Fresno
10,Cáceres,902,1,Vegaviana
10,Cáceres,206,5,Viandar de la Vera
10,Cáceres,207,1,Villa del Campo
10,Cáceres,208,7,Villa del Rey
10,Cáceres,209,0,Villamesías
10,Cáceres,210,4,Villamiel
10,Cáceres,211,1,Villanueva de la Sierra
10,Cáceres,212,6,Villanueva de la Vera
10,Cáceres,214,7,Villar de Plasencia
10,Cáceres,213,2,Villar del Pedroso
10,Cáceres,215,0,Villasbuenas de Gata
10,Cáceres,216,3,Zarza de Granadilla
10,Cáceres,217,9,Zarza de Montánchez
10,Cáceres,218,5,Zarza la Mayor
10,Cáceres,219,8,Zorita
11,Cádiz,000,0,Sanlucar de Barrameda
11,Cádiz,000,0,Cadiz
11,Cádiz,001,2,Alcalá de los Gazules
11,Cádiz,002,7,Alcalá del Valle
11,Cádiz,003,3,Algar
11,Cádiz,004,8,Algeciras
11,Cádiz,005,1,Algodonales
11,Cádiz,006,4,Arcos de la Frontera
11,Cádiz,007,0,Barbate
11,Cádiz,008,6,"Barrios, Los"
11,Cádiz,901,3,Benalup-Casas Viejas
11,Cádiz,009,9,Benaocaz
11,Cádiz,010,3,Bornos
11,Cádiz,011,0,"Bosque, El"
11,Cádiz,012,5,Cádiz
11,Cádiz,013,1,Castellar de la Frontera
11,Cádiz,015,9,Chiclana de la Frontera
11,Cádiz,016,2,Chipiona
11,Cádiz,014,6,Conil de la Frontera
11,Cádiz,017,8,Espera
11,Cádiz,018,4,"Gastor, El"
11,Cádiz,019,7,Grazalema
11,Cádiz,020,1,Jerez de la Frontera
11,Cádiz,021,8,Jimena de la Frontera
11,Cádiz,022,3,"Línea de la Concepción, La"
11,Cádiz,023,9,Medina-Sidonia
11,Cádiz,024,4,Olvera
11,Cádiz,025,7,Paterna de Rivera
11,Cádiz,026,0,Prado del Rey
11,Cádiz,027,6,"Puerto de Santa María, El"
11,Cádiz,028,2,Puerto Real
11,Cádiz,029,5,Puerto Serrano
11,Cádiz,030,9,Rota
11,Cádiz,031,6,San Fernando
11,Cádiz,902,8,San José del Valle
11,Cádiz,033,7,San Roque
11,Cádiz,032,1,Sanlúcar de Barrameda
11,Cádiz,034,2,Setenil de las Bodegas
11,Cádiz,035,5,Tarifa
11,Cádiz,036,8,Torre Alháquime
11,Cádiz,037,4,Trebujena
11,Cádiz,038,0,Ubrique
11,Cádiz,039,3,Vejer de la Frontera
11,Cádiz,040,7,Villaluenga del Rosario
11,Cádiz,041,4,Villamartín
11,Cádiz,042,9,Zahara
12,Castellón,000,0,Castellon
12,Castellón,002,2,Aín
12,Castellón,003,8,Albocàsser
12,Castellón,004,3,Alcalà de Xivert
12,Castellón,005,6,"Alcora, l'"
12,Castellón,006,9,Alcudia de Veo
12,Castellón,007,5,Alfondeguilla
12,Castellón,008,1,Algimia de Almonacid
12,Castellón,009,4,Almazora/Almassora
12,Castellón,010,8,Almedíjar
12,Castellón,011,5,Almenara
12,Castellón,901,8,Alquerías del Niño Perdido
12,Castellón,012,0,Altura
12,Castellón,013,6,Arañuel
12,Castellón,014,1,Ares del Maestrat
12,Castellón,015,4,Argelita
12,Castellón,016,7,Artana
12,Castellón,001,7,Atzeneta del Maestrat
12,Castellón,017,3,Ayódar
12,Castellón,018,9,Azuébar
12,Castellón,020,6,Barracas
12,Castellón,022,8,Bejís
12,Castellón,024,9,Benafer
12,Castellón,025,2,Benafigos
12,Castellón,026,5,Benasal
12,Castellón,027,1,Benicarló
12,Castellón,028,7,Benicasim/Benicàssim
12,Castellón,029,0,Benlloch
12,Castellón,021,3,Betxí
12,Castellón,032,6,Borriana/Burriana
12,Castellón,031,1,Borriol
12,Castellón,033,2,Cabanes
12,Castellón,034,7,Càlig
12,Castellón,036,3,Canet lo Roig
12,Castellón,037,9,Castell de Cabres
12,Castellón,038,5,Castellfort
12,Castellón,039,8,Castellnovo
12,Castellón,040,2,Castellón de la Plana/Castelló de la Plana
12,Castellón,041,9,Castillo de Villamalefa
12,Castellón,042,4,Catí
12,Castellón,043,0,Caudiel
12,Castellón,044,5,Cervera del Maestre
12,Castellón,052,1,Chert/Xert
12,Castellón,053,7,Chilches/Xilxes
12,Castellón,055,5,Chodos/Xodos
12,Castellón,056,8,Chóvar
12,Castellón,045,8,Cinctorres
12,Castellón,046,1,Cirat
12,Castellón,048,3,Cortes de Arenoso
12,Castellón,049,6,Costur
12,Castellón,050,9,"Coves de Vinromà, les"
12,Castellón,051,6,Culla
12,Castellón,057,4,Eslida
12,Castellón,058,0,Espadilla
12,Castellón,059,3,Fanzara
12,Castellón,060,7,Figueroles
12,Castellón,061,4,Forcall
12,Castellón,063,5,Fuente la Reina
12,Castellón,064,0,Fuentes de Ayódar
12,Castellón,065,3,Gaibiel
12,Castellón,067,2,Geldo
12,Castellón,068,8,Herbés
12,Castellón,069,1,Higueras
12,Castellón,070,5,"Jana, la"
12,Castellón,071,2,Jérica
12,Castellón,074,8,"Llosa, la"
12,Castellón,072,7,Lucena del Cid
12,Castellón,073,3,Ludiente
12,Castellón,075,1,"Mata de Morella, la"
12,Castellón,076,4,Matet
12,Castellón,077,0,Moncofa
12,Castellón,078,6,Montán
12,Castellón,079,9,Montanejos
12,Castellón,080,3,Morella
12,Castellón,081,0,Navajas
12,Castellón,082,5,Nules
12,Castellón,083,1,Olocau del Rey
12,Castellón,084,6,Onda
12,Castellón,085,9,Oropesa del Mar/Orpesa
12,Castellón,087,8,Palanques
12,Castellón,088,4,Pavías
12,Castellón,089,7,Peníscola/Peñíscola
12,Castellón,090,1,Pina de Montalgrao
12,Castellón,093,9,"Pobla de Benifassà, la"
12,Castellón,094,4,"Pobla Tornesa, la"
12,Castellón,091,8,Portell de Morella
12,Castellón,092,3,Puebla de Arenoso
12,Castellón,095,7,Ribesalbes
12,Castellón,096,0,Rossell
12,Castellón,097,6,Sacañet
12,Castellón,098,2,"Salzadella, la"
12,Castellón,101,6,San Rafael del Río
12,Castellón,902,3,Sant Joan de Moró
12,Castellón,099,5,Sant Jordi/San Jorge
12,Castellón,100,9,Sant Mateu
12,Castellón,102,1,Santa Magdalena de Pulpis
12,Castellón,103,7,Sarratella
12,Castellón,104,2,Segorbe
12,Castellón,105,5,Sierra Engarcerán
12,Castellón,106,8,Soneja
12,Castellón,107,4,Sot de Ferrer
12,Castellón,108,0,Sueras/Suera
12,Castellón,109,3,Tales
12,Castellón,110,7,Teresa
12,Castellón,111,4,Tírig
12,Castellón,112,9,Todolella
12,Castellón,113,5,Toga
12,Castellón,114,0,Torás
12,Castellón,115,3,"Toro, El"
12,Castellón,116,6,Torralba del Pinar
12,Castellón,119,1,"Torre d'En Besora, la"
12,Castellón,120,5,"Torre d'en Doménec, la"
12,Castellón,117,2,Torreblanca
12,Castellón,118,8,Torrechiva
12,Castellón,121,2,Traiguera
12,Castellón,122,7,"Useras/Useres, les"
12,Castellón,124,8,Vall d'Alba
12,Castellón,125,1,Vall de Almonacid
12,Castellón,126,4,"Vall d'Uixó, la"
12,Castellón,123,3,Vallat
12,Castellón,127,0,Vallibona
12,Castellón,128,6,Vilafamés
12,Castellón,132,5,Vilanova d'Alcolea
12,Castellón,134,6,Vilar de Canes
12,Castellón,135,9,Vila-real
12,Castellón,136,2,"Vilavella, la"
12,Castellón,129,9,Villafranca del Cid/Vilafranca
12,Castellón,130,3,Villahermosa del Río
12,Castellón,131,0,Villamalur
12,Castellón,133,1,Villanueva de Viver
12,Castellón,137,8,Villores
12,Castellón,138,4,Vinaròs
12,Castellón,139,7,Vistabella del Maestrazgo
12,Castellón,140,1,Viver
12,Castellón,141,8,Zorita del Maestrazgo
12,Castellón,142,3,Zucaina
13,Ciudad Real,001,3,Abenójar
13,Ciudad Real,002,8,Agudo
13,Ciudad Real,003,4,Alamillo
13,Ciudad Real,004,9,Albaladejo
13,Ciudad Real,005,2,Alcázar de San Juan
13,Ciudad Real,006,5,Alcoba
13,Ciudad Real,007,1,Alcolea de Calatrava
13,Ciudad Real,008,7,Alcubillas
13,Ciudad Real,009,0,Aldea del Rey
13,Ciudad Real,010,4,Alhambra
13,Ciudad Real,011,1,Almadén
13,Ciudad Real,012,6,Almadenejos
13,Ciudad Real,013,2,Almagro
13,Ciudad Real,014,7,Almedina
13,Ciudad Real,015,0,Almodóvar del Campo
13,Ciudad Real,016,3,Almuradiel
13,Ciudad Real,017,9,Anchuras
13,Ciudad Real,903,5,Arenales de San Gregorio
13,Ciudad Real,018,5,Arenas de San Juan
13,Ciudad Real,019,8,Argamasilla de Alba
13,Ciudad Real,020,2,Argamasilla de Calatrava
13,Ciudad Real,021,9,Arroba de los Montes
13,Ciudad Real,022,4,Ballesteros de Calatrava
13,Ciudad Real,023,0,Bolaños de Calatrava
13,Ciudad Real,024,5,Brazatortas
13,Ciudad Real,025,8,Cabezarados
13,Ciudad Real,026,1,Cabezarrubias del Puerto
13,Ciudad Real,027,7,Calzada de Calatrava
13,Ciudad Real,028,3,Campo de Criptana
13,Ciudad Real,029,6,Cañada de Calatrava
13,Ciudad Real,030,0,Caracuel de Calatrava
13,Ciudad Real,031,7,Carrión de Calatrava
13,Ciudad Real,032,2,Carrizosa
13,Ciudad Real,033,8,Castellar de Santiago
13,Ciudad Real,038,1,Chillón
13,Ciudad Real,034,3,Ciudad Real
13,Ciudad Real,035,6,Corral de Calatrava
13,Ciudad Real,036,9,"Cortijos, Los"
13,Ciudad Real,037,5,Cózar
13,Ciudad Real,039,4,Daimiel
13,Ciudad Real,040,8,Fernán Caballero
13,Ciudad Real,041,5,Fontanarejo
13,Ciudad Real,042,0,Fuencaliente
13,Ciudad Real,043,6,Fuenllana
13,Ciudad Real,044,1,Fuente el Fresno
13,Ciudad Real,045,4,Granátula de Calatrava
13,Ciudad Real,046,7,Guadalmez
13,Ciudad Real,047,3,Herencia
13,Ciudad Real,048,9,Hinojosas de Calatrava
13,Ciudad Real,049,2,Horcajo de los Montes
13,Ciudad Real,050,5,"Labores, Las"
13,Ciudad Real,904,0,Llanos del Caudillo
13,Ciudad Real,051,2,Luciana
13,Ciudad Real,052,7,Malagón
13,Ciudad Real,053,3,Manzanares
13,Ciudad Real,054,8,Membrilla
13,Ciudad Real,055,1,Mestanza
13,Ciudad Real,056,4,Miguelturra
13,Ciudad Real,057,0,Montiel
13,Ciudad Real,058,6,Moral de Calatrava
13,Ciudad Real,059,9,Navalpino
13,Ciudad Real,060,3,Navas de Estena
13,Ciudad Real,061,0,Pedro Muñoz
13,Ciudad Real,062,5,Picón
13,Ciudad Real,063,1,Piedrabuena
13,Ciudad Real,064,6,Poblete
13,Ciudad Real,065,9,Porzuna
13,Ciudad Real,066,2,Pozuelo de Calatrava
13,Ciudad Real,067,8,"Pozuelos de Calatrava, Los"
13,Ciudad Real,068,4,Puebla de Don Rodrigo
13,Ciudad Real,069,7,Puebla del Príncipe
13,Ciudad Real,070,1,Puerto Lápice
13,Ciudad Real,071,8,Puertollano
13,Ciudad Real,072,3,Retuerta del Bullaque
13,Ciudad Real,901,4,"Robledo, El"
13,Ciudad Real,902,9,Ruidera
13,Ciudad Real,073,9,Saceruela
13,Ciudad Real,074,4,San Carlos del Valle
13,Ciudad Real,075,7,San Lorenzo de Calatrava
13,Ciudad Real,076,0,Santa Cruz de los Cáñamos
13,Ciudad Real,077,6,Santa Cruz de Mudela
13,Ciudad Real,078,2,Socuéllamos
13,Ciudad Real,079,5,"Solana, La"
13,Ciudad Real,080,9,Solana del Pino
13,Ciudad Real,081,6,Terrinches
13,Ciudad Real,082,1,Tomelloso
13,Ciudad Real,083,7,Torralba de Calatrava
13,Ciudad Real,084,2,Torre de Juan Abad
13,Ciudad Real,085,5,Torrenueva
13,Ciudad Real,086,8,Valdemanco del Esteras
13,Ciudad Real,087,4,Valdepeñas
13,Ciudad Real,088,0,Valenzuela de Calatrava
13,Ciudad Real,089,3,Villahermosa
13,Ciudad Real,090,7,Villamanrique
13,Ciudad Real,091,4,Villamayor de Calatrava
13,Ciudad Real,092,9,Villanueva de la Fuente
13,Ciudad Real,093,5,Villanueva de los Infantes
13,Ciudad Real,094,0,Villanueva de San Carlos
13,Ciudad Real,095,3,Villar del Pozo
13,Ciudad Real,096,6,Villarrubia de los Ojos
13,Ciudad Real,097,2,Villarta de San Juan
13,Ciudad Real,098,8,Viso del Marqués
14,Córdoba,001,8,Cordoba
14,Córdoba,001,8,Crdoba
14,Córdoba,001,8,Adamuz
14,Córdoba,002,3,Aguilar de la Frontera
14,Córdoba,003,9,Alcaracejos
14,Córdoba,004,4,Almedinilla
14,Córdoba,005,7,Almodóvar del Río
14,Córdoba,006,0,Añora
14,Córdoba,007,6,Baena
14,Córdoba,008,2,Belalcázar
14,Córdoba,009,5,Belmez
14,Córdoba,010,9,Benamejí
14,Córdoba,011,6,"Blázquez, Los"
14,Córdoba,012,1,Bujalance
14,Córdoba,013,7,Cabra
14,Córdoba,014,2,Cañete de las Torres
14,Córdoba,015,5,Carcabuey
14,Córdoba,016,8,Cardeña
14,Córdoba,017,4,"Carlota, La"
14,Córdoba,018,0,"Carpio, El"
14,Córdoba,019,3,Castro del Río
14,Córdoba,020,7,Conquista
14,Córdoba,021,4,Córdoba
14,Córdoba,022,9,Doña Mencía
14,Córdoba,023,5,Dos Torres
14,Córdoba,024,0,Encinas Reales
14,Córdoba,025,3,Espejo
14,Córdoba,026,6,Espiel
14,Córdoba,027,2,Fernán-Núñez
14,Córdoba,028,8,Fuente la Lancha
14,Córdoba,029,1,Fuente Obejuna
14,Córdoba,030,5,Fuente Palmera
14,Córdoba,031,2,Fuente-Tójar
14,Córdoba,032,7,"Granjuela, La"
14,Córdoba,033,3,Guadalcázar
14,Córdoba,034,8,"Guijo, El"
14,Córdoba,035,1,Hinojosa del Duque
14,Córdoba,036,4,Hornachuelos
14,Córdoba,037,0,Iznájar
14,Córdoba,038,6,Lucena
14,Córdoba,039,9,Luque
14,Córdoba,040,3,Montalbán de Córdoba
14,Córdoba,041,0,Montemayor
14,Córdoba,042,5,Montilla
14,Córdoba,043,1,Montoro
14,Córdoba,044,6,Monturque
14,Córdoba,045,9,Moriles
14,Córdoba,046,2,Nueva Carteya
14,Córdoba,047,8,Obejo
14,Córdoba,048,4,Palenciana
14,Córdoba,049,7,Palma del Río
14,Córdoba,050,0,Pedro Abad
14,Córdoba,051,7,Pedroche
14,Córdoba,052,2,Peñarroya-Pueblonuevo
14,Córdoba,053,8,Posadas
14,Córdoba,054,3,Pozoblanco
14,Córdoba,055,6,Priego de Córdoba
14,Córdoba,056,9,Puente Genil
14,Córdoba,057,5,"Rambla, La"
14,Córdoba,058,1,Rute
14,Córdoba,059,4,San Sebastián de los Ballesteros
14,Córdoba,061,5,Santa Eufemia
14,Córdoba,060,8,Santaella
14,Córdoba,062,0,Torrecampo
14,Córdoba,063,6,Valenzuela
14,Córdoba,064,1,Valsequillo
14,Córdoba,065,4,"Victoria, La"
14,Córdoba,066,7,Villa del Río
14,Córdoba,067,3,Villafranca de Córdoba
14,Córdoba,068,9,Villaharta
14,Córdoba,069,2,Villanueva de Córdoba
14,Córdoba,070,6,Villanueva del Duque
14,Córdoba,071,3,Villanueva del Rey
14,Córdoba,072,8,Villaralto
14,Córdoba,073,4,Villaviciosa de Córdoba
14,Córdoba,074,9,"Viso, El"
14,Córdoba,075,2,Zuheros
15,A Coruña,000,0,A Coruna
15,A Coruña,000,0,A Corua
15,A Coruña,000,0,La Corua
15,A Coruña,001,1,Abegondo
15,A Coruña,002,6,Ames
15,A Coruña,003,2,Aranga
15,A Coruña,004,7,Ares
15,A Coruña,005,0,Arteixo
15,A Coruña,006,3,Arzúa
15,A Coruña,007,9,"Baña, A"
15,A Coruña,008,5,Bergondo
15,A Coruña,009,8,Betanzos
15,A Coruña,010,2,Boimorto
15,A Coruña,011,9,Boiro
15,A Coruña,012,4,Boqueixón
15,A Coruña,013,0,Brión
15,A Coruña,014,5,Cabana de Bergantiños
15,A Coruña,015,8,Cabanas
15,A Coruña,016,1,Camariñas
15,A Coruña,017,7,Cambre
15,A Coruña,018,3,"Capela, A"
15,A Coruña,019,6,Carballo
15,A Coruña,901,2,Cariño
15,A Coruña,020,0,Carnota
15,A Coruña,021,7,Carral
15,A Coruña,022,2,Cedeira
15,A Coruña,023,8,Cee
15,A Coruña,024,3,Cerceda
15,A Coruña,025,6,Cerdido
15,A Coruña,026,9,Cesuras
15,A Coruña,027,5,Coirós
15,A Coruña,028,1,Corcubión
15,A Coruña,029,4,Coristanco
15,A Coruña,030,8,"Coruña, A"
15,A Coruña,031,5,Culleredo
15,A Coruña,032,0,Curtis
15,A Coruña,033,6,Dodro
15,A Coruña,034,1,Dumbría
15,A Coruña,035,4,Fene
15,A Coruña,036,7,Ferrol
15,A Coruña,037,3,Fisterra
15,A Coruña,038,9,Frades
15,A Coruña,039,2,Irixoa
15,A Coruña,041,3,"Laracha, A"
15,A Coruña,040,6,Laxe
15,A Coruña,042,8,Lousame
15,A Coruña,043,4,Malpica de Bergantiños
15,A Coruña,044,9,Mañón
15,A Coruña,045,2,Mazaricos
15,A Coruña,046,5,Melide
15,A Coruña,047,1,Mesía
15,A Coruña,048,7,Miño
15,A Coruña,049,0,Moeche
15,A Coruña,050,3,Monfero
15,A Coruña,051,0,Mugardos
15,A Coruña,053,1,Muros
15,A Coruña,052,5,Muxía
15,A Coruña,054,6,Narón
15,A Coruña,055,9,Neda
15,A Coruña,056,2,Negreira
15,A Coruña,057,8,Noia
15,A Coruña,058,4,Oleiros
15,A Coruña,059,7,Ordes
15,A Coruña,060,1,Oroso
15,A Coruña,061,8,Ortigueira
15,A Coruña,062,3,Outes
15,A Coruña,063,9,Oza dos Ríos
15,A Coruña,064,4,Paderne
15,A Coruña,065,7,Padrón
15,A Coruña,066,0,"Pino, O"
15,A Coruña,067,6,"Pobra do Caramiñal, A"
15,A Coruña,068,2,Ponteceso
15,A Coruña,069,5,Pontedeume
15,A Coruña,070,9,"Pontes de García Rodríguez, As"
15,A Coruña,071,6,Porto do Son
15,A Coruña,072,1,Rianxo
15,A Coruña,073,7,Ribeira
15,A Coruña,074,2,Rois
15,A Coruña,075,5,Sada
15,A Coruña,076,8,San Sadurniño
15,A Coruña,077,4,Santa Comba
15,A Coruña,078,0,Santiago de Compostela
15,A Coruña,079,3,Santiso
15,A Coruña,080,7,Sobrado
15,A Coruña,081,4,"Somozas, As"
15,A Coruña,082,9,Teo
15,A Coruña,083,5,Toques
15,A Coruña,084,0,Tordoia
15,A Coruña,085,3,Touro
15,A Coruña,086,6,Trazo
15,A Coruña,088,8,Val do Dubra
15,A Coruña,087,2,Valdoviño
15,A Coruña,089,1,Vedra
15,A Coruña,091,2,Vilarmaior
15,A Coruña,090,5,Vilasantar
15,A Coruña,092,7,Vimianzo
15,A Coruña,093,3,Zas
16,Cuenca,001,4,Abia de la Obispalía
16,Cuenca,002,9,"Acebrón, El"
16,Cuenca,003,5,Alarcón
16,Cuenca,004,0,Albaladejo del Cuende
16,Cuenca,005,3,Albalate de las Nogueras
16,Cuenca,006,6,Albendea
16,Cuenca,007,2,"Alberca de Záncara, La"
16,Cuenca,008,8,Alcalá de la Vega
16,Cuenca,009,1,Alcantud
16,Cuenca,010,5,Alcázar del Rey
16,Cuenca,011,2,Alcohujate
16,Cuenca,012,7,Alconchel de la Estrella
16,Cuenca,013,3,Algarra
16,Cuenca,014,8,Aliaguilla
16,Cuenca,015,1,"Almarcha, La"
16,Cuenca,016,4,Almendros
16,Cuenca,017,0,Almodóvar del Pinar
16,Cuenca,018,6,Almonacid del Marquesado
16,Cuenca,019,9,Altarejos
16,Cuenca,020,3,Arandilla del Arroyo
16,Cuenca,905,4,Arcas del Villar
16,Cuenca,022,5,Arcos de la Sierra
16,Cuenca,024,6,Arguisuelas
16,Cuenca,025,9,Arrancacepas
16,Cuenca,026,2,Atalaya del Cañavate
16,Cuenca,027,8,Barajas de Melo
16,Cuenca,029,7,Barchín del Hoyo
16,Cuenca,030,1,Bascuñana de San Pedro
16,Cuenca,031,8,Beamud
16,Cuenca,032,3,Belinchón
16,Cuenca,033,9,Belmonte
16,Cuenca,034,4,Belmontejo
16,Cuenca,035,7,Beteta
16,Cuenca,036,0,Boniches
16,Cuenca,038,2,Buciegas
16,Cuenca,039,5,Buenache de Alarcón
16,Cuenca,040,9,Buenache de la Sierra
16,Cuenca,041,6,Buendía
16,Cuenca,042,1,Campillo de Altobuey
16,Cuenca,043,7,Campillos-Paravientos
16,Cuenca,044,2,Campillos-Sierra
16,Cuenca,901,5,Campos del Paraíso
16,Cuenca,045,5,Canalejas del Arroyo
16,Cuenca,046,8,Cañada del Hoyo
16,Cuenca,047,4,Cañada Juncosa
16,Cuenca,048,0,Cañamares
16,Cuenca,049,3,"Cañavate, El"
16,Cuenca,050,6,Cañaveras
16,Cuenca,051,3,Cañaveruelas
16,Cuenca,052,8,Cañete
16,Cuenca,053,4,Cañizares
16,Cuenca,055,2,Carboneras de Guadazaón
16,Cuenca,056,5,Cardenete
16,Cuenca,057,1,Carrascosa
16,Cuenca,058,7,Carrascosa de Haro
16,Cuenca,060,4,Casas de Benítez
16,Cuenca,061,1,Casas de Fernando Alonso
16,Cuenca,062,6,Casas de Garcimolina
16,Cuenca,063,2,Casas de Guijarro
16,Cuenca,064,7,Casas de Haro
16,Cuenca,065,0,Casas de los Pinos
16,Cuenca,066,3,Casasimarro
16,Cuenca,067,9,Castejón
16,Cuenca,068,5,Castillejo de Iniesta
16,Cuenca,070,2,Castillejo-Sierra
16,Cuenca,072,4,Castillo de Garcimuñoz
16,Cuenca,071,9,Castillo-Albaráñez
16,Cuenca,073,0,Cervera del Llano
16,Cuenca,023,1,Chillarón de Cuenca
16,Cuenca,081,7,Chumillas
16,Cuenca,074,5,"Cierva, La"
16,Cuenca,078,3,Cuenca
16,Cuenca,079,6,Cueva del Hierro
16,Cuenca,082,2,Enguídanos
16,Cuenca,083,8,Fresneda de Altarejos
16,Cuenca,084,3,Fresneda de la Sierra
16,Cuenca,085,6,"Frontera, La"
16,Cuenca,086,9,Fuente de Pedro Naharro
16,Cuenca,087,5,Fuentelespino de Haro
16,Cuenca,088,1,Fuentelespino de Moya
16,Cuenca,904,1,Fuentenava de Jábaga
16,Cuenca,089,4,Fuentes
16,Cuenca,091,5,Fuertescusa
16,Cuenca,092,0,Gabaldón
16,Cuenca,093,6,Garaballa
16,Cuenca,094,1,Gascueña
16,Cuenca,095,4,Graja de Campalbo
16,Cuenca,096,7,Graja de Iniesta
16,Cuenca,097,3,Henarejos
16,Cuenca,098,9,"Herrumblar, El"
16,Cuenca,099,2,"Hinojosa, La"
16,Cuenca,100,6,"Hinojosos, Los"
16,Cuenca,101,3,"Hito, El"
16,Cuenca,102,8,Honrubia
16,Cuenca,103,4,Hontanaya
16,Cuenca,104,9,Hontecillas
16,Cuenca,106,5,Horcajo de Santiago
16,Cuenca,107,1,Huélamo
16,Cuenca,108,7,Huelves
16,Cuenca,109,0,Huérguina
16,Cuenca,110,4,Huerta de la Obispalía
16,Cuenca,111,1,Huerta del Marquesado
16,Cuenca,112,6,Huete
16,Cuenca,113,2,Iniesta
16,Cuenca,115,0,Laguna del Marquesado
16,Cuenca,116,3,Lagunaseca
16,Cuenca,117,9,Landete
16,Cuenca,118,5,Ledaña
16,Cuenca,119,8,Leganiel
16,Cuenca,121,9,"Majadas, Las"
16,Cuenca,122,4,Mariana
16,Cuenca,123,0,Masegosa
16,Cuenca,124,5,"Mesas, Las"
16,Cuenca,125,8,Minglanilla
16,Cuenca,126,1,Mira
16,Cuenca,128,3,Monreal del Llano
16,Cuenca,129,6,Montalbanejo
16,Cuenca,130,0,Montalbo
16,Cuenca,131,7,Monteagudo de las Salinas
16,Cuenca,132,2,Mota de Altarejos
16,Cuenca,133,8,Mota del Cuervo
16,Cuenca,134,3,Motilla del Palancar
16,Cuenca,135,6,Moya
16,Cuenca,137,5,Narboneta
16,Cuenca,139,4,Olivares de Júcar
16,Cuenca,140,8,Olmeda de la Cuesta
16,Cuenca,141,5,Olmeda del Rey
16,Cuenca,142,0,Olmedilla de Alarcón
16,Cuenca,143,6,Olmedilla de Eliz
16,Cuenca,145,4,Osa de la Vega
16,Cuenca,146,7,Pajarón
16,Cuenca,147,3,Pajaroncillo
16,Cuenca,148,9,Palomares del Campo
16,Cuenca,149,2,Palomera
16,Cuenca,150,5,Paracuellos
16,Cuenca,151,2,Paredes
16,Cuenca,152,7,"Parra de las Vegas, La"
16,Cuenca,153,3,"Pedernoso, El"
16,Cuenca,154,8,"Pedroñeras, Las"
16,Cuenca,155,1,"Peral, El"
16,Cuenca,156,4,"Peraleja, La"
16,Cuenca,157,0,"Pesquera, La"
16,Cuenca,158,6,"Picazo, El"
16,Cuenca,159,9,Pinarejo
16,Cuenca,160,3,Pineda de Gigüela
16,Cuenca,161,0,Piqueras del Castillo
16,Cuenca,162,5,Portalrubio de Guadamejud
16,Cuenca,163,1,Portilla
16,Cuenca,165,9,Poyatos
16,Cuenca,166,2,Pozoamargo
16,Cuenca,908,9,Pozorrubielos de la Mancha
16,Cuenca,167,8,Pozorrubio
16,Cuenca,169,7,"Pozuelo, El"
16,Cuenca,170,1,Priego
16,Cuenca,171,8,"Provencio, El"
16,Cuenca,172,3,Puebla de Almenara
16,Cuenca,174,4,Puebla del Salvador
16,Cuenca,175,7,Quintanar del Rey
16,Cuenca,176,0,Rada de Haro
16,Cuenca,177,6,Reíllo
16,Cuenca,181,6,Rozalén del Monte
16,Cuenca,185,5,Saceda-Trasierra
16,Cuenca,186,8,Saelices
16,Cuenca,187,4,Salinas del Manzano
16,Cuenca,188,0,Salmeroncillos
16,Cuenca,189,3,Salvacañete
16,Cuenca,190,7,San Clemente
16,Cuenca,191,4,San Lorenzo de la Parrilla
16,Cuenca,192,9,San Martín de Boniches
16,Cuenca,193,5,San Pedro Palmiches
16,Cuenca,194,0,Santa Cruz de Moya
16,Cuenca,196,6,Santa María de los Llanos
16,Cuenca,195,3,Santa María del Campo Rus
16,Cuenca,197,2,Santa María del Val
16,Cuenca,198,8,Sisante
16,Cuenca,199,1,Solera de Gabaldón
16,Cuenca,909,2,Sotorribas
16,Cuenca,202,7,Talayuelas
16,Cuenca,203,3,Tarancón
16,Cuenca,204,8,Tébar
16,Cuenca,205,1,Tejadillos
16,Cuenca,206,4,Tinajas
16,Cuenca,209,9,Torralba
16,Cuenca,211,0,Torrejoncillo del Rey
16,Cuenca,212,5,Torrubia del Campo
16,Cuenca,213,1,Torrubia del Castillo
16,Cuenca,215,9,Tragacete
16,Cuenca,216,2,Tresjuncos
16,Cuenca,217,8,Tribaldos
16,Cuenca,218,4,Uclés
16,Cuenca,219,7,Uña
16,Cuenca,906,7,"Valdecolmenas, Los"
16,Cuenca,224,4,Valdemeca
16,Cuenca,225,7,Valdemorillo de la Sierra
16,Cuenca,227,6,Valdemoro-Sierra
16,Cuenca,228,2,Valdeolivas
16,Cuenca,902,0,Valdetórtola
16,Cuenca,903,6,"Valeras, Las"
16,Cuenca,231,6,Valhermoso de la Fuente
16,Cuenca,173,9,"Valle de Altomira, El"
16,Cuenca,234,2,Valsalobre
16,Cuenca,236,8,Valverde de Júcar
16,Cuenca,237,4,Valverdejo
16,Cuenca,238,0,Vara de Rey
16,Cuenca,239,3,Vega del Codorno
16,Cuenca,240,7,Vellisca
16,Cuenca,242,9,Villaconejos de Trabaque
16,Cuenca,243,5,Villaescusa de Haro
16,Cuenca,244,0,Villagarcía del Llano
16,Cuenca,245,3,Villalba de la Sierra
16,Cuenca,246,6,Villalba del Rey
16,Cuenca,247,2,Villalgordo del Marquesado
16,Cuenca,248,8,Villalpardo
16,Cuenca,249,1,Villamayor de Santiago
16,Cuenca,250,4,Villanueva de Guadamejud
16,Cuenca,251,1,Villanueva de la Jara
16,Cuenca,253,2,Villar de Cañas
16,Cuenca,254,7,Villar de Domingo García
16,Cuenca,255,0,Villar de la Encina
16,Cuenca,263,0,Villar de Olalla
16,Cuenca,258,5,Villar del Humo
16,Cuenca,259,8,Villar del Infantado
16,Cuenca,910,6,Villar y Velasco
16,Cuenca,264,5,Villarejo de Fuentes
16,Cuenca,265,8,Villarejo de la Peñuela
16,Cuenca,266,1,Villarejo-Periesteban
16,Cuenca,269,6,Villares del Saz
16,Cuenca,270,0,Villarrubio
16,Cuenca,271,7,Villarta
16,Cuenca,272,2,Villas de la Ventosa
16,Cuenca,273,8,Villaverde y Pasaconsol
16,Cuenca,274,3,Víllora
16,Cuenca,275,6,Vindel
16,Cuenca,276,9,Yémeda
16,Cuenca,277,5,Zafra de Záncara
16,Cuenca,278,1,Zafrilla
16,Cuenca,279,4,Zarza de Tajo
16,Cuenca,280,8,Zarzuela
17,Girona,001,0,Agullana
17,Girona,002,5,Aiguaviva
17,Girona,003,1,Albanyà
17,Girona,004,6,Albons
17,Girona,006,2,Alp
17,Girona,007,8,Amer
17,Girona,008,4,Anglès
17,Girona,009,7,Arbúcies
17,Girona,010,1,Argelaguer
17,Girona,011,8,"Armentera, L'"
17,Girona,012,3,Avinyonet de Puigventós
17,Girona,015,7,Banyoles
17,Girona,016,0,Bàscara
17,Girona,013,9,Begur
17,Girona,018,2,Bellcaire d'Empordà
17,Girona,019,5,Besalú
17,Girona,020,9,Bescanó
17,Girona,021,6,Beuda
17,Girona,022,1,"Bisbal d'Empordà, La"
17,Girona,234,8,Biure
17,Girona,023,7,Blanes
17,Girona,029,3,Boadella i les Escaules
17,Girona,024,2,Bolvir
17,Girona,025,5,Bordils
17,Girona,026,8,Borrassà
17,Girona,027,4,Breda
17,Girona,028,0,Brunyola
17,Girona,031,4,Cabanelles
17,Girona,030,7,Cabanes
17,Girona,032,9,Cadaqués
17,Girona,033,5,Caldes de Malavella
17,Girona,034,0,Calonge
17,Girona,035,3,Camós
17,Girona,036,6,Campdevànol
17,Girona,037,2,Campelles
17,Girona,038,8,Campllong
17,Girona,039,1,Camprodon
17,Girona,040,5,Canet d'Adri
17,Girona,041,2,Cantallops
17,Girona,042,7,Capmany
17,Girona,044,8,Cassà de la Selva
17,Girona,046,4,Castellfollit de la Roca
17,Girona,047,0,Castelló d'Empúries
17,Girona,048,6,Castell-Platja d'Aro
17,Girona,189,9,"Cellera de Ter, La"
17,Girona,049,9,Celrà
17,Girona,050,2,Cervià de Ter
17,Girona,051,9,Cistella
17,Girona,054,5,Colera
17,Girona,055,8,Colomers
17,Girona,057,7,Corçà
17,Girona,056,1,Cornellà del Terri
17,Girona,058,3,Crespià
17,Girona,901,1,"Cruïlles, Monells i Sant Sadurní de l'Heura"
17,Girona,060,0,Darnius
17,Girona,061,7,Das
17,Girona,062,2,"Escala, L'"
17,Girona,063,8,Espinelves
17,Girona,064,3,Espolla
17,Girona,065,6,Esponellà
17,Girona,005,9,"Far d'Empordà, El"
17,Girona,066,9,Figueres
17,Girona,067,5,Flaçà
17,Girona,068,1,Foixà
17,Girona,069,4,Fontanals de Cerdanya
17,Girona,070,8,Fontanilles
17,Girona,071,5,Fontcoberta
17,Girona,902,6,Forallac
17,Girona,073,6,Fornells de la Selva
17,Girona,074,1,Fortià
17,Girona,075,4,Garrigàs
17,Girona,076,7,Garrigoles
17,Girona,077,3,Garriguella
17,Girona,078,9,Ger
17,Girona,079,2,Girona
17,Girona,080,6,Gombrèn
17,Girona,081,3,Gualta
17,Girona,082,8,Guils de Cerdanya
17,Girona,083,4,Hostalric
17,Girona,084,9,Isòvol
17,Girona,085,2,Jafre
17,Girona,086,5,"Jonquera, La"
17,Girona,087,1,Juià
17,Girona,088,7,Lladó
17,Girona,089,0,Llagostera
17,Girona,090,4,Llambilles
17,Girona,091,1,Llanars
17,Girona,092,6,Llançà
17,Girona,093,2,Llers
17,Girona,094,7,Llívia
17,Girona,095,0,Lloret de Mar
17,Girona,096,3,"Llosses, Les"
17,Girona,102,4,Maçanet de Cabrenys
17,Girona,103,0,Maçanet de la Selva
17,Girona,097,9,Madremanya
17,Girona,098,5,Maià de Montcal
17,Girona,100,2,Masarac
17,Girona,101,9,Massanes
17,Girona,099,8,Meranges
17,Girona,105,8,Mieres
17,Girona,106,1,Mollet de Peralada
17,Girona,107,7,Molló
17,Girona,109,6,Montagut i Oix
17,Girona,110,0,Mont-ras
17,Girona,111,7,Navata
17,Girona,112,2,Ogassa
17,Girona,114,3,Olot
17,Girona,115,6,Ordis
17,Girona,116,9,Osor
17,Girona,117,5,Palafrugell
17,Girona,118,1,Palamós
17,Girona,119,4,Palau de Santa Eulàlia
17,Girona,121,5,Palau-sator
17,Girona,120,8,Palau-saverdera
17,Girona,123,6,Palol de Revardit
17,Girona,124,1,Pals
17,Girona,125,4,Pardines
17,Girona,126,7,Parlavà
17,Girona,128,9,Pau
17,Girona,129,2,Pedret i Marzà
17,Girona,130,6,"Pera, La"
17,Girona,132,8,Peralada
17,Girona,133,4,"Planes d'Hostoles, Les"
17,Girona,134,9,Planoles
17,Girona,135,2,Pont de Molins
17,Girona,136,5,Pontós
17,Girona,137,1,Porqueres
17,Girona,140,4,"Port de la Selva, El"
17,Girona,138,7,Portbou
17,Girona,139,0,"Preses, Les"
17,Girona,141,1,Puigcerdà
17,Girona,142,6,Quart
17,Girona,043,3,Queralbs
17,Girona,143,2,Rabós
17,Girona,144,7,Regencós
17,Girona,145,0,Ribes de Freser
17,Girona,146,3,Riells i Viabrea
17,Girona,147,9,Ripoll
17,Girona,148,5,Riudarenes
17,Girona,149,8,Riudaura
17,Girona,150,1,Riudellots de la Selva
17,Girona,151,8,Riumors
17,Girona,152,3,Roses
17,Girona,153,9,Rupià
17,Girona,154,4,Sales de Llierca
17,Girona,155,7,Salt
17,Girona,157,6,Sant Andreu Salou
17,Girona,183,3,Sant Aniol de Finestres
17,Girona,158,2,Sant Climent Sescebes
17,Girona,159,5,Sant Feliu de Buixalleu
17,Girona,160,9,Sant Feliu de Guíxols
17,Girona,161,6,Sant Feliu de Pallerols
17,Girona,162,1,Sant Ferriol
17,Girona,163,7,Sant Gregori
17,Girona,164,2,Sant Hilari Sacalm
17,Girona,165,5,Sant Jaume de Llierca
17,Girona,167,4,Sant Joan de les Abadesses
17,Girona,168,0,Sant Joan de Mollet
17,Girona,185,1,Sant Joan les Fonts
17,Girona,166,8,Sant Jordi Desvalls
17,Girona,169,3,Sant Julià de Ramis
17,Girona,903,2,Sant Julià del Llor i Bonmatí
17,Girona,171,4,Sant Llorenç de la Muga
17,Girona,172,9,Sant Martí de Llémena
17,Girona,173,5,Sant Martí Vell
17,Girona,174,0,Sant Miquel de Campmajor
17,Girona,175,3,Sant Miquel de Fluvià
17,Girona,176,6,Sant Mori
17,Girona,177,2,Sant Pau de Segúries
17,Girona,178,8,Sant Pere Pescador
17,Girona,180,5,Santa Coloma de Farners
17,Girona,181,2,Santa Cristina d'Aro
17,Girona,182,7,Santa Llogaia d'Àlguema
17,Girona,184,8,Santa Pau
17,Girona,186,4,Sarrià de Ter
17,Girona,187,0,"Saus, Camallera i Llampaies"
17,Girona,188,6,"Selva de Mar, La"
17,Girona,190,3,Serinyà
17,Girona,191,0,Serra de Daró
17,Girona,192,5,Setcases
17,Girona,193,1,Sils
17,Girona,052,4,Siurana
17,Girona,194,6,Susqueda
17,Girona,195,9,"Tallada d'Empordà, La"
17,Girona,196,2,Terrades
17,Girona,197,8,Torrent
17,Girona,198,4,Torroella de Fluvià
17,Girona,199,7,Torroella de Montgrí
17,Girona,200,1,Tortellà
17,Girona,201,8,Toses
17,Girona,202,3,Tossa de Mar
17,Girona,204,4,Ullà
17,Girona,205,7,Ullastret
17,Girona,203,9,Ultramort
17,Girona,206,0,Urús
17,Girona,014,4,"Vajol, La"
17,Girona,208,2,"Vall de Bianya, La"
17,Girona,207,6,"Vall d'en Bas, La"
17,Girona,170,7,Vallfogona de Ripollès
17,Girona,209,5,Vall-llobrega
17,Girona,210,9,Ventalló
17,Girona,211,6,Verges
17,Girona,212,1,Vidrà
17,Girona,213,7,Vidreres
17,Girona,214,2,Vilabertran
17,Girona,215,5,Vilablareix
17,Girona,217,4,Viladamat
17,Girona,216,8,Viladasens
17,Girona,218,0,Vilademuls
17,Girona,220,7,Viladrau
17,Girona,221,4,Vilafant
17,Girona,223,5,Vilajuïga
17,Girona,224,0,Vilallonga de Ter
17,Girona,225,3,Vilamacolum
17,Girona,226,6,Vilamalla
17,Girona,227,2,Vilamaniscle
17,Girona,228,8,Vilanant
17,Girona,230,5,Vila-sacra
17,Girona,222,9,Vilaür
17,Girona,233,3,Vilobí d'Onyar
17,Girona,232,7,Vilopriu
18,Granada,001,6,Agrón
18,Granada,002,1,Alamedilla
18,Granada,003,7,Albolote
18,Granada,004,2,Albondón
18,Granada,005,5,Albuñán
18,Granada,006,8,Albuñol
18,Granada,007,4,Albuñuelas
18,Granada,010,7,Aldeire
18,Granada,011,4,Alfacar
18,Granada,012,9,Algarinejo
18,Granada,013,5,Alhama de Granada
18,Granada,014,0,Alhendín
18,Granada,015,3,Alicún de Ortega
18,Granada,016,6,Almegíjar
18,Granada,017,2,Almuñécar
18,Granada,904,3,Alpujarra de la Sierra
18,Granada,018,8,Alquife
18,Granada,020,5,Arenas del Rey
18,Granada,021,2,Armilla
18,Granada,022,7,Atarfe
18,Granada,023,3,Baza
18,Granada,024,8,Beas de Granada
18,Granada,025,1,Beas de Guadix
18,Granada,027,0,Benalúa
18,Granada,028,6,Benalúa de las Villas
18,Granada,029,9,Benamaurel
18,Granada,030,3,Bérchules
18,Granada,032,5,Bubión
18,Granada,033,1,Busquístar
18,Granada,034,6,Cacín
18,Granada,035,9,Cádiar
18,Granada,036,2,Cájar
18,Granada,114,9,"Calahorra, La"
18,Granada,037,8,Calicasas
18,Granada,038,4,Campotéjar
18,Granada,039,7,Caniles
18,Granada,040,1,Cáñar
18,Granada,042,3,Capileira
18,Granada,043,9,Carataunas
18,Granada,044,4,Cástaras
18,Granada,045,7,Castilléjar
18,Granada,046,0,Castril
18,Granada,047,6,Cenes de la Vega
18,Granada,059,2,Chauchina
18,Granada,061,3,Chimeneas
18,Granada,062,8,Churriana de la Vega
18,Granada,048,2,Cijuela
18,Granada,049,5,Cogollos de Guadix
18,Granada,050,8,Cogollos de la Vega
18,Granada,051,5,Colomera
18,Granada,053,6,Cortes de Baza
18,Granada,054,1,Cortes y Graena
18,Granada,912,0,Cuevas del Campo
18,Granada,056,7,Cúllar
18,Granada,057,3,Cúllar Vega
18,Granada,063,4,Darro
18,Granada,064,9,Dehesas de Guadix
18,Granada,066,5,Deifontes
18,Granada,067,1,Diezma
18,Granada,068,7,Dílar
18,Granada,069,0,Dólar
18,Granada,070,4,Dúdar
18,Granada,071,1,Dúrcal
18,Granada,072,6,Escúzar
18,Granada,074,7,Ferreira
18,Granada,076,3,Fonelas
18,Granada,078,5,Freila
18,Granada,079,8,Fuente Vaqueros
18,Granada,905,6,"Gabias, Las"
18,Granada,082,4,Galera
18,Granada,083,0,Gobernador
18,Granada,084,5,Gójar
18,Granada,085,8,Gor
18,Granada,086,1,Gorafe
18,Granada,087,7,Granada
18,Granada,088,3,Guadahortuna
18,Granada,089,6,Guadix
18,Granada,906,9,"Guajares, Los"
18,Granada,093,8,Gualchos
18,Granada,094,3,Güejar Sierra
18,Granada,095,6,Güevéjar
18,Granada,096,9,Huélago
18,Granada,097,5,Huéneja
18,Granada,098,1,Huéscar
18,Granada,099,4,Huétor de Santillán
18,Granada,100,8,Huétor Tájar
18,Granada,101,5,Huétor Vega
18,Granada,102,0,Illora
18,Granada,103,6,Itrabo
18,Granada,105,4,Iznalloz
18,Granada,107,3,Jayena
18,Granada,108,9,Jerez del Marquesado
18,Granada,109,2,Jete
18,Granada,111,3,Jun
18,Granada,112,8,Juviles
18,Granada,115,2,Láchar
18,Granada,116,5,Lanjarón
18,Granada,117,1,Lanteira
18,Granada,119,0,Lecrín
18,Granada,120,4,Lentegí
18,Granada,121,1,Lobras
18,Granada,122,6,Loja
18,Granada,123,2,Lugros
18,Granada,124,7,Lújar
18,Granada,126,3,"Malahá, La"
18,Granada,127,9,Maracena
18,Granada,128,5,Marchal
18,Granada,132,4,Moclín
18,Granada,133,0,Molvízar
18,Granada,134,5,Monachil
18,Granada,135,8,Montefrío
18,Granada,136,1,Montejícar
18,Granada,137,7,Montillana
18,Granada,138,3,Moraleda de Zafayona
18,Granada,909,4,Morelábor
18,Granada,140,0,Motril
18,Granada,141,7,Murtas
18,Granada,903,8,Nevada
18,Granada,143,8,Nigüelas
18,Granada,144,3,Nívar
18,Granada,145,6,Ogíjares
18,Granada,146,9,Orce
18,Granada,147,5,Órgiva
18,Granada,148,1,Otívar
18,Granada,149,4,Otura
18,Granada,150,7,Padul
18,Granada,151,4,Pampaneira
18,Granada,152,9,Pedro Martínez
18,Granada,153,5,Peligros
18,Granada,154,0,"Peza, La"
18,Granada,910,8,"Pinar, El"
18,Granada,157,2,Pinos Genil
18,Granada,158,8,Pinos Puente
18,Granada,159,1,Píñar
18,Granada,161,2,Polícar
18,Granada,162,7,Polopos
18,Granada,163,3,Pórtugos
18,Granada,164,8,Puebla de Don Fadrique
18,Granada,165,1,Pulianas
18,Granada,167,0,Purullena
18,Granada,168,6,Quéntar
18,Granada,170,3,Rubite
18,Granada,171,0,Salar
18,Granada,173,1,Salobreña
18,Granada,174,6,Santa Cruz del Comercio
18,Granada,175,9,Santa Fe
18,Granada,176,2,Soportújar
18,Granada,177,8,Sorvilán
18,Granada,901,7,"Taha, La"
18,Granada,178,4,Torre-Cardela
18,Granada,179,7,Torvizcón
18,Granada,180,1,Trevélez
18,Granada,181,8,Turón
18,Granada,182,3,Ugíjar
18,Granada,902,2,"Valle, El"
18,Granada,907,5,Valle del Zalabí
18,Granada,183,9,Válor
18,Granada,911,5,Vegas del Genil
18,Granada,184,4,Vélez de Benaudalla
18,Granada,185,7,Ventas de Huelma
18,Granada,908,1,Villamena
18,Granada,187,6,Villanueva de las Torres
18,Granada,188,2,Villanueva Mesía
18,Granada,189,5,Víznar
18,Granada,192,1,Zafarraya
18,Granada,913,6,Zagra
18,Granada,193,7,"Zubia, La"
18,Granada,194,2,Zújar
19,Guadalajara,001,9,Abánades
19,Guadalajara,002,4,Ablanque
19,Guadalajara,003,0,Adobes
19,Guadalajara,004,5,Alaminos
19,Guadalajara,005,8,Alarilla
19,Guadalajara,006,1,Albalate de Zorita
19,Guadalajara,007,7,Albares
19,Guadalajara,008,3,Albendiego
19,Guadalajara,009,6,Alcocer
19,Guadalajara,010,0,Alcolea de las Peñas
19,Guadalajara,011,7,Alcolea del Pinar
19,Guadalajara,013,8,Alcoroches
19,Guadalajara,015,6,Aldeanueva de Guadalajara
19,Guadalajara,016,9,Algar de Mesa
19,Guadalajara,017,5,Algora
19,Guadalajara,018,1,Alhóndiga
19,Guadalajara,019,4,Alique
19,Guadalajara,020,8,Almadrones
19,Guadalajara,021,5,Almoguera
19,Guadalajara,022,0,Almonacid de Zorita
19,Guadalajara,023,6,Alocén
19,Guadalajara,024,1,Alovera
19,Guadalajara,027,3,Alustante
19,Guadalajara,031,3,Angón
19,Guadalajara,032,8,Anguita
19,Guadalajara,033,4,Anquela del Ducado
19,Guadalajara,034,9,Anquela del Pedregal
19,Guadalajara,036,5,Aranzueque
19,Guadalajara,037,1,Arbancón
19,Guadalajara,038,7,Arbeteta
19,Guadalajara,039,0,Argecilla
19,Guadalajara,040,4,Armallones
19,Guadalajara,041,1,Armuña de Tajuña
19,Guadalajara,042,6,Arroyo de las Fraguas
19,Guadalajara,043,2,Atanzón
19,Guadalajara,044,7,Atienza
19,Guadalajara,045,0,Auñón
19,Guadalajara,046,3,Azuqueca de Henares
19,Guadalajara,047,9,Baides
19,Guadalajara,048,5,Baños de Tajo
19,Guadalajara,049,8,Bañuelos
19,Guadalajara,050,1,Barriopedro
19,Guadalajara,051,8,Berninches
19,Guadalajara,052,3,"Bodera, La"
19,Guadalajara,053,9,Brihuega
19,Guadalajara,054,4,Budia
19,Guadalajara,055,7,Bujalaro
19,Guadalajara,057,6,Bustares
19,Guadalajara,058,2,Cabanillas del Campo
19,Guadalajara,059,5,Campillo de Dueñas
19,Guadalajara,060,9,Campillo de Ranas
19,Guadalajara,061,6,Campisábalos
19,Guadalajara,064,2,Canredondo
19,Guadalajara,065,5,Cantalojas
19,Guadalajara,066,8,Cañizar
19,Guadalajara,067,4,"Cardoso de la Sierra, El"
19,Guadalajara,070,7,Casa de Uceda
19,Guadalajara,071,4,"Casar, El"
19,Guadalajara,073,5,Casas de San Galindo
19,Guadalajara,074,0,Caspueñas
19,Guadalajara,075,3,Castejón de Henares
19,Guadalajara,076,6,Castellar de la Muela
19,Guadalajara,078,8,Castilforte
19,Guadalajara,079,1,Castilnuevo
19,Guadalajara,080,5,Cendejas de Enmedio
19,Guadalajara,081,2,Cendejas de la Torre
19,Guadalajara,082,7,Centenera
19,Guadalajara,103,9,Checa
19,Guadalajara,104,4,Chequilla
19,Guadalajara,106,0,Chillarón del Rey
19,Guadalajara,105,7,Chiloeches
19,Guadalajara,086,4,Cifuentes
19,Guadalajara,087,0,Cincovillas
19,Guadalajara,088,6,Ciruelas
19,Guadalajara,089,9,Ciruelos del Pinar
19,Guadalajara,090,3,Cobeta
19,Guadalajara,091,0,Cogollor
19,Guadalajara,092,5,Cogolludo
19,Guadalajara,095,9,Condemios de Abajo
19,Guadalajara,096,2,Condemios de Arriba
19,Guadalajara,097,8,Congostrina
19,Guadalajara,098,4,Copernal
19,Guadalajara,099,7,Corduente
19,Guadalajara,102,3,"Cubillo de Uceda, El"
19,Guadalajara,107,6,Driebes
19,Guadalajara,108,2,Durón
19,Guadalajara,109,5,Embid
19,Guadalajara,110,9,Escamilla
19,Guadalajara,111,6,Escariche
19,Guadalajara,112,1,Escopete
19,Guadalajara,113,7,Espinosa de Henares
19,Guadalajara,114,2,Esplegares
19,Guadalajara,115,5,Establés
19,Guadalajara,116,8,Estriégana
19,Guadalajara,117,4,Fontanar
19,Guadalajara,118,0,Fuembellida
19,Guadalajara,119,3,Fuencemillán
19,Guadalajara,120,7,Fuentelahiguera de Albatages
19,Guadalajara,121,4,Fuentelencina
19,Guadalajara,122,9,Fuentelsaz
19,Guadalajara,123,5,Fuentelviejo
19,Guadalajara,124,0,Fuentenovilla
19,Guadalajara,125,3,Gajanejos
19,Guadalajara,126,6,Galápagos
19,Guadalajara,127,2,Galve de Sorbe
19,Guadalajara,129,1,Gascueña de Bornova
19,Guadalajara,130,5,Guadalajara
19,Guadalajara,132,7,Henche
19,Guadalajara,133,3,Heras de Ayuso
19,Guadalajara,134,8,Herrería
19,Guadalajara,135,1,Hiendelaencina
19,Guadalajara,136,4,Hijes
19,Guadalajara,138,6,Hita
19,Guadalajara,139,9,Hombrados
19,Guadalajara,142,5,Hontoba
19,Guadalajara,143,1,Horche
19,Guadalajara,145,9,Hortezuela de Océn
19,Guadalajara,146,2,"Huerce, La"
19,Guadalajara,147,8,Huérmeces del Cerro
19,Guadalajara,148,4,Huertahernando
19,Guadalajara,150,0,Hueva
19,Guadalajara,151,7,Humanes
19,Guadalajara,152,2,Illana
19,Guadalajara,153,8,Iniéstola
19,Guadalajara,154,3,"Inviernas, Las"
19,Guadalajara,155,6,Irueste
19,Guadalajara,156,9,Jadraque
19,Guadalajara,157,5,Jirueque
19,Guadalajara,159,4,Ledanca
19,Guadalajara,160,8,Loranca de Tajuña
19,Guadalajara,161,5,Lupiana
19,Guadalajara,162,0,Luzaga
19,Guadalajara,163,6,Luzón
19,Guadalajara,165,4,Majaelrayo
19,Guadalajara,166,7,Málaga del Fresno
19,Guadalajara,167,3,Malaguilla
19,Guadalajara,168,9,Mandayona
19,Guadalajara,169,2,Mantiel
19,Guadalajara,170,6,Maranchón
19,Guadalajara,171,3,Marchamalo
19,Guadalajara,172,8,Masegoso de Tajuña
19,Guadalajara,173,4,Matarrubia
19,Guadalajara,174,9,Matillas
19,Guadalajara,175,2,Mazarete
19,Guadalajara,176,5,Mazuecos
19,Guadalajara,177,1,Medranda
19,Guadalajara,178,7,Megina
19,Guadalajara,179,0,Membrillera
19,Guadalajara,181,1,Miedes de Atienza
19,Guadalajara,182,6,"Mierla, La"
19,Guadalajara,184,7,Millana
19,Guadalajara,183,2,Milmarcos
19,Guadalajara,185,0,"Miñosa, La"
19,Guadalajara,186,3,Mirabueno
19,Guadalajara,187,9,Miralrío
19,Guadalajara,188,5,Mochales
19,Guadalajara,189,8,Mohernando
19,Guadalajara,190,2,Molina de Aragón
19,Guadalajara,191,9,Monasterio
19,Guadalajara,192,4,Mondéjar
19,Guadalajara,193,0,Montarrón
19,Guadalajara,194,5,Moratilla de los Meleros
19,Guadalajara,195,8,Morenilla
19,Guadalajara,196,1,Muduex
19,Guadalajara,197,7,"Navas de Jadraque, Las"
19,Guadalajara,198,3,Negredo
19,Guadalajara,199,6,Ocentejo
19,Guadalajara,200,0,"Olivar, El"
19,Guadalajara,201,7,Olmeda de Cobeta
19,Guadalajara,202,2,"Olmeda de Jadraque, La"
19,Guadalajara,203,8,"Ordial, El"
19,Guadalajara,204,3,Orea
19,Guadalajara,208,1,Pálmaces de Jadraque
19,Guadalajara,209,4,Pardos
19,Guadalajara,210,8,Paredes de Sigüenza
19,Guadalajara,211,5,Pareja
19,Guadalajara,212,0,Pastrana
19,Guadalajara,213,6,"Pedregal, El"
19,Guadalajara,214,1,Peñalén
19,Guadalajara,215,4,Peñalver
19,Guadalajara,216,7,Peralejos de las Truchas
19,Guadalajara,217,3,Peralveche
19,Guadalajara,218,9,Pinilla de Jadraque
19,Guadalajara,219,2,Pinilla de Molina
19,Guadalajara,220,6,Pioz
19,Guadalajara,221,3,Piqueras
19,Guadalajara,222,8,"Pobo de Dueñas, El"
19,Guadalajara,223,4,Poveda de la Sierra
19,Guadalajara,224,9,Pozo de Almoguera
19,Guadalajara,225,2,Pozo de Guadalajara
19,Guadalajara,226,5,Prádena de Atienza
19,Guadalajara,227,1,Prados Redondos
19,Guadalajara,228,7,Puebla de Beleña
19,Guadalajara,229,0,Puebla de Valles
19,Guadalajara,230,4,Quer
19,Guadalajara,231,1,Rebollosa de Jadraque
19,Guadalajara,232,6,"Recuenco, El"
19,Guadalajara,233,2,Renera
19,Guadalajara,234,7,Retiendas
19,Guadalajara,235,0,Riba de Saelices
19,Guadalajara,237,9,Rillo de Gallo
19,Guadalajara,238,5,Riofrío del Llano
19,Guadalajara,239,8,Robledillo de Mohernando
19,Guadalajara,240,2,Robledo de Corpes
19,Guadalajara,241,9,Romanillos de Atienza
19,Guadalajara,242,4,Romanones
19,Guadalajara,243,0,Rueda de la Sierra
19,Guadalajara,244,5,Sacecorbo
19,Guadalajara,245,8,Sacedón
19,Guadalajara,246,1,Saelices de la Sal
19,Guadalajara,247,7,Salmerón
19,Guadalajara,248,3,San Andrés del Congosto
19,Guadalajara,249,6,San Andrés del Rey
19,Guadalajara,250,9,Santiuste
19,Guadalajara,251,6,Saúca
19,Guadalajara,252,1,Sayatón
19,Guadalajara,254,2,Selas
19,Guadalajara,901,0,Semillas
19,Guadalajara,255,5,Setiles
19,Guadalajara,256,8,Sienes
19,Guadalajara,257,4,Sigüenza
19,Guadalajara,258,0,Solanillos del Extremo
19,Guadalajara,259,3,Somolinos
19,Guadalajara,260,7,"Sotillo, El"
19,Guadalajara,261,4,Sotodosos
19,Guadalajara,262,9,Tamajón
19,Guadalajara,263,5,Taragudo
19,Guadalajara,264,0,Taravilla
19,Guadalajara,265,3,Tartanedo
19,Guadalajara,266,6,Tendilla
19,Guadalajara,267,2,Terzaga
19,Guadalajara,268,8,Tierzo
19,Guadalajara,269,1,"Toba, La"
19,Guadalajara,271,2,Tordellego
19,Guadalajara,270,5,Tordelrábano
19,Guadalajara,272,7,Tordesilos
19,Guadalajara,274,8,Torija
19,Guadalajara,279,9,Torre del Burgo
19,Guadalajara,277,0,Torrecuadrada de Molina
19,Guadalajara,278,6,Torrecuadradilla
19,Guadalajara,280,3,Torrejón del Rey
19,Guadalajara,281,0,Torremocha de Jadraque
19,Guadalajara,282,5,Torremocha del Campo
19,Guadalajara,283,1,Torremocha del Pinar
19,Guadalajara,284,6,Torremochuela
19,Guadalajara,285,9,Torrubia
19,Guadalajara,286,2,Tórtola de Henares
19,Guadalajara,287,8,Tortuera
19,Guadalajara,288,4,Tortuero
19,Guadalajara,289,7,Traíd
19,Guadalajara,290,1,Trijueque
19,Guadalajara,291,8,Trillo
19,Guadalajara,293,9,Uceda
19,Guadalajara,294,4,Ujados
19,Guadalajara,296,0,Utande
19,Guadalajara,297,6,Valdarachas
19,Guadalajara,298,2,Valdearenas
19,Guadalajara,299,5,Valdeavellano
19,Guadalajara,300,9,Valdeaveruelo
19,Guadalajara,301,6,Valdeconcha
19,Guadalajara,302,1,Valdegrudas
19,Guadalajara,303,7,Valdelcubo
19,Guadalajara,304,2,Valdenuño Fernández
19,Guadalajara,305,5,Valdepeñas de la Sierra
19,Guadalajara,306,8,Valderrebollo
19,Guadalajara,307,4,Valdesotos
19,Guadalajara,308,0,Valfermoso de Tajuña
19,Guadalajara,309,3,Valhermoso
19,Guadalajara,310,7,Valtablado del Río
19,Guadalajara,311,4,Valverde de los Arroyos
19,Guadalajara,314,0,Viana de Jadraque
19,Guadalajara,317,2,Villanueva de Alcorón
19,Guadalajara,318,8,Villanueva de Argecilla
19,Guadalajara,319,1,Villanueva de la Torre
19,Guadalajara,321,2,Villares de Jadraque
19,Guadalajara,322,7,Villaseca de Henares
19,Guadalajara,323,3,Villaseca de Uceda
19,Guadalajara,324,8,Villel de Mesa
19,Guadalajara,325,1,Viñuelas
19,Guadalajara,326,4,Yebes
19,Guadalajara,327,0,Yebra
19,Guadalajara,329,9,Yélamos de Abajo
19,Guadalajara,330,3,Yélamos de Arriba
19,Guadalajara,331,0,Yunquera de Henares
19,Guadalajara,332,5,"Yunta, La"
19,Guadalajara,333,1,Zaorejas
19,Guadalajara,334,6,Zarzuela de Jadraque
19,Guadalajara,335,9,Zorita de los Canes
20,Gipuzkoa,000,0,San Sebastian
20,Gipuzkoa,001,3,Abaltzisketa
20,Gipuzkoa,002,8,Aduna
20,Gipuzkoa,016,3,Aia
20,Gipuzkoa,003,4,Aizarnazabal
20,Gipuzkoa,004,9,Albiztur
20,Gipuzkoa,005,2,Alegia
20,Gipuzkoa,006,5,Alkiza
20,Gipuzkoa,906,6,Altzaga
20,Gipuzkoa,007,1,Altzo
20,Gipuzkoa,008,7,Amezketa
20,Gipuzkoa,009,0,Andoain
20,Gipuzkoa,010,4,Anoeta
20,Gipuzkoa,011,1,Antzuola
20,Gipuzkoa,012,6,Arama
20,Gipuzkoa,013,2,Aretxabaleta
20,Gipuzkoa,055,1,Arrasate/Mondragón
20,Gipuzkoa,014,7,Asteasu
20,Gipuzkoa,903,5,Astigarraga
20,Gipuzkoa,015,0,Ataun
20,Gipuzkoa,017,9,Azkoitia
20,Gipuzkoa,018,5,Azpeitia
20,Gipuzkoa,904,0,Baliarrain
20,Gipuzkoa,019,8,Beasain
20,Gipuzkoa,020,2,Beizama
20,Gipuzkoa,021,9,Belauntza
20,Gipuzkoa,022,4,Berastegi
20,Gipuzkoa,074,4,Bergara
20,Gipuzkoa,023,0,Berrobi
20,Gipuzkoa,024,5,Bidegoian
20,Gipuzkoa,029,6,Deba
20,Gipuzkoa,069,7,Donostia-San Sebastián
20,Gipuzkoa,030,0,Eibar
20,Gipuzkoa,031,7,Elduain
20,Gipuzkoa,033,8,Elgeta
20,Gipuzkoa,032,2,Elgoibar
20,Gipuzkoa,067,8,Errenteria
20,Gipuzkoa,066,2,Errezil
20,Gipuzkoa,034,3,Eskoriatza
20,Gipuzkoa,035,6,Ezkio-Itsaso
20,Gipuzkoa,038,1,Gabiria
20,Gipuzkoa,037,5,Gaintza
20,Gipuzkoa,907,2,Gaztelu
20,Gipuzkoa,039,4,Getaria
20,Gipuzkoa,040,8,Hernani
20,Gipuzkoa,041,5,Hernialde
20,Gipuzkoa,036,9,Hondarribia
20,Gipuzkoa,042,0,Ibarra
20,Gipuzkoa,043,6,Idiazabal
20,Gipuzkoa,044,1,Ikaztegieta
20,Gipuzkoa,045,4,Irun
20,Gipuzkoa,046,7,Irura
20,Gipuzkoa,047,3,Itsasondo
20,Gipuzkoa,048,9,Larraul
20,Gipuzkoa,902,9,Lasarte-Oria
20,Gipuzkoa,049,2,Lazkao
20,Gipuzkoa,050,5,Leaburu
20,Gipuzkoa,051,2,Legazpi
20,Gipuzkoa,052,7,Legorreta
20,Gipuzkoa,068,4,Leintz-Gatzaga
20,Gipuzkoa,053,3,Lezo
20,Gipuzkoa,054,8,Lizartza
20,Gipuzkoa,901,4,Mendaro
20,Gipuzkoa,057,0,Mutiloa
20,Gipuzkoa,056,4,Mutriku
20,Gipuzkoa,063,1,Oiartzun
20,Gipuzkoa,058,6,Olaberria
20,Gipuzkoa,059,9,Oñati
20,Gipuzkoa,076,0,Ordizia
20,Gipuzkoa,905,3,Orendain
20,Gipuzkoa,060,3,Orexa
20,Gipuzkoa,061,0,Orio
20,Gipuzkoa,062,5,Ormaiztegi
20,Gipuzkoa,064,6,Pasaia
20,Gipuzkoa,070,1,Segura
20,Gipuzkoa,065,9,Soraluze/Placencia de las Armas
20,Gipuzkoa,071,8,Tolosa
20,Gipuzkoa,072,3,Urnieta
20,Gipuzkoa,077,6,Urretxu
20,Gipuzkoa,073,9,Usurbil
20,Gipuzkoa,075,7,Villabona
20,Gipuzkoa,078,2,Zaldibia
20,Gipuzkoa,079,5,Zarautz
20,Gipuzkoa,025,8,Zegama
20,Gipuzkoa,026,1,Zerain
20,Gipuzkoa,027,7,Zestoa
20,Gipuzkoa,028,3,Zizurkil
20,Gipuzkoa,081,6,Zumaia
20,Gipuzkoa,080,9,Zumarraga
21,Huelva,001,0,Alájar
21,Huelva,002,5,Aljaraque
21,Huelva,003,1,"Almendro, El"
21,Huelva,004,6,Almonaster la Real
21,Huelva,005,9,Almonte
21,Huelva,006,2,Alosno
21,Huelva,007,8,Aracena
21,Huelva,008,4,Aroche
21,Huelva,009,7,Arroyomolinos de León
21,Huelva,010,1,Ayamonte
21,Huelva,011,8,Beas
21,Huelva,012,3,Berrocal
21,Huelva,013,9,Bollullos Par del Condado
21,Huelva,014,4,Bonares
21,Huelva,015,7,Cabezas Rubias
21,Huelva,016,0,Cala
21,Huelva,017,6,Calañas
21,Huelva,018,2,"Campillo, El"
21,Huelva,019,5,Campofrío
21,Huelva,020,9,Cañaveral de León
21,Huelva,021,6,Cartaya
21,Huelva,022,1,Castaño del Robledo
21,Huelva,023,7,"Cerro de Andévalo, El"
21,Huelva,030,7,Chucena
21,Huelva,024,2,Corteconcepción
21,Huelva,025,5,Cortegana
21,Huelva,026,8,Cortelazor
21,Huelva,027,4,Cumbres de Enmedio
21,Huelva,028,0,Cumbres de San Bartolomé
21,Huelva,029,3,Cumbres Mayores
21,Huelva,031,4,Encinasola
21,Huelva,032,9,Escacena del Campo
21,Huelva,033,5,Fuenteheridos
21,Huelva,034,0,Galaroza
21,Huelva,035,3,Gibraleón
21,Huelva,036,6,"Granada de Río-Tinto, La"
21,Huelva,037,2,"Granado, El"
21,Huelva,038,8,Higuera de la Sierra
21,Huelva,039,1,Hinojales
21,Huelva,040,5,Hinojos
21,Huelva,041,2,Huelva
21,Huelva,042,7,Isla Cristina
21,Huelva,043,3,Jabugo
21,Huelva,044,8,Lepe
21,Huelva,045,1,Linares de la Sierra
21,Huelva,046,4,Lucena del Puerto
21,Huelva,047,0,Manzanilla
21,Huelva,048,6,"Marines, Los"
21,Huelva,049,9,Minas de Riotinto
21,Huelva,050,2,Moguer
21,Huelva,051,9,"Nava, La"
21,Huelva,052,4,Nerva
21,Huelva,053,0,Niebla
21,Huelva,054,5,"Palma del Condado, La"
21,Huelva,055,8,Palos de la Frontera
21,Huelva,056,1,Paterna del Campo
21,Huelva,057,7,Paymogo
21,Huelva,058,3,Puebla de Guzmán
21,Huelva,059,6,Puerto Moral
21,Huelva,060,0,Punta Umbría
21,Huelva,061,7,Rociana del Condado
21,Huelva,062,2,Rosal de la Frontera
21,Huelva,063,8,San Bartolomé de la Torre
21,Huelva,064,3,San Juan del Puerto
21,Huelva,066,9,San Silvestre de Guzmán
21,Huelva,065,6,Sanlúcar de Guadiana
21,Huelva,067,5,Santa Ana la Real
21,Huelva,068,1,Santa Bárbara de Casa
21,Huelva,069,4,Santa Olalla del Cala
21,Huelva,070,8,Trigueros
21,Huelva,071,5,Valdelarco
21,Huelva,072,0,Valverde del Camino
21,Huelva,073,6,Villablanca
21,Huelva,074,1,Villalba del Alcor
21,Huelva,075,4,Villanueva de las Cruces
21,Huelva,076,7,Villanueva de los Castillejos
21,Huelva,077,3,Villarrasa
21,Huelva,078,9,Zalamea la Real
21,Huelva,079,2,Zufre
22,Huesca,001,5,Abiego
22,Huesca,002,0,Abizanda
22,Huesca,003,6,Adahuesca
22,Huesca,004,1,Agüero
22,Huesca,907,4,Aínsa-Sobrarbe
22,Huesca,006,7,Aisa
22,Huesca,007,3,Albalate de Cinca
22,Huesca,008,9,Albalatillo
22,Huesca,009,2,Albelda
22,Huesca,011,3,Albero Alto
22,Huesca,012,8,Albero Bajo
22,Huesca,013,4,Alberuela de Tubo
22,Huesca,014,9,Alcalá de Gurrea
22,Huesca,015,2,Alcalá del Obispo
22,Huesca,016,5,Alcampell
22,Huesca,017,1,Alcolea de Cinca
22,Huesca,018,7,Alcubierre
22,Huesca,019,0,Alerre
22,Huesca,020,4,Alfántega
22,Huesca,021,1,Almudévar
22,Huesca,022,6,Almunia de San Juan
22,Huesca,023,2,Almuniente
22,Huesca,024,7,Alquézar
22,Huesca,025,0,Altorricón
22,Huesca,027,9,Angüés
22,Huesca,028,5,Ansó
22,Huesca,029,8,Antillón
22,Huesca,032,4,Aragüés del Puerto
22,Huesca,035,8,Arén
22,Huesca,036,1,Argavieso
22,Huesca,037,7,Arguis
22,Huesca,039,6,Ayerbe
22,Huesca,040,0,Azanuy-Alins
22,Huesca,041,7,Azara
22,Huesca,042,2,Azlor
22,Huesca,043,8,Baélls
22,Huesca,044,3,Bailo
22,Huesca,045,6,Baldellou
22,Huesca,046,9,Ballobar
22,Huesca,047,5,Banastás
22,Huesca,048,1,Barbastro
22,Huesca,049,4,Barbués
22,Huesca,050,7,Barbuñales
22,Huesca,051,4,Bárcabo
22,Huesca,052,9,Belver de Cinca
22,Huesca,053,5,Benabarre
22,Huesca,054,0,Benasque
22,Huesca,055,3,Berbegal
22,Huesca,057,2,Bielsa
22,Huesca,058,8,Bierge
22,Huesca,059,1,Biescas
22,Huesca,060,5,Binaced
22,Huesca,061,2,Binéfar
22,Huesca,062,7,Bisaurri
22,Huesca,063,3,Biscarrués
22,Huesca,064,8,Blecua y Torres
22,Huesca,066,4,Boltaña
22,Huesca,067,0,Bonansa
22,Huesca,068,6,Borau
22,Huesca,069,9,Broto
22,Huesca,072,5,Caldearenas
22,Huesca,074,6,Campo
22,Huesca,075,9,Camporrélls
22,Huesca,076,2,Canal de Berdún
22,Huesca,077,8,Candasnos
22,Huesca,078,4,Canfranc
22,Huesca,079,7,Capdesaso
22,Huesca,080,1,Capella
22,Huesca,081,8,Casbas de Huesca
22,Huesca,083,9,Castejón de Monegros
22,Huesca,084,4,Castejón de Sos
22,Huesca,082,3,Castejón del Puente
22,Huesca,085,7,Castelflorite
22,Huesca,086,0,Castiello de Jaca
22,Huesca,087,6,Castigaleu
22,Huesca,088,2,Castillazuelo
22,Huesca,089,5,Castillonroy
22,Huesca,094,2,Chalamera
22,Huesca,095,5,Chía
22,Huesca,096,8,Chimillas
22,Huesca,090,9,Colungo
22,Huesca,099,3,Esplús
22,Huesca,102,9,Estada
22,Huesca,103,5,Estadilla
22,Huesca,105,3,Estopiñán del Castillo
22,Huesca,106,6,Fago
22,Huesca,107,2,Fanlo
22,Huesca,109,1,Fiscal
22,Huesca,110,5,Fonz
22,Huesca,111,2,Foradada del Toscar
22,Huesca,112,7,Fraga
22,Huesca,113,3,"Fueva, La"
22,Huesca,114,8,Gistaín
22,Huesca,115,1,"Grado, El"
22,Huesca,116,4,Grañén
22,Huesca,117,0,Graus
22,Huesca,119,9,Gurrea de Gállego
22,Huesca,122,5,Hoz de Jaca
22,Huesca,908,0,Hoz y Costean
22,Huesca,124,6,Huerto
22,Huesca,125,9,Huesca
22,Huesca,126,2,Ibieca
22,Huesca,127,8,Igriés
22,Huesca,128,4,Ilche
22,Huesca,129,7,Isábena
22,Huesca,130,1,Jaca
22,Huesca,131,8,Jasa
22,Huesca,133,9,Labuerda
22,Huesca,135,7,Laluenga
22,Huesca,136,0,Lalueza
22,Huesca,137,6,Lanaja
22,Huesca,139,5,Laperdiguera
22,Huesca,141,6,Lascellas-Ponzano
22,Huesca,142,1,Lascuarre
22,Huesca,143,7,Laspaúles
22,Huesca,144,2,Laspuña
22,Huesca,149,3,Loarre
22,Huesca,150,6,Loporzano
22,Huesca,151,3,Loscorrales
22,Huesca,905,5,Lupiñén-Ortilla
22,Huesca,155,2,Monesma y Cajigar
22,Huesca,156,5,Monflorite-Lascasas
22,Huesca,157,1,Montanuy
22,Huesca,158,7,Monzón
22,Huesca,160,4,Naval
22,Huesca,162,6,Novales
22,Huesca,163,2,Nueno
22,Huesca,164,7,Olvena
22,Huesca,165,0,Ontiñena
22,Huesca,167,9,Osso de Cinca
22,Huesca,168,5,Palo
22,Huesca,170,2,Panticosa
22,Huesca,172,4,Peñalba
22,Huesca,173,0,"Peñas de Riglos, Las"
22,Huesca,174,5,Peralta de Alcofea
22,Huesca,175,8,Peralta de Calasanz
22,Huesca,176,1,Peraltilla
22,Huesca,177,7,Perarrúa
22,Huesca,178,3,Pertusa
22,Huesca,181,7,Piracés
22,Huesca,182,2,Plan
22,Huesca,184,3,Poleñino
22,Huesca,186,9,Pozán de Vero
22,Huesca,187,5,"Puebla de Castro, La"
22,Huesca,188,1,Puente de Montañana
22,Huesca,902,1,Puente la Reina de Jaca
22,Huesca,189,4,Puértolas
22,Huesca,190,8,"Pueyo de Araguás, El"
22,Huesca,193,6,Pueyo de Santa Cruz
22,Huesca,195,4,Quicena
22,Huesca,197,3,Robres
22,Huesca,199,2,Sabiñánigo
22,Huesca,200,6,Sahún
22,Huesca,201,3,Salas Altas
22,Huesca,202,8,Salas Bajas
22,Huesca,203,4,Salillas
22,Huesca,204,9,Sallent de Gállego
22,Huesca,205,2,San Esteban de Litera
22,Huesca,207,1,San Juan de Plan
22,Huesca,903,7,San Miguel del Cinca
22,Huesca,206,5,Sangarrén
22,Huesca,208,7,Santa Cilia
22,Huesca,209,0,Santa Cruz de la Serós
22,Huesca,906,8,Santa María de Dulcis
22,Huesca,212,6,Santaliestra y San Quílez
22,Huesca,213,2,Sariñena
22,Huesca,214,7,Secastilla
22,Huesca,215,0,Seira
22,Huesca,217,9,Sena
22,Huesca,218,5,Senés de Alcubierre
22,Huesca,220,2,Sesa
22,Huesca,221,9,Sesué
22,Huesca,222,4,Siétamo
22,Huesca,223,0,Sopeira
22,Huesca,904,2,"Sotonera, La"
22,Huesca,225,8,Tamarite de Litera
22,Huesca,226,1,Tardienta
22,Huesca,227,7,Tella-Sin
22,Huesca,228,3,Tierz
22,Huesca,229,6,Tolva
22,Huesca,230,0,Torla
22,Huesca,232,2,Torralba de Aragón
22,Huesca,233,8,Torre la Ribera
22,Huesca,234,3,Torrente de Cinca
22,Huesca,235,6,Torres de Alcanadre
22,Huesca,236,9,Torres de Barbués
22,Huesca,239,4,Tramaced
22,Huesca,242,0,Valfarta
22,Huesca,243,6,Valle de Bardají
22,Huesca,901,6,Valle de Hecho
22,Huesca,244,1,Valle de Lierp
22,Huesca,245,4,Velilla de Cinca
22,Huesca,909,3,Vencillón
22,Huesca,246,7,Veracruz
22,Huesca,247,3,Viacamp y Litera
22,Huesca,248,9,Vicién
22,Huesca,249,2,Villanova
22,Huesca,250,5,Villanúa
22,Huesca,251,2,Villanueva de Sigena
22,Huesca,252,7,Yebra de Basa
22,Huesca,253,3,Yésero
22,Huesca,254,8,Zaidín
23,Jaén,000,0,Jaen
23,Jaén,001,1,Albanchez de Mágina
23,Jaén,002,6,Alcalá la Real
23,Jaén,003,2,Alcaudete
23,Jaén,004,7,Aldeaquemada
23,Jaén,005,0,Andújar
23,Jaén,006,3,Arjona
23,Jaén,007,9,Arjonilla
23,Jaén,008,5,Arquillos
23,Jaén,905,1,Arroyo del Ojanco
23,Jaén,009,8,Baeza
23,Jaén,010,2,Bailén
23,Jaén,011,9,Baños de la Encina
23,Jaén,012,4,Beas de Segura
23,Jaén,902,7,Bedmar y Garcíez
23,Jaén,014,5,Begíjar
23,Jaén,015,8,Bélmez de la Moraleda
23,Jaén,016,1,Benatae
23,Jaén,017,7,Cabra del Santo Cristo
23,Jaén,018,3,Cambil
23,Jaén,019,6,Campillo de Arenas
23,Jaén,020,0,Canena
23,Jaén,021,7,Carboneros
23,Jaén,901,2,Cárcheles
23,Jaén,024,3,"Carolina, La"
23,Jaén,025,6,Castellar
23,Jaén,026,9,Castillo de Locubín
23,Jaén,027,5,Cazalilla
23,Jaén,028,1,Cazorla
23,Jaén,029,4,Chiclana de Segura
23,Jaén,030,8,Chilluévar
23,Jaén,031,5,Escañuela
23,Jaén,032,0,Espelúy
23,Jaén,033,6,Frailes
23,Jaén,034,1,Fuensanta de Martos
23,Jaén,035,4,Fuerte del Rey
23,Jaén,037,3,Génave
23,Jaén,038,9,"Guardia de Jaén, La"
23,Jaén,039,2,Guarromán
23,Jaén,041,3,Higuera de Calatrava
23,Jaén,042,8,Hinojares
23,Jaén,043,4,Hornos
23,Jaén,044,9,Huelma
23,Jaén,045,2,Huesa
23,Jaén,046,5,Ibros
23,Jaén,047,1,"Iruela, La"
23,Jaén,048,7,Iznatoraf
23,Jaén,049,0,Jabalquinto
23,Jaén,050,3,Jaén
23,Jaén,051,0,Jamilena
23,Jaén,052,5,Jimena
23,Jaén,053,1,Jódar
23,Jaén,040,6,Lahiguera
23,Jaén,054,6,Larva
23,Jaén,055,9,Linares
23,Jaén,056,2,Lopera
23,Jaén,057,8,Lupión
23,Jaén,058,4,Mancha Real
23,Jaén,059,7,Marmolejo
23,Jaén,060,1,Martos
23,Jaén,061,8,Mengíbar
23,Jaén,062,3,Montizón
23,Jaén,063,9,Navas de San Juan
23,Jaén,064,4,Noalejo
23,Jaén,065,7,Orcera
23,Jaén,066,0,Peal de Becerro
23,Jaén,067,6,Pegalajar
23,Jaén,069,5,Porcuna
23,Jaén,070,9,Pozo Alcón
23,Jaén,071,6,Puente de Génave
23,Jaén,072,1,"Puerta de Segura, La"
23,Jaén,073,7,Quesada
23,Jaén,074,2,Rus
23,Jaén,075,5,Sabiote
23,Jaén,076,8,Santa Elena
23,Jaén,077,4,Santiago de Calatrava
23,Jaén,904,8,Santiago-Pontones
23,Jaén,079,3,Santisteban del Puerto
23,Jaén,080,7,Santo Tomé
23,Jaén,081,4,Segura de la Sierra
23,Jaén,082,9,Siles
23,Jaén,084,0,Sorihuela del Guadalimar
23,Jaén,086,6,Torre del Campo
23,Jaén,085,3,Torreblascopedro
23,Jaén,087,2,Torredonjimeno
23,Jaén,088,8,Torreperogil
23,Jaén,090,5,Torres
23,Jaén,091,2,Torres de Albánchez
23,Jaén,092,7,Úbeda
23,Jaén,093,3,Valdepeñas de Jaén
23,Jaén,094,8,Vilches
23,Jaén,095,1,Villacarrillo
23,Jaén,096,4,Villanueva de la Reina
23,Jaén,097,0,Villanueva del Arzobispo
23,Jaén,098,6,Villardompardo
23,Jaén,099,9,"Villares, Los"
23,Jaén,101,0,Villarrodrigo
23,Jaén,903,3,Villatorres
24,León,000,0,Leon
24,León,001,6,Acebedo
24,León,002,1,Algadefe
24,León,003,7,Alija del Infantado
24,León,004,2,Almanza
24,León,005,5,"Antigua, La"
24,León,006,8,Ardón
24,León,007,4,Arganza
24,León,008,0,Astorga
24,León,009,3,Balboa
24,León,010,7,"Bañeza, La"
24,León,011,4,Barjas
24,León,012,9,"Barrios de Luna, Los"
24,León,014,0,Bembibre
24,León,015,3,Benavides
24,León,016,6,Benuza
24,León,017,2,Bercianos del Páramo
24,León,018,8,Bercianos del Real Camino
24,León,019,1,Berlanga del Bierzo
24,León,020,5,Boca de Huérgano
24,León,021,2,Boñar
24,León,022,7,Borrenes
24,León,023,3,Brazuelo
24,León,024,8,"Burgo Ranero, El"
24,León,025,1,Burón
24,León,026,4,Bustillo del Páramo
24,León,027,0,Cabañas Raras
24,León,028,6,Cabreros del Río
24,León,029,9,Cabrillanes
24,León,030,3,Cacabelos
24,León,031,0,Calzada del Coto
24,León,032,5,Campazas
24,León,033,1,Campo de Villavidel
24,León,034,6,Camponaraya
24,León,036,2,Candín
24,León,037,8,Cármenes
24,León,038,4,Carracedelo
24,León,039,7,Carrizo
24,León,040,1,Carrocera
24,León,041,8,Carucedo
24,León,042,3,Castilfalé
24,León,043,9,Castrillo de Cabrera
24,León,044,4,Castrillo de la Valduerna
24,León,046,0,Castrocalbón
24,León,047,6,Castrocontrigo
24,León,049,5,Castropodame
24,León,050,8,Castrotierra de Valmadrigal
24,León,051,5,Cea
24,León,052,0,Cebanico
24,León,053,6,Cebrones del Río
24,León,065,2,Chozas de Abajo
24,León,054,1,Cimanes de la Vega
24,León,055,4,Cimanes del Tejar
24,León,056,7,Cistierna
24,León,057,3,Congosto
24,León,058,9,Corbillos de los Oteros
24,León,059,2,Corullón
24,León,060,6,Crémenes
24,León,061,3,Cuadros
24,León,062,8,Cubillas de los Oteros
24,León,063,4,Cubillas de Rueda
24,León,064,9,Cubillos del Sil
24,León,066,5,Destriana
24,León,067,1,Encinedo
24,León,068,7,"Ercina, La"
24,León,069,0,Escobar de Campos
24,León,070,4,Fabero
24,León,071,1,Folgoso de la Ribera
24,León,073,2,Fresno de la Vega
24,León,074,7,Fuentes de Carbajal
24,León,076,3,Garrafe de Torío
24,León,077,9,Gordaliza del Pino
24,León,078,5,Gordoncillo
24,León,079,8,Gradefes
24,León,080,2,Grajal de Campos
24,León,081,9,Gusendos de los Oteros
24,León,082,4,Hospital de Órbigo
24,León,083,0,Igüeña
24,León,084,5,Izagre
24,León,086,1,Joarilla de las Matas
24,León,087,7,Laguna Dalga
24,León,088,3,Laguna de Negrillos
24,León,089,6,León
24,León,092,2,Llamas de la Ribera
24,León,090,0,Lucillo
24,León,091,7,Luyego
24,León,093,8,Magaz de Cepeda
24,León,094,3,Mansilla de las Mulas
24,León,095,6,Mansilla Mayor
24,León,096,9,Maraña
24,León,097,5,Matadeón de los Oteros
24,León,098,1,Matallana de Torío
24,León,099,4,Matanza
24,León,100,8,Molinaseca
24,León,101,5,Murias de Paredes
24,León,102,0,Noceda del Bierzo
24,León,103,6,Oencia
24,León,104,1,"Omañas, Las"
24,León,105,4,Onzonilla
24,León,106,7,Oseja de Sajambre
24,León,107,3,Pajares de los Oteros
24,León,108,9,Palacios de la Valduerna
24,León,109,2,Palacios del Sil
24,León,110,6,Páramo del Sil
24,León,112,8,Peranzanes
24,León,113,4,Pobladura de Pelayo García
24,León,114,9,"Pola de Gordón, La"
24,León,115,2,Ponferrada
24,León,116,5,Posada de Valdeón
24,León,117,1,Pozuelo del Páramo
24,León,118,7,Prado de la Guzpeña
24,León,119,0,Priaranza del Bierzo
24,León,120,4,Prioro
24,León,121,1,Puebla de Lillo
24,León,122,6,Puente de Domingo Flórez
24,León,123,2,Quintana del Castillo
24,León,124,7,Quintana del Marco
24,León,125,0,Quintana y Congosto
24,León,127,9,Regueras de Arriba
24,León,129,8,Reyero
24,León,130,2,Riaño
24,León,131,9,Riego de la Vega
24,León,132,4,Riello
24,León,133,0,Rioseco de Tapia
24,León,134,5,"Robla, La"
24,León,136,1,Roperuelos del Páramo
24,León,137,7,Sabero
24,León,139,6,Sahagún
24,León,141,7,San Adrián del Valle
24,León,142,2,San Andrés del Rabanedo
24,León,144,3,San Cristóbal de la Polantera
24,León,145,6,San Emiliano
24,León,146,9,San Esteban de Nogales
24,León,148,1,San Justo de la Vega
24,León,149,4,San Millán de los Caballeros
24,León,150,7,San Pedro Bercianos
24,León,143,8,Sancedo
24,León,151,4,Santa Colomba de Curueño
24,León,152,9,Santa Colomba de Somoza
24,León,153,5,Santa Cristina de Valmadrigal
24,León,154,0,Santa Elena de Jamuz
24,León,155,3,Santa María de la Isla
24,León,158,8,Santa María de Ordás
24,León,156,6,Santa María del Monte de Cea
24,León,157,2,Santa María del Páramo
24,León,159,1,Santa Marina del Rey
24,León,160,5,Santas Martas
24,León,161,2,Santiago Millas
24,León,162,7,Santovenia de la Valdoncina
24,León,163,3,Sariegos
24,León,164,8,Sena de Luna
24,León,165,1,Sobrado
24,León,166,4,Soto de la Vega
24,León,167,0,Soto y Amío
24,León,168,6,Toral de los Guzmanes
24,León,206,6,Toral de los Vados
24,León,169,9,Toreno
24,León,170,3,Torre del Bierzo
24,León,171,0,Trabadelo
24,León,172,5,Truchas
24,León,173,1,Turcia
24,León,174,6,Urdiales del Páramo
24,León,185,7,Val de San Lorenzo
24,León,175,9,Valdefresno
24,León,176,2,Valdefuentes del Páramo
24,León,177,8,Valdelugueros
24,León,178,4,Valdemora
24,León,179,7,Valdepiélago
24,León,180,1,Valdepolo
24,León,181,8,Valderas
24,León,182,3,Valderrey
24,León,183,9,Valderrueda
24,León,184,4,Valdesamario
24,León,187,6,Valdevimbre
24,León,188,2,Valencia de Don Juan
24,León,191,6,Vallecillo
24,León,189,5,Valverde de la Virgen
24,León,190,9,Valverde-Enrique
24,León,193,7,"Vecilla, La"
24,León,196,8,Vega de Espinareda
24,León,197,4,Vega de Infanzones
24,León,198,0,Vega de Valcarce
24,León,194,2,Vegacervera
24,León,199,3,Vegaquemada
24,León,201,4,Vegas del Condado
24,León,202,9,Villablino
24,León,203,5,Villabraz
24,León,205,3,Villadangos del Páramo
24,León,207,2,Villademor de la Vega
24,León,209,1,Villafranca del Bierzo
24,León,210,5,Villagatón
24,León,211,2,Villamandos
24,León,901,7,Villamanín
24,León,212,7,Villamañán
24,León,213,3,Villamartín de Don Sancho
24,León,214,8,Villamejil
24,León,215,1,Villamol
24,León,216,4,Villamontán de la Valduerna
24,León,217,0,Villamoratiel de las Matas
24,León,218,6,Villanueva de las Manzanas
24,León,219,9,Villaobispo de Otero
24,León,902,2,Villaornate y Castro
24,León,221,0,Villaquejida
24,León,222,5,Villaquilambre
24,León,223,1,Villarejo de Órbigo
24,León,224,6,Villares de Órbigo
24,León,225,9,Villasabariego
24,León,226,2,Villaselán
24,León,227,8,Villaturiel
24,León,228,4,Villazala
24,León,229,7,Villazanzo de Valderaduey
24,León,230,1,Zotes del Páramo
25,Lleida,001,9,Abella de la Conca
25,Lleida,002,4,Àger
25,Lleida,003,0,Agramunt
25,Lleida,038,7,Aitona
25,Lleida,004,5,"Alamús, Els"
25,Lleida,005,8,Alàs i Cerc
25,Lleida,006,1,"Albagés, L'"
25,Lleida,007,7,Albatàrrec
25,Lleida,008,3,Albesa
25,Lleida,009,6,"Albi, L'"
25,Lleida,010,0,Alcanó
25,Lleida,011,7,Alcarràs
25,Lleida,012,2,Alcoletge
25,Lleida,013,8,Alfarràs
25,Lleida,014,3,Alfés
25,Lleida,015,6,Algerri
25,Lleida,016,9,Alguaire
25,Lleida,017,5,Alins
25,Lleida,019,4,Almacelles
25,Lleida,020,8,Almatret
25,Lleida,021,5,Almenar
25,Lleida,022,0,Alòs de Balaguer
25,Lleida,023,6,Alpicat
25,Lleida,024,1,Alt Àneu
25,Lleida,027,3,Anglesola
25,Lleida,029,2,Arbeca
25,Lleida,031,3,Arres
25,Lleida,032,8,Arsèguel
25,Lleida,033,4,Artesa de Lleida
25,Lleida,034,9,Artesa de Segre
25,Lleida,036,5,Aspa
25,Lleida,037,1,"Avellanes i Santa Linya, Les"
25,Lleida,039,0,Baix Pallars
25,Lleida,040,4,Balaguer
25,Lleida,041,1,Barbens
25,Lleida,042,6,"Baronia de Rialb, La"
25,Lleida,044,7,Bassella
25,Lleida,045,0,Bausen
25,Lleida,046,3,Belianes
25,Lleida,170,6,Bellaguarda
25,Lleida,047,9,Bellcaire d'Urgell
25,Lleida,048,5,Bell-lloc d'Urgell
25,Lleida,049,8,Bellmunt d'Urgell
25,Lleida,050,1,Bellpuig
25,Lleida,051,8,Bellver de Cerdanya
25,Lleida,052,3,Bellvís
25,Lleida,053,9,Benavent de Segrià
25,Lleida,055,7,Biosca
25,Lleida,057,6,"Bòrdes, Es"
25,Lleida,058,2,"Borges Blanques, Les"
25,Lleida,059,5,Bossòst
25,Lleida,056,0,Bovera
25,Lleida,060,9,Cabanabona
25,Lleida,061,6,Cabó
25,Lleida,062,1,Camarasa
25,Lleida,063,7,Canejan
25,Lleida,904,6,Castell de Mur
25,Lleida,064,2,Castellar de la Ribera
25,Lleida,067,4,Castelldans
25,Lleida,068,0,Castellnou de Seana
25,Lleida,069,3,Castelló de Farfanya
25,Lleida,070,7,Castellserà
25,Lleida,071,4,Cava
25,Lleida,072,9,Cervera
25,Lleida,073,5,Cervià de les Garrigues
25,Lleida,074,0,Ciutadilla
25,Lleida,075,3,Clariana de Cardener
25,Lleida,076,6,"Cogul, El"
25,Lleida,077,2,Coll de Nargó
25,Lleida,163,6,"Coma i la Pedra, La"
25,Lleida,161,5,Conca de Dalt
25,Lleida,078,8,Corbins
25,Lleida,079,1,Cubells
25,Lleida,081,2,"Espluga Calba, L'"
25,Lleida,082,7,Espot
25,Lleida,088,6,Estamariu
25,Lleida,085,1,Estaràs
25,Lleida,086,4,Esterri d'Àneu
25,Lleida,087,0,Esterri de Cardós
25,Lleida,089,9,Farrera
25,Lleida,908,4,Fígols i Alinyà
25,Lleida,092,5,"Floresta, La"
25,Lleida,093,1,Fondarella
25,Lleida,094,6,Foradada
25,Lleida,096,2,"Fuliola, La"
25,Lleida,097,8,Fulleda
25,Lleida,098,4,Gavet de la Conca
25,Lleida,912,3,Gimenells i el Pla de la Font
25,Lleida,099,7,Golmés
25,Lleida,100,1,Gósol
25,Lleida,101,8,"Granadella, La"
25,Lleida,102,3,"Granja d'Escarp, La"
25,Lleida,103,9,Granyanella
25,Lleida,105,7,Granyena de les Garrigues
25,Lleida,104,4,Granyena de Segarra
25,Lleida,109,5,Guimerà
25,Lleida,903,1,"Guingueta d'Àneu, La"
25,Lleida,110,9,Guissona
25,Lleida,111,6,Guixers
25,Lleida,115,5,Isona i Conca Dellà
25,Lleida,112,1,Ivars de Noguera
25,Lleida,113,7,Ivars d'Urgell
25,Lleida,114,2,Ivorra
25,Lleida,910,1,Josa i Tuixén
25,Lleida,118,0,Juncosa
25,Lleida,119,3,Juneda
25,Lleida,121,4,Les
25,Lleida,122,9,Linyola
25,Lleida,123,5,Lladorre
25,Lleida,124,0,Lladurs
25,Lleida,125,3,Llardecans
25,Lleida,126,6,Llavorsí
25,Lleida,120,7,Lleida
25,Lleida,127,2,Lles de Cerdanya
25,Lleida,128,8,Llimiana
25,Lleida,129,1,Llobera
25,Lleida,133,3,Maials
25,Lleida,130,5,Maldà
25,Lleida,131,2,Massalcoreig
25,Lleida,132,7,Massoteres
25,Lleida,134,8,Menàrguens
25,Lleida,135,1,Miralcamp
25,Lleida,137,0,Mollerussa
25,Lleida,136,4,"Molsosa, La"
25,Lleida,139,9,Montellà i Martinet
25,Lleida,140,3,Montferrer i Castellbò
25,Lleida,138,6,Montgai
25,Lleida,142,5,Montoliu de Lleida
25,Lleida,141,0,Montoliu de Segarra
25,Lleida,143,1,Montornès de Segarra
25,Lleida,145,9,Nalec
25,Lleida,025,4,Naut Aran
25,Lleida,146,2,Navès
25,Lleida,148,4,Odèn
25,Lleida,149,7,Oliana
25,Lleida,150,0,Oliola
25,Lleida,151,7,Olius
25,Lleida,152,2,"Oluges, Les"
25,Lleida,153,8,"Omellons, Els"
25,Lleida,154,3,"Omells de na Gaia, Els"
25,Lleida,155,6,Organyà
25,Lleida,156,9,Os de Balaguer
25,Lleida,157,5,Ossó de Sió
25,Lleida,158,1,"Palau d'Anglesola, El"
25,Lleida,164,1,Penelles
25,Lleida,165,4,Peramola
25,Lleida,166,7,Pinell de Solsonès
25,Lleida,167,3,Pinós
25,Lleida,911,8,"Plans de Sió, Els"
25,Lleida,168,9,"Poal, El"
25,Lleida,169,2,"Pobla de Cérvoles, La"
25,Lleida,171,3,"Pobla de Segur, La"
25,Lleida,030,6,"Pont de Bar, El"
25,Lleida,173,4,"Pont de Suert, El"
25,Lleida,172,8,Ponts
25,Lleida,174,9,"Portella, La"
25,Lleida,175,2,Prats i Sansor
25,Lleida,176,5,Preixana
25,Lleida,177,1,Preixens
25,Lleida,179,0,Prullans
25,Lleida,180,4,Puiggròs
25,Lleida,181,1,Puigverd d'Agramunt
25,Lleida,182,6,Puigverd de Lleida
25,Lleida,183,2,Rialp
25,Lleida,905,9,Ribera d'Ondara
25,Lleida,185,0,Ribera d'Urgellet
25,Lleida,186,3,Riner
25,Lleida,913,9,Riu de Cerdanya
25,Lleida,189,8,Rosselló
25,Lleida,190,2,Salàs de Pallars
25,Lleida,191,9,Sanaüja
25,Lleida,196,1,Sant Esteve de la Sarga
25,Lleida,192,4,Sant Guim de Freixenet
25,Lleida,197,7,Sant Guim de la Plana
25,Lleida,193,0,Sant Llorenç de Morunys
25,Lleida,902,5,Sant Martí de Riucorb
25,Lleida,194,5,Sant Ramon
25,Lleida,201,7,Sarroca de Bellera
25,Lleida,200,0,Sarroca de Lleida
25,Lleida,202,2,Senterada
25,Lleida,035,2,"Sentiu de Sió, La"
25,Lleida,204,3,Seròs
25,Lleida,203,8,"Seu d'Urgell, La"
25,Lleida,205,6,Sidamon
25,Lleida,206,9,"Soleràs, El"
25,Lleida,207,5,Solsona
25,Lleida,208,1,Soriguera
25,Lleida,209,4,Sort
25,Lleida,210,8,Soses
25,Lleida,211,5,Sudanell
25,Lleida,212,0,Sunyer
25,Lleida,215,4,Talarn
25,Lleida,216,7,Talavera
25,Lleida,217,3,Tàrrega
25,Lleida,218,9,Tarrés
25,Lleida,219,2,Tarroja de Segarra
25,Lleida,220,6,Térmens
25,Lleida,221,3,Tírvia
25,Lleida,222,8,Tiurana
25,Lleida,223,4,Torà
25,Lleida,224,9,"Torms, Els"
25,Lleida,225,2,Tornabous
25,Lleida,227,1,"Torre de Cabdella, La"
25,Lleida,226,5,Torrebesses
25,Lleida,228,7,Torrefarrera
25,Lleida,907,8,Torrefeta i Florejacs
25,Lleida,230,4,Torregrossa
25,Lleida,231,1,Torrelameu
25,Lleida,232,6,Torres de Segre
25,Lleida,233,2,Torre-serona
25,Lleida,234,7,Tremp
25,Lleida,043,2,"Vall de Boí, La"
25,Lleida,901,0,Vall de Cardós
25,Lleida,238,5,Vallbona de les Monges
25,Lleida,240,2,Vallfogona de Balaguer
25,Lleida,906,2,"Valls d'Aguilar, Les"
25,Lleida,239,8,"Valls de Valira, Les"
25,Lleida,909,7,"Vansa i Fórnols, La"
25,Lleida,242,4,Verdú
25,Lleida,243,0,Vielha e Mijaran
25,Lleida,244,5,Vilagrassa
25,Lleida,245,8,Vilaller
25,Lleida,247,7,Vilamòs
25,Lleida,248,3,Vilanova de Bellpuig
25,Lleida,254,2,Vilanova de la Barca
25,Lleida,249,6,Vilanova de l'Aguda
25,Lleida,250,9,Vilanova de Meià
25,Lleida,251,6,Vilanova de Segrià
25,Lleida,252,1,Vila-sana
25,Lleida,253,7,"Vilosell, El"
25,Lleida,255,5,Vinaixa
26,La Rioja,000,0,Logrono
26,La Rioja,000,0,Logroo
26,La Rioja,001,2,Ábalos
26,La Rioja,002,7,Agoncillo
26,La Rioja,003,3,Aguilar del Río Alhama
26,La Rioja,004,8,Ajamil de Cameros
26,La Rioja,005,1,Albelda de Iregua
26,La Rioja,006,4,Alberite
26,La Rioja,007,0,Alcanadre
26,La Rioja,008,6,Aldeanueva de Ebro
26,La Rioja,009,9,Alesanco
26,La Rioja,010,3,Alesón
26,La Rioja,011,0,Alfaro
26,La Rioja,012,5,Almarza de Cameros
26,La Rioja,013,1,Anguciana
26,La Rioja,014,6,Anguiano
26,La Rioja,015,9,Arenzana de Abajo
26,La Rioja,016,2,Arenzana de Arriba
26,La Rioja,017,8,Arnedillo
26,La Rioja,018,4,Arnedo
26,La Rioja,019,7,Arrúbal
26,La Rioja,020,1,Ausejo
26,La Rioja,021,8,Autol
26,La Rioja,022,3,Azofra
26,La Rioja,023,9,Badarán
26,La Rioja,024,4,Bañares
26,La Rioja,026,0,Baños de Río Tobía
26,La Rioja,025,7,Baños de Rioja
26,La Rioja,027,6,Berceo
26,La Rioja,028,2,Bergasa
26,La Rioja,029,5,Bergasillas Bajera
26,La Rioja,030,9,Bezares
26,La Rioja,031,6,Bobadilla
26,La Rioja,032,1,Brieva de Cameros
26,La Rioja,033,7,Briñas
26,La Rioja,034,2,Briones
26,La Rioja,035,5,Cabezón de Cameros
26,La Rioja,036,8,Calahorra
26,La Rioja,037,4,Camprovín
26,La Rioja,038,0,Canales de la Sierra
26,La Rioja,039,3,Canillas de Río Tuerto
26,La Rioja,040,7,Cañas
26,La Rioja,041,4,Cárdenas
26,La Rioja,042,9,Casalarreina
26,La Rioja,043,5,Castañares de Rioja
26,La Rioja,044,0,Castroviejo
26,La Rioja,045,3,Cellorigo
26,La Rioja,046,6,Cenicero
26,La Rioja,047,2,Cervera del Río Alhama
26,La Rioja,048,8,Cidamón
26,La Rioja,049,1,Cihuri
26,La Rioja,050,4,Cirueña
26,La Rioja,051,1,Clavijo
26,La Rioja,052,6,Cordovín
26,La Rioja,053,2,Corera
26,La Rioja,054,7,Cornago
26,La Rioja,055,0,Corporales
26,La Rioja,056,3,Cuzcurrita de Río Tirón
26,La Rioja,057,9,Daroca de Rioja
26,La Rioja,058,5,Enciso
26,La Rioja,059,8,Entrena
26,La Rioja,060,2,Estollo
26,La Rioja,061,9,Ezcaray
26,La Rioja,062,4,Foncea
26,La Rioja,063,0,Fonzaleche
26,La Rioja,064,5,Fuenmayor
26,La Rioja,065,8,Galbárruli
26,La Rioja,066,1,Galilea
26,La Rioja,067,7,Gallinero de Cameros
26,La Rioja,068,3,Gimileo
26,La Rioja,069,6,Grañón
26,La Rioja,070,0,Grávalos
26,La Rioja,071,7,Haro
26,La Rioja,072,2,Herce
26,La Rioja,073,8,Herramélluri
26,La Rioja,074,3,Hervías
26,La Rioja,075,6,Hormilla
26,La Rioja,076,9,Hormilleja
26,La Rioja,077,5,Hornillos de Cameros
26,La Rioja,078,1,Hornos de Moncalvillo
26,La Rioja,079,4,Huércanos
26,La Rioja,080,8,Igea
26,La Rioja,081,5,Jalón de Cameros
26,La Rioja,082,0,Laguna de Cameros
26,La Rioja,083,6,Lagunilla del Jubera
26,La Rioja,084,1,Lardero
26,La Rioja,086,7,Ledesma de la Cogolla
26,La Rioja,087,3,Leiva
26,La Rioja,088,9,Leza de Río Leza
26,La Rioja,089,2,Logroño
26,La Rioja,091,3,Lumbreras
26,La Rioja,092,8,Manjarrés
26,La Rioja,093,4,Mansilla de la Sierra
26,La Rioja,094,9,Manzanares de Rioja
26,La Rioja,095,2,Matute
26,La Rioja,096,5,Medrano
26,La Rioja,098,7,Munilla
26,La Rioja,099,0,Murillo de Río Leza
26,La Rioja,100,4,Muro de Aguas
26,La Rioja,101,1,Muro en Cameros
26,La Rioja,102,6,Nájera
26,La Rioja,103,2,Nalda
26,La Rioja,104,7,Navajún
26,La Rioja,105,0,Navarrete
26,La Rioja,106,3,Nestares
26,La Rioja,107,9,Nieva de Cameros
26,La Rioja,109,8,Ochánduri
26,La Rioja,108,5,Ocón
26,La Rioja,110,2,Ojacastro
26,La Rioja,111,9,Ollauri
26,La Rioja,112,4,Ortigosa de Cameros
26,La Rioja,113,0,Pazuengos
26,La Rioja,114,5,Pedroso
26,La Rioja,115,8,Pinillos
26,La Rioja,117,7,Pradejón
26,La Rioja,118,3,Pradillo
26,La Rioja,119,6,Préjano
26,La Rioja,120,0,Quel
26,La Rioja,121,7,Rabanera
26,La Rioja,122,2,"Rasillo de Cameros, El"
26,La Rioja,123,8,"Redal, El"
26,La Rioja,124,3,Ribafrecha
26,La Rioja,125,6,Rincón de Soto
26,La Rioja,126,9,Robres del Castillo
26,La Rioja,127,5,Rodezno
26,La Rioja,128,1,Sajazarra
26,La Rioja,129,4,San Asensio
26,La Rioja,130,8,San Millán de la Cogolla
26,La Rioja,131,5,San Millán de Yécora
26,La Rioja,132,0,San Román de Cameros
26,La Rioja,139,2,San Torcuato
26,La Rioja,142,8,San Vicente de la Sonsierra
26,La Rioja,134,1,Santa Coloma
26,La Rioja,135,4,Santa Engracia del Jubera
26,La Rioja,136,7,Santa Eulalia Bajera
26,La Rioja,138,9,Santo Domingo de la Calzada
26,La Rioja,140,6,Santurde de Rioja
26,La Rioja,141,3,Santurdejo
26,La Rioja,143,4,Sojuela
26,La Rioja,144,9,Sorzano
26,La Rioja,145,2,Sotés
26,La Rioja,146,5,Soto en Cameros
26,La Rioja,147,1,Terroba
26,La Rioja,148,7,Tirgo
26,La Rioja,149,0,Tobía
26,La Rioja,150,3,Tormantos
26,La Rioja,153,1,Torre en Cameros
26,La Rioja,151,0,Torrecilla en Cameros
26,La Rioja,152,5,Torrecilla sobre Alesanco
26,La Rioja,154,6,Torremontalbo
26,La Rioja,155,9,Treviana
26,La Rioja,157,8,Tricio
26,La Rioja,158,4,Tudelilla
26,La Rioja,160,1,Uruñuela
26,La Rioja,161,8,Valdemadera
26,La Rioja,162,3,Valgañón
26,La Rioja,163,9,Ventosa
26,La Rioja,164,4,Ventrosa
26,La Rioja,165,7,Viguera
26,La Rioja,166,0,Villalba de Rioja
26,La Rioja,167,6,Villalobar de Rioja
26,La Rioja,168,2,Villamediana de Iregua
26,La Rioja,169,5,Villanueva de Cameros
26,La Rioja,170,9,"Villar de Arnedo, El"
26,La Rioja,171,6,Villar de Torre
26,La Rioja,172,1,Villarejo
26,La Rioja,173,7,Villarroya
26,La Rioja,174,2,Villarta-Quintana
26,La Rioja,175,5,Villavelayo
26,La Rioja,176,8,Villaverde de Rioja
26,La Rioja,177,4,Villoslada de Cameros
26,La Rioja,178,0,Viniegra de Abajo
26,La Rioja,179,3,Viniegra de Arriba
26,La Rioja,180,7,Zarratón
26,La Rioja,181,4,Zarzosa
26,La Rioja,183,5,Zorraquín
27,Lugo,001,8,Abadín
27,Lugo,002,3,Alfoz
27,Lugo,003,9,Antas de Ulla
27,Lugo,004,4,Baleira
27,Lugo,901,9,Baralla
27,Lugo,005,7,Barreiros
27,Lugo,006,0,Becerreá
27,Lugo,007,6,Begonte
27,Lugo,008,2,Bóveda
27,Lugo,902,4,Burela
27,Lugo,009,5,Carballedo
27,Lugo,010,9,Castro de Rei
27,Lugo,011,6,Castroverde
27,Lugo,012,1,Cervantes
27,Lugo,013,7,Cervo
27,Lugo,016,8,Chantada
27,Lugo,014,2,"Corgo, O"
27,Lugo,015,5,Cospeito
27,Lugo,017,4,Folgoso do Courel
27,Lugo,018,0,"Fonsagrada, A"
27,Lugo,019,3,Foz
27,Lugo,020,7,Friol
27,Lugo,022,9,Guitiriz
27,Lugo,023,5,Guntín
27,Lugo,024,0,"Incio, O"
27,Lugo,026,6,Láncara
27,Lugo,027,2,Lourenzá
27,Lugo,028,8,Lugo
27,Lugo,029,1,Meira
27,Lugo,030,5,Mondoñedo
27,Lugo,031,2,Monforte de Lemos
27,Lugo,032,7,Monterroso
27,Lugo,033,3,Muras
27,Lugo,034,8,Navia de Suarna
27,Lugo,035,1,Negueira de Muñiz
27,Lugo,037,0,"Nogais, As"
27,Lugo,038,6,Ourol
27,Lugo,039,9,Outeiro de Rei
27,Lugo,040,3,Palas de Rei
27,Lugo,041,0,Pantón
27,Lugo,042,5,Paradela
27,Lugo,043,1,"Páramo, O"
27,Lugo,044,6,"Pastoriza, A"
27,Lugo,045,9,Pedrafita do Cebreiro
27,Lugo,047,8,"Pobra do Brollón, A"
27,Lugo,046,2,Pol
27,Lugo,048,4,"Pontenova, A"
27,Lugo,049,7,Portomarín
27,Lugo,050,0,Quiroga
27,Lugo,056,9,Rábade
27,Lugo,051,7,Ribadeo
27,Lugo,052,2,Ribas de Sil
27,Lugo,053,8,Ribeira de Piquín
27,Lugo,054,3,Riotorto
27,Lugo,055,6,Samos
27,Lugo,057,5,Sarria
27,Lugo,058,1,"Saviñao, O"
27,Lugo,059,4,Sober
27,Lugo,060,8,Taboada
27,Lugo,061,5,Trabada
27,Lugo,062,0,Triacastela
27,Lugo,063,6,"Valadouro, O"
27,Lugo,064,1,"Vicedo, O"
27,Lugo,065,4,Vilalba
27,Lugo,066,7,Viveiro
27,Lugo,021,4,Xermade
27,Lugo,025,3,Xove
28,Madrid,000,0,Mostoles
28,Madrid,000,0,Alcala de Henares
28,Madrid,000,0,Torrejon de Ardoz
28,Madrid,000,0,Alcorcon
28,Madrid,000,0,Leganes
28,Madrid,001,4,"Acebeda, La"
28,Madrid,002,9,Ajalvir
28,Madrid,003,5,Alameda del Valle
28,Madrid,004,0,"Álamo, El"
28,Madrid,005,3,Alcalá de Henares
28,Madrid,006,6,Alcobendas
28,Madrid,007,2,Alcorcón
28,Madrid,008,8,Aldea del Fresno
28,Madrid,009,1,Algete
28,Madrid,010,5,Alpedrete
28,Madrid,011,2,Ambite
28,Madrid,012,7,Anchuelo
28,Madrid,013,3,Aranjuez
28,Madrid,014,8,Arganda del Rey
28,Madrid,015,1,Arroyomolinos
28,Madrid,016,4,"Atazar, El"
28,Madrid,017,0,Batres
28,Madrid,018,6,Becerril de la Sierra
28,Madrid,019,9,Belmonte de Tajo
28,Madrid,021,0,"Berrueco, El"
28,Madrid,020,3,Berzosa del Lozoya
28,Madrid,022,5,Boadilla del Monte
28,Madrid,023,1,"Boalo, El"
28,Madrid,024,6,Braojos
28,Madrid,025,9,Brea de Tajo
28,Madrid,026,2,Brunete
28,Madrid,027,8,Buitrago del Lozoya
28,Madrid,028,4,Bustarviejo
28,Madrid,029,7,Cabanillas de la Sierra
28,Madrid,030,1,"Cabrera, La"
28,Madrid,031,8,Cadalso de los Vidrios
28,Madrid,032,3,Camarma de Esteruelas
28,Madrid,033,9,Campo Real
28,Madrid,034,4,Canencia
28,Madrid,035,7,Carabaña
28,Madrid,036,0,Casarrubuelos
28,Madrid,037,6,Cenicientos
28,Madrid,038,2,Cercedilla
28,Madrid,039,5,Cervera de Buitrago
28,Madrid,051,3,Chapinería
28,Madrid,052,8,Chinchón
28,Madrid,040,9,Ciempozuelos
28,Madrid,041,6,Cobeña
28,Madrid,046,8,Collado Mediano
28,Madrid,047,4,Collado Villalba
28,Madrid,043,7,Colmenar de Oreja
28,Madrid,042,1,Colmenar del Arroyo
28,Madrid,045,5,Colmenar Viejo
28,Madrid,044,2,Colmenarejo
28,Madrid,048,0,Corpa
28,Madrid,049,3,Coslada
28,Madrid,050,6,Cubas de la Sagra
28,Madrid,053,4,Daganzo de Arriba
28,Madrid,054,9,"Escorial, El"
28,Madrid,055,2,Estremera
28,Madrid,056,5,Fresnedillas de la Oliva
28,Madrid,057,1,Fresno de Torote
28,Madrid,058,7,Fuenlabrada
28,Madrid,059,0,Fuente el Saz de Jarama
28,Madrid,060,4,Fuentidueña de Tajo
28,Madrid,061,1,Galapagar
28,Madrid,062,6,Garganta de los Montes
28,Madrid,063,2,Gargantilla del Lozoya y Pinilla de Buitrago
28,Madrid,064,7,Gascones
28,Madrid,065,0,Getafe
28,Madrid,066,3,Griñón
28,Madrid,067,9,Guadalix de la Sierra
28,Madrid,068,5,Guadarrama
28,Madrid,069,8,"Hiruela, La"
28,Madrid,070,2,Horcajo de la Sierra-Aoslos
28,Madrid,071,9,Horcajuelo de la Sierra
28,Madrid,072,4,Hoyo de Manzanares
28,Madrid,073,0,Humanes de Madrid
28,Madrid,074,5,Leganés
28,Madrid,075,8,Loeches
28,Madrid,076,1,Lozoya
28,Madrid,901,5,Lozoyuela-Navas-Sieteiglesias
28,Madrid,078,3,Madarcos
28,Madrid,079,6,Madrid
28,Madrid,080,0,Majadahonda
28,Madrid,082,2,Manzanares el Real
28,Madrid,083,8,Meco
28,Madrid,084,3,Mejorada del Campo
28,Madrid,085,6,Miraflores de la Sierra
28,Madrid,086,9,"Molar, El"
28,Madrid,087,5,"Molinos, Los"
28,Madrid,088,1,Montejo de la Sierra
28,Madrid,089,4,Moraleja de Enmedio
28,Madrid,090,8,Moralzarzal
28,Madrid,091,5,Morata de Tajuña
28,Madrid,092,0,Móstoles
28,Madrid,093,6,Navacerrada
28,Madrid,094,1,Navalafuente
28,Madrid,095,4,Navalagamella
28,Madrid,096,7,Navalcarnero
28,Madrid,097,3,Navarredonda y San Mamés
28,Madrid,099,2,Navas del Rey
28,Madrid,100,6,Nuevo Baztán
28,Madrid,101,3,Olmeda de las Fuentes
28,Madrid,102,8,Orusco de Tajuña
28,Madrid,104,9,Paracuellos de Jarama
28,Madrid,106,5,Parla
28,Madrid,107,1,Patones
28,Madrid,108,7,Pedrezuela
28,Madrid,109,0,Pelayos de la Presa
28,Madrid,110,4,Perales de Tajuña
28,Madrid,111,1,Pezuela de las Torres
28,Madrid,112,6,Pinilla del Valle
28,Madrid,113,2,Pinto
28,Madrid,114,7,Piñuécar-Gandullas
28,Madrid,115,0,Pozuelo de Alarcón
28,Madrid,116,3,Pozuelo del Rey
28,Madrid,117,9,Prádena del Rincón
28,Madrid,118,5,Puebla de la Sierra
28,Madrid,902,0,Puentes Viejas
28,Madrid,119,8,Quijorna
28,Madrid,120,2,Rascafría
28,Madrid,121,9,Redueña
28,Madrid,122,4,Ribatejada
28,Madrid,123,0,Rivas-Vaciamadrid
28,Madrid,124,5,Robledillo de la Jara
28,Madrid,125,8,Robledo de Chavela
28,Madrid,126,1,Robregordo
28,Madrid,127,7,"Rozas de Madrid, Las"
28,Madrid,128,3,Rozas de Puerto Real
28,Madrid,129,6,San Agustín del Guadalix
28,Madrid,130,0,San Fernando de Henares
28,Madrid,131,7,San Lorenzo de El Escorial
28,Madrid,132,2,San Martín de la Vega
28,Madrid,133,8,San Martín de Valdeiglesias
28,Madrid,134,3,San Sebastián de los Reyes
28,Madrid,135,6,Santa María de la Alameda
28,Madrid,136,9,Santorcaz
28,Madrid,137,5,"Santos de la Humosa, Los"
28,Madrid,138,1,"Serna del Monte, La"
28,Madrid,140,8,Serranillos del Valle
28,Madrid,141,5,Sevilla la Nueva
28,Madrid,143,6,Somosierra
28,Madrid,144,1,Soto del Real
28,Madrid,145,4,Talamanca de Jarama
28,Madrid,146,7,Tielmes
28,Madrid,147,3,Titulcia
28,Madrid,148,9,Torrejón de Ardoz
28,Madrid,149,2,Torrejón de la Calzada
28,Madrid,150,5,Torrejón de Velasco
28,Madrid,151,2,Torrelaguna
28,Madrid,152,7,Torrelodones
28,Madrid,153,3,Torremocha de Jarama
28,Madrid,154,8,Torres de la Alameda
28,Madrid,903,6,Tres Cantos
28,Madrid,155,1,Valdaracete
28,Madrid,156,4,Valdeavero
28,Madrid,157,0,Valdelaguna
28,Madrid,158,6,Valdemanco
28,Madrid,159,9,Valdemaqueda
28,Madrid,160,3,Valdemorillo
28,Madrid,161,0,Valdemoro
28,Madrid,162,5,Valdeolmos-Alalpardo
28,Madrid,163,1,Valdepiélagos
28,Madrid,164,6,Valdetorres de Jarama
28,Madrid,165,9,Valdilecha
28,Madrid,166,2,Valverde de Alcalá
28,Madrid,167,8,Velilla de San Antonio
28,Madrid,168,4,"Vellón, El"
28,Madrid,169,7,Venturada
28,Madrid,171,8,Villa del Prado
28,Madrid,170,1,Villaconejos
28,Madrid,172,3,Villalbilla
28,Madrid,173,9,Villamanrique de Tajo
28,Madrid,174,4,Villamanta
28,Madrid,175,7,Villamantilla
28,Madrid,176,0,Villanueva de la Cañada
28,Madrid,178,2,Villanueva de Perales
28,Madrid,177,6,Villanueva del Pardillo
28,Madrid,179,5,Villar del Olmo
28,Madrid,180,9,Villarejo de Salvanés
28,Madrid,181,6,Villaviciosa de Odón
28,Madrid,182,1,Villavieja del Lozoya
28,Madrid,183,7,Zarzalejo
29,Málaga,000,0,Malaga
29,Málaga,001,7,Alameda
29,Málaga,002,2,Alcaucín
29,Málaga,003,8,Alfarnate
29,Málaga,004,3,Alfarnatejo
29,Málaga,005,6,Algarrobo
29,Málaga,006,9,Algatocín
29,Málaga,007,5,Alhaurín de la Torre
29,Málaga,008,1,Alhaurín el Grande
29,Málaga,009,4,Almáchar
29,Málaga,010,8,Almargen
29,Málaga,011,5,Almogía
29,Málaga,012,0,Álora
29,Málaga,013,6,Alozaina
29,Málaga,014,1,Alpandeire
29,Málaga,015,4,Antequera
29,Málaga,016,7,Árchez
29,Málaga,017,3,Archidona
29,Málaga,018,9,Ardales
29,Málaga,019,2,Arenas
29,Málaga,020,6,Arriate
29,Málaga,021,3,Atajate
29,Málaga,022,8,Benadalid
29,Málaga,023,4,Benahavís
29,Málaga,024,9,Benalauría
29,Málaga,025,2,Benalmádena
29,Málaga,026,5,Benamargosa
29,Málaga,027,1,Benamocarra
29,Málaga,028,7,Benaoján
29,Málaga,029,0,Benarrabá
29,Málaga,030,4,"Borge, El"
29,Málaga,031,1,"Burgo, El"
29,Málaga,032,6,Campillos
29,Málaga,033,2,Canillas de Aceituno
29,Málaga,034,7,Canillas de Albaida
29,Málaga,035,0,Cañete la Real
29,Málaga,036,3,Carratraca
29,Málaga,037,9,Cartajima
29,Málaga,038,5,Cártama
29,Málaga,039,8,Casabermeja
29,Málaga,040,2,Casarabonela
29,Málaga,041,9,Casares
29,Málaga,042,4,Coín
29,Málaga,043,0,Colmenar
29,Málaga,044,5,Comares
29,Málaga,045,8,Cómpeta
29,Málaga,046,1,Cortes de la Frontera
29,Málaga,047,7,Cuevas Bajas
29,Málaga,049,6,Cuevas de San Marcos
29,Málaga,048,3,Cuevas del Becerro
29,Málaga,050,9,Cútar
29,Málaga,051,6,Estepona
29,Málaga,052,1,Faraján
29,Málaga,053,7,Frigiliana
29,Málaga,054,2,Fuengirola
29,Málaga,055,5,Fuente de Piedra
29,Málaga,056,8,Gaucín
29,Málaga,057,4,Genalguacil
29,Málaga,058,0,Guaro
29,Málaga,059,3,Humilladero
29,Málaga,060,7,Igualeja
29,Málaga,061,4,Istán
29,Málaga,062,9,Iznate
29,Málaga,063,5,Jimera de Líbar
29,Málaga,064,0,Jubrique
29,Málaga,065,3,Júzcar
29,Málaga,066,6,Macharaviaya
29,Málaga,067,2,Málaga
29,Málaga,068,8,Manilva
29,Málaga,069,1,Marbella
29,Málaga,070,5,Mijas
29,Málaga,071,2,Moclinejo
29,Málaga,072,7,Mollina
29,Málaga,073,3,Monda
29,Málaga,074,8,Montejaque
29,Málaga,075,1,Nerja
29,Málaga,076,4,Ojén
29,Málaga,077,0,Parauta
29,Málaga,079,9,Periana
29,Málaga,080,3,Pizarra
29,Málaga,081,0,Pujerra
29,Málaga,082,5,Rincón de la Victoria
29,Málaga,083,1,Riogordo
29,Málaga,084,6,Ronda
29,Málaga,085,9,Salares
29,Málaga,086,2,Sayalonga
29,Málaga,087,8,Sedella
29,Málaga,088,4,Sierra de Yeguas
29,Málaga,089,7,Teba
29,Málaga,090,1,Tolox
29,Málaga,901,8,Torremolinos
29,Málaga,091,8,Torrox
29,Málaga,092,3,Totalán
29,Málaga,093,9,Valle de Abdalajís
29,Málaga,094,4,Vélez-Málaga
29,Málaga,095,7,Villanueva de Algaidas
29,Málaga,902,3,Villanueva de la Concepción
29,Málaga,098,2,Villanueva de Tapia
29,Málaga,096,0,Villanueva del Rosario
29,Málaga,097,6,Villanueva del Trabuco
29,Málaga,099,5,Viñuela
29,Málaga,100,9,Yunquera
30,Murcia,001,1,Abanilla
30,Murcia,002,6,Abarán
30,Murcia,003,2,Águilas
30,Murcia,004,7,Albudeite
30,Murcia,005,0,Alcantarilla
30,Murcia,902,7,"Alcázares, Los"
30,Murcia,006,3,Aledo
30,Murcia,007,9,Alguazas
30,Murcia,008,5,Alhama de Murcia
30,Murcia,009,8,Archena
30,Murcia,010,2,Beniel
30,Murcia,011,9,Blanca
30,Murcia,012,4,Bullas
30,Murcia,013,0,Calasparra
30,Murcia,014,5,Campos del Río
30,Murcia,015,8,Caravaca de la Cruz
30,Murcia,016,1,Cartagena
30,Murcia,017,7,Cehegín
30,Murcia,018,3,Ceutí
30,Murcia,019,6,Cieza
30,Murcia,020,0,Fortuna
30,Murcia,021,7,Fuente Álamo de Murcia
30,Murcia,022,2,Jumilla
30,Murcia,023,8,Librilla
30,Murcia,024,3,Lorca
30,Murcia,025,6,Lorquí
30,Murcia,026,9,Mazarrón
30,Murcia,027,5,Molina de Segura
30,Murcia,028,1,Moratalla
30,Murcia,029,4,Mula
30,Murcia,030,8,Murcia
30,Murcia,031,5,Ojós
30,Murcia,032,0,Pliego
30,Murcia,033,6,Puerto Lumbreras
30,Murcia,034,1,Ricote
30,Murcia,035,4,San Javier
30,Murcia,036,7,San Pedro del Pinatar
30,Murcia,901,2,Santomera
30,Murcia,037,3,Torre-Pacheco
30,Murcia,038,9,"Torres de Cotillas, Las"
30,Murcia,039,2,Totana
30,Murcia,040,6,Ulea
30,Murcia,041,3,"Unión, La"
30,Murcia,042,8,Villanueva del Río Segura
30,Murcia,043,4,Yecla
31,Navarra,000,0,Pamplona
31,Navarra,001,8,Abáigar
31,Navarra,002,3,Abárzuza
31,Navarra,003,9,Abaurregaina/Abaurrea Alta
31,Navarra,004,4,Abaurrepea/Abaurrea Baja
31,Navarra,005,7,Aberin
31,Navarra,006,0,Ablitas
31,Navarra,007,6,Adiós
31,Navarra,008,2,Aguilar de Codés
31,Navarra,009,5,Aibar/Oibar
31,Navarra,011,6,Allín
31,Navarra,012,1,Allo
31,Navarra,010,9,Altsasu/Alsasua
31,Navarra,013,7,Améscoa Baja
31,Navarra,014,2,Ancín/Antzin
31,Navarra,015,5,Andosilla
31,Navarra,016,8,Ansoáin/Antsoain
31,Navarra,017,4,Anue
31,Navarra,018,0,Añorbe
31,Navarra,019,3,Aoiz/Agoitz
31,Navarra,020,7,Araitz
31,Navarra,025,3,Arakil
31,Navarra,021,4,Aranarache/Aranaratxe
31,Navarra,023,5,Aranguren
31,Navarra,024,0,Arano
31,Navarra,022,9,Arantza
31,Navarra,026,6,Aras
31,Navarra,027,2,Arbizu
31,Navarra,028,8,Arce/Artzi
31,Navarra,029,1,"Arcos, Los"
31,Navarra,030,5,Arellano
31,Navarra,031,2,Areso
31,Navarra,032,7,Arguedas
31,Navarra,033,3,Aria
31,Navarra,034,8,Aribe
31,Navarra,035,1,Armañanzas
31,Navarra,036,4,Arróniz
31,Navarra,037,0,Arruazu
31,Navarra,038,6,Artajona
31,Navarra,039,9,Artazu
31,Navarra,040,3,Atez
31,Navarra,058,1,Auritz/Burguete
31,Navarra,041,0,Ayegui/Aiegi
31,Navarra,042,5,Azagra
31,Navarra,043,1,Azuelo
31,Navarra,044,6,Bakaiku
31,Navarra,901,9,Barañain
31,Navarra,045,9,Barásoain
31,Navarra,046,2,Barbarin
31,Navarra,047,8,Bargota
31,Navarra,048,4,Barillas
31,Navarra,049,7,Basaburua
31,Navarra,050,0,Baztan
31,Navarra,137,9,Beintza-Labaien
31,Navarra,051,7,Beire
31,Navarra,052,2,Belascoáin
31,Navarra,250,8,Bera
31,Navarra,053,8,Berbinzana
31,Navarra,905,8,Beriáin
31,Navarra,902,4,Berrioplano
31,Navarra,903,0,Berriozar
31,Navarra,054,3,Bertizarana
31,Navarra,055,6,Betelu
31,Navarra,253,6,Bidaurreta
31,Navarra,056,9,Biurrun-Olcoz
31,Navarra,057,5,Buñuel
31,Navarra,059,4,Burgui/Burgi
31,Navarra,060,8,Burlada/Burlata
31,Navarra,061,5,"Busto, El"
31,Navarra,062,0,Cabanillas
31,Navarra,063,6,Cabredo
31,Navarra,064,1,Cadreita
31,Navarra,065,4,Caparroso
31,Navarra,066,7,Cárcar
31,Navarra,067,3,Carcastillo
31,Navarra,068,9,Cascante
31,Navarra,069,2,Cáseda
31,Navarra,070,6,Castejón
31,Navarra,071,3,Castillonuevo
31,Navarra,193,9,Cendea de Olza/Oltza Zendea
31,Navarra,072,8,Cintruénigo
31,Navarra,074,9,Cirauqui/Zirauki
31,Navarra,075,2,Ciriza
31,Navarra,076,5,Cizur
31,Navarra,077,1,Corella
31,Navarra,078,7,Cortes
31,Navarra,079,0,Desojo
31,Navarra,080,4,Dicastillo
31,Navarra,081,1,Donamaria
31,Navarra,221,2,Doneztebe/Santesteban
31,Navarra,083,2,Echarri
31,Navarra,086,3,Egüés
31,Navarra,087,9,Elgorriaga
31,Navarra,089,8,Enériz/Eneritz
31,Navarra,090,2,Eratsun
31,Navarra,091,9,Ergoiena
31,Navarra,092,4,Erro
31,Navarra,094,5,Eslava
31,Navarra,095,8,Esparza de Salazar/Espartza Zaraitzu
31,Navarra,096,1,Espronceda
31,Navarra,097,7,Estella/Lizarra
31,Navarra,098,3,Esteribar
31,Navarra,099,6,Etayo
31,Navarra,082,6,Etxalar
31,Navarra,084,7,Etxarri-Aranatz
31,Navarra,085,0,Etxauri
31,Navarra,100,0,Eulate
31,Navarra,101,7,Ezcabarte
31,Navarra,093,0,Ezcároz/Ezkaroze
31,Navarra,102,2,Ezkurra
31,Navarra,103,8,Ezprogui
31,Navarra,104,3,Falces
31,Navarra,105,6,Fitero
31,Navarra,106,9,Fontellas
31,Navarra,107,5,Funes
31,Navarra,108,1,Fustiñana
31,Navarra,109,4,Galar
31,Navarra,110,8,Gallipienzo/Galipentzu
31,Navarra,111,5,Gallués/Galoze
31,Navarra,112,0,Garaioa
31,Navarra,113,6,Garde
31,Navarra,114,1,Garínoain
31,Navarra,115,4,Garralda
31,Navarra,116,7,Genevilla
31,Navarra,117,3,Goizueta
31,Navarra,118,9,Goñi
31,Navarra,119,2,Güesa/Gorza
31,Navarra,120,6,Guesálaz/Gesalatz
31,Navarra,121,3,Guirguillano
31,Navarra,256,7,Hiriberri/Villanueva de Aezkoa
31,Navarra,122,8,Huarte/Uharte
31,Navarra,124,9,Ibargoiti
31,Navarra,259,2,Igantzi
31,Navarra,125,2,Igúzquiza
31,Navarra,126,5,Imotz
31,Navarra,127,1,Irañeta
31,Navarra,904,5,Irurtzun
31,Navarra,128,7,Isaba/Izaba
31,Navarra,129,0,Ituren
31,Navarra,130,4,Iturmendi
31,Navarra,131,1,Iza/Itza
31,Navarra,132,6,Izagaondoa
31,Navarra,133,2,Izalzu/Itzaltzu
31,Navarra,134,7,Jaurrieta
31,Navarra,135,0,Javier
31,Navarra,136,3,Juslapeña
31,Navarra,138,5,Lakuntza
31,Navarra,139,8,Lana
31,Navarra,140,2,Lantz
31,Navarra,141,9,Lapoblación
31,Navarra,142,4,Larraga
31,Navarra,143,0,Larraona
31,Navarra,144,5,Larraun
31,Navarra,145,8,Lazagurría
31,Navarra,146,1,Leache
31,Navarra,147,7,Legarda
31,Navarra,148,3,Legaria
31,Navarra,149,6,Leitza
31,Navarra,908,3,Lekunberri
31,Navarra,150,9,Leoz/Leotz
31,Navarra,151,6,Lerga
31,Navarra,152,1,Lerín
31,Navarra,153,7,Lesaka
31,Navarra,154,2,Lezáun
31,Navarra,155,5,Liédena
31,Navarra,156,8,Lizoáin-Arriasgoiti
31,Navarra,157,4,Lodosa
31,Navarra,158,0,Lónguida/Longida
31,Navarra,159,3,Lumbier
31,Navarra,160,7,Luquin
31,Navarra,248,2,Luzaide/Valcarlos
31,Navarra,161,4,Mañeru
31,Navarra,162,9,Marañón
31,Navarra,163,5,Marcilla
31,Navarra,164,0,Mélida
31,Navarra,165,3,Mendavia
31,Navarra,166,6,Mendaza
31,Navarra,167,2,Mendigorría
31,Navarra,168,8,Metauten
31,Navarra,169,1,Milagro
31,Navarra,170,5,Mirafuentes
31,Navarra,171,2,Miranda de Arga
31,Navarra,172,7,Monreal
31,Navarra,173,3,Monteagudo
31,Navarra,174,8,Morentin
31,Navarra,175,1,Mues
31,Navarra,176,4,Murchante
31,Navarra,177,0,Murieta
31,Navarra,178,6,Murillo el Cuende
31,Navarra,179,9,Murillo el Fruto
31,Navarra,180,3,Muruzábal
31,Navarra,181,0,Navascués
31,Navarra,182,5,Nazar
31,Navarra,088,5,Noáin (Valle de Elorz)/Noain (Elortzibar)
31,Navarra,183,1,Obanos
31,Navarra,185,9,Ochagavía/Otsagabia
31,Navarra,184,6,Oco
31,Navarra,186,2,Odieta
31,Navarra,187,8,Oitz
31,Navarra,188,4,Olaibar
31,Navarra,189,7,Olazti/Olazagutía
31,Navarra,190,1,Olejua
31,Navarra,191,8,Olite/Erriberri
31,Navarra,194,4,Ollo
31,Navarra,192,3,Olóriz/Oloritz
31,Navarra,195,7,Orbaitzeta
31,Navarra,196,0,Orbara
31,Navarra,197,6,Orísoain
31,Navarra,906,1,Orkoien
31,Navarra,198,2,Oronz/Orontze
31,Navarra,199,5,Oroz-Betelu/Orotz-Betelu
31,Navarra,211,4,Orreaga/Roncesvalles
31,Navarra,200,9,Oteiza
31,Navarra,201,6,Pamplona/Iruña
31,Navarra,202,1,Peralta/Azkoien
31,Navarra,203,7,Petilla de Aragón
31,Navarra,204,2,Piedramillera
31,Navarra,205,5,Pitillas
31,Navarra,206,8,Puente la Reina/Gares
31,Navarra,207,4,Pueyo
31,Navarra,208,0,Ribaforada
31,Navarra,209,3,Romanzado
31,Navarra,210,7,Roncal/Erronkari
31,Navarra,212,9,Sada
31,Navarra,213,5,Saldías
31,Navarra,214,0,Salinas de Oro/Jaitz
31,Navarra,215,3,San Adrián
31,Navarra,217,2,San Martín de Unx
31,Navarra,216,6,Sangüesa/Zangoza
31,Navarra,219,1,Sansol
31,Navarra,220,5,Santacara
31,Navarra,222,7,Sarriés/Sartze
31,Navarra,223,3,Sartaguda
31,Navarra,224,8,Sesma
31,Navarra,225,1,Sorlada
31,Navarra,226,4,Sunbilla
31,Navarra,227,0,Tafalla
31,Navarra,228,6,Tiebas-Muruarte de Reta
31,Navarra,229,9,Tirapu
31,Navarra,230,3,Torralba del Río
31,Navarra,231,0,Torres del Río
31,Navarra,232,5,Tudela
31,Navarra,233,1,Tulebras
31,Navarra,234,6,Ucar
31,Navarra,123,4,Uharte-Arakil
31,Navarra,235,9,Ujué
31,Navarra,236,2,Ultzama
31,Navarra,237,8,Unciti
31,Navarra,238,4,Unzué
31,Navarra,239,7,Urdazubi/Urdax
31,Navarra,240,1,Urdiain
31,Navarra,241,8,Urraul Alto
31,Navarra,242,3,Urraul Bajo
31,Navarra,244,4,Urrotz
31,Navarra,243,9,Urroz-Villa
31,Navarra,245,7,Urzainqui/Urzainki
31,Navarra,246,0,Uterga
31,Navarra,247,6,Uztárroz/Uztarroze
31,Navarra,260,6,Valle de Yerri/Deierri
31,Navarra,249,5,Valtierra
31,Navarra,251,5,Viana
31,Navarra,252,0,Vidángoz/Bidankoze
31,Navarra,254,1,Villafranca
31,Navarra,255,4,Villamayor de Monjardín
31,Navarra,257,3,Villatuerta
31,Navarra,258,9,Villava/Atarrabia
31,Navarra,261,3,Yesa
31,Navarra,262,8,Zabalza/Zabaltza
31,Navarra,073,4,Ziordia
31,Navarra,907,7,Zizur Mayor/Zizur Nagusia
31,Navarra,263,4,Zubieta
31,Navarra,264,9,Zugarramurdi
31,Navarra,265,2,Zúñiga
32,Ourense,001,3,Allariz
32,Ourense,002,8,Amoeiro
32,Ourense,003,4,"Arnoia, A"
32,Ourense,004,9,Avión
32,Ourense,005,2,Baltar
32,Ourense,006,5,Bande
32,Ourense,007,1,Baños de Molgas
32,Ourense,008,7,Barbadás
32,Ourense,009,0,"Barco de Valdeorras, O"
32,Ourense,010,4,Beade
32,Ourense,011,1,Beariz
32,Ourense,012,6,"Blancos, Os"
32,Ourense,013,2,Boborás
32,Ourense,014,7,"Bola, A"
32,Ourense,015,0,"Bolo, O"
32,Ourense,016,3,Calvos de Randín
32,Ourense,018,5,Carballeda de Avia
32,Ourense,017,9,Carballeda de Valdeorras
32,Ourense,019,8,"Carballiño, O"
32,Ourense,020,2,Cartelle
32,Ourense,022,4,Castrelo de Miño
32,Ourense,021,9,Castrelo do Val
32,Ourense,023,0,Castro Caldelas
32,Ourense,024,5,Celanova
32,Ourense,025,8,Cenlle
32,Ourense,029,6,Chandrexa de Queixa
32,Ourense,026,1,Coles
32,Ourense,027,7,Cortegada
32,Ourense,028,3,Cualedro
32,Ourense,030,0,Entrimo
32,Ourense,031,7,Esgos
32,Ourense,033,8,Gomesende
32,Ourense,034,3,"Gudiña, A"
32,Ourense,035,6,"Irixo, O"
32,Ourense,038,1,Larouco
32,Ourense,039,4,Laza
32,Ourense,040,8,Leiro
32,Ourense,041,5,Lobeira
32,Ourense,042,0,Lobios
32,Ourense,043,6,Maceda
32,Ourense,044,1,Manzaneda
32,Ourense,045,4,Maside
32,Ourense,046,7,Melón
32,Ourense,047,3,"Merca, A"
32,Ourense,048,9,"Mezquita, A"
32,Ourense,049,2,Montederramo
32,Ourense,050,5,Monterrei
32,Ourense,051,2,Muíños
32,Ourense,052,7,Nogueira de Ramuín
32,Ourense,053,3,Oímbra
32,Ourense,054,8,Ourense
32,Ourense,055,1,Paderne de Allariz
32,Ourense,056,4,Padrenda
32,Ourense,057,0,Parada de Sil
32,Ourense,058,6,"Pereiro de Aguiar, O"
32,Ourense,059,9,"Peroxa, A"
32,Ourense,060,3,Petín
32,Ourense,061,0,Piñor
32,Ourense,063,1,"Pobra de Trives, A"
32,Ourense,064,6,Pontedeva
32,Ourense,062,5,Porqueira
32,Ourense,065,9,Punxín
32,Ourense,066,2,Quintela de Leirado
32,Ourense,067,8,Rairiz de Veiga
32,Ourense,068,4,Ramirás
32,Ourense,069,7,Ribadavia
32,Ourense,071,8,Riós
32,Ourense,072,3,"Rúa, A"
32,Ourense,073,9,Rubiá
32,Ourense,074,4,San Amaro
32,Ourense,075,7,San Cibrao das Viñas
32,Ourense,076,0,San Cristovo de Cea
32,Ourense,070,1,San Xoán de Río
32,Ourense,077,6,Sandiás
32,Ourense,078,2,Sarreaus
32,Ourense,079,5,Taboadela
32,Ourense,080,9,"Teixeira, A"
32,Ourense,081,6,Toén
32,Ourense,082,1,Trasmiras
32,Ourense,083,7,"Veiga, A"
32,Ourense,084,2,Verea
32,Ourense,085,5,Verín
32,Ourense,086,8,Viana do Bolo
32,Ourense,087,4,Vilamarín
32,Ourense,088,0,Vilamartín de Valdeorras
32,Ourense,089,3,Vilar de Barrio
32,Ourense,090,7,Vilar de Santos
32,Ourense,091,4,Vilardevós
32,Ourense,092,9,Vilariño de Conso
32,Ourense,032,2,Xinzo de Limia
32,Ourense,036,9,Xunqueira de Ambía
32,Ourense,037,5,Xunqueira de Espadanedo
33,Asturias,000,0,Gijon
33,Asturias,000,0,Aviles
33,Asturias,000,0,Gijn
33,Asturias,001,9,Allande
33,Asturias,002,4,Aller
33,Asturias,003,0,Amieva
33,Asturias,004,5,Avilés
33,Asturias,005,8,Belmonte de Miranda
33,Asturias,006,1,Bimenes
33,Asturias,007,7,Boal
33,Asturias,008,3,Cabrales
33,Asturias,009,6,Cabranes
33,Asturias,010,0,Candamo
33,Asturias,012,2,Cangas de Onís
33,Asturias,011,7,Cangas del Narcea
33,Asturias,013,8,Caravia
33,Asturias,014,3,Carreño
33,Asturias,015,6,Caso
33,Asturias,016,9,Castrillón
33,Asturias,017,5,Castropol
33,Asturias,018,1,Coaña
33,Asturias,019,4,Colunga
33,Asturias,020,8,Corvera de Asturias
33,Asturias,021,5,Cudillero
33,Asturias,022,0,Degaña
33,Asturias,023,6,"Franco, El"
33,Asturias,024,1,Gijón
33,Asturias,025,4,Gozón
33,Asturias,026,7,Grado
33,Asturias,027,3,Grandas de Salime
33,Asturias,028,9,Ibias
33,Asturias,029,2,Illano
33,Asturias,030,6,Illas
33,Asturias,031,3,Langreo
33,Asturias,032,8,Laviana
33,Asturias,033,4,Lena
33,Asturias,035,2,Llanera
33,Asturias,036,5,Llanes
33,Asturias,037,1,Mieres
33,Asturias,038,7,Morcín
33,Asturias,039,0,Muros de Nalón
33,Asturias,040,4,Nava
33,Asturias,041,1,Navia
33,Asturias,042,6,Noreña
33,Asturias,043,2,Onís
33,Asturias,044,7,Oviedo
33,Asturias,045,0,Parres
33,Asturias,046,3,Peñamellera Alta
33,Asturias,047,9,Peñamellera Baja
33,Asturias,048,5,Pesoz
33,Asturias,049,8,Piloña
33,Asturias,050,1,Ponga
33,Asturias,051,8,Pravia
33,Asturias,052,3,Proaza
33,Asturias,053,9,Quirós
33,Asturias,054,4,"Regueras, Las"
33,Asturias,055,7,Ribadedeva
33,Asturias,056,0,Ribadesella
33,Asturias,057,6,Ribera de Arriba
33,Asturias,058,2,Riosa
33,Asturias,059,5,Salas
33,Asturias,061,6,San Martín de Oscos
33,Asturias,060,9,San Martín del Rey Aurelio
33,Asturias,063,7,San Tirso de Abres
33,Asturias,062,1,Santa Eulalia de Oscos
33,Asturias,064,2,Santo Adriano
33,Asturias,065,5,Sariego
33,Asturias,066,8,Siero
33,Asturias,067,4,Sobrescobio
33,Asturias,068,0,Somiedo
33,Asturias,069,3,Soto del Barco
33,Asturias,070,7,Tapia de Casariego
33,Asturias,071,4,Taramundi
33,Asturias,072,9,Teverga
33,Asturias,073,5,Tineo
33,Asturias,034,9,Valdés
33,Asturias,074,0,Vegadeo
33,Asturias,075,3,Villanueva de Oscos
33,Asturias,076,6,Villaviciosa
33,Asturias,077,2,Villayón
33,Asturias,078,8,Yernes y Tameza
34,Palencia,001,4,Abarca de Campos
34,Palencia,003,5,Abia de las Torres
34,Palencia,004,0,Aguilar de Campoo
34,Palencia,005,3,Alar del Rey
34,Palencia,006,6,Alba de Cerrato
34,Palencia,009,1,Amayuelas de Arriba
34,Palencia,010,5,Ampudia
34,Palencia,011,2,Amusco
34,Palencia,012,7,Antigüedad
34,Palencia,015,1,Arconada
34,Palencia,017,0,Astudillo
34,Palencia,018,6,Autilla del Pino
34,Palencia,019,9,Autillo de Campos
34,Palencia,020,3,Ayuela
34,Palencia,022,5,Baltanás
34,Palencia,024,6,Baquerín de Campos
34,Palencia,025,9,Bárcena de Campos
34,Palencia,027,8,Barruelo de Santullán
34,Palencia,028,4,Báscones de Ojeda
34,Palencia,029,7,Becerril de Campos
34,Palencia,031,8,Belmonte de Campos
34,Palencia,032,3,Berzosilla
34,Palencia,033,9,Boada de Campos
34,Palencia,035,7,Boadilla de Rioseco
34,Palencia,034,4,Boadilla del Camino
34,Palencia,036,0,Brañosera
34,Palencia,037,6,Buenavista de Valdavia
34,Palencia,038,2,Bustillo de la Vega
34,Palencia,039,5,Bustillo del Páramo de Carrión
34,Palencia,041,6,Calahorra de Boedo
34,Palencia,042,1,Calzada de los Molinos
34,Palencia,045,5,Capillas
34,Palencia,046,8,Cardeñosa de Volpejera
34,Palencia,047,4,Carrión de los Condes
34,Palencia,048,0,Castil de Vela
34,Palencia,049,3,Castrejón de la Peña
34,Palencia,050,6,Castrillo de Don Juan
34,Palencia,051,3,Castrillo de Onielo
34,Palencia,052,8,Castrillo de Villavega
34,Palencia,053,4,Castromocho
34,Palencia,055,2,Cervatos de la Cueza
34,Palencia,056,5,Cervera de Pisuerga
34,Palencia,057,1,Cevico de la Torre
34,Palencia,058,7,Cevico Navero
34,Palencia,059,0,Cisneros
34,Palencia,060,4,Cobos de Cerrato
34,Palencia,061,1,Collazos de Boedo
34,Palencia,062,6,Congosto de Valdavia
34,Palencia,063,2,Cordovilla la Real
34,Palencia,066,3,Cubillas de Cerrato
34,Palencia,067,9,Dehesa de Montejo
34,Palencia,068,5,Dehesa de Romanos
34,Palencia,069,8,Dueñas
34,Palencia,070,2,Espinosa de Cerrato
34,Palencia,071,9,Espinosa de Villagonzalo
34,Palencia,072,4,Frechilla
34,Palencia,073,0,Fresno del Río
34,Palencia,074,5,Frómista
34,Palencia,076,1,Fuentes de Nava
34,Palencia,077,7,Fuentes de Valdepero
34,Palencia,079,6,Grijota
34,Palencia,080,0,Guardo
34,Palencia,081,7,Guaza de Campos
34,Palencia,082,2,Hérmedes de Cerrato
34,Palencia,083,8,Herrera de Pisuerga
34,Palencia,084,3,Herrera de Valdecañas
34,Palencia,086,9,Hontoria de Cerrato
34,Palencia,087,5,Hornillos de Cerrato
34,Palencia,088,1,Husillos
34,Palencia,089,4,Itero de la Vega
34,Palencia,091,5,Lagartos
34,Palencia,092,0,Lantadilla
34,Palencia,094,1,Ledigos
34,Palencia,903,6,Loma de Ucieza
34,Palencia,096,7,Lomas
34,Palencia,098,9,Magaz de Pisuerga
34,Palencia,099,2,Manquillos
34,Palencia,100,6,Mantinos
34,Palencia,101,3,Marcilla de Campos
34,Palencia,102,8,Mazariegos
34,Palencia,103,4,Mazuecos de Valdeginate
34,Palencia,104,9,Melgar de Yuso
34,Palencia,106,5,Meneses de Campos
34,Palencia,107,1,Micieces de Ojeda
34,Palencia,108,7,Monzón de Campos
34,Palencia,109,0,Moratinos
34,Palencia,110,4,Mudá
34,Palencia,112,6,Nogal de las Huertas
34,Palencia,113,2,Olea de Boedo
34,Palencia,114,7,Olmos de Ojeda
34,Palencia,116,3,Osornillo
34,Palencia,901,5,Osorno la Mayor
34,Palencia,120,2,Palencia
34,Palencia,121,9,Palenzuela
34,Palencia,122,4,Páramo de Boedo
34,Palencia,123,0,Paredes de Nava
34,Palencia,124,5,Payo de Ojeda
34,Palencia,125,8,Pedraza de Campos
34,Palencia,126,1,Pedrosa de la Vega
34,Palencia,127,7,Perales
34,Palencia,904,1,"Pernía, La"
34,Palencia,129,6,Pino del Río
34,Palencia,130,0,Piña de Campos
34,Palencia,131,7,Población de Arroyo
34,Palencia,132,2,Población de Campos
34,Palencia,133,8,Población de Cerrato
34,Palencia,134,3,Polentinos
34,Palencia,135,6,Pomar de Valdivia
34,Palencia,136,9,Poza de la Vega
34,Palencia,137,5,Pozo de Urama
34,Palencia,139,4,Prádanos de Ojeda
34,Palencia,140,8,"Puebla de Valdavia, La"
34,Palencia,141,5,Quintana del Puente
34,Palencia,143,6,Quintanilla de Onsoña
34,Palencia,146,7,Reinoso de Cerrato
34,Palencia,147,3,Renedo de la Vega
34,Palencia,149,2,Requena de Campos
34,Palencia,151,2,Respenda de la Peña
34,Palencia,152,7,Revenga de Campos
34,Palencia,154,8,Revilla de Collazos
34,Palencia,155,1,Ribas de Campos
34,Palencia,156,4,Riberos de la Cueza
34,Palencia,157,0,Saldaña
34,Palencia,158,6,Salinas de Pisuerga
34,Palencia,159,9,San Cebrián de Campos
34,Palencia,160,3,San Cebrián de Mudá
34,Palencia,161,0,San Cristóbal de Boedo
34,Palencia,163,1,San Mamés de Campos
34,Palencia,165,9,San Román de la Cuba
34,Palencia,167,8,Santa Cecilia del Alcor
34,Palencia,168,4,Santa Cruz de Boedo
34,Palencia,169,7,Santervás de la Vega
34,Palencia,170,1,Santibáñez de Ecla
34,Palencia,171,8,Santibáñez de la Peña
34,Palencia,174,4,Santoyo
34,Palencia,175,7,"Serna, La"
34,Palencia,177,6,Soto de Cerrato
34,Palencia,176,0,Sotobañado y Priorato
34,Palencia,178,2,Tabanera de Cerrato
34,Palencia,179,5,Tabanera de Valdavia
34,Palencia,180,9,Támara de Campos
34,Palencia,181,6,Tariego de Cerrato
34,Palencia,182,1,Torquemada
34,Palencia,184,2,Torremormojón
34,Palencia,185,5,Triollo
34,Palencia,186,8,Valbuena de Pisuerga
34,Palencia,189,3,Valdeolmillos
34,Palencia,190,7,Valderrábano
34,Palencia,192,9,Valde-Ucieza
34,Palencia,196,6,Valle de Cerrato
34,Palencia,902,0,Valle del Retortillo
34,Palencia,199,1,Velilla del Río Carrión
34,Palencia,023,1,Venta de Baños
34,Palencia,201,2,Vertavillo
34,Palencia,093,6,"Vid de Ojeda, La"
34,Palencia,202,7,Villabasta de Valdavia
34,Palencia,204,8,Villacidaler
34,Palencia,205,1,Villaconancio
34,Palencia,206,4,Villada
34,Palencia,208,6,Villaeles de Valdavia
34,Palencia,210,3,Villahán
34,Palencia,211,0,Villaherreros
34,Palencia,213,1,Villalaco
34,Palencia,214,6,Villalba de Guardo
34,Palencia,215,9,Villalcázar de Sirga
34,Palencia,216,2,Villalcón
34,Palencia,217,8,Villalobón
34,Palencia,218,4,Villaluenga de la Vega
34,Palencia,220,1,Villamartín de Campos
34,Palencia,221,8,Villamediana
34,Palencia,222,3,Villameriel
34,Palencia,223,9,Villamoronta
34,Palencia,224,4,Villamuera de la Cueza
34,Palencia,225,7,Villamuriel de Cerrato
34,Palencia,227,6,Villanueva del Rebollar
34,Palencia,228,2,Villanuño de Valdavia
34,Palencia,229,5,Villaprovedo
34,Palencia,230,9,Villarmentero de Campos
34,Palencia,231,6,Villarrabé
34,Palencia,232,1,Villarramiel
34,Palencia,233,7,Villasarracino
34,Palencia,234,2,Villasila de Valdavia
34,Palencia,236,8,Villaturde
34,Palencia,237,4,Villaumbrales
34,Palencia,238,0,Villaviudas
34,Palencia,240,7,Villerías de Campos
34,Palencia,241,4,Villodre
34,Palencia,242,9,Villodrigo
34,Palencia,243,5,Villoldo
34,Palencia,245,3,Villota del Páramo
34,Palencia,246,6,Villovieco
35,Las Palmas,000,0,Las Palmas
35,Las Palmas,000,0,Las Palmas de G.C.
35,Las Palmas,000,0,Las Palmas G.C.
35,Las Palmas,000,0,Las Palmas G.C
35,Las Palmas,000,0,Las Palmas de Gran Canaria
35,Las Palmas,000,0,Las Palmas Gran Canaria
35,Las Palmas,000,0,Las Palmas de G.Canaria 
35,Las Palmas,000,0,Las Palmas de Gran  Canar
35,Las Palmas,000,0,Las Palmas de Gran Canria
35,Las Palmas,000,0,Las Palmas de Gran Canari
35,Las Palmas,000,0,Las Palmas de Gran Canar
35,Las Palmas,001,7,Agaete
35,Las Palmas,002,2,Agüimes
35,Las Palmas,020,6,"Aldea de San Nicolás, La"
35,Las Palmas,003,8,Antigua
35,Las Palmas,004,3,Arrecife
35,Las Palmas,005,6,Artenara
35,Las Palmas,006,9,Arucas
35,Las Palmas,007,5,Betancuria
35,Las Palmas,008,1,Firgas
35,Las Palmas,009,4,Gáldar
35,Las Palmas,010,8,Haría
35,Las Palmas,011,5,Ingenio
35,Las Palmas,012,0,Mogán
35,Las Palmas,013,6,Moya
35,Las Palmas,014,1,"Oliva, La"
35,Las Palmas,015,4,Pájara
35,Las Palmas,016,7,"Palmas de Gran Canaria, Las"
35,Las Palmas,017,3,Puerto del Rosario
35,Las Palmas,018,9,San Bartolomé
35,Las Palmas,019,2,San Bartolomé de Tirajana
35,Las Palmas,021,3,Santa Brígida
35,Las Palmas,022,8,Santa Lucía de Tirajana
35,Las Palmas,023,4,Santa María de Guía de Gran Canaria
35,Las Palmas,024,9,Teguise
35,Las Palmas,025,2,Tejeda
35,Las Palmas,026,5,Telde
35,Las Palmas,027,1,Teror
35,Las Palmas,028,7,Tías
35,Las Palmas,029,0,Tinajo
35,Las Palmas,030,4,Tuineje
35,Las Palmas,032,6,Valleseco
35,Las Palmas,031,1,Valsequillo de Gran Canaria
35,Las Palmas,033,2,Vega de San Mateo
35,Las Palmas,034,7,Yaiza
36,Pontevedra,020,9,Agolada
36,Pontevedra,001,0,Arbo
36,Pontevedra,003,1,Baiona
36,Pontevedra,002,5,Barro
36,Pontevedra,004,6,Bueu
36,Pontevedra,005,9,Caldas de Reis
36,Pontevedra,006,2,Cambados
36,Pontevedra,007,8,Campo Lameiro
36,Pontevedra,008,4,Cangas
36,Pontevedra,009,7,"Cañiza, A"
36,Pontevedra,010,1,Catoira
36,Pontevedra,011,8,Cerdedo
36,Pontevedra,012,3,Cotobade
36,Pontevedra,013,9,Covelo
36,Pontevedra,014,4,Crecente
36,Pontevedra,015,7,Cuntis
36,Pontevedra,016,0,Dozón
36,Pontevedra,017,6,"Estrada, A"
36,Pontevedra,018,2,Forcarei
36,Pontevedra,019,5,Fornelos de Montes
36,Pontevedra,021,6,Gondomar
36,Pontevedra,022,1,"Grove, O"
36,Pontevedra,023,7,"Guarda, A"
36,Pontevedra,901,1,"Illa de Arousa, A"
36,Pontevedra,024,2,Lalín
36,Pontevedra,025,5,"Lama, A"
36,Pontevedra,026,8,Marín
36,Pontevedra,027,4,Meaño
36,Pontevedra,028,0,Meis
36,Pontevedra,029,3,Moaña
36,Pontevedra,030,7,Mondariz
36,Pontevedra,031,4,Mondariz-Balneario
36,Pontevedra,032,9,Moraña
36,Pontevedra,033,5,Mos
36,Pontevedra,034,0,"Neves, As"
36,Pontevedra,035,3,Nigrán
36,Pontevedra,036,6,Oia
36,Pontevedra,037,2,Pazos de Borbén
36,Pontevedra,041,2,Poio
36,Pontevedra,043,3,Ponte Caldelas
36,Pontevedra,042,7,Ponteareas
36,Pontevedra,044,8,Pontecesures
36,Pontevedra,038,8,Pontevedra
36,Pontevedra,039,1,"Porriño, O"
36,Pontevedra,040,5,Portas
36,Pontevedra,045,1,Redondela
36,Pontevedra,046,4,Ribadumia
36,Pontevedra,047,0,Rodeiro
36,Pontevedra,048,6,"Rosal, O"
36,Pontevedra,049,9,Salceda de Caselas
36,Pontevedra,050,2,Salvaterra de Miño
36,Pontevedra,051,9,Sanxenxo
36,Pontevedra,052,4,Silleda
36,Pontevedra,053,0,Soutomaior
36,Pontevedra,054,5,Tomiño
36,Pontevedra,055,8,Tui
36,Pontevedra,056,1,Valga
36,Pontevedra,057,7,Vigo
36,Pontevedra,059,6,Vila de Cruces
36,Pontevedra,058,3,Vilaboa
36,Pontevedra,060,0,Vilagarcía de Arousa
36,Pontevedra,061,7,Vilanova de Arousa
37,Salamanca,001,6,Abusejo
37,Salamanca,002,1,Agallas
37,Salamanca,003,7,Ahigal de los Aceiteros
37,Salamanca,004,2,Ahigal de Villarino
37,Salamanca,005,5,"Alameda de Gardón, La"
37,Salamanca,006,8,"Alamedilla, La"
37,Salamanca,007,4,Alaraz
37,Salamanca,008,0,Alba de Tormes
37,Salamanca,009,3,Alba de Yeltes
37,Salamanca,010,7,"Alberca, La"
37,Salamanca,011,4,"Alberguería de Argañán, La"
37,Salamanca,012,9,Alconada
37,Salamanca,015,3,Aldea del Obispo
37,Salamanca,013,5,Aldeacipreste
37,Salamanca,014,0,Aldeadávila de la Ribera
37,Salamanca,016,6,Aldealengua
37,Salamanca,017,2,Aldeanueva de Figueroa
37,Salamanca,018,8,Aldeanueva de la Sierra
37,Salamanca,019,1,Aldearrodrigo
37,Salamanca,020,5,Aldearrubia
37,Salamanca,021,2,Aldeaseca de Alba
37,Salamanca,022,7,Aldeaseca de la Frontera
37,Salamanca,023,3,Aldeatejada
37,Salamanca,024,8,Aldeavieja de Tormes
37,Salamanca,025,1,Aldehuela de la Bóveda
37,Salamanca,026,4,Aldehuela de Yeltes
37,Salamanca,027,0,Almenara de Tormes
37,Salamanca,028,6,Almendra
37,Salamanca,029,9,Anaya de Alba
37,Salamanca,030,3,Añover de Tormes
37,Salamanca,031,0,Arabayona de Mógica
37,Salamanca,032,5,Arapiles
37,Salamanca,033,1,Arcediano
37,Salamanca,034,6,"Arco, El"
37,Salamanca,035,9,Armenteros
37,Salamanca,037,8,"Atalaya, La"
37,Salamanca,038,4,Babilafuente
37,Salamanca,039,7,Bañobárez
37,Salamanca,040,1,Barbadillo
37,Salamanca,041,8,Barbalos
37,Salamanca,042,3,Barceo
37,Salamanca,044,4,Barruecopardo
37,Salamanca,045,7,"Bastida, La"
37,Salamanca,046,0,Béjar
37,Salamanca,047,6,Beleña
37,Salamanca,049,5,Bermellar
37,Salamanca,050,8,Berrocal de Huebra
37,Salamanca,051,5,Berrocal de Salvatierra
37,Salamanca,052,0,Boada
37,Salamanca,054,1,"Bodón, El"
37,Salamanca,055,4,Bogajo
37,Salamanca,056,7,"Bouza, La"
37,Salamanca,057,3,Bóveda del Río Almar
37,Salamanca,058,9,Brincones
37,Salamanca,059,2,Buenamadre
37,Salamanca,060,6,Buenavista
37,Salamanca,061,3,"Cabaco, El"
37,Salamanca,063,4,"Cabeza de Béjar, La"
37,Salamanca,065,2,Cabeza del Caballo
37,Salamanca,062,8,Cabezabellosa de la Calzada
37,Salamanca,067,1,Cabrerizos
37,Salamanca,068,7,Cabrillas
37,Salamanca,069,0,Calvarrasa de Abajo
37,Salamanca,070,4,Calvarrasa de Arriba
37,Salamanca,071,1,"Calzada de Béjar, La"
37,Salamanca,072,6,Calzada de Don Diego
37,Salamanca,073,2,Calzada de Valdunciel
37,Salamanca,074,7,Campillo de Azaba
37,Salamanca,077,9,"Campo de Peñaranda, El"
37,Salamanca,078,5,Candelario
37,Salamanca,079,8,Canillas de Abajo
37,Salamanca,080,2,Cantagallo
37,Salamanca,081,9,Cantalapiedra
37,Salamanca,082,4,Cantalpino
37,Salamanca,083,0,Cantaracillo
37,Salamanca,085,8,Carbajosa de la Sagrada
37,Salamanca,086,1,Carpio de Azaba
37,Salamanca,087,7,Carrascal de Barregas
37,Salamanca,088,3,Carrascal del Obispo
37,Salamanca,089,6,Casafranca
37,Salamanca,090,0,"Casas del Conde, Las"
37,Salamanca,091,7,Casillas de Flores
37,Salamanca,092,2,Castellanos de Moriscos
37,Salamanca,185,7,Castellanos de Villiquera
37,Salamanca,096,9,Castillejo de Martín Viejo
37,Salamanca,097,5,Castraz
37,Salamanca,098,1,Cepeda
37,Salamanca,099,4,Cereceda de la Sierra
37,Salamanca,100,8,Cerezal de Peñahorcada
37,Salamanca,101,5,Cerralbo
37,Salamanca,102,0,"Cerro, El"
37,Salamanca,103,6,Cespedosa de Tormes
37,Salamanca,114,9,Chagarcía Medianero
37,Salamanca,104,1,Cilleros de la Bastida
37,Salamanca,106,7,Cipérez
37,Salamanca,107,3,Ciudad Rodrigo
37,Salamanca,108,9,Coca de Alba
37,Salamanca,109,2,Colmenar de Montemayor
37,Salamanca,110,6,Cordovilla
37,Salamanca,112,8,Cristóbal
37,Salamanca,113,4,"Cubo de Don Sancho, El"
37,Salamanca,115,2,Dios le Guarde
37,Salamanca,116,5,Doñinos de Ledesma
37,Salamanca,117,1,Doñinos de Salamanca
37,Salamanca,118,7,Ejeme
37,Salamanca,119,0,"Encina, La"
37,Salamanca,120,4,Encina de San Silvestre
37,Salamanca,121,1,Encinas de Abajo
37,Salamanca,122,6,Encinas de Arriba
37,Salamanca,123,2,Encinasola de los Comendadores
37,Salamanca,124,7,Endrinal
37,Salamanca,125,0,Escurial de la Sierra
37,Salamanca,126,3,Espadaña
37,Salamanca,127,9,Espeja
37,Salamanca,128,5,Espino de la Orbada
37,Salamanca,129,8,Florida de Liébana
37,Salamanca,130,2,Forfoleda
37,Salamanca,131,9,Frades de la Sierra
37,Salamanca,132,4,"Fregeneda, La"
37,Salamanca,133,0,Fresnedoso
37,Salamanca,134,5,Fresno Alhándiga
37,Salamanca,135,8,"Fuente de San Esteban, La"
37,Salamanca,136,1,Fuenteguinaldo
37,Salamanca,137,7,Fuenteliante
37,Salamanca,138,3,Fuenterroble de Salvatierra
37,Salamanca,139,6,Fuentes de Béjar
37,Salamanca,140,0,Fuentes de Oñoro
37,Salamanca,141,7,Gajates
37,Salamanca,142,2,Galindo y Perahuy
37,Salamanca,143,8,Galinduste
37,Salamanca,144,3,Galisancho
37,Salamanca,145,6,Gallegos de Argañán
37,Salamanca,146,9,Gallegos de Solmirón
37,Salamanca,147,5,Garcibuey
37,Salamanca,148,1,Garcihernández
37,Salamanca,149,4,Garcirrey
37,Salamanca,150,7,Gejuelo del Barro
37,Salamanca,151,4,Golpejas
37,Salamanca,152,9,Gomecello
37,Salamanca,154,0,Guadramiro
37,Salamanca,155,3,Guijo de Ávila
37,Salamanca,156,6,Guijuelo
37,Salamanca,157,2,Herguijuela de Ciudad Rodrigo
37,Salamanca,158,8,Herguijuela de la Sierra
37,Salamanca,159,1,Herguijuela del Campo
37,Salamanca,160,5,Hinojosa de Duero
37,Salamanca,161,2,Horcajo de Montemayor
37,Salamanca,162,7,Horcajo Medianero
37,Salamanca,163,3,"Hoya, La"
37,Salamanca,164,8,Huerta
37,Salamanca,165,1,Iruelos
37,Salamanca,166,4,Ituero de Azaba
37,Salamanca,167,0,Juzbado
37,Salamanca,168,6,Lagunilla
37,Salamanca,169,9,Larrodrigo
37,Salamanca,170,3,Ledesma
37,Salamanca,171,0,Ledrada
37,Salamanca,172,5,Linares de Riofrío
37,Salamanca,173,1,Lumbrales
37,Salamanca,175,9,Machacón
37,Salamanca,174,6,Macotera
37,Salamanca,176,2,Madroñal
37,Salamanca,177,8,"Maíllo, El"
37,Salamanca,178,4,Malpartida
37,Salamanca,179,7,Mancera de Abajo
37,Salamanca,180,1,"Manzano, El"
37,Salamanca,181,8,Martiago
37,Salamanca,183,9,Martín de Yeltes
37,Salamanca,182,3,Martinamor
37,Salamanca,184,4,Masueco
37,Salamanca,186,0,"Mata de Ledesma, La"
37,Salamanca,187,6,Matilla de los Caños del Río
37,Salamanca,188,2,"Maya, La"
37,Salamanca,189,5,Membribe de la Sierra
37,Salamanca,190,9,Mieza
37,Salamanca,191,6,"Milano, El"
37,Salamanca,192,1,Miranda de Azán
37,Salamanca,193,7,Miranda del Castañar
37,Salamanca,194,2,Mogarraz
37,Salamanca,195,5,Molinillo
37,Salamanca,196,8,Monforte de la Sierra
37,Salamanca,197,4,Monleón
37,Salamanca,198,0,Monleras
37,Salamanca,199,3,Monsagro
37,Salamanca,200,7,Montejo
37,Salamanca,201,4,Montemayor del Río
37,Salamanca,202,9,Monterrubio de Armuña
37,Salamanca,203,5,Monterrubio de la Sierra
37,Salamanca,204,0,Morasverdes
37,Salamanca,205,3,Morille
37,Salamanca,206,6,Moríñigo
37,Salamanca,207,2,Moriscos
37,Salamanca,208,8,Moronta
37,Salamanca,209,1,Mozárbez
37,Salamanca,211,2,Narros de Matalayegua
37,Salamanca,213,3,Nava de Béjar
37,Salamanca,214,8,Nava de Francia
37,Salamanca,215,1,Nava de Sotrobal
37,Salamanca,212,7,Navacarros
37,Salamanca,216,4,Navales
37,Salamanca,217,0,Navalmoral de Béjar
37,Salamanca,218,6,Navamorales
37,Salamanca,219,9,Navarredonda de la Rinconada
37,Salamanca,221,0,Navasfrías
37,Salamanca,222,5,Negrilla de Palencia
37,Salamanca,223,1,Olmedo de Camaces
37,Salamanca,224,6,"Orbada, La"
37,Salamanca,225,9,Pajares de la Laguna
37,Salamanca,226,2,Palacios del Arzobispo
37,Salamanca,228,4,Palaciosrubios
37,Salamanca,229,7,Palencia de Negrilla
37,Salamanca,230,1,Parada de Arriba
37,Salamanca,231,8,Parada de Rubiales
37,Salamanca,232,3,Paradinas de San Juan
37,Salamanca,233,9,Pastores
37,Salamanca,234,4,"Payo, El"
37,Salamanca,235,7,Pedraza de Alba
37,Salamanca,236,0,Pedrosillo de Alba
37,Salamanca,237,6,Pedrosillo de los Aires
37,Salamanca,238,2,Pedrosillo el Ralo
37,Salamanca,239,5,"Pedroso de la Armuña, El"
37,Salamanca,240,9,Pelabravo
37,Salamanca,241,6,Pelarrodríguez
37,Salamanca,242,1,Pelayos
37,Salamanca,243,7,"Peña, La"
37,Salamanca,244,2,Peñacaballera
37,Salamanca,245,5,Peñaparda
37,Salamanca,246,8,Peñaranda de Bracamonte
37,Salamanca,247,4,Peñarandilla
37,Salamanca,248,0,Peralejos de Abajo
37,Salamanca,249,3,Peralejos de Arriba
37,Salamanca,250,6,Pereña de la Ribera
37,Salamanca,251,3,Peromingo
37,Salamanca,252,8,Pinedas
37,Salamanca,253,4,"Pino de Tormes, El"
37,Salamanca,254,9,Pitiegua
37,Salamanca,255,2,Pizarral
37,Salamanca,256,5,Poveda de las Cintas
37,Salamanca,257,1,Pozos de Hinojo
37,Salamanca,258,7,Puebla de Azaba
37,Salamanca,259,0,Puebla de San Medel
37,Salamanca,260,4,Puebla de Yeltes
37,Salamanca,261,1,Puente del Congosto
37,Salamanca,262,6,Puertas
37,Salamanca,263,2,Puerto de Béjar
37,Salamanca,264,7,Puerto Seguro
37,Salamanca,265,0,Rágama
37,Salamanca,266,3,"Redonda, La"
37,Salamanca,267,9,Retortillo
37,Salamanca,268,5,"Rinconada de la Sierra, La"
37,Salamanca,269,8,Robleda
37,Salamanca,270,2,Robliza de Cojos
37,Salamanca,271,9,Rollán
37,Salamanca,272,4,Saelices el Chico
37,Salamanca,273,0,"Sagrada, La"
37,Salamanca,303,4,"Sahugo, El"
37,Salamanca,274,5,Salamanca
37,Salamanca,275,8,Saldeana
37,Salamanca,276,1,Salmoral
37,Salamanca,277,7,Salvatierra de Tormes
37,Salamanca,278,3,San Cristóbal de la Cuesta
37,Salamanca,284,3,San Esteban de la Sierra
37,Salamanca,285,6,San Felices de los Gallegos
37,Salamanca,286,9,San Martín del Castañar
37,Salamanca,287,5,San Miguel de Valero
37,Salamanca,036,2,San Miguel del Robledo
37,Salamanca,288,1,San Morales
37,Salamanca,289,4,San Muñoz
37,Salamanca,291,5,San Pedro de Rozados
37,Salamanca,290,8,San Pedro del Valle
37,Salamanca,292,0,San Pelayo de Guareña
37,Salamanca,280,0,Sanchón de la Ribera
37,Salamanca,281,7,Sanchón de la Sagrada
37,Salamanca,282,2,Sanchotello
37,Salamanca,279,6,Sancti-Spíritus
37,Salamanca,283,8,Sando
37,Salamanca,293,6,Santa María de Sando
37,Salamanca,294,1,Santa Marta de Tormes
37,Salamanca,296,7,Santiago de la Puebla
37,Salamanca,297,3,Santibáñez de Béjar
37,Salamanca,298,9,Santibáñez de la Sierra
37,Salamanca,299,2,Santiz
37,Salamanca,300,6,"Santos, Los"
37,Salamanca,301,3,Sardón de los Frailes
37,Salamanca,302,8,Saucelle
37,Salamanca,304,9,Sepulcro-Hilario
37,Salamanca,305,2,Sequeros
37,Salamanca,306,5,Serradilla del Arroyo
37,Salamanca,307,1,Serradilla del Llano
37,Salamanca,309,0,"Sierpe, La"
37,Salamanca,310,4,Sieteiglesias de Tormes
37,Salamanca,311,1,Sobradillo
37,Salamanca,312,6,Sorihuela
37,Salamanca,313,2,Sotoserrano
37,Salamanca,314,7,Tabera de Abajo
37,Salamanca,315,0,"Tala, La"
37,Salamanca,316,3,Tamames
37,Salamanca,317,9,Tarazona de Guareña
37,Salamanca,318,5,Tardáguila
37,Salamanca,319,8,"Tejado, El"
37,Salamanca,320,2,Tejeda y Segoyuela
37,Salamanca,321,9,Tenebrón
37,Salamanca,322,4,Terradillos
37,Salamanca,323,0,Topas
37,Salamanca,324,5,Tordillos
37,Salamanca,325,8,"Tornadizo, El"
37,Salamanca,327,7,Torresmenudas
37,Salamanca,328,3,Trabanca
37,Salamanca,329,6,Tremedal de Tormes
37,Salamanca,330,0,Valdecarros
37,Salamanca,331,7,Valdefuentes de Sangusín
37,Salamanca,332,2,Valdehijaderos
37,Salamanca,333,8,Valdelacasa
37,Salamanca,334,3,Valdelageve
37,Salamanca,335,6,Valdelosa
37,Salamanca,336,9,Valdemierque
37,Salamanca,337,5,Valderrodrigo
37,Salamanca,338,1,Valdunciel
37,Salamanca,339,4,Valero
37,Salamanca,343,6,Vallejera de Riofrío
37,Salamanca,340,8,Valsalabroso
37,Salamanca,341,5,Valverde de Valdelacasa
37,Salamanca,342,0,Valverdón
37,Salamanca,344,1,Vecinos
37,Salamanca,345,4,Vega de Tirados
37,Salamanca,346,7,"Veguillas, Las"
37,Salamanca,347,3,"Vellés, La"
37,Salamanca,348,9,Ventosa del Río Almar
37,Salamanca,349,2,"Vídola, La"
37,Salamanca,351,2,Villaflores
37,Salamanca,352,7,Villagonzalo de Tormes
37,Salamanca,353,3,Villalba de los Llanos
37,Salamanca,354,8,Villamayor
37,Salamanca,355,1,Villanueva del Conde
37,Salamanca,356,4,Villar de Argañán
37,Salamanca,357,0,Villar de Ciervo
37,Salamanca,358,6,Villar de Gallimazo
37,Salamanca,359,9,Villar de la Yegua
37,Salamanca,360,3,Villar de Peralonso
37,Salamanca,361,0,Villar de Samaniego
37,Salamanca,362,5,Villares de la Reina
37,Salamanca,363,1,Villares de Yeltes
37,Salamanca,364,6,Villarino de los Aires
37,Salamanca,365,9,Villarmayor
37,Salamanca,366,2,Villarmuerto
37,Salamanca,367,8,Villasbuenas
37,Salamanca,368,4,Villasdardo
37,Salamanca,369,7,Villaseco de los Gamitos
37,Salamanca,370,1,Villaseco de los Reyes
37,Salamanca,371,8,Villasrubias
37,Salamanca,372,3,Villaverde de Guareña
37,Salamanca,373,9,Villavieja de Yeltes
37,Salamanca,374,4,Villoria
37,Salamanca,375,7,Villoruela
37,Salamanca,350,5,Vilvestre
37,Salamanca,376,0,Vitigudino
37,Salamanca,377,6,Yecla de Yeltes
37,Salamanca,378,2,Zamarra
37,Salamanca,379,5,Zamayón
37,Salamanca,380,9,Zarapicos
37,Salamanca,381,6,"Zarza de Pumareda, La"
37,Salamanca,382,1,Zorita de la Frontera
38,Santa Cruz de Tenerife,000,0,S/C de Tenerife
38,Santa Cruz de Tenerife,001,2,Adeje
38,Santa Cruz de Tenerife,002,7,Agulo
38,Santa Cruz de Tenerife,003,3,Alajeró
38,Santa Cruz de Tenerife,004,8,Arafo
38,Santa Cruz de Tenerife,005,1,Arico
38,Santa Cruz de Tenerife,006,4,Arona
38,Santa Cruz de Tenerife,007,0,Barlovento
38,Santa Cruz de Tenerife,008,6,Breña Alta
38,Santa Cruz de Tenerife,009,9,Breña Baja
38,Santa Cruz de Tenerife,010,3,Buenavista del Norte
38,Santa Cruz de Tenerife,011,0,Candelaria
38,Santa Cruz de Tenerife,012,5,Fasnia
38,Santa Cruz de Tenerife,013,1,Frontera
38,Santa Cruz de Tenerife,014,6,Fuencaliente de la Palma
38,Santa Cruz de Tenerife,015,9,Garachico
38,Santa Cruz de Tenerife,016,2,Garafía
38,Santa Cruz de Tenerife,017,8,Granadilla de Abona
38,Santa Cruz de Tenerife,018,4,"Guancha, La"
38,Santa Cruz de Tenerife,019,7,Guía de Isora
38,Santa Cruz de Tenerife,020,1,Güímar
38,Santa Cruz de Tenerife,021,8,Hermigua
38,Santa Cruz de Tenerife,022,3,Icod de los Vinos
38,Santa Cruz de Tenerife,024,4,"Llanos de Aridane, Los"
38,Santa Cruz de Tenerife,025,7,"Matanza de Acentejo, La"
38,Santa Cruz de Tenerife,026,0,"Orotava, La"
38,Santa Cruz de Tenerife,027,6,"Paso, El"
38,Santa Cruz de Tenerife,901,3,"Pinar de El Hierro, El"
38,Santa Cruz de Tenerife,028,2,Puerto de la Cruz
38,Santa Cruz de Tenerife,029,5,Puntagorda
38,Santa Cruz de Tenerife,030,9,Puntallana
38,Santa Cruz de Tenerife,031,6,"Realejos, Los"
38,Santa Cruz de Tenerife,032,1,"Rosario, El"
38,Santa Cruz de Tenerife,033,7,San Andrés y Sauces
38,Santa Cruz de Tenerife,023,9,San Cristóbal de La Laguna
38,Santa Cruz de Tenerife,034,2,San Juan de la Rambla
38,Santa Cruz de Tenerife,035,5,San Miguel de Abona
38,Santa Cruz de Tenerife,036,8,San Sebastián de la Gomera
38,Santa Cruz de Tenerife,037,4,Santa Cruz de la Palma
38,Santa Cruz de Tenerife,038,0,Santa Cruz de Tenerife
38,Santa Cruz de Tenerife,039,3,Santa Úrsula
38,Santa Cruz de Tenerife,040,7,Santiago del Teide
38,Santa Cruz de Tenerife,041,4,"Sauzal, El"
38,Santa Cruz de Tenerife,042,9,"Silos, Los"
38,Santa Cruz de Tenerife,043,5,Tacoronte
38,Santa Cruz de Tenerife,044,0,"Tanque, El"
38,Santa Cruz de Tenerife,045,3,Tazacorte
38,Santa Cruz de Tenerife,046,6,Tegueste
38,Santa Cruz de Tenerife,047,2,Tijarafe
38,Santa Cruz de Tenerife,049,1,Valle Gran Rey
38,Santa Cruz de Tenerife,050,4,Vallehermoso
38,Santa Cruz de Tenerife,048,8,Valverde
38,Santa Cruz de Tenerife,051,1,"Victoria de Acentejo, La"
38,Santa Cruz de Tenerife,052,6,Vilaflor
38,Santa Cruz de Tenerife,053,2,Villa de Mazo
39,Cantabria,001,5,Alfoz de Lloredo
39,Cantabria,002,0,Ampuero
39,Cantabria,003,6,Anievas
39,Cantabria,004,1,Arenas de Iguña
39,Cantabria,005,4,Argoños
39,Cantabria,006,7,Arnuero
39,Cantabria,007,3,Arredondo
39,Cantabria,008,9,"Astillero, El"
39,Cantabria,009,2,Bárcena de Cicero
39,Cantabria,010,6,Bárcena de Pie de Concha
39,Cantabria,011,3,Bareyo
39,Cantabria,012,8,Cabezón de la Sal
39,Cantabria,013,4,Cabezón de Liébana
39,Cantabria,014,9,Cabuérniga
39,Cantabria,015,2,Camaleño
39,Cantabria,016,5,Camargo
39,Cantabria,027,9,Campoo de Enmedio
39,Cantabria,017,1,Campoo de Yuso
39,Cantabria,018,7,Cartes
39,Cantabria,019,0,Castañeda
39,Cantabria,020,4,Castro-Urdiales
39,Cantabria,021,1,Cieza
39,Cantabria,022,6,Cillorigo de Liébana
39,Cantabria,023,2,Colindres
39,Cantabria,024,7,Comillas
39,Cantabria,025,0,"Corrales de Buelna, Los"
39,Cantabria,026,3,Corvera de Toranzo
39,Cantabria,028,5,Entrambasaguas
39,Cantabria,029,8,Escalante
39,Cantabria,030,2,Guriezo
39,Cantabria,031,9,Hazas de Cesto
39,Cantabria,032,4,Hermandad de Campoo de Suso
39,Cantabria,033,0,Herrerías
39,Cantabria,034,5,Lamasón
39,Cantabria,035,8,Laredo
39,Cantabria,036,1,Liendo
39,Cantabria,037,7,Liérganes
39,Cantabria,038,3,Limpias
39,Cantabria,039,6,Luena
39,Cantabria,040,0,Marina de Cudeyo
39,Cantabria,041,7,Mazcuerras
39,Cantabria,042,2,Medio Cudeyo
39,Cantabria,043,8,Meruelo
39,Cantabria,044,3,Miengo
39,Cantabria,045,6,Miera
39,Cantabria,046,9,Molledo
39,Cantabria,047,5,Noja
39,Cantabria,048,1,Penagos
39,Cantabria,049,4,Peñarrubia
39,Cantabria,050,7,Pesaguero
39,Cantabria,051,4,Pesquera
39,Cantabria,052,9,Piélagos
39,Cantabria,053,5,Polaciones
39,Cantabria,054,0,Polanco
39,Cantabria,055,3,Potes
39,Cantabria,056,6,Puente Viesgo
39,Cantabria,057,2,Ramales de la Victoria
39,Cantabria,058,8,Rasines
39,Cantabria,059,1,Reinosa
39,Cantabria,060,5,Reocín
39,Cantabria,061,2,Ribamontán al Mar
39,Cantabria,062,7,Ribamontán al Monte
39,Cantabria,063,3,Rionansa
39,Cantabria,064,8,Riotuerto
39,Cantabria,065,1,"Rozas de Valdearroyo, Las"
39,Cantabria,066,4,Ruente
39,Cantabria,067,0,Ruesga
39,Cantabria,068,6,Ruiloba
39,Cantabria,069,9,San Felices de Buelna
39,Cantabria,070,3,San Miguel de Aguayo
39,Cantabria,071,0,San Pedro del Romeral
39,Cantabria,072,5,San Roque de Riomiera
39,Cantabria,080,1,San Vicente de la Barquera
39,Cantabria,073,1,Santa Cruz de Bezana
39,Cantabria,074,6,Santa María de Cayón
39,Cantabria,075,9,Santander
39,Cantabria,076,2,Santillana del Mar
39,Cantabria,077,8,Santiurde de Reinosa
39,Cantabria,078,4,Santiurde de Toranzo
39,Cantabria,079,7,Santoña
39,Cantabria,081,8,Saro
39,Cantabria,082,3,Selaya
39,Cantabria,083,9,Soba
39,Cantabria,084,4,Solórzano
39,Cantabria,085,7,Suances
39,Cantabria,086,0,"Tojos, Los"
39,Cantabria,087,6,Torrelavega
39,Cantabria,088,2,Tresviso
39,Cantabria,089,5,Tudanca
39,Cantabria,090,9,Udías
39,Cantabria,095,5,Val de San Vicente
39,Cantabria,091,6,Valdáliga
39,Cantabria,092,1,Valdeolea
39,Cantabria,093,7,Valdeprado del Río
39,Cantabria,094,2,Valderredible
39,Cantabria,101,4,Valle de Villaverde
39,Cantabria,096,8,Vega de Liébana
39,Cantabria,097,4,Vega de Pas
39,Cantabria,098,0,Villacarriedo
39,Cantabria,099,3,Villaescusa
39,Cantabria,100,7,Villafufre
39,Cantabria,102,9,Voto
40,Segovia,001,9,Abades
40,Segovia,002,4,Adrada de Pirón
40,Segovia,003,0,Adrados
40,Segovia,004,5,Aguilafuente
40,Segovia,005,8,Alconada de Maderuelo
40,Segovia,012,2,Aldea Real
40,Segovia,006,1,Aldealcorvo
40,Segovia,007,7,Aldealengua de Pedraza
40,Segovia,008,3,Aldealengua de Santa María
40,Segovia,009,6,Aldeanueva de la Serrezuela
40,Segovia,010,0,Aldeanueva del Codonal
40,Segovia,013,8,Aldeasoña
40,Segovia,014,3,Aldehorno
40,Segovia,015,6,Aldehuela del Codonal
40,Segovia,016,9,Aldeonte
40,Segovia,017,5,Anaya
40,Segovia,018,1,Añe
40,Segovia,019,4,Arahuetes
40,Segovia,020,8,Arcones
40,Segovia,021,5,Arevalillo de Cega
40,Segovia,022,0,Armuña
40,Segovia,024,1,Ayllón
40,Segovia,025,4,Barbolla
40,Segovia,026,7,Basardilla
40,Segovia,028,9,Bercial
40,Segovia,029,2,Bercimuel
40,Segovia,030,6,Bernardos
40,Segovia,031,3,Bernuy de Porreros
40,Segovia,032,8,Boceguillas
40,Segovia,033,4,Brieva
40,Segovia,034,9,Caballar
40,Segovia,035,2,Cabañas de Polendos
40,Segovia,036,5,Cabezuela
40,Segovia,037,1,Calabazas de Fuentidueña
40,Segovia,039,0,Campo de San Pedro
40,Segovia,040,4,Cantalejo
40,Segovia,041,1,Cantimpalos
40,Segovia,043,2,Carbonero el Mayor
40,Segovia,044,7,Carrascal del Río
40,Segovia,045,0,Casla
40,Segovia,046,3,Castillejo de Mesleón
40,Segovia,047,9,Castro de Fuentidueña
40,Segovia,048,5,Castrojimeno
40,Segovia,049,8,Castroserna de Abajo
40,Segovia,051,8,Castroserracín
40,Segovia,052,3,Cedillo de la Torre
40,Segovia,053,9,Cerezo de Abajo
40,Segovia,054,4,Cerezo de Arriba
40,Segovia,065,5,Chañe
40,Segovia,055,7,Cilleruelo de San Mamés
40,Segovia,056,0,Cobos de Fuentidueña
40,Segovia,057,6,Coca
40,Segovia,058,2,Codorniz
40,Segovia,059,5,Collado Hermoso
40,Segovia,060,9,Condado de Castilnovo
40,Segovia,061,6,Corral de Ayllón
40,Segovia,902,5,Cozuelos de Fuentidueña
40,Segovia,062,1,Cubillo
40,Segovia,063,7,Cuéllar
40,Segovia,905,9,Cuevas de Provanco
40,Segovia,068,0,Domingo García
40,Segovia,069,3,Donhierro
40,Segovia,070,7,Duruelo
40,Segovia,071,4,Encinas
40,Segovia,072,9,Encinillas
40,Segovia,073,5,Escalona del Prado
40,Segovia,074,0,Escarabajosa de Cabezas
40,Segovia,075,3,Escobar de Polendos
40,Segovia,076,6,"Espinar, El"
40,Segovia,077,2,Espirdo
40,Segovia,078,8,Fresneda de Cuéllar
40,Segovia,079,1,Fresno de Cantespino
40,Segovia,080,5,Fresno de la Fuente
40,Segovia,081,2,Frumales
40,Segovia,082,7,Fuente de Santa Cruz
40,Segovia,083,3,Fuente el Olmo de Fuentidueña
40,Segovia,084,8,Fuente el Olmo de Íscar
40,Segovia,086,4,Fuentepelayo
40,Segovia,087,0,Fuentepiñel
40,Segovia,088,6,Fuenterrebollo
40,Segovia,089,9,Fuentesaúco de Fuentidueña
40,Segovia,091,0,Fuentesoto
40,Segovia,092,5,Fuentidueña
40,Segovia,093,1,Gallegos
40,Segovia,094,6,Garcillán
40,Segovia,095,9,Gomezserracín
40,Segovia,097,8,Grajera
40,Segovia,099,7,Honrubia de la Cuesta
40,Segovia,100,1,Hontalbilla
40,Segovia,101,8,Hontanares de Eresma
40,Segovia,103,9,"Huertos, Los"
40,Segovia,104,4,Ituero y Lama
40,Segovia,105,7,Juarros de Riomoros
40,Segovia,106,0,Juarros de Voltoya
40,Segovia,107,6,Labajos
40,Segovia,108,2,Laguna de Contreras
40,Segovia,109,5,Languilla
40,Segovia,110,9,Lastras de Cuéllar
40,Segovia,111,6,Lastras del Pozo
40,Segovia,112,1,"Lastrilla, La"
40,Segovia,113,7,"Losa, La"
40,Segovia,115,5,Maderuelo
40,Segovia,903,1,Marazoleja
40,Segovia,118,0,Marazuela
40,Segovia,119,3,Martín Miguel
40,Segovia,120,7,Martín Muñoz de la Dehesa
40,Segovia,121,4,Martín Muñoz de las Posadas
40,Segovia,122,9,Marugán
40,Segovia,124,0,Mata de Cuéllar
40,Segovia,123,5,Matabuena
40,Segovia,125,3,"Matilla, La"
40,Segovia,126,6,Melque de Cercos
40,Segovia,127,2,Membibre de la Hoz
40,Segovia,128,8,Migueláñez
40,Segovia,129,1,Montejo de Arévalo
40,Segovia,130,5,Montejo de la Vega de la Serrezuela
40,Segovia,131,2,Monterrubio
40,Segovia,132,7,Moral de Hornuez
40,Segovia,134,8,Mozoncillo
40,Segovia,135,1,Muñopedro
40,Segovia,136,4,Muñoveros
40,Segovia,138,6,Nava de la Asunción
40,Segovia,139,9,Navafría
40,Segovia,140,3,Navalilla
40,Segovia,141,0,Navalmanzano
40,Segovia,142,5,Navares de Ayuso
40,Segovia,143,1,Navares de Enmedio
40,Segovia,144,6,Navares de las Cuevas
40,Segovia,145,9,Navas de Oro
40,Segovia,904,6,Navas de Riofrío
40,Segovia,146,2,Navas de San Antonio
40,Segovia,148,4,Nieva
40,Segovia,149,7,Olombrada
40,Segovia,150,0,Orejana
40,Segovia,151,7,Ortigosa de Pestaño
40,Segovia,901,0,Ortigosa del Monte
40,Segovia,152,2,Otero de Herreros
40,Segovia,154,3,Pajarejos
40,Segovia,155,6,Palazuelos de Eresma
40,Segovia,156,9,Pedraza
40,Segovia,157,5,Pelayos del Arroyo
40,Segovia,158,1,Perosillo
40,Segovia,159,4,Pinarejos
40,Segovia,160,8,Pinarnegrillo
40,Segovia,161,5,Pradales
40,Segovia,162,0,Prádena
40,Segovia,163,6,Puebla de Pedraza
40,Segovia,164,1,Rapariegos
40,Segovia,181,1,Real Sitio de San Ildefonso
40,Segovia,165,4,Rebollo
40,Segovia,166,7,Remondo
40,Segovia,168,9,Riaguas de San Bartolomé
40,Segovia,170,6,Riaza
40,Segovia,171,3,Ribota
40,Segovia,172,8,Riofrío de Riaza
40,Segovia,173,4,Roda de Eresma
40,Segovia,174,9,Sacramenia
40,Segovia,176,5,Samboal
40,Segovia,177,1,San Cristóbal de Cuéllar
40,Segovia,178,7,San Cristóbal de la Vega
40,Segovia,906,2,San Cristóbal de Segovia
40,Segovia,182,6,San Martín y Mudrián
40,Segovia,183,2,San Miguel de Bernuy
40,Segovia,184,7,San Pedro de Gaíllos
40,Segovia,179,0,Sanchonuño
40,Segovia,180,4,Sangarcía
40,Segovia,185,0,Santa María la Real de Nieva
40,Segovia,186,3,Santa Marta del Cerro
40,Segovia,188,5,Santiuste de Pedraza
40,Segovia,189,8,Santiuste de San Juan Bautista
40,Segovia,190,2,Santo Domingo de Pirón
40,Segovia,191,9,Santo Tomé del Puerto
40,Segovia,192,4,Sauquillo de Cabezas
40,Segovia,193,0,Sebúlcor
40,Segovia,194,5,Segovia
40,Segovia,195,8,Sepúlveda
40,Segovia,196,1,Sequera de Fresno
40,Segovia,198,3,Sotillo
40,Segovia,199,6,Sotosalbos
40,Segovia,200,0,Tabanera la Luenga
40,Segovia,201,7,Tolocirio
40,Segovia,206,9,Torre Val de San Pedro
40,Segovia,202,2,Torreadrada
40,Segovia,203,8,Torrecaballeros
40,Segovia,204,3,Torrecilla del Pinar
40,Segovia,205,6,Torreiglesias
40,Segovia,207,5,Trescasas
40,Segovia,208,1,Turégano
40,Segovia,210,8,Urueñas
40,Segovia,211,5,Valdeprados
40,Segovia,212,0,Valdevacas de Montejo
40,Segovia,213,6,Valdevacas y Guijar
40,Segovia,218,9,Valle de Tabladillo
40,Segovia,219,2,Vallelado
40,Segovia,220,6,Valleruela de Pedraza
40,Segovia,221,3,Valleruela de Sepúlveda
40,Segovia,214,1,Valseca
40,Segovia,215,4,Valtiendas
40,Segovia,216,7,Valverde del Majano
40,Segovia,222,8,Veganzones
40,Segovia,223,4,Vegas de Matute
40,Segovia,224,9,Ventosilla y Tejadilla
40,Segovia,225,2,Villacastín
40,Segovia,228,7,Villaverde de Íscar
40,Segovia,229,0,Villaverde de Montejo
40,Segovia,230,4,Villeguillo
40,Segovia,231,1,Yanguas de Eresma
40,Segovia,233,2,Zarzuela del Monte
40,Segovia,234,7,Zarzuela del Pinar
41,Sevilla,001,6,Aguadulce
41,Sevilla,002,1,Alanís
41,Sevilla,003,7,Albaida del Aljarafe
41,Sevilla,004,2,Alcalá de Guadaíra
41,Sevilla,005,5,Alcalá del Río
41,Sevilla,006,8,Alcolea del Río
41,Sevilla,007,4,"Algaba, La"
41,Sevilla,008,0,Algámitas
41,Sevilla,009,3,Almadén de la Plata
41,Sevilla,010,7,Almensilla
41,Sevilla,011,4,Arahal
41,Sevilla,012,9,Aznalcázar
41,Sevilla,013,5,Aznalcóllar
41,Sevilla,014,0,Badolatosa
41,Sevilla,015,3,Benacazón
41,Sevilla,016,6,Bollullos de la Mitación
41,Sevilla,017,2,Bormujos
41,Sevilla,018,8,Brenes
41,Sevilla,019,1,Burguillos
41,Sevilla,020,5,"Cabezas de San Juan, Las"
41,Sevilla,021,2,Camas
41,Sevilla,022,7,"Campana, La"
41,Sevilla,023,3,Cantillana
41,Sevilla,901,7,Cañada Rosal
41,Sevilla,024,8,Carmona
41,Sevilla,025,1,Carrión de los Céspedes
41,Sevilla,026,4,Casariche
41,Sevilla,027,0,Castilblanco de los Arroyos
41,Sevilla,028,6,Castilleja de Guzmán
41,Sevilla,029,9,Castilleja de la Cuesta
41,Sevilla,030,3,Castilleja del Campo
41,Sevilla,031,0,"Castillo de las Guardas, El"
41,Sevilla,032,5,Cazalla de la Sierra
41,Sevilla,033,1,Constantina
41,Sevilla,034,6,Coria del Río
41,Sevilla,035,9,Coripe
41,Sevilla,036,2,"Coronil, El"
41,Sevilla,037,8,"Corrales, Los"
41,Sevilla,903,8,"Cuervo de Sevilla, El"
41,Sevilla,038,4,Dos Hermanas
41,Sevilla,039,7,Écija
41,Sevilla,040,1,Espartinas
41,Sevilla,041,8,Estepa
41,Sevilla,042,3,Fuentes de Andalucía
41,Sevilla,043,9,"Garrobo, El"
41,Sevilla,044,4,Gelves
41,Sevilla,045,7,Gerena
41,Sevilla,046,0,Gilena
41,Sevilla,047,6,Gines
41,Sevilla,048,2,Guadalcanal
41,Sevilla,049,5,Guillena
41,Sevilla,050,8,Herrera
41,Sevilla,051,5,Huévar del Aljarafe
41,Sevilla,902,2,Isla Mayor
41,Sevilla,052,0,"Lantejuela, La"
41,Sevilla,053,6,Lebrija
41,Sevilla,054,1,Lora de Estepa
41,Sevilla,055,4,Lora del Río
41,Sevilla,056,7,"Luisiana, La"
41,Sevilla,057,3,"Madroño, El"
41,Sevilla,058,9,Mairena del Alcor
41,Sevilla,059,2,Mairena del Aljarafe
41,Sevilla,060,6,Marchena
41,Sevilla,061,3,Marinaleda
41,Sevilla,062,8,Martín de la Jara
41,Sevilla,063,4,"Molares, Los"
41,Sevilla,064,9,Montellano
41,Sevilla,065,2,Morón de la Frontera
41,Sevilla,066,5,"Navas de la Concepción, Las"
41,Sevilla,067,1,Olivares
41,Sevilla,068,7,Osuna
41,Sevilla,069,0,"Palacios y Villafranca, Los"
41,Sevilla,070,4,Palomares del Río
41,Sevilla,071,1,Paradas
41,Sevilla,072,6,Pedrera
41,Sevilla,073,2,"Pedroso, El"
41,Sevilla,074,7,Peñaflor
41,Sevilla,075,0,Pilas
41,Sevilla,076,3,Pruna
41,Sevilla,077,9,"Puebla de Cazalla, La"
41,Sevilla,078,5,"Puebla de los Infantes, La"
41,Sevilla,079,8,"Puebla del Río, La"
41,Sevilla,080,2,"Real de la Jara, El"
41,Sevilla,081,9,"Rinconada, La"
41,Sevilla,082,4,"Roda de Andalucía, La"
41,Sevilla,083,0,"Ronquillo, El"
41,Sevilla,084,5,"Rubio, El"
41,Sevilla,085,8,Salteras
41,Sevilla,086,1,San Juan de Aznalfarache
41,Sevilla,088,3,San Nicolás del Puerto
41,Sevilla,087,7,Sanlúcar la Mayor
41,Sevilla,089,6,Santiponce
41,Sevilla,090,0,"Saucejo, El"
41,Sevilla,091,7,Sevilla
41,Sevilla,092,2,Tocina
41,Sevilla,093,8,Tomares
41,Sevilla,094,3,Umbrete
41,Sevilla,095,6,Utrera
41,Sevilla,096,9,Valencina de la Concepción
41,Sevilla,097,5,Villamanrique de la Condesa
41,Sevilla,100,8,Villanueva de San Juan
41,Sevilla,098,1,Villanueva del Ariscal
41,Sevilla,099,4,Villanueva del Río y Minas
41,Sevilla,101,5,Villaverde del Río
41,Sevilla,102,0,"Viso del Alcor, El"
42,Soria,001,1,Abejar
42,Soria,003,2,Adradas
42,Soria,004,7,Ágreda
42,Soria,006,3,Alconaba
42,Soria,007,9,Alcubilla de Avellaneda
42,Soria,008,5,Alcubilla de las Peñas
42,Soria,009,8,Aldealafuente
42,Soria,010,2,Aldealices
42,Soria,011,9,Aldealpozo
42,Soria,012,4,Aldealseñor
42,Soria,013,0,Aldehuela de Periáñez
42,Soria,014,5,"Aldehuelas, Las"
42,Soria,015,8,Alentisque
42,Soria,016,1,Aliud
42,Soria,017,7,Almajano
42,Soria,018,3,Almaluez
42,Soria,019,6,Almarza
42,Soria,020,0,Almazán
42,Soria,021,7,Almazul
42,Soria,022,2,Almenar de Soria
42,Soria,023,8,Alpanseque
42,Soria,024,3,Arancón
42,Soria,025,6,Arcos de Jalón
42,Soria,026,9,Arenillas
42,Soria,027,5,Arévalo de la Sierra
42,Soria,028,1,Ausejo de la Sierra
42,Soria,029,4,Baraona
42,Soria,030,8,Barca
42,Soria,031,5,Barcones
42,Soria,032,0,Bayubas de Abajo
42,Soria,033,6,Bayubas de Arriba
42,Soria,034,1,Beratón
42,Soria,035,4,Berlanga de Duero
42,Soria,036,7,Blacos
42,Soria,037,3,Bliecos
42,Soria,038,9,Borjabad
42,Soria,039,2,Borobia
42,Soria,041,3,Buberos
42,Soria,042,8,Buitrago
42,Soria,043,4,Burgo de Osma-Ciudad de Osma
42,Soria,044,9,Cabrejas del Campo
42,Soria,045,2,Cabrejas del Pinar
42,Soria,046,5,Calatañazor
42,Soria,048,7,Caltojar
42,Soria,049,0,Candilichera
42,Soria,050,3,Cañamaque
42,Soria,051,0,Carabantes
42,Soria,052,5,Caracena
42,Soria,053,1,Carrascosa de Abajo
42,Soria,054,6,Carrascosa de la Sierra
42,Soria,055,9,Casarejos
42,Soria,056,2,Castilfrío de la Sierra
42,Soria,058,4,Castillejo de Robledo
42,Soria,057,8,Castilruiz
42,Soria,059,7,Centenera de Andaluz
42,Soria,060,1,Cerbón
42,Soria,061,8,Cidones
42,Soria,062,3,Cigudosa
42,Soria,063,9,Cihuela
42,Soria,064,4,Ciria
42,Soria,065,7,Cirujales del Río
42,Soria,068,2,Coscurita
42,Soria,069,5,Covaleda
42,Soria,070,9,Cubilla
42,Soria,071,6,Cubo de la Solana
42,Soria,073,7,Cueva de Ágreda
42,Soria,075,5,Dévanos
42,Soria,076,8,Deza
42,Soria,078,0,Duruelo de la Sierra
42,Soria,079,3,Escobosa de Almazán
42,Soria,080,7,Espeja de San Marcelino
42,Soria,081,4,Espejón
42,Soria,082,9,Estepa de San Juan
42,Soria,083,5,Frechilla de Almazán
42,Soria,084,0,Fresno de Caracena
42,Soria,085,3,Fuentearmegil
42,Soria,086,6,Fuentecambrón
42,Soria,087,2,Fuentecantos
42,Soria,088,8,Fuentelmonge
42,Soria,089,1,Fuentelsaz de Soria
42,Soria,090,5,Fuentepinilla
42,Soria,092,7,Fuentes de Magaña
42,Soria,093,3,Fuentestrún
42,Soria,094,8,Garray
42,Soria,095,1,Golmayo
42,Soria,096,4,Gómara
42,Soria,097,0,Gormaz
42,Soria,098,6,Herrera de Soria
42,Soria,100,3,Hinojosa del Campo
42,Soria,103,1,Langa de Duero
42,Soria,105,9,Liceras
42,Soria,106,2,"Losilla, La"
42,Soria,107,8,Magaña
42,Soria,108,4,Maján
42,Soria,110,1,Matalebreras
42,Soria,111,8,Matamala de Almazán
42,Soria,113,9,Medinaceli
42,Soria,115,7,Miño de Medinaceli
42,Soria,116,0,Miño de San Esteban
42,Soria,117,6,Molinos de Duero
42,Soria,118,2,Momblona
42,Soria,119,5,Monteagudo de las Vicarías
42,Soria,120,9,Montejo de Tiermes
42,Soria,121,6,Montenegro de Cameros
42,Soria,123,7,Morón de Almazán
42,Soria,124,2,Muriel de la Fuente
42,Soria,125,5,Muriel Viejo
42,Soria,127,4,Nafría de Ucero
42,Soria,128,0,Narros
42,Soria,129,3,Navaleno
42,Soria,130,7,Nepas
42,Soria,131,4,Nolay
42,Soria,132,9,Noviercas
42,Soria,134,0,Ólvega
42,Soria,135,3,Oncala
42,Soria,139,1,Pinilla del Campo
42,Soria,140,5,Portillo de Soria
42,Soria,141,2,"Póveda de Soria, La"
42,Soria,142,7,Pozalmuro
42,Soria,144,8,Quintana Redonda
42,Soria,145,1,Quintanas de Gormaz
42,Soria,148,6,Quiñonería
42,Soria,149,9,"Rábanos, Los"
42,Soria,151,9,Rebollar
42,Soria,152,4,Recuerda
42,Soria,153,0,Rello
42,Soria,154,5,Renieblas
42,Soria,155,8,Retortillo de Soria
42,Soria,156,1,Reznos
42,Soria,157,7,"Riba de Escalote, La"
42,Soria,158,3,Rioseco de Soria
42,Soria,159,6,Rollamienta
42,Soria,160,0,"Royo, El"
42,Soria,161,7,Salduero
42,Soria,162,2,San Esteban de Gormaz
42,Soria,163,8,San Felices
42,Soria,164,3,San Leonardo de Yagüe
42,Soria,165,6,San Pedro Manrique
42,Soria,166,9,Santa Cruz de Yanguas
42,Soria,167,5,Santa María de Huerta
42,Soria,168,1,Santa María de las Hoyas
42,Soria,171,5,Serón de Nágima
42,Soria,172,0,Soliedra
42,Soria,173,6,Soria
42,Soria,174,1,Sotillo del Rincón
42,Soria,175,4,Suellacabras
42,Soria,176,7,Tajahuerce
42,Soria,177,3,Tajueco
42,Soria,178,9,Talveila
42,Soria,181,3,Tardelcuende
42,Soria,182,8,Taroda
42,Soria,183,4,Tejado
42,Soria,184,9,Torlengua
42,Soria,185,2,Torreblacos
42,Soria,187,1,Torrubia de Soria
42,Soria,188,7,Trévago
42,Soria,189,0,Ucero
42,Soria,190,4,Vadillo
42,Soria,191,1,Valdeavellano de Tera
42,Soria,192,6,Valdegeña
42,Soria,193,2,Valdelagua del Cerro
42,Soria,194,7,Valdemaluque
42,Soria,195,0,Valdenebro
42,Soria,196,3,Valdeprado
42,Soria,197,9,Valderrodilla
42,Soria,198,5,Valtajeros
42,Soria,200,2,Velamazán
42,Soria,201,9,Velilla de la Sierra
42,Soria,202,4,Velilla de los Ajos
42,Soria,204,5,Viana de Duero
42,Soria,205,8,Villaciervos
42,Soria,206,1,Villanueva de Gormaz
42,Soria,207,7,Villar del Ala
42,Soria,208,3,Villar del Campo
42,Soria,209,6,Villar del Río
42,Soria,211,7,"Villares de Soria, Los"
42,Soria,212,2,Villasayas
42,Soria,213,8,Villaseca de Arciel
42,Soria,215,6,Vinuesa
42,Soria,216,9,Vizmanos
42,Soria,217,5,Vozmediano
42,Soria,218,1,Yanguas
42,Soria,219,4,Yelo
43,Tarragona,001,7,Aiguamúrcia
43,Tarragona,002,2,Albinyana
43,Tarragona,003,8,"Albiol, L'"
43,Tarragona,004,3,Alcanar
43,Tarragona,005,6,Alcover
43,Tarragona,904,4,"Aldea, L'"
43,Tarragona,006,9,Aldover
43,Tarragona,007,5,"Aleixar, L'"
43,Tarragona,008,1,Alfara de Carles
43,Tarragona,009,4,Alforja
43,Tarragona,010,8,Alió
43,Tarragona,011,5,Almoster
43,Tarragona,012,0,Altafulla
43,Tarragona,013,6,"Ametlla de Mar, L'"
43,Tarragona,906,0,"Ampolla, L'"
43,Tarragona,014,1,Amposta
43,Tarragona,016,7,"Arboç, L'"
43,Tarragona,015,4,Arbolí
43,Tarragona,017,3,"Argentera, L'"
43,Tarragona,018,9,Arnes
43,Tarragona,019,2,Ascó
43,Tarragona,020,6,Banyeres del Penedès
43,Tarragona,021,3,Barberà de la Conca
43,Tarragona,022,8,Batea
43,Tarragona,023,4,Bellmunt del Priorat
43,Tarragona,024,9,Bellvei
43,Tarragona,025,2,Benifallet
43,Tarragona,026,5,Benissanet
43,Tarragona,027,1,"Bisbal de Falset, La"
43,Tarragona,028,7,"Bisbal del Penedès, La"
43,Tarragona,029,0,Blancafort
43,Tarragona,030,4,Bonastre
43,Tarragona,031,1,"Borges del Camp, Les"
43,Tarragona,032,6,Bot
43,Tarragona,033,2,Botarell
43,Tarragona,034,7,Bràfim
43,Tarragona,035,0,Cabacés
43,Tarragona,036,3,Cabra del Camp
43,Tarragona,037,9,Calafell
43,Tarragona,903,9,Camarles
43,Tarragona,038,5,Cambrils
43,Tarragona,907,6,"Canonja, La"
43,Tarragona,039,8,Capafonts
43,Tarragona,040,2,Capçanes
43,Tarragona,041,9,Caseres
43,Tarragona,042,4,Castellvell del Camp
43,Tarragona,043,0,"Catllar, El"
43,Tarragona,045,8,Colldejou
43,Tarragona,046,1,Conesa
43,Tarragona,047,7,Constantí
43,Tarragona,048,3,Corbera d'Ebre
43,Tarragona,049,6,Cornudella de Montsant
43,Tarragona,050,9,Creixell
43,Tarragona,051,6,Cunit
43,Tarragona,901,8,Deltebre
43,Tarragona,053,7,Duesaigües
43,Tarragona,054,2,"Espluga de Francolí, L'"
43,Tarragona,055,5,Falset
43,Tarragona,056,8,"Fatarella, La"
43,Tarragona,057,4,"Febró, La"
43,Tarragona,058,0,"Figuera, La"
43,Tarragona,059,3,Figuerola del Camp
43,Tarragona,060,7,Flix
43,Tarragona,061,4,Forès
43,Tarragona,062,9,Freginals
43,Tarragona,063,5,"Galera, La"
43,Tarragona,064,0,Gandesa
43,Tarragona,065,3,Garcia
43,Tarragona,066,6,"Garidells, Els"
43,Tarragona,067,2,Ginestar
43,Tarragona,068,8,Godall
43,Tarragona,069,1,Gratallops
43,Tarragona,070,5,"Guiamets, Els"
43,Tarragona,071,2,Horta de Sant Joan
43,Tarragona,072,7,"Lloar, El"
43,Tarragona,073,3,Llorac
43,Tarragona,074,8,Llorenç del Penedès
43,Tarragona,076,4,Marçà
43,Tarragona,075,1,Margalef
43,Tarragona,077,0,Mas de Barberans
43,Tarragona,078,6,Masdenverge
43,Tarragona,079,9,Masllorenç
43,Tarragona,080,3,"Masó, La"
43,Tarragona,081,0,Maspujols
43,Tarragona,082,5,"Masroig, El"
43,Tarragona,083,1,"Milà, El"
43,Tarragona,084,6,Miravet
43,Tarragona,085,9,"Molar, El"
43,Tarragona,086,2,Montblanc
43,Tarragona,088,4,Montbrió del Camp
43,Tarragona,089,7,Montferri
43,Tarragona,090,1,"Montmell, El"
43,Tarragona,091,8,Mont-ral
43,Tarragona,092,3,Mont-roig del Camp
43,Tarragona,093,9,Móra d'Ebre
43,Tarragona,094,4,Móra la Nova
43,Tarragona,095,7,"Morell, El"
43,Tarragona,096,0,"Morera de Montsant, La"
43,Tarragona,097,6,"Nou de Gaià, La"
43,Tarragona,098,2,Nulles
43,Tarragona,100,9,"Pallaresos, Els"
43,Tarragona,099,5,"Palma d'Ebre, La"
43,Tarragona,101,6,Passanant i Belltall
43,Tarragona,102,1,Paüls
43,Tarragona,103,7,Perafort
43,Tarragona,104,2,"Perelló, El"
43,Tarragona,105,5,"Piles, Les"
43,Tarragona,106,8,"Pinell de Brai, El"
43,Tarragona,107,4,Pira
43,Tarragona,108,0,"Pla de Santa Maria, El"
43,Tarragona,109,3,"Pobla de Mafumet, La"
43,Tarragona,110,7,"Pobla de Massaluca, La"
43,Tarragona,111,4,"Pobla de Montornès, La"
43,Tarragona,112,9,Poboleda
43,Tarragona,113,5,"Pont d'Armentera, El"
43,Tarragona,141,8,Pontils
43,Tarragona,114,0,Porrera
43,Tarragona,115,3,Pradell de la Teixeta
43,Tarragona,116,6,Prades
43,Tarragona,117,2,Prat de Comte
43,Tarragona,118,8,Pratdip
43,Tarragona,119,1,Puigpelat
43,Tarragona,120,5,Querol
43,Tarragona,121,2,Rasquera
43,Tarragona,122,7,Renau
43,Tarragona,123,3,Reus
43,Tarragona,124,8,"Riba, La"
43,Tarragona,125,1,Riba-roja d'Ebre
43,Tarragona,126,4,"Riera de Gaià, La"
43,Tarragona,127,0,Riudecanyes
43,Tarragona,128,6,Riudecols
43,Tarragona,129,9,Riudoms
43,Tarragona,130,3,Rocafort de Queralt
43,Tarragona,131,0,Roda de Barà
43,Tarragona,132,5,Rodonyà
43,Tarragona,133,1,Roquetes
43,Tarragona,134,6,"Rourell, El"
43,Tarragona,135,9,Salomó
43,Tarragona,905,7,Salou
43,Tarragona,136,2,Sant Carles de la Ràpita
43,Tarragona,137,8,Sant Jaume dels Domenys
43,Tarragona,902,3,Sant Jaume d'Enveja
43,Tarragona,138,4,Santa Bàrbara
43,Tarragona,139,7,Santa Coloma de Queralt
43,Tarragona,140,1,Santa Oliva
43,Tarragona,142,3,Sarral
43,Tarragona,143,9,Savallà del Comtat
43,Tarragona,144,4,"Secuita, La"
43,Tarragona,145,7,"Selva del Camp, La"
43,Tarragona,146,0,Senan
43,Tarragona,044,5,"Sénia, La"
43,Tarragona,147,6,Solivella
43,Tarragona,148,2,Tarragona
43,Tarragona,149,5,Tivenys
43,Tarragona,150,8,Tivissa
43,Tarragona,151,5,"Torre de Fontaubella, La"
43,Tarragona,152,0,"Torre de l'Espanyol, La"
43,Tarragona,153,6,Torredembarra
43,Tarragona,154,1,Torroja del Priorat
43,Tarragona,155,4,Tortosa
43,Tarragona,156,7,Ulldecona
43,Tarragona,157,3,Ulldemolins
43,Tarragona,158,9,Vallclara
43,Tarragona,159,2,Vallfogona de Riucorb
43,Tarragona,160,6,Vallmoll
43,Tarragona,161,3,Valls
43,Tarragona,162,8,Vandellòs i l'Hospitalet de l'Infant
43,Tarragona,163,4,"Vendrell, El"
43,Tarragona,164,9,Vespella de Gaià
43,Tarragona,165,2,Vilabella
43,Tarragona,175,0,Vilalba dels Arcs
43,Tarragona,166,5,Vilallonga del Camp
43,Tarragona,168,7,Vilanova de Prades
43,Tarragona,167,1,Vilanova d'Escornalbou
43,Tarragona,169,0,Vilaplana
43,Tarragona,170,4,Vila-rodona
43,Tarragona,171,1,Vila-seca
43,Tarragona,172,6,Vilaverd
43,Tarragona,173,2,"Vilella Alta, La"
43,Tarragona,174,7,"Vilella Baixa, La"
43,Tarragona,176,3,Vimbodí i Poblet
43,Tarragona,177,9,Vinebre
43,Tarragona,178,5,Vinyols i els Arcs
43,Tarragona,052,1,Xerta
44,Teruel,001,2,Ababuj
44,Teruel,002,7,Abejuela
44,Teruel,003,3,Aguatón
44,Teruel,004,8,Aguaviva
44,Teruel,005,1,Aguilar del Alfambra
44,Teruel,006,4,Alacón
44,Teruel,007,0,Alba
44,Teruel,008,6,Albalate del Arzobispo
44,Teruel,009,9,Albarracín
44,Teruel,010,3,Albentosa
44,Teruel,011,0,Alcaine
44,Teruel,012,5,Alcalá de la Selva
44,Teruel,013,1,Alcañiz
44,Teruel,014,6,Alcorisa
44,Teruel,016,2,Alfambra
44,Teruel,017,8,Aliaga
44,Teruel,021,8,Allepuz
44,Teruel,022,3,Alloza
44,Teruel,023,9,Allueva
44,Teruel,018,4,Almohaja
44,Teruel,019,7,Alobras
44,Teruel,020,1,Alpeñés
44,Teruel,024,4,Anadón
44,Teruel,025,7,Andorra
44,Teruel,026,0,Arcos de las Salinas
44,Teruel,027,6,Arens de Lledó
44,Teruel,028,2,Argente
44,Teruel,029,5,Ariño
44,Teruel,031,6,Azaila
44,Teruel,032,1,Bádenas
44,Teruel,033,7,Báguena
44,Teruel,034,2,Bañón
44,Teruel,035,5,Barrachina
44,Teruel,036,8,Bea
44,Teruel,037,4,Beceite
44,Teruel,039,3,Bello
44,Teruel,038,0,Belmonte de San José
44,Teruel,040,7,Berge
44,Teruel,041,4,Bezas
44,Teruel,042,9,Blancas
44,Teruel,043,5,Blesa
44,Teruel,044,0,Bordón
44,Teruel,045,3,Bronchales
44,Teruel,046,6,Bueña
44,Teruel,047,2,Burbáguena
44,Teruel,048,8,Cabra de Mora
44,Teruel,049,1,Calaceite
44,Teruel,050,4,Calamocha
44,Teruel,051,1,Calanda
44,Teruel,052,6,Calomarde
44,Teruel,053,2,Camañas
44,Teruel,054,7,Camarena de la Sierra
44,Teruel,055,0,Camarillas
44,Teruel,056,3,Caminreal
44,Teruel,059,8,Cantavieja
44,Teruel,060,2,Cañada de Benatanduz
44,Teruel,061,9,"Cañada de Verich, La"
44,Teruel,062,4,Cañada Vellida
44,Teruel,063,0,Cañizar del Olivar
44,Teruel,064,5,Cascante del Río
44,Teruel,065,8,Castejón de Tornos
44,Teruel,066,1,Castel de Cabra
44,Teruel,070,0,"Castellar, El"
44,Teruel,071,7,Castellote
44,Teruel,067,7,Castelnou
44,Teruel,068,3,Castelserás
44,Teruel,074,3,Cedrillas
44,Teruel,075,6,Celadas
44,Teruel,076,9,Cella
44,Teruel,077,5,"Cerollera, La"
44,Teruel,080,8,"Codoñera, La"
44,Teruel,082,0,Corbalán
44,Teruel,084,1,Cortes de Aragón
44,Teruel,085,4,Cosa
44,Teruel,086,7,Cretas
44,Teruel,087,3,Crivillén
44,Teruel,088,9,"Cuba, La"
44,Teruel,089,2,Cubla
44,Teruel,090,6,Cucalón
44,Teruel,092,8,"Cuervo, El"
44,Teruel,093,4,Cuevas de Almudén
44,Teruel,094,9,Cuevas Labradas
44,Teruel,096,5,Ejulve
44,Teruel,097,1,Escorihuela
44,Teruel,099,0,Escucha
44,Teruel,100,4,Estercuel
44,Teruel,101,1,Ferreruela de Huerva
44,Teruel,102,6,Fonfría
44,Teruel,103,2,Formiche Alto
44,Teruel,105,0,Fórnoles
44,Teruel,106,3,Fortanete
44,Teruel,107,9,Foz-Calanda
44,Teruel,108,5,"Fresneda, La"
44,Teruel,109,8,Frías de Albarracín
44,Teruel,110,2,Fuenferrada
44,Teruel,111,9,Fuentes Calientes
44,Teruel,112,4,Fuentes Claras
44,Teruel,113,0,Fuentes de Rubielos
44,Teruel,114,5,Fuentespalda
44,Teruel,115,8,Galve
44,Teruel,116,1,Gargallo
44,Teruel,117,7,Gea de Albarracín
44,Teruel,118,3,"Ginebrosa, La"
44,Teruel,119,6,Griegos
44,Teruel,120,0,Guadalaviar
44,Teruel,121,7,Gúdar
44,Teruel,122,2,Híjar
44,Teruel,123,8,Hinojosa de Jarque
44,Teruel,124,3,"Hoz de la Vieja, La"
44,Teruel,125,6,Huesa del Común
44,Teruel,126,9,"Iglesuela del Cid, La"
44,Teruel,127,5,Jabaloyas
44,Teruel,128,1,Jarque de la Val
44,Teruel,129,4,Jatiel
44,Teruel,130,8,Jorcas
44,Teruel,131,5,Josa
44,Teruel,132,0,Lagueruela
44,Teruel,133,6,Lanzuela
44,Teruel,135,4,Libros
44,Teruel,136,7,Lidón
44,Teruel,137,3,Linares de Mora
44,Teruel,141,3,Lledó
44,Teruel,138,9,Loscos
44,Teruel,142,8,Maicas
44,Teruel,143,4,Manzanera
44,Teruel,144,9,Martín del Río
44,Teruel,145,2,Mas de las Matas
44,Teruel,146,5,"Mata de los Olmos, La"
44,Teruel,147,1,Mazaleón
44,Teruel,148,7,Mezquita de Jarque
44,Teruel,149,0,Mirambel
44,Teruel,150,3,Miravete de la Sierra
44,Teruel,151,0,Molinos
44,Teruel,152,5,Monforte de Moyuela
44,Teruel,153,1,Monreal del Campo
44,Teruel,154,6,Monroyo
44,Teruel,155,9,Montalbán
44,Teruel,156,2,Monteagudo del Castillo
44,Teruel,157,8,Monterde de Albarracín
44,Teruel,158,4,Mora de Rubielos
44,Teruel,159,7,Moscardón
44,Teruel,160,1,Mosqueruela
44,Teruel,161,8,Muniesa
44,Teruel,163,9,Noguera de Albarracín
44,Teruel,164,4,Nogueras
44,Teruel,165,7,Nogueruelas
44,Teruel,167,6,Obón
44,Teruel,168,2,Odón
44,Teruel,169,5,Ojos Negros
44,Teruel,171,6,Olba
44,Teruel,172,1,Oliete
44,Teruel,173,7,"Olmos, Los"
44,Teruel,174,2,Orihuela del Tremedal
44,Teruel,175,5,Orrios
44,Teruel,176,8,Palomar de Arroyos
44,Teruel,177,4,Pancrudo
44,Teruel,178,0,"Parras de Castellote, Las"
44,Teruel,179,3,Peñarroya de Tastavins
44,Teruel,180,7,Peracense
44,Teruel,181,4,Peralejos
44,Teruel,182,9,Perales del Alfambra
44,Teruel,183,5,Pitarque
44,Teruel,184,0,Plou
44,Teruel,185,3,"Pobo, El"
44,Teruel,187,2,"Portellada, La"
44,Teruel,189,1,Pozondón
44,Teruel,190,5,Pozuel del Campo
44,Teruel,191,2,"Puebla de Híjar, La"
44,Teruel,192,7,"Puebla de Valverde, La"
44,Teruel,193,3,Puertomingalvo
44,Teruel,194,8,Ráfales
44,Teruel,195,1,Rillo
44,Teruel,196,4,Riodeva
44,Teruel,197,0,Ródenas
44,Teruel,198,6,Royuela
44,Teruel,199,9,Rubiales
44,Teruel,200,3,Rubielos de la Cérida
44,Teruel,201,0,Rubielos de Mora
44,Teruel,203,1,Salcedillo
44,Teruel,204,6,Saldón
44,Teruel,205,9,Samper de Calanda
44,Teruel,206,2,San Agustín
44,Teruel,207,8,San Martín del Río
44,Teruel,208,4,Santa Cruz de Nogueras
44,Teruel,209,7,Santa Eulalia
44,Teruel,210,1,Sarrión
44,Teruel,211,8,Segura de los Baños
44,Teruel,212,3,Seno
44,Teruel,213,9,Singra
44,Teruel,215,7,Terriente
44,Teruel,216,0,Teruel
44,Teruel,217,6,Toril y Masegoso
44,Teruel,218,2,Tormón
44,Teruel,219,5,Tornos
44,Teruel,220,9,Torralba de los Sisones
44,Teruel,223,7,Torre de Arcas
44,Teruel,224,2,Torre de las Arcas
44,Teruel,225,5,Torre del Compte
44,Teruel,227,4,Torre los Negros
44,Teruel,221,6,Torrecilla de Alcañiz
44,Teruel,222,1,Torrecilla del Rebollar
44,Teruel,226,8,Torrelacárcel
44,Teruel,228,0,Torremocha de Jiloca
44,Teruel,229,3,Torres de Albarracín
44,Teruel,230,7,Torrevelilla
44,Teruel,231,4,Torrijas
44,Teruel,232,9,Torrijo del Campo
44,Teruel,234,0,Tramacastiel
44,Teruel,235,3,Tramacastilla
44,Teruel,236,6,Tronchón
44,Teruel,237,2,Urrea de Gaén
44,Teruel,238,8,Utrillas
44,Teruel,239,1,Valacloche
44,Teruel,240,5,Valbona
44,Teruel,241,2,Valdealgorfa
44,Teruel,243,3,Valdecuenca
44,Teruel,244,8,Valdelinares
44,Teruel,245,1,Valdeltormo
44,Teruel,246,4,Valderrobres
44,Teruel,247,0,Valjunquera
44,Teruel,249,9,"Vallecillo, El"
44,Teruel,250,2,Veguillas de la Sierra
44,Teruel,251,9,Villafranca del Campo
44,Teruel,252,4,Villahermosa del Campo
44,Teruel,256,1,Villanueva del Rebollar de la Sierra
44,Teruel,257,7,Villar del Cobo
44,Teruel,258,3,Villar del Salz
44,Teruel,260,0,Villarluengo
44,Teruel,261,7,Villarquemado
44,Teruel,262,2,Villarroya de los Pinares
44,Teruel,263,8,Villastar
44,Teruel,264,3,Villel
44,Teruel,265,6,Vinaceite
44,Teruel,266,9,Visiedo
44,Teruel,267,5,Vivel del Río Martín
44,Teruel,268,1,"Zoma, La"
45,Toledo,001,5,Ajofrín
45,Toledo,002,0,Alameda de la Sagra
45,Toledo,003,6,Albarreal de Tajo
45,Toledo,004,1,Alcabón
45,Toledo,005,4,Alcañizo
45,Toledo,006,7,Alcaudete de la Jara
45,Toledo,007,3,Alcolea de Tajo
45,Toledo,008,9,Aldea en Cabo
45,Toledo,009,2,Aldeanueva de Barbarroya
45,Toledo,010,6,Aldeanueva de San Bartolomé
45,Toledo,011,3,Almendral de la Cañada
45,Toledo,012,8,Almonacid de Toledo
45,Toledo,013,4,Almorox
45,Toledo,014,9,Añover de Tajo
45,Toledo,015,2,Arcicóllar
45,Toledo,016,5,Argés
45,Toledo,017,1,Azután
45,Toledo,018,7,Barcience
45,Toledo,019,0,Bargas
45,Toledo,020,4,Belvís de la Jara
45,Toledo,021,1,Borox
45,Toledo,022,6,Buenaventura
45,Toledo,023,2,Burguillos de Toledo
45,Toledo,024,7,Burujón
45,Toledo,025,0,Cabañas de la Sagra
45,Toledo,026,3,Cabañas de Yepes
45,Toledo,027,9,Cabezamesada
45,Toledo,028,5,Calera y Chozas
45,Toledo,029,8,Caleruela
45,Toledo,030,2,Calzada de Oropesa
45,Toledo,031,9,Camarena
45,Toledo,032,4,Camarenilla
45,Toledo,033,0,"Campillo de la Jara, El"
45,Toledo,034,5,Camuñas
45,Toledo,035,8,Cardiel de los Montes
45,Toledo,036,1,Carmena
45,Toledo,037,7,"Carpio de Tajo, El"
45,Toledo,038,3,Carranque
45,Toledo,039,6,Carriches
45,Toledo,040,0,"Casar de Escalona, El"
45,Toledo,041,7,Casarrubios del Monte
45,Toledo,042,2,Casasbuenas
45,Toledo,043,8,Castillo de Bayuela
45,Toledo,045,6,Cazalegas
45,Toledo,046,9,Cebolla
45,Toledo,047,5,Cedillo del Condado
45,Toledo,048,1,"Cerralbos, Los"
45,Toledo,049,4,Cervera de los Montes
45,Toledo,056,6,Chozas de Canales
45,Toledo,057,2,Chueca
45,Toledo,050,7,Ciruelos
45,Toledo,051,4,Cobeja
45,Toledo,052,9,Cobisa
45,Toledo,053,5,Consuegra
45,Toledo,054,0,Corral de Almaguer
45,Toledo,055,3,Cuerva
45,Toledo,058,8,Domingo Pérez
45,Toledo,059,1,Dosbarrios
45,Toledo,060,5,Erustes
45,Toledo,061,2,Escalona
45,Toledo,062,7,Escalonilla
45,Toledo,063,3,Espinoso del Rey
45,Toledo,064,8,Esquivias
45,Toledo,065,1,"Estrella, La"
45,Toledo,066,4,Fuensalida
45,Toledo,067,0,Gálvez
45,Toledo,068,6,Garciotum
45,Toledo,069,9,Gerindote
45,Toledo,070,3,Guadamur
45,Toledo,071,0,"Guardia, La"
45,Toledo,072,5,"Herencias, Las"
45,Toledo,073,1,Herreruela de Oropesa
45,Toledo,074,6,Hinojosa de San Vicente
45,Toledo,075,9,Hontanar
45,Toledo,076,2,Hormigos
45,Toledo,077,8,Huecas
45,Toledo,078,4,Huerta de Valdecarábanos
45,Toledo,079,7,"Iglesuela, La"
45,Toledo,080,1,Illán de Vacas
45,Toledo,081,8,Illescas
45,Toledo,082,3,Lagartera
45,Toledo,083,9,Layos
45,Toledo,084,4,Lillo
45,Toledo,085,7,Lominchar
45,Toledo,086,0,Lucillos
45,Toledo,087,6,Madridejos
45,Toledo,088,2,Magán
45,Toledo,089,5,Malpica de Tajo
45,Toledo,090,9,Manzaneque
45,Toledo,091,6,Maqueda
45,Toledo,092,1,Marjaliza
45,Toledo,093,7,Marrupe
45,Toledo,094,2,Mascaraque
45,Toledo,095,5,"Mata, La"
45,Toledo,096,8,Mazarambroz
45,Toledo,097,4,Mejorada
45,Toledo,098,0,Menasalbas
45,Toledo,099,3,Méntrida
45,Toledo,100,7,Mesegar de Tajo
45,Toledo,101,4,Miguel Esteban
45,Toledo,102,9,Mocejón
45,Toledo,103,5,Mohedas de la Jara
45,Toledo,104,0,Montearagón
45,Toledo,105,3,Montesclaros
45,Toledo,106,6,Mora
45,Toledo,107,2,Nambroca
45,Toledo,108,8,"Nava de Ricomalillo, La"
45,Toledo,109,1,Navahermosa
45,Toledo,110,5,Navalcán
45,Toledo,111,2,Navalmoralejo
45,Toledo,112,7,"Navalmorales, Los"
45,Toledo,113,3,"Navalucillos, Los"
45,Toledo,114,8,Navamorcuende
45,Toledo,115,1,Noblejas
45,Toledo,116,4,Noez
45,Toledo,117,0,Nombela
45,Toledo,118,6,Novés
45,Toledo,119,9,Numancia de la Sagra
45,Toledo,120,3,Nuño Gómez
45,Toledo,121,0,Ocaña
45,Toledo,122,5,Olías del Rey
45,Toledo,123,1,Ontígola
45,Toledo,124,6,Orgaz
45,Toledo,125,9,Oropesa
45,Toledo,126,2,Otero
45,Toledo,127,8,Palomeque
45,Toledo,128,4,Pantoja
45,Toledo,129,7,Paredes de Escalona
45,Toledo,130,1,Parrillas
45,Toledo,131,8,Pelahustán
45,Toledo,132,3,Pepino
45,Toledo,133,9,Polán
45,Toledo,134,4,Portillo de Toledo
45,Toledo,135,7,"Puebla de Almoradiel, La"
45,Toledo,136,0,"Puebla de Montalbán, La"
45,Toledo,137,6,"Pueblanueva, La"
45,Toledo,138,2,"Puente del Arzobispo, El"
45,Toledo,139,5,Puerto de San Vicente
45,Toledo,140,9,Pulgar
45,Toledo,141,6,Quero
45,Toledo,142,1,Quintanar de la Orden
45,Toledo,143,7,Quismondo
45,Toledo,144,2,"Real de San Vicente, El"
45,Toledo,145,5,Recas
45,Toledo,146,8,Retamoso de la Jara
45,Toledo,147,4,Rielves
45,Toledo,148,0,Robledo del Mazo
45,Toledo,149,3,"Romeral, El"
45,Toledo,150,6,San Bartolomé de las Abiertas
45,Toledo,151,3,San Martín de Montalbán
45,Toledo,152,8,San Martín de Pusa
45,Toledo,153,4,San Pablo de los Montes
45,Toledo,154,9,San Román de los Montes
45,Toledo,155,2,Santa Ana de Pusa
45,Toledo,156,5,Santa Cruz de la Zarza
45,Toledo,157,1,Santa Cruz del Retamar
45,Toledo,158,7,Santa Olalla
45,Toledo,901,6,Santo Domingo-Caudilla
45,Toledo,159,0,Sartajada
45,Toledo,160,4,Segurilla
45,Toledo,161,1,Seseña
45,Toledo,162,6,Sevilleja de la Jara
45,Toledo,163,2,Sonseca
45,Toledo,164,7,Sotillo de las Palomas
45,Toledo,165,0,Talavera de la Reina
45,Toledo,166,3,Tembleque
45,Toledo,167,9,"Toboso, El"
45,Toledo,168,5,Toledo
45,Toledo,169,8,Torralba de Oropesa
45,Toledo,171,9,"Torre de Esteban Hambrán, La"
45,Toledo,170,2,Torrecilla de la Jara
45,Toledo,172,4,Torrico
45,Toledo,173,0,Torrijos
45,Toledo,174,5,Totanés
45,Toledo,175,8,Turleque
45,Toledo,176,1,Ugena
45,Toledo,177,7,Urda
45,Toledo,179,6,Valdeverdeja
45,Toledo,180,0,Valmojado
45,Toledo,181,7,Velada
45,Toledo,182,2,"Ventas con Peña Aguilera, Las"
45,Toledo,183,8,"Ventas de Retamosa, Las"
45,Toledo,184,3,"Ventas de San Julián, Las"
45,Toledo,186,9,"Villa de Don Fadrique, La"
45,Toledo,185,6,Villacañas
45,Toledo,187,5,Villafranca de los Caballeros
45,Toledo,188,1,Villaluenga de la Sagra
45,Toledo,189,4,Villamiel de Toledo
45,Toledo,190,8,Villaminaya
45,Toledo,191,5,Villamuelas
45,Toledo,192,0,Villanueva de Alcardete
45,Toledo,193,6,Villanueva de Bogas
45,Toledo,194,1,Villarejo de Montalbán
45,Toledo,195,4,Villarrubia de Santiago
45,Toledo,196,7,Villaseca de la Sagra
45,Toledo,197,3,Villasequilla
45,Toledo,198,9,Villatobas
45,Toledo,199,2,"Viso de San Juan, El"
45,Toledo,200,6,"Yébenes, Los"
45,Toledo,201,3,Yeles
45,Toledo,202,8,Yepes
45,Toledo,203,4,Yuncler
45,Toledo,204,9,Yunclillos
45,Toledo,205,2,Yuncos
46,Valencia,001,8,Ademuz
46,Valencia,002,3,Ador
46,Valencia,004,4,Agullent
46,Valencia,042,5,Aielo de Malferit
46,Valencia,043,1,Aielo de Rugat
46,Valencia,005,7,Alaquàs
46,Valencia,006,0,Albaida
46,Valencia,007,6,Albal
46,Valencia,008,2,Albalat de la Ribera
46,Valencia,009,5,Albalat dels Sorells
46,Valencia,010,9,Albalat dels Tarongers
46,Valencia,011,6,Alberic
46,Valencia,012,1,Alborache
46,Valencia,013,7,Alboraya
46,Valencia,014,2,Albuixech
46,Valencia,016,8,Alcàntera de Xúquer
46,Valencia,015,5,Alcàsser
46,Valencia,018,0,Alcublas
46,Valencia,019,3,"Alcúdia, l'"
46,Valencia,020,7,"Alcúdia de Crespins, l'"
46,Valencia,021,4,Aldaia
46,Valencia,022,9,Alfafar
46,Valencia,024,0,Alfara de la Baronia
46,Valencia,025,3,Alfara del Patriarca
46,Valencia,026,6,Alfarp
46,Valencia,027,2,Alfarrasí
46,Valencia,023,5,Alfauir
46,Valencia,028,8,Algar de Palancia
46,Valencia,029,1,Algemesí
46,Valencia,030,5,Algimia de Alfara
46,Valencia,031,2,Alginet
46,Valencia,032,7,Almàssera
46,Valencia,033,3,Almiserà
46,Valencia,034,8,Almoines
46,Valencia,035,1,Almussafes
46,Valencia,036,4,Alpuente
46,Valencia,037,0,"Alqueria de la Comtessa, l'"
46,Valencia,017,4,Alzira
46,Valencia,038,6,Andilla
46,Valencia,039,9,Anna
46,Valencia,040,3,Antella
46,Valencia,041,0,Aras de los Olmos
46,Valencia,003,9,Atzeneta d'Albaida
46,Valencia,044,6,Ayora
46,Valencia,046,2,Barx
46,Valencia,045,9,Barxeta
46,Valencia,047,8,Bèlgida
46,Valencia,048,4,Bellreguard
46,Valencia,049,7,Bellús
46,Valencia,050,0,Benagéber
46,Valencia,051,7,Benaguasil
46,Valencia,052,2,Benavites
46,Valencia,053,8,Beneixida
46,Valencia,054,3,Benetússer
46,Valencia,055,6,Beniarjó
46,Valencia,056,9,Beniatjar
46,Valencia,057,5,Benicolet
46,Valencia,904,5,Benicull de Xúquer
46,Valencia,060,8,Benifaió
46,Valencia,059,4,Benifairó de la Valldigna
46,Valencia,058,1,Benifairó de les Valls
46,Valencia,061,5,Beniflá
46,Valencia,062,0,Benigánim
46,Valencia,063,6,Benimodo
46,Valencia,064,1,Benimuslem
46,Valencia,065,4,Beniparrell
46,Valencia,066,7,Benirredrà
46,Valencia,067,3,Benisanó
46,Valencia,068,9,Benissoda
46,Valencia,069,2,Benisuera
46,Valencia,070,6,Bétera
46,Valencia,071,3,Bicorp
46,Valencia,072,8,Bocairent
46,Valencia,073,4,Bolbaite
46,Valencia,074,9,Bonrepòs i Mirambell
46,Valencia,075,2,Bufali
46,Valencia,076,5,Bugarra
46,Valencia,077,1,Buñol
46,Valencia,078,7,Burjassot
46,Valencia,079,0,Calles
46,Valencia,080,4,Camporrobles
46,Valencia,081,1,Canals
46,Valencia,082,6,Canet d'En Berenguer
46,Valencia,083,2,Carcaixent
46,Valencia,084,7,Càrcer
46,Valencia,085,0,Carlet
46,Valencia,086,3,Carrícola
46,Valencia,087,9,Casas Altas
46,Valencia,088,5,Casas Bajas
46,Valencia,089,8,Casinos
46,Valencia,090,2,Castelló de Rugat
46,Valencia,091,9,Castellonet de la Conquesta
46,Valencia,092,4,Castielfabib
46,Valencia,093,0,Catadau
46,Valencia,094,5,Catarroja
46,Valencia,095,8,Caudete de las Fuentes
46,Valencia,096,1,Cerdà
46,Valencia,107,5,Chella
46,Valencia,106,9,Chelva
46,Valencia,108,1,Chera
46,Valencia,109,4,Cheste
46,Valencia,111,5,Chiva
46,Valencia,112,0,Chulilla
46,Valencia,097,7,Cofrentes
46,Valencia,098,3,Corbera
46,Valencia,099,6,Cortes de Pallás
46,Valencia,100,0,Cotes
46,Valencia,105,6,Cullera
46,Valencia,113,6,Daimús
46,Valencia,114,1,Domeño
46,Valencia,115,4,Dos Aguas
46,Valencia,116,7,"Eliana, l'"
46,Valencia,117,3,Emperador
46,Valencia,118,9,Enguera
46,Valencia,119,2,"Ènova, l'"
46,Valencia,120,6,Estivella
46,Valencia,121,3,Estubeny
46,Valencia,122,8,Faura
46,Valencia,123,4,Favara
46,Valencia,126,5,Foios
46,Valencia,128,7,"Font de la Figuera, la"
46,Valencia,127,1,"Font d'En Carròs, la"
46,Valencia,124,9,Fontanars dels Alforins
46,Valencia,125,2,Fortaleny
46,Valencia,129,0,Fuenterrobles
46,Valencia,131,1,Gandia
46,Valencia,902,4,Gátova
46,Valencia,130,4,Gavarda
46,Valencia,132,6,Genovés
46,Valencia,133,2,Gestalgar
46,Valencia,134,7,Gilet
46,Valencia,135,0,Godella
46,Valencia,136,3,Godelleta
46,Valencia,137,9,"Granja de la Costera, la"
46,Valencia,138,5,Guadasequies
46,Valencia,139,8,Guadassuar
46,Valencia,140,2,Guardamar de la Safor
46,Valencia,141,9,Higueruelas
46,Valencia,142,4,Jalance
46,Valencia,144,5,Jarafuel
46,Valencia,154,2,Llanera de Ranes
46,Valencia,155,5,Llaurí
46,Valencia,147,7,Llíria
46,Valencia,152,1,Llocnou de la Corona
46,Valencia,153,7,Llocnou de Sant Jeroni
46,Valencia,151,6,Llocnou d'En Fenollet
46,Valencia,156,8,Llombai
46,Valencia,157,4,"Llosa de Ranes, la"
46,Valencia,150,9,Llutxent
46,Valencia,148,3,Loriguilla
46,Valencia,149,6,Losa del Obispo
46,Valencia,158,0,Macastre
46,Valencia,159,3,Manises
46,Valencia,160,7,Manuel
46,Valencia,161,4,Marines
46,Valencia,162,9,Massalavés
46,Valencia,163,5,Massalfassar
46,Valencia,164,0,Massamagrell
46,Valencia,165,3,Massanassa
46,Valencia,166,6,Meliana
46,Valencia,167,2,Millares
46,Valencia,168,8,Miramar
46,Valencia,169,1,Mislata
46,Valencia,170,5,Mogente/Moixent
46,Valencia,171,2,Moncada
46,Valencia,173,3,Montaverner
46,Valencia,174,8,Montesa
46,Valencia,175,1,Montitxelvo/Montichelvo
46,Valencia,176,4,Montroy
46,Valencia,172,7,Montserrat
46,Valencia,177,0,Museros
46,Valencia,178,6,Náquera
46,Valencia,179,9,Navarrés
46,Valencia,180,3,Novelé/Novetlè
46,Valencia,181,0,Oliva
46,Valencia,183,1,"Olleria, l'"
46,Valencia,182,5,Olocau
46,Valencia,184,6,Ontinyent
46,Valencia,185,9,Otos
46,Valencia,186,2,Paiporta
46,Valencia,187,8,Palma de Gandía
46,Valencia,188,4,Palmera
46,Valencia,189,7,"Palomar, el"
46,Valencia,190,1,Paterna
46,Valencia,191,8,Pedralba
46,Valencia,192,3,Petrés
46,Valencia,193,9,Picanya
46,Valencia,194,4,Picassent
46,Valencia,195,7,Piles
46,Valencia,196,0,Pinet
46,Valencia,199,5,"Pobla de Farnals, la"
46,Valencia,202,1,"Pobla de Vallbona, la"
46,Valencia,200,9,"Pobla del Duc, la"
46,Valencia,203,7,"Pobla Llarga, la"
46,Valencia,197,6,Polinyà de Xúquer
46,Valencia,198,2,Potríes
46,Valencia,205,5,Puçol
46,Valencia,201,6,Puebla de San Miguel
46,Valencia,204,2,Puig
46,Valencia,101,7,Quart de les Valls
46,Valencia,102,2,Quart de Poblet
46,Valencia,103,8,Quartell
46,Valencia,104,3,Quatretonda
46,Valencia,206,8,Quesa
46,Valencia,207,4,Rafelbuñol/Rafelbunyol
46,Valencia,208,0,Rafelcofer
46,Valencia,209,3,Rafelguaraf
46,Valencia,210,7,Ráfol de Salem
46,Valencia,212,9,Real
46,Valencia,211,4,Real de Gandía
46,Valencia,213,5,Requena
46,Valencia,214,0,Riba-roja de Túria
46,Valencia,215,3,Riola
46,Valencia,216,6,Rocafort
46,Valencia,217,2,Rotglà i Corberà
46,Valencia,218,8,Rótova
46,Valencia,219,1,Rugat
46,Valencia,220,5,Sagunto/Sagunt
46,Valencia,221,2,Salem
46,Valencia,903,0,San Antonio de Benagéber
46,Valencia,222,7,Sant Joanet
46,Valencia,223,3,Sedaví
46,Valencia,224,8,Segart
46,Valencia,225,1,Sellent
46,Valencia,226,4,Sempere
46,Valencia,227,0,Senyera
46,Valencia,228,6,Serra
46,Valencia,229,9,Siete Aguas
46,Valencia,230,3,Silla
46,Valencia,231,0,Simat de la Valldigna
46,Valencia,232,5,Sinarcas
46,Valencia,233,1,Sollana
46,Valencia,234,6,Sot de Chera
46,Valencia,235,9,Sueca
46,Valencia,236,2,Sumacàrcer
46,Valencia,237,8,Tavernes Blanques
46,Valencia,238,4,Tavernes de la Valldigna
46,Valencia,239,7,Teresa de Cofrentes
46,Valencia,240,1,Terrateig
46,Valencia,241,8,Titaguas
46,Valencia,242,3,Torrebaja
46,Valencia,243,9,Torrella
46,Valencia,244,4,Torrent
46,Valencia,245,7,Torres Torres
46,Valencia,246,0,Tous
46,Valencia,247,6,Tuéjar
46,Valencia,248,2,Turís
46,Valencia,249,5,Utiel
46,Valencia,250,8,Valencia
46,Valencia,251,5,Vallada
46,Valencia,252,0,Vallanca
46,Valencia,253,6,Vallés
46,Valencia,254,1,Venta del Moro
46,Valencia,256,7,Vilamarxant
46,Valencia,255,4,Villalonga
46,Valencia,257,3,Villanueva de Castellón
46,Valencia,258,9,Villar del Arzobispo
46,Valencia,259,2,Villargordo del Cabriel
46,Valencia,260,6,Vinalesa
46,Valencia,145,8,Xàtiva
46,Valencia,143,0,Xeraco
46,Valencia,146,1,Xeresa
46,Valencia,110,8,Xirivella
46,Valencia,261,3,Yátova
46,Valencia,262,8,"Yesa, La"
46,Valencia,263,4,Zarra
47,Valladolid,001,4,Adalia
47,Valladolid,002,9,Aguasal
47,Valladolid,003,5,Aguilar de Campos
47,Valladolid,004,0,Alaejos
47,Valladolid,005,3,Alcazarén
47,Valladolid,006,6,Aldea de San Miguel
47,Valladolid,007,2,Aldeamayor de San Martín
47,Valladolid,008,8,Almenara de Adaja
47,Valladolid,009,1,Amusquillo
47,Valladolid,010,5,Arroyo de la Encomienda
47,Valladolid,011,2,Ataquines
47,Valladolid,012,7,Bahabón
47,Valladolid,013,3,Barcial de la Loma
47,Valladolid,014,8,Barruelo del Valle
47,Valladolid,015,1,Becilla de Valderaduey
47,Valladolid,016,4,Benafarces
47,Valladolid,017,0,Bercero
47,Valladolid,018,6,Berceruelo
47,Valladolid,019,9,Berrueces
47,Valladolid,020,3,Bobadilla del Campo
47,Valladolid,021,0,Bocigas
47,Valladolid,022,5,Bocos de Duero
47,Valladolid,023,1,Boecillo
47,Valladolid,024,6,Bolaños de Campos
47,Valladolid,025,9,Brahojos de Medina
47,Valladolid,026,2,Bustillo de Chaves
47,Valladolid,027,8,Cabezón de Pisuerga
47,Valladolid,028,4,Cabezón de Valderaduey
47,Valladolid,029,7,Cabreros del Monte
47,Valladolid,030,1,Campaspero
47,Valladolid,031,8,"Campillo, El"
47,Valladolid,032,3,Camporredondo
47,Valladolid,033,9,Canalejas de Peñafiel
47,Valladolid,034,4,Canillas de Esgueva
47,Valladolid,035,7,Carpio
47,Valladolid,036,0,Casasola de Arión
47,Valladolid,037,6,Castrejón de Trabancos
47,Valladolid,038,2,Castrillo de Duero
47,Valladolid,039,5,Castrillo-Tejeriego
47,Valladolid,040,9,Castrobol
47,Valladolid,041,6,Castrodeza
47,Valladolid,042,1,Castromembibre
47,Valladolid,043,7,Castromonte
47,Valladolid,044,2,Castronuevo de Esgueva
47,Valladolid,045,5,Castronuño
47,Valladolid,046,8,Castroponce
47,Valladolid,047,4,Castroverde de Cerrato
47,Valladolid,048,0,Ceinos de Campos
47,Valladolid,049,3,Cervillego de la Cruz
47,Valladolid,050,6,Cigales
47,Valladolid,051,3,Ciguñuela
47,Valladolid,052,8,Cistérniga
47,Valladolid,053,4,Cogeces de Íscar
47,Valladolid,054,9,Cogeces del Monte
47,Valladolid,055,2,Corcos
47,Valladolid,056,5,Corrales de Duero
47,Valladolid,057,1,Cubillas de Santa Marta
47,Valladolid,058,7,Cuenca de Campos
47,Valladolid,059,0,Curiel de Duero
47,Valladolid,060,4,Encinas de Esgueva
47,Valladolid,061,1,Esguevillas de Esgueva
47,Valladolid,062,6,Fombellida
47,Valladolid,063,2,Fompedraza
47,Valladolid,064,7,Fontihoyuelo
47,Valladolid,065,0,Fresno el Viejo
47,Valladolid,066,3,Fuensaldaña
47,Valladolid,067,9,Fuente el Sol
47,Valladolid,068,5,Fuente-Olmedo
47,Valladolid,069,8,Gallegos de Hornija
47,Valladolid,070,2,Gatón de Campos
47,Valladolid,071,9,Geria
47,Valladolid,073,0,Herrín de Campos
47,Valladolid,074,5,Hornillos de Eresma
47,Valladolid,075,8,Íscar
47,Valladolid,076,1,Laguna de Duero
47,Valladolid,077,7,Langayo
47,Valladolid,079,6,Llano de Olmedo
47,Valladolid,078,3,Lomoviejo
47,Valladolid,080,0,Manzanillo
47,Valladolid,081,7,Marzales
47,Valladolid,082,2,Matapozuelos
47,Valladolid,083,8,Matilla de los Caños
47,Valladolid,084,3,Mayorga
47,Valladolid,086,9,Medina de Rioseco
47,Valladolid,085,6,Medina del Campo
47,Valladolid,087,5,Megeces
47,Valladolid,088,1,Melgar de Abajo
47,Valladolid,089,4,Melgar de Arriba
47,Valladolid,090,8,Mojados
47,Valladolid,091,5,Monasterio de Vega
47,Valladolid,092,0,Montealegre de Campos
47,Valladolid,093,6,Montemayor de Pililla
47,Valladolid,094,1,Moral de la Reina
47,Valladolid,095,4,Moraleja de las Panaderas
47,Valladolid,096,7,Morales de Campos
47,Valladolid,097,3,Mota del Marqués
47,Valladolid,098,9,Mucientes
47,Valladolid,099,2,"Mudarra, La"
47,Valladolid,100,6,Muriel
47,Valladolid,101,3,Nava del Rey
47,Valladolid,102,8,Nueva Villa de las Torres
47,Valladolid,103,4,Olivares de Duero
47,Valladolid,104,9,Olmedo
47,Valladolid,105,2,Olmos de Esgueva
47,Valladolid,106,5,Olmos de Peñafiel
47,Valladolid,109,0,Palazuelo de Vedija
47,Valladolid,110,4,"Parrilla, La"
47,Valladolid,111,1,"Pedraja de Portillo, La"
47,Valladolid,112,6,Pedrajas de San Esteban
47,Valladolid,113,2,Pedrosa del Rey
47,Valladolid,114,7,Peñafiel
47,Valladolid,115,0,Peñaflor de Hornija
47,Valladolid,116,3,Pesquera de Duero
47,Valladolid,117,9,Piña de Esgueva
47,Valladolid,118,5,Piñel de Abajo
47,Valladolid,119,8,Piñel de Arriba
47,Valladolid,121,9,Pollos
47,Valladolid,122,4,Portillo
47,Valladolid,123,0,Pozal de Gallinas
47,Valladolid,124,5,Pozaldez
47,Valladolid,125,8,Pozuelo de la Orden
47,Valladolid,126,1,Puras
47,Valladolid,127,7,Quintanilla de Arriba
47,Valladolid,129,6,Quintanilla de Onésimo
47,Valladolid,130,0,Quintanilla de Trigueros
47,Valladolid,128,3,Quintanilla del Molar
47,Valladolid,131,7,Rábano
47,Valladolid,132,2,Ramiro
47,Valladolid,133,8,Renedo de Esgueva
47,Valladolid,134,3,Roales de Campos
47,Valladolid,135,6,Robladillo
47,Valladolid,137,5,Roturas
47,Valladolid,138,1,Rubí de Bracamonte
47,Valladolid,139,4,Rueda
47,Valladolid,140,8,Saelices de Mayorga
47,Valladolid,141,5,Salvador de Zapardiel
47,Valladolid,142,0,San Cebrián de Mazote
47,Valladolid,143,6,San Llorente
47,Valladolid,144,1,San Martín de Valvení
47,Valladolid,145,4,San Miguel del Arroyo
47,Valladolid,146,7,San Miguel del Pino
47,Valladolid,147,3,San Pablo de la Moraleja
47,Valladolid,148,9,San Pedro de Latarce
47,Valladolid,149,2,San Pelayo
47,Valladolid,150,5,San Román de Hornija
47,Valladolid,151,2,San Salvador
47,Valladolid,156,4,San Vicente del Palacio
47,Valladolid,152,7,Santa Eufemia del Arroyo
47,Valladolid,153,3,Santervás de Campos
47,Valladolid,154,8,Santibáñez de Valcorba
47,Valladolid,155,1,Santovenia de Pisuerga
47,Valladolid,157,0,Sardón de Duero
47,Valladolid,158,6,"Seca, La"
47,Valladolid,159,9,Serrada
47,Valladolid,160,3,Siete Iglesias de Trabancos
47,Valladolid,161,0,Simancas
47,Valladolid,162,5,Tamariz de Campos
47,Valladolid,163,1,Tiedra
47,Valladolid,164,6,Tordehumos
47,Valladolid,165,9,Tordesillas
47,Valladolid,169,7,Torre de Esgueva
47,Valladolid,170,1,Torre de Peñafiel
47,Valladolid,166,2,Torrecilla de la Abadesa
47,Valladolid,167,8,Torrecilla de la Orden
47,Valladolid,168,4,Torrecilla de la Torre
47,Valladolid,171,8,Torrelobatón
47,Valladolid,172,3,Torrescárcela
47,Valladolid,173,9,Traspinedo
47,Valladolid,174,4,Trigueros del Valle
47,Valladolid,175,7,Tudela de Duero
47,Valladolid,176,0,"Unión de Campos, La"
47,Valladolid,177,6,Urones de Castroponce
47,Valladolid,178,2,Urueña
47,Valladolid,179,5,Valbuena de Duero
47,Valladolid,180,9,Valdearcos de la Vega
47,Valladolid,181,6,Valdenebro de los Valles
47,Valladolid,182,1,Valdestillas
47,Valladolid,183,7,Valdunquillo
47,Valladolid,186,8,Valladolid
47,Valladolid,184,2,Valoria la Buena
47,Valladolid,185,5,Valverde de Campos
47,Valladolid,187,4,Vega de Ruiponce
47,Valladolid,188,0,Vega de Valdetronco
47,Valladolid,189,3,Velascálvaro
47,Valladolid,190,7,Velilla
47,Valladolid,191,4,Velliza
47,Valladolid,192,9,Ventosa de la Cuesta
47,Valladolid,193,5,Viana de Cega
47,Valladolid,195,3,Villabáñez
47,Valladolid,196,6,Villabaruz de Campos
47,Valladolid,197,2,Villabrágima
47,Valladolid,198,8,Villacarralón
47,Valladolid,199,1,Villacid de Campos
47,Valladolid,200,5,Villaco
47,Valladolid,203,3,Villafrades de Campos
47,Valladolid,204,8,Villafranca de Duero
47,Valladolid,205,1,Villafrechós
47,Valladolid,206,4,Villafuerte
47,Valladolid,207,0,Villagarcía de Campos
47,Valladolid,208,6,Villagómez la Nueva
47,Valladolid,209,9,Villalán de Campos
47,Valladolid,210,3,Villalar de los Comuneros
47,Valladolid,211,0,Villalba de la Loma
47,Valladolid,212,5,Villalba de los Alcores
47,Valladolid,213,1,Villalbarba
47,Valladolid,214,6,Villalón de Campos
47,Valladolid,215,9,Villamuriel de Campos
47,Valladolid,216,2,Villán de Tordesillas
47,Valladolid,217,8,Villanubla
47,Valladolid,218,4,Villanueva de Duero
47,Valladolid,219,7,Villanueva de la Condesa
47,Valladolid,220,1,Villanueva de los Caballeros
47,Valladolid,221,8,Villanueva de los Infantes
47,Valladolid,222,3,Villanueva de San Mancio
47,Valladolid,223,9,Villardefrades
47,Valladolid,224,4,Villarmentero de Esgueva
47,Valladolid,225,7,Villasexmir
47,Valladolid,226,0,Villavaquerín
47,Valladolid,227,6,Villavellid
47,Valladolid,228,2,Villaverde de Medina
47,Valladolid,229,5,Villavicencio de los Caballeros
47,Valladolid,194,0,Viloria
47,Valladolid,230,9,Wamba
47,Valladolid,231,6,Zaratán
47,Valladolid,232,1,"Zarza, La"
48,Vizcaya,001,0,Abadiño
48,Vizcaya,002,5,Abanto y Ciérvana-Abanto Zierbena
48,Vizcaya,911,9,Ajangiz
48,Vizcaya,912,4,Alonsotegi
48,Vizcaya,003,1,Amorebieta-Etxano
48,Vizcaya,004,6,Amoroto
48,Vizcaya,005,9,Arakaldo
48,Vizcaya,006,2,Arantzazu
48,Vizcaya,093,2,Areatza
48,Vizcaya,009,7,Arrankudiaga
48,Vizcaya,914,5,Arratzu
48,Vizcaya,010,1,Arrieta
48,Vizcaya,011,8,Arrigorriaga
48,Vizcaya,023,7,Artea
48,Vizcaya,008,4,Artzentales
48,Vizcaya,091,1,Atxondo
48,Vizcaya,070,8,Aulesti
48,Vizcaya,012,3,Bakio
48,Vizcaya,090,4,Balmaseda
48,Vizcaya,013,9,Barakaldo
48,Vizcaya,014,4,Barrika
48,Vizcaya,015,7,Basauri
48,Vizcaya,092,6,Bedia
48,Vizcaya,016,0,Berango
48,Vizcaya,017,6,Bermeo
48,Vizcaya,018,2,Berriatua
48,Vizcaya,019,5,Berriz
48,Vizcaya,020,9,Bilbao
48,Vizcaya,021,6,Busturia
48,Vizcaya,901,1,Derio
48,Vizcaya,026,8,Dima
48,Vizcaya,027,4,Durango
48,Vizcaya,028,0,Ea
48,Vizcaya,031,4,Elantxobe
48,Vizcaya,032,9,Elorrio
48,Vizcaya,902,6,Erandio
48,Vizcaya,033,5,Ereño
48,Vizcaya,034,0,Ermua
48,Vizcaya,079,2,Errigoiti
48,Vizcaya,029,3,Etxebarri
48,Vizcaya,030,7,Etxebarria
48,Vizcaya,906,3,Forua
48,Vizcaya,035,3,Fruiz
48,Vizcaya,036,6,Galdakao
48,Vizcaya,037,2,Galdames
48,Vizcaya,038,8,Gamiz-Fika
48,Vizcaya,039,1,Garai
48,Vizcaya,040,5,Gatika
48,Vizcaya,041,2,Gautegiz Arteaga
48,Vizcaya,046,4,Gernika-Lumo
48,Vizcaya,044,8,Getxo
48,Vizcaya,047,0,Gizaburuaga
48,Vizcaya,042,7,Gordexola
48,Vizcaya,043,3,Gorliz
48,Vizcaya,045,1,Güeñes
48,Vizcaya,048,6,Ibarrangelu
48,Vizcaya,094,7,Igorre
48,Vizcaya,049,9,Ispaster
48,Vizcaya,910,2,Iurreta
48,Vizcaya,050,2,Izurtza
48,Vizcaya,022,1,Karrantza Harana/Valle de Carranza
48,Vizcaya,907,9,Kortezubi
48,Vizcaya,051,9,Lanestosa
48,Vizcaya,052,4,Larrabetzu
48,Vizcaya,053,0,Laukiz
48,Vizcaya,054,5,Leioa
48,Vizcaya,057,7,Lekeitio
48,Vizcaya,055,8,Lemoa
48,Vizcaya,056,1,Lemoiz
48,Vizcaya,081,3,Lezama
48,Vizcaya,903,2,Loiu
48,Vizcaya,058,3,Mallabia
48,Vizcaya,059,6,Mañaria
48,Vizcaya,060,0,Markina-Xemein
48,Vizcaya,061,7,Maruri-Jatabe
48,Vizcaya,062,2,Mendata
48,Vizcaya,063,8,Mendexa
48,Vizcaya,064,3,Meñaka
48,Vizcaya,066,9,Morga
48,Vizcaya,068,1,Mundaka
48,Vizcaya,069,4,Mungia
48,Vizcaya,007,8,Munitibar-Arbatzegi Gerrikaitz
48,Vizcaya,908,5,Murueta
48,Vizcaya,071,5,Muskiz
48,Vizcaya,067,5,Muxika
48,Vizcaya,909,8,Nabarniz
48,Vizcaya,073,6,Ondarroa
48,Vizcaya,075,4,Orozko
48,Vizcaya,083,4,Ortuella
48,Vizcaya,072,0,Otxandio
48,Vizcaya,077,3,Plentzia
48,Vizcaya,078,9,Portugalete
48,Vizcaya,082,8,Santurtzi
48,Vizcaya,084,9,Sestao
48,Vizcaya,904,7,Sondika
48,Vizcaya,085,2,Sopelana
48,Vizcaya,086,5,Sopuerta
48,Vizcaya,076,7,Sukarrieta
48,Vizcaya,087,1,Trucios-Turtzioz
48,Vizcaya,088,7,Ubide
48,Vizcaya,065,6,Ugao-Miraballes
48,Vizcaya,089,0,Urduliz
48,Vizcaya,074,1,Urduña/Orduña
48,Vizcaya,080,6,Valle de Trápaga-Trapagaran
48,Vizcaya,095,0,Zaldibar
48,Vizcaya,096,3,Zalla
48,Vizcaya,905,0,Zamudio
48,Vizcaya,097,9,Zaratamo
48,Vizcaya,024,2,Zeanuri
48,Vizcaya,025,5,Zeberio
48,Vizcaya,913,0,Zierbena
48,Vizcaya,915,8,Ziortza-Bolibar
49,Zamora,002,8,Abezames
49,Zamora,003,4,Alcañices
49,Zamora,004,9,Alcubilla de Nogales
49,Zamora,005,2,Alfaraz de Sayago
49,Zamora,006,5,Algodre
49,Zamora,007,1,Almaraz de Duero
49,Zamora,008,7,Almeida de Sayago
49,Zamora,009,0,Andavías
49,Zamora,010,4,Arcenillas
49,Zamora,011,1,Arcos de la Polvorosa
49,Zamora,012,6,Argañín
49,Zamora,013,2,Argujillo
49,Zamora,014,7,Arquillinos
49,Zamora,015,0,Arrabalde
49,Zamora,016,3,Aspariegos
49,Zamora,017,9,Asturianos
49,Zamora,018,5,Ayoó de Vidriales
49,Zamora,019,8,Barcial del Barco
49,Zamora,020,2,Belver de los Montes
49,Zamora,021,9,Benavente
49,Zamora,022,4,Benegiles
49,Zamora,023,0,Bermillo de Sayago
49,Zamora,024,5,"Bóveda de Toro, La"
49,Zamora,025,8,Bretó
49,Zamora,026,1,Bretocino
49,Zamora,027,7,Brime de Sog
49,Zamora,028,3,Brime de Urz
49,Zamora,029,6,Burganes de Valverde
49,Zamora,030,0,Bustillo del Oro
49,Zamora,031,7,Cabañas de Sayago
49,Zamora,032,2,Calzadilla de Tera
49,Zamora,033,8,Camarzana de Tera
49,Zamora,034,3,Cañizal
49,Zamora,035,6,Cañizo
49,Zamora,036,9,Carbajales de Alba
49,Zamora,037,5,Carbellino
49,Zamora,038,1,Casaseca de Campeán
49,Zamora,039,4,Casaseca de las Chanas
49,Zamora,040,8,Castrillo de la Guareña
49,Zamora,041,5,Castrogonzalo
49,Zamora,042,0,Castronuevo
49,Zamora,043,6,Castroverde de Campos
49,Zamora,044,1,Cazurra
49,Zamora,046,7,Cerecinos de Campos
49,Zamora,047,3,Cerecinos del Carrizal
49,Zamora,048,9,Cernadilla
49,Zamora,050,5,Cobreros
49,Zamora,052,7,Coomonte
49,Zamora,053,3,Coreses
49,Zamora,054,8,Corrales del Vino
49,Zamora,055,1,Cotanes del Monte
49,Zamora,056,4,Cubillos
49,Zamora,057,0,Cubo de Benavente
49,Zamora,058,6,"Cubo de Tierra del Vino, El"
49,Zamora,059,9,Cuelgamures
49,Zamora,061,0,Entrala
49,Zamora,062,5,Espadañedo
49,Zamora,063,1,Faramontanos de Tábara
49,Zamora,064,6,Fariza
49,Zamora,065,9,Fermoselle
49,Zamora,066,2,Ferreras de Abajo
49,Zamora,067,8,Ferreras de Arriba
49,Zamora,068,4,Ferreruela
49,Zamora,069,7,Figueruela de Arriba
49,Zamora,071,8,Fonfría
49,Zamora,075,7,Fresno de la Polvorosa
49,Zamora,076,0,Fresno de la Ribera
49,Zamora,077,6,Fresno de Sayago
49,Zamora,078,2,Friera de Valverde
49,Zamora,079,5,Fuente Encalada
49,Zamora,080,9,Fuentelapeña
49,Zamora,082,1,Fuentes de Ropel
49,Zamora,081,6,Fuentesaúco
49,Zamora,083,7,Fuentesecas
49,Zamora,084,2,Fuentespreadas
49,Zamora,085,5,Galende
49,Zamora,086,8,Gallegos del Pan
49,Zamora,087,4,Gallegos del Río
49,Zamora,088,0,Gamones
49,Zamora,090,7,Gema
49,Zamora,091,4,Granja de Moreruela
49,Zamora,092,9,Granucillo
49,Zamora,093,5,Guarrate
49,Zamora,094,0,Hermisende
49,Zamora,095,3,"Hiniesta, La"
49,Zamora,096,6,Jambrina
49,Zamora,097,2,Justel
49,Zamora,098,8,Losacino
49,Zamora,099,1,Losacio
49,Zamora,100,5,Lubián
49,Zamora,101,2,Luelmo
49,Zamora,102,7,"Maderal, El"
49,Zamora,103,3,Madridanos
49,Zamora,104,8,Mahide
49,Zamora,105,1,Maire de Castroponce
49,Zamora,107,0,Malva
49,Zamora,108,6,Manganeses de la Lampreana
49,Zamora,109,9,Manganeses de la Polvorosa
49,Zamora,110,3,Manzanal de Arriba
49,Zamora,112,5,Manzanal de los Infantes
49,Zamora,111,0,Manzanal del Barco
49,Zamora,113,1,Matilla de Arzón
49,Zamora,114,6,Matilla la Seca
49,Zamora,115,9,Mayalde
49,Zamora,116,2,Melgar de Tera
49,Zamora,117,8,Micereces de Tera
49,Zamora,118,4,Milles de la Polvorosa
49,Zamora,119,7,Molacillos
49,Zamora,120,1,Molezuelas de la Carballeda
49,Zamora,121,8,Mombuey
49,Zamora,122,3,Monfarracinos
49,Zamora,123,9,Montamarta
49,Zamora,124,4,Moral de Sayago
49,Zamora,126,0,Moraleja de Sayago
49,Zamora,125,7,Moraleja del Vino
49,Zamora,128,2,Morales de Rey
49,Zamora,129,5,Morales de Toro
49,Zamora,130,9,Morales de Valverde
49,Zamora,127,6,Morales del Vino
49,Zamora,131,6,Moralina
49,Zamora,132,1,Moreruela de los Infanzones
49,Zamora,133,7,Moreruela de Tábara
49,Zamora,134,2,Muelas de los Caballeros
49,Zamora,135,5,Muelas del Pan
49,Zamora,136,8,Muga de Sayago
49,Zamora,137,4,Navianos de Valverde
49,Zamora,138,0,Olmillos de Castro
49,Zamora,139,3,Otero de Bodas
49,Zamora,141,4,Pajares de la Lampreana
49,Zamora,143,5,Palacios de Sanabria
49,Zamora,142,9,Palacios del Pan
49,Zamora,145,3,Pedralba de la Pradería
49,Zamora,146,6,"Pego, El"
49,Zamora,147,2,Peleagonzalo
49,Zamora,148,8,Peleas de Abajo
49,Zamora,149,1,Peñausende
49,Zamora,150,4,Peque
49,Zamora,151,1,"Perdigón, El"
49,Zamora,152,6,Pereruela
49,Zamora,153,2,Perilla de Castro
49,Zamora,154,7,Pías
49,Zamora,155,0,Piedrahita de Castro
49,Zamora,156,3,Pinilla de Toro
49,Zamora,157,9,Pino del Oro
49,Zamora,158,5,"Piñero, El"
49,Zamora,160,2,Pobladura de Valderaduey
49,Zamora,159,8,Pobladura del Valle
49,Zamora,162,4,Porto
49,Zamora,163,0,Pozoantiguo
49,Zamora,164,5,Pozuelo de Tábara
49,Zamora,165,8,Prado
49,Zamora,166,1,Puebla de Sanabria
49,Zamora,167,7,Pueblica de Valverde
49,Zamora,170,0,Quintanilla de Urz
49,Zamora,168,3,Quintanilla del Monte
49,Zamora,169,6,Quintanilla del Olmo
49,Zamora,171,7,Quiruelas de Vidriales
49,Zamora,172,2,Rabanales
49,Zamora,173,8,Rábano de Aliste
49,Zamora,174,3,Requejo
49,Zamora,175,6,Revellinos
49,Zamora,176,9,Riofrío de Aliste
49,Zamora,177,5,Rionegro del Puente
49,Zamora,178,1,Roales
49,Zamora,179,4,Robleda-Cervantes
49,Zamora,180,8,Roelos de Sayago
49,Zamora,181,5,Rosinos de la Requejada
49,Zamora,183,6,Salce
49,Zamora,184,1,Samir de los Caños
49,Zamora,185,4,San Agustín del Pozo
49,Zamora,186,7,San Cebrián de Castro
49,Zamora,187,3,San Cristóbal de Entreviñas
49,Zamora,188,9,San Esteban del Molar
49,Zamora,189,2,San Justo
49,Zamora,190,6,San Martín de Valderaduey
49,Zamora,191,3,San Miguel de la Ribera
49,Zamora,192,8,San Miguel del Valle
49,Zamora,193,4,San Pedro de Ceque
49,Zamora,194,9,San Pedro de la Nave-Almendra
49,Zamora,208,5,San Vicente de la Cabeza
49,Zamora,209,8,San Vitero
49,Zamora,197,1,Santa Clara de Avedillo
49,Zamora,199,0,Santa Colomba de las Monjas
49,Zamora,200,4,Santa Cristina de la Polvorosa
49,Zamora,201,1,Santa Croya de Tera
49,Zamora,202,6,Santa Eufemia del Barco
49,Zamora,203,2,Santa María de la Vega
49,Zamora,204,7,Santa María de Valverde
49,Zamora,205,0,Santibáñez de Tera
49,Zamora,206,3,Santibáñez de Vidriales
49,Zamora,207,9,Santovenia
49,Zamora,210,2,Sanzoles
49,Zamora,214,5,Tábara
49,Zamora,216,1,Tapioles
49,Zamora,219,6,Toro
49,Zamora,220,0,"Torre del Valle, La"
49,Zamora,221,7,Torregamones
49,Zamora,222,2,Torres del Carrizal
49,Zamora,223,8,Trabazos
49,Zamora,224,3,Trefacio
49,Zamora,225,6,Uña de Quintana
49,Zamora,226,9,Vadillo de la Guareña
49,Zamora,227,5,Valcabado
49,Zamora,228,1,Valdefinjas
49,Zamora,229,4,Valdescorriel
49,Zamora,230,8,Vallesa de la Guareña
49,Zamora,231,5,Vega de Tera
49,Zamora,232,0,Vega de Villalobos
49,Zamora,233,6,Vegalatrave
49,Zamora,234,1,Venialbo
49,Zamora,235,4,Vezdemarbán
49,Zamora,236,7,Vidayanes
49,Zamora,237,3,Videmala
49,Zamora,238,9,Villabrázaro
49,Zamora,239,2,Villabuena del Puente
49,Zamora,240,6,Villadepera
49,Zamora,241,3,Villaescusa
49,Zamora,242,8,Villafáfila
49,Zamora,243,4,Villaferrueña
49,Zamora,244,9,Villageriz
49,Zamora,245,2,Villalazán
49,Zamora,246,5,Villalba de la Lampreana
49,Zamora,247,1,Villalcampo
49,Zamora,248,7,Villalobos
49,Zamora,249,0,Villalonso
49,Zamora,250,3,Villalpando
49,Zamora,251,0,Villalube
49,Zamora,252,5,Villamayor de Campos
49,Zamora,255,9,Villamor de los Escuderos
49,Zamora,256,2,Villanázar
49,Zamora,257,8,Villanueva de Azoague
49,Zamora,258,4,Villanueva de Campeán
49,Zamora,259,7,Villanueva de las Peras
49,Zamora,260,1,Villanueva del Campo
49,Zamora,263,9,Villar de Fallaves
49,Zamora,264,4,Villar del Buey
49,Zamora,261,8,Villaralbo
49,Zamora,262,3,Villardeciervos
49,Zamora,265,7,Villardiegua de la Ribera
49,Zamora,266,0,Villárdiga
49,Zamora,267,6,Villardondiego
49,Zamora,268,2,Villarrín de Campos
49,Zamora,269,5,Villaseco del Pan
49,Zamora,270,9,Villavendimio
49,Zamora,272,1,Villaveza de Valverde
49,Zamora,271,6,Villaveza del Agua
49,Zamora,273,7,Viñas
49,Zamora,275,5,Zamora
50,Zaragoza,001,6,Abanto
50,Zaragoza,002,1,Acered
50,Zaragoza,003,7,Agón
50,Zaragoza,004,2,Aguarón
50,Zaragoza,005,5,Aguilón
50,Zaragoza,006,8,Ainzón
50,Zaragoza,007,4,Aladrén
50,Zaragoza,008,0,Alagón
50,Zaragoza,009,3,Alarba
50,Zaragoza,010,7,Alberite de San Juan
50,Zaragoza,011,4,Albeta
50,Zaragoza,012,9,Alborge
50,Zaragoza,013,5,Alcalá de Ebro
50,Zaragoza,014,0,Alcalá de Moncayo
50,Zaragoza,015,3,Alconchel de Ariza
50,Zaragoza,016,6,Aldehuela de Liestos
50,Zaragoza,017,2,Alfajarín
50,Zaragoza,018,8,Alfamén
50,Zaragoza,019,1,Alforque
50,Zaragoza,020,5,Alhama de Aragón
50,Zaragoza,021,2,Almochuel
50,Zaragoza,022,7,"Almolda, La"
50,Zaragoza,023,3,Almonacid de la Cuba
50,Zaragoza,024,8,Almonacid de la Sierra
50,Zaragoza,025,1,"Almunia de Doña Godina, La"
50,Zaragoza,026,4,Alpartir
50,Zaragoza,027,0,Ambel
50,Zaragoza,028,6,Anento
50,Zaragoza,029,9,Aniñón
50,Zaragoza,030,3,Añón de Moncayo
50,Zaragoza,031,0,Aranda de Moncayo
50,Zaragoza,032,5,Arándiga
50,Zaragoza,033,1,Ardisa
50,Zaragoza,034,6,Ariza
50,Zaragoza,035,9,Artieda
50,Zaragoza,036,2,Asín
50,Zaragoza,037,8,Atea
50,Zaragoza,038,4,Ateca
50,Zaragoza,039,7,Azuara
50,Zaragoza,040,1,Badules
50,Zaragoza,041,8,Bagüés
50,Zaragoza,042,3,Balconchán
50,Zaragoza,043,9,Bárboles
50,Zaragoza,044,4,Bardallur
50,Zaragoza,045,7,Belchite
50,Zaragoza,046,0,Belmonte de Gracián
50,Zaragoza,047,6,Berdejo
50,Zaragoza,048,2,Berrueco
50,Zaragoza,901,7,Biel
50,Zaragoza,050,8,Bijuesca
50,Zaragoza,051,5,Biota
50,Zaragoza,052,0,Bisimbre
50,Zaragoza,053,6,Boquiñeni
50,Zaragoza,054,1,Bordalba
50,Zaragoza,055,4,Borja
50,Zaragoza,056,7,Botorrita
50,Zaragoza,057,3,Brea de Aragón
50,Zaragoza,058,9,Bubierca
50,Zaragoza,059,2,Bujaraloz
50,Zaragoza,060,6,Bulbuente
50,Zaragoza,061,3,Bureta
50,Zaragoza,062,8,"Burgo de Ebro, El"
50,Zaragoza,063,4,"Buste, El"
50,Zaragoza,064,9,Cabañas de Ebro
50,Zaragoza,065,2,Cabolafuente
50,Zaragoza,066,5,Cadrete
50,Zaragoza,067,1,Calatayud
50,Zaragoza,068,7,Calatorao
50,Zaragoza,069,0,Calcena
50,Zaragoza,070,4,Calmarza
50,Zaragoza,071,1,Campillo de Aragón
50,Zaragoza,072,6,Carenas
50,Zaragoza,073,2,Cariñena
50,Zaragoza,074,7,Caspe
50,Zaragoza,075,0,Castejón de Alarba
50,Zaragoza,076,3,Castejón de las Armas
50,Zaragoza,077,9,Castejón de Valdejasa
50,Zaragoza,078,5,Castiliscar
50,Zaragoza,079,8,Cervera de la Cañada
50,Zaragoza,080,2,Cerveruela
50,Zaragoza,081,9,Cetina
50,Zaragoza,092,2,Chiprana
50,Zaragoza,093,8,Chodes
50,Zaragoza,082,4,Cimballa
50,Zaragoza,083,0,Cinco Olivas
50,Zaragoza,084,5,Clarés de Ribota
50,Zaragoza,085,8,Codo
50,Zaragoza,086,1,Codos
50,Zaragoza,087,7,Contamina
50,Zaragoza,088,3,Cosuenda
50,Zaragoza,089,6,Cuarte de Huerva
50,Zaragoza,090,0,Cubel
50,Zaragoza,091,7,"Cuerlas, Las"
50,Zaragoza,094,3,Daroca
50,Zaragoza,095,6,Ejea de los Caballeros
50,Zaragoza,096,9,Embid de Ariza
50,Zaragoza,098,1,Encinacorba
50,Zaragoza,099,4,Épila
50,Zaragoza,100,8,Erla
50,Zaragoza,101,5,Escatrón
50,Zaragoza,102,0,Fabara
50,Zaragoza,104,1,Farlete
50,Zaragoza,105,4,Fayón
50,Zaragoza,106,7,"Fayos, Los"
50,Zaragoza,107,3,Figueruelas
50,Zaragoza,108,9,Fombuena
50,Zaragoza,109,2,"Frago, El"
50,Zaragoza,110,6,"Frasno, El"
50,Zaragoza,111,3,Fréscano
50,Zaragoza,113,4,Fuendejalón
50,Zaragoza,114,9,Fuendetodos
50,Zaragoza,115,2,Fuentes de Ebro
50,Zaragoza,116,5,Fuentes de Jiloca
50,Zaragoza,117,1,Gallocanta
50,Zaragoza,118,7,Gallur
50,Zaragoza,119,0,Gelsa
50,Zaragoza,120,4,Godojos
50,Zaragoza,121,1,Gotor
50,Zaragoza,122,6,Grisel
50,Zaragoza,123,2,Grisén
50,Zaragoza,124,7,Herrera de los Navarros
50,Zaragoza,125,0,Ibdes
50,Zaragoza,126,3,Illueca
50,Zaragoza,128,5,Isuerre
50,Zaragoza,129,8,Jaraba
50,Zaragoza,130,2,Jarque
50,Zaragoza,131,9,Jaulín
50,Zaragoza,132,4,"Joyosa, La"
50,Zaragoza,133,0,Lagata
50,Zaragoza,134,5,Langa del Castillo
50,Zaragoza,135,8,Layana
50,Zaragoza,136,1,Lécera
50,Zaragoza,138,3,Lechón
50,Zaragoza,137,7,Leciñena
50,Zaragoza,139,6,Letux
50,Zaragoza,140,0,Litago
50,Zaragoza,141,7,Lituénigo
50,Zaragoza,142,2,Lobera de Onsella
50,Zaragoza,143,8,Longares
50,Zaragoza,144,3,Longás
50,Zaragoza,146,9,Lucena de Jalón
50,Zaragoza,147,5,Luceni
50,Zaragoza,148,1,Luesia
50,Zaragoza,149,4,Luesma
50,Zaragoza,150,7,Lumpiaque
50,Zaragoza,151,4,Luna
50,Zaragoza,152,9,Maella
50,Zaragoza,153,5,Magallón
50,Zaragoza,154,0,Mainar
50,Zaragoza,155,3,Malanquilla
50,Zaragoza,156,6,Maleján
50,Zaragoza,160,5,Mallén
50,Zaragoza,157,2,Malón
50,Zaragoza,159,1,Maluenda
50,Zaragoza,161,2,Manchones
50,Zaragoza,162,7,Mara
50,Zaragoza,163,3,María de Huerva
50,Zaragoza,902,2,Marracos
50,Zaragoza,164,8,Mediana de Aragón
50,Zaragoza,165,1,Mequinenza
50,Zaragoza,166,4,Mesones de Isuela
50,Zaragoza,167,0,Mezalocha
50,Zaragoza,168,6,Mianos
50,Zaragoza,169,9,Miedes de Aragón
50,Zaragoza,170,3,Monegrillo
50,Zaragoza,171,0,Moneva
50,Zaragoza,172,5,Monreal de Ariza
50,Zaragoza,173,1,Monterde
50,Zaragoza,174,6,Montón
50,Zaragoza,175,9,Morata de Jalón
50,Zaragoza,176,2,Morata de Jiloca
50,Zaragoza,177,8,Morés
50,Zaragoza,178,4,Moros
50,Zaragoza,179,7,Moyuela
50,Zaragoza,180,1,Mozota
50,Zaragoza,181,8,Muel
50,Zaragoza,182,3,"Muela, La"
50,Zaragoza,183,9,Munébrega
50,Zaragoza,184,4,Murero
50,Zaragoza,185,7,Murillo de Gállego
50,Zaragoza,186,0,Navardún
50,Zaragoza,187,6,Nigüella
50,Zaragoza,188,2,Nombrevilla
50,Zaragoza,189,5,Nonaspe
50,Zaragoza,190,9,Novallas
50,Zaragoza,191,6,Novillas
50,Zaragoza,192,1,Nuévalos
50,Zaragoza,193,7,Nuez de Ebro
50,Zaragoza,194,2,Olvés
50,Zaragoza,195,5,Orcajo
50,Zaragoza,196,8,Orera
50,Zaragoza,197,4,Orés
50,Zaragoza,198,0,Oseja
50,Zaragoza,199,3,Osera de Ebro
50,Zaragoza,200,7,Paniza
50,Zaragoza,201,4,Paracuellos de Jiloca
50,Zaragoza,202,9,Paracuellos de la Ribera
50,Zaragoza,203,5,Pastriz
50,Zaragoza,204,0,Pedrola
50,Zaragoza,205,3,"Pedrosas, Las"
50,Zaragoza,206,6,Perdiguera
50,Zaragoza,207,2,Piedratajada
50,Zaragoza,208,8,Pina de Ebro
50,Zaragoza,209,1,Pinseque
50,Zaragoza,210,5,"Pintanos, Los"
50,Zaragoza,211,2,Plasencia de Jalón
50,Zaragoza,212,7,Pleitas
50,Zaragoza,213,3,Plenas
50,Zaragoza,214,8,Pomer
50,Zaragoza,215,1,Pozuel de Ariza
50,Zaragoza,216,4,Pozuelo de Aragón
50,Zaragoza,217,0,Pradilla de Ebro
50,Zaragoza,218,6,Puebla de Albortón
50,Zaragoza,219,9,"Puebla de Alfindén, La"
50,Zaragoza,220,3,Puendeluna
50,Zaragoza,221,0,Purujosa
50,Zaragoza,222,5,Quinto
50,Zaragoza,223,1,Remolinos
50,Zaragoza,224,6,Retascón
50,Zaragoza,225,9,Ricla
50,Zaragoza,227,8,Romanos
50,Zaragoza,228,4,Rueda de Jalón
50,Zaragoza,229,7,Ruesca
50,Zaragoza,241,6,Sabiñán
50,Zaragoza,230,1,Sádaba
50,Zaragoza,231,8,Salillas de Jalón
50,Zaragoza,232,3,Salvatierra de Esca
50,Zaragoza,233,9,Samper del Salz
50,Zaragoza,234,4,San Martín de la Virgen de Moncayo
50,Zaragoza,235,7,San Mateo de Gállego
50,Zaragoza,236,0,Santa Cruz de Grío
50,Zaragoza,237,6,Santa Cruz de Moncayo
50,Zaragoza,238,2,Santa Eulalia de Gállego
50,Zaragoza,239,5,Santed
50,Zaragoza,240,9,Sástago
50,Zaragoza,242,1,Sediles
50,Zaragoza,243,7,Sestrica
50,Zaragoza,244,2,Sierra de Luna
50,Zaragoza,245,5,Sigüés
50,Zaragoza,246,8,Sisamón
50,Zaragoza,247,4,Sobradiel
50,Zaragoza,248,0,Sos del Rey Católico
50,Zaragoza,249,3,Tabuenca
50,Zaragoza,250,6,Talamantes
50,Zaragoza,251,3,Tarazona
50,Zaragoza,252,8,Tauste
50,Zaragoza,253,4,Terrer
50,Zaragoza,254,9,Tierga
50,Zaragoza,255,2,Tobed
50,Zaragoza,256,5,Torralba de los Frailes
50,Zaragoza,257,1,Torralba de Ribota
50,Zaragoza,258,7,Torralbilla
50,Zaragoza,259,0,Torrehermosa
50,Zaragoza,260,4,Torrelapaja
50,Zaragoza,261,1,Torrellas
50,Zaragoza,262,6,Torres de Berrellén
50,Zaragoza,263,2,Torrijo de la Cañada
50,Zaragoza,264,7,Tosos
50,Zaragoza,265,0,Trasmoz
50,Zaragoza,266,3,Trasobares
50,Zaragoza,267,9,Uncastillo
50,Zaragoza,268,5,Undués de Lerda
50,Zaragoza,269,8,Urrea de Jalón
50,Zaragoza,270,2,Urriés
50,Zaragoza,271,9,Used
50,Zaragoza,272,4,Utebo
50,Zaragoza,274,5,Val de San Martín
50,Zaragoza,273,0,Valdehorna
50,Zaragoza,275,8,Valmadrid
50,Zaragoza,276,1,Valpalmas
50,Zaragoza,277,7,Valtorres
50,Zaragoza,278,3,Velilla de Ebro
50,Zaragoza,279,6,Velilla de Jiloca
50,Zaragoza,280,0,Vera de Moncayo
50,Zaragoza,281,7,Vierlas
50,Zaragoza,283,8,Villadoz
50,Zaragoza,284,3,Villafeliche
50,Zaragoza,285,6,Villafranca de Ebro
50,Zaragoza,286,9,Villalba de Perejil
50,Zaragoza,287,5,Villalengua
50,Zaragoza,903,8,Villamayor de Gállego
50,Zaragoza,288,1,Villanueva de Gállego
50,Zaragoza,290,8,Villanueva de Huerva
50,Zaragoza,289,4,Villanueva de Jiloca
50,Zaragoza,291,5,Villar de los Navarros
50,Zaragoza,292,0,Villarreal de Huerva
50,Zaragoza,293,6,Villarroya de la Sierra
50,Zaragoza,294,1,Villarroya del Campo
50,Zaragoza,282,2,"Vilueña, La"
50,Zaragoza,295,4,Vistabella
50,Zaragoza,296,7,"Zaida, La"
50,Zaragoza,297,3,Zaragoza
50,Zaragoza,298,9,Zuera
51,Ceuta,001,3,Ceuta
52,Melilla,001,8,Melilla`,
	)
