<script setup lang="ts">
import Container from './ui/Container.vue';
defineProps<{ imageSrc?: string, title: string, bread: string, content: string, show: boolean }>();
</script>

<template>
    <Container v-if="show" class="tooltip">
        <div class="tooltip__header">
            <img class="tooltip__image" v-if="imageSrc" :src="imageSrc" alt="">
            <div>
                <p class="tooltip__title">{{ title }}</p>
                <p class="tooltip__bread">{{ bread }}</p>
            </div>
        </div>
        <div v-html="content"></div>
    </Container>
</template>

<style scoped>
.tooltip {
    position: absolute;
    right: 50%;
    transform: translateX(-25%);
}

.tooltip--top {
    bottom: 100%;
}

.tooltip--bottom {
    top: 100%;
}

.tooltip__header {
    display: flex;
}

.tooltip__image {
    margin-right: var(--margin-base);
}

.tooltip__title {
    color: #ffffff;
}

.tooltip__bread {
    color: #ffffff87;
}
</style>
