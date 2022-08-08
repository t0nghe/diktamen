import { describe, it, expect, beforeAll, afterAll } from "vitest";
// import { nextTick } from "vue";

import { mount } from "@vue/test-utils";
import NavSidebar from "../NavSidebar.vue";

describe("if there's a token, sidebar renders correctly", () => {
  beforeAll(() => {
    window.localStorage.setItem("token", "<tokenstring>");
  });

  it("sidebar renders correctly", () => {
    const wrapper = mount(NavSidebar);
    expect(wrapper.text()).toContain("side bar initialized");
  });

  it.todo("sidebar #sidebar-control-area is rendered");

  it.todo("sidebar #sidebar-display-area is rendered");

  it.todo(
    "if #sidebar-display-area .sidebar-narrow, 3 svgs for nav and 1 svg for logging out are rendered"
  );

  it.todo(
    "if #sidebar-display-area .sidebar-wide, 3 strings for nav & 1 svg for logging out are rendered"
  );

  it.todo("clicking #sidebar-nav-list, router goes to /articles/");

  it.todo("clicking #sidebar-nav-learn, router goes to /learn/");

  it.todo("clicking #sidebar-nav-review, router goes to /review/");

  it.todo(
    "clicking #sidebar-button-logout, router goes to / and token is cleared"
  );

  it.todo(
    "clicking #sidebar-control-area will toggle .sidebar-wide vs .sidebar-narrow"
  );

  afterAll(()=> {
    window.localStorage.removeItem("token");
  });
});

describe("if there's not a token, sidebar renders correctly when not logged in", () => {
  beforeAll(() => {
    window.localStorage.removeItem("token");
  });

  it.todo("#sidebar-control-area is rendered");

  it.todo("#sidebar-display-area is rendered");

  it.todo(
    "when #sidebar-control-area is .sidebar-narrow, 1 svg icon for logging in is rendered"
  );

  it.todo(
    "when #sidebar-control-area is .sidebar-wide, 1 svg and a string is rendered"
  );

  it.todo("clicking #sidebar-button-login router goes to /login");

  it.todo(
    "clicking #sidebar-control-area will toggle .sidebar-wide vs .sidebar-narrow"
  );
});
