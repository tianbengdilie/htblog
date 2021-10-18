<template>
  <v-app>
    <router-view />
  </v-app>
</template>

<script>
export default {
  name: "App",

  components: {},

  meta: {
    // if no subcomponents specify a metaInfo.title, this title will be used
    title: "Default Title",
    // all titles will be injected into this template
    titleTemplate: "%s | My Awesome Webapp"
  },
  data: () => ({
    drawer: false,
    group: null,
    website: {
      title: "HT的个人博客"
    },
    paths: ["home", "viewer"]
  }),
  methods: {},
  watch: {
    group() {
      this.drawer = false;
    }
  }
};
</script>
