import { createRouter, createWebHistory } from "vue-router";

import HomeView from "../views/HomeView.vue";
import SummonerView from "../views/SummonerView.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "home",
            meta: {
                title: "Nashor"
            },
            component: HomeView,
        },
        {
            path: "/summoner/:region/:gameName/:tagLine",
            name: "summoner",
            props: true,
            component: SummonerView,
        },
    ],
});

router.beforeEach((to, _, next) => {
    if (to.name === "summoner") {
        const title = `Nashor - ${to.params.gameName}#${to.params.tagLine}`;

        document.title = title;
        next();
    }
})

export default router;
