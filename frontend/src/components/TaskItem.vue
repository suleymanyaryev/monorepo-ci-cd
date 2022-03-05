<script setup>
defineProps({
    item: {
        type: Object,
        required: true,
    },
});

defineEmits(["remove", "toggle-state"]);
</script>

<template>
    <div class="p-2">
        <div
            class="px-3 py-2 bg-white shadow rounded cursor-pointer select-none"
            @dblclick="$emit('remove', item.name)"
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
                    @click.stop="$emit('toggle-state', item.name)"
                >
                    Undone
                </button>

                <button
                    v-else
                    class="px-2 py-0.5 ml-auto text-sm bg-green-500 text-white shadow-sm rounded-sm"
                    @dblclick.stop
                    @click.stop="$emit('toggle-state', item.name)"
                >
                    Done
                </button>
            </div>
        </div>
    </div>
</template>
