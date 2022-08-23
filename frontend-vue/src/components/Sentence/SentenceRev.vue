<script setup lang="ts">
import { ref, watch, computed, onMounted } from "vue";
import {
  sentWord,
  userTrySent,
  inputValuesType,
  inputFieldsType,
} from "../../types";
import PlayPauseRev from "../Interaction/PlayPauseRev.vue";
import ArrowBlueRev from "../Interaction/ArrowBlueRev.vue";

const props = defineProps<{
  sentId: number;
  mediaUri: string;
  sentWords: sentWord[];
  finnsNext: boolean;
}>();

const emit = defineEmits<{ (e: "submit-sent", payload: userTrySent): void }>();

const playPauseCounter = ref(0);

const isComplete = ref<boolean>(false);
const inputFields = ref<inputFieldsType>({});
const inputValues = ref<inputValuesType>({});

const newSentIdRef = computed(() => props.sentId);
watch(newSentIdRef, () => {
  inputValues.value = {};
  isComplete.value = false;
});

const focusPrevInput = (idxSent: string) => {
  const keys = Object.keys(inputFields.value);
  const prevIdxSent = keys[keys.indexOf(idxSent) - 1];
  if (prevIdxSent) {
    const prevField = inputFields.value[prevIdxSent];
    prevField.focus();
  }
};

const focusNextInput = (idxSent: string) => {
  const keys = Object.keys(inputFields.value);
  const nextIdxSent = keys[keys.indexOf(idxSent) + 1];
  if (nextIdxSent) {
    const nextField = inputFields.value[nextIdxSent];
    nextField.focus();
  } else {
    // if this input box is the last available,
    // this sentence seems to be finished.
    trySubmitSent();
  }
};

const readyToSubmit = (): boolean => {
  // BUG NOTE (FIXED): The following line is a problem. It seems:
  // Each input element is bound to a value in this object.
  // and before interaction, the value is NOT appended to this object.
  // In this object, only the keys of fields that HAVE a value are present.
  // const keys = Object.keys(inputValues.value);
  // Using this instead, will list the keys to all the input fields.
  const keys = Object.keys(inputFields.value);
  for (let i = 0; i < keys.length; i++) {
    const key = keys[i];
    if (!inputValues.value[key]) {
      inputFields.value[key].classList.add("word-user-input-missing");
      inputFields.value[key].focus();
      return false;
    }
  }
  isComplete.value = true;
  return true;
};

const trySubmitSent = () => {
  // when press enter:
  // if all input fields contain a value:
  // - emit: submit-sent event, among other things, include { index-in-sent, userinput } in data;
  const userInputWords = Object.keys(inputValues.value).map((key) => ({
    indexInSent: parseInt(key),
    inputText: inputValues.value[key],
  }));

  const payload = {
    sentId: props.sentId,
    userInputWords,
  };
  if (readyToSubmit()) {
    emit("submit-sent", payload);
  }
};

const onInputHandler = (ev) => {
  // if user input value contains space:
  // - trigger: play-sound
  // - access ref and remove space from value
  const inputContent: string = ev.target.value.toLowerCase();
  const idxSent: string = ev.target.attributes["data-index"].value;
  const thisWord = props.sentWords.find(
    (item) => item.indexInSent === parseInt(idxSent)
  );
  const canonWordForm: string = thisWord ? thisWord.wordform.toLowerCase() : "";
  const canonWordLength: number = thisWord ? thisWord.length : 1;
  const thisField = inputFields.value[idxSent];

  // When the user presses spacebar, re-play sound
  if (inputContent.includes(" ")) {
    inputValues.value[idxSent] = inputContent.replace(" ", "");
    playPauseCounter.value++;
    return null;
  }

  // check if user input is contained in wordform
  // - if not: access ref and apply `word-user-input-wrong` class
  if (!canonWordForm.includes(inputContent)) {
    thisField.classList.add("word-user-input-wrong");
  } else {
    thisField.classList.remove("word-user-input-wrong");
  }

  // check if the input length is as long as data-length, if so focus next.
  if (
    inputContent.length >= canonWordLength &&
    !inputContent.includes("˚") &&
    !inputContent.includes("¨")
  ) {
    focusNextInput(idxSent);
  }

  if (inputContent.length > 0) {
    thisField.classList.remove("word-user-input-missing");
  }
};

const onUpHandler = (ev) => {
  const idxSent: string = ev.target.attributes["data-index"].value;
  focusPrevInput(idxSent);
};

const onDownHandler = (ev) => {
  const idxSent: string = ev.target.attributes["data-index"].value;
  focusNextInput(idxSent);
};

onMounted(() => trySubmitSent());
</script>

<template>
  <div class="div-review-sent">
    <div>
      <play-pause-rev
        :media-uri="props.mediaUri"
        :parent-play-pause="playPauseCounter"
        :dimension="150"
      />
    </div>
    <!-- <div>
      {{ props.mediaUri }}
      {{ props.sentWords.length }}
      props.sentId: {{ props.sentId }}
      {{ props.finnsNext }}
    </div> -->
    <template v-if="props.sentWords && props.sentWords.length > 0">
      <div class="rev-sentence-input">
        <span
          v-for="word in props.sentWords"
          :key="`${word.indexInSent}` + word.wordform"
          class="learn-sent-word"
        >
          <span
            v-if="
              !word.isCloze ||
              (word.lastInputScore && word.lastInputScore > 0.998)
            "
            style="font-weight: bold"
          >
            {{ word.wordform }}
          </span>
          <input
            v-else
            :data-length="word.length"
            :data-index="word.indexInSent"
            class="word-user-input"
            v-model="inputValues[word.indexInSent]"
            :ref="(el) => (inputFields[word.indexInSent] = el)"
            @keyup.enter="trySubmitSent"
            @keyup.up="onUpHandler"
            @keyup.down="onDownHandler"
            @input="onInputHandler"
            :style="{ width: `${word.length * 1.3}rem` }"
            autocomplete="off"
          />
        </span>
        <span v-if="isComplete" class="small-finish-tick">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
            <path
              d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM371.8 211.8C382.7 200.9 382.7 183.1 371.8 172.2C360.9 161.3 343.1 161.3 332.2 172.2L224 280.4L179.8 236.2C168.9 225.3 151.1 225.3 140.2 236.2C129.3 247.1 129.3 264.9 140.2 275.8L204.2 339.8C215.1 350.7 232.9 350.7 243.8 339.8L371.8 211.8z"
            />
          </svg>
        </span>
      </div>
    </template>
    <template v-if="props.finnsNext">
      <arrow-blue-rev @click="trySubmitSent" />
    </template>
    <template v-else>
      <button @click="trySubmitSent" class="rev-finish-button">Finish</button>
    </template>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.div-review-sent {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 10px;
  max-width: 700px;
}

.rev-sentence-input {
  font-size: 1.2rem;
  color: $blue-primary;
  text-align: center;
  line-height: 2.2rem;
  margin: 80px 20px;

  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
}

.learn-sent-word {
  margin: 0.1rem 0.2em;
}

.word-user-input {
  border: 1px solid $azure-primary;
  box-shadow: 2px 2px $blue-secondary;
  background-color: $azure-secondary;
  height: 1.8rem;
  color: $blue-primary;
  font-size: 1.2rem;
}

.word-user-input-missing {
  box-shadow: 0px 4px $yellow-gold;
}

.word-user-input-wrong {
  box-shadow: 0px 4px $red-secondary;
}

.small-finish-tick {
  margin-left: 10px;
  margin-right: 10px;
  svg {
    width: 18px;
    height: 18px;
    path {
      fill: $blue-secondary;
    }
  }
}

.rev-finish-button {
  text-align: center;
  background-color: $blue-primary;
  border: 2px solid $yellow-canary;
  cursor: pointer;
  padding: 0rem 3rem;
  height: 3rem;
  margin-left: 2rem;
  margin-right: 2rem;
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
  text-decoration: none;
  color: $yellow-gold;
  font-size: 1.5rem;
  font-weight: bold;
}
</style>
