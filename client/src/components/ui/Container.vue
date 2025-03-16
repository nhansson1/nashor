<template>
  <div class="container">
    <slot></slot>
  </div>
</template>

<style scoped>
.container {
  background-color: var(--foreground);
  padding: var(--padding-base);
  border-radius: var(--rounding-base);
}
</style>
