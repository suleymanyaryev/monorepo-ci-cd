import { describe, it, expect } from "vitest";

import { mount } from "@vue/test-utils";

import TaskForm from "../TaskForm.vue";

describe("TaskForm", () => {
    it("renders properly", () => {
        const wrapper = mount(TaskForm, {
            props: { modelValue: [] },
        });
        expect(wrapper.text()).toContain("Add");
    });
});
