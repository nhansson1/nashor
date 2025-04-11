<script setup lang="ts">
import Icon from './ui/Icon.vue';
import Tooltip from './Tooltip.vue';
import { nextTick, ref, useTemplateRef } from 'vue';

defineProps<{ iconSrc: string, content: string, title: string, bread: string }>()
const showTooltip = ref(false);

const tooltipClass = ref({ 'tooltip--hide': true, 'tooltip--top': true, 'tooltip--bottom': false, 'tooltip--right': false });
const tooltipRef = useTemplateRef<InstanceType<typeof Tooltip>>("tooltip-ref");

const handleMouseEnter = () => {
    showTooltip.value = true;

    nextTick(() => {
        const rect = tooltipRef.value?.tooltip?.getBoundingClientRect();

        if (rect) {
            const withinTop = rect.top > 0;
            console.log(rect);

            tooltipClass.value = { 'tooltip--hide': false, 'tooltip--top': withinTop, 'tooltip--bottom': !withinTop, 'tooltip--right': rect.right > window.innerWidth };
        }
    });
};

const handleMouseLeave = () => {
    showTooltip.value = false;
    tooltipClass.value = { 'tooltip--hide': true, 'tooltip--top': true, 'tooltip--bottom': false, 'tooltip--right': false };
};
</script>

<template>
    <div class="item-tooltip-container">
        <Icon @m-enter="handleMouseEnter" @m-leave="handleMouseLeave" :icon-src="iconSrc" />
        <Tooltip ref="tooltip-ref" :class="tooltipClass" :image-src="iconSrc" :content="content" :title="title"
            :bread="bread" :show="showTooltip" />
    </div>
</template>

<style scoped>
.item-tooltip-container {
    position: relative;
}
</style>
