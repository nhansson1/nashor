<script setup lang="ts">
import { useTemplateRef } from "vue";
import Icon from "./ui/Icon.vue";
defineProps<{
  imageSrc?: string;
  title: string;
  bread: string;
  content: string;
  show: boolean;
}>();

const tooltip = useTemplateRef("tooltip");
defineExpose({
  tooltip,
});
</script>

<template>
  <div ref="tooltip" v-if="show" class="tooltip">
    <div class="tooltip__header">
      <Icon :icon-src="imageSrc" />
      <div>
        <p class="tooltip__title">{{ title }}</p>
        <p class="tooltip__bread">{{ bread }}</p>
      </div>
    </div>
    <div v-html="content"></div>
  </div>
</template>

<style scoped>
.tooltip {
  position: absolute;
  z-index: 10;
  min-width: 10rem;
  background-color: var(--foreground);
  padding: var(--margin-base);
  border-radius: var(--rounding-base);
  left: 50%;
  transform: translateX(-50%);
  width: clamp(10rem, 15vw, 20rem);
  color: #ffffff87;
  font-size: clamp(0.6rem, 1vw, 0.75rem);
}

.tooltip--top {
  bottom: 100%;
}

.tooltip--bottom {
  top: 100%;
}

.tooltip--right {
  transform: translateX(-100%);
}

.tooltip__header {
  display: flex;
  align-items: center;
  gap: var(--margin-base);
  margin-bottom: var(--margin-base);
}

.tooltip__title {
  color: #ffffff;
}
</style>
