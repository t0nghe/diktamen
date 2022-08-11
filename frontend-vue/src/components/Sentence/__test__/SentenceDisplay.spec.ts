import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import {
  examineCorrectSents as IncorrSents,
  examineCorrectSents as CorrSents,
  displayUnseenSents as UnseenSents,
} from "../mockData";

import SentenceNew from "../SentenceNew.vue";
import SentenceTried from "../SentenceTried.vue";

describe("correct seen sents displayed correctly", () => {
  const correctSeenSent = CorrSents[0];
  const wrapper = mount(SentenceTried, {
    props: { ...correctSeenSent, isCorrect: true },
  });

  it("style of tried sentences", () => {
    // This classname to distringuish from sentences shown on summary page.
    expect(wrapper.classes()).toContain("learn-sent-tried");
  });

  it("indexInArticle: styling and content correct", () => {
    const iiaSpan = wrapper.find(".learn-sent-index-correct");
    expect(iiaSpan).toBeDefined();
    expect(iiaSpan.text()).toEqual("3");
  });

  it("text: styling and content correct", () => {
    const tryTextSpan = wrapper.find(".learn-sent-text");
    expect(tryTextSpan).toBeDefined();
    expect(tryTextSpan.text()).toEqual(correctSeenSent.tryText);
  });
});

describe("incorrect seen sents displayed correctly", () => {
  const incorrectSeenSent = IncorrSents[0];
  const wrapper = mount(SentenceTried, {
    props: { ...incorrectSeenSent, isCorrect: false },
  });

  it("indexInArticle: styling and content correct", () => {
    const iiaSpanIncorr = wrapper.find(".learn-sent-index-incorrect");
    expect(iiaSpanIncorr).toBeDefined();
    expect(iiaSpanIncorr).toEqual("1");
  });

  it("incorrect words are highlighted", () => {
    const errorWords = wrapper.findAll(".word-error");
    const canonicalWords = wrapper.findAll(".word-canonical");
    // .token-normal would be used to display tokens the user got right.
    expect(errorWords[0]).toEqual("Dxt");
    expect(canonicalWords[0]).toEqual("Det");
    expect(errorWords[1]).toEqual("xr");
    expect(canonicalWords[1]).toEqual("är");
    expect(errorWords[4]).toEqual("melxxx");
    expect(canonicalWords[4]).toEqual("mellan");
  });
});

describe("unseen sents displayed correctly", () => {
  const sentId20 = UnseenSents[2];
  const wrapper = mount(SentenceNew, { props: { ...sentId20 } });

  it("styling and content of sent correct", () => {
    expect(wrapper.classes()).toContain("learn-new-sent");
    expect(wrapper.text()).not.toContain("har");
    expect(wrapper.text()).not.toContain("Det");
    expect(wrapper.text()).toContain(",");
    expect(wrapper.text()).toContain(".");
    expect(wrapper.text()).toContain("_________");
  });
});
