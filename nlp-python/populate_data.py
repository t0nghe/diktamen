import spacy
import mysql.connector
from dotenv import dotenv_values
from mysql.connector import Error


CLOUDFRONT_PREFIX = "http://d1zg52ope8o24c.cloudfront.net/"
ARTICLES = [
    { "slug": "meningsfullt/meningsfullt_",
    "title": "Vad gör livet meningsfullt? ", "sentences": [
        "Familjen ger livet dess värde.", "Och inom familjen värderar vi barnen långt högre än en eventuell partner.", "Men vi uppskattar familj och partner olika beroende på faktorer som exempelvis ålder, kön och var i världen vi bor. ", "Den framgår av en ny internationell studie av befolkningarna i 17 relativt rika länder.", "Familj, men även arbete, välstånd och hälsa. Det är vad som ger oss mening i livet ...", "Svaren kanske inte förvånar, men frågan har nu analyserats i en stor enkätundersökning av Pew Research Center i USA.", "Under våren 2021 intervjuades nästan 19 000 vuxna personer i 17 relativt rika länder i fyra världsdelar.", "Sverige var ett av länderna i studien.", "Huvudfrågan var öppen, och löd:", "Vi är intresserade av att undersöka vad det betyder att leva ett meningsfullt liv.", "Vilka aspekter av ditt nuvarande liv finner du meningsfulla och tillfredsställande?", "Det återkommande och vanligaste svaret var familjen.", "I 14 av de 17 länderna kom familjen allra först.", "Men i de två östasiatiska länderna Sydkorea och Taiwan placerades familjen först på tredje plats."
    ] },
    {
        "title": "Zelenskyj är en man som kvinnor längtar efter",
        "slug": "zelenskyj/zelenskyj_",
        "sentences": [
            "Det är 75 mil mellan mitt skrivbord i Malmö och närmaste gränsstad till Ukraina, Medyka, en liten polsk by med några tusen invånare.", "På andra sidan gränsen pågår ett krig, inte ett sådant som utspelar sig i spalter och debattprogram, utan ett sådant som dödar människor och utplånar städer. ", "Mot bakgrund av detta ter sig samtalet i den svenska offentligheten som vi har vant oss vid det som bisarrt.", "I åratal har ett oförsonligt kulturkrig rasat, med offer till både höger och vänster,", "båda sidor lika övertygade om sin förträfflighet och om att just deras världsbild är korrekt.", "Det har framställts  som en kamp mellan mörker och ljus, och människor har, ofta på tämligen dunkla grunder, dubbats till goda eller dömts som onda.", "Hjältar, martyrer och profeter – som Greta Thunberg och Cissi Wallin, Katerina Janouch, Hanif Bali och Joakim Lamotte, allt efter tycke och smak – har kommit och gått.", "Varenda människa som någon gång publicerat en åsikt har nogsamt sorterats in, oavsett om man vill vara en del av detta fiktiva krig eller inte.", "Diskussioner om vilka ord som är tillåtna och otillåtna, uppförandekoder för standup-branschen,", "vem som är ond respektive god och vilka pronomen som är lämpliga ter sig nu som väldigt obsoleta. ", "Världen är inte bara en berättelse. Den är.", "Människan är inte blott en konstruktion. Hon är en kropp som känner, lider och dör.", "Krig förs inte bara på Twitter, utan på gator och torg, vapen är inte bara ord, utan gevär och bomber.", "Trots dessa uppenbara påpekanden ser jag hur det pågående kriget i Ukraina fiktionaliseras.", "Vi har två huvudpersoner, den hjältemodige president Zelenskyj och den galne diktatorn Vladimir Putin.", "Sociala medier fylls av hyllningar till Zelenskyj och nidbilder av Putin,", "reportage om den osannolike presidenten, från komiker till statschef skrivs i rasande fart, liksom analyser som avfärdar Putin som ”sjuk”, ”vansinnig” och ”galen”."
        ]
    }
]

nlp = spacy.load('sv_core_news_lg')
config = dotenv_values(".env")

def create_connection(host_name, username, password, db_name): 
    try:
        connection = mysql.connector.connect(
            host=host_name,
            user=username,
            passwd=password,
            database=db_name
            )
        print("Connection established.")
        return connection
    except Error as err:
        print("Error", err)
        return None

def execute_query(connection, query):
    print(query)
    cursor = connection.cursor()
    try:
        ret = cursor.execute(query)
        connection.commit()
        print("Query successful")
        return ret
    except Error as err:
        print(f"Error: '{err}'")

# def execute_query(connection, query):
#     print(query)

connection = create_connection(config['DBHOST'], config['DBUSERNAME'], config['DBPASSWORD'], config['DBNAME'])

art_id = 4
sent_id = 0
word_id = 0

for k in range(len(ARTICLES)):
    art_id += 1
    article=ARTICLES[k]
    slug = article['slug']
    sent_count = len(article['sentences'])
    title = article['title']

    create_article_query = f"INSERT INTO article (id, title, sent_count) VALUES ({art_id}, '{title}', {sent_count});"
    execute_query(connection, create_article_query)
    
    for i in range(len(article['sentences'])):
        sent_id += 1
        index_in_article = i + 1
        media_uri = CLOUDFRONT_PREFIX + slug + "{:02d}".format(i+1) + ".mp3"
        text = article['sentences'][i]
        doc = nlp(text)
        word_count = len(doc)
        insert_sentence_query = f"INSERT INTO sent (id, article_id, index_in_article, media_uri, text, word_count) VALUES ({sent_id}, {art_id}, {index_in_article}, '{media_uri}', '{text}', {word_count});"
        execute_query(connection, insert_sentence_query)

        for j in range(word_count):
            word_id += 1
            tok = doc[j]
            wordform = tok.text
            if tok.ent_iob != 2 or tok.pos_ == "PROPN" or tok.pos_ == "PUNCT":
                is_cloze = False
            else:
                is_cloze = True
            insert_word_query = f"INSERT INTO word (id, wordform, length, sent_id, index_in_sent, is_cloze) VALUES ({word_id}, '{wordform}', {len(wordform)}, {sent_id}, {j+1}, {is_cloze});"
            execute_query(connection, insert_word_query)