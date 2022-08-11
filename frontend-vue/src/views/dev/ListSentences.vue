<script setup lang="ts">
import SentenceTried from "../../components/Sentence/SentenceTried.vue";
import SentenceNew from "../../components/Sentence/SentenceNew.vue";

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
</template>

<style></style>
