<script setup lang="ts">
import Divider from "@/components/ui/Divider.vue";
import { getQueueNameById } from "@/utils/league-utils";
defineProps<{
    matchOutcome: string;
    timeSinceMatch: string;
    matchDuration: string;
    queueId: number;
}>();
</script>

<template>
    <div class="match__details">
        <p :class="['match__outcome', `match__outcome--${matchOutcome}`]">
            {{ matchOutcome }}
        </p>
        <p class="match__time-since">{{ timeSinceMatch }}</p>
        <Divider />
        <p class="match__queue-type">{{ getQueueNameById(queueId) }}</p>
        <p class="match__duration">{{ matchDuration }}</p>
    </div>
</template>

<style scoped>
.match__details {
    display: none;
}

.divider {
    display: none;
}

.match__details--mobile .match__info {
    display: flex;
}

.match__details--mobile {
    display: flex;
    align-items: center;
    justify-content: space-between;
    color: #ffffff67;
    margin-bottom: var(--margin-base);
    font-size: 0.75rem;
}

.match__outcome {
    margin-left: var(--margin-base);
    text-transform: capitalize;
}

.match__outcome--victory {
    color: var(--col-accent);
}

.match__outcome--defeat {
    color: var(--col-defeat);
}

.match__queue-type {
    width: 10ch;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

@media screen and (min-width: 1024px) {
    .match__details {
        display: block;
    }

    .divider {
        display: block;
    }

    .match__details--mobile {
        display: none;
    }

    .match__outcome {
        margin: 0;
        font-size: 1rem;
    }
}
</style>
