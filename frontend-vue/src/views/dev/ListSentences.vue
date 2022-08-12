<script setup lang="ts">
import SentenceTried from "../../components/Sentence/SentenceTried.vue";
import SentenceNew from "../../components/Sentence/SentenceNew.vue";
import SentenceTrying from "../../components/Sentence/SentenceTrying.vue";

const correctExample = {
  sentId: 17,
  indexInArticle: 3,
  tryText:
    "Mot bakgrund av detta ter sig samtalet i den svenska offentligheten som vi har vant oss vid det som bisarrt .",
};

const incorrectExample = {
  sentId: 15,
  indexInArticle: 1,
  sentWords: [
    {
      length: 3,
      isCloze: true,
      wordform: "Det",
      indexInSent: 1,
      lastInputText: "Det",
      lastInputScore: 1,
    },
    {
      length: 2,
      isCloze: true,
      wordform: "är",
      indexInSent: 2,
      lastInputText: "är",
      lastInputScore: 1,
    },
    {
      length: 2,
      isCloze: true,
      wordform: "75",
      indexInSent: 3,
      lastInputText: "85",
      lastInputScore: 0.5,
    },
    {
      length: 3,
      isCloze: true,
      wordform: "mil",
      indexInSent: 4,
      lastInputText: "mil",
      lastInputScore: 1,
    },
    {
      length: 6,
      isCloze: true,
      wordform: "mellan",
      indexInSent: 5,
      lastInputText: "melxxx",
      lastInputScore: 0.5,
    },
    {
      length: 4,
      isCloze: true,
      wordform: "mitt",
      indexInSent: 6,
      lastInputText: "mixx",
      lastInputScore: 0.5,
    },
    {
      length: 9,
      isCloze: true,
      wordform: "skrivbord",
      indexInSent: 7,
      lastInputText: "skrivbord",
      lastInputScore: 1,
    },
    {
      length: 1,
      isCloze: true,
      wordform: "i",
      indexInSent: 8,
      lastInputText: "i",
      lastInputScore: 1,
    },
    {
      length: 5,
      isCloze: false,
      wordform: "Malmö",
      indexInSent: 9,
      lastInputText: "",
      lastInputScore: 0,
    },
    {
      length: 3,
      isCloze: true,
      wordform: "och",
      indexInSent: 10,
      lastInputText: "och",
      lastInputScore: 1,
    },
    {
      length: 8,
      isCloze: true,
      wordform: "närmaste",
      indexInSent: 11,
      lastInputText: "xäxmxsxe",
      lastInputScore: 0.5555556,
    },
    {
      length: 9,
      isCloze: true,
      wordform: "gränsstad",
      indexInSent: 12,
      lastInputText: "xrxnxsxax",
      lastInputScore: 0.5,
    },
    {
      length: 4,
      isCloze: true,
      wordform: "till",
      indexInSent: 13,
      lastInputText: "till",
      lastInputScore: 1,
    },
    {
      length: 7,
      isCloze: false,
      wordform: "Ukraina",
      indexInSent: 14,
      lastInputText: "",
      lastInputScore: 0,
    },
    {
      length: 1,
      isCloze: false,
      wordform: ",",
      indexInSent: 15,
      lastInputText: "",
      lastInputScore: 0,
    },
    {
      length: 6,
      isCloze: true,
      wordform: "Medyka",
      indexInSent: 16,
      lastInputText: "Mxdxkx",
      lastInputScore: 0.5,
    },
    {
      length: 1,
      isCloze: false,
      wordform: ",",
      indexInSent: 17,
      lastInputText: "",
      lastInputScore: 0,
    },
    {
      length: 2,
      isCloze: true,
      wordform: "en",
      indexInSent: 18,
      lastInputText: "en",
      lastInputScore: 1,
    },
    {
      length: 5,
      isCloze: true,
      wordform: "liten",
      indexInSent: 19,
      lastInputText: "liten",
      lastInputScore: 1,
    },
    {
      length: 5,
      isCloze: true,
      wordform: "polsk",
      indexInSent: 20,
      lastInputText: "xoxsx",
      lastInputScore: 0.4,
    },
    {
      length: 2,
      isCloze: true,
      wordform: "by",
      indexInSent: 21,
      lastInputText: "bx",
      lastInputScore: 0.5,
    },
    {
      length: 3,
      isCloze: true,
      wordform: "med",
      indexInSent: 22,
      lastInputText: "med",
      lastInputScore: 1,
    },
    {
      length: 5,
      isCloze: true,
      wordform: "några",
      indexInSent: 23,
      lastInputText: "några",
      lastInputScore: 1,
    },
    {
      length: 5,
      isCloze: true,
      wordform: "tusen",
      indexInSent: 24,
      lastInputText: "tusen",
      lastInputScore: 1,
    },
    {
      length: 8,
      isCloze: true,
      wordform: "invånare",
      indexInSent: 25,
      lastInputText: "invånare",
      lastInputScore: 1,
    },
    {
      length: 1,
      isCloze: false,
      wordform: ".",
      indexInSent: 26,
      lastInputText: "",
      lastInputScore: 0,
    },
  ],
};

const unseenSentExample = {
  sentId: 18,
  indexInArticle: 4,
  mediaUri: "https://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_04.mp3",
  sentWords: [
    {
      length: 1,
      isCloze: true,
      wordform: "I",
      indexInSent: 1,
    },
    {
      length: 6,
      isCloze: false,
      wordform: "åratal",
      indexInSent: 2,
    },
    {
      length: 3,
      isCloze: true,
      wordform: "har",
      indexInSent: 3,
    },
    {
      length: 3,
      isCloze: false,
      wordform: "ett",
      indexInSent: 4,
    },
    {
      length: 11,
      isCloze: true,
      wordform: "oförsonligt",
      indexInSent: 5,
    },
    {
      length: 10,
      isCloze: false,
      wordform: "kulturkrig",
      indexInSent: 6,
    },
    {
      length: 5,
      isCloze: true,
      wordform: "rasat",
      indexInSent: 7,
    },
    {
      length: 1,
      isCloze: false,
      wordform: ",",
      indexInSent: 8,
    },
    {
      length: 3,
      isCloze: true,
      wordform: "med",
      indexInSent: 9,
    },
    {
      length: 5,
      isCloze: false,
      wordform: "offer",
      indexInSent: 10,
    },
    {
      length: 4,
      isCloze: true,
      wordform: "till",
      indexInSent: 11,
    },
    {
      length: 4,
      isCloze: false,
      wordform: "både",
      indexInSent: 12,
    },
    {
      length: 5,
      isCloze: true,
      wordform: "höger",
      indexInSent: 13,
    },
    {
      length: 3,
      isCloze: false,
      wordform: "och",
      indexInSent: 14,
    },
    {
      length: 7,
      isCloze: true,
      wordform: "vänster",
      indexInSent: 15,
    },
    {
      length: 1,
      isCloze: false,
      wordform: ",",
      indexInSent: 16,
    },
  ],
};

const shouldPlay = () => {
  console.log("should play sound");
};

const shouldSubmitSent = (payload) => {
  console.log(payload);
};
</script>

<template>
  <h1>List of sentence components</h1>
  <hr />
  <h2>Sentence Tried (correct)</h2>
  <SentenceTried
    :isCorrect="true"
    :sentId="correctExample.sentId"
    :indexInArticle="correctExample.indexInArticle"
    :tryText="correctExample.tryText"
  />
  <h2>Sentence Tried (incorrect)</h2>
  <SentenceTried
    :isCorrect="false"
    :sentId="incorrectExample.sentId"
    :indexInArticle="incorrectExample.indexInArticle"
    :sentWords="incorrectExample.sentWords"
  />
  <h2>Sentence New - that awaits you to listen!</h2>
  <SentenceNew
    :sent-id="unseenSentExample.sentId"
    :index-in-article="unseenSentExample.indexInArticle"
    :sent-words="unseenSentExample.sentWords"
  />
  <h2>Sentence Trying - enjoy it!</h2>
  <div>
    &nbsp;&nbsp;
    <span v-for="w in unseenSentExample.sentWords" 
      >{{ w.wordform }} &nbsp;</span
    >
  </div>
  <SentenceTrying
    :sent-id="unseenSentExample.sentId"
    :index-in-article="unseenSentExample.indexInArticle"
    :sent-words="unseenSentExample.sentWords"
    @play-sound="shouldPlay"
    @submit-sent="shouldSubmitSent"
  />
  <h2>...</h2>
</template>

<style></style>
