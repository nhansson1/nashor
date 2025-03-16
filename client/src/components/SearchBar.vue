<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const regions = [
  {
    name: "euw",
    value: "euw1",
  },
  {
    name: "eune",
    value: "eun1",
  },
  {
    name: "br",
    value: "br1",
  },
  {
    name: "jp",
    value: "jp1",
  },
  {
    name: "kr",
    value: "kr1",
  },
  {
    name: "na",
    value: "na1",
  },
];

const searchModel = defineModel({ type: String });
const selectedRegion = ref(regions[0].value);

const handleChange = (e: Event) => {
  const target = e.target as HTMLSelectElement;
  selectedRegion.value = target.value;
};

const getProfile = async (gameName: string, tagLine: string): Promise<any> => {
  const url = new URL(
    `http://localhost:8080/api/v1/account/by-riot-id/${gameName}/${tagLine}`
  );

  try {
    const res = await fetch(url);

    if (!res.ok) {
      if (res.status === 404) {
        router.push("/404");
      }
      return;
    }

    const data = await res.json();

    return data;
  } catch (err) {
    console.log(err);
  }
};

const handleSubmit = async (e: Event) => {
  if (!searchModel.value) return;

  e.preventDefault();

  if (!searchModel.value.includes("#")) return;

  const [gameName, tagLine] = searchModel.value.split("#");

  const profile = await getProfile(gameName, tagLine);

  if (!profile) return;

  router.push({
    name: "summoner",
    params: { region: selectedRegion.value, gameName, tagLine },
  });
};
</script>

<!-- can take modifier classes: header -->
<template>
  <form @submit="handleSubmit" class="searchbar">
    <select
      @change="handleChange"
      name="region"
      id="region"
      class="searchbar__region"
    >
      <option v-for="region in regions" :value="region.value">
        {{ region.name }}
      </option>
    </select>
    <input
      v-model="searchModel"
      placeholder="Search a Summoner"
      type="text"
      class="searchbar__input"
    />
  </form>
</template>

<style scoped>
.searchbar {
  border-radius: var(--rounding-base);
  margin: var(--margin-base);
  background-color: var(--foreground);
  max-width: fit-content;
}

.searchbar__region {
  text-transform: uppercase;
  text-align: center;
}

.searchbar__input,
.searchbar__region {
  padding: var(--margin-large) var(--margin-small);
  background-color: inherit;
  color: white;
  outline: none;
  border: none;
  border-radius: var(--rounding-base);
  font-size: clamp(0.5rem, 3vw, 1.25rem);
}

.searchbar--header .searchbar__input,
.searchbar--header .searchbar__region {
  background-color: var(--background);
  padding: var(--margin-base) var(--margin-small);
  font-size: clamp(0.5rem, 3vw, 1rem);
}

@media screen and (min-width: 1024px) {
  .searchbar__input,
  .searchbar__region {
    padding: var(--margin-xl) var(--margin-base);
  }
}
</style>
