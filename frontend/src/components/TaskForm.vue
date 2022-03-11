<script setup>
import axios from "axios";
import { ref, toRefs } from "vue";

const props = defineProps({
    modelValue: {
        type: Array,
        required: true,
    },
});
const { modelValue } = toRefs(props);
const emit = defineEmits(["update:modelValue"]);
const name = ref("");

async function onSubmit() {
    try {
        const formData = new FormData();
        formData.append("name", name.value);
        const response = await axios.post("/api/v1/todo/create", formData);
        emit("update:modelValue", [...modelValue.value, response.data.data]);
        name.value = "";
    } catch (error) {
        console.warn(error);
    }
}
</script>

<template>
    <form
        class="flex w-full items-center bg-gray-100 p-3 rounded shadow mb-4"
        @submit.prevent="onSubmit"
    >
        <input
            v-model="name"
            v-cy="'form-input'"
            type="text"
            placeholder="Type here your task"
            class="px-3 py-2 border border-blue-400 grow outline-none rounded-tl rounded-bl"
            required
        />
        <button
            type="submit"
            v-cy="'form-submit'"
            class="w-20 self-stretch bg-blue-400 text-white shrink-0 rounded-tr rounded-br"
        >
            Add
        </button>
    </form>
</template>
