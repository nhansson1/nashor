import { createRouter, createWebHistory } from "vue-router";

import HomeView from "../views/HomeView.vue";
import SummonerView from "../views/SummonerView.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "home",
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

export default router;
