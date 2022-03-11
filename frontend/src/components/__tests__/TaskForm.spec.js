import { describe, it, expect } from "vitest";

import { mount } from "@vue/test-utils";

import TaskForm from "../TaskForm.vue";

describe("TaskForm", () => {
    it("renders properly", () => {
        const wrapper = mount(TaskForm, {
            props: { modelValue: [] },
            directives: {
                cy() {
                    /* stub */
                },
            },
        });
        expect(wrapper.text()).toContain("Add");
    });
});
