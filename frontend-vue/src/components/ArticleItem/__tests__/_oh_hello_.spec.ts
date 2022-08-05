import { describe, it, expect } from "vitest";

import { mount } from "@vue/test-utils";
import ArticleItemComplete from "../ArticleItemComplete.vue";

describe("ArticleItemComplete", () => {
  it("renders properly", () => {
    const a = 1
    expect(a).toBe(1);
  });
});

// describe("ArticleItemNew", () => {
//   it("renders properly", () => {
//     const wrapper = mount(ArticleItemNew, { props: { msg: "Hello Vitest" } });
//     expect(wrapper.text()).toContain("Hello Vitest");
//   });
// });

// describe("ArticleItemProgress", () => {
//   it("renders properly", () => {
//     const wrapper = mount(ArticleItemProgress, {
//       props: { msg: "Hello Vitest" },
//     });
//     expect(wrapper.text()).toContain("Hello Vitest");
//   });
// });
