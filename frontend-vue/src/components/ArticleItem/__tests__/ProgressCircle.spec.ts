import { describe, it, expect } from "vitest";

import { mount } from "@vue/test-utils";
import ProgressCircle from "../ProgressCircle.vue";

describe("ProgressCircle", () => {

  const testProps = [
    {primary: "#ffcc00", secondary: "#00ccff", prog: 1, goal: 10},
    {primary: "#ffcc00", secondary: "#00ccff", prog: -10, goal: 10},
    {primary: "#ffcc00", secondary: "#00ccff", prog: 12, goal: 10}
  ]

  it("renders properly, displays prog/goal", ()=>{
    const wrapper = mount(ProgressCircle, {props: testProps[0]})
    expect(wrapper.text()).toContain("1/10")
    expect(wrapper).toMatchSnapshot()
  });

  it.todo("when prog is negative, displays 0/goal", ()=>{
    const wrapper = mount(ProgressCircle, {props: testProps[1]})
    expect(wrapper.text()).toContain("0/10")
    expect(wrapper).toMatchSnapshot()
  });

  it.todo("when prog is greater than goal, displays goal/goal", ()=>{
    const wrapper = mount(ProgressCircle, {props: testProps[2]})
    expect(wrapper.text()).toContain("10/10")
    expect(wrapper).toMatchSnapshot()
  });
});
