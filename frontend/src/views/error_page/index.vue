<template>
  <v-container class="container--fluid fill-height primary">
    <v-row no-gutters align="center" justify="center">
      <v-col cols="12" sm="8" md="4" lg="4">
        <v-card class="elevation-5 pa-3">
          <v-card-text>
            <div class="layout column align-center">
              <img src="img/icons/error-404.png" width="120" height="120" />
              <h1 class="text-center my-4 primary--text">Ht Website</h1>
              <h3 class="text-center my-4 primary--text">
                The requested URL {{ $route.params.pathMatch }} was not found on
                this server.
              </h3>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "error404",
  components: {}
};
</script>