CREATE TABLE "articles" (
  "article_id" serial,
  "author" varchar(50),
  "audio_seconds" int,
  PRIMARY KEY ("article_id")
);

CREATE TABLE "learners" (
  "learner_id" varchar(50),
  "email" varchar(50),
  "passwordHash" varchar(100),
  PRIMARY KEY ("learner_id")
);

CREATE TABLE "sentences" (
  "sentence_id" serial,
  "audio_url" varchar(50),
  "text" text,
  "text_wordcount" int,
  "audio_seconds" int,
  "text_lemmas" text[],
  "article_id" bigint,
  "index_in_article" int,
  PRIMARY KEY ("sentence_id"),
  CONSTRAINT "FK_sentences.article_id"
    FOREIGN KEY ("article_id")
      REFERENCES "articles"("article_id")
);

CREATE TABLE "learner_sentence" (
  "learner_id" varchar(50),
  "sentence_id" int,
  "learner_heard_text" text,
  "learner_heard_correct" boolean,
  "learner_activity_id" bigserial,
  "learner_hearing_time" timestamp,
  "incorrect_words" text[],
  PRIMARY KEY ("learner_activity_id"),
  CONSTRAINT "FK_learner_sentence.learner_id"
    FOREIGN KEY ("learner_id")
      REFERENCES "learners"("learner_id"),
  CONSTRAINT "FK_learner_sentence.sentence_id"
    FOREIGN KEY ("sentence_id")
      REFERENCES "sentences"("sentence_id")
);

