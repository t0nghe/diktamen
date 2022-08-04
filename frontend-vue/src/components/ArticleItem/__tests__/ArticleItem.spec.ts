import { describe, it, expect } from "vitest";

import { mount } from "@vue/test-utils";
import ArticleItemComplete from "../ArticleItemComplete.vue";
import ArticleItemNew from "../ArticleItemNew.vue";
import ArticleItemProgress from "../ArticleItemProgress.vue";
import ProgressCircle from "../ProgressCircle.vue";

describe("ProgressCircle", () => {
  it.todo("renders properly, displays prog/goal");

  it.todo("takes four props, primary color, secondary color, prog, goal");

  it.todo("when prog is zero, circle is only secondary color, displays 0/goal");

  it.todo("when prog is negative, circle is secondary color, displays 0/goal");

  it.todo("when prog equals goal, circle is primary color");
  // presumably we can `getComputedStyle` and see the rotation

  it.todo(
    "when prog is greater than goal, circle is primary color, displays goal/goal"
  );
});

describe("ArticleItemComplete", () => {
  it("renders properly", () => {
    const wrapper = mount(ArticleItemComplete, {
      props: { msg: "Hello Vitest" },
    });
    expect(wrapper.text()).toContain("Hello Vitest");
  });
});

describe("ArticleItemNew", () => {
  it("renders properly", () => {
    const wrapper = mount(ArticleItemNew, { props: { msg: "Hello Vitest" } });
    expect(wrapper.text()).toContain("Hello Vitest");
  });
});

describe("ArticleItemProgress", () => {
  it("renders properly", () => {
    const wrapper = mount(ArticleItemProgress, {
      props: { msg: "Hello Vitest" },
    });
    expect(wrapper.text()).toContain("Hello Vitest");
  });
});
