import { describe, it, expect, beforeAll, afterAll, vi } from "vitest";
import { nextTick } from "vue";

import { mount } from "@vue/test-utils";
import NavSidebar from "../NavSidebar.vue";

describe("sidebar renders correctly when logged in", () => {
  beforeAll(() => {
    window.localStorage.setItem("token", "<tokenstring>");
  });

  it("sidebar renders correctly", () => {
    const wrapper = mount(NavSidebar);
    expect(wrapper.find("#sidebar-control-area").exists()).toBe(true);
    expect(wrapper.find("#sidebar-display-area").exists()).toBe(true);
  });

  it("in wide mode, labels are rendered for naviation and an icon is rendered for logging out", () => {
    const wrapper = mount(NavSidebar, { props: { wide: true } });
    expect(wrapper.find("#sidebar-nav-learn").text()).toContain("learn");
    expect(wrapper.find(".svg-nav-learn").exists()).toBe(true);
    expect(wrapper.find("#sidebar-nav-review").text()).toContain("review");
    expect(wrapper.find(".svg-nav-review").exists()).toBe(true);
    expect(wrapper.find("#sidebar-button-logout").exists()).toBe(true);
    expect(wrapper.find(".svg-button-logout").exists()).toBe(true);
  });

  it("in narrow mode, 3 icons are rendered for naviation & 1 icon is rendered for logging out", () => {
    const wrapper = mount(NavSidebar, { props: { wide: false } });
    expect(wrapper.find("#sidebar-nav-learn").text()).not.toContain("learn");
    expect(wrapper.find(".svg-nav-learn").exists()).toBe(true);
    expect(wrapper.find("#sidebar-nav-review").text()).not.toContain("review");
    expect(wrapper.find(".svg-nav-review").exists()).toBe(true);
    expect(wrapper.find("#sidebar-button-logout").exists()).toBe(true);
    expect(wrapper.find(".svg-button-logout").exists()).toBe(true);
  });

  it("clicking #sidebar-control-area will turn .sidebar-wide to .sidebar-narrow", async () => {
    const wrapper = mount(NavSidebar, { props: { wide: true } });
    wrapper.get("#sidebar-control-area").trigger("click");
    await nextTick();
    expect(wrapper.find(".sidebar-narrow").exists()).toBe(true);
    expect(wrapper.find(".sidebar-wide").exists()).toBe(false);
  });

  it("clicking #sidebar-control-area will turn .sidebar-narrow to .sidebar-wide", async () => {
    const wrapper = mount(NavSidebar, { props: { wide: false } });
    wrapper.get("#sidebar-control-area").trigger("click");
    await nextTick();
    expect(wrapper.find(".sidebar-narrow").exists()).toBe(false);
    expect(wrapper.find(".sidebar-wide").exists()).toBe(true);
  });

  // it("clicking learn on nav sidebar, router.push is called with /articles", async () => {
  //   const mockRouter = { push: vi.fn() };
  //   const wrapper = mount(NavSidebar, {
  //     props: { wide: false },
  //     global: {
  //       mocks: {
  //         $router: mockRouter,
  //       },
  //     },
  //   });

  //   wrapper.get("#sidebar-nav-learn").trigger("click");
  //   await nextTick();
  //   expect(mockRouter.push).toHaveBeenCalledTimes(1);
  //   expect(mockRouter.push).toHaveBeenCalledWith("/articles");
  // });

  // it("clicking review on nav sidebar, router.push is called with /review", async () => {
  //   const mockRouter = { push: vi.fn() };
  //   const wrapper = mount(NavSidebar, {
  //     props: { wide: false },
  //     global: {
  //       mocks: {
  //         $router: mockRouter,
  //       },
  //     },
  //   });

  //   wrapper.get("#sidebar-nav-review").trigger("click");
  //   await nextTick();
  //   expect(mockRouter.push).toHaveBeenCalledTimes(1);
  //   expect(mockRouter.push).toHaveBeenCalledWith("/review");
  // });

  // how to test there's no token
  it("clicking #sidebar-button-logout, router goes to / and token is cleared", async () => {
    // const mockRouter = { push: vi.fn() };
    const wrapper = mount(NavSidebar, {
      props: { wide: false },
      // global: {
      //   mocks: {
      //     $router: mockRouter,
      //   },
      // },
    });
    wrapper.get("#sidebar-button-logout").trigger("click");
    await nextTick();
    // expect(mockRouter.push).toHaveBeenCalledTimes(1);
    // expect(mockRouter.push).toHaveBeenCalledWith("/");
    const token = window.localStorage.getItem("token");
    expect(token).toBe(null);
  });

  afterAll(() => {
    window.localStorage.removeItem("token");
  });
});

describe("sidebar renders correctly when not logged in", () => {
  beforeAll(() => {
    window.localStorage.removeItem("token");
  });

  it("sidebar renders correctly in wide mode", () => {
    const wrapper = mount(NavSidebar, { props: { wide: true } });
    expect(wrapper.find("#sidebar-control-area").exists()).toBe(true);
    expect(wrapper.find("#sidebar-display-area").exists()).toBe(true);

    expect(wrapper.find("#sidebar-nav-learn").exists()).toBe(false);
    expect(wrapper.find("#sidebar-nav-review").exists()).toBe(false);
    expect(wrapper.find("#sidebar-button-logout").exists()).toBe(false);

    expect(wrapper.find("#sidebar-button-login").exists()).toBe(true);
    expect(wrapper.find("#sidebar-button-login").text()).toContain("log in");
  });

  it("sidebar renders correctly in narrow mode", () => {
    const wrapper = mount(NavSidebar, { props: { wide: false } });
    expect(wrapper.find("#sidebar-control-area").exists()).toBe(true);
    expect(wrapper.find("#sidebar-display-area").exists()).toBe(true);

    expect(wrapper.find("#sidebar-nav-learn").exists()).toBe(false);
    expect(wrapper.find("#sidebar-nav-review").exists()).toBe(false);
    expect(wrapper.find("#sidebar-button-logout").exists()).toBe(false);

    expect(wrapper.find("#sidebar-button-login").exists()).toBe(true);
    expect(wrapper.find("#sidebar-button-login").text()).not.toContain(
      "log in"
    );
    expect(wrapper.find(".svg-button-login").exists()).toBe(true);
  });

  // it("clicking #sidebar-button-login router goes to /login", async () => {
  //   const mockRouter = { push: vi.fn() };
  //   const wrapper = mount(NavSidebar, {
  //     props: { wide: false },
  //     global: {
  //       mocks: { $router: mockRouter },
  //     },
  //   });
  //   wrapper.find("#sidebar-button-login").trigger("click");
  //   await nextTick();
  //   expect(mockRouter.push).toHaveBeenCalledTimes(1);
  //   expect(mockRouter.push).toHaveBeenCalledWith("/login");
  // });
  
  // it("clicking #sidebar-button-login router goes to /login", async () => {
  //   const mockRouterPush = vi.fn();
  //   vi.mock("vue-router", () => ({
  //     useRouter: () => ({
  //       push: mockRouterPush,
  //     }),
  //   }));

  //   const wrapper = mount(NavSidebar, {
  //     props: { wide: false },
  //   });
  //   wrapper.find("#sidebar-button-login").trigger("click");
  //   await nextTick();
  //   expect(mockRouterPush).toHaveBeenCalledTimes(1);
  //   expect(mockRouterPush).toHaveBeenCalledWith("/login");
  // });
});
