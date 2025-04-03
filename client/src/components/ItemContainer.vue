<script setup lang="ts">
import Icon from "@/components/ui/Icon.vue";
const props = defineProps<{
    items: number[];
    foreground: boolean;
}>();

const itemSources = props.items.map((item) =>
    item ? `https://cdn.nashor.gg/assets/15.7.1/img/item/${item}.png` : ""
);
</script>

<template>
    <div class="item-container">
        <Icon :class="{ 'icon--foreground': foreground }" v-for="itemSrc in itemSources" :icon-src="itemSrc" />
    </div>
</template>

<style scoped>
.item-container {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: var(--margin-small);
    margin: var(--margin-base);
}
</style>
