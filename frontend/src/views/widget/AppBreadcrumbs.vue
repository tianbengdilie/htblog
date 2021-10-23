<template>
  <v-breadcrumbs class="hidden-sm-and-down" :items="levelList" divider=" ">
    <template v-slot:item="props">
      <a @click.prevent="handleLink(props.item.path)">
        {{ generateTitle(props.item.title) }}
      </a>
    </template>
  </v-breadcrumbs>
</template>

<script>
export default {
  name: "AppBreadcrumbs",
  data: () => ({
    levelList: [],
    rootPath: {
      title: "/", //
      path: "/"
    }
  }),
  watch: {
    $route() {
      this.getBreadcrumb();
    }
  },
  created() {
    this.getBreadcrumb();
  },
  methods: {
    generateTitle(title) {
      const hasKey = this.$te(`${title}`);
      if (hasKey) {
        const translatedTitle = this.$t(`${title}`);
        return translatedTitle;
      }
      return title;
    },
    getBreadcrumb() {
      let nodeList = this.$route.fullPath
        .split("/")
        .filter(e => e && e.length > 0);
      if (nodeList.length < 2) return;
      if (nodeList[0] !== "blog") return;

      let basePath = "";
      this.levelList = nodeList.map(e => {
        basePath = basePath + "/" + e;
        return {
          title: e,
          path: basePath
        };
      });

      this.levelList = this.levelList.slice(1);
    },
    handleLink(path) {
      if (this.$route.fullPath === path) return;
      this.$router.push(path);
    }
  }
};
</script>
