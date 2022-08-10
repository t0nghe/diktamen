import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import TopBar from "../TopBar.vue";

describe("topbar renders correctly", () => {
  it("waitlist: style and anchor are correct", () => {
    const wrapper = mount(TopBar, {
      props: { state: "waitlist", ahref: "https://dummy.com" },
    });

    const anchor = wrapper.find("a");
    expect(wrapper.text()).toContain("join waitlist");
    expect(wrapper.classes()).toContain("top-bar-waitlist");
    expect(anchor.element.href).toEqual("https://dummy.com/");
  });

  it("articles: style and heading are correct", () => {
    const wrapper = mount(TopBar, {
      props: { state: "articles" },
    });

    expect(wrapper.text()).toContain("learning progress");
    expect(wrapper.classes()).toContain("top-bar-articles");
  });

  it("summary: style and heading are correct", () => {
    const wrapper = mount(TopBar, {
      props: { state: "summary" },
    });

    expect(wrapper.text()).toContain("summary");
    expect(wrapper.classes()).toContain("top-bar-summary");
  });

  it("learn: style, heading and article title are correct", () => {
    const wrapper = mount(TopBar, {
      props: { state: "learn", title: "Men varför inte?" },
    });

    expect(wrapper.text()).toContain("learning:");
    expect(wrapper.text()).toContain("Men varför inte?");
    expect(wrapper.classes()).toContain("top-bar-learn");
  });

  it("review: style, heading, progress are correct", () => {
    const wrapper = mount(TopBar, {
      props: { state: "review", revcount: 0, duecount: 10 },
    });

    expect(wrapper.text()).toContain("review");
    expect(wrapper.text()).toContain("0/10");
    expect(wrapper.classes()).toContain("top-bar-review");
  });
});
