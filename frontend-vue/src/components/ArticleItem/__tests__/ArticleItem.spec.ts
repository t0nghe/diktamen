import { describe, it, expect } from "vitest";
import { nextTick } from "vue";

import { mount } from "@vue/test-utils";
import ArticleItem from "../ArticleItem.vue";

describe("ArticleItem: in-progress article", () => {
  const testData = {
    articleId: 6,
    articleTitle: "Zelenskyj är en man som kvinnor längtar efter",
    articleSentCount: 17,
    articleDescription:
      "Nr han nu vl tar plats, den traditionelle mannen, vcker han bde kvinnors tr och mns nskan att vara som han",
    userFinishedIndex: 3,
  };

  const wrapper = mount(ArticleItem, {
    props: {
      id: testData.articleId,
      title: testData.articleTitle,
      description: testData.articleDescription,
      progress: testData.userFinishedIndex,
      goal: testData.articleSentCount,
    },
  });

  it("render article title and info correctly", () => {
    expect(wrapper.text()).toContain(
      "Zelenskyj är en man som kvinnor längtar efter"
    );
    expect(wrapper.text()).toContain("den traditionelle mannen");
  });

  it("in-progress articles should display progress circles with currentIndex/goal", () => {
    expect(wrapper.text()).toContain(
      `${testData.userFinishedIndex}/${testData.articleSentCount}`
    );
  });

  it("in-progress articles should have classnames that start with progress", () => {
    expect(wrapper.classes()).toContain("article-item-progress");
    expect(wrapper.classes("article-item-new")).toBe(false);
    expect(wrapper.classes("article-item-complete")).toBe(false);
  });

  it("clicking on an in-progress article, go-to-article event should emitted", async () => {
    const title = wrapper.find(".article-item__title");
    title.trigger("click");
    expect(wrapper.emitted()).toHaveProperty("click");
    await nextTick();
    expect(wrapper.emitted()).toHaveProperty("go-to-article");
    expect(wrapper.emitted()["go-to-article"][0]).toEqual([testData.articleId]);
  });
});

describe("ArticleItem: new article", () => {
  const testData = {
    articleId: 6,
    articleTitle: "Zelenskyj är en man som kvinnor längtar efter",
    articleSentCount: 17,
    articleDescription:
      "Nr han nu vl tar plats, den traditionelle mannen, vcker han bde kvinnors tr och mns nskan att vara som han",
    userFinishedIndex: 0,
  };

  const wrapper = mount(ArticleItem, {
    props: {
      id: testData.articleId,
      title: testData.articleTitle,
      description: testData.articleDescription,
      progress: testData.userFinishedIndex,
      goal: testData.articleSentCount,
    },
  });

  it("render article title and info correctly", () => {
    expect(wrapper.text()).toContain(
      "Zelenskyj är en man som kvinnor längtar efter"
    );
    expect(wrapper.text()).toContain("den traditionelle mannen");
  });

  it("new articles should have classnames that start with new", () => {
    expect(wrapper.classes()).toContain("article-item-new");
    expect(wrapper.classes("article-item-progress")).toBe(false);
    expect(wrapper.classes("article-item-complete")).toBe(false);
  });

  it("new articles should display progress circles with 0/goal", () => {
    expect(wrapper.text()).toContain(`0/${testData.articleSentCount}`);
  });

  it("clicking on a new article, go-to-article event should emitted", async () => {
    const title = wrapper.find(".article-item__title");
    title.trigger("click");
    expect(wrapper.emitted()).toHaveProperty("click");
    await nextTick();
    expect(wrapper.emitted()).toHaveProperty("go-to-article");
    expect(wrapper.emitted()["go-to-article"][0]).toEqual([testData.articleId]);
  });
});

describe("ArticleItem: complete", () => {
  const testData = {
    articleId: 6,
    articleTitle: "Zelenskyj är en man som kvinnor längtar efter",
    articleSentCount: 17,
    articleDescription:
      "Nr han nu vl tar plats, den traditionelle mannen, vcker han bde kvinnors tr och mns nskan att vara som han",
    userFinishedIndex: 17,
  };

  const wrapper = mount(ArticleItem, {
    props: {
      id: testData.articleId,
      title: testData.articleTitle,
      description: testData.articleDescription,
      progress: testData.userFinishedIndex,
      goal: testData.articleSentCount,
    },
  });

  it("render article title and info correctly", () => {
    expect(wrapper.text()).toContain(
      "Zelenskyj är en man som kvinnor längtar efter"
    );
    expect(wrapper.text()).toContain("den traditionelle mannen");
  });

  it("complete articles should have classnames that start with complete", () => {
    expect(wrapper.classes()).toContain("article-item-complete");
    expect(wrapper.classes("article-item-new")).toBe(false);
    expect(wrapper.classes("article-item-progress")).toBe(false);
  });

  it("complete articles should display progress circle with goal/goal", () => {
    expect(wrapper.text()).toContain(
      `${testData.articleSentCount}/${testData.articleSentCount}`
    );
  });

  it("clicking on a complete article, no event should be emitted", async () => {
    wrapper.get(".article-item__title").trigger("click");
    await nextTick();
    expect(wrapper.emitted()).toBeNull;
  });
});
