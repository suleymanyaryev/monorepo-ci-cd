import { describe, it, expect } from "vitest";

import { mount } from "@vue/test-utils";

import TaskItem from "../TaskItem.vue";

describe("TaskItem", () => {
    it("renders properly", () => {
        const wrapper = mount(TaskItem, {
            props: {
                item: { name: "task", status: "done" },
            },
            directives: {
                cy() {
                    /* stub */
                },
            },
        });
        expect(wrapper.text()).toContain("task");
    });
});
