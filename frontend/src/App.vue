<script setup>
import axios from "axios";
import { ref, computed, onMounted } from "vue";
import TaskForm from "./components/TaskForm.vue";
import TaskItem from "./components/TaskItem.vue";

const list = ref([]);

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
            <TaskForm v-model="list" />
            <div v-cy="'todo-wrapper'" class="bg-gray-100 p-4 rounded shadow">
                <TransitionGroup name="list" tag="div">
                    <TaskItem
                        v-for="item in orderedList"
                        v-cy="'todo-item'"
                        :item="item"
                        :key="item.name"
                        @remove="remove"
                        @toggle-state="toggleState"
                    />
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
