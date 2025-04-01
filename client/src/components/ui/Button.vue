<script setup lang="ts">
defineProps<{ text: string }>()
</script>

<template>
    <button class="button" @click="$emit('buttonClick')">
        {{ text }}
    </button>
</template>

<style scoped>
.button {
    border-radius: var(--rounding-base);
    padding: var(--margin-base);
    background-color: var(--background);
    color: #ffffff;
    border: none;
    outline: none;
    max-width: fit-content;
    cursor: pointer;
}
</style>
