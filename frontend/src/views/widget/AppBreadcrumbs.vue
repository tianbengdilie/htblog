<template>
  <v-breadcrumbs class="hidden-sm-and-down" :items="levelList" divider="/">
    <template v-slot:item="props">
      <a @click.prevent="handleLink(props.item)">
        {{ generateTitle(props.item.title) }}
      </a>
    </template>
  </v-breadcrumbs>
</template>

<script>
const pathToRegexp = require("path");

export default {
  name: "AppBreadcrumbs",
  data: () => ({
    levelList: null
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
      // console.log("matched: ", this.$route.matched);
      // let matched = this.$route.matched.filter(item => item.name);

      // // const first = matched[0];
      // // if (first && first.name.trim().toLocaleLowerCase() !== "dashboard") {
      // //   matched = [
      // //     { path: "/dashboard", meta: { title: "route.dashboard" } }
      // //   ].concat(matched);
      // // }

      // this.levelList = matched.map(item => {
      //   return {
      //     text: item.name
      //   };
      // });
      // console.log("levelList: ", this.levelList);

      console.log(this.$route.fullPath);
      let nodeList = this.$route.fullPath
        .split("/")
        .filter(e => e && e.length > 0);
      console.log(nodeList);
      if (nodeList.length < 2) return;
      if (nodeList[0] !== "blog") return;

      let basePath = nodeList[0];
      this.levelList = nodeList.slice(1).map(e => {
        basePath = basePath + "/" + e;
        return {
          title: e,
          path: basePath
        };
      });

      console.log(this.$route);
      console.log(this.levelList);
      // this.levelList = [
      //   {
      //     meta: { title: "Dashboard" },
      //     disabled: false,
      //     href: "breadcrumbs_dashboard"
      //   },
      //   {
      //     meta: { title: "Link 1" },
      //     disabled: false,
      //     href: "breadcrumbs_link_1"
      //   },
      //   {
      //     meta: { title: "Link 2" },
      //     disabled: true,
      //     href: "breadcrumbs_link_2"
      //   }
      // ];
    },
    pathCompile(path) {
      const { params } = this.$route;
      const toPath = pathToRegexp.compile(path);
      return toPath(params);
    },
    handleLink(item) {
      console.groupCollapsed("handleLink");
      console.log(item);
      const { redirect, path } = item;
      console.log(`redirect=${redirect}, path=${path}`);
      if (redirect) {
        console.log("redirect");
        this.$router.push(redirect);
        console.groupEnd();
        return;
      }
      this.$router.push(path);
      console.groupEnd();
    }
  }
};
</script>
