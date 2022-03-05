<script setup>
import axios from "axios";
import { ref, computed, onMounted } from "vue";

const list = ref([]);
const name = ref("");

const orderedList = computed(() => {
    const done = [];
    const undone = [];
    for (const item of list.value) {
        if (item.status === "undone") {
            undone.push(item);
        } else {
            done.push(item);
        }
    }
    return [...undone, ...done];
});

onMounted(() => {
    axios
        .get("/api/v1/todo/list")
        .then((response) => {
            list.value = response.data.data;
        })
        .catch((error) => {
            console.warn(error.message);
        });
});

async function onSubmit() {
    try {
        const formData = new FormData();
        formData.append("name", name.value);
        const response = await axios.post("/api/v1/todo/create", formData);
        list.value.push(response.data.data);
        name.value = "";
    } catch (error) {
        console.warn(error.mesage);
    }
}

async function remove(name) {
    try {
        const index = list.value.findIndex((item) => item.name === name);
        const formData = new FormData();
        formData.append("name", list.value[index].name);
        await axios.post("/api/v1/todo/delete", formData);
        list.value.splice(index, 1);
    } catch (error) {
        console.warn(error.message);
    }
}

async function toggleState(name) {
    try {
        const index = list.value.findIndex((item) => item.name === name);
        const newStatus =
            list.value[index].status === "done" ? "undone" : "done";
        const formData = new FormData();
        formData.append("name", list.value[index].name);
        formData.append("status", newStatus);
        await axios.post("/api/v1/todo/change-status", formData);
        list.value[index].status = newStatus;
    } catch (error) {
        console.warn(error.message);
    }
}
</script>

<template>
    <div class="w-full min-h-screen h-full bg-blue-200">
        <div class="p-4 mx-auto w-120">
            <form
                @submit.prevent="onSubmit"
                class="flex w-full items-center bg-gray-100 p-3 rounded shadow mb-4"
            >
                <input
                    v-model="name"
                    type="text"
                    placeholder="Type here your task"
                    class="px-3 py-2 border border-blue-400 grow outline-none rounded-tl rounded-bl"
                    required
                />
                <button
                    type="submit"
                    class="w-20 self-stretch bg-blue-400 text-white shrink-0 rounded-tr rounded-br"
                >
                    Add
                </button>
            </form>
            <div class="bg-gray-100 p-4 rounded shadow">
                <TransitionGroup name="list" tag="div">
                    <div
                        v-for="item in orderedList"
                        :key="item.name"
                        class="p-2"
                    >
                        <div
                            class="px-3 py-2 bg-white shadow rounded cursor-pointer select-none"
                            @dblclick="remove(item.name)"
                        >
                            <p
                                :class="{
                                    'line-through': item.status === 'done',
                                }"
                            >
                                {{ item.name }}
                            </p>
                            <div class="flex mt-2">
                                <button
                                    v-if="item.status === 'done'"
                                    class="px-2 py-0.5 ml-auto text-sm bg-red-500 text-white shadow-sm rounded-sm"
                                    @dblclick.stop
                                    @click.stop="toggleState(item.name)"
                                >
                                    Undone
                                </button>

                                <button
                                    v-else
                                    class="px-2 py-0.5 ml-auto text-sm bg-green-500 text-white shadow-sm rounded-sm"
                                    @dblclick.stop
                                    @click.stop="toggleState(item.name)"
                                >
                                    Done
                                </button>
                            </div>
                        </div>
                    </div>
                </TransitionGroup>

                <div v-if="orderedList.length === 0">
                    <p
                        class="text-sm uppercase text-center text-gray-400 select-none"
                    >
                        No tasks added yet
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.list-enter-active,
.list-leave-active {
    transition: all 0.3s ease;
}

.list-move {
    transition: all 0.3s ease;
}

.list-enter-from,
.list-leave-to {
    opacity: 0;
    transform: translateX(30px);
}
</style>
