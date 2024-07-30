<script setup>
import HelloWorld from "./components/HelloWorld.vue";

const links = ["Dashboard", "Messages", "Profile", "Updates"];
</script>
<!--<HelloWorld msg="You did it!" />-->
<template>
  <v-app id="inspire">
    <v-app-bar class="px-3" flat density="compact">
      <v-avatar
        color="grey-darken-1"
        class="hidden-md-and-up"
        size="32"
      ></v-avatar>

      <v-spacer></v-spacer>

      <v-tabs centered color="grey-darken-2">
        <v-tab v-for="link in links" :key="link" :text="link"></v-tab>
      </v-tabs>
      <v-spacer></v-spacer>

      <v-avatar
        class="hidden-sm-and-down"
        color="grey-darken-1"
        size="32"
      ></v-avatar>
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col offset-sm="2" cols="12" md="6">
            <v-sheet min-height="70vh" rounded="lg">
              <!--  -->

              <v-form v-model="valid">
                <v-container>
                  <v-row>
                    <v-col cols="12" md="12">
                      <v-text-field
                        v-model="firstname"
                        :rules="nameRules"
                        :counter="10"
                        label="First name"
                        required
                        hide-details
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" md="12">
                      <v-text-field
                        v-model="firstname"
                        :rules="nameRules"
                        :counter="10"
                        label="First name"
                        required
                        hide-details
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
              </v-form>
            </v-sheet>
          </v-col>

          <v-col cols="12" md="2">
            <v-sheet rounded="lg" min-height="268">
              <!--  -->
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    links: ["Dashboard", "Messages", "Profile", "Updates"],
  }),
};
</script>
