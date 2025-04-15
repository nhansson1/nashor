<script setup lang="ts">
import { getItemDataById } from "@/utils/league-utils";
import IconWithTooltip from "./IconWithTooltip.vue";
const props = defineProps<{
    items: number[];
}>();

interface ItemData {
    price: number;
    name: string;
    src: string;
    content: string;
}

const itemData: ItemData[] = props.items.map(itemId => {
    const itemData = { price: 0, name: "", src: "", content: "" };

    if (!itemId)
        return itemData;

    const itemObj = getItemDataById(itemId);

    if (!itemObj)
        return itemData;

    itemData.price = itemObj.gold.total;
    itemData.name = itemObj.name;
    itemData.content = itemObj.description;
    itemData.src = `${import.meta.env.VITE_ASSETS_BASE}/img/item/${itemId}.png`;

    return itemData;
});

</script>

<template>
    <div class="item-container">
        <IconWithTooltip v-for="(item, idx) in itemData" :title="item.name" :bread="item.price.toString()"
            :icon-src="item.src" :content="item.content" :class="{ 'trinket': idx === itemData.length - 1 }" />
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
