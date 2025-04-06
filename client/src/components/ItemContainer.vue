<script setup lang="ts">
import Icon from "@/components/ui/Icon.vue";
defineProps<{
    items: number[];
}>();

</script>

<template>
    <div class="item-container">
        <Icon :class="{ 'trinket': idx === items.length - 1 }" v-for="(item, idx) in items"
            :icon-src="item ? `https://cdn.nashor.gg/assets/15.7.1/img/item/${item}.png` : ''" />
    </div>
</template>

<style scoped>
.item-container {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: var(--margin-small);
    margin: var(--margin-base);
}

.item-container .trinket {
    grid-row: 1;
    grid-column: 4;
}


@media screen and (min-width: 1024px) {
    .item-container--row {
        grid-template-columns: repeat(7, 1fr);
    }

    .item-container--row .trinket {
        grid-column: 7;
    }
}
</style>
