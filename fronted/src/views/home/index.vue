<template>
  <v-container class="grow d-flex flex-column flex-nowrap">
    <v-row class="shrink">
      <v-col cols="12">
        <v-card elevation="2" class="pa-3"> Row 1 </v-card>
      </v-col>
    </v-row>
    <v-row class="grow">
      <v-col cols="12">
        <v-card elevation="2" class="pa-3 fill-height"> Row 2 </v-card>
      </v-col>
    </v-row>
    <v-row class="shrink">
      <v-col cols="12">
        <v-card elevation="2" class="pa-3"> Row 3 </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "homepage",
  components: {}
};
</script>

<style lang="less" scoped>
#main-view {
  height: 100%;
}
</style>
