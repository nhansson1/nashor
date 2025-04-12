<script setup lang="ts"></script>

<template>
    <footer class="footer">
        <div class="footer__main">
            <p class="footer__legal">
                Nashor.gg is not endorsed by Riot Games and does not reflect the views or opinions of Riot Games
                or anyone officially involved in producing or managing Riot Games properties. Riot Games and all
                associated
                properties are trademarks or registered trademarks of Riot Games, Inc
            </p>
        </div>
    </footer>
</template>

<style scoped>
.footer {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: var(--margin-base);
    background-color: var(--foreground);
}

.footer__legal {
    text-align: center;
    color: #ffffff7a;
}

@media screen and (min-width: 1024px) {
    .footer__main {
        width: 50vw;
    }
}
</style>
